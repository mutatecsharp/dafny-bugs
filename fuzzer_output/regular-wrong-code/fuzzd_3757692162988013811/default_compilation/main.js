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
      return new BigNumber(-290);
    };
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ln")).length), new BigNumber(589), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(873)), function (_0_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      })).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), function (_1_i1) {
        return new BigNumber(-916);
      })), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(2)), function (_2_i2) {
        return new BigNumber((function () {
          let _coll0 = new _dafny.Map();
          for (const _compr_0 of (_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))).Elements) {
            let _3_v0 = _compr_0;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0))), _3_v0)) {
              _coll0.push([_3_v0,new BigNumber(-708)]);
            }
          }
          return _coll0;
        }()).length);
      }), (_module.D5.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), function (_4_i3) {
  return new BigNumber((_dafny.Set.fromElements(true, true)).length);
}))).dtor_cf25));
    };
    static fm2(p0, globalState) {
      return ((_module.D7.create_DC20(_dafny.MultiSet.fromElements(new BigNumber((function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).Elements) {
    let _5_v0 = _compr_1;
    if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).contains(_5_v0)) {
      _coll1.push([_5_v0,true]);
    }
  }
  return _coll1;
}()).length)))).dtor_cf30).equals((_dafny.MultiSet.fromElements(new BigNumber(-675), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(429), new BigNumber(870))).cardinality()), new BigNumber(495))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-209)))));
    };
    static fm3(p0, globalState) {
      return _module.D0.create_DC2(new BigNumber((_dafny.Seq.of(new BigNumber(-501))).length), (_dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)))).IsSubsetOf(_dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)))));
    };
    static fm4(p0, p1, globalState) {
      return _module.D0.create_DC0(((true) ? (true) : (true)));
    };
    static fm13(p0, globalState) {
      return (_dafny.Set.fromElements(true)).Intersect((_dafny.Set.fromElements(true, false)).Intersect(_dafny.Set.fromElements(true, false, false, false)));
    };
    static fm15(p0, p1, p2, p3, globalState) {
      if (!((true) && (false))) {
        return (_module.D4.create_DC12(_dafny.Seq.of(false))).dtor_cf18;
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(!(!(!(!(!(!(!(false)))))))));
      }
    };
    static fm16(p0, globalState) {
      let _source0 = _module.D8.create_DC27(_module.D8.create_DC26(!(false)));
      if (_source0.is_DC24) {
        let _6___mcc_h0 = (_source0).cf36;
        let _7_cf36 = _6___mcc_h0;
        if (_7_cf36) {
          return _module.D2.create_DC5(new BigNumber(-268), _7_cf36, _7_cf36);
        } else {
          return _module.D2.create_DC5(new BigNumber(417), true, !(_7_cf36));
        }
      } else if (_source0.is_DC25) {
        let _8___mcc_h1 = (_source0).cf37;
        let _9___mcc_h2 = (_source0).cf38;
        let _10___mcc_h3 = (_source0).cf39;
        let _11___mcc_h4 = (_source0).cf40;
        let _12_cf40 = _11___mcc_h4;
        let _13_cf39 = _10___mcc_h3;
        let _14_cf38 = _9___mcc_h2;
        let _15_cf37 = _8___mcc_h1;
        return _module.D2.create_DC5(_14_cf38, _15_cf37, _15_cf37);
      } else if (_source0.is_DC26) {
        let _16___mcc_h5 = (_source0).cf41;
        let _17_cf41 = _16___mcc_h5;
        return _module.D2.create_DC5(new BigNumber(-893), _17_cf41, _17_cf41);
      } else if (_source0.is_DC23) {
        let _18___mcc_h6 = (_source0).cf35;
        let _19_cf35 = _18___mcc_h6;
        return _module.D2.create_DC5((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality())))).length)), true, true);
      } else {
        let _20___mcc_h7 = (_source0).cf42;
        let _21_cf42 = _20___mcc_h7;
        return _module.D2.create_DC5(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false)).length), (_module.D21.create_DC58(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),true)).length))).dtor_cf89, true);
      }
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(((!(false)) ? (_dafny.Seq.of(true)) : (_dafny.Seq.of(false))), _dafny.Seq.Concat(_dafny.Seq.of(true, !(true), false, !(true)), _dafny.Seq.of(!(!(false)))));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC0((function () {
  let _coll2 = new _dafny.Set();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(837), new BigNumber(-677))) {
    let _22_v0 = _compr_2;
    if (((new BigNumber(837)).isLessThanOrEqualTo(_22_v0)) && ((_22_v0).isLessThan(new BigNumber(-677)))) {
      _coll2.add(_module.__default.safeDivisionInt(_22_v0, new BigNumber(509)));
    }
  }
  return _coll2;
}()).IsProperSubsetOf(function () {
  let _coll3 = new _dafny.Set();
  for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-800), new BigNumber(437))) {
    let _23_v1 = _compr_3;
    if (((new BigNumber(-800)).isLessThanOrEqualTo(_23_v1)) && ((_23_v1).isLessThan(new BigNumber(437)))) {
      _coll3.add(_module.__default.safeDivisionInt(_23_v1, new BigNumber(642)));
    }
  }
  return _coll3;
}()));
    };
    static fm20(p0, p1, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).Union(_dafny.MultiSet.fromElements(false))).Difference((_dafny.MultiSet.fromElements(!(!(true)), true)).Intersect(_dafny.MultiSet.fromElements(false)));
    };
    static fm21(p0, globalState) {
      return _dafny.Seq.UnicodeFromString("mxqour");
    };
    static fm22(p0, globalState) {
      return (_dafny.MultiSet.fromElements(false, true)).Union(_dafny.MultiSet.fromElements(!(false), true));
    };
    static fm23(p0, p1, globalState) {
      return _module.D2.create_DC4(_dafny.Seq.of(_dafny.Seq.of((_module.D22.create_DC60(false, false, !(true))).dtor_cf93)));
    };
    static fm24(p0, p1, p2, globalState) {
      return _module.D2.create_DC5((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(218),new BigNumber((_dafny.Seq.UnicodeFromString("gikgwt")).length))).length)))), true, true);
    };
    static fm25(p0, globalState) {
      if (!(new BigNumber(815)).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(16)), function (_24_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length))) {
        return _module.D3.create_DC9(new BigNumber(33), new BigNumber((_dafny.Seq.UnicodeFromString("gkarebnxj")).length));
      } else {
        return _module.D3.create_DC9(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wwdkvrqt")).length), new BigNumber(992))).cardinality()), new BigNumber(731));
      }
    };
    static fm27(globalState) {
      return ((_dafny.Set.fromElements(false)).Union(_dafny.Set.fromElements(true, false))).Intersect((_dafny.Set.fromElements(true)));
    };
    static fm28(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(618), new BigNumber(-394))) {
          let _25_v0 = _compr_4;
          if (((new BigNumber(618)).isLessThanOrEqualTo(_25_v0)) && ((_25_v0).isLessThan(new BigNumber(-394)))) {
            _coll4.push([_module.__default.safeDivisionInt(_25_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC37(_dafny.Seq.UnicodeFromString("k"), true),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(930)))).length)),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("eskdxn")).length), new BigNumber(383), new BigNumber(492))]);
          }
        }
        return _coll4;
      }()).length), new BigNumber(-861)), _dafny.Seq.of(new BigNumber(789), new BigNumber(210))), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-831)))));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D23.create_DC64();
      if (_source1.is_DC64) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("txelk"), _dafny.Seq.UnicodeFromString("pj"));
      } else {
        let _26___mcc_h0 = (_source1).cf98;
        let _27_cf98 = _26___mcc_h0;
        return _dafny.Seq.UnicodeFromString("bnvlqggk");
      }
    };
    static fm30(p0, p1, globalState) {
      let _source2 = new _dafny.CodePoint('v'.codePointAt(0));
      let _28___mcc_h0 = _source2;
      let _29_cf44 = _28___mcc_h0;
      return _module.D6.create_DC19(_module.D6.create_DC18((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(867), new BigNumber(843))).length))));
    };
    static fm31(p0, globalState) {
      return _module.D0.create_DC2(new BigNumber(230), (new BigNumber((_dafny.Seq.UnicodeFromString("lpmdnlj")).length)).isLessThan(new BigNumber(-160)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(!(false))), _dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(true, !(false), !(true))));
    };
    static fm35(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(691)), function (_30_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      });
    };
    static fm36(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(false, (_module.D17.create_DC49(false)).dtor_cf73))).Difference(_dafny.Set.fromElements(_dafny.Seq.of(!(!(false))), _dafny.Seq.of(true, true, true), _dafny.Seq.of(true, false)))).Union(_dafny.Set.fromElements(_dafny.Seq.of(false, true, false, false, true), _dafny.Seq.of(!(true)), _dafny.Seq.of(true), _dafny.Seq.of(!(true), true, !(!(true)))));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC12(_dafny.Seq.of(true, true));
    };
    static fm39(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('h'.codePointAt(0));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), function (_31_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      });
    };
    static fm41(p0, p1, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.of(true));
    };
    static fm42(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_dafny.MultiSet.FromArray((_module.D5.create_DC15(_dafny.Seq.of(new BigNumber(-461), new BigNumber((_dafny.Seq.of(function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of (_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)))).Elements) {
    let _32_v0 = _compr_5;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0))), _32_v0)) {
      _coll5.push([_32_v0,new BigNumber(112)]);
    }
  }
  return _coll5;
}())).length), new BigNumber((_dafny.Seq.UnicodeFromString("pu")).length), new BigNumber(485)))).dtor_cf25))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(303)))));
    };
    static fm43(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, true), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(false, !(false))));
    };
    static fm44(p0, p1, globalState) {
      if (true) {
        return _module.D2.create_DC5((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)), false, false);
      } else {
        return _module.D2.create_DC5(new BigNumber(-672), false, false);
      }
    };
    static fm45(p0, p1, p2, globalState) {
      return (((true) ? (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(310))).length), new BigNumber(244))).cardinality()),new BigNumber(464))).length))) : (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-60)));
    };
    static fm46(p0, p1, p2, globalState) {
      if (!(!(true)) || (false)) {
        return _module.D6.create_DC18(new BigNumber(779));
      } else {
        return _module.D6.create_DC18(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true), _dafny.Seq.of(!(false), true, true), _dafny.Seq.of(true, false))).length));
      }
    };
    static fm47(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), function (_33_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }),_dafny.Seq.of(new BigNumber(562), new BigNumber(654), new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(30), new BigNumber(-859))) {
          let _34_v0 = _compr_6;
          if (((new BigNumber(30)).isLessThanOrEqualTo(_34_v0)) && ((_34_v0).isLessThan(new BigNumber(-859)))) {
            _coll6.push([(_34_v0).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-906), new BigNumber(953))).length))),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)]);
          }
        }
        return _coll6;
      }()).length)))).Merge(function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ryegaxv"), _dafny.Seq.UnicodeFromString("kpurpehd"), _dafny.Seq.UnicodeFromString("vrthtypa"))).Elements) {
          let _35_v1 = _compr_7;
          if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ryegaxv"), _dafny.Seq.UnicodeFromString("kpurpehd"), _dafny.Seq.UnicodeFromString("vrthtypa"))).contains(_35_v1)) {
            _coll7.push([_35_v1,_dafny.Seq.of(new BigNumber(854), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(998))).cardinality()))]);
          }
        }
        return _coll7;
      }());
    };
    static fm48(globalState) {
      return _module.D11.create_DC30(function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-656), new BigNumber(360))) {
    let _36_v0 = _compr_8;
    if (((new BigNumber(-656)).isLessThanOrEqualTo(_36_v0)) && ((_36_v0).isLessThan(new BigNumber(360)))) {
      _coll8.push([(_36_v0).plus(new BigNumber(-774)),_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).length))]);
    }
  }
  return _coll8;
}());
    };
    static fm49(p0, p1, p2, p3, globalState) {
      if (!(((false) ? (true) : (!(!(true)))))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(139),!(!(false)));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(377))).length),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(409),false));
      }
    };
    static fm50(p0, p1, globalState) {
      return _module.D4.create_DC14(true);
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(540), new BigNumber(589))) {
            let _37_v0 = _compr_10;
            if (((new BigNumber(540)).isLessThanOrEqualTo(_37_v0)) && ((_37_v0).isLessThan(new BigNumber(589)))) {
              _coll10.push([(_37_v0).multipliedBy(new BigNumber(394)),new _dafny.CodePoint('a'.codePointAt(0))]);
            }
          }
          return _coll10;
        }()).Keys.Elements) {
          let _38_v1 = _compr_9;
          if ((function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of _dafny.IntegerRange(new BigNumber(540), new BigNumber(589))) {
              let _37_v0 = _compr_11;
              if (((new BigNumber(540)).isLessThanOrEqualTo(_37_v0)) && ((_37_v0).isLessThan(new BigNumber(589)))) {
                _coll11.push([(_37_v0).multipliedBy(new BigNumber(394)),new _dafny.CodePoint('a'.codePointAt(0))]);
              }
            }
            return _coll11;
          }()).contains(_38_v1)) {
            _coll9.add(_module.__default.safeModuloInt(_38_v1, new BigNumber(130)));
          }
        }
        return _coll9;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(new BigNumber(717))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length))));
    };
    static fm52(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-114),new BigNumber(398))).length))).length),true), function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Seq.of(new BigNumber(617))).Elements) {
          let _39_v0 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(617)), _39_v0)) {
            _coll12.push([_module.__default.safeDivisionInt(_39_v0, new BigNumber(96)),false]);
          }
        }
        return _coll12;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),!(true)))).Union(function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(600), new BigNumber(102))) {
            let _40_v1 = _compr_14;
            if (((new BigNumber(600)).isLessThanOrEqualTo(_40_v1)) && ((_40_v1).isLessThan(new BigNumber(102)))) {
              _coll14.push([(_40_v1).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality())),new BigNumber(-401)]);
            }
          }
          return _coll14;
        }()).length))).cardinality())),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(443),true), function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of (_dafny.Seq.of(new BigNumber(503))).Elements) {
            let _41_v2 = _compr_15;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(503)), _41_v2)) {
              _coll15.push([(_41_v2).plus(new BigNumber(-371)),false]);
            }
          }
          return _coll15;
        }(), function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(915), new BigNumber(-501))) {
            let _42_v3 = _compr_16;
            if (((new BigNumber(915)).isLessThanOrEqualTo(_42_v3)) && ((_42_v3).isLessThan(new BigNumber(-501)))) {
              _coll16.push([_module.__default.safeModuloInt(_42_v3, new BigNumber((_dafny.Seq.UnicodeFromString("swbdrpg")).length)),false]);
            }
          }
          return _coll16;
        }())).Elements) {
          let _43_v4 = _compr_13;
          if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of _dafny.IntegerRange(new BigNumber(600), new BigNumber(102))) {
              let _40_v1 = _compr_17;
              if (((new BigNumber(600)).isLessThanOrEqualTo(_40_v1)) && ((_40_v1).isLessThan(new BigNumber(102)))) {
                _coll17.push([(_40_v1).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality())),new BigNumber(-401)]);
              }
            }
            return _coll17;
          }()).length))).cardinality())),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(443),true), function () {
            let _coll18 = new _dafny.Map();
            for (const _compr_18 of (_dafny.Seq.of(new BigNumber(503))).Elements) {
              let _41_v2 = _compr_18;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(503)), _41_v2)) {
                _coll18.push([(_41_v2).plus(new BigNumber(-371)),false]);
              }
            }
            return _coll18;
          }(), function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(915), new BigNumber(-501))) {
              let _42_v3 = _compr_19;
              if (((new BigNumber(915)).isLessThanOrEqualTo(_42_v3)) && ((_42_v3).isLessThan(new BigNumber(-501)))) {
                _coll19.push([_module.__default.safeModuloInt(_42_v3, new BigNumber((_dafny.Seq.UnicodeFromString("swbdrpg")).length)),false]);
              }
            }
            return _coll19;
          }())).contains(_43_v4)) {
            _coll13.add(_43_v4);
          }
        }
        return _coll13;
      }())).Difference(_dafny.Set.fromElements(function () {
        let _coll20 = new _dafny.Map();
        for (const _compr_20 of (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)))).length))).Elements) {
          let _44_v5 = _compr_20;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)))).length))).contains(_44_v5)) {
            _coll20.push([(_44_v5).multipliedBy(new BigNumber(512)),!(!(false))]);
          }
        }
        return _coll20;
      }()));
    };
    static fm53(p0, p1, globalState) {
      return _module.D2.create_DC4(_dafny.Seq.of(_dafny.Seq.of(true)));
    };
    static fm54(p0, globalState) {
      if (!(!(((true) ? (true) : (!(true)))))) {
        return _module.D3.create_DC8((_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(!(true),false))).dtor_cf11);
      } else {
        return _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),true));
      }
    };
    static fm55(globalState) {
      return _module.D4.create_DC13(!(true), (new BigNumber(-318)).isEqualTo(new BigNumber((_dafny.Set.fromElements(true, !(!(!(true))), false, !(!(false)), true)).length)), new BigNumber((function () {
  let _coll21 = new _dafny.Set();
  for (const _compr_21 of (_dafny.Set.fromElements(_module.D14.create_DC41(_dafny.Seq.UnicodeFromString("osq"), new BigNumber(762)), _module.D14.create_DC41(_dafny.Seq.UnicodeFromString("vnwyed"), new BigNumber((_dafny.Seq.UnicodeFromString("skj")).length)))).Elements) {
    let _45_v0 = _compr_21;
    if ((_dafny.Set.fromElements(_module.D14.create_DC41(_dafny.Seq.UnicodeFromString("osq"), new BigNumber(762)), _module.D14.create_DC41(_dafny.Seq.UnicodeFromString("vnwyed"), new BigNumber((_dafny.Seq.UnicodeFromString("skj")).length)))).contains(_45_v0)) {
      _coll21.add(_45_v0);
    }
  }
  return _coll21;
}()).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nnji"), _dafny.Seq.UnicodeFromString("jvs"))).length), !(!(true)));
    };
    static fm56(p0, globalState) {
      if (false) {
        return _dafny.MultiSet.fromElements(new BigNumber(364), new BigNumber((_dafny.Seq.UnicodeFromString("hpnljp")).length));
      } else {
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(738)));
      }
    };
    static fm57(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of _dafny.IntegerRange(new BigNumber(228), new BigNumber(940))) {
          let _46_v0 = _compr_22;
          if (((new BigNumber(228)).isLessThanOrEqualTo(_46_v0)) && ((_46_v0).isLessThan(new BigNumber(940)))) {
            _coll22.push([(_46_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("huket")).length)),new BigNumber(757)]);
          }
        }
        return _coll22;
      }(), (function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of _dafny.IntegerRange(new BigNumber(846), new BigNumber(847))) {
          let _47_v1 = _compr_23;
          if (((new BigNumber(846)).isLessThanOrEqualTo(_47_v1)) && ((_47_v1).isLessThan(new BigNumber(847)))) {
            _coll23.push([_module.__default.safeDivisionInt(_47_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(635))).length)),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, true)).length)))]);
          }
        }
        return _coll23;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-619),new BigNumber(-169))));
    };
    static fm58(p0, globalState) {
      return _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(650)), function (_48_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }));
    };
    static fm59(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,(_module.D2.create_DC5(new BigNumber((_dafny.Seq.UnicodeFromString("ulwymqieg")).length), true, true)).dtor_cf7)));
    };
    static fm60(p0, globalState) {
      let _source3 = _module.D12.create_DC35(false);
      if (_source3.is_DC35) {
        let _49___mcc_h0 = (_source3).cf54;
        let _50_cf54 = _49___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_50_cf54,_dafny.Map.Empty.slice().updateUnsafe(_50_cf54,_50_cf54));
      } else if (_source3.is_DC36) {
        let _51___mcc_h1 = (_source3).cf55;
        let _52___mcc_h2 = (_source3).cf56;
        let _53___mcc_h3 = (_source3).cf57;
        let _54___mcc_h4 = (_source3).cf58;
        let _55_cf58 = _54___mcc_h4;
        let _56_cf57 = _53___mcc_h3;
        let _57_cf56 = _52___mcc_h2;
        let _58_cf55 = _51___mcc_h1;
        return (_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Map.Empty.slice().updateUnsafe(true,!(false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(true,true)));
      } else if (_source3.is_DC37) {
        let _59___mcc_h5 = (_source3).cf59;
        let _60___mcc_h6 = (_source3).cf60;
        let _61_cf60 = _60___mcc_h6;
        let _62_cf59 = _59___mcc_h5;
        return _dafny.Map.Empty.slice().updateUnsafe(!(_61_cf60),_dafny.Map.Empty.slice().updateUnsafe(_61_cf60,_61_cf60));
      } else if (_source3.is_DC34) {
        let _63___mcc_h7 = (_source3).cf53;
        let _64_cf53 = _63___mcc_h7;
        return _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(true,!(false)));
      } else {
        let _65___mcc_h8 = (_source3).cf61;
        let _66_cf61 = _65___mcc_h8;
        return _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(true,true));
      }
    };
    static fm61(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,true),_dafny.Seq.UnicodeFromString("gsgynkkcj"))).Merge(function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of ((_module.D27.create_DC73(function () {
  let _coll25 = new _dafny.Set();
  for (const _compr_25 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).Elements) {
    let _67_v1 = _compr_25;
    if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).contains(_67_v1)) {
      _coll25.add(_67_v1);
    }
  }
  return _coll25;
}())).dtor_cf113).Elements) {
          let _68_v0 = _compr_24;
          if (((_module.D27.create_DC73(function () {
  let _coll26 = new _dafny.Set();
  for (const _compr_26 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).Elements) {
    let _69_v1 = _compr_26;
    if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).contains(_69_v1)) {
      _coll26.add(_69_v1);
    }
  }
  return _coll26;
}())).dtor_cf113).contains(_68_v0)) {
            _coll24.push([_68_v0,_dafny.Seq.UnicodeFromString("ax")]);
          }
        }
        return _coll24;
      }());
    };
    static fm62(p0, globalState) {
      return function () {
        let _coll27 = new _dafny.Map();
        for (const _compr_27 of _dafny.IntegerRange(new BigNumber(348), new BigNumber(296))) {
          let _70_v0 = _compr_27;
          if (((new BigNumber(348)).isLessThanOrEqualTo(_70_v0)) && ((_70_v0).isLessThan(new BigNumber(296)))) {
            _coll27.push([_module.__default.safeModuloInt(_70_v0, new BigNumber(490)),_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(57))]);
          }
        }
        return _coll27;
      }();
    };
    static fm63(p0, globalState) {
      return _module.D0.create_DC1();
    };
    static fm64(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true))).Difference(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, false), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(false, true, true)));
    };
    static fm65(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),new _dafny.CodePoint('p'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new _dafny.CodePoint('h'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),new _dafny.CodePoint('u'.codePointAt(0))));
    };
    static fm66(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(54),_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(false, true, true)).length), new BigNumber((_dafny.Seq.of(new BigNumber(911), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("bybpefmp")).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length)), new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0)))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(63), new BigNumber(840))).cardinality()))).cardinality()))).length)));
    };
    static fm67(p0, p1, p2, globalState) {
      return _module.D8.create_DC24(false);
    };
    static fm68(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),true)).Merge(function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).Elements) {
          let _71_v0 = _compr_28;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0))), _71_v0)) {
            _coll28.push([_71_v0,false]);
          }
        }
        return _coll28;
      }())).Merge((function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), function (_72_i0) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).Elements) {
          let _73_v1 = _compr_29;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(47)), function (_72_i0) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }), _73_v1)) {
            _coll29.push([_73_v1,false]);
          }
        }
        return _coll29;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),false)));
    };
    static fm69(p0, p1, p2, globalState) {
      if ((function () {
        let _coll30 = new _dafny.Set();
        for (const _compr_30 of _dafny.IntegerRange(new BigNumber(-13), new BigNumber(636))) {
          let _74_v0 = _compr_30;
          if (((new BigNumber(-13)).isLessThanOrEqualTo(_74_v0)) && ((_74_v0).isLessThan(new BigNumber(636)))) {
            _coll30.add((_74_v0).minus(new BigNumber(722)));
          }
        }
        return _coll30;
      }()).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(144)), function (_75_i0) {
        return new BigNumber(-237);
      })).length)))) {
        return _module.D20.create_DC54((new _dafny.CodePoint('f'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-555),false), _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(578)), function (_76_i1) {
  return new _dafny.CodePoint('j'.codePointAt(0));
})), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("yjsq")).length)));
      } else {
        return _module.D20.create_DC54(new _dafny.CodePoint('x'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-840),false), _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("aquqvkpq")), new BigNumber(471));
      }
    };
    static fm70(p0, p1, p2, p3, globalState) {
      return (((!(true)) ? (_module.D6.create_DC17(_dafny.Set.fromElements(new BigNumber(199), new BigNumber((_dafny.Set.fromElements(false, true)).length), new BigNumber(937), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(225)), function (_77_i0) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length), new BigNumber(754)))) : (_module.D6.create_DC17(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length))).cardinality())))))).dtor_cf27;
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _78_v0;
      let _init0 = function (_79_i0) {
        return _module.__default.safeModuloInt(_79_i0, new BigNumber((_dafny.Seq.of(new BigNumber(515), new BigNumber(-90))).length));
      };
      let _nw0 = Array((_dafny.ONE).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _78_v0 = _nw0;
      let _80_v1;
      _80_v1 = _dafny.MultiSet.fromElements(p0, p0, p2);
      let _81_v2;
      _81_v2 = new BigNumber(788);
      let _index0 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length));
      (_78_v0)[_index0] = (((_80_v1).contains(p0)) ? ((_80_v1).get(p0)) : (_81_v2));
      let _82_v3;
      let _nw1 = new _module.C7();
      _nw1.__ctor();
      _82_v3 = _nw1;
      let _83_v4;
      _83_v4 = _dafny.Seq.of(_82_v3);
      let _84_v5;
      _84_v5 = _dafny.Map.Empty.slice().updateUnsafe(_83_v4,_dafny.Seq.UnicodeFromString("cn"));
      let _85_v6;
      _85_v6 = _dafny.Seq.UnicodeFromString("kvva");
      let _index1 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length));
      let _rhs0 = new BigNumber(286);
      let _rhs1 = (_84_v5).Merge(_84_v5);
      let _rhs2 = _85_v6;
      let _lhs0 = _78_v0;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length));
      _lhs0[_lhs1] = _rhs0;
      _84_v5 = _rhs1;
      _85_v6 = _rhs2;
      r2 = (_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))];
      let _86_v7;
      _86_v7 = ((p2) ? (_80_v1) : (_80_v1));
      let _87_v8;
      _87_v8 = _dafny.Set.fromElements(_81_v2);
      let _88_v9;
      _88_v9 = _module.D3.create_DC9(_81_v2, _81_v2);
      let _89_v10;
      _89_v10 = _dafny.Map.Empty.slice().updateUnsafe(p2,_81_v2);
      let _90_v11;
      _90_v11 = new _dafny.CodePoint('s'.codePointAt(0));
      let _rhs3 = (new BigNumber((_87_v8).length)).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,(_88_v9).dtor_cf13)).Merge(_89_v10)).length));
      let _rhs4 = false;
      let _rhs5 = _86_v7;
      let _rhs6 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_85_v6, _module.__default.safeIndex(_81_v2, new BigNumber((_85_v6).length)), _90_v11), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gtswhfxei"), _85_v6));
      let _rhs7 = (_87_v8).Union(_module.__default.fm70(p0, _dafny.Seq.UnicodeFromString("qecc"), _81_v2, p2, globalState));
      let _lhs2 = globalState;
      r2 = _rhs3;
      _lhs2.f4 = _rhs4;
      _86_v7 = _rhs5;
      r1 = _rhs6;
      _87_v8 = _rhs7;
      let _91_i1;
      _91_i1 = _dafny.ZERO;
      L0: {
        while (p2) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_91_i1)) {
              break L0;
            }
            _91_i1 = (_91_i1).plus(_dafny.ONE);
            let _92_v12;
            let _init1 = ((_93_v2, _94_v0) => function (_95_i2) {
              return (_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(_93_v2), _module.__default.safeIndex((_94_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_94_v0).length))], new BigNumber((_dafny.Seq.of(_93_v2)).length)), (_94_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_94_v0).length))]))).IsDisjointFrom(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_93_v2), new BigNumber((_dafny.MultiSet.fromElements(_93_v2, (_94_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_94_v0).length))])).cardinality())));
            })(_81_v2, _78_v0);
            let _nw2 = Array((new BigNumber(10)).toNumber());
            for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw2.length); _i0_1++) {
              _nw2[_i0_1] = _init1(new BigNumber(_i0_1));
            }
            _92_v12 = _nw2;
            let _index2 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_92_v12).length));
            (_92_v12)[_index2] = p2;
            let _96_v13;
            let _nw3 = new _module.C4();
            _nw3.__ctor();
            _96_v13 = _nw3;
            let _97_v14;
            _97_v14 = _dafny.Map.Empty.slice().updateUnsafe(_96_v13,_87_v8);
            let _98_v15;
            _98_v15 = _dafny.Set.fromElements(_87_v8, (((_97_v14).contains(_96_v13)) ? ((_97_v14).get(_96_v13)) : (_87_v8)));
            let _99_v16;
            _99_v16 = _dafny.Seq.of(p2, p2, false);
            let _100_v17;
            _100_v17 = _dafny.Map.Empty.slice().updateUnsafe((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))],_89_v10);
            let _101_v18;
            _101_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(497),_dafny.Set.fromElements(_dafny.Set.fromElements((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))], _81_v2, _module.__default.fm0((_99_v16)[_module.__default.safeIndex(_81_v2, new BigNumber((_99_v16).length))], _100_v17, p2, globalState), (_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))])));
            let _index3 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_92_v12).length));
            (_92_v12)[_index3] = !(((_98_v15).Difference(_dafny.Set.fromElements(_87_v8, _87_v8, _dafny.Set.fromElements(_81_v2), _87_v8, _87_v8))).equals((((_101_v18).contains(new BigNumber(809))) ? ((_101_v18).get(new BigNumber(809))) : (_98_v15))));
            let _102_v19;
            _102_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_92_v12)[_module.__default.safeIndex(new BigNumber(530), new BigNumber((_92_v12).length))]);
            let _103_v20;
            _103_v20 = _dafny.Map.Empty.slice().updateUnsafe(p2,_102_v19);
            let _104_v21;
            _104_v21 = _dafny.MultiSet.fromElements(_81_v2, _81_v2, new BigNumber(213), (_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
            let _105_v22;
            _105_v22 = _dafny.Seq.of(_104_v21);
            let _106_v23;
            _106_v23 = _dafny.Set.fromElements(p2);
            let _107_v24;
            let _nw4 = Array((new BigNumber(28)).toNumber());
            _nw4[(_dafny.ZERO).toNumber()] = _104_v21;
            _nw4[(_dafny.ONE).toNumber()] = _104_v21;
            _nw4[(new BigNumber(2)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(3)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(4)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(5)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(6)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(7)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(8)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
            _nw4[(new BigNumber(10)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(11)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(12)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(13)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(14)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
            _nw4[(new BigNumber(16)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(17)).toNumber()] = (_105_v22)[_module.__default.safeIndex(new BigNumber(523), new BigNumber((_105_v22).length))];
            _nw4[(new BigNumber(18)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements(_81_v2, new BigNumber((_106_v23).length));
            _nw4[(new BigNumber(20)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(21)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(22)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(23)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(24)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(25)).toNumber()] = _104_v21;
            _nw4[(new BigNumber(26)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_81_v2));
            _nw4[(new BigNumber(27)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(926)));
            _107_v24 = _nw4;
            let _108_v25;
            let _nw5 = new _module.C2();
            _nw5.__ctor(_dafny.Seq.of(_81_v2), p2, _103_v20, _99_v16, _107_v24);
            _108_v25 = _nw5;
            let _109_v26;
            _109_v26 = _dafny.Map.Empty.slice().updateUnsafe(_108_v25,_104_v21);
            let _110_v27;
            _110_v27 = _dafny.Seq.of((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
            let _111_v28;
            let _nw6 = Array((new BigNumber(19)).toNumber());
            _nw6[(_dafny.ZERO).toNumber()] = _104_v21;
            _nw6[(_dafny.ONE).toNumber()] = _104_v21;
            _nw6[(new BigNumber(2)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(3)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(4)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(5)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(6)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(7)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_112_i3) {
              return new BigNumber(157);
            }));
            _nw6[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber(602));
            _nw6[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
            _nw6[(new BigNumber(11)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_81_v2);
            _nw6[(new BigNumber(13)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))], _81_v2, new BigNumber(873), _81_v2);
            _nw6[(new BigNumber(15)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(16)).toNumber()] = _104_v21;
            _nw6[(new BigNumber(17)).toNumber()] = (((_109_v26).contains(_108_v25)) ? ((_109_v26).get(_108_v25)) : (_104_v21));
            _nw6[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.FromArray(_110_v27);
            _111_v28 = _nw6;
            let _113_v29;
            let _nw7 = new _module.C1();
            _nw7.__ctor(p0, true, (_103_v20).Merge(_103_v20), _dafny.Seq.Concat(_99_v16, _99_v16), _107_v24);
            _113_v29 = _nw7;
            _113_v29 = _113_v29;
            let _index4 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length));
            (_78_v0)[_index4] = _81_v2;
            let _114_v30;
            let _init2 = ((_115_v6) => function (_116_i4) {
              return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("j"), _115_v6);
            })(_85_v6);
            let _nw8 = Array((new BigNumber(16)).toNumber());
            for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw8.length); _i0_2++) {
              _nw8[_i0_2] = _init2(new BigNumber(_i0_2));
            }
            _114_v30 = _nw8;
            let _117_v31;
            let _nw9 = new _module.C2();
            _nw9.__ctor(_110_v27, (_87_v8).IsDisjointFrom(_87_v8), (((_113_v29).f17) ? (_103_v20) : (_103_v20)), _99_v16, _111_v28);
            _117_v31 = _nw9;
            let _rhs8 = _114_v30;
            let _rhs9 = _81_v2;
            let _rhs10 = _117_v31;
            let _rhs11 = (_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))];
            let _rhs12 = _92_v12;
            _114_v30 = _rhs8;
            _81_v2 = _rhs9;
            _117_v31 = _rhs10;
            _81_v2 = _rhs11;
            _92_v12 = _rhs12;
          }
        }
      }
      let _118_v32;
      _118_v32 = _dafny.Seq.of((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
      let _119_v33;
      _119_v33 = _dafny.Seq.of(p2);
      let _120_v34;
      _120_v34 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
      let _121_v35;
      _121_v35 = _dafny.Map.Empty.slice().updateUnsafe((_119_v33)[_module.__default.safeIndex(new BigNumber((_120_v34).length), new BigNumber((_119_v33).length))],p0);
      let _122_v36;
      _122_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_121_v35);
      let _123_v37;
      _123_v37 = _dafny.MultiSet.fromElements((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
      let _124_v38;
      _124_v38 = _dafny.Map.Empty.slice().updateUnsafe(false,_123_v37);
      let _125_v39;
      _125_v39 = _dafny.Map.Empty.slice().updateUnsafe(_119_v33,p2);
      let _126_v40;
      let _nw10 = Array((new BigNumber(19)).toNumber()).fill(_dafny.MultiSet.Empty);
      _126_v40 = _nw10;
      let _127_v41;
      let _nw11 = new _module.C5();
      _nw11.__ctor(p0, _126_v40, _122_v36, _119_v33);
      _127_v41 = _nw11;
      let _128_v42;
      _128_v42 = _dafny.Map.Empty.slice().updateUnsafe(p2,_127_v41);
      let _129_v43;
      _129_v43 = _dafny.MultiSet.fromElements(_128_v42);
      let _130_v44;
      _130_v44 = _module.D14.create_DC41(_85_v6, _81_v2);
      let _131_v45;
      _131_v45 = _dafny.Seq.of(_130_v44);
      let _132_v46;
      _132_v46 = _module.D26.create_DC70(_131_v45);
      let _133_v47;
      let _nw12 = Array((new BigNumber(25)).toNumber());
      _nw12[(_dafny.ZERO).toNumber()] = _123_v37;
      _nw12[(_dafny.ONE).toNumber()] = _123_v37;
      _nw12[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.FromArray(_118_v32);
      _nw12[(new BigNumber(3)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(4)).toNumber()] = (((_124_v38).contains((((_125_v39).contains(_119_v33)) ? ((_125_v39).get(_119_v33)) : (!(p0))))) ? ((_124_v38).get((((_125_v39).contains(_119_v33)) ? ((_125_v39).get(_119_v33)) : (!(p0))))) : (_123_v37));
      _nw12[(new BigNumber(5)).toNumber()] = (((_124_v38).contains(p2)) ? ((_124_v38).get(p2)) : (_123_v37));
      _nw12[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements((((_129_v43).contains(_128_v42)) ? ((_129_v43).get(_128_v42)) : ((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))])), new BigNumber((_85_v6).length));
      _nw12[(new BigNumber(7)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(8)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(9)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(10)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(11)).toNumber()] = (((_124_v38).contains(true)) ? ((_124_v38).get(true)) : (_dafny.MultiSet.fromElements(_81_v2, new BigNumber((_123_v37).cardinality()), new BigNumber((_119_v33).length))));
      _nw12[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_81_v2);
      _nw12[(new BigNumber(13)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(14)).toNumber()] = (_dafny.MultiSet.FromArray(_118_v32)).update(new BigNumber(274), _module.__default.abs(new BigNumber(((_132_v46).dtor_cf106).length)));
      _nw12[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]);
      _nw12[(new BigNumber(16)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(17)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(18)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(19)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(20)).toNumber()] = _module.__default.fm28(_85_v6, p0, _81_v2, globalState);
      _nw12[(new BigNumber(21)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(22)).toNumber()] = _123_v37;
      _nw12[(new BigNumber(23)).toNumber()] = (_module.__default.fm56(_90_v11, globalState)).update(_81_v2, _module.__default.abs((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))]));
      _nw12[(new BigNumber(24)).toNumber()] = _dafny.MultiSet.fromElements(_81_v2);
      _133_v47 = _nw12;
      let _134_v48;
      let _nw13 = new _module.C2();
      _nw13.__ctor(_dafny.Seq.Concat(_118_v32, _118_v32), !(p0), _122_v36, _119_v33, _133_v47);
      _134_v48 = _nw13;
      if (!(p0)) {
        let _135_v49;
        let _136_v50;
        let _137_v51;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_82_v3).m2((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))], globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _135_v49 = _out0;
        _136_v50 = _out1;
        _137_v51 = _out2;
        let _138_v52;
        let _nw14 = Array((new BigNumber(21)).toNumber()).fill(_dafny.MultiSet.Empty);
        _138_v52 = _nw14;
        let _index5 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_138_v52).length));
        (_138_v52)[_index5] = _80_v1;
        let _index6 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_138_v52).length));
        (_138_v52)[_index6] = _80_v1;
        let _139_v53;
        _139_v53 = _dafny.Map.Empty.slice().updateUnsafe(_85_v6,_78_v0);
        let _140_v54;
        _140_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(123),_78_v0);
        let _141_v55;
        _141_v55 = _dafny.Seq.of((((_140_v54).contains(_136_v50)) ? ((_140_v54).get(_136_v50)) : (_78_v0)), _78_v0);
        _139_v53 = (_139_v53).update(_dafny.Seq.Concat(_85_v6, _85_v6), (_141_v55)[_module.__default.safeIndex((_78_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_78_v0).length))], new BigNumber((_141_v55).length))]);
        let _142_v56;
        let _nw15 = Array((new BigNumber(8)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = (_127_v41).f19;
        _nw15[(_dafny.ONE).toNumber()] = p0;
        _nw15[(new BigNumber(2)).toNumber()] = (_127_v41).f19;
        _nw15[(new BigNumber(3)).toNumber()] = (_127_v41).f19;
        _nw15[(new BigNumber(4)).toNumber()] = (_127_v41).f19;
        _nw15[(new BigNumber(5)).toNumber()] = p2;
        _nw15[(new BigNumber(6)).toNumber()] = (_127_v41).f19;
        _nw15[(new BigNumber(7)).toNumber()] = p0;
        _142_v56 = _nw15;
        let _143_v57;
        _143_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-756),p2);
        let _index7 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_142_v56).length));
        (_142_v56)[_index7] = !(_143_v57).contains(_135_v49);
        let _index8 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_142_v56).length));
        (_142_v56)[_index8] = p0;
        let _144_v58;
        let _nw16 = Array((new BigNumber(11)).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
        _nw16[(_dafny.ONE).toNumber()] = _90_v11;
        _nw16[(new BigNumber(2)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(3)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(4)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(5)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(6)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(7)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(8)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(9)).toNumber()] = _90_v11;
        _nw16[(new BigNumber(10)).toNumber()] = _90_v11;
        _144_v58 = _nw16;
        let _145_v59;
        _145_v59 = _dafny.Seq.of(_144_v58);
        _143_v57 = (_143_v57).update(new BigNumber((_dafny.Seq.Concat(_145_v59, _145_v59)).length), (_127_v41).f19);
      } else {
        let _146_v60;
        let _init3 = ((_147_v37) => function (_148_i5) {
          return (_147_v37).IsProperSubsetOf(_147_v37);
        })(_123_v37);
        let _nw17 = Array((new BigNumber(24)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw17.length); _i0_3++) {
          _nw17[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _146_v60 = _nw17;
        let _149_v61;
        let _nw18 = new _module.C4();
        _nw18.__ctor();
        _149_v61 = _nw18;
        let _150_v62;
        _150_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_85_v6).length),_dafny.Set.fromElements(_149_v61, _149_v61));
        let _151_v63;
        _151_v63 = _dafny.Set.fromElements(_149_v61);
        let _152_v64;
        _152_v64 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_149_v61),_90_v11);
        let _index9 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_146_v60).length));
        (_146_v60)[_index9] = ((_dafny.Map.Empty.slice().updateUnsafe((((_150_v62).contains(new BigNumber((_123_v37).cardinality()))) ? ((_150_v62).get(new BigNumber((_123_v37).cardinality()))) : (_151_v63)),_90_v11)).update(_151_v63, _90_v11)).equals(_152_v64);
        let _index10 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_146_v60).length));
        (_146_v60)[_index10] = p2;
        let _153_v65;
        let _nw19 = Array((new BigNumber(15)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw19[(_dafny.ONE).toNumber()] = _90_v11;
        _nw19[(new BigNumber(2)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(3)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(4)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(5)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(6)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(7)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(8)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(9)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(10)).toNumber()] = (_85_v6)[_module.__default.safeIndex(_81_v2, new BigNumber((_85_v6).length))];
        _nw19[(new BigNumber(11)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(12)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(13)).toNumber()] = _90_v11;
        _nw19[(new BigNumber(14)).toNumber()] = _90_v11;
        _153_v65 = _nw19;
        let _154_v66;
        _154_v66 = _dafny.Map.Empty.slice().updateUnsafe(_146_v60,(((_146_v60)[_module.__default.safeIndex(new BigNumber(640), new BigNumber((_146_v60).length))]) ? (_153_v65) : (_153_v65)));
        _154_v66 = (_154_v66).update(_146_v60, _153_v65);
        _85_v6 = _85_v6;
        _89_v10 = (_89_v10).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_81_v2));
        r3 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("fpbqponc"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(359)), ((_155_v11) => function (_156_i6) {
          return _155_v11;
        })(_90_v11)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-418)), ((_157_v11) => function (_158_i7) {
          return _157_v11;
        })(_90_v11))));
      }
      r0 = _dafny.Seq.Concat(_118_v32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(75)), ((_159_v0) => function (_160_i8) {
        return (_159_v0)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_159_v0).length))];
      })(_78_v0)));
      r1 = p2;
      let _161_v67;
      _161_v67 = _module.D8.create_DC25((_127_v41).f19, new BigNumber((_119_v33).length), _81_v2, _89_v10);
      r2 = _module.__default.safeModuloInt((_81_v2).minus(new BigNumber((_module.__default.fm35(new BigNumber((_80_v1).cardinality()), _161_v67, globalState)).length)), (_134_v48).fm8((_127_v41).f19, p2, globalState));
      r3 = p0;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _162_v0;
      let _nw20 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _162_v0 = _nw20;
      let _163_v1;
      _163_v1 = false;
      let _164_v2;
      _164_v2 = _dafny.MultiSet.fromElements(_163_v1);
      let _165_v3;
      _165_v3 = new BigNumber(846);
      let _166_v4;
      _166_v4 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_165_v3);
      let _167_globalState;
      let _nw21 = new _module.GlobalState();
      _nw21.__ctor(false, _162_v0, new BigNumber(-689), new BigNumber(665), false, _164_v2, new BigNumber(409), new BigNumber(-44), false, false, _166_v4, false);
      _167_globalState = _nw21;
      let _hi0 = _module.__default.fm0(_163_v1, function () {
        let _coll31 = new _dafny.Map();
        for (const _compr_31 of _dafny.IntegerRange(new BigNumber(481), new BigNumber(-667))) {
          let _168_v5 = _compr_31;
          if (((new BigNumber(481)).isLessThanOrEqualTo(_168_v5)) && ((_168_v5).isLessThan(new BigNumber(-667)))) {
            _coll31.push([(_168_v5).plus(new BigNumber((_dafny.Seq.of(_163_v1)).length)),_dafny.Map.Empty.slice().updateUnsafe(_163_v1,_165_v3)]);
          }
        }
        return _coll31;
      }(), _163_v1, _167_globalState);
      for (let _169_i0 = _165_v3; _169_i0.isLessThan(_hi0); _169_i0 = _169_i0.plus(_dafny.ONE)) {
        let _170_v7;
        let _init4 = ((_171_v3) => function (_172_i1) {
          return function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of _dafny.IntegerRange(new BigNumber(-584), new BigNumber(-15))) {
              let _173_v6 = _compr_32;
              if (((new BigNumber(-584)).isLessThanOrEqualTo(_173_v6)) && ((_173_v6).isLessThan(new BigNumber(-15)))) {
                _coll32.add((_173_v6).plus((_dafny.ZERO).minus(_171_v3)));
              }
            }
            return _coll32;
          }();
        })(_165_v3);
        let _nw22 = Array((new BigNumber(7)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw22.length); _i0_4++) {
          _nw22[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _170_v7 = _nw22;
        _170_v7 = _170_v7;
        _163_v1 = _163_v1;
        let _174_v8;
        _174_v8 = new _dafny.CodePoint('k'.codePointAt(0));
        let _175_v9;
        _175_v9 = _dafny.Map.Empty.slice().updateUnsafe(_174_v8,new _dafny.CodePoint('o'.codePointAt(0)));
        let _176_v10;
        _176_v10 = _dafny.Seq.of(new BigNumber((_175_v9).length));
        let _177_v11;
        _177_v11 = _module.D0.create_DC2(new BigNumber((_164_v2).cardinality()), _163_v1);
        let _178_v12;
        _178_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_176_v10, _module.__default.fm1((_177_v11).dtor_cf2, _163_v1, _167_globalState)),_163_v1);
        _178_v12 = (_178_v12).update(false, _module.__default.fm2(_163_v1, _167_globalState));
        if (_163_v1) {
          _163_v1 = _163_v1;
          let _179_v13;
          _179_v13 = _dafny.Seq.of(_163_v1, _163_v1);
          let _180_v14;
          _180_v14 = _dafny.Map.Empty.slice().updateUnsafe(_176_v10,_163_v1);
          let _181_v15;
          _181_v15 = _dafny.Set.fromElements(new BigNumber((_180_v14).length));
          _165_v3 = (_module.__default.safeDivisionInt(new BigNumber((_179_v13).length), new BigNumber((_181_v15).length))).multipliedBy((new BigNumber(890)).plus(new BigNumber(701)));
          (_167_globalState).f4 = _163_v1;
          _165_v3 = _165_v3;
          let _182_v16;
          _182_v16 = _module.D0.create_DC1();
          _182_v16 = _182_v16;
        } else {
          let _183_v17;
          let _184_v18;
          let _185_v19;
          let _186_v20;
          let _out3;
          let _out4;
          let _out5;
          let _out6;
          let _outcollector1 = _module.__default.m0(((_163_v1) ? (_163_v1) : (_163_v1)), _177_v11, (_169_i0).isLessThanOrEqualTo(_169_i0), _167_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _out6 = _outcollector1[3];
          _183_v17 = _out3;
          _184_v18 = _out4;
          _185_v19 = _out5;
          _186_v20 = _out6;
          let _187_v21;
          _187_v21 = _164_v2;
          _184_v18 = ((_187_v21)).IsSubsetOf(_164_v2);
          let _188_v22;
          _188_v22 = _dafny.Seq.UnicodeFromString("lomvp");
          let _189_v23;
          let _nw23 = Array((new BigNumber(23)).toNumber()).fill(_dafny.MultiSet.Empty);
          _189_v23 = _nw23;
          let _index11 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_189_v23).length));
          (_189_v23)[_index11] = ((_184_v18) ? (_164_v2) : (_164_v2));
          let _index12 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_189_v23).length));
          let _rhs13 = !(_186_v20) || (!(_module.__default.fm2(_186_v20, _167_globalState)));
          let _rhs14 = _module.D0.create_DC2(new BigNumber((_176_v10).length), _184_v18);
          let _rhs15 = _188_v22;
          let _rhs16 = ((_163_v1) ? (_dafny.MultiSet.fromElements(_186_v20, false, _163_v1)) : (((_184_v18) ? (_dafny.MultiSet.fromElements(true, _163_v1)) : (_164_v2))));
          let _lhs3 = _189_v23;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_189_v23).length));
          _184_v18 = _rhs13;
          _177_v11 = _rhs14;
          _188_v22 = _rhs15;
          _lhs3[_lhs4] = _rhs16;
          let _190_v24;
          _190_v24 = _dafny.Map.Empty.slice().updateUnsafe(_188_v22,_163_v1);
          (_167_globalState).f4 = (((_190_v24).contains(_188_v22)) ? ((_190_v24).get(_188_v22)) : (_184_v18));
          let _191_v25;
          _191_v25 = _dafny.Set.fromElements(_163_v1);
          _191_v25 = _191_v25;
        }
      }
      let _hi1 = _165_v3;
      for (let _192_i2 = _165_v3; _192_i2.isLessThan(_hi1); _192_i2 = _192_i2.plus(_dafny.ONE)) {
        let _193_v26;
        _193_v26 = _dafny.Set.fromElements(false, _163_v1, _163_v1, _163_v1, _module.__default.fm2(_163_v1, _167_globalState));
        let _194_v27;
        _194_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-77),_163_v1);
        let _195_v28;
        _195_v28 = _dafny.Seq.of(_163_v1, _163_v1, _163_v1, _163_v1);
        let _196_v29;
        _196_v29 = _dafny.Seq.of(_163_v1, (((_194_v27).contains((_dafny.ZERO).minus(new BigNumber((_195_v28).length)))) ? ((_194_v27).get((_dafny.ZERO).minus(new BigNumber((_195_v28).length)))) : (_163_v1)), _163_v1, _163_v1, !(_163_v1));
        let _197_v30;
        _197_v30 = _dafny.Map.Empty.slice().updateUnsafe((_193_v26).Intersect(_193_v26),_dafny.areEqual(_195_v28, _dafny.Seq.of(_163_v1)));
        _197_v30 = (_197_v30).update(_193_v26, _163_v1);
        let _198_v31;
        let _init5 = ((_199_i2) => function (_200_i3) {
          return _dafny.MultiSet.fromElements(_199_i2, new BigNumber((_dafny.Seq.UnicodeFromString("tywvlxp")).length));
        })(_192_i2);
        let _nw24 = Array((new BigNumber(20)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw24.length); _i0_5++) {
          _nw24[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _198_v31 = _nw24;
        let _201_v32;
        _201_v32 = _dafny.Map.Empty.slice().updateUnsafe(_193_v26,new BigNumber(305));
        let _202_v33;
        _202_v33 = _dafny.MultiSet.fromElements(_165_v3, new BigNumber((_201_v32).length));
        let _203_v34;
        _203_v34 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_module.__default.fm2(_163_v1, _167_globalState));
        let _204_v35;
        _204_v35 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_203_v34);
        let _205_v36;
        let _nw25 = new _module.C3();
        _nw25.__ctor(_module.__default.fm2(_163_v1, _167_globalState), _198_v31, ((((_202_v33).contains((_dafny.ZERO).minus(_192_i2))) ? ((_202_v33).get((_dafny.ZERO).minus(_192_i2))) : (_192_i2))).isLessThan(_192_i2), _204_v35, _dafny.Seq.of(true, _163_v1, _163_v1));
        _205_v36 = _nw25;
        _163_v1 = ((_163_v1) ? (false) : (!((((_203_v34).contains((_205_v36).f20)) ? ((_203_v34).get((_205_v36).f20)) : ((_195_v28)[_module.__default.safeIndex(_192_i2, new BigNumber((_195_v28).length))])))));
        let _206_v37;
        _206_v37 = _dafny.Seq.of(_192_i2, _192_i2);
        let _207_v38;
        let _nw26 = new _module.C2();
        _nw26.__ctor(_206_v37, true, _204_v35, _195_v28, _198_v31);
        _207_v38 = _nw26;
        let _208_v39;
        _208_v39 = _dafny.Map.Empty.slice().updateUnsafe(_207_v38,(_205_v36).f20);
        _163_v1 = (((_205_v36).f20) ? ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_192_i2), new BigNumber(881))).IsProperSubsetOf(_202_v33)) : ((((_208_v39).contains(_207_v38)) ? ((_208_v39).get(_207_v38)) : (_163_v1))));
      }
      let _209_v40;
      _209_v40 = _dafny.Set.fromElements(_163_v1, _163_v1);
      let _hi2 = _165_v3;
      for (let _210_i4 = new BigNumber((_209_v40).length); _210_i4.isLessThan(_hi2); _210_i4 = _210_i4.plus(_dafny.ONE)) {
        let _211_v41;
        let _nw27 = Array((new BigNumber(5)).toNumber()).fill([]);
        _211_v41 = _nw27;
        let _212_v42;
        let _nw28 = Array((_dafny.ONE).toNumber());
        _nw28[(_dafny.ZERO).toNumber()] = _module.__default.fm2(false, _167_globalState);
        _212_v42 = _nw28;
        let _index13 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_211_v41).length));
        (_211_v41)[_index13] = _212_v42;
        let _index14 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_211_v41).length));
        (_211_v41)[_index14] = _212_v42;
        if (_163_v1) {
          let _213_v43;
          _213_v43 = _dafny.Seq.of(_163_v1, _163_v1, _163_v1);
          (_167_globalState).f4 = (((_213_v43)[_module.__default.safeIndex(new BigNumber(233), new BigNumber((_213_v43).length))]) ? (_163_v1) : (_163_v1));
          (_167_globalState).f4 = _163_v1;
          let _214_v44;
          _214_v44 = _dafny.Seq.UnicodeFromString("nwe");
          let _215_v45;
          _215_v45 = new _dafny.CodePoint('n'.codePointAt(0));
          let _216_v46;
          _216_v46 = _dafny.Seq.of(_165_v3, _210_i4);
          let _217_v47;
          _217_v47 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_163_v1);
          let _218_v48;
          _218_v48 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_217_v47);
          let _219_v49;
          let _nw29 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _219_v49 = _nw29;
          let _220_v50;
          let _nw30 = new _module.C6();
          _nw30.__ctor(_216_v46, _163_v1, _218_v48, _213_v43, _219_v49);
          _220_v50 = _nw30;
          let _221_v51;
          _221_v51 = _dafny.Seq.of(_220_v50);
          let _222_v52;
          _222_v52 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.update(_214_v44, _module.__default.safeIndex(_165_v3, new BigNumber((_214_v44).length)), _215_v45)).length)).multipliedBy(new BigNumber((_221_v51).length)),_215_v45);
          _222_v52 = (_222_v52).update(_165_v3, _215_v45);
          let _223_v53;
          _223_v53 = _dafny.Map.Empty.slice().updateUnsafe(_210_i4,_214_v44);
          let _224_v54;
          let _out7;
          _out7 = (_220_v50).m5(_209_v40, _dafny.Seq.Concat(_214_v44, (((_223_v53).contains(new BigNumber(761))) ? ((_223_v53).get(new BigNumber(761))) : (_214_v44))), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(718), new BigNumber((_214_v44).length))), _167_globalState);
          _224_v54 = _out7;
          (_167_globalState).f4 = (_210_i4).isLessThanOrEqualTo(_210_i4);
        } else {
          let _225_v55;
          _225_v55 = _dafny.Seq.of(_165_v3, _210_i4);
          let _226_v56;
          _226_v56 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,new BigNumber((_225_v55).length));
          let _227_v57;
          _227_v57 = _dafny.Seq.of(_module.__default.fm0(false, _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_226_v56), _163_v1, _167_globalState));
          _227_v57 = _dafny.Seq.Concat(_225_v55, _225_v55);
          let _228_v58;
          let _init6 = ((_229_v1) => function (_230_i5) {
            return _module.D4.create_DC12(_dafny.Seq.of(_229_v1));
          })(_163_v1);
          let _nw31 = Array((new BigNumber(13)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw31.length); _i0_6++) {
            _nw31[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _228_v58 = _nw31;
          let _index15 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_162_v0).length));
          (_162_v0)[_index15] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_210_i4)), _165_v3);
          let _index16 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_162_v0).length));
          (_162_v0)[_index16] = (_165_v3).plus(_165_v3);
          let _231_v59;
          _231_v59 = _dafny.Seq.of(_228_v58, _228_v58, _228_v58);
          let _232_v60;
          _232_v60 = _dafny.Map.Empty.slice().updateUnsafe(_210_i4,_163_v1);
          let _233_v61;
          _233_v61 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_226_v56);
          let _index17 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_162_v0).length));
          let _index18 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_162_v0).length));
          let _rhs17 = (_231_v59)[_module.__default.safeIndex(new BigNumber((((_163_v1) ? (_232_v60) : (_232_v60))).length), new BigNumber((_231_v59).length))];
          let _rhs18 = _module.__default.fm0(true, _233_v61, (((((_232_v60).contains(_210_i4)) ? ((_232_v60).get(_210_i4)) : (_163_v1))) ? (true) : (_163_v1)), _167_globalState);
          let _rhs19 = (_210_i4).minus(_module.__default.safeModuloInt(_165_v3, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm68(_163_v1, _167_globalState)).length))));
          let _rhs20 = (_210_i4).minus(_165_v3);
          let _lhs5 = _162_v0;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_162_v0).length));
          let _lhs7 = _162_v0;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_162_v0).length));
          _228_v58 = _rhs17;
          _lhs5[_lhs6] = _rhs18;
          _165_v3 = _rhs19;
          _lhs7[_lhs8] = _rhs20;
          (_167_globalState).f4 = _163_v1;
          let _234_v62;
          _234_v62 = _dafny.Seq.of(_212_v42);
          let _235_v63;
          _235_v63 = _module.D21.create_DC55(_234_v62);
          let _236_v64;
          _236_v64 = _module.D0.create_DC2(new BigNumber(((_235_v63).dtor_cf84).length), true);
          let _237_v65;
          _237_v65 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_163_v1);
          let _238_v66;
          let _239_v67;
          let _240_v68;
          let _241_v69;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0(_163_v1, _236_v64, (((_237_v65).contains(true)) ? ((_237_v65).get(true)) : (_163_v1)), _167_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _238_v66 = _out8;
          _239_v67 = _out9;
          _240_v68 = _out10;
          _241_v69 = _out11;
          let _242_v70;
          _242_v70 = _dafny.Map.Empty.slice().updateUnsafe(false,_237_v65);
          let _243_v71;
          _243_v71 = _dafny.Seq.of(_239_v67);
          let _244_v72;
          let _init7 = ((_245_v55) => function (_246_i6) {
            return _dafny.MultiSet.FromArray(_245_v55);
          })(_225_v55);
          let _nw32 = Array((new BigNumber(9)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw32.length); _i0_7++) {
            _nw32[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _244_v72 = _nw32;
          let _247_v73;
          let _nw33 = new _module.C1();
          _nw33.__ctor(_163_v1, _163_v1, _242_v70, _243_v71, _244_v72);
          _247_v73 = _nw33;
          let _248_v74;
          let _nw34 = new _module.C5();
          _nw34.__ctor(false, _247_v73.f12, (_247_v73).f13, (_247_v73).f14);
          _248_v74 = _nw34;
          let _249_v75;
          _249_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_165_v3,_248_v74)).update((_162_v0)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_162_v0).length))], _248_v74)).length),_212_v42);
          let _250_v76;
          _250_v76 = _dafny.Map.Empty.slice().updateUnsafe(_247_v73,(((_249_v75).contains((_dafny.ZERO).minus(_210_i4))) ? ((_249_v75).get((_dafny.ZERO).minus(_210_i4))) : ((_211_v41)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_211_v41).length))])));
          let _251_v77;
          _251_v77 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_247_v73);
          _250_v76 = (_250_v76).update((((_251_v77).contains((_248_v74).f19)) ? ((_251_v77).get((_248_v74).f19)) : (_247_v73)), (_211_v41)[_module.__default.safeIndex(new BigNumber(552), new BigNumber((_211_v41).length))]);
        }
        let _252_v78;
        _252_v78 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,false);
        let _253_v79;
        _253_v79 = _dafny.Map.Empty.slice().updateUnsafe(false,_252_v78);
        let _254_v80;
        let _init8 = ((_255_i4) => function (_256_i7) {
          return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mkxripbcw")).length), _255_i4);
        })(_210_i4);
        let _nw35 = Array((new BigNumber(23)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw35.length); _i0_8++) {
          _nw35[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _254_v80 = _nw35;
        let _257_v81;
        let _nw36 = new _module.C1();
        _nw36.__ctor(!(false), _163_v1, _253_v79, _dafny.Seq.of(_163_v1), _254_v80);
        _257_v81 = _nw36;
        let _258_v82;
        _258_v82 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_257_v81);
        let _259_v83;
        _259_v83 = _dafny.Map.Empty.slice().updateUnsafe(_258_v82,_163_v1);
        let _260_v84;
        _260_v84 = _dafny.Map.Empty.slice().updateUnsafe(false,_210_i4);
        let _261_v85;
        _261_v85 = _dafny.Map.Empty.slice().updateUnsafe(_210_i4,_260_v84);
        _259_v83 = (_259_v83).update((_258_v82).update(_165_v3, _257_v81), (_module.__default.fm0(_163_v1, _261_v85, _163_v1, _167_globalState)).isLessThanOrEqualTo(new BigNumber((_260_v84).length)));
        let _index19 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_211_v41).length));
        (_211_v41)[_index19] = _212_v42;
      }
      (_167_globalState).f4 = (_164_v2).IsSubsetOf(_164_v2);
      if (_163_v1) {
        let _index20 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_162_v0).length));
        (_162_v0)[_index20] = _165_v3;
        let _262_v86;
        _262_v86 = new _dafny.CodePoint('i'.codePointAt(0));
        let _263_v87;
        _263_v87 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_262_v86);
        let _index21 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_162_v0).length));
        (_162_v0)[_index21] = new BigNumber((_263_v87).length);
        let _264_v88;
        _264_v88 = _dafny.Seq.of(_163_v1);
        let _265_v89;
        _265_v89 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_163_v1, _167_globalState),new BigNumber((_264_v88).length)));
        let _266_v90;
        _266_v90 = _dafny.MultiSet.fromElements(_module.D11.create_DC30(_265_v89));
        let _267_v91;
        _267_v91 = _module.D4.create_DC13(_163_v1, _163_v1, _165_v3, _165_v3, (_264_v88)[_module.__default.safeIndex(new BigNumber((_266_v90).cardinality()), new BigNumber((_264_v88).length))]);
        let _268_v92;
        _268_v92 = _dafny.Seq.UnicodeFromString("b");
        let _269_v93;
        _269_v93 = _module.D4.create_DC14(_163_v1);
        let _270_v94;
        _270_v94 = _module.D12.create_DC37(_dafny.Seq.UnicodeFromString("de"), _163_v1);
        let _271_v95;
        _271_v95 = _dafny.MultiSet.fromElements(_module.__default.fm29((_267_v91).dtor_cf23, _163_v1, _163_v1, new BigNumber((_268_v92).length), _167_globalState), _dafny.Seq.update(_dafny.Seq.Concat(_268_v92, _module.__default.fm40(_269_v93, _165_v3, _163_v1, !(_163_v1), _167_globalState)), _module.__default.safeIndex((_dafny.ZERO).minus(_165_v3), new BigNumber((_dafny.Seq.Concat(_268_v92, _module.__default.fm40(_269_v93, _165_v3, _163_v1, !(_163_v1), _167_globalState))).length)), _module.__default.fm39(new BigNumber(-338), _163_v1, _165_v3, (_270_v94).dtor_cf59, _167_globalState)));
        let _index22 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_162_v0).length));
        let _index23 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_162_v0).length));
        let _rhs21 = _165_v3;
        let _rhs22 = (((_271_v95).contains(_268_v92)) ? ((_271_v95).get(_268_v92)) : (_165_v3));
        let _rhs23 = _module.__default.safeModuloInt(new BigNumber((_164_v2).cardinality()), (_165_v3).plus(new BigNumber((_264_v88).length)));
        let _rhs24 = _module.__default.safeModuloInt(new BigNumber(393), new BigNumber((_dafny.Seq.Concat(_264_v88, _264_v88)).length));
        let _lhs9 = _162_v0;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_162_v0).length));
        let _lhs11 = _162_v0;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_162_v0).length));
        _lhs9[_lhs10] = _rhs21;
        _lhs11[_lhs12] = _rhs22;
        _165_v3 = _rhs23;
        _165_v3 = _rhs24;
        let _272_v96;
        _272_v96 = _dafny.Map.Empty.slice().updateUnsafe(!(_163_v1),_165_v3);
        let _273_v97;
        _273_v97 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_163_v1);
        let _274_v98;
        _274_v98 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_273_v97);
        let _275_v99;
        let _nw37 = Array((new BigNumber(28)).toNumber()).fill(_dafny.MultiSet.Empty);
        _275_v99 = _nw37;
        let _276_v100;
        _276_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(348),_275_v99);
        let _277_v101;
        let _nw38 = new _module.C6();
        _nw38.__ctor(_dafny.Seq.update(_dafny.Seq.of(_165_v3, (_dafny.ZERO).minus(_165_v3)), _module.__default.safeIndex(_module.__default.fm0(true, _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_272_v96), _163_v1, _167_globalState), new BigNumber((_dafny.Seq.of(_165_v3, (_dafny.ZERO).minus(_165_v3))).length)), new BigNumber((_268_v92).length)), _163_v1, (_274_v98).Merge(_dafny.Map.Empty.slice().updateUnsafe(_163_v1,_273_v97)), _module.__default.fm15(_163_v1, _165_v3, new BigNumber((_268_v92).length), new BigNumber(-254), _167_globalState), (((_276_v100).contains((_162_v0)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_162_v0).length))])) ? ((_276_v100).get((_162_v0)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_162_v0).length))])) : (_275_v99)));
        _277_v101 = _nw38;
        let _278_v102;
        _278_v102 = _dafny.Seq.of(_165_v3);
        _278_v102 = _278_v102;
        _262_v86 = _262_v86;
        let _279_v103;
        _279_v103 = _dafny.MultiSet.fromElements(new BigNumber(237), _165_v3);
        let _280_v104;
        let _out12;
        _out12 = (_277_v101).m4(_165_v3, (new BigNumber(594)).plus(_165_v3), _268_v92, ((_279_v103).update((_dafny.ZERO).minus((_162_v0)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_162_v0).length))]), _module.__default.abs(new BigNumber((_166_v4).length)))).Intersect(_279_v103), _167_globalState);
        _280_v104 = _out12;
      } else {
        let _281_v105;
        _281_v105 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_165_v3,_163_v1)).length));
        _281_v105 = _281_v105;
        let _282_v106;
        _282_v106 = _dafny.Seq.UnicodeFromString("isu");
        let _283_v107;
        _283_v107 = _dafny.MultiSet.fromElements(new BigNumber(631));
        let _284_v108;
        _284_v108 = _module.D7.create_DC20(_283_v107);
        let _285_v109;
        _285_v109 = _module.D11.create_DC32(new BigNumber(780), _165_v3);
        _165_v3 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber((_282_v106).length), new BigNumber(((_284_v108).dtor_cf30).cardinality())), (_285_v109).dtor_cf47);
        let _286_v110;
        let _nw39 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _286_v110 = _nw39;
        let _287_v111;
        _287_v111 = new _dafny.CodePoint('g'.codePointAt(0));
        let _index24 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_286_v110).length));
        (_286_v110)[_index24] = _287_v111;
        let _index25 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_286_v110).length));
        (_286_v110)[_index25] = ((_163_v1) ? (_287_v111) : (_287_v111));
        let _288_v112;
        let _nw40 = new _module.C7();
        _nw40.__ctor();
        _288_v112 = _nw40;
        let _289_v113;
        let _init9 = ((_290_v107) => function (_291_i8) {
          return _290_v107;
        })(_283_v107);
        let _nw41 = Array((new BigNumber(7)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw41.length); _i0_9++) {
          _nw41[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _289_v113 = _nw41;
        let _292_v114;
        _292_v114 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,true);
        let _293_v115;
        _293_v115 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_292_v114);
        let _294_v116;
        _294_v116 = _dafny.Seq.of(_163_v1);
        let _295_v117;
        let _nw42 = new _module.C5();
        _nw42.__ctor(_163_v1, _289_v113, _293_v115, _294_v116);
        _295_v117 = _nw42;
        let _296_v118;
        _296_v118 = _dafny.Map.Empty.slice().updateUnsafe(_295_v117,_288_v112);
        let _297_v119;
        let _nw43 = Array((new BigNumber(17)).toNumber());
        _nw43[(_dafny.ZERO).toNumber()] = _288_v112;
        _nw43[(_dafny.ONE).toNumber()] = _288_v112;
        _nw43[(new BigNumber(2)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(3)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(4)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(5)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(6)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(7)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(8)).toNumber()] = (((_296_v118).contains(_295_v117)) ? ((_296_v118).get(_295_v117)) : (_288_v112));
        _nw43[(new BigNumber(9)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(10)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(11)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(12)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(13)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(14)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(15)).toNumber()] = _288_v112;
        _nw43[(new BigNumber(16)).toNumber()] = _288_v112;
        _297_v119 = _nw43;
        let _298_v120;
        _298_v120 = _dafny.Set.fromElements(_297_v119, _297_v119, _297_v119, _297_v119, _297_v119);
        _298_v120 = _dafny.Set.fromElements(_297_v119);
        let _299_v121;
        let _nw44 = Array((new BigNumber(11)).toNumber());
        _299_v121 = _nw44;
        let _300_v122;
        let _nw45 = new _module.C3();
        _nw45.__ctor(_163_v1, _295_v117.f12, _163_v1, (_295_v117).f13, _dafny.Seq.of(_163_v1, _163_v1, _163_v1));
        _300_v122 = _nw45;
        let _index26 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_299_v121).length));
        (_299_v121)[_index26] = _300_v122;
        let _index27 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_299_v121).length));
        (_299_v121)[_index27] = _300_v122;
      }
      let _301_i9;
      _301_i9 = _dafny.ZERO;
      L1: {
        while (_163_v1) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_301_i9)) {
              break L1;
            }
            _301_i9 = (_301_i9).plus(_dafny.ONE);
            let _302_v123;
            _302_v123 = new _dafny.CodePoint('n'.codePointAt(0));
            let _303_v124;
            _303_v124 = _dafny.Map.Empty.slice().updateUnsafe(_302_v123,_165_v3);
            let _304_v125;
            _304_v125 = _dafny.Seq.of(new BigNumber(42));
            let _305_v126;
            _305_v126 = _dafny.Map.Empty.slice().updateUnsafe(_303_v124,_304_v125);
            _165_v3 = new BigNumber((_305_v126).length);
            _165_v3 = (_165_v3).minus(_165_v3);
            let _306_v127;
            let _init10 = ((_307_v125) => function (_308_i10) {
              return _dafny.MultiSet.FromArray(_307_v125);
            })(_304_v125);
            let _nw46 = Array((new BigNumber(12)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw46.length); _i0_10++) {
              _nw46[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _306_v127 = _nw46;
            let _309_v128;
            let _nw47 = new _module.C5();
            _nw47.__ctor(_163_v1, _306_v127, _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(_163_v1,!(_163_v1))), _dafny.Seq.of(!(!(_163_v1)), _163_v1));
            _309_v128 = _nw47;
            let _310_v129;
            let _nw48 = Array((new BigNumber(21)).toNumber());
            _nw48[(_dafny.ZERO).toNumber()] = _309_v128;
            _nw48[(_dafny.ONE).toNumber()] = _309_v128;
            _nw48[(new BigNumber(2)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(3)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(4)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(5)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(6)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(7)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(8)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(9)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(10)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(11)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(12)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(13)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(14)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(15)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(16)).toNumber()] = ((!((_309_v128).f19)) ? (_309_v128) : (_309_v128));
            _nw48[(new BigNumber(17)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(18)).toNumber()] = _309_v128;
            _nw48[(new BigNumber(19)).toNumber()] = (_module.D22.create_DC59(_309_v128)).dtor_cf91;
            _nw48[(new BigNumber(20)).toNumber()] = _309_v128;
            _310_v129 = _nw48;
            let _nw49 = Array((new BigNumber(6)).toNumber());
            _nw49[(_dafny.ZERO).toNumber()] = _309_v128;
            _nw49[(_dafny.ONE).toNumber()] = _309_v128;
            _nw49[(new BigNumber(2)).toNumber()] = _309_v128;
            _nw49[(new BigNumber(3)).toNumber()] = _309_v128;
            _nw49[(new BigNumber(4)).toNumber()] = _309_v128;
            _nw49[(new BigNumber(5)).toNumber()] = (((_309_v128).f19) ? (_309_v128) : (_309_v128));
            _310_v129 = _nw49;
            let _311_v130;
            _311_v130 = _dafny.Seq.UnicodeFromString("nqmjqhts");
            let _312_v131;
            _312_v131 = _dafny.Seq.of(_311_v130, _dafny.Seq.UnicodeFromString("th"));
            let _313_v132;
            _313_v132 = _dafny.Seq.of(!((_309_v128).f19), (_dafny.MultiSet.FromArray(_312_v131)).IsSubsetOf(_dafny.MultiSet.fromElements(_311_v130)));
            let _314_v133;
            _314_v133 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_163_v1);
            let _315_v134;
            _315_v134 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,(_314_v133).update((_309_v128).f19, (_309_v128).f19));
            let _316_v135;
            let _nw50 = new _module.C2();
            _nw50.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(485)), ((_317_v3) => function (_318_i11) {
              return _317_v3;
            })(_165_v3)), (_309_v128).f19, _315_v134, _313_v132, _306_v127);
            _316_v135 = _nw50;
            let _319_v136;
            _319_v136 = _dafny.Map.Empty.slice().updateUnsafe(_162_v0,_316_v135);
            let _320_v137;
            _320_v137 = _module.D23.create_DC63(_316_v135);
            let _321_v138;
            let _nw51 = Array((new BigNumber(22)).toNumber());
            _nw51[(_dafny.ZERO).toNumber()] = _316_v135;
            _nw51[(_dafny.ONE).toNumber()] = (((_319_v136).contains(_162_v0)) ? ((_319_v136).get(_162_v0)) : (_316_v135));
            _nw51[(new BigNumber(2)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(3)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(4)).toNumber()] = (_320_v137).dtor_cf98;
            _nw51[(new BigNumber(5)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(6)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(7)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(8)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(9)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(10)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(11)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(12)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(13)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(14)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(15)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(16)).toNumber()] = (_320_v137).dtor_cf98;
            _nw51[(new BigNumber(17)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(18)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(19)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(20)).toNumber()] = _316_v135;
            _nw51[(new BigNumber(21)).toNumber()] = _316_v135;
            _321_v138 = _nw51;
            let _index28 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_321_v138).length));
            (_321_v138)[_index28] = _316_v135;
            let _index29 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_321_v138).length));
            let _rhs25 = _dafny.Seq.Concat(_dafny.Seq.Concat(_313_v132, _313_v132), _dafny.Seq.Concat(_313_v132, _313_v132));
            let _rhs26 = _316_v135;
            let _rhs27 = new BigNumber(306);
            let _rhs28 = !(true);
            let _rhs29 = _313_v132;
            let _lhs13 = _321_v138;
            let _lhs14 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_321_v138).length));
            let _lhs15 = _167_globalState;
            _313_v132 = _rhs25;
            _lhs13[_lhs14] = _rhs26;
            _165_v3 = _rhs27;
            _lhs15.f4 = _rhs28;
            _313_v132 = _rhs29;
          }
        }
      }
      let _322_v139;
      _322_v139 = _dafny.Map.Empty.slice().updateUnsafe(true,_165_v3);
      let _323_v140;
      _323_v140 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_322_v139);
      _165_v3 = ((_163_v1) ? (_module.__default.fm0(_163_v1, _323_v140, _163_v1, _167_globalState)) : (_165_v3));
      let _324_v141;
      _324_v141 = _dafny.Seq.UnicodeFromString("r");
      let _325_v142;
      _325_v142 = _dafny.Seq.of(_162_v0);
      let _326_v143;
      _326_v143 = _dafny.Seq.of(_325_v142, _325_v142, _325_v142, _325_v142, _325_v142);
      let _hi3 = new BigNumber(((_326_v143)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_165_v3)), new BigNumber((_326_v143).length))]).length);
      for (let _327_i12 = new BigNumber((_324_v141).length); _327_i12.isLessThan(_hi3); _327_i12 = _327_i12.plus(_dafny.ONE)) {
        let _328_v145;
        let _init11 = ((_329_v2, _330_v1) => function (_331_i13) {
          return (function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of (_dafny.MultiSet.fromElements(_329_v2, _dafny.MultiSet.fromElements(_330_v1))).Elements) {
              let _332_v144 = _compr_33;
              if ((_dafny.MultiSet.fromElements(_329_v2, _dafny.MultiSet.fromElements(_330_v1))).contains(_332_v144)) {
                _coll33.add(_332_v144);
              }
            }
            return _coll33;
          }()).Difference(_dafny.Set.fromElements(_329_v2));
        })(_164_v2, _163_v1);
        let _nw52 = Array((new BigNumber(3)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw52.length); _i0_11++) {
          _nw52[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _328_v145 = _nw52;
        let _333_v146;
        _333_v146 = _dafny.Set.fromElements(_164_v2);
        let _index30 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_328_v145).length));
        (_328_v145)[_index30] = _333_v146;
        let _index31 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_328_v145).length));
        (_328_v145)[_index31] = _333_v146;
        let _334_v147;
        let _nw53 = Array((new BigNumber(15)).toNumber());
        _334_v147 = _nw53;
        let _335_v148;
        let _nw54 = new _module.C7();
        _nw54.__ctor();
        _335_v148 = _nw54;
        let _index32 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_334_v147).length));
        (_334_v147)[_index32] = _335_v148;
        let _index33 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_334_v147).length));
        (_334_v147)[_index33] = (((_163_v1) === (!(_163_v1))) ? (_335_v148) : (_335_v148));
        if (_163_v1) {
          let _336_v149;
          _336_v149 = _dafny.MultiSet.fromElements(_327_i12);
          let _337_v150;
          _337_v150 = _dafny.Seq.of((_dafny.ZERO).minus(_327_i12));
          let _338_v151;
          let _nw55 = Array((new BigNumber(5)).toNumber());
          _nw55[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_165_v3));
          _nw55[(_dafny.ONE).toNumber()] = _336_v149;
          _nw55[(new BigNumber(2)).toNumber()] = _336_v149;
          _nw55[(new BigNumber(3)).toNumber()] = _336_v149;
          _nw55[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_337_v150);
          _338_v151 = _nw55;
          let _339_v152;
          _339_v152 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_163_v1);
          let _340_v153;
          _340_v153 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_339_v152);
          let _341_v154;
          _341_v154 = _dafny.Seq.of(_163_v1);
          let _342_v155;
          let _nw56 = new _module.C3();
          _nw56.__ctor(_163_v1, _338_v151, _163_v1, _340_v153, _341_v154);
          _342_v155 = _nw56;
          let _343_v156;
          _343_v156 = _module.D24.create_DC65(_342_v155);
          _342_v155 = (_343_v156).dtor_cf99;
          (_167_globalState).f4 = !(_module.__default.fm2((_342_v155).f20, _167_globalState));
          let _344_v157;
          let _345_v158;
          let _346_v159;
          let _347_v160;
          let _out13;
          let _out14;
          let _out15;
          let _out16;
          let _outcollector3 = _module.__default.m0((_342_v155).f20, _module.D0.create_DC2(_165_v3, (_342_v155).f20), ((_342_v155).f20) || (_163_v1), _167_globalState);
          _out13 = _outcollector3[0];
          _out14 = _outcollector3[1];
          _out15 = _outcollector3[2];
          _out16 = _outcollector3[3];
          _344_v157 = _out13;
          _345_v158 = _out14;
          _346_v159 = _out15;
          _347_v160 = _out16;
          let _348_v161;
          let _349_v162;
          let _350_v163;
          let _351_v164;
          let _out17;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector4 = (_342_v155).m9(true, !(false), _324_v141, _167_globalState);
          _out17 = _outcollector4[0];
          _out18 = _outcollector4[1];
          _out19 = _outcollector4[2];
          _out20 = _outcollector4[3];
          _348_v161 = _out17;
          _349_v162 = _out18;
          _350_v163 = _out19;
          _351_v164 = _out20;
          let _352_v165;
          let _nw57 = new _module.C2();
          _nw57.__ctor(_344_v157, (_349_v162).isLessThan(_327_i12), _340_v153, _341_v154, _338_v151);
          _352_v165 = _nw57;
        } else {
          _165_v3 = _module.__default.fm0(!(_163_v1), _323_v140, _dafny.areEqual(_324_v141, _324_v141), _167_globalState);
          let _353_v167;
          let _init12 = ((_354_i12, _355_v1, _356_v141, _357_v4, _358_v3) => function (_359_i14) {
            return ((false) ? (_module.D20.create_DC54(new _dafny.CodePoint('e'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(_354_i12,_355_v1), _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), function (_360_i15) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})), new BigNumber(-472))) : (_module.D20.create_DC54(new _dafny.CodePoint('g'.codePointAt(0)), function () {
  let _coll34 = new _dafny.Map();
  for (const _compr_34 of (_357_v4).Keys.Elements) {
    let _361_v166 = _compr_34;
    if ((_357_v4).contains(_361_v166)) {
      _coll34.push([(_361_v166).multipliedBy(new BigNumber((_dafny.Seq.of(false)).length)),_355_v1]);
    }
  }
  return _coll34;
}(), _dafny.Map.Empty.slice().updateUnsafe(_355_v1,_356_v141), _358_v3)));
          })(_327_i12, _163_v1, _324_v141, _166_v4, _165_v3);
          let _nw58 = Array((new BigNumber(15)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw58.length); _i0_12++) {
            _nw58[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _353_v167 = _nw58;
          let _362_v168;
          _362_v168 = _dafny.MultiSet.fromElements(_327_i12, _327_i12);
          let _index34 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_353_v167).length));
          (_353_v167)[_index34] = _module.__default.fm69(new BigNumber(141), (((_362_v168).contains(_327_i12)) ? ((_362_v168).get(_327_i12)) : (_165_v3)), _165_v3, _167_globalState);
          let _363_v169;
          _363_v169 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_327_i12),true);
          let _364_v170;
          _364_v170 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_324_v141);
          let _365_v171;
          _365_v171 = _dafny.Set.fromElements(_165_v3);
          let _366_v172;
          _366_v172 = _dafny.Set.fromElements(_365_v171);
          let _367_v173;
          _367_v173 = _module.D20.create_DC54(new _dafny.CodePoint('p'.codePointAt(0)), _363_v169, _364_v170, new BigNumber((_366_v172).length));
          let _index35 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_353_v167).length));
          (_353_v167)[_index35] = _367_v173;
          _165_v3 = _327_i12;
          let _368_v174;
          let _init13 = ((_369_v141) => function (_370_i16) {
            return _369_v141;
          })(_324_v141);
          let _nw59 = Array((new BigNumber(28)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw59.length); _i0_13++) {
            _nw59[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _368_v174 = _nw59;
          let _371_v175;
          _371_v175 = _dafny.Seq.of(_165_v3, _165_v3, new BigNumber(-477), _165_v3, _327_i12);
          let _372_v176;
          _372_v176 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(898)), ((_373_i12) => function (_374_i17) {
            return _373_i12;
          })(_327_i12)), _371_v175, _371_v175, _371_v175, _371_v175);
          let _375_v177;
          _375_v177 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(772),_362_v168);
          let _376_v178;
          let _nw60 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
          _376_v178 = _nw60;
          let _377_v179;
          let _nw61 = new _module.C6();
          _nw61.__ctor(_dafny.Seq.Concat(_371_v175, (_372_v176)[_module.__default.safeIndex(_165_v3, new BigNumber((_372_v176).length))]), (_163_v1) && (_163_v1), _dafny.Map.Empty.slice().updateUnsafe(true,_module.__default.fm59((((_375_v177).contains(_327_i12)) ? ((_375_v177).get(_327_i12)) : (_362_v168)), _163_v1, _163_v1, _165_v3, _167_globalState)), _dafny.Seq.of(_163_v1), _376_v178);
          _377_v179 = _nw61;
          let _index36 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_162_v0).length));
          (_162_v0)[_index36] = new BigNumber(985);
          let _index37 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_162_v0).length));
          let _rhs30 = !(_163_v1);
          let _rhs31 = _368_v174;
          let _rhs32 = _377_v179;
          let _rhs33 = _327_i12;
          let _rhs34 = _324_v141;
          let _lhs16 = _162_v0;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_162_v0).length));
          _163_v1 = _rhs30;
          _368_v174 = _rhs31;
          _377_v179 = _rhs32;
          _lhs16[_lhs17] = _rhs33;
          _324_v141 = _rhs34;
          let _378_v180;
          _378_v180 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_163_v1);
          let _379_v181;
          _379_v181 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_378_v180);
          let _380_v182;
          _380_v182 = _dafny.Seq.of(_163_v1, _163_v1, _163_v1, _163_v1);
          let _381_v183;
          let _nw62 = new _module.C3();
          _nw62.__ctor(_163_v1, _376_v178, _163_v1, (_379_v181).Merge(_module.__default.fm60(new BigNumber((_378_v180).length), _167_globalState)), _dafny.Seq.Concat(_380_v182, _dafny.Seq.of(_163_v1)));
          _381_v183 = _nw62;
          _381_v183 = _381_v183;
        }
        let _382_v184;
        let _383_v185;
        let _384_v186;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector5 = (_335_v148).m2(new BigNumber(993), _167_globalState);
        _out21 = _outcollector5[0];
        _out22 = _outcollector5[1];
        _out23 = _outcollector5[2];
        _382_v184 = _out21;
        _383_v185 = _out22;
        _384_v186 = _out23;
      }
      let _385_v187;
      let _nw63 = Array((new BigNumber(27)).toNumber()).fill(_dafny.MultiSet.Empty);
      _385_v187 = _nw63;
      let _386_v188;
      _386_v188 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_163_v1);
      let _387_v189;
      _387_v189 = _dafny.Seq.of(_163_v1);
      let _388_v190;
      let _nw64 = new _module.C3();
      _nw64.__ctor(_163_v1, _385_v187, _163_v1, _dafny.Map.Empty.slice().updateUnsafe(_163_v1,_386_v188), _387_v189);
      _388_v190 = _nw64;
      let _389_v191;
      _389_v191 = _module.D14.create_DC40(_388_v190);
      let _pat_let_tv0 = _388_v190;
      let _source4 = function (_pat_let0_0) {
        return function (_390_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_391_dt__update_hcf63_h0) {
              return _module.D14.create_DC40(_391_dt__update_hcf63_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_389_v191);
      if (_source4.is_DC41) {
        let _392___mcc_h0 = (_source4).cf64;
        let _393___mcc_h1 = (_source4).cf65;
        let _394_cf65 = _393___mcc_h1;
        let _395_cf64 = _392___mcc_h0;
        if ((_388_v190).f15) {
          (_167_globalState).f4 = (_165_v3).isLessThan(_394_cf65);
          _324_v141 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(243)), function (_396_i18) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), _324_v141);
          let _397_v192;
          let _nw65 = Array((new BigNumber(8)).toNumber());
          _nw65[(_dafny.ZERO).toNumber()] = false;
          _nw65[(_dafny.ONE).toNumber()] = (_388_v190).f15;
          _nw65[(new BigNumber(2)).toNumber()] = _module.__default.fm2((_388_v190).f15, _167_globalState);
          _nw65[(new BigNumber(3)).toNumber()] = !(_163_v1);
          _nw65[(new BigNumber(4)).toNumber()] = _163_v1;
          _nw65[(new BigNumber(5)).toNumber()] = false;
          _nw65[(new BigNumber(6)).toNumber()] = false;
          _nw65[(new BigNumber(7)).toNumber()] = (_388_v190).f15;
          _397_v192 = _nw65;
          let _398_v193;
          _398_v193 = _dafny.Map.Empty.slice().updateUnsafe(_397_v192,_163_v1);
          let _399_v194;
          let _nw66 = Array((new BigNumber(17)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _163_v1;
          _nw66[(_dafny.ONE).toNumber()] = (((_398_v193).contains(_397_v192)) ? ((_398_v193).get(_397_v192)) : ((_388_v190).f15));
          _nw66[(new BigNumber(2)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(3)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(4)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(5)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(6)).toNumber()] = !(!_dafny.areEqual(_324_v141, _395_cf64));
          _nw66[(new BigNumber(7)).toNumber()] = _163_v1;
          _nw66[(new BigNumber(8)).toNumber()] = _163_v1;
          _nw66[(new BigNumber(9)).toNumber()] = _module.__default.fm2((_388_v190).f15, _167_globalState);
          _nw66[(new BigNumber(10)).toNumber()] = ((_388_v190).f15) === ((_388_v190).f15);
          _nw66[(new BigNumber(11)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(12)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(13)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(14)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(15)).toNumber()] = (_388_v190).f15;
          _nw66[(new BigNumber(16)).toNumber()] = (_388_v190).f15;
          _399_v194 = _nw66;
          let _400_v195;
          _400_v195 = _dafny.MultiSet.fromElements(new BigNumber((_324_v141).length), _165_v3);
          let _401_v196;
          _401_v196 = _dafny.Map.Empty.slice().updateUnsafe(_400_v195,_165_v3);
          let _402_v197;
          _402_v197 = _dafny.Set.fromElements(_394_cf65);
          let _403_v198;
          _403_v198 = _dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)));
          let _404_v199;
          _404_v199 = _module.D5.create_DC16((_388_v190).f15);
          let _nw67 = Array((new BigNumber(20)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = !(!(_402_v197).contains(new BigNumber((_401_v196).length)));
          _nw67[(_dafny.ONE).toNumber()] = _163_v1;
          _nw67[(new BigNumber(2)).toNumber()] = !((_388_v190).f15);
          _nw67[(new BigNumber(3)).toNumber()] = (_165_v3).isLessThan((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_164_v2).cardinality()),_403_v198)).length)));
          _nw67[(new BigNumber(4)).toNumber()] = !(_163_v1);
          _nw67[(new BigNumber(5)).toNumber()] = (_388_v190).f15;
          _nw67[(new BigNumber(6)).toNumber()] = !((_388_v190).f15);
          _nw67[(new BigNumber(7)).toNumber()] = (_388_v190).f15;
          _nw67[(new BigNumber(8)).toNumber()] = (_388_v190).f15;
          _nw67[(new BigNumber(9)).toNumber()] = (_388_v190).f15;
          _nw67[(new BigNumber(10)).toNumber()] = _163_v1;
          _nw67[(new BigNumber(11)).toNumber()] = (_388_v190).f15;
          _nw67[(new BigNumber(12)).toNumber()] = ((_388_v190).f15) === (!(_163_v1));
          _nw67[(new BigNumber(13)).toNumber()] = _163_v1;
          _nw67[(new BigNumber(14)).toNumber()] = _163_v1;
          _nw67[(new BigNumber(15)).toNumber()] = ((_163_v1) ? ((_388_v190).f15) : ((_388_v190).f15));
          _nw67[(new BigNumber(16)).toNumber()] = (_388_v190).f15;
          _nw67[(new BigNumber(17)).toNumber()] = (false) && (_163_v1);
          _nw67[(new BigNumber(18)).toNumber()] = (_388_v190).f15;
          _nw67[(new BigNumber(19)).toNumber()] = (_404_v199).dtor_cf26;
          _399_v194 = _nw67;
          _394_cf65 = ((!_dafny.areEqual(_395_cf64, _324_v141)) ? (_165_v3) : ((_388_v190).fm8(_module.__default.fm2(true, _167_globalState), _163_v1, _167_globalState)));
          _165_v3 = _394_cf65;
        } else {
          let _405_v200;
          let _init14 = ((_406_v1) => function (_407_i19) {
            return _406_v1;
          })(_163_v1);
          let _nw68 = Array((new BigNumber(2)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw68.length); _i0_14++) {
            _nw68[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _405_v200 = _nw68;
          let _index38 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_405_v200).length));
          (_405_v200)[_index38] = _module.__default.fm2(false, _167_globalState);
          let _408_v201;
          let _init15 = function (_409_i20) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          };
          let _nw69 = Array((new BigNumber(15)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw69.length); _i0_15++) {
            _nw69[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _408_v201 = _nw69;
          let _410_v202;
          _410_v202 = _dafny.Seq.of(_408_v201);
          let _411_v203;
          _411_v203 = _410_v202;
          let _index39 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_405_v200).length));
          let _rhs35 = (_388_v190).f15;
          let _rhs36 = (_dafny.ZERO).minus(_394_cf65);
          let _rhs37 = (_388_v190).f15;
          let _rhs38 = (_388_v190).f15;
          let _rhs39 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_410_v202, _dafny.Seq.update(_410_v202, _module.__default.safeIndex(_394_cf65, new BigNumber((_410_v202).length)), _408_v201)), (_411_v203));
          let _lhs18 = _167_globalState;
          let _lhs19 = _405_v200;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_405_v200).length));
          let _lhs21 = _167_globalState;
          _163_v1 = _rhs35;
          _394_cf65 = _rhs36;
          _lhs18.f4 = _rhs37;
          _lhs19[_lhs20] = _rhs38;
          _lhs21.f4 = _rhs39;
          let _412_v204;
          let _nw70 = new _module.C4();
          _nw70.__ctor();
          _412_v204 = _nw70;
          let _413_v205;
          let _nw71 = Array((new BigNumber(6)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = _412_v204;
          _nw71[(_dafny.ONE).toNumber()] = _412_v204;
          _nw71[(new BigNumber(2)).toNumber()] = _412_v204;
          _nw71[(new BigNumber(3)).toNumber()] = _412_v204;
          _nw71[(new BigNumber(4)).toNumber()] = _412_v204;
          _nw71[(new BigNumber(5)).toNumber()] = _412_v204;
          _413_v205 = _nw71;
          let _414_v206;
          _414_v206 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_388_v190).f15, _167_globalState),_412_v204);
          let _index40 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_413_v205).length));
          (_413_v205)[_index40] = (((_414_v206).contains((_405_v200)[_module.__default.safeIndex(new BigNumber(430), new BigNumber((_405_v200).length))])) ? ((_414_v206).get((_405_v200)[_module.__default.safeIndex(new BigNumber(430), new BigNumber((_405_v200).length))])) : (_412_v204));
          let _index41 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_413_v205).length));
          (_413_v205)[_index41] = _412_v204;
          _165_v3 = _module.__default.safeDivisionInt(((((_164_v2).contains((_388_v190).f15)) ? ((_164_v2).get((_388_v190).f15)) : ((_388_v190).fm5(false, _167_globalState)))).multipliedBy(new BigNumber((_324_v141).length)), _165_v3);
          _394_cf65 = _394_cf65;
          _165_v3 = _module.__default.safeDivisionInt(_394_cf65, _165_v3);
        }
        let _415_v207;
        let _nw72 = new _module.C5();
        _nw72.__ctor((_388_v190).f15, _388_v190.f12, ((_388_v190).f13).update((_388_v190).f15, _386_v188), _dafny.Seq.of(_163_v1));
        _415_v207 = _nw72;
        let _416_v208;
        _416_v208 = _dafny.Map.Empty.slice().updateUnsafe(_415_v207,_394_cf65);
        _165_v3 = _module.__default.safeModuloInt(_394_cf65, _module.__default.safeModuloInt(new BigNumber(((_416_v208).update(_415_v207, _394_cf65)).length), _165_v3));
        let _417_v209;
        let _nw73 = Array((new BigNumber(11)).toNumber());
        _417_v209 = _nw73;
        let _index42 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_417_v209).length));
        (_417_v209)[_index42] = _388_v190;
        let _418_v210;
        _418_v210 = _module.D21.create_DC56(_388_v190);
        let _index43 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_417_v209).length));
        let _rhs40 = (_418_v210).dtor_cf85;
        let _rhs41 = _162_v0;
        let _rhs42 = (_163_v1) === (_163_v1);
        let _rhs43 = new BigNumber((_324_v141).length);
        let _rhs44 = _394_cf65;
        let _lhs22 = _417_v209;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_417_v209).length));
        let _lhs24 = _167_globalState;
        _lhs22[_lhs23] = _rhs40;
        _162_v0 = _rhs41;
        _lhs24.f4 = _rhs42;
        _394_cf65 = _rhs43;
        _165_v3 = _rhs44;
        let _419_v211;
        let _nw74 = Array((new BigNumber(2)).toNumber()).fill(false);
        _419_v211 = _nw74;
        let _420_v212;
        _420_v212 = _dafny.Set.fromElements(_419_v211);
        let _421_v213;
        _421_v213 = _module.D22.create_DC62(_419_v211);
        let _422_v214;
        _422_v214 = _dafny.Seq.of(_419_v211, _419_v211, (_421_v213).dtor_cf97, _419_v211, _419_v211);
        (_167_globalState).f4 = ((_420_v212).Intersect(_dafny.Set.fromElements((_422_v214)[_module.__default.safeIndex(_165_v3, new BigNumber((_422_v214).length))]))).contains(_419_v211);
      } else if (_source4.is_DC42) {
        let _423_v215;
        _423_v215 = _module.D0.create_DC2(_165_v3, (_388_v190).f15);
        let _424_v216;
        let _425_v217;
        let _426_v218;
        let _427_v219;
        let _out24;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector6 = _module.__default.m0(_163_v1, _423_v215, (_388_v190).f15, _167_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _out26 = _outcollector6[2];
        _out27 = _outcollector6[3];
        _424_v216 = _out24;
        _425_v217 = _out25;
        _426_v218 = _out26;
        _427_v219 = _out27;
        let _428_v220;
        _428_v220 = _dafny.Seq.of(_388_v190);
        (_167_globalState).f4 = (_426_v218).isLessThanOrEqualTo(new BigNumber((_428_v220).length));
        _426_v218 = (_388_v190).fm5((_388_v190).f15, _167_globalState);
        let _429_v221;
        _429_v221 = _module.D8.create_DC25(!((_388_v190).f15), new BigNumber((_424_v216).length), _426_v218, _322_v139);
        let _430_v222;
        _430_v222 = _module.D8.create_DC27(_429_v221);
        let _source5 = _430_v222;
        if (_source5.is_DC24) {
          let _431___mcc_h4 = (_source5).cf36;
          let _432_cf36 = _431___mcc_h4;
          let _index44 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_162_v0).length));
          (_162_v0)[_index44] = (_388_v190).fm5((_388_v190).f15, _167_globalState);
          let _index45 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_162_v0).length));
          (_162_v0)[_index45] = (_dafny.ZERO).minus(_165_v3);
          let _433_v223;
          let _nw75 = Array((new BigNumber(25)).toNumber()).fill(false);
          _433_v223 = _nw75;
          let _index46 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_433_v223).length));
          (_433_v223)[_index46] = (_426_v218).isEqualTo(_426_v218);
          let _434_v224;
          _434_v224 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,(_388_v190).f15);
          let _index47 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_433_v223).length));
          (_433_v223)[_index47] = (_module.__default.fm2((_388_v190).f15, _167_globalState)) === ((new BigNumber((_434_v224).length)).isEqualTo((_162_v0)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_162_v0).length))]));
          _426_v218 = new BigNumber((_324_v141).length);
          let _435_v225;
          _435_v225 = new _dafny.CodePoint('m'.codePointAt(0));
          let _436_v226;
          _436_v226 = _dafny.MultiSet.fromElements((_162_v0)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_162_v0).length))]);
          let _index48 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_162_v0).length));
          (_162_v0)[_index48] = (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_435_v225,(_388_v190).f15)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_435_v225,(_388_v190).f15))).length)).plus((_165_v3).plus(new BigNumber((_436_v226).cardinality())));
        } else if (_source5.is_DC25) {
          let _437___mcc_h5 = (_source5).cf37;
          let _438___mcc_h6 = (_source5).cf38;
          let _439___mcc_h7 = (_source5).cf39;
          let _440___mcc_h8 = (_source5).cf40;
          let _441_cf40 = _440___mcc_h8;
          let _442_cf39 = _439___mcc_h7;
          let _443_cf38 = _438___mcc_h6;
          let _444_cf37 = _437___mcc_h5;
          (_167_globalState).f4 = !(false);
          let _445_v227;
          let _init16 = ((_446_v141) => function (_447_i21) {
            return _446_v141;
          })(_324_v141);
          let _nw76 = Array((new BigNumber(27)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw76.length); _i0_16++) {
            _nw76[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _445_v227 = _nw76;
          _445_v227 = _445_v227;
          let _448_v228;
          let _nw77 = new _module.C6();
          _nw77.__ctor(_dafny.Seq.of(new BigNumber((_164_v2).cardinality())), _163_v1, (_388_v190).f13, (_388_v190).f14, ((!(_444_cf37)) ? (_385_v187) : (_388_v190.f12)));
          _448_v228 = _nw77;
          _442_cf39 = _426_v218;
        } else if (_source5.is_DC26) {
          let _449___mcc_h9 = (_source5).cf41;
          let _450_cf41 = _449___mcc_h9;
          _426_v218 = _165_v3;
          _426_v218 = _426_v218;
          let _451_v229;
          _451_v229 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),(_388_v190).f15);
          let _452_v230;
          _452_v230 = new _dafny.CodePoint('a'.codePointAt(0));
          _451_v229 = _dafny.Map.Empty.slice().updateUnsafe(_452_v230,false);
          _165_v3 = (_426_v218).multipliedBy(new BigNumber(71));
        } else if (_source5.is_DC23) {
          let _453___mcc_h10 = (_source5).cf35;
          let _454_cf35 = _453___mcc_h10;
          let _455_v231;
          let _nw78 = new _module.C3();
          _nw78.__ctor(_163_v1, _388_v190.f12, false, (_388_v190).f13, _387_v189);
          _455_v231 = _nw78;
          let _456_v232;
          _456_v232 = _dafny.Map.Empty.slice().updateUnsafe(_455_v231,(_388_v190).f15);
          let _457_v233;
          let _458_v234;
          let _459_v235;
          let _460_v236;
          let _out28;
          let _out29;
          let _out30;
          let _out31;
          let _outcollector7 = _module.__default.m0((_456_v232).contains(_455_v231), _423_v215, (_455_v231).f20, _167_globalState);
          _out28 = _outcollector7[0];
          _out29 = _outcollector7[1];
          _out30 = _outcollector7[2];
          _out31 = _outcollector7[3];
          _457_v233 = _out28;
          _458_v234 = _out29;
          _459_v235 = _out30;
          _460_v236 = _out31;
          (_167_globalState).f4 = (_388_v190).f15;
          let _461_v237;
          let _462_v238;
          let _463_v239;
          let _464_v240;
          let _out32;
          let _out33;
          let _out34;
          let _out35;
          let _outcollector8 = (_455_v231).m9((new BigNumber((_dafny.MultiSet.fromElements(_455_v231, _455_v231, _455_v231)).cardinality())).isLessThan(_426_v218), (_388_v190).f15, _dafny.Seq.UnicodeFromString("fgwsvo"), _167_globalState);
          _out32 = _outcollector8[0];
          _out33 = _outcollector8[1];
          _out34 = _outcollector8[2];
          _out35 = _outcollector8[3];
          _461_v237 = _out32;
          _462_v238 = _out33;
          _463_v239 = _out34;
          _464_v240 = _out35;
          let _465_v241;
          let _466_v242;
          let _467_v243;
          let _468_v244;
          let _out36;
          let _out37;
          let _out38;
          let _out39;
          let _outcollector9 = _module.__default.m0(_163_v1, _module.D0.create_DC2(_462_v238, _460_v236), ((_388_v190).f14)[_module.__default.safeIndex(_459_v235, new BigNumber(((_388_v190).f14).length))], _167_globalState);
          _out36 = _outcollector9[0];
          _out37 = _outcollector9[1];
          _out38 = _outcollector9[2];
          _out39 = _outcollector9[3];
          _465_v241 = _out36;
          _466_v242 = _out37;
          _467_v243 = _out38;
          _468_v244 = _out39;
        } else {
          let _469___mcc_h11 = (_source5).cf42;
          let _470_cf42 = _469___mcc_h11;
          (_167_globalState).f4 = (_388_v190).f15;
          _387_v189 = (_388_v190).f14;
          let _index49 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_162_v0).length));
          (_162_v0)[_index49] = _module.__default.safeDivisionInt(_426_v218, new BigNumber((_209_v40).length));
          let _471_v245;
          let _nw79 = Array((new BigNumber(6)).toNumber()).fill(false);
          _471_v245 = _nw79;
          let _index50 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_471_v245).length));
          (_471_v245)[_index50] = (new BigNumber((_166_v4).length)).isEqualTo(_165_v3);
          let _index51 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_162_v0).length));
          let _index52 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_471_v245).length));
          let _rhs45 = (((_388_v190).f15) ? (_322_v139) : (((_dafny.Map.Empty.slice().updateUnsafe(_163_v1,_426_v218)).update((_388_v190).f15, _165_v3)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_388_v190).f15,new BigNumber(969)))));
          let _rhs46 = !(_163_v1);
          let _rhs47 = _426_v218;
          let _rhs48 = !(_163_v1);
          let _lhs25 = _162_v0;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_162_v0).length));
          let _lhs27 = _471_v245;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_471_v245).length));
          _322_v139 = _rhs45;
          _425_v217 = _rhs46;
          _lhs25[_lhs26] = _rhs47;
          _lhs27[_lhs28] = _rhs48;
          _163_v1 = (_388_v190).f15;
        }
      } else if (_source4.is_DC40) {
        let _472___mcc_h2 = (_source4).cf63;
        let _473_cf63 = _472___mcc_h2;
        _163_v1 = _module.__default.fm2(true, _167_globalState);
        let _474_v246;
        let _nw80 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _474_v246 = _nw80;
        let _475_v247;
        _475_v247 = _dafny.Seq.of(_474_v246);
        let _476_v248;
        _476_v248 = _dafny.Map.Empty.slice().updateUnsafe((_475_v247)[_module.__default.safeIndex(_165_v3, new BigNumber((_475_v247).length))],(_dafny.ZERO).minus((((_166_v4).contains(_165_v3)) ? ((_166_v4).get(_165_v3)) : ((_473_cf63).fm5((_388_v190).f15, _167_globalState)))));
        let _477_v249;
        _477_v249 = _dafny.Set.fromElements(_165_v3, new BigNumber(((_473_cf63).f14).length));
        _476_v248 = (_476_v248).update(_474_v246, new BigNumber(((_477_v249).Intersect(_477_v249)).length));
        let _478_v250;
        let _479_v251;
        let _480_v252;
        let _481_v253;
        let _out40;
        let _out41;
        let _out42;
        let _out43;
        let _outcollector10 = (_388_v190).m3(_167_globalState);
        _out40 = _outcollector10[0];
        _out41 = _outcollector10[1];
        _out42 = _outcollector10[2];
        _out43 = _outcollector10[3];
        _478_v250 = _out40;
        _479_v251 = _out41;
        _480_v252 = _out42;
        _481_v253 = _out43;
        (_167_globalState).f4 = _478_v250;
      } else {
        let _482___mcc_h3 = (_source4).cf66;
        let _483_cf66 = _482___mcc_h3;
        let _484_v254;
        _484_v254 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_388_v190).fm5((_388_v190).f15, _167_globalState))));
        _484_v254 = _dafny.Seq.Concat(_484_v254, _dafny.Seq.Concat(_484_v254, _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_388_v190).f15,_163_v1)).length), _165_v3, _165_v3), _module.__default.safeIndex(new BigNumber((_324_v141).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_388_v190).f15,_163_v1)).length), _165_v3, _165_v3)).length)), _165_v3)));
        let _485_v255;
        _485_v255 = _module.D12.create_DC37(_324_v141, (_388_v190).f15);
        let _486_v256;
        _486_v256 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_324_v141);
        _163_v1 = _dafny.areEqual(_dafny.Seq.Concat((_485_v255).dtor_cf59, (((_486_v256).contains((_484_v254)[_module.__default.safeIndex(_165_v3, new BigNumber((_484_v254).length))])) ? ((_486_v256).get((_484_v254)[_module.__default.safeIndex(_165_v3, new BigNumber((_484_v254).length))])) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(503)), function (_487_i22) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })))), _324_v141);
        let _488_v257;
        let _nw81 = new _module.C0();
        _nw81.__ctor();
        _488_v257 = _nw81;
        let _489_v258;
        _489_v258 = _module.D21.create_DC56(_388_v190);
        let _source6 = _489_v258;
        if (_source6.is_DC56) {
          let _490___mcc_h12 = (_source6).cf85;
          let _491_cf85 = _490___mcc_h12;
          let _492_v259;
          let _493_v260;
          let _494_v261;
          let _495_v262;
          let _out44;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector11 = (_388_v190).m3(_167_globalState);
          _out44 = _outcollector11[0];
          _out45 = _outcollector11[1];
          _out46 = _outcollector11[2];
          _out47 = _outcollector11[3];
          _492_v259 = _out44;
          _493_v260 = _out45;
          _494_v261 = _out46;
          _495_v262 = _out47;
          let _496_v264;
          _496_v264 = _dafny.Map.Empty.slice().updateUnsafe(_491_cf85,_493_v260);
          let _497_v265;
          _497_v265 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_484_v254).length),new BigNumber((_496_v264).length)), _166_v4);
          let _498_v266;
          _498_v266 = _module.D11.create_DC32(_493_v260, _165_v3);
          let _499_v267;
          let _nw82 = Array((new BigNumber(7)).toNumber());
          _nw82[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), ((_500_v254) => function (_501_i23) {
            return function () {
              let _coll35 = new _dafny.Map();
              for (const _compr_35 of _dafny.IntegerRange(new BigNumber(-646), new BigNumber(701))) {
                let _502_v263 = _compr_35;
                if (((new BigNumber(-646)).isLessThanOrEqualTo(_502_v263)) && ((_502_v263).isLessThan(new BigNumber(701)))) {
                  _coll35.push([(_502_v263).plus(new BigNumber(814)),new BigNumber((_500_v254).length)]);
                }
              }
              return _coll35;
            }();
          })(_484_v254));
          _nw82[(_dafny.ONE).toNumber()] = _497_v265;
          _nw82[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_497_v265, _497_v265);
          _nw82[(new BigNumber(3)).toNumber()] = _497_v265;
          _nw82[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_497_v265, _dafny.Seq.of(_166_v4));
          _nw82[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), ((_503_v3) => function (_504_i24) {
            return _dafny.Map.Empty.slice().updateUnsafe(_503_v3,_503_v3);
          })(_165_v3)), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_498_v266).dtor_cf46,(_dafny.ZERO).minus(_165_v3)), (_497_v265)[_module.__default.safeIndex((_dafny.ZERO).minus(_493_v260), new BigNumber((_497_v265).length))], _166_v4));
          _nw82[(new BigNumber(6)).toNumber()] = _497_v265;
          _499_v267 = _nw82;
          let _index53 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_499_v267).length));
          (_499_v267)[_index53] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(384)), ((_505_v4) => function (_506_i25) {
            return _505_v4;
          })(_166_v4));
          let _index54 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_499_v267).length));
          (_499_v267)[_index54] = _497_v265;
          let _507_v268;
          let _nw83 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _507_v268 = _nw83;
          let _508_v269;
          _508_v269 = _dafny.Map.Empty.slice().updateUnsafe((_493_v260).multipliedBy(_493_v260),_507_v268);
          let _509_v270;
          let _nw84 = new _module.C3();
          _nw84.__ctor((_388_v190).f15, _388_v190.f12, (_388_v190).f15, (_388_v190).f13, (_388_v190).f14);
          _509_v270 = _nw84;
          let _510_v271;
          _510_v271 = _module.D24.create_DC65(_509_v270);
          let _511_v272;
          _511_v272 = _dafny.Map.Empty.slice().updateUnsafe(_510_v271,_507_v268);
          _508_v269 = (_508_v269).update(new BigNumber((_324_v141).length), (((_511_v272).contains(_510_v271)) ? ((_511_v272).get(_510_v271)) : (_507_v268)));
          _488_v257 = _488_v257;
        } else if (_source6.is_DC57) {
          let _512___mcc_h13 = (_source6).cf86;
          let _513___mcc_h14 = (_source6).cf87;
          let _514___mcc_h15 = (_source6).cf88;
          let _515_cf88 = _514___mcc_h15;
          let _516_cf87 = _513___mcc_h14;
          let _517_cf86 = _512___mcc_h13;
          let _518_v273;
          let _nw85 = new _module.C7();
          _nw85.__ctor();
          _518_v273 = _nw85;
          let _nw86 = new _module.C7();
          _nw86.__ctor();
          _518_v273 = _nw86;
          _163_v1 = !(!(((((_164_v2).contains(true)) ? ((_164_v2).get(true)) : ((_dafny.ZERO).minus(_165_v3)))).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_165_v3, new BigNumber(-986)))));
          let _519_v274;
          _519_v274 = _dafny.Set.fromElements(_165_v3);
          let _520_v275;
          _520_v275 = _dafny.Map.Empty.slice().updateUnsafe(_488_v257,_519_v274);
          let _521_v276;
          let _nw87 = new _module.C2();
          _nw87.__ctor(_484_v254, _163_v1, _dafny.Map.Empty.slice().updateUnsafe((_388_v190).f15,_dafny.Map.Empty.slice().updateUnsafe(!(true),(_388_v190).f15)), _387_v189, _388_v190.f12);
          _521_v276 = _nw87;
          let _522_v277;
          _522_v277 = _dafny.Seq.of(_521_v276);
          _520_v275 = (_520_v275).update(_488_v257, _dafny.Set.fromElements(new BigNumber((_522_v277).length), (((_164_v2).contains((_388_v190).f15)) ? ((_164_v2).get((_388_v190).f15)) : ((_dafny.ZERO).minus(new BigNumber(((_module.D14.create_DC41(_324_v141, new BigNumber((_519_v274).length))).dtor_cf64).length)))), (_dafny.ZERO).minus(_165_v3)));
          let _523_v278;
          let _nw88 = Array((new BigNumber(18)).toNumber()).fill(false);
          _523_v278 = _nw88;
          let _524_v279;
          _524_v279 = _module.D2.create_DC6(_523_v278, _523_v278);
          let _525_v280;
          let _nw89 = Array((new BigNumber(15)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _523_v278;
          _nw89[(_dafny.ONE).toNumber()] = _523_v278;
          _nw89[(new BigNumber(2)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(3)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(4)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(5)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(6)).toNumber()] = (_524_v279).dtor_cf9;
          _nw89[(new BigNumber(7)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(8)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(9)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(10)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(11)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(12)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(13)).toNumber()] = _523_v278;
          _nw89[(new BigNumber(14)).toNumber()] = _523_v278;
          _525_v280 = _nw89;
          let _index55 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_525_v280).length));
          (_525_v280)[_index55] = _523_v278;
          let _index56 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_525_v280).length));
          (_525_v280)[_index56] = _523_v278;
        } else if (_source6.is_DC58) {
          let _526___mcc_h16 = (_source6).cf89;
          let _527___mcc_h17 = (_source6).cf90;
          let _528_cf90 = _527___mcc_h17;
          let _529_cf89 = _526___mcc_h16;
          _166_v4 = _dafny.Map.Empty.slice().updateUnsafe(_528_cf90,_528_cf90);
          let _530_v281;
          let _init17 = ((_531_v141) => function (_532_i26) {
            return _531_v141;
          })(_324_v141);
          let _nw90 = Array((new BigNumber(7)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw90.length); _i0_17++) {
            _nw90[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _530_v281 = _nw90;
          let _index57 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_530_v281).length));
          (_530_v281)[_index57] = _324_v141;
          let _533_v282;
          _533_v282 = _dafny.MultiSet.fromElements(_165_v3);
          let _index58 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_530_v281).length));
          let _rhs49 = (_dafny.MultiSet.fromElements(_528_cf90)).IsSubsetOf((_533_v282).update(_528_cf90, _module.__default.abs(_528_cf90)));
          let _rhs50 = !(false);
          let _rhs51 = (_dafny.ZERO).minus((_528_cf90).multipliedBy(_165_v3));
          let _rhs52 = _dafny.Seq.Concat(_dafny.Seq.Concat(_324_v141, _module.__default.fm21(_163_v1, _167_globalState)), _324_v141);
          let _lhs29 = _167_globalState;
          let _lhs30 = _530_v281;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_530_v281).length));
          _529_cf89 = _rhs49;
          _lhs29.f4 = _rhs50;
          _528_cf90 = _rhs51;
          _lhs30[_lhs31] = _rhs52;
          let _534_v283;
          _534_v283 = _dafny.Map.Empty.slice().updateUnsafe(_528_cf90,(_388_v190).f15);
          let _535_v284;
          _535_v284 = _164_v2;
          let _536_v285;
          _536_v285 = _dafny.Seq.of(_535_v284);
          let _537_v286;
          _537_v286 = new _dafny.CodePoint('n'.codePointAt(0));
          let _538_v287;
          let _539_v288;
          let _out48;
          let _out49;
          let _outcollector12 = (_488_v257).m7(_module.__default.safeModuloInt(_165_v3, new BigNumber((_534_v283).length)), _536_v285, _165_v3, _537_v286, _167_globalState);
          _out48 = _outcollector12[0];
          _out49 = _outcollector12[1];
          _538_v287 = _out48;
          _539_v288 = _out49;
          let _540_v289;
          _540_v289 = _dafny.Seq.of((_530_v281)[_module.__default.safeIndex(new BigNumber(651), new BigNumber((_530_v281).length))], (_530_v281)[_module.__default.safeIndex(new BigNumber(651), new BigNumber((_530_v281).length))], _324_v141, (_530_v281)[_module.__default.safeIndex(new BigNumber(651), new BigNumber((_530_v281).length))]);
          (_167_globalState).f4 = _dafny.Seq.IsPrefixOf(_540_v289, _dafny.Seq.update(_540_v289, _module.__default.safeIndex(new BigNumber(722), new BigNumber((_540_v289).length)), (_530_v281)[_module.__default.safeIndex(new BigNumber(651), new BigNumber((_530_v281).length))]));
        } else {
          let _541___mcc_h18 = (_source6).cf84;
          let _542_cf84 = _541___mcc_h18;
          (_167_globalState).f4 = _module.__default.fm2((_388_v190).f15, _167_globalState);
          let _index59 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_162_v0).length));
          (_162_v0)[_index59] = _165_v3;
          let _index60 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_162_v0).length));
          (_162_v0)[_index60] = new BigNumber((_324_v141).length);
          let _543_v290;
          let _nw91 = new _module.C5();
          _nw91.__ctor((_165_v3).isLessThan(_165_v3), _388_v190.f12, (_388_v190).f13, _387_v189);
          _543_v290 = _nw91;
          _543_v290 = _543_v290;
          let _544_v291;
          let _nw92 = new _module.C2();
          _nw92.__ctor(_484_v254, (_388_v190).f15, _dafny.Map.Empty.slice().updateUnsafe((_543_v290).f19,_386_v188), _dafny.Seq.Concat(_387_v189, (_388_v190).f14), _388_v190.f12);
          _544_v291 = _nw92;
          _544_v291 = _544_v291;
        }
      }
      let _545_i27;
      _545_i27 = _dafny.ZERO;
      L2: {
        while ((_388_v190).f15) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_545_i27)) {
              break L2;
            }
            _545_i27 = (_545_i27).plus(_dafny.ONE);
            let _546_v292;
            let _init18 = ((_547_v3) => function (_548_i28) {
              return _dafny.Seq.Concat(_dafny.Seq.of(_547_v3), _dafny.Seq.of(_547_v3));
            })(_165_v3);
            let _nw93 = Array((new BigNumber(24)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw93.length); _i0_18++) {
              _nw93[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _546_v292 = _nw93;
            let _549_v293;
            _549_v293 = _dafny.Seq.of(new BigNumber(-548));
            let _index61 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_546_v292).length));
            (_546_v292)[_index61] = _549_v293;
            let _index62 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_546_v292).length));
            (_546_v292)[_index62] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), ((_550_v3) => function (_551_i29) {
              return _550_v3;
            })(_165_v3)), _549_v293);
            let _552_v294;
            let _nw94 = new _module.C0();
            _nw94.__ctor();
            _552_v294 = _nw94;
            _552_v294 = _552_v294;
            let _553_v295;
            let _init19 = ((_554_v141) => function (_555_i30) {
              return _554_v141;
            })(_324_v141);
            let _nw95 = Array((new BigNumber(17)).toNumber());
            for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw95.length); _i0_19++) {
              _nw95[_i0_19] = _init19(new BigNumber(_i0_19));
            }
            _553_v295 = _nw95;
            let _index63 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_553_v295).length));
            (_553_v295)[_index63] = _324_v141;
            let _556_v296;
            _556_v296 = new _dafny.CodePoint('u'.codePointAt(0));
            let _index64 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_553_v295).length));
            (_553_v295)[_index64] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_324_v141, _module.__default.safeIndex(_165_v3, new BigNumber((_324_v141).length)), _556_v296), _324_v141), _324_v141);
            let _557_v297;
            _557_v297 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_324_v141).length),(_388_v190).f15);
            _557_v297 = (_module.__default.fm49((_388_v190).f15, false, !((_388_v190).f15), new _dafny.CodePoint('n'.codePointAt(0)), _167_globalState)).Merge((_557_v297).Merge(_557_v297));
          }
        }
      }
      if (_163_v1) {
        let _index65 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length));
        (_162_v0)[_index65] = new BigNumber(708);
        let _index66 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length));
        (_162_v0)[_index66] = ((new BigNumber(-942)).plus(_165_v3)).plus((_388_v190).fm5((_388_v190).f15, _167_globalState));
        let _558_v298;
        let _init20 = function (_559_i31) {
          return _dafny.Seq.UnicodeFromString("saum");
        };
        let _nw96 = Array((new BigNumber(8)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw96.length); _i0_20++) {
          _nw96[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _558_v298 = _nw96;
        let _560_v299;
        _560_v299 = _dafny.Seq.of(_558_v298, _558_v298);
        let _561_v300;
        _561_v300 = new _dafny.CodePoint('n'.codePointAt(0));
        let _562_v301;
        _562_v301 = _dafny.Map.Empty.slice().updateUnsafe(_561_v300,_558_v298);
        let _563_v302;
        _563_v302 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_166_v4).length),_558_v298);
        let _564_v303;
        _564_v303 = _dafny.Map.Empty.slice().updateUnsafe((_162_v0)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length))],(((_563_v302).contains((_162_v0)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length))])) ? ((_563_v302).get((_162_v0)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length))])) : (_558_v298)));
        let _565_v304;
        let _nw97 = Array((new BigNumber(22)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = _558_v298;
        _nw97[(_dafny.ONE).toNumber()] = _558_v298;
        _nw97[(new BigNumber(2)).toNumber()] = (_560_v299)[_module.__default.safeIndex(new BigNumber((_324_v141).length), new BigNumber((_560_v299).length))];
        _nw97[(new BigNumber(3)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(4)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(5)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(6)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(7)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(8)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(9)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(10)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(11)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(12)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(13)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(14)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(15)).toNumber()] = (((_562_v301).contains(_561_v300)) ? ((_562_v301).get(_561_v300)) : ((((_563_v302).contains(_165_v3)) ? ((_563_v302).get(_165_v3)) : (_558_v298))));
        _nw97[(new BigNumber(16)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(17)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(18)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(19)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(20)).toNumber()] = _558_v298;
        _nw97[(new BigNumber(21)).toNumber()] = _558_v298;
        _565_v304 = _nw97;
        let _index67 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_565_v304).length));
        (_565_v304)[_index67] = _558_v298;
        let _index68 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_565_v304).length));
        (_565_v304)[_index68] = _558_v298;
        _163_v1 = ((_162_v0)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length))]).isLessThan(new BigNumber((_324_v141).length));
        let _566_v305;
        let _nw98 = Array((new BigNumber(24)).toNumber()).fill(false);
        _566_v305 = _nw98;
        let _index69 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_566_v305).length));
        (_566_v305)[_index69] = _module.__default.fm2(_163_v1, _167_globalState);
        let _index70 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_566_v305).length));
        (_566_v305)[_index70] = (_388_v190).f15;
        let _567_v306;
        let _init21 = ((_568_v300) => function (_569_i32) {
          return _568_v300;
        })(_561_v300);
        let _nw99 = Array((new BigNumber(21)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw99.length); _i0_21++) {
          _nw99[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _567_v306 = _nw99;
        let _570_v307;
        _570_v307 = _dafny.Seq.of(_567_v306, _567_v306, _567_v306);
        let _571_v308;
        _571_v308 = _module.D20.create_DC52((_570_v307)[_module.__default.safeIndex((_dafny.ZERO).minus(_165_v3), new BigNumber((_570_v307).length))]);
        let _572_v309;
        _572_v309 = _dafny.Map.Empty.slice().updateUnsafe((_162_v0)[_module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length))],_571_v308);
        let _573_v310;
        _573_v310 = _dafny.Map.Empty.slice().updateUnsafe(_572_v309,_165_v3);
        let _574_v312;
        _574_v312 = _module.D7.create_DC20(_dafny.MultiSet.fromElements(_165_v3));
        let _575_v313;
        _575_v313 = _dafny.Map.Empty.slice().updateUnsafe(_574_v312,_165_v3);
        let _576_v314;
        _576_v314 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(150)), function () {
          let _coll36 = new _dafny.Set();
          for (const _compr_36 of _dafny.IntegerRange(new BigNumber(730), new BigNumber(922))) {
            let _577_v311 = _compr_36;
            if (((new BigNumber(730)).isLessThanOrEqualTo(_577_v311)) && ((_577_v311).isLessThan(new BigNumber(922)))) {
              _coll36.add((_577_v311).minus(_165_v3));
            }
          }
          return _coll36;
        }())).length),_575_v313);
        let _index71 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length));
        let _rhs53 = new BigNumber(((((((_576_v314).contains(_165_v3)) ? ((_576_v314).get(_165_v3)) : (_dafny.Map.Empty.slice().updateUnsafe(_574_v312,(_388_v190).fm8((_388_v190).f15, (_388_v190).f15, _167_globalState))))).Merge(_575_v313)).Merge(_575_v313)).length);
        let _rhs54 = _573_v310;
        let _rhs55 = (_388_v190).fm8((_388_v190).f15, !((((_566_v305)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_566_v305).length))]) ? ((_388_v190).f15) : ((_388_v190).f15))), _167_globalState);
        let _lhs32 = _162_v0;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_162_v0).length));
        _lhs32[_lhs33] = _rhs53;
        _573_v310 = _rhs54;
        _165_v3 = _rhs55;
      } else {
        let _578_v315;
        let _579_v316;
        let _580_v317;
        let _581_v318;
        let _out50;
        let _out51;
        let _out52;
        let _out53;
        let _outcollector13 = (_388_v190).m3(_167_globalState);
        _out50 = _outcollector13[0];
        _out51 = _outcollector13[1];
        _out52 = _outcollector13[2];
        _out53 = _outcollector13[3];
        _578_v315 = _out50;
        _579_v316 = _out51;
        _580_v317 = _out52;
        _581_v318 = _out53;
        let _582_v319;
        _582_v319 = _dafny.Seq.of(_165_v3, _165_v3);
        let _583_v320;
        let _nw100 = new _module.C2();
        _nw100.__ctor(_582_v319, (_388_v190).f15, (_388_v190).f13, _dafny.Seq.update((_388_v190).f14, _module.__default.safeIndex(new BigNumber((_324_v141).length), new BigNumber(((_388_v190).f14).length)), !((_388_v190).f15)), _388_v190.f12);
        _583_v320 = _nw100;
        let _584_v321;
        let _out54;
        _out54 = (_583_v320).m6(_167_globalState);
        _584_v321 = _out54;
        let _585_v322;
        let _586_v323;
        let _587_v324;
        let _588_v325;
        let _out55;
        let _out56;
        let _out57;
        let _out58;
        let _outcollector14 = (_388_v190).m3(_167_globalState);
        _out55 = _outcollector14[0];
        _out56 = _outcollector14[1];
        _out57 = _outcollector14[2];
        _out58 = _outcollector14[3];
        _585_v322 = _out55;
        _586_v323 = _out56;
        _587_v324 = _out57;
        _588_v325 = _out58;
        let _589_v326;
        _589_v326 = new _dafny.CodePoint('d'.codePointAt(0));
        let _590_v327;
        let _nw101 = Array((new BigNumber(2)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = (_324_v141)[_module.__default.safeIndex(_579_v316, new BigNumber((_324_v141).length))];
        _nw101[(_dafny.ONE).toNumber()] = _589_v326;
        _590_v327 = _nw101;
        let _591_v328;
        _591_v328 = _dafny.Seq.of(_590_v327, _590_v327);
        let _592_v329;
        _592_v329 = _591_v328;
        let _593_v330;
        _593_v330 = _dafny.Seq.of(_592_v329, _592_v329, _592_v329, _592_v329, _591_v328);
        let _rhs56 = _324_v141;
        let _rhs57 = _593_v330;
        let _rhs58 = _324_v141;
        _324_v141 = _rhs56;
        _593_v330 = _rhs57;
        _324_v141 = _rhs58;
      }
      let _594_i33;
      _594_i33 = _dafny.ZERO;
      L3: {
        while (_163_v1) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_594_i33)) {
              break L3;
            }
            _594_i33 = (_594_i33).plus(_dafny.ONE);
            (_167_globalState).f4 = (_387_v189)[_module.__default.safeIndex(_165_v3, new BigNumber((_387_v189).length))];
            if (true) {
              let _595_v331;
              _595_v331 = _dafny.Map.Empty.slice().updateUnsafe((_388_v190).f15,_162_v0);
              let _rhs59 = new BigNumber((_325_v142).length);
              let _rhs60 = !(((_388_v190).f14)[_module.__default.safeIndex(new BigNumber((_595_v331).length), new BigNumber(((_388_v190).f14).length))]) || ((_163_v1) || ((_388_v190).f15));
              let _rhs61 = (new BigNumber((_dafny.Seq.UnicodeFromString("v")).length)).isLessThan(new BigNumber(23));
              let _rhs62 = _163_v1;
              let _lhs34 = _167_globalState;
              let _lhs35 = _167_globalState;
              _165_v3 = _rhs59;
              _lhs34.f4 = _rhs60;
              _lhs35.f4 = _rhs61;
              _163_v1 = _rhs62;
              let _596_v332;
              _596_v332 = _dafny.Map.Empty.slice().updateUnsafe((_325_v142)[_module.__default.safeIndex(new BigNumber(66), new BigNumber((_325_v142).length))],false);
              let _597_v333;
              _597_v333 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_165_v3, _165_v3, _165_v3, _165_v3)).length)), _165_v3);
              let _598_v334;
              let _nw102 = Array((new BigNumber(21)).toNumber());
              _nw102[(_dafny.ZERO).toNumber()] = _163_v1;
              _nw102[(_dafny.ONE).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(2)).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(3)).toNumber()] = (_596_v332).contains(_162_v0);
              _nw102[(new BigNumber(4)).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(5)).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(6)).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(7)).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(8)).toNumber()] = _163_v1;
              _nw102[(new BigNumber(9)).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(10)).toNumber()] = true;
              _nw102[(new BigNumber(11)).toNumber()] = (_388_v190).f15;
              _nw102[(new BigNumber(12)).toNumber()] = (!((_388_v190).f15)) || ((_388_v190).f15);
              _nw102[(new BigNumber(13)).toNumber()] = ((_388_v190).f15) || (!((_388_v190).f15));
              _nw102[(new BigNumber(14)).toNumber()] = (_597_v333).IsProperSubsetOf(_597_v333);
              _nw102[(new BigNumber(15)).toNumber()] = true;
              _nw102[(new BigNumber(16)).toNumber()] = _163_v1;
              _nw102[(new BigNumber(17)).toNumber()] = !(_163_v1);
              _nw102[(new BigNumber(18)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_324_v141, _324_v141);
              _nw102[(new BigNumber(19)).toNumber()] = _163_v1;
              _nw102[(new BigNumber(20)).toNumber()] = !((_388_v190).f15);
              _598_v334 = _nw102;
              let _index72 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_598_v334).length));
              (_598_v334)[_index72] = _163_v1;
              let _index73 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_598_v334).length));
              (_598_v334)[_index73] = (_388_v190).f15;
              let _599_v335;
              _599_v335 = _module.D12.create_DC37(_dafny.Seq.Concat(_324_v141, _324_v141), _163_v1);
              _599_v335 = _599_v335;
              let _600_v336;
              _600_v336 = _dafny.Map.Empty.slice().updateUnsafe((_388_v190).f14,(_166_v4).Merge(_dafny.Map.Empty.slice().updateUnsafe(_165_v3,_165_v3)));
              _600_v336 = _600_v336;
              let _601_v337;
              _601_v337 = new _dafny.CodePoint('o'.codePointAt(0));
              _601_v337 = (_324_v141)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber(140), _165_v3), new BigNumber((_324_v141).length))];
            } else {
              let _602_v338;
              _602_v338 = _dafny.MultiSet.fromElements(_165_v3);
              let _603_v339;
              _603_v339 = _dafny.Map.Empty.slice().updateUnsafe(_602_v338,_165_v3);
              _603_v339 = _603_v339;
              _163_v1 = !((_209_v40).Difference(_209_v40)).contains((_388_v190).f15);
              let _604_v340;
              let _nw103 = new _module.C1();
              _nw103.__ctor((_388_v190).f15, _163_v1, _module.__default.fm60(_165_v3, _167_globalState), (_388_v190).f14, _385_v187);
              _604_v340 = _nw103;
              let _605_v341;
              _605_v341 = _module.D11.create_DC33((_388_v190).f15, new BigNumber(766), (_388_v190).f15, _604_v340, (_388_v190).f15);
              _605_v341 = _605_v341;
              let _606_v342;
              let _607_v343;
              let _608_v344;
              let _609_v345;
              let _out59;
              let _out60;
              let _out61;
              let _out62;
              let _outcollector15 = (_388_v190).m3(_167_globalState);
              _out59 = _outcollector15[0];
              _out60 = _outcollector15[1];
              _out61 = _outcollector15[2];
              _out62 = _outcollector15[3];
              _606_v342 = _out59;
              _607_v343 = _out60;
              _608_v344 = _out61;
              _609_v345 = _out62;
              let _610_v346;
              _610_v346 = _dafny.Seq.of(_607_v343, _165_v3, _165_v3);
              let _611_v347;
              _611_v347 = _dafny.Set.fromElements(new BigNumber((_386_v188).length), _607_v343, _607_v343);
              let _612_v349;
              let _nw104 = new _module.C6();
              _nw104.__ctor(_610_v346, !(_module.__default.fm70((_388_v190).f15, _324_v141, _165_v3, (_388_v190).f15, _167_globalState)).equals(function () {
                let _coll37 = new _dafny.Set();
                for (const _compr_37 of (_611_v347).Elements) {
                  let _613_v348 = _compr_37;
                  if ((_611_v347).contains(_613_v348)) {
                    _coll37.add(_module.__default.safeDivisionInt(_613_v348, new BigNumber(22)));
                  }
                }
                return _coll37;
              }()), ((_388_v190).f13).Merge((_388_v190).f13), _dafny.Seq.of((_388_v190).f15, (_388_v190).f15), _385_v187);
              _612_v349 = _nw104;
            }
            _165_v3 = _165_v3;
            let _614_v350;
            let _nw105 = Array((new BigNumber(21)).toNumber()).fill(false);
            _614_v350 = _nw105;
            let _615_v351;
            _615_v351 = _dafny.Seq.of(_614_v350);
            let _616_v352;
            _616_v352 = _dafny.Map.Empty.slice().updateUnsafe((_615_v351)[_module.__default.safeIndex(new BigNumber(966), new BigNumber((_615_v351).length))],(_388_v190).f15);
            _616_v352 = (_616_v352).update(_614_v350, (_388_v190).f15);
          }
        }
      }
      let _617_v353;
      _617_v353 = new _dafny.CodePoint('b'.codePointAt(0));
      let _618_v354;
      _618_v354 = _dafny.Map.Empty.slice().updateUnsafe(_617_v353,(_388_v190).f15);
      let _619_v355;
      _619_v355 = _dafny.MultiSet.fromElements(_618_v354);
      let _620_i34;
      _620_i34 = _dafny.ZERO;
      L4: {
        while (_module.__default.fm2((_619_v355).IsDisjointFrom(_619_v355), _167_globalState)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_620_i34)) {
              break L4;
            }
            _620_i34 = (_620_i34).plus(_dafny.ONE);
            let _621_v356;
            _621_v356 = _module.D4.create_DC12(_dafny.Seq.of(true));
            let _source7 = _621_v356;
            if (_source7.is_DC13) {
              let _622___mcc_h19 = (_source7).cf19;
              let _623___mcc_h20 = (_source7).cf20;
              let _624___mcc_h21 = (_source7).cf21;
              let _625___mcc_h22 = (_source7).cf22;
              let _626___mcc_h23 = (_source7).cf23;
              let _627_cf23 = _626___mcc_h23;
              let _628_cf22 = _625___mcc_h22;
              let _629_cf21 = _624___mcc_h21;
              let _630_cf20 = _623___mcc_h20;
              let _631_cf19 = _622___mcc_h19;
              let _632_v357;
              _632_v357 = _module.D0.create_DC2(_628_cf22, (_388_v190).f15);
              let _633_v358;
              _633_v358 = _dafny.Set.fromElements(_324_v141, _324_v141);
              let _634_v359;
              let _635_v360;
              let _636_v361;
              let _637_v362;
              let _out63;
              let _out64;
              let _out65;
              let _out66;
              let _outcollector16 = _module.__default.m0(_630_cf20, _632_v357, (_633_v358).IsDisjointFrom(_633_v358), _167_globalState);
              _out63 = _outcollector16[0];
              _out64 = _outcollector16[1];
              _out65 = _outcollector16[2];
              _out66 = _outcollector16[3];
              _634_v359 = _out63;
              _635_v360 = _out64;
              _636_v361 = _out65;
              _637_v362 = _out66;
              _323_v140 = (_323_v140).update(_module.__default.fm0(!(_631_cf19), _323_v140, _635_v360, _167_globalState), _322_v139);
              _628_cf22 = _module.__default.safeModuloInt(_629_cf21, _629_cf21);
              let _638_v363;
              _638_v363 = _dafny.Map.Empty.slice().updateUnsafe(_388_v190,_324_v141);
              _324_v141 = (((_638_v363).contains(_388_v190)) ? ((_638_v363).get(_388_v190)) : (_324_v141));
            } else if (_source7.is_DC14) {
              let _639___mcc_h24 = (_source7).cf24;
              let _640_cf24 = _639___mcc_h24;
              _324_v141 = _dafny.Seq.Concat(_324_v141, _dafny.Seq.Create(_module.__default.abs(new BigNumber(594)), ((_641_v353) => function (_642_i35) {
                return _641_v353;
              })(_617_v353)));
              let _643_v364;
              let _644_v365;
              let _645_v366;
              let _646_v367;
              let _out67;
              let _out68;
              let _out69;
              let _out70;
              let _outcollector17 = (_388_v190).m3(_167_globalState);
              _out67 = _outcollector17[0];
              _out68 = _outcollector17[1];
              _out69 = _outcollector17[2];
              _out70 = _outcollector17[3];
              _643_v364 = _out67;
              _644_v365 = _out68;
              _645_v366 = _out69;
              _646_v367 = _out70;
              let _647_v368;
              let _nw106 = new _module.C2();
              _nw106.__ctor((_388_v190).fm6(_324_v141, _167_globalState), _646_v367, (_dafny.Map.Empty.slice().updateUnsafe(true,_386_v188)).Merge((_388_v190).f13), (_388_v190).f14, _388_v190.f12);
              _647_v368 = _nw106;
              _647_v368 = _647_v368;
              let _648_v369;
              _648_v369 = _module.D0.create_DC2(_165_v3, (_388_v190).f15);
              let _649_v370;
              let _650_v371;
              let _651_v372;
              let _652_v373;
              let _out71;
              let _out72;
              let _out73;
              let _out74;
              let _outcollector18 = _module.__default.m0((_388_v190).f15, _648_v369, (_388_v190).f15, _167_globalState);
              _out71 = _outcollector18[0];
              _out72 = _outcollector18[1];
              _out73 = _outcollector18[2];
              _out74 = _outcollector18[3];
              _649_v370 = _out71;
              _650_v371 = _out72;
              _651_v372 = _out73;
              _652_v373 = _out74;
            } else {
              let _653___mcc_h25 = (_source7).cf18;
              let _654_cf18 = _653___mcc_h25;
              let _655_v374;
              let _nw107 = new _module.C3();
              _nw107.__ctor(_163_v1, _388_v190.f12, (_388_v190).f15, (_388_v190).f13, (_388_v190).f14);
              _655_v374 = _nw107;
              let _656_v375;
              _656_v375 = _dafny.Seq.of(_655_v374);
              let _657_v376;
              _657_v376 = _dafny.Map.Empty.slice().updateUnsafe(_656_v375,_324_v141);
              _163_v1 = _dafny.areEqual((((_657_v376).contains(_656_v375)) ? ((_657_v376).get(_656_v375)) : (_324_v141)), _324_v141);
              let _658_v377;
              _658_v377 = _dafny.Seq.of(_165_v3);
              (_167_globalState).f4 = !(_dafny.areEqual(_658_v377, _658_v377));
              (_167_globalState).f4 = (_655_v374).f20;
              let _659_v378;
              let _nw108 = new _module.C5();
              _nw108.__ctor((_388_v190).f15, _385_v187, (_388_v190).f13, _dafny.Seq.Concat(_387_v189, _654_cf18));
              _659_v378 = _nw108;
            }
            let _660_v379;
            _660_v379 = _module.D5.create_DC16(_163_v1);
            let _661_v380;
            _661_v380 = _dafny.Set.fromElements(_660_v379);
            let _662_v381;
            _662_v381 = _dafny.Seq.of(_660_v379);
            let _663_v383;
            _663_v383 = _dafny.Seq.of(_661_v380, (_661_v380).Intersect(_dafny.Set.fromElements(_660_v379)), _dafny.Set.fromElements(_660_v379, _660_v379), (_661_v380).Difference(function () {
              let _coll38 = new _dafny.Set();
              for (const _compr_38 of (_662_v381).Elements) {
                let _664_v382 = _compr_38;
                if (_dafny.Seq.contains(_662_v381, _664_v382)) {
                  _coll38.add(_664_v382);
                }
              }
              return _coll38;
            }()));
            _663_v383 = _663_v383;
            if (!(true)) {
              let _665_v384;
              _665_v384 = _dafny.Map.Empty.slice().updateUnsafe(_162_v0,(_388_v190).f15);
              let _666_v385;
              let _nw109 = new _module.C1();
              _nw109.__ctor((_388_v190).f15, true, (_388_v190).f13, ((_163_v1) ? ((_388_v190).f14) : (_387_v189)), (((((_665_v384).contains(_162_v0)) ? ((_665_v384).get(_162_v0)) : ((_388_v190).f15))) ? (_388_v190.f12) : (_388_v190.f12)));
              _666_v385 = _nw109;
              let _667_v386;
              _667_v386 = _dafny.Map.Empty.slice().updateUnsafe(_162_v0,_165_v3);
              let _668_v387;
              _668_v387 = _dafny.Map.Empty.slice().updateUnsafe((_666_v385).fm5(_163_v1, _167_globalState),(_667_v386).contains(_162_v0));
              _668_v387 = (_668_v387).update(_165_v3, (_666_v385).f17);
              _165_v3 = (_165_v3).plus((new BigNumber(438)).multipliedBy(_165_v3));
              let _669_v388;
              let _nw110 = new _module.C6();
              _nw110.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(353)), ((_670_v139, _671_v3) => function (_672_i36) {
                return (_module.D20.create_DC53(_670_v139, _671_v3, _671_v3)).dtor_cf79;
              })(_322_v139, _165_v3)), (_388_v190).f15, (_388_v190).f13, (_388_v190).f14, _388_v190.f12);
              _669_v388 = _nw110;
              let _673_v389;
              _673_v389 = _dafny.Map.Empty.slice().updateUnsafe(_669_v388,(_388_v190).f15);
              let _674_v390;
              _674_v390 = _dafny.Seq.of(_165_v3, new BigNumber((_673_v389).length));
              let _675_v391;
              _675_v391 = _dafny.Set.fromElements(_165_v3, _module.__default.safeModuloInt(_165_v3, new BigNumber((_674_v390).length)));
              _675_v391 = _675_v391;
              _165_v3 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(new BigNumber((_164_v2).cardinality()), _165_v3)).minus(_165_v3));
            } else {
              let _676_v392;
              let _nw111 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
              _676_v392 = _nw111;
              let _index74 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_676_v392).length));
              (_676_v392)[_index74] = ((_386_v188).update((_388_v190).f15, (_388_v190).f15)).update((_388_v190).f15, _module.__default.fm2((_388_v190).f15, _167_globalState));
              let _677_v393;
              let _nw112 = Array((new BigNumber(14)).toNumber());
              _677_v393 = _nw112;
              let _678_v394;
              let _nw113 = new _module.C5();
              _nw113.__ctor((_388_v190).f15, _388_v190.f12, (_388_v190).f13, (_388_v190).f14);
              _678_v394 = _nw113;
              let _index75 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_677_v393).length));
              (_677_v393)[_index75] = _678_v394;
              let _index76 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_676_v392).length));
              let _index77 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_677_v393).length));
              let _rhs63 = _163_v1;
              let _rhs64 = _386_v188;
              let _rhs65 = _165_v3;
              let _rhs66 = _678_v394;
              let _lhs36 = _167_globalState;
              let _lhs37 = _676_v392;
              let _lhs38 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_676_v392).length));
              let _lhs39 = _677_v393;
              let _lhs40 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_677_v393).length));
              _lhs36.f4 = _rhs63;
              _lhs37[_lhs38] = _rhs64;
              _165_v3 = _rhs65;
              _lhs39[_lhs40] = _rhs66;
              _165_v3 = _165_v3;
              let _679_v395;
              _679_v395 = _dafny.MultiSet.fromElements(_165_v3);
              let _680_v396;
              _680_v396 = _module.D0.create_DC2(_165_v3, _163_v1);
              let _681_v397;
              let _682_v398;
              let _683_v399;
              let _684_v400;
              let _out75;
              let _out76;
              let _out77;
              let _out78;
              let _outcollector19 = _module.__default.m0((_679_v395).equals(_679_v395), _680_v396, (_388_v190).f15, _167_globalState);
              _out75 = _outcollector19[0];
              _out76 = _outcollector19[1];
              _out77 = _outcollector19[2];
              _out78 = _outcollector19[3];
              _681_v397 = _out75;
              _682_v398 = _out76;
              _683_v399 = _out77;
              _684_v400 = _out78;
              _684_v400 = false;
              let _index78 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_162_v0).length));
              (_162_v0)[_index78] = _683_v399;
              let _index79 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_162_v0).length));
              (_162_v0)[_index79] = (_388_v190).fm7(_683_v399, _167_globalState);
            }
            let _685_v401;
            _685_v401 = _module.D8.create_DC26((_388_v190).f15);
            let _686_v402;
            let _nw114 = Array((new BigNumber(6)).toNumber());
            _nw114[(_dafny.ZERO).toNumber()] = _685_v401;
            _nw114[(_dafny.ONE).toNumber()] = _685_v401;
            _nw114[(new BigNumber(2)).toNumber()] = _module.D8.create_DC26(_163_v1);
            _nw114[(new BigNumber(3)).toNumber()] = _685_v401;
            _nw114[(new BigNumber(4)).toNumber()] = _685_v401;
            _nw114[(new BigNumber(5)).toNumber()] = _685_v401;
            _686_v402 = _nw114;
            let _index80 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_686_v402).length));
            (_686_v402)[_index80] = _module.D8.create_DC26((_388_v190).f15);
            let _pat_let_tv1 = _388_v190;
            let _pat_let_tv2 = _163_v1;
            let _index81 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_686_v402).length));
            (_686_v402)[_index81] = function (_pat_let2_0) {
              return function (_687_dt__update__tmp_h2) {
                return function (_pat_let3_0) {
                  return function (_688_dt__update_hcf41_h0) {
                    return _module.D8.create_DC26(_688_dt__update_hcf41_h0);
                  }(_pat_let3_0);
                }((!((_pat_let_tv1).f15)) && (_pat_let_tv2));
              }(_pat_let2_0);
            }(_685_v401);
          }
        }
      }
      let _689_v403;
      let _nw115 = Array((new BigNumber(6)).toNumber()).fill(false);
      _689_v403 = _nw115;
      let _690_v404;
      _690_v404 = _dafny.MultiSet.fromElements(_689_v403, _689_v403);
      let _index82 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_689_v403).length));
      (_689_v403)[_index82] = (_690_v404).IsProperSubsetOf(_dafny.MultiSet.fromElements(_689_v403));
      let _index83 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_689_v403).length));
      (_689_v403)[_index83] = (_165_v3).isEqualTo(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_324_v141, _324_v141), _module.__default.safeIndex(_165_v3, new BigNumber((_dafny.Seq.Concat(_324_v141, _324_v141)).length)), _617_v353)).length));
      let _691_v405;
      let _nw116 = new _module.C2();
      _nw116.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(342)), ((_692_v141) => function (_693_i37) {
        return new BigNumber((_692_v141).length);
      })(_324_v141)), ((_388_v190).f15) === ((_388_v190).f15), (_388_v190).f13, (_388_v190).f14, _388_v190.f12);
      _691_v405 = _nw116;
      _691_v405 = _691_v405;
      let _694_v406;
      let _nw117 = new _module.C3();
      _nw117.__ctor((_388_v190).f15, _388_v190.f12, false, _dafny.Map.Empty.slice().updateUnsafe(false,_386_v188), _387_v189);
      _694_v406 = _nw117;
      let _695_v407;
      _695_v407 = _dafny.MultiSet.fromElements(_694_v406);
      let _696_v408;
      _696_v408 = _dafny.Seq.of(_165_v3, (((_695_v407).contains(_694_v406)) ? ((_695_v407).get(_694_v406)) : (_165_v3)), _165_v3);
      let _697_v409;
      _697_v409 = _dafny.Map.Empty.slice().updateUnsafe(_165_v3,_696_v408);
      let _698_v410;
      _698_v410 = _module.D5.create_DC15(_696_v408);
      let _699_v411;
      _699_v411 = _dafny.Map.Empty.slice().updateUnsafe(_163_v1,(_698_v410).dtor_cf25);
      let _700_v412;
      let _nw118 = Array((new BigNumber(12)).toNumber());
      _nw118[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_696_v408, _696_v408);
      _nw118[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_696_v408)[_module.__default.safeIndex(_165_v3, new BigNumber((_696_v408).length))], _165_v3);
      _nw118[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_696_v408, _696_v408);
      _nw118[(new BigNumber(3)).toNumber()] = _696_v408;
      _nw118[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_696_v408, _dafny.Seq.of(_165_v3, new BigNumber((_dafny.Seq.UnicodeFromString("rkdkja")).length)));
      _nw118[(new BigNumber(5)).toNumber()] = _696_v408;
      _nw118[(new BigNumber(6)).toNumber()] = _696_v408;
      _nw118[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_696_v408, (_module.D5.create_DC15((((_697_v409).contains(_165_v3)) ? ((_697_v409).get(_165_v3)) : ((((_699_v411).contains(_163_v1)) ? ((_699_v411).get(_163_v1)) : (_dafny.Seq.of(_165_v3))))))).dtor_cf25);
      _nw118[(new BigNumber(8)).toNumber()] = (_694_v406).fm6(_324_v141, _167_globalState);
      _nw118[(new BigNumber(9)).toNumber()] = _696_v408;
      _nw118[(new BigNumber(10)).toNumber()] = _696_v408;
      _nw118[(new BigNumber(11)).toNumber()] = _696_v408;
      _700_v412 = _nw118;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_700_v412).length))) {
        let _701_i38 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_701_i38)) && ((_701_i38).isLessThan(new BigNumber((_700_v412).length))))) {
          (_700_v412)[(_701_i38)] = _dafny.Seq.Concat(_dafny.Seq.Concat(_696_v408, _696_v408), _696_v408);
        }
      }
      process.stdout.write(_dafny.toString((_162_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_163_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v2).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_165_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(846),new BigNumber(846)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_167_globalState).f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_167_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_167_globalState).f5).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_167_globalState).f10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(846),new BigNumber(846)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_209_v40).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_301_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v139).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_323_v140).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(3))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_324_v141).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_325_v142).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_326_v143).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_386_v188).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_387_v189, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_388_v190).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_388_v190).f13).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(false,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_388_v190).f14, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_389_v191).dtor_cf63).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_389_v191).dtor_cf63).f13).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(false,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_389_v191).dtor_cf63).f14, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_545_i27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_594_i33));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_617_v353));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_618_v354).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_619_v355).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_620_i34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_689_v403)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_690_v404).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_694_v406).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_695_v407).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_696_v408, _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_697_v409).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),_dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_698_v410).dtor_cf25, _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_699_v411).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[_dafny.ONE], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(2)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(3)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(4)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(5)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(6)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(7)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(8)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(9)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(10)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_700_v412)[new BigNumber(11)], _dafny.Seq.of(new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3), new BigNumber(3), _dafny.ONE, new BigNumber(3)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC2(cf1, cf2) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
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
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
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
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC3(cf3) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
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
      return _dafny.MultiSet.Empty;
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
    static create_DC5(cf5, cf6, cf7) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC6(cf8, cf9) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC4(cf4) {
      let $dt = new D2(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC7(cf10) {
      let $dt = new D2(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6 && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf8 === other.cf8 && this.cf9 === other.cf9;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(_dafny.ZERO, false, false);
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
    static create_DC9(cf12, cf13) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC10(cf14, cf15, cf16) {
      let $dt = new D3(1);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC8(cf11) {
      let $dt = new D3(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D3(3);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC13(cf19, cf20, cf21, cf22, cf23) {
      let $dt = new D4(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC14(cf24) {
      let $dt = new D4(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D4(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf24) + ")";
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
        return other.$tag === 0 && this.cf19 === other.cf19 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(false, false, _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC16(cf26) {
      let $dt = new D5(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC15(cf25) {
      let $dt = new D5(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC16(false);
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
    static create_DC18(cf28) {
      let $dt = new D6(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC17(cf27) {
      let $dt = new D6(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC19(cf29) {
      let $dt = new D6(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18(_dafny.ZERO);
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
    static create_DC21(cf31, cf32, cf33) {
      let $dt = new D7(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC22(cf34) {
      let $dt = new D7(1);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC20(cf30) {
      let $dt = new D7(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && this.cf33 === other.cf33;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf34 === other.cf34;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC21(_dafny.Map.Empty, false, false);
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
    static create_DC24(cf36) {
      let $dt = new D8(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC25(cf37, cf38, cf39, cf40) {
      let $dt = new D8(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf41) {
      let $dt = new D8(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC23(cf35) {
      let $dt = new D8(3);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC27(cf42) {
      let $dt = new D8(4);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get is_DC27() { return this.$tag === 4; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC27" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39) && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf41 === other.cf41;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC24(false);
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
    static create_DC28(cf43) {
      let $dt = new D9(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf43) + ")";
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
      return [];
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
    static create_DC29(cf44) {
      let $dt = new D10(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return new _dafny.CodePoint('D'.codePointAt(0));
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
    static create_DC31() {
      let $dt = new D11(0);
      return $dt;
    }
    static create_DC32(cf46, cf47) {
      let $dt = new D11(1);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC33(cf48, cf49, cf50, cf51, cf52) {
      let $dt = new D11(2);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC30(cf45) {
      let $dt = new D11(3);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get is_DC30() { return this.$tag === 3; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31";
      } else if (this.$tag === 1) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf45) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf48 === other.cf48 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50 && this.cf51 === other.cf51 && this.cf52 === other.cf52;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31();
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
    static create_DC35(cf54) {
      let $dt = new D12(0);
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC36(cf55, cf56, cf57, cf58) {
      let $dt = new D12(1);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC37(cf59, cf60) {
      let $dt = new D12(2);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC34(cf53) {
      let $dt = new D12(3);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC38(cf61) {
      let $dt = new D12(4);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC34() { return this.$tag === 3; }
    get is_DC38() { return this.$tag === 4; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC37" + "(" + this.cf59.toVerbatimString(true) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 4) {
        return "D12.DC38" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf54 === other.cf54;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf55 === other.cf55 && _dafny.areEqual(this.cf56, other.cf56) && _dafny.areEqual(this.cf57, other.cf57) && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC35(false);
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
    static create_DC39(cf62) {
      let $dt = new D13(0);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC39" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf62 === other.cf62;
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41(cf64, cf65) {
      let $dt = new D14(0);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC42() {
      let $dt = new D14(1);
      return $dt;
    }
    static create_DC40(cf63) {
      let $dt = new D14(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC43(cf66) {
      let $dt = new D14(3);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get is_DC43() { return this.$tag === 3; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC41" + "(" + this.cf64.toVerbatimString(true) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC42";
      } else if (this.$tag === 2) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC43" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf63 === other.cf63;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC41(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
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
    static create_DC44(cf67) {
      let $dt = new D15(0);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf67) + ")";
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
      return _dafny.Set.Empty;
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
    static create_DC45(cf68) {
      let $dt = new D16(0);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf68 === other.cf68;
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
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC47(cf70) {
      let $dt = new D17(0);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC48(cf71, cf72) {
      let $dt = new D17(1);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC49(cf73) {
      let $dt = new D17(2);
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC46(cf69) {
      let $dt = new D17(3);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC49() { return this.$tag === 2; }
    get is_DC46() { return this.$tag === 3; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC47" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC48" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC49" + "(" + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf70 === other.cf70;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf71 === other.cf71 && this.cf72 === other.cf72;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf73 === other.cf73;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf69, other.cf69);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC47(false);
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
    static create_DC50(cf74) {
      let $dt = new D18(0);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74;
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
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC51(cf75) {
      let $dt = new D19(0);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC51" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf75 === other.cf75;
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
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf77, cf78, cf79) {
      let $dt = new D20(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC54(cf80, cf81, cf82, cf83) {
      let $dt = new D20(1);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC52(cf76) {
      let $dt = new D20(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
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
        return "D20.DC53" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC54" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC52" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80) && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82) && _dafny.areEqual(this.cf83, other.cf83);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf76 === other.cf76;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC53(_dafny.Map.Empty, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC56(cf85) {
      let $dt = new D21(0);
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC57(cf86, cf87, cf88) {
      let $dt = new D21(1);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC58(cf89, cf90) {
      let $dt = new D21(2);
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC55(cf84) {
      let $dt = new D21(3);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC58() { return this.$tag === 2; }
    get is_DC55() { return this.$tag === 3; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC56" + "(" + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC57" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC58" + "(" + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf85 === other.cf85;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf86, other.cf86) && this.cf87 === other.cf87 && this.cf88 === other.cf88;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf89 === other.cf89 && _dafny.areEqual(this.cf90, other.cf90);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf84, other.cf84);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC56(null);
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
    static create_DC60(cf92, cf93, cf94) {
      let $dt = new D22(0);
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC61(cf95, cf96) {
      let $dt = new D22(1);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC62(cf97) {
      let $dt = new D22(2);
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC59(cf91) {
      let $dt = new D22(3);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC61() { return this.$tag === 1; }
    get is_DC62() { return this.$tag === 2; }
    get is_DC59() { return this.$tag === 3; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC60" + "(" + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC61" + "(" + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC62" + "(" + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 3) {
        return "D22.DC59" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf92 === other.cf92 && this.cf93 === other.cf93 && this.cf94 === other.cf94;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf95, other.cf95) && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf97 === other.cf97;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf91 === other.cf91;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC60(false, false, false);
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
    static create_DC64() {
      let $dt = new D23(0);
      return $dt;
    }
    static create_DC63(cf98) {
      let $dt = new D23(1);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC64() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC64";
      } else if (this.$tag === 1) {
        return "D23.DC63" + "(" + _dafny.toString(this.cf98) + ")";
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
        return other.$tag === 1 && this.cf98 === other.cf98;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC64();
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
    static create_DC66(cf100) {
      let $dt = new D24(0);
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC67(cf101, cf102, cf103) {
      let $dt = new D24(1);
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC65(cf99) {
      let $dt = new D24(2);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC68(cf104) {
      let $dt = new D24(3);
      $dt.cf104 = cf104;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get is_DC67() { return this.$tag === 1; }
    get is_DC65() { return this.$tag === 2; }
    get is_DC68() { return this.$tag === 3; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf104() { return this.cf104; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC66" + "(" + _dafny.toString(this.cf100) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC67" + "(" + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC65" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 3) {
        return "D24.DC68" + "(" + _dafny.toString(this.cf104) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf100 === other.cf100;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf101 === other.cf101 && _dafny.areEqual(this.cf102, other.cf102) && _dafny.areEqual(this.cf103, other.cf103);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf99 === other.cf99;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf104, other.cf104);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC66(false);
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
    static create_DC69(cf105) {
      let $dt = new D25(0);
      $dt.cf105 = cf105;
      return $dt;
    }
    get is_DC69() { return this.$tag === 0; }
    get dtor_cf105() { return this.cf105; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC69" + "(" + _dafny.toString(this.cf105) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf105, other.cf105);
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
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC71(cf107, cf108, cf109, cf110, cf111) {
      let $dt = new D26(0);
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC70(cf106) {
      let $dt = new D26(1);
      $dt.cf106 = cf106;
      return $dt;
    }
    static create_DC72(cf112) {
      let $dt = new D26(2);
      $dt.cf112 = cf112;
      return $dt;
    }
    get is_DC71() { return this.$tag === 0; }
    get is_DC70() { return this.$tag === 1; }
    get is_DC72() { return this.$tag === 2; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf112() { return this.cf112; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC71" + "(" + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ", " + _dafny.toString(this.cf109) + ", " + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC70" + "(" + _dafny.toString(this.cf106) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC72" + "(" + _dafny.toString(this.cf112) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf107, other.cf107) && this.cf108 === other.cf108 && this.cf109 === other.cf109 && _dafny.areEqual(this.cf110, other.cf110) && this.cf111 === other.cf111;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf106, other.cf106);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf112, other.cf112);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC71(_dafny.ZERO, false, [], _dafny.Set.Empty, null);
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
    static create_DC74(cf114) {
      let $dt = new D27(0);
      $dt.cf114 = cf114;
      return $dt;
    }
    static create_DC75(cf115, cf116) {
      let $dt = new D27(1);
      $dt.cf115 = cf115;
      $dt.cf116 = cf116;
      return $dt;
    }
    static create_DC73(cf113) {
      let $dt = new D27(2);
      $dt.cf113 = cf113;
      return $dt;
    }
    get is_DC74() { return this.$tag === 0; }
    get is_DC75() { return this.$tag === 1; }
    get is_DC73() { return this.$tag === 2; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf113() { return this.cf113; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC74" + "(" + _dafny.toString(this.cf114) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC75" + "(" + this.cf115.toVerbatimString(true) + ", " + _dafny.toString(this.cf116) + ")";
      } else if (this.$tag === 2) {
        return "D27.DC73" + "(" + _dafny.toString(this.cf113) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf114, other.cf114);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf115, other.cf115) && this.cf116 === other.cf116;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf113, other.cf113);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC74(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
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

  $module.T4 = class T4 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f4 = false;
      this._f0 = false;
      this._f1 = [];
      this._f2 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f5 = _dafny.MultiSet.Empty;
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.ZERO;
      this._f8 = false;
      this._f9 = false;
      this._f10 = _dafny.Map.Empty;
      this._f11 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
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
    get f11() {
      let _this = this;
      return _this._f11;
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
    fm14(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fihm"), _dafny.Seq.UnicodeFromString("ajtstts")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(166)), function (_702_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }))).length);
    };
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.ZERO;
      let _703_v0;
      let _init22 = ((_704_p2, _705_p0) => function (_706_i0) {
        return (_704_p2).isLessThanOrEqualTo(_705_p0);
      })(p2, p0);
      let _nw119 = Array((new BigNumber(3)).toNumber());
      for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw119.length); _i0_22++) {
        _nw119[_i0_22] = _init22(new BigNumber(_i0_22));
      }
      _703_v0 = _nw119;
      let _707_v1;
      _707_v1 = true;
      let _index84 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length));
      (_703_v0)[_index84] = _707_v1;
      let _index85 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length));
      (_703_v0)[_index85] = _707_v1;
      let _708_v2;
      _708_v2 = _dafny.Map.Empty.slice().updateUnsafe(_707_v1,false);
      let _709_v3;
      _709_v3 = _dafny.Seq.UnicodeFromString("juqts");
      let _710_v4;
      _710_v4 = _dafny.Map.Empty.slice().updateUnsafe(_709_v3,p3);
      let _711_v5;
      _711_v5 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0));
      _708_v2 = (_708_v2).update(_707_v1, (_711_v5).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_710_v4).length), (_dafny.ZERO).minus(new BigNumber((_711_v5).length)))));
      let _712_v6;
      _712_v6 = _module.D0.create_DC2((new BigNumber((_dafny.MultiSet.fromElements(p2, (_this).fm14(p0, globalState))).cardinality())).plus(new BigNumber((_708_v2).length)), (_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))]);
      let _source8 = _712_v6;
      if (_source8.is_DC1) {
        (globalState).f4 = _707_v1;
        r1 = _module.__default.safeDivisionInt(p2, p2);
        _707_v1 = !(!((p0).isLessThanOrEqualTo((_dafny.ZERO).minus(p0))) || (!(_707_v1) || (_707_v1)));
        _709_v3 = ((_707_v1) ? (_709_v3) : (_dafny.Seq.update(_709_v3, _module.__default.safeIndex(p0, new BigNumber((_709_v3).length)), p3)));
      } else if (_source8.is_DC2) {
        let _713___mcc_h0 = (_source8).cf1;
        let _714___mcc_h1 = (_source8).cf2;
        let _715_cf2 = _714___mcc_h1;
        let _716_cf1 = _713___mcc_h0;
        let _717_v7;
        let _nw120 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
        _717_v7 = _nw120;
        let _718_v8;
        _718_v8 = _dafny.Seq.of(_707_v1, _715_cf2);
        let _719_v9;
        _719_v9 = _dafny.Seq.of(_718_v8, _718_v8);
        let _index86 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_717_v7).length));
        (_717_v7)[_index86] = _719_v9;
        let _720_v10;
        _720_v10 = _module.D2.create_DC4(_719_v9);
        let _721_v11;
        _721_v11 = _dafny.MultiSet.fromElements(_707_v1);
        let _722_v12;
        _722_v12 = _dafny.Seq.of(new BigNumber((_721_v11).cardinality()));
        let _index87 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_717_v7).length));
        (_717_v7)[_index87] = _dafny.Seq.update(_dafny.Seq.update((_720_v10).dtor_cf4, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p0, new BigNumber(585)), _722_v12)).length), new BigNumber(((_720_v10).dtor_cf4).length)), _module.__default.fm15((_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))], p2, p0, p2, globalState)), _module.__default.safeIndex(new BigNumber((_709_v3).length), new BigNumber((_dafny.Seq.update((_720_v10).dtor_cf4, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p0, new BigNumber(585)), _722_v12)).length), new BigNumber(((_720_v10).dtor_cf4).length)), _module.__default.fm15((_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))], p2, p0, p2, globalState))).length)), _dafny.Seq.update(_dafny.Seq.Concat(_718_v8, _dafny.Seq.of(true)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat(_718_v8, _dafny.Seq.of(true))).length)), _707_v1));
        (globalState).f4 = !(_715_cf2);
        let _723_v13;
        let _nw121 = Array((new BigNumber(15)).toNumber()).fill([]);
        _723_v13 = _nw121;
        let _index88 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_723_v13).length));
        (_723_v13)[_index88] = _703_v0;
        let _index89 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_723_v13).length));
        (_723_v13)[_index89] = _703_v0;
        let _724_v14;
        _724_v14 = (_721_v11).update(false, _module.__default.abs(new BigNumber((_721_v11).cardinality())));
        let _725_v15;
        let _nw122 = Array((new BigNumber(16)).toNumber());
        _nw122[(_dafny.ZERO).toNumber()] = _724_v14;
        _nw122[(_dafny.ONE).toNumber()] = _724_v14;
        _nw122[(new BigNumber(2)).toNumber()] = _721_v11;
        _nw122[(new BigNumber(3)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(4)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_715_cf2);
        _nw122[(new BigNumber(6)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(7)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(8)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(9)).toNumber()] = ((false) ? (_721_v11) : (_724_v14));
        _nw122[(new BigNumber(10)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(11)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(12)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(13)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(14)).toNumber()] = _724_v14;
        _nw122[(new BigNumber(15)).toNumber()] = _724_v14;
        _725_v15 = _nw122;
        _725_v15 = _725_v15;
      } else {
        let _726___mcc_h2 = (_source8).cf0;
        let _727_cf0 = _726___mcc_h2;
        let _index90 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length));
        (_703_v0)[_index90] = (_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))];
        let _728_v16;
        _728_v16 = _dafny.Map.Empty.slice().updateUnsafe(p2,_709_v3);
        let _729_v17;
        _729_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_728_v16).length),(_711_v5).Difference(_711_v5));
        let _730_v18;
        _730_v18 = _dafny.MultiSet.fromElements(_707_v1);
        _729_v17 = (_729_v17).update(new BigNumber(((_730_v18).Difference(_dafny.MultiSet.fromElements(!(true)))).cardinality()), (_711_v5).Intersect(function () {
          let _coll39 = new _dafny.Set();
          for (const _compr_39 of _dafny.IntegerRange(new BigNumber(667), new BigNumber(946))) {
            let _731_v19 = _compr_39;
            if (((new BigNumber(667)).isLessThanOrEqualTo(_731_v19)) && ((_731_v19).isLessThan(new BigNumber(946)))) {
              _coll39.add(_module.__default.safeModuloInt(_731_v19, new BigNumber(199)));
            }
          }
          return _coll39;
        }()));
        r1 = _module.__default.safeModuloInt(new BigNumber((_709_v3).length), p2);
        r1 = (_dafny.ZERO).minus(p0);
      }
      r1 = p0;
      let _732_i1;
      _732_i1 = _dafny.ZERO;
      L5: {
        let _pat_let_tv3 = _707_v1;
        let _pat_let_tv4 = _703_v0;
        let _pat_let_tv5 = _703_v0;
        let _pat_let_tv6 = _707_v1;
        while (function (_source9) {
          if (_source9.is_DC1) {
            return (_module.D0.create_DC0(_pat_let_tv3)).dtor_cf0;
          } else if (_source9.is_DC2) {
            let _746___mcc_h3 = (_source9).cf1;
            let _747___mcc_h4 = (_source9).cf2;
            let _748_cf2 = _747___mcc_h4;
            let _749_cf1 = _746___mcc_h3;
            return (_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_pat_let_tv4).length))];
          } else {
            let _750___mcc_h5 = (_source9).cf0;
            let _751_cf0 = _750___mcc_h5;
            return _pat_let_tv6;
          }
        }(_712_v6)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_732_i1)) {
              break L5;
            }
            _732_i1 = (_732_i1).plus(_dafny.ONE);
            (globalState).f4 = _707_v1;
            let _733_v20;
            _733_v20 = _module.D0.create_DC0((_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))]);
            let _734_v21;
            _734_v21 = _dafny.Map.Empty.slice().updateUnsafe((_733_v20).dtor_cf0,_703_v0);
            _734_v21 = _734_v21;
            if ((_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))]) {
              let _735_v22;
              _735_v22 = _dafny.Seq.of(p2, new BigNumber(601));
              r1 = ((_735_v22)[_module.__default.safeIndex(p0, new BigNumber((_735_v22).length))]).plus(p0);
              r1 = p0;
              let _736_v23;
              _736_v23 = _dafny.Seq.of(_707_v1);
              let _737_v24;
              _737_v24 = _dafny.Seq.of(_736_v23);
              let _738_v25;
              _738_v25 = _module.D2.create_DC7(_module.D2.create_DC4(_737_v24));
              _738_v25 = _738_v25;
              let _739_v26;
              let _740_v27;
              let _741_v28;
              let _742_v29;
              let _out79;
              let _out80;
              let _out81;
              let _out82;
              let _outcollector20 = _module.__default.m0(_707_v1, _712_v6, (_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))], globalState);
              _out79 = _outcollector20[0];
              _out80 = _outcollector20[1];
              _out81 = _outcollector20[2];
              _out82 = _outcollector20[3];
              _739_v26 = _out79;
              _740_v27 = _out80;
              _741_v28 = _out81;
              _742_v29 = _out82;
              (globalState).f4 = (new BigNumber(608)).isLessThanOrEqualTo(new BigNumber(929));
            } else {
              let _index91 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length));
              (_703_v0)[_index91] = _707_v1;
              let _743_v30;
              _743_v30 = _dafny.Seq.of(_703_v0);
              r1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_743_v30).length), (p0).plus(p0)));
              let _744_v31;
              _744_v31 = _dafny.MultiSet.fromElements(_707_v1);
              let _745_v32;
              _745_v32 = _dafny.Seq.of((_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))]);
              r1 = (_this).fm14((new BigNumber(((_744_v31).update((_745_v32)[_module.__default.safeIndex(p0, new BigNumber((_745_v32).length))], _module.__default.abs(p2))).cardinality())).minus(p2), globalState);
              _711_v5 = (_711_v5).Difference(_dafny.Set.fromElements(p0));
              (globalState).f4 = (_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))];
            }
            (globalState).f4 = _707_v1;
          }
        }
      }
      _707_v1 = _707_v1;
      let _752_v33;
      let _nw123 = Array((new BigNumber(22)).toNumber());
      _nw123[(_dafny.ZERO).toNumber()] = _703_v0;
      _nw123[(_dafny.ONE).toNumber()] = _703_v0;
      _nw123[(new BigNumber(2)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(3)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(4)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(5)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(6)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(7)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(8)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(9)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(10)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(11)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(12)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(13)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(14)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(15)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(16)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(17)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(18)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(19)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(20)).toNumber()] = _703_v0;
      _nw123[(new BigNumber(21)).toNumber()] = _703_v0;
      _752_v33 = _nw123;
      let _753_v34;
      _753_v34 = _dafny.Seq.of((((_703_v0)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_703_v0).length))]) ? (_752_v33) : (_752_v33)));
      r0 = (_753_v34)[_module.__default.safeIndex(p0, new BigNumber((_753_v34).length))];
      r1 = (_module.__default.safeModuloInt((_dafny.ZERO).minus(p2), p0)).minus(p2);
      return [r0, r1];
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f12 = [];
      this._f13 = _dafny.Map.Empty;
      this._f14 = _dafny.Seq.of();
      this.f18 = false;
      this._f17 = false;
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
    __ctor(f17, f18, f13, f14, f12) {
      let _this = this;
      (_this)._f17 = f17;
      (_this).f18 = f18;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f12 = f12;
      return;
    }
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-74)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("oqldcxev")).length)));
    };
    fm5(p0, globalState) {
      let _this = this;
      return new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, (_this).f17),new BigNumber(-572))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_this.f18),new BigNumber(-820)))).length);
    };
    fm17(p0, p1, globalState) {
      let _this = this;
      return (_this).f17;
    };
    m3(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let r3 = false;
      let _754_v0;
      let _nw124 = new _module.C0();
      _nw124.__ctor();
      _754_v0 = _nw124;
      let _755_v1;
      let _nw125 = new _module.C0();
      _nw125.__ctor();
      _755_v1 = _nw125;
      let _756_v2;
      let _nw126 = Array((new BigNumber(11)).toNumber());
      _nw126[(_dafny.ZERO).toNumber()] = _755_v1;
      _nw126[(_dafny.ONE).toNumber()] = _754_v0;
      _nw126[(new BigNumber(2)).toNumber()] = _755_v1;
      _nw126[(new BigNumber(3)).toNumber()] = ((_this.f18) ? (_754_v0) : (_755_v1));
      _nw126[(new BigNumber(4)).toNumber()] = _754_v0;
      _nw126[(new BigNumber(5)).toNumber()] = _755_v1;
      _nw126[(new BigNumber(6)).toNumber()] = _754_v0;
      _nw126[(new BigNumber(7)).toNumber()] = _754_v0;
      _nw126[(new BigNumber(8)).toNumber()] = _754_v0;
      _nw126[(new BigNumber(9)).toNumber()] = _755_v1;
      _nw126[(new BigNumber(10)).toNumber()] = _755_v1;
      _756_v2 = _nw126;
      let _index92 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_756_v2).length));
      (_756_v2)[_index92] = _754_v0;
      let _index93 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_756_v2).length));
      (_756_v2)[_index93] = _754_v0;
      let _757_v3;
      _757_v3 = new BigNumber(18);
      let _758_v4;
      _758_v4 = _dafny.Seq.UnicodeFromString("rgpt");
      r1 = (((_module.__default.fm2(_this.f18, globalState)) || (!((_this).f17))) ? ((((_this).f17) ? ((_dafny.ZERO).minus(_757_v3)) : (_757_v3))) : (new BigNumber((_758_v4).length)));
      let _759_v5;
      _759_v5 = _dafny.Seq.of((_this).f14, (_this).f14);
      let _760_v6;
      _760_v6 = _module.D2.create_DC4(_759_v5);
      let _761_i0;
      _761_i0 = _dafny.ZERO;
      L6: {
        let _pat_let_tv7 = _757_v3;
        let _pat_let_tv8 = _757_v3;
        while (function (_source10) {
          if (_source10.is_DC5) {
            let _768___mcc_h0 = (_source10).cf5;
            let _769___mcc_h1 = (_source10).cf6;
            let _770___mcc_h2 = (_source10).cf7;
            let _771_cf7 = _770___mcc_h2;
            let _772_cf6 = _769___mcc_h1;
            let _773_cf5 = _768___mcc_h0;
            return (_dafny.Set.fromElements(_771_cf7)).IsSubsetOf(_dafny.Set.fromElements(_771_cf7, _771_cf7));
          } else if (_source10.is_DC6) {
            let _774___mcc_h3 = (_source10).cf8;
            let _775___mcc_h4 = (_source10).cf9;
            let _776_cf9 = _775___mcc_h4;
            let _777_cf8 = _774___mcc_h3;
            return (_this).f17;
          } else if (_source10.is_DC4) {
            let _778___mcc_h5 = (_source10).cf4;
            let _779_cf4 = _778___mcc_h5;
            return _this.f18;
          } else {
            let _780___mcc_h6 = (_source10).cf10;
            let _781_cf10 = _780___mcc_h6;
            return (_pat_let_tv7).isEqualTo(_pat_let_tv8);
          }
        }(_760_v6)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_761_i0)) {
              break L6;
            }
            _761_i0 = (_761_i0).plus(_dafny.ONE);
            let _762_v7;
            let _nw127 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _762_v7 = _nw127;
            let _763_v8;
            _763_v8 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_757_v3),_762_v7);
            _763_v8 = (_763_v8).update(_757_v3, _762_v7);
            let _764_v9;
            _764_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f18,_757_v3);
            _764_v9 = (_764_v9).update(_module.__default.fm2(true, globalState), _757_v3);
            let _765_v10;
            _765_v10 = _dafny.MultiSet.fromElements(true, _this.f18);
            let _766_v11;
            _766_v11 = _dafny.Set.fromElements((((_this).f17) ? ((_754_v0).fm14(_757_v3, globalState)) : (new BigNumber((_765_v10).cardinality()))), (_dafny.ZERO).minus((_757_v3).minus(_757_v3)), _757_v3);
            _766_v11 = (_766_v11).Intersect((_dafny.Set.fromElements(_757_v3)).Union(_766_v11));
            let _767_v12;
            _767_v12 = _dafny.Seq.of(_762_v7, _762_v7, _762_v7);
            _762_v7 = (_767_v12)[_module.__default.safeIndex(_757_v3, new BigNumber((_767_v12).length))];
          }
        }
      }
      let _782_v13;
      _782_v13 = _dafny.MultiSet.fromElements(true, (_this).f17);
      let _783_v14;
      _783_v14 = _dafny.Seq.of((((_782_v13).contains(_this.f18)) ? ((_782_v13).get(_this.f18)) : (_757_v3)), _757_v3);
      let _784_v15;
      _784_v15 = _module.D0.create_DC2(_757_v3, (_this).f17);
      let _785_v16;
      _785_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f17, globalState),(_784_v15).dtor_cf1);
      r3 = !(!((_this).fm17(_dafny.Seq.Concat(_dafny.Seq.update(_783_v14, _module.__default.safeIndex(_757_v3, new BigNumber((_783_v14).length)), _757_v3), _dafny.Seq.of(_757_v3)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_785_v16).contains(false)) ? ((_785_v16).get(false)) : (new BigNumber((_783_v14).length))),_this.f18)).length), globalState)));
      r0 = _module.__default.fm2((_this).f17, globalState);
      r1 = _757_v3;
      let _786_v17;
      _786_v17 = _dafny.Set.fromElements((_this).f14, _dafny.Seq.Concat((_this).f14, _module.__default.fm18(_this.f18, _757_v3, false, new BigNumber((_dafny.Seq.of(_757_v3)).length), globalState)), (_this).f14, (_this).f14, _dafny.Seq.Concat(_dafny.Seq.of((_this).fm17(_783_v14, new BigNumber((_758_v4).length), globalState)), (_this).f14));
      r2 = _786_v17;
      r3 = _this.f18;
      return [r0, r1, r2, r3];
    }
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f12 = [];
      this._f16 = _dafny.Seq.of();
      this._f15 = false;
      this._f13 = _dafny.Map.Empty;
      this._f14 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T3, _module.T0, _module.T4, _module.T2, _module.T1];
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
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    __ctor(f16, f15, f13, f14, f12) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f15 = f15;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f12 = f12;
      return;
    }
    fm9(globalState) {
      let _this = this;
      let _source11 = _module.D8.create_DC26((_this).f15);
      if (_source11.is_DC24) {
        let _787___mcc_h0 = (_source11).cf36;
        let _788_cf36 = _787___mcc_h0;
        return _dafny.MultiSet.FromArray((_this).f14);
      } else if (_source11.is_DC25) {
        let _789___mcc_h1 = (_source11).cf37;
        let _790___mcc_h2 = (_source11).cf38;
        let _791___mcc_h3 = (_source11).cf39;
        let _792___mcc_h4 = (_source11).cf40;
        let _793_cf40 = _792___mcc_h4;
        let _794_cf39 = _791___mcc_h3;
        let _795_cf38 = _790___mcc_h2;
        let _796_cf37 = _789___mcc_h1;
        return _dafny.MultiSet.fromElements(false, false);
      } else if (_source11.is_DC26) {
        let _797___mcc_h5 = (_source11).cf41;
        let _798_cf41 = _797___mcc_h5;
        if ((_this).f15) {
          return _dafny.MultiSet.fromElements(_798_cf41);
        } else {
          return _dafny.MultiSet.fromElements((_module.D0.create_DC0(_798_cf41)).dtor_cf0);
        }
      } else if (_source11.is_DC23) {
        let _799___mcc_h6 = (_source11).cf35;
        let _800_cf35 = _799___mcc_h6;
        return _dafny.MultiSet.fromElements((_this).f15, (_this).f15, false);
      } else {
        let _801___mcc_h7 = (_source11).cf42;
        let _802_cf42 = _801___mcc_h7;
        return (_dafny.MultiSet.FromArray((_this).f14));
      }
    };
    fm7(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((new BigNumber((function () {
        let _coll40 = new _dafny.Map();
        for (const _compr_40 of _dafny.IntegerRange(new BigNumber(209), new BigNumber(497))) {
          let _803_v0 = _compr_40;
          if (((new BigNumber(209)).isLessThanOrEqualTo(_803_v0)) && ((_803_v0).isLessThan(new BigNumber(497)))) {
            _coll40.push([_module.__default.safeDivisionInt(_803_v0, new BigNumber(494)),(_this).f15]);
          }
        }
        return _coll40;
      }()).length)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-566),(_this).f16)).length)), new BigNumber(239));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(844);
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.update((_this).f16, _module.__default.safeIndex((_dafny.ZERO).minus(((true) ? ((_dafny.ZERO).minus(new BigNumber(((_module.D6.create_DC17(_dafny.Set.fromElements(new BigNumber(553), new BigNumber(605)))).dtor_cf27).length))) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(517)), function (_804_i0) {
        return _dafny.Seq.UnicodeFromString("nan");
      })).length))))), new BigNumber(((_this).f16).length)), new BigNumber((function () {
        let _coll41 = new _dafny.Map();
        for (const _compr_41 of (function () {
          let _coll42 = new _dafny.Set();
          for (const _compr_42 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pjnkal"),(_this).f15)).Keys.Elements) {
            let _805_v1 = _compr_42;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pjnkal"),(_this).f15)).contains(_805_v1)) {
              _coll42.add(_805_v1);
            }
          }
          return _coll42;
        }()).Elements) {
          let _806_v0 = _compr_41;
          if ((function () {
            let _coll43 = new _dafny.Set();
            for (const _compr_43 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pjnkal"),(_this).f15)).Keys.Elements) {
              let _807_v1 = _compr_43;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pjnkal"),(_this).f15)).contains(_807_v1)) {
                _coll43.add(_807_v1);
              }
            }
            return _coll43;
          }()).contains(_806_v0)) {
            _coll41.push([_806_v0,new BigNumber((_dafny.Seq.UnicodeFromString("oy")).length)]);
          }
        }
        return _coll41;
      }()).length));
    };
    fm5(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Concat((_this).f14, (_this).f14)).length);
    };
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Set.fromElements(new BigNumber(-129), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(((_this).f14).length), new BigNumber(((_this).f14).length), new BigNumber(527), new BigNumber(772))).cardinality()));
    };
    fm38(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((((_module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15))).dtor_cf11).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f15))).length)).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(999), new BigNumber(912)))));
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _808_v0;
      _808_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
      let _809_v1;
      _809_v1 = _module.D3.create_DC8(_808_v0);
      let _810_v3;
      _810_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f15);
      let _811_v4;
      _811_v4 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p1,_809_v1), function () {
        let _coll44 = new _dafny.Map();
        for (const _compr_44 of (_810_v3).Keys.Elements) {
          let _812_v2 = _compr_44;
          if ((_810_v3).contains(_812_v2)) {
            _coll44.push([(_812_v2).plus(p1),_809_v1]);
          }
        }
        return _coll44;
      }());
      let _813_v5;
      _813_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(724),_module.D3.create_DC8(_808_v0));
      (globalState).f4 = !((_dafny.Seq.IsProperPrefixOf(_811_v4, _dafny.Seq.update(_811_v4, _module.__default.safeIndex(p0, new BigNumber((_811_v4).length)), _813_v5))) === ((_this).f15));
      let _hi4 = new BigNumber(414);
      for (let _814_i0 = new BigNumber(((_this).f14).length); _814_i0.isLessThan(_hi4); _814_i0 = _814_i0.plus(_dafny.ONE)) {
        let _815_v6;
        _815_v6 = new _dafny.CodePoint('m'.codePointAt(0));
        let _816_v7;
        _816_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_815_v6);
        let _817_v8;
        let _nw128 = Array((new BigNumber(13)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
        _nw128[(_dafny.ONE).toNumber()] = _815_v6;
        _nw128[(new BigNumber(2)).toNumber()] = _815_v6;
        _nw128[(new BigNumber(3)).toNumber()] = (p2)[_module.__default.safeIndex(_814_i0, new BigNumber((p2).length))];
        _nw128[(new BigNumber(4)).toNumber()] = _815_v6;
        _nw128[(new BigNumber(5)).toNumber()] = _815_v6;
        _nw128[(new BigNumber(6)).toNumber()] = _815_v6;
        _nw128[(new BigNumber(7)).toNumber()] = _815_v6;
        _nw128[(new BigNumber(8)).toNumber()] = _815_v6;
        _nw128[(new BigNumber(9)).toNumber()] = (((_816_v7).contains(false)) ? ((_816_v7).get(false)) : (_815_v6));
        _nw128[(new BigNumber(10)).toNumber()] = (((_this).f15) ? (_815_v6) : (_815_v6));
        _nw128[(new BigNumber(11)).toNumber()] = _815_v6;
        _nw128[(new BigNumber(12)).toNumber()] = _815_v6;
        _817_v8 = _nw128;
        let _index94 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_817_v8).length));
        (_817_v8)[_index94] = ((false) ? (_815_v6) : (_module.__default.fm39(p1, !((_this).f15), p0, p2, globalState)));
        let _index95 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_817_v8).length));
        (_817_v8)[_index95] = _815_v6;
        let _818_v9;
        _818_v9 = _dafny.MultiSet.fromElements((((_this).f15) ? (new BigNumber((p2).length)) : ((_dafny.ZERO).minus(p0))));
        _818_v9 = _dafny.MultiSet.FromArray(_dafny.Seq.of(p0));
        (globalState).f4 = !((_this).f15);
        r0 = (_dafny.ZERO).minus((p0).plus((p0).minus(p0)));
      }
      let _819_v10;
      let _init23 = function (_820_i1) {
        return (_this).f15;
      };
      let _nw129 = Array((new BigNumber(13)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw129.length); _i0_23++) {
        _nw129[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _819_v10 = _nw129;
      let _index96 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length));
      (_819_v10)[_index96] = (_this).f15;
      let _821_v11;
      _821_v11 = new _dafny.CodePoint('h'.codePointAt(0));
      let _index97 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length));
      (_819_v10)[_index97] = _dafny.Seq.contains(p2, _821_v11);
      (globalState).f4 = false;
      let _822_v12;
      _822_v12 = _module.D0.create_DC2(p1, (_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))]);
      let _823_i2;
      _823_i2 = _dafny.ZERO;
      L7: {
        let _pat_let_tv9 = p1;
        while (function (_source12) {
          if (_source12.is_DC1) {
            return (_this).f15;
          } else if (_source12.is_DC2) {
            let _837___mcc_h0 = (_source12).cf1;
            let _838___mcc_h1 = (_source12).cf2;
            let _839_cf2 = _838___mcc_h1;
            let _840_cf1 = _837___mcc_h0;
            return (_dafny.Set.fromElements((_this).f16)).IsDisjointFrom(_dafny.Set.fromElements((_this).f16));
          } else {
            let _841___mcc_h2 = (_source12).cf0;
            let _842_cf0 = _841___mcc_h2;
            return !(_dafny.areEqual(_dafny.Seq.of(_842_cf0), (_this).f14));
          }
        }(function (_pat_let4_0) {
          return function (_843_dt__update__tmp_h0) {
            return function (_pat_let5_0) {
              return function (_844_dt__update_hcf1_h0) {
                return _module.D0.create_DC2(_844_dt__update_hcf1_h0, (_843_dt__update__tmp_h0).dtor_cf2);
              }(_pat_let5_0);
            }(_pat_let_tv9);
          }(_pat_let4_0);
        }(_822_v12))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_823_i2)) {
              break L7;
            }
            _823_i2 = (_823_i2).plus(_dafny.ONE);
            r0 = new BigNumber((p2).length);
            let _824_v13;
            _824_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,p1);
            let _825_v14;
            _825_v14 = _dafny.Seq.of(_824_v13, _824_v13, _824_v13);
            let _826_v15;
            _826_v15 = _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,p0));
            let _827_v16;
            _827_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),p1);
            let _828_v17;
            let _nw130 = Array((new BigNumber(15)).toNumber());
            _nw130[(_dafny.ZERO).toNumber()] = (_825_v14)[_module.__default.safeIndex(p1, new BigNumber((_825_v14).length))];
            _nw130[(_dafny.ONE).toNumber()] = (_824_v13).Merge(_824_v13);
            _nw130[(new BigNumber(2)).toNumber()] = (_824_v13).update((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], p0);
            _nw130[(new BigNumber(3)).toNumber()] = ((false) ? (_dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber(827))) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber(235))));
            _nw130[(new BigNumber(4)).toNumber()] = (_826_v15).dtor_cf35;
            _nw130[(new BigNumber(5)).toNumber()] = _824_v13;
            _nw130[(new BigNumber(6)).toNumber()] = (_824_v13).Merge(_824_v13);
            _nw130[(new BigNumber(7)).toNumber()] = (_824_v13).update((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], new BigNumber(((_this).f16).length));
            _nw130[(new BigNumber(8)).toNumber()] = (_824_v13).update((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], p0);
            _nw130[(new BigNumber(9)).toNumber()] = (_824_v13).update((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], new BigNumber((p3).cardinality()));
            _nw130[(new BigNumber(10)).toNumber()] = _824_v13;
            _nw130[(new BigNumber(11)).toNumber()] = (_824_v13).update((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], (((_827_v16).contains(p1)) ? ((_827_v16).get(p1)) : (p1)));
            _nw130[(new BigNumber(12)).toNumber()] = (_824_v13).update((_this).f15, p0);
            _nw130[(new BigNumber(13)).toNumber()] = _824_v13;
            _nw130[(new BigNumber(14)).toNumber()] = (((_this).f15) ? (_dafny.Map.Empty.slice().updateUnsafe((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))],new BigNumber(97))) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f15,p1)));
            _828_v17 = _nw130;
            let _index98 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_828_v17).length));
            (_828_v17)[_index98] = _824_v13;
            let _index99 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_828_v17).length));
            (_828_v17)[_index99] = (_824_v13).Merge(_824_v13);
            let _829_v18;
            let _nw131 = new _module.C1();
            _nw131.__ctor(!((_this).f15), (_this).f15, (_this).f13, (_this).f14, _this.f12);
            _829_v18 = _nw131;
            let _830_v19;
            _830_v19 = _dafny.MultiSet.fromElements(_829_v18);
            let _831_v20;
            _831_v20 = _dafny.Seq.of(_829_v18, _829_v18);
            let _832_v21;
            _832_v21 = _dafny.Map.Empty.slice().updateUnsafe((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))],_829_v18);
            let _833_v22;
            _833_v22 = _dafny.Map.Empty.slice().updateUnsafe(_821_v11,(_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))]);
            let _834_v23;
            let _nw132 = Array((new BigNumber(19)).toNumber());
            _nw132[(_dafny.ZERO).toNumber()] = _830_v19;
            _nw132[(_dafny.ONE).toNumber()] = _830_v19;
            _nw132[(new BigNumber(2)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(3)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_831_v20);
            _nw132[(new BigNumber(5)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(6)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(7)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.fromElements(_829_v18, _829_v18, _829_v18, _829_v18, _829_v18)).Difference(_830_v19);
            _nw132[(new BigNumber(9)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(10)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_829_v18, (((_832_v21).contains((_this).f15)) ? ((_832_v21).get((_this).f15)) : (_829_v18)), _829_v18);
            _nw132[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_829_v18, _829_v18);
            _nw132[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_829_v18, _829_v18, _829_v18);
            _nw132[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(_829_v18, (((_832_v21).contains((((_833_v22).contains(_821_v11)) ? ((_833_v22).get(_821_v11)) : ((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))])))) ? ((_832_v21).get((((_833_v22).contains(_821_v11)) ? ((_833_v22).get(_821_v11)) : ((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))])))) : (_829_v18)));
            _nw132[(new BigNumber(15)).toNumber()] = (_830_v19).Difference(_830_v19);
            _nw132[(new BigNumber(16)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(17)).toNumber()] = _830_v19;
            _nw132[(new BigNumber(18)).toNumber()] = _830_v19;
            _834_v23 = _nw132;
            let _835_v24;
            _835_v24 = _dafny.Map.Empty.slice().updateUnsafe((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))],_834_v23);
            let _836_v25;
            _836_v25 = _dafny.Seq.of((((_835_v24).contains(_module.__default.fm2((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], globalState))) ? ((_835_v24).get(_module.__default.fm2((_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], globalState))) : (_834_v23)), _834_v23, _834_v23);
            let _rhs67 = p1;
            let _rhs68 = p0;
            let _rhs69 = (_836_v25)[_module.__default.safeIndex(p1, new BigNumber((_836_v25).length))];
            r0 = _rhs67;
            r0 = _rhs68;
            _834_v23 = _rhs69;
            let _index100 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length));
            (_819_v10)[_index100] = _module.__default.fm2(_dafny.areEqual((_this).f14, (_829_v18).f14), globalState);
          }
        }
      }
      r0 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((p2).length)), _821_v11)).length), p1);
      r0 = (_this).fm38(p1, (_dafny.ZERO).minus(((_this).f16)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("qifnlwcgl")).length), new BigNumber(((_this).f16).length))]), p0, (_819_v10)[_module.__default.safeIndex(new BigNumber(859), new BigNumber((_819_v10).length))], globalState);
      return r0;
    }
    m3(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let r3 = false;
      let _845_v0;
      _845_v0 = _module.D2.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), function (_846_i0) {
  return (_this).f14;
}));
      _845_v0 = _845_v0;
      let _847_v1;
      _847_v1 = new BigNumber(-924);
      let _848_v2;
      _848_v2 = _dafny.Set.fromElements(_847_v1);
      let _849_v3;
      _849_v3 = _module.D8.create_DC24((_this).f15);
      let _850_v4;
      _850_v4 = _module.D8.create_DC27(_849_v3);
      let _851_v5;
      _851_v5 = new _dafny.CodePoint('d'.codePointAt(0));
      let _852_v6;
      _852_v6 = _dafny.Map.Empty.slice().updateUnsafe(_850_v4,_851_v5);
      let _853_v7;
      _853_v7 = _dafny.Seq.of(_852_v6);
      let _854_v8;
      _854_v8 = _dafny.Seq.UnicodeFromString("qjlnmny");
      let _855_v9;
      let _nw133 = Array((new BigNumber(27)).toNumber());
      _nw133[(_dafny.ZERO).toNumber()] = (_this).f15;
      _nw133[(_dafny.ONE).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(2)).toNumber()] = (_848_v2).IsSubsetOf(_848_v2);
      _nw133[(new BigNumber(3)).toNumber()] = !((_this).f15);
      _nw133[(new BigNumber(4)).toNumber()] = ((_this).f15) && (((_this).f14)[_module.__default.safeIndex(_847_v1, new BigNumber(((_this).f14).length))]);
      _nw133[(new BigNumber(5)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(6)).toNumber()] = (_847_v1).isLessThan(new BigNumber((_853_v7).length));
      _nw133[(new BigNumber(7)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(8)).toNumber()] = _module.__default.fm2((_this).f15, globalState);
      _nw133[(new BigNumber(9)).toNumber()] = !((_this).f15) || ((_this).f15);
      _nw133[(new BigNumber(10)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(11)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(12)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(13)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(14)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(15)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(16)).toNumber()] = !_dafny.areEqual(_854_v8, _854_v8);
      _nw133[(new BigNumber(17)).toNumber()] = false;
      _nw133[(new BigNumber(18)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(19)).toNumber()] = false;
      _nw133[(new BigNumber(20)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(21)).toNumber()] = false;
      _nw133[(new BigNumber(22)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(23)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(24)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(722))).length)).isLessThan(_847_v1);
      _nw133[(new BigNumber(25)).toNumber()] = (_this).f15;
      _nw133[(new BigNumber(26)).toNumber()] = false;
      _855_v9 = _nw133;
      let _856_v10;
      _856_v10 = _module.D7.create_DC22((_this).f15);
      let _index101 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length));
      (_855_v9)[_index101] = (_856_v10).dtor_cf34;
      let _857_v11;
      _857_v11 = _dafny.MultiSet.fromElements(_847_v1, _847_v1);
      let _index102 = _module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length));
      (_855_v9)[_index102] = ((_dafny.MultiSet.fromElements(_847_v1, _847_v1, _847_v1)).Difference((_857_v11).update(_847_v1, _module.__default.abs(new BigNumber((_854_v8).length))))).IsProperSubsetOf(_857_v11);
      (_this).m12(!((_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))]), globalState);
      let _858_v12;
      _858_v12 = _dafny.Map.Empty.slice().updateUnsafe((((_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))]) ? (_847_v1) : (new BigNumber(271))),_847_v1);
      _858_v12 = (_858_v12).update(_847_v1, _847_v1);
      let _859_v13;
      _859_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))], globalState),_dafny.Seq.update(_854_v8, _module.__default.safeIndex(new BigNumber(((_this).f16).length), new BigNumber((_854_v8).length)), _851_v5));
      let _860_v14;
      _860_v14 = _dafny.Set.fromElements((_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))], (_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))], (_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))]);
      let _861_v15;
      _861_v15 = _dafny.Seq.of(new BigNumber((_859_v13).length), new BigNumber(((_860_v14).Intersect(_dafny.Set.fromElements(false, (_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))]))).length), (((_857_v11).contains(new BigNumber((_dafny.Seq.UnicodeFromString("jefu")).length))) ? ((_857_v11).get(new BigNumber((_dafny.Seq.UnicodeFromString("jefu")).length))) : (_847_v1)), new BigNumber(520), (new BigNumber(((_this).f14).length)).minus(_847_v1));
      let _rhs70 = _851_v5;
      let _rhs71 = _dafny.Seq.update((_this).f16, _module.__default.safeIndex(_847_v1, new BigNumber(((_this).f16).length)), _847_v1);
      _851_v5 = _rhs70;
      _861_v15 = _rhs71;
      let _862_v16;
      _862_v16 = _module.D4.create_DC14((_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))]);
      _854_v8 = ((!(false)) ? (_854_v8) : (_module.__default.fm40(_862_v16, _847_v1, true, (_this).f15, globalState)));
      r0 = (_module.__default.safeModuloInt(_847_v1, new BigNumber((_858_v12).length))).isLessThan(_847_v1);
      let _863_v17;
      _863_v17 = _dafny.MultiSet.fromElements((_this).f14);
      r1 = (new BigNumber((_863_v17).cardinality())).multipliedBy(_847_v1);
      r2 = _module.__default.fm41((new BigNumber((_859_v13).length)).minus((_dafny.ZERO).minus(_847_v1)), _851_v5, globalState);
      r3 = (_855_v9)[_module.__default.safeIndex(new BigNumber(4), new BigNumber((_855_v9).length))];
      return [r0, r1, r2, r3];
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _864_v0;
      _864_v0 = new _dafny.CodePoint('t'.codePointAt(0));
      let _865_v1;
      _865_v1 = _dafny.MultiSet.fromElements(new BigNumber(((_this).f14).length), new BigNumber(((_this).f14).length), new BigNumber(660));
      let _866_v2;
      _866_v2 = _dafny.Map.Empty.slice().updateUnsafe(_864_v0,_865_v1);
      let _867_v3;
      _867_v3 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.fm2((_this).f15, globalState));
      let _868_v4;
      _868_v4 = _module.D7.create_DC21(_867_v3, (_this).f15, (_this).f15);
      _866_v2 = _module.__default.fm42(_868_v4, p1, globalState);
      let _869_v5;
      _869_v5 = _dafny.MultiSet.fromElements((_this).f15);
      let _hi5 = new BigNumber((_dafny.Seq.Concat(p1, p1)).length);
      for (let _870_i0 = (((_869_v5).contains((_this).f15)) ? ((_869_v5).get((_this).f15)) : (p2)); _870_i0.isLessThan(_hi5); _870_i0 = _870_i0.plus(_dafny.ONE)) {
        let _871_v6;
        let _nw134 = Array((new BigNumber(21)).toNumber()).fill(false);
        _871_v6 = _nw134;
        let _pat_let_tv10 = _871_v6;
        let _source13 = function (_pat_let6_0) {
          return function (_872_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_873_dt__update_hcf8_h0) {
                return _module.D2.create_DC6(_873_dt__update_hcf8_h0, (_872_dt__update__tmp_h0).dtor_cf9);
              }(_pat_let7_0);
            }(_pat_let_tv10);
          }(_pat_let6_0);
        }(_module.D2.create_DC6(_871_v6, _871_v6));
        if (_source13.is_DC5) {
          let _874___mcc_h0 = (_source13).cf5;
          let _875___mcc_h1 = (_source13).cf6;
          let _876___mcc_h2 = (_source13).cf7;
          let _877_cf7 = _876___mcc_h2;
          let _878_cf6 = _875___mcc_h1;
          let _879_cf5 = _874___mcc_h0;
          let _880_v8;
          let _nw135 = Array((new BigNumber(19)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = p2;
          _nw135[(_dafny.ONE).toNumber()] = _879_cf5;
          _nw135[(new BigNumber(2)).toNumber()] = (((_this).f15) ? (_879_cf5) : ((_this).fm5(_878_cf6, globalState)));
          _nw135[(new BigNumber(3)).toNumber()] = new BigNumber(-356);
          _nw135[(new BigNumber(4)).toNumber()] = _870_i0;
          _nw135[(new BigNumber(5)).toNumber()] = _879_cf5;
          _nw135[(new BigNumber(6)).toNumber()] = ((!(_877_cf7)) ? (_879_cf5) : (_879_cf5));
          _nw135[(new BigNumber(7)).toNumber()] = p2;
          _nw135[(new BigNumber(8)).toNumber()] = (p2).plus(_879_cf5);
          _nw135[(new BigNumber(9)).toNumber()] = p2;
          _nw135[(new BigNumber(10)).toNumber()] = new BigNumber((p1).length);
          _nw135[(new BigNumber(11)).toNumber()] = _870_i0;
          _nw135[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), ((_881_v0) => function (_882_i1) {
            return _881_v0;
          })(_864_v0))).length);
          _nw135[(new BigNumber(13)).toNumber()] = (new BigNumber(209)).minus(_870_i0);
          _nw135[(new BigNumber(14)).toNumber()] = p2;
          _nw135[(new BigNumber(15)).toNumber()] = _879_cf5;
          _nw135[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((function () {
            let _coll45 = new _dafny.Map();
            for (const _compr_45 of _dafny.IntegerRange(new BigNumber(624), new BigNumber(717))) {
              let _883_v7 = _compr_45;
              if (((new BigNumber(624)).isLessThanOrEqualTo(_883_v7)) && ((_883_v7).isLessThan(new BigNumber(717)))) {
                _coll45.push([(_883_v7).minus(p2),_864_v0]);
              }
            }
            return _coll45;
          }()).length), p2);
          _nw135[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(_879_cf5, new BigNumber((p1).length));
          _nw135[(new BigNumber(18)).toNumber()] = _870_i0;
          _880_v8 = _nw135;
          _880_v8 = _880_v8;
          _879_cf5 = _module.__default.safeDivisionInt(p2, (_879_cf5).minus(new BigNumber((_dafny.Seq.UnicodeFromString("mmitbisb")).length)));
          (globalState).f4 = _877_cf7;
          _879_cf5 = (new BigNumber(-39)).plus(p2);
        } else if (_source13.is_DC6) {
          let _884___mcc_h3 = (_source13).cf8;
          let _885___mcc_h4 = (_source13).cf9;
          let _886_cf9 = _885___mcc_h4;
          let _887_cf8 = _884___mcc_h3;
          let _888_v9;
          _888_v9 = _dafny.Seq.of((new BigNumber(60)).multipliedBy(_870_i0));
          _888_v9 = _dafny.Seq.Concat(_dafny.Seq.of(p2, p2), _888_v9);
          let _889_v10;
          _889_v10 = _869_v5;
          let _890_v11;
          _890_v11 = _module.D7.create_DC22((_this).f15);
          let _891_v12;
          _891_v12 = _dafny.Seq.of(_module.__default.fm43(_889_v10, _890_v11, globalState));
          let _index103 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_871_v6).length));
          (_871_v6)[_index103] = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_891_v12, _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_891_v12).length)), (_this).f14), _891_v12);
          let _index104 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_871_v6).length));
          (_871_v6)[_index104] = _module.__default.fm2((_this).f15, globalState);
          let _892_v13;
          _892_v13 = _module.D5.create_DC16((_871_v6)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_871_v6).length))]);
          let _893_v14;
          _893_v14 = _dafny.Map.Empty.slice().updateUnsafe(_892_v13,_this.f12);
          let _894_v15;
          let _nw136 = new _module.C1();
          _nw136.__ctor((_871_v6)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_871_v6).length))], (_871_v6)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_871_v6).length))], (_this).f13, _module.__default.fm43(_889_v10, _890_v11, globalState), (((_893_v14).contains(_892_v13)) ? ((_893_v14).get(_892_v13)) : (_this.f12)));
          _894_v15 = _nw136;
          let _895_v16;
          let _init24 = ((_896_i0) => function (_897_i2) {
            return _module.__default.safeDivisionInt(_897_i2, _896_i0);
          })(_870_i0);
          let _nw137 = Array((new BigNumber(20)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw137.length); _i0_24++) {
            _nw137[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _895_v16 = _nw137;
          let _898_v17;
          _898_v17 = _dafny.Seq.of(_895_v16, _895_v16);
          let _899_v18;
          _899_v18 = _dafny.Map.Empty.slice().updateUnsafe(false,_898_v17);
          let _900_v19;
          _900_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,_895_v16);
          let _index105 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_871_v6).length));
          (_871_v6)[_index105] = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.update((((_899_v18).contains(_894_v15.f18)) ? ((_899_v18).get(_894_v15.f18)) : (_dafny.Seq.of(_895_v16, (((_900_v19).contains(p2)) ? ((_900_v19).get(p2)) : (_895_v16))))), _module.__default.safeIndex((((_865_v1).contains(p2)) ? ((_865_v1).get(p2)) : (new BigNumber(-225))), new BigNumber(((((_899_v18).contains(_894_v15.f18)) ? ((_899_v18).get(_894_v15.f18)) : (_dafny.Seq.of(_895_v16, (((_900_v19).contains(p2)) ? ((_900_v19).get(p2)) : (_895_v16)))))).length)), _895_v16), _898_v17), _898_v17);
        } else if (_source13.is_DC4) {
          let _901___mcc_h5 = (_source13).cf4;
          let _902_cf4 = _901___mcc_h5;
          (globalState).f4 = !(_867_v3).equals(_867_v3);
          let _903_v20;
          _903_v20 = _dafny.Seq.UnicodeFromString("pdnbe");
          _903_v20 = _dafny.Seq.Concat(_dafny.Seq.Concat(_903_v20, _dafny.Seq.Create(_module.__default.abs(new BigNumber(233)), ((_904_v0) => function (_905_i3) {
            return _904_v0;
          })(_864_v0))), _dafny.Seq.Concat(_module.__default.fm40(_module.D4.create_DC14((_this).f15), _870_i0, (_this).f15, (_this).f15, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-820)), ((_906_v0) => function (_907_i4) {
            return _906_v0;
          })(_864_v0))));
          let _908_v21;
          _908_v21 = _864_v0;
          _908_v21 = _908_v21;
          let _909_v22;
          _909_v22 = new BigNumber(123);
          _909_v22 = _909_v22;
        } else {
          let _910___mcc_h6 = (_source13).cf10;
          let _911_cf10 = _910___mcc_h6;
          let _912_v23;
          let _nw138 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _912_v23 = _nw138;
          let _index106 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_912_v23).length));
          (_912_v23)[_index106] = (_dafny.ZERO).minus(new BigNumber(((_this).f16).length));
          let _index107 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_912_v23).length));
          (_912_v23)[_index107] = (_dafny.ZERO).minus((_870_i0).minus(p2));
          let _913_v24;
          _913_v24 = _dafny.Seq.UnicodeFromString("ak");
          _913_v24 = _913_v24;
          _871_v6 = (((_this).f15) ? (_871_v6) : (_871_v6));
          let _index108 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_912_v23).length));
          (_912_v23)[_index108] = (_dafny.ZERO).minus(((_912_v23)[_module.__default.safeIndex(new BigNumber(973), new BigNumber((_912_v23).length))]).multipliedBy(p2));
        }
        r0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((_this).f16, (_this).f16), (_this).f16);
        let _914_v25;
        let _nw139 = Array((new BigNumber(13)).toNumber()).fill([]);
        _914_v25 = _nw139;
        let _915_v26;
        let _init25 = ((_916_p2) => function (_917_i5) {
          return (_917_i5).plus(_916_p2);
        })(p2);
        let _nw140 = Array((new BigNumber(3)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw140.length); _i0_25++) {
          _nw140[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _915_v26 = _nw140;
        let _index109 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_914_v25).length));
        (_914_v25)[_index109] = _915_v26;
        let _index110 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_914_v25).length));
        (_914_v25)[_index110] = _915_v26;
        (globalState).f4 = (new BigNumber((p0).length)).isLessThan(p2);
      }
      let _918_v27;
      _918_v27 = _module.D4.create_DC12((_this).f14);
      let _source14 = _module.__default.fm44(_918_v27, (_this).f15, globalState);
      if (_source14.is_DC5) {
        let _919___mcc_h7 = (_source14).cf5;
        let _920___mcc_h8 = (_source14).cf6;
        let _921___mcc_h9 = (_source14).cf7;
        let _922_cf7 = _921___mcc_h9;
        let _923_cf6 = _920___mcc_h8;
        let _924_cf5 = _919___mcc_h7;
        _924_cf5 = (_924_cf5).minus(p2);
        let _925_v28;
        let _init26 = ((_926_p2) => function (_927_i6) {
          return (_927_i6).plus(_926_p2);
        })(p2);
        let _nw141 = Array((new BigNumber(29)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw141.length); _i0_26++) {
          _nw141[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _925_v28 = _nw141;
        let _index111 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_925_v28).length));
        (_925_v28)[_index111] = (_dafny.ZERO).minus(_924_cf5);
        let _928_v29;
        _928_v29 = _dafny.Map.Empty.slice().updateUnsafe(_864_v0,p2);
        let _index112 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_925_v28).length));
        let _rhs72 = _923_cf6;
        let _rhs73 = new BigNumber(927);
        let _rhs74 = _924_cf5;
        let _rhs75 = ((((_928_v29).contains(_864_v0)) ? ((_928_v29).get(_864_v0)) : (p2))).plus(p2);
        let _lhs41 = _925_v28;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_925_v28).length));
        r0 = _rhs72;
        _lhs41[_lhs42] = _rhs73;
        _924_cf5 = _rhs74;
        _924_cf5 = _rhs75;
        let _929_v30;
        _929_v30 = _dafny.Map.Empty.slice().updateUnsafe((_925_v28)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_925_v28).length))],p2);
        let _930_v31;
        _930_v31 = _dafny.Map.Empty.slice().updateUnsafe(_929_v30,_864_v0);
        let _931_v32;
        _931_v32 = _dafny.Map.Empty.slice().updateUnsafe(_867_v3,new BigNumber((_930_v31).length));
        let _932_v33;
        _932_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber((_931_v32).length));
        _932_v33 = (_932_v33).update(!((_this).f15) || ((_this).f15), (((((_867_v3).contains((_dafny.ZERO).minus((_925_v28)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_925_v28).length))]))) ? ((_867_v3).get((_dafny.ZERO).minus((_925_v28)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_925_v28).length))]))) : (_923_cf6))) ? ((_dafny.ZERO).minus(p2)) : (p2)));
        let _index113 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_925_v28).length));
        (_925_v28)[_index113] = _924_cf5;
      } else if (_source14.is_DC6) {
        let _933___mcc_h10 = (_source14).cf8;
        let _934___mcc_h11 = (_source14).cf9;
        let _935_cf9 = _934___mcc_h11;
        let _936_cf8 = _933___mcc_h10;
        let _937_v34;
        _937_v34 = _dafny.Set.fromElements(p2, p2);
        r0 = !((_937_v34).IsDisjointFrom((_937_v34).Union(_dafny.Set.fromElements(p2))));
        _864_v0 = _864_v0;
        let _938_v35;
        _938_v35 = _dafny.Seq.UnicodeFromString("jvd");
        _938_v35 = _938_v35;
        let _rhs76 = (p2).isLessThan(new BigNumber(859));
        let _rhs77 = (_this).f15;
        let _lhs43 = globalState;
        let _lhs44 = globalState;
        _lhs43.f4 = _rhs76;
        _lhs44.f4 = _rhs77;
      } else if (_source14.is_DC4) {
        let _939___mcc_h12 = (_source14).cf4;
        let _940_cf4 = _939___mcc_h12;
        let _941_v36;
        _941_v36 = _module.D0.create_DC2(new BigNumber(((_this).f16).length), (_this).f15);
        let _942_v37;
        let _943_v38;
        let _944_v39;
        let _945_v40;
        let _out83;
        let _out84;
        let _out85;
        let _out86;
        let _outcollector21 = _module.__default.m0((((_this).f15) ? ((_this).f15) : ((_this).f15)), function (_pat_let8_0) {
          return function (_946_dt__update__tmp_h1) {
            return function (_pat_let9_0) {
              return function (_947_dt__update_hcf2_h0) {
                return _module.D0.create_DC2((_946_dt__update__tmp_h1).dtor_cf1, _947_dt__update_hcf2_h0);
              }(_pat_let9_0);
            }((_this).f15);
          }(_pat_let8_0);
        }(_941_v36), (_this).f15, globalState);
        _out83 = _outcollector21[0];
        _out84 = _outcollector21[1];
        _out85 = _outcollector21[2];
        _out86 = _outcollector21[3];
        _942_v37 = _out83;
        _943_v38 = _out84;
        _944_v39 = _out85;
        _945_v40 = _out86;
        _918_v27 = _918_v27;
        let _948_v41;
        let _nw142 = new _module.C0();
        _nw142.__ctor();
        _948_v41 = _nw142;
        let _949_v42;
        let _init27 = ((_950_v39) => function (_951_i7) {
          return (_951_i7).plus(_950_v39);
        })(_944_v39);
        let _nw143 = Array((new BigNumber(8)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw143.length); _i0_27++) {
          _nw143[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _949_v42 = _nw143;
        let _index114 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_949_v42).length));
        (_949_v42)[_index114] = p2;
        let _index115 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_949_v42).length));
        (_949_v42)[_index115] = (((_869_v5).contains(_943_v38)) ? ((_869_v5).get(_943_v38)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), ((_952_v0) => function (_953_i8) {
          return _952_v0;
        })(_864_v0))).length)));
      } else {
        let _954___mcc_h13 = (_source14).cf10;
        let _955_cf10 = _954___mcc_h13;
        let _956_v43;
        _956_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_dafny.ZERO).minus(new BigNumber((p1).length)));
        let _957_v44;
        _957_v44 = _dafny.Seq.of(_956_v43, _956_v43, _module.__default.fm45((_this).f15, ((_this).f16)[_module.__default.safeIndex(p2, new BigNumber(((_this).f16).length))], p2, globalState), (_dafny.Map.Empty.slice().updateUnsafe(true,p2)).Merge(_956_v43));
        _957_v44 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_956_v43, _956_v43), _dafny.Seq.of(_956_v43)), _module.__default.safeIndex(new BigNumber(((_956_v43).Merge(_956_v43)).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_956_v43, _956_v43), _dafny.Seq.of(_956_v43))).length)), _956_v43);
        let _958_v45;
        _958_v45 = new BigNumber(138);
        _958_v45 = (_dafny.ZERO).minus(_958_v45);
        let _959_v46;
        _959_v46 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(p2, new BigNumber((_dafny.Seq.UnicodeFromString("ahgwcjry")).length)));
        _959_v46 = (_959_v46).Intersect(_dafny.Set.fromElements(_958_v45, _958_v45));
        _864_v0 = new _dafny.CodePoint('n'.codePointAt(0));
      }
      if (!(p2).isEqualTo(p2)) {
        let _960_v47;
        _960_v47 = new BigNumber(92);
        _960_v47 = p2;
        if ((_this).f15) {
          let _961_v48;
          let _nw144 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _961_v48 = _nw144;
          let _index116 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_961_v48).length));
          (_961_v48)[_index116] = p2;
          let _index117 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_961_v48).length));
          (_961_v48)[_index117] = (p2).minus(_module.__default.safeModuloInt(p2, new BigNumber((p1).length)));
          let _962_v49;
          let _nw145 = Array((new BigNumber(3)).toNumber()).fill(false);
          _962_v49 = _nw145;
          let _index118 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v49).length));
          (_962_v49)[_index118] = (_this).f15;
          let _963_v50;
          _963_v50 = _dafny.Set.fromElements(_961_v48);
          let _964_v51;
          _964_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_963_v50).length),p2);
          let _965_v52;
          _965_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f15, globalState),new BigNumber((_964_v51).length));
          let _966_v53;
          _966_v53 = _dafny.Seq.of(_965_v52, (_965_v52).update((_module.D0.create_DC2(new BigNumber((p1).length), (_this).f15)).dtor_cf2, p2));
          let _967_v54;
          _967_v54 = _dafny.Map.Empty.slice().updateUnsafe(_961_v48,_dafny.MultiSet.FromArray(_966_v53));
          let _968_v55;
          _968_v55 = _dafny.MultiSet.fromElements(_965_v52, _dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber(401)), _965_v52);
          let _index119 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v49).length));
          (_962_v49)[_index119] = !(((((_967_v54).contains(_961_v48)) ? ((_967_v54).get(_961_v48)) : (_968_v55))).IsDisjointFrom(_968_v55));
          let _index120 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_961_v48).length));
          (_961_v48)[_index120] = _960_v47;
          let _969_v56;
          _969_v56 = _module.D0.create_DC0((_this).f15);
          _969_v56 = _969_v56;
          let _970_v57;
          _970_v57 = _dafny.Map.Empty.slice().updateUnsafe((_962_v49)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v49).length))],true);
          let _971_v58;
          let _nw146 = new _module.C1();
          _nw146.__ctor(false, !((_962_v49)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v49).length))]), _dafny.Map.Empty.slice().updateUnsafe(false,_970_v57), _dafny.Seq.of((_962_v49)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_962_v49).length))]), _this.f12);
          _971_v58 = _nw146;
          let _972_v59;
          let _nw147 = new _module.C1();
          _nw147.__ctor((_dafny.Map.Empty.slice().updateUnsafe(_971_v58,(_961_v48)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_961_v48).length))])).equals(_dafny.Map.Empty.slice().updateUnsafe(_971_v58,_960_v47)), _971_v58.f18, (_this).f13, (_this).f14, _this.f12);
          _972_v59 = _nw147;
          _972_v59 = _971_v58;
        } else {
          let _rhs78 = _864_v0;
          let _rhs79 = (_this).f15;
          let _lhs45 = globalState;
          _864_v0 = _rhs78;
          _lhs45.f4 = _rhs79;
          let _973_v60;
          let _nw148 = Array((new BigNumber(26)).toNumber()).fill(false);
          _973_v60 = _nw148;
          let _index121 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length));
          (_973_v60)[_index121] = false;
          let _974_v61;
          let _nw149 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _974_v61 = _nw149;
          let _index122 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length));
          (_974_v61)[_index122] = new BigNumber(-21);
          let _975_v62;
          _975_v62 = _dafny.Seq.of(_974_v61, _974_v61);
          let _976_v63;
          _976_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), ((_977_p1) => function (_978_i9) {
            return (_977_p1)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_this).f15)).length), new BigNumber((_977_p1).length))];
          })(p1))).length))));
          let _index123 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length));
          let _index124 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length));
          let _rhs80 = (_this).f15;
          let _rhs81 = new BigNumber(((_this).f14).length);
          let _rhs82 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_975_v62, _dafny.Seq.of(_974_v61, _974_v61, _974_v61))).length)), p2);
          let _rhs83 = ((!(true)) ? (p2) : ((new BigNumber((p1).length)).plus(_960_v47)));
          let _rhs84 = new BigNumber(((_976_v63).Merge(_976_v63)).length);
          let _lhs46 = _973_v60;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length));
          let _lhs48 = _974_v61;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length));
          _lhs46[_lhs47] = _rhs80;
          _lhs48[_lhs49] = _rhs81;
          _960_v47 = _rhs82;
          _960_v47 = _rhs83;
          _960_v47 = _rhs84;
          let _979_v64;
          let _nw150 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
          _979_v64 = _nw150;
          let _980_v65;
          _980_v65 = _module.D3.create_DC8(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))]));
          let _981_v66;
          _981_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,!((_this).f15));
          let _982_v67;
          _982_v67 = _dafny.Seq.of(_module.D3.create_DC8(_981_v66));
          let _index125 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_979_v64).length));
          (_979_v64)[_index125] = _dafny.Seq.Concat(_dafny.Seq.of(_980_v65, _980_v65), _982_v67);
          let _983_v68;
          _983_v68 = _dafny.Set.fromElements((_974_v61)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length))]);
          let _index126 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_979_v64).length));
          (_979_v64)[_index126] = _dafny.Seq.update(_dafny.Seq.of(_980_v65, _980_v65), _module.__default.safeIndex(new BigNumber((_983_v68).length), new BigNumber((_dafny.Seq.of(_980_v65, _980_v65)).length)), _980_v65);
          let _984_v69;
          _984_v69 = _module.D6.create_DC18((_974_v61)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length))]);
          let _985_v70;
          _985_v70 = _module.D5.create_DC16((_this).f15);
          let _986_v71;
          _986_v71 = _dafny.Map.Empty.slice().updateUnsafe(_985_v70,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(((_this).f16).length))).cardinality()));
          let _987_v72;
          _987_v72 = _module.D2.create_DC5((((_986_v71).contains(_module.D5.create_DC16((_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))]))) ? ((_986_v71).get(_module.D5.create_DC16((_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))]))) : ((_974_v61)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length))])), (_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))], ((_this).f14)[_module.__default.safeIndex(_960_v47, new BigNumber(((_this).f14).length))]);
          let _988_v73;
          _988_v73 = _dafny.Map.Empty.slice().updateUnsafe(_984_v69,_987_v72);
          let _989_v75;
          _989_v75 = _dafny.MultiSet.fromElements(_module.D6.create_DC18(_960_v47), _984_v69, _984_v69);
          let _990_v76;
          let _nw151 = Array((new BigNumber(5)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = (_988_v73).update(_984_v69, _987_v72);
          _nw151[(_dafny.ONE).toNumber()] = function () {
            let _coll46 = new _dafny.Map();
            for (const _compr_46 of (((_989_v75).update(_module.D6.create_DC18((_974_v61)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length))]), _module.__default.abs(p2))).update(_984_v69, _module.__default.abs(p2))).Elements) {
              let _991_v74 = _compr_46;
              if ((((_989_v75).update(_module.D6.create_DC18((_974_v61)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length))]), _module.__default.abs(p2))).update(_984_v69, _module.__default.abs(p2))).contains(_991_v74)) {
                _coll46.push([_991_v74,_987_v72]);
              }
            }
            return _coll46;
          }();
          _nw151[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_984_v69,_987_v72);
          _nw151[(new BigNumber(3)).toNumber()] = _988_v73;
          _nw151[(new BigNumber(4)).toNumber()] = _988_v73;
          _990_v76 = _nw151;
          let _index127 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_990_v76).length));
          (_990_v76)[_index127] = _988_v73;
          let _992_v77;
          _992_v77 = _dafny.Map.Empty.slice().updateUnsafe(_960_v47,p2);
          let _index128 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_990_v76).length));
          let _index129 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length));
          let _rhs85 = _988_v73;
          let _rhs86 = (_module.__default.safeModuloInt(new BigNumber((_992_v77).length), _960_v47)).isLessThanOrEqualTo(new BigNumber((p1).length));
          let _rhs87 = p2;
          let _rhs88 = ((_this).f14)[_module.__default.safeIndex((_974_v61)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length))], new BigNumber(((_this).f14).length))];
          let _rhs89 = (_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))];
          let _lhs50 = _990_v76;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_990_v76).length));
          let _lhs52 = _974_v61;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_974_v61).length));
          let _lhs54 = globalState;
          let _lhs55 = globalState;
          _lhs50[_lhs51] = _rhs85;
          r0 = _rhs86;
          _lhs52[_lhs53] = _rhs87;
          _lhs54.f4 = _rhs88;
          _lhs55.f4 = _rhs89;
          let _993_v78;
          let _nw152 = new _module.C1();
          _nw152.__ctor(((_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))]) && ((_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))]), !((_973_v60)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_973_v60).length))]), (_this).f13, (_this).f14, _this.f12);
          _993_v78 = _nw152;
          _993_v78 = _993_v78;
        }
        let _994_v79;
        let _nw153 = new _module.C0();
        _nw153.__ctor();
        _994_v79 = _nw153;
        let _995_v80;
        _995_v80 = _dafny.Seq.of(_this.f12, _this.f12, _this.f12);
        let _996_v81;
        let _nw154 = new _module.C1();
        _nw154.__ctor((_this).f15, (_this).f15, (_this).f13, _dafny.Seq.Concat((_this).f14, (_this).f14), (_995_v80)[_module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_995_v80).length))]);
        _996_v81 = _nw154;
        (globalState).f4 = _996_v81.f18;
      } else {
        (globalState).f4 = (_this).f15;
        let _997_v82;
        let _init28 = ((_998_p2) => function (_999_i10) {
          return _module.__default.safeDivisionInt(_999_i10, _998_p2);
        })(p2);
        let _nw155 = Array((new BigNumber(16)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw155.length); _i0_28++) {
          _nw155[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _997_v82 = _nw155;
        let _1000_v83;
        _1000_v83 = _869_v5;
        let _1001_v84;
        _1001_v84 = _dafny.Map.Empty.slice().updateUnsafe(_997_v82,_1000_v83);
        let _index130 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_997_v82).length));
        (_997_v82)[_index130] = new BigNumber((_1001_v84).length);
        let _1002_v85;
        _1002_v85 = _module.D7.create_DC22((_this).f15);
        let _1003_v86;
        _1003_v86 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let10_0) {
          return function (_1004_dt__update__tmp_h2) {
            return function (_pat_let11_0) {
              return function (_1005_dt__update_hcf34_h0) {
                return _module.D7.create_DC22(_1005_dt__update_hcf34_h0);
              }(_pat_let11_0);
            }((_this).f15);
          }(_pat_let10_0);
        }(_1002_v85),new BigNumber(-494));
        let _index131 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_997_v82).length));
        (_997_v82)[_index131] = (((_1003_v86).contains(_module.D7.create_DC22(((_this).f14)[_module.__default.safeIndex(p2, new BigNumber(((_this).f14).length))]))) ? ((_1003_v86).get(_module.D7.create_DC22(((_this).f14)[_module.__default.safeIndex(p2, new BigNumber(((_this).f14).length))]))) : (p2));
        let _1006_v87;
        let _nw156 = Array((new BigNumber(16)).toNumber()).fill(false);
        _1006_v87 = _nw156;
        let _1007_v88;
        _1007_v88 = _dafny.Seq.of(_1006_v87, _1006_v87, _1006_v87);
        _1007_v88 = _dafny.Seq.update(_1007_v88, _module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_1007_v88).length)), _1006_v87);
        let _1008_v89;
        let _nw157 = new _module.C1();
        _nw157.__ctor((((_867_v3).contains(p2)) ? ((_867_v3).get(p2)) : (_module.__default.fm2((_this).f15, globalState))), (_this).f15, (_this).f13, (_this).f14, _this.f12);
        _1008_v89 = _nw157;
        let _index132 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_997_v82).length));
        (_997_v82)[_index132] = (_997_v82)[_module.__default.safeIndex(new BigNumber(811), new BigNumber((_997_v82).length))];
      }
      let _1009_i11;
      _1009_i11 = _dafny.ZERO;
      L8: {
        while ((_this).f15) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1009_i11)) {
              break L8;
            }
            _1009_i11 = (_1009_i11).plus(_dafny.ONE);
            let _1010_v90;
            _1010_v90 = new BigNumber(683);
            let _1011_v91;
            _1011_v91 = _dafny.Seq.of((_this).f15, !((_this).f15), (_this).f15, (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15)).update((_this).f15, (_this).f15)).length)).isEqualTo(_1010_v90), (_this).f15);
            let _1012_v92;
            let _nw158 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _1012_v92 = _nw158;
            let _rhs90 = new BigNumber(109);
            let _rhs91 = (_this).f14;
            let _rhs92 = p2;
            let _rhs93 = _1012_v92;
            _1010_v90 = _rhs90;
            _1011_v91 = _rhs91;
            _1010_v90 = _rhs92;
            _1012_v92 = _rhs93;
            let _1013_v93;
            _1013_v93 = _module.D6.create_DC18(_1010_v90);
            let _1014_v94;
            _1014_v94 = _dafny.Set.fromElements(_867_v3);
            _1013_v93 = _module.__default.fm46(p1, (_this).f15, _1014_v94, globalState);
            _1010_v90 = _1010_v90;
            let _1015_v95;
            let _nw159 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _1015_v95 = _nw159;
            let _1016_v96;
            _1016_v96 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.UnicodeFromString("naxhcjn"));
            let _1017_v97;
            _1017_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
            let _1018_v98;
            _1018_v98 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f15, globalState),p2);
            let _index133 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1015_v95).length));
            (_1015_v95)[_index133] = (((_1016_v96).contains(_module.__default.fm0((_this).f15, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1017_v97).length),_1018_v98), (_this).f15, globalState))) ? ((_1016_v96).get(_module.__default.fm0((_this).f15, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1017_v97).length),_1018_v98), (_this).f15, globalState))) : (p1));
            let _index134 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1015_v95).length));
            (_1015_v95)[_index134] = p1;
          }
        }
      }
      let _1019_v99;
      let _nw160 = new _module.C0();
      _nw160.__ctor();
      _1019_v99 = _nw160;
      r0 = (_this).f15;
      return r0;
    }
    m6(globalState) {
      let _this = this;
      let r0 = false;
      let _1020_v0;
      _1020_v0 = new BigNumber(58);
      let _1021_v1;
      _1021_v1 = _dafny.Set.fromElements(_1020_v0);
      let _1022_v2;
      _1022_v2 = _dafny.Map.Empty.slice().updateUnsafe(true,_1020_v0);
      let _1023_v3;
      _1023_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1021_v1).length),_1022_v2);
      let _1024_v4;
      _1024_v4 = _module.D11.create_DC30(_1023_v3);
      let _1025_v5;
      _1025_v5 = _dafny.Set.fromElements((_this).f15, (_this).f15, (_this).f15);
      let _1026_v6;
      let _nw161 = Array((new BigNumber(26)).toNumber());
      _nw161[(_dafny.ZERO).toNumber()] = _1020_v0;
      _nw161[(_dafny.ONE).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(2)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(3)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1020_v0));
      _nw161[(new BigNumber(5)).toNumber()] = new BigNumber(((_this).f14).length);
      _nw161[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_1020_v0);
      _nw161[(new BigNumber(7)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(8)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(9)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(10)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(11)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(12)).toNumber()] = new BigNumber(((_this).f16).length);
      _nw161[(new BigNumber(13)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(14)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(15)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(16)).toNumber()] = _module.__default.fm0(_module.__default.fm2((_this).f15, globalState), (_1024_v4).dtor_cf45, (_this).f15, globalState);
      _nw161[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_1020_v0);
      _nw161[(new BigNumber(18)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(19)).toNumber()] = new BigNumber(522);
      _nw161[(new BigNumber(20)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(21)).toNumber()] = new BigNumber((_1025_v5).length);
      _nw161[(new BigNumber(22)).toNumber()] = _1020_v0;
      _nw161[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("kjmqcatm")).length);
      _nw161[(new BigNumber(24)).toNumber()] = new BigNumber(-426);
      _nw161[(new BigNumber(25)).toNumber()] = _1020_v0;
      _1026_v6 = _nw161;
      let _1027_v7;
      _1027_v7 = _dafny.Seq.of(_1026_v6, _1026_v6);
      let _1028_v8;
      let _nw162 = Array((new BigNumber(22)).toNumber()).fill([]);
      _1028_v8 = _nw162;
      let _1029_v9;
      let _nw163 = Array((new BigNumber(8)).toNumber()).fill(false);
      _1029_v9 = _nw163;
      let _index135 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1028_v8).length));
      (_1028_v8)[_index135] = _1029_v9;
      let _1030_v10;
      _1030_v10 = _module.D0.create_DC1();
      let _pat_let_tv11 = _1025_v5;
      let _pat_let_tv12 = _1020_v0;
      let _pat_let_tv13 = _1020_v0;
      let _index136 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1028_v8).length));
      let _rhs94 = _1027_v7;
      let _rhs95 = _1029_v9;
      let _rhs96 = new BigNumber((function (_source15) {
        if (_source15.is_DC1) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(767)), ((_1031_v5) => function (_1032_i0) {
            return new BigNumber((_1031_v5).length);
          })(_pat_let_tv11));
        } else if (_source15.is_DC2) {
          let _1033___mcc_h0 = (_source15).cf1;
          let _1034___mcc_h1 = (_source15).cf2;
          let _1035_cf2 = _1034___mcc_h1;
          let _1036_cf1 = _1033___mcc_h0;
          return _dafny.Seq.Concat(_dafny.Seq.update((_this).f16, _module.__default.safeIndex(_pat_let_tv12, new BigNumber(((_this).f16).length)), _1036_cf1), _dafny.Seq.of(_1036_cf1, _pat_let_tv13));
        } else {
          let _1037___mcc_h2 = (_source15).cf0;
          let _1038_cf0 = _1037___mcc_h2;
          return (_this).f16;
        }
      }(_1030_v10)).length);
      let _rhs97 = _module.__default.fm2((_this).f15, globalState);
      let _lhs56 = _1028_v8;
      let _lhs57 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1028_v8).length));
      let _lhs58 = globalState;
      _1027_v7 = _rhs94;
      _lhs56[_lhs57] = _rhs95;
      _1020_v0 = _rhs96;
      _lhs58.f4 = _rhs97;
      let _rhs98 = (_this).f15;
      let _rhs99 = ((_this).f16)[_module.__default.safeIndex(new BigNumber((_1025_v5).length), new BigNumber(((_this).f16).length))];
      let _lhs59 = globalState;
      _lhs59.f4 = _rhs98;
      _1020_v0 = _rhs99;
      let _1039_i1;
      _1039_i1 = _dafny.ZERO;
      L9: {
        while ((_this).f15) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1039_i1)) {
              break L9;
            }
            _1039_i1 = (_1039_i1).plus(_dafny.ONE);
            let _1040_v11;
            _1040_v11 = _dafny.MultiSet.fromElements((_this).f15, (_this).f15);
            let _1041_v12;
            _1041_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1020_v0,(_this).fm38(new BigNumber(233), new BigNumber((_1040_v11).cardinality()), _1020_v0, (_this).f15, globalState));
            _1041_v12 = (_1041_v12).update(_1020_v0, _1020_v0);
            let _1042_v13;
            _1042_v13 = _module.D2.create_DC6((_1028_v8)[_module.__default.safeIndex(new BigNumber(820), new BigNumber((_1028_v8).length))], _1029_v9);
            let _source16 = _1042_v13;
            if (_source16.is_DC5) {
              let _1043___mcc_h3 = (_source16).cf5;
              let _1044___mcc_h4 = (_source16).cf6;
              let _1045___mcc_h5 = (_source16).cf7;
              let _1046_cf7 = _1045___mcc_h5;
              let _1047_cf6 = _1044___mcc_h4;
              let _1048_cf5 = _1043___mcc_h3;
              let _1049_v14;
              let _nw164 = new _module.C0();
              _nw164.__ctor();
              _1049_v14 = _nw164;
              let _1050_v15;
              let _nw165 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
              _1050_v15 = _nw165;
              let _1051_v16;
              _1051_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
              let _index137 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1050_v15).length));
              (_1050_v15)[_index137] = _1051_v16;
              let _index138 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_1050_v15).length));
              (_1050_v15)[_index138] = ((((_this).f13).contains(_1046_cf7)) ? (((_this).f13).get(_1046_cf7)) : (_1051_v16));
              let _1052_v17;
              let _nw166 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
              _1052_v17 = _nw166;
              let _1053_v18;
              _1053_v18 = _dafny.MultiSet.fromElements(_1020_v0);
              let _1054_v19;
              _1054_v19 = _dafny.Seq.of(_module.D7.create_DC20(_1053_v18), _module.D7.create_DC20(_1053_v18));
              let _1055_v20;
              _1055_v20 = _dafny.Seq.of(_1054_v19, _1054_v19, _dafny.Seq.update(_1054_v19, _module.__default.safeIndex(_1020_v0, new BigNumber((_1054_v19).length)), _module.D7.create_DC20(_1053_v18)), _1054_v19, _1054_v19);
              let _index139 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1052_v17).length));
              (_1052_v17)[_index139] = _dafny.Seq.Concat(_1055_v20, _1055_v20);
              let _index140 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_1052_v17).length));
              (_1052_v17)[_index140] = _1055_v20;
              let _1056_v21;
              let _init29 = ((_1057_v11) => function (_1058_i2) {
                return _1057_v11;
              })(_1040_v11);
              let _nw167 = Array((new BigNumber(11)).toNumber());
              for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw167.length); _i0_29++) {
                _nw167[_i0_29] = _init29(new BigNumber(_i0_29));
              }
              _1056_v21 = _nw167;
              let _index141 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1056_v21).length));
              (_1056_v21)[_index141] = _1040_v11;
              let _index142 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1056_v21).length));
              (_1056_v21)[_index142] = ((_1040_v11).update(true, _module.__default.abs(new BigNumber(-767)))).Union(_1040_v11);
            } else if (_source16.is_DC6) {
              let _1059___mcc_h6 = (_source16).cf8;
              let _1060___mcc_h7 = (_source16).cf9;
              let _1061_cf9 = _1060___mcc_h7;
              let _1062_cf8 = _1059___mcc_h6;
              let _1063_v22;
              _1063_v22 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f15);
              let _1064_v23;
              _1064_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1063_v22,true);
              _1064_v23 = (_1064_v23).update(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15), (_this).f15);
              let _1065_v24;
              _1065_v24 = _dafny.Seq.UnicodeFromString("jrgw");
              _1020_v0 = new BigNumber((_dafny.Seq.Concat(_1065_v24, _dafny.Seq.Concat(_1065_v24, _1065_v24))).length);
              _1065_v24 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cvucslniw"), _1065_v24), _dafny.Seq.Concat(_1065_v24, _dafny.Seq.UnicodeFromString("kjffmj")));
              let _1066_v25;
              _1066_v25 = _module.D8.create_DC26((_this).f15);
              let _1067_v26;
              _1067_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1062_cf8,_1066_v25);
              let _1068_v27;
              _1068_v27 = _module.D12.create_DC34(_1067_v26);
              _1067_v26 = (_1068_v27).dtor_cf53;
            } else if (_source16.is_DC4) {
              let _1069___mcc_h8 = (_source16).cf4;
              let _1070_cf4 = _1069___mcc_h8;
              (globalState).f4 = (_this).f15;
              (globalState).f4 = (_this).f15;
              let _1071_v28;
              let _init30 = function (_1072_i3) {
                return new _dafny.CodePoint('s'.codePointAt(0));
              };
              let _nw168 = Array((new BigNumber(28)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw168.length); _i0_30++) {
                _nw168[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _1071_v28 = _nw168;
              let _index143 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1071_v28).length));
              (_1071_v28)[_index143] = new _dafny.CodePoint('y'.codePointAt(0));
              let _1073_v29;
              _1073_v29 = new _dafny.CodePoint('a'.codePointAt(0));
              let _index144 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1071_v28).length));
              (_1071_v28)[_index144] = _1073_v29;
              let _1074_v30;
              _1074_v30 = _dafny.Seq.UnicodeFromString("dqykhwa");
              _1020_v0 = new BigNumber((_dafny.Seq.Concat(_1074_v30, _dafny.Seq.UnicodeFromString("gpcuoqb"))).length);
            } else {
              let _1075___mcc_h9 = (_source16).cf10;
              let _1076_cf10 = _1075___mcc_h9;
              let _index145 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1026_v6).length));
              (_1026_v6)[_index145] = _1020_v0;
              let _index146 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1026_v6).length));
              (_1026_v6)[_index146] = (_dafny.ZERO).minus(_1020_v0);
              (globalState).f4 = (_this).f15;
              (globalState).f4 = !(_1020_v0).isEqualTo((_1026_v6)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_1026_v6).length))]);
              let _1077_v31;
              _1077_v31 = _module.D2.create_DC5(new BigNumber((_dafny.Seq.UnicodeFromString("vljqrh")).length), (_this).f15, (_this).f15);
              let _index147 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_1026_v6).length));
              (_1026_v6)[_index147] = ((true) ? (_1020_v0) : ((_1077_v31).dtor_cf5));
            }
            let _1078_v32;
            let _init31 = function (_1079_i4) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-301)), function (_1080_i5) {
                return new _dafny.CodePoint('q'.codePointAt(0));
              });
            };
            let _nw169 = Array((new BigNumber(11)).toNumber());
            for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw169.length); _i0_31++) {
              _nw169[_i0_31] = _init31(new BigNumber(_i0_31));
            }
            _1078_v32 = _nw169;
            let _1081_v33;
            _1081_v33 = _module.D12.create_DC36(_1078_v32, _1020_v0, _module.__default.fm47(globalState), _dafny.Set.fromElements(_1029_v9, _1029_v9));
            let _index148 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1029_v9).length));
            (_1029_v9)[_index148] = _module.__default.fm2(!((_this).f15), globalState);
            let _1082_v34;
            _1082_v34 = _dafny.Seq.UnicodeFromString("npgudenrk");
            let _index149 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1029_v9).length));
            let _rhs100 = (_this).f15;
            let _rhs101 = _1081_v33;
            let _rhs102 = _dafny.Seq.IsProperPrefixOf(_1082_v34, (((_this).f15) ? (_1082_v34) : (_dafny.Seq.UnicodeFromString("eghjdwmk"))));
            let _rhs103 = ((_this).f15) || ((_this).f15);
            let _lhs60 = globalState;
            let _lhs61 = globalState;
            let _lhs62 = _1029_v9;
            let _lhs63 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1029_v9).length));
            _lhs60.f4 = _rhs100;
            _1081_v33 = _rhs101;
            _lhs61.f4 = _rhs102;
            _lhs62[_lhs63] = _rhs103;
            _1020_v0 = _1020_v0;
          }
        }
      }
      _1020_v0 = _1020_v0;
      let _pat_let_tv14 = _1023_v3;
      let _pat_let_tv15 = _1023_v3;
      let _1083_v35;
      let _nw170 = Array((new BigNumber(15)).toNumber());
      _nw170[(_dafny.ZERO).toNumber()] = _1024_v4;
      _nw170[(_dafny.ONE).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(2)).toNumber()] = ((true) ? (_1024_v4) : (_1024_v4));
      _nw170[(new BigNumber(3)).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(4)).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(5)).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(6)).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(7)).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(8)).toNumber()] = function (_pat_let12_0) {
        return function (_1086_dt__update__tmp_h1) {
          return function (_pat_let15_0) {
            return function (_1087_dt__update_hcf45_h1) {
              return _module.D11.create_DC30(_1087_dt__update_hcf45_h1);
            }(_pat_let15_0);
          }(_pat_let_tv15);
        }(_pat_let12_0);
      }(function (_pat_let13_0) {
        return function (_1084_dt__update__tmp_h0) {
          return function (_pat_let14_0) {
            return function (_1085_dt__update_hcf45_h0) {
              return _module.D11.create_DC30(_1085_dt__update_hcf45_h0);
            }(_pat_let14_0);
          }(_pat_let_tv14);
        }(_pat_let13_0);
      }(_1024_v4));
      _nw170[(new BigNumber(9)).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(10)).toNumber()] = _1024_v4;
      _nw170[(new BigNumber(11)).toNumber()] = _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(_1020_v0,_1022_v2));
      _nw170[(new BigNumber(12)).toNumber()] = _module.__default.fm48(globalState);
      _nw170[(new BigNumber(13)).toNumber()] = _module.D11.create_DC30(_1023_v3);
      _nw170[(new BigNumber(14)).toNumber()] = _module.__default.fm48(globalState);
      _1083_v35 = _nw170;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1083_v35).length))) {
        let _1088_i6 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1088_i6)) && ((_1088_i6).isLessThan(new BigNumber((_1083_v35).length))))) {
          (_1083_v35)[(_1088_i6)] = _1024_v4;
        }
      }
      if ((_this).f15) {
        _1022_v2 = (_1022_v2).Merge((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f15),_1020_v0)).Merge(_1022_v2));
        let _1089_v36;
        let _nw171 = new _module.C1();
        _nw171.__ctor((_this).f15, _dafny.areEqual((_this).f16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(202)), ((_1090_v0) => function (_1091_i7) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1090_v0,_1090_v0)).length);
        })(_1020_v0))), (_this).f13, _dafny.Seq.update((_this).f14, _module.__default.safeIndex(_1020_v0, new BigNumber(((_this).f14).length)), (_this).f15), _this.f12);
        _1089_v36 = _nw171;
        let _index150 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1028_v8).length));
        (_1028_v8)[_index150] = _1029_v9;
        let _1092_v37;
        _1092_v37 = _module.D11.create_DC31();
        _1092_v37 = _module.D11.create_DC31();
        _1020_v0 = (_dafny.ZERO).minus(new BigNumber(-572));
      } else {
        let _1093_v38;
        _1093_v38 = _module.D6.create_DC18(_1020_v0);
        let _1094_v39;
        _1094_v39 = _dafny.Seq.of((_1020_v0).plus(new BigNumber(((_this).f14).length)));
        let _rhs104 = _1093_v38;
        let _rhs105 = _dafny.Seq.Concat(_dafny.Seq.update((_this).f16, _module.__default.safeIndex(new BigNumber(-437), new BigNumber(((_this).f16).length)), _1020_v0), (_this).f16);
        _1093_v38 = _rhs104;
        _1094_v39 = _rhs105;
        _1020_v0 = new BigNumber(709);
        let _1095_v40;
        _1095_v40 = _dafny.Seq.of(_1029_v9, _1029_v9);
        let _1096_v41;
        _1096_v41 = _dafny.MultiSet.fromElements((_1095_v40)[_module.__default.safeIndex(_1020_v0, new BigNumber((_1095_v40).length))], _1029_v9);
        let _1097_v42;
        _1097_v42 = new _dafny.CodePoint('b'.codePointAt(0));
        let _1098_v43;
        _1098_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1096_v41).Union(_1096_v41),_1097_v42);
        _1098_v43 = (_1098_v43).update((_1096_v41).Union(_1096_v41), _1097_v42);
        let _1099_v44;
        let _nw172 = Array((new BigNumber(15)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = _1026_v6;
        _nw172[(_dafny.ONE).toNumber()] = ((_module.__default.fm2((_this).f15, globalState)) ? ((_1027_v7)[_module.__default.safeIndex(_1020_v0, new BigNumber((_1027_v7).length))]) : (_1026_v6));
        _nw172[(new BigNumber(2)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(3)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(4)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(5)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(6)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(7)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(8)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(9)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(10)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(11)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(12)).toNumber()] = _1026_v6;
        _nw172[(new BigNumber(13)).toNumber()] = (_1027_v7)[_module.__default.safeIndex(_1020_v0, new BigNumber((_1027_v7).length))];
        _nw172[(new BigNumber(14)).toNumber()] = _1026_v6;
        _1099_v44 = _nw172;
        _1099_v44 = (((_this).f15) ? (_1099_v44) : (_1099_v44));
        let _1100_v45;
        _1100_v45 = _module.D5.create_DC16((_this).f15);
        let _source17 = _1100_v45;
        if (_source17.is_DC16) {
          let _1101___mcc_h10 = (_source17).cf26;
          let _1102_cf26 = _1101___mcc_h10;
          _1020_v0 = (_dafny.ZERO).minus(_1020_v0);
          let _1103_v46;
          _1103_v46 = _module.D5.create_DC15((_this).f16);
          let _index151 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length));
          (_1029_v9)[_index151] = _1102_cf26;
          let _1104_v47;
          _1104_v47 = _dafny.MultiSet.fromElements((_this).f16, _dafny.Seq.of(_1020_v0), _dafny.Seq.update((_this).f16, _module.__default.safeIndex(_1020_v0, new BigNumber(((_this).f16).length)), _1020_v0), (_this).f16);
          let _1105_v48;
          _1105_v48 = _dafny.Seq.of(_1104_v47, _1104_v47);
          let _index152 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length));
          let _rhs106 = _1103_v46;
          let _rhs107 = ((_1104_v47).Intersect(_1104_v47)).IsSubsetOf((_1104_v47).Union((_1105_v48)[_module.__default.safeIndex(new BigNumber(149), new BigNumber((_1105_v48).length))]));
          let _rhs108 = (((_this).f15) === (_1102_cf26)) && ((_this).f15);
          let _rhs109 = _1093_v38;
          let _rhs110 = ((_1102_cf26) ? (!((_this).f15)) : (!((_this).f15)));
          let _lhs64 = globalState;
          let _lhs65 = globalState;
          let _lhs66 = _1029_v9;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length));
          _1103_v46 = _rhs106;
          _lhs64.f4 = _rhs107;
          _lhs65.f4 = _rhs108;
          _1093_v38 = _rhs109;
          _lhs66[_lhs67] = _rhs110;
          (globalState).f4 = _1102_cf26;
          let _1106_v49;
          _1106_v49 = _dafny.Seq.of((_1029_v9)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length))], (_1029_v9)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length))], (_1029_v9)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length))]);
          _1106_v49 = _dafny.Seq.Concat(_dafny.Seq.of((_1029_v9)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length))], _1102_cf26, true, !((_1029_v9)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_1029_v9).length))]), _module.__default.fm2((_1106_v49)[_module.__default.safeIndex(new BigNumber(-489), new BigNumber((_1106_v49).length))], globalState)), _1106_v49);
        } else {
          let _1107___mcc_h11 = (_source17).cf25;
          let _1108_cf25 = _1107___mcc_h11;
          _1020_v0 = new BigNumber(((_this).f14).length);
          let _1109_v50;
          _1109_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1029_v9);
          _1109_v50 = (_1109_v50).update((_1021_v1).IsProperSubsetOf(_1021_v1), (_1028_v8)[_module.__default.safeIndex(new BigNumber(820), new BigNumber((_1028_v8).length))]);
          let _1110_v51;
          _1110_v51 = _dafny.Seq.of(_1097_v42);
          let _1111_v52;
          _1111_v52 = _module.D8.create_DC25((_1025_v5).IsSubsetOf(_1025_v5), (((_this).f15) ? (new BigNumber((_1110_v51).length)) : (_1020_v0)), _1020_v0, (_1022_v2).Merge(_1022_v2));
          _1111_v52 = _1111_v52;
          let _1112_v53;
          let _1113_v54;
          let _out87;
          let _out88;
          let _outcollector22 = (_this).m11(globalState);
          _out87 = _outcollector22[0];
          _out88 = _outcollector22[1];
          _1112_v53 = _out87;
          _1113_v54 = _out88;
        }
      }
      let _1114_v55;
      let _nw173 = new _module.C1();
      _nw173.__ctor((_this).f15, true, (_this).f13, (_this).f14, _this.f12);
      _1114_v55 = _nw173;
      let _1115_v56;
      _1115_v56 = _module.D3.create_DC10(!((_this).f15), _1020_v0, _1114_v55);
      r0 = !((_1115_v56).dtor_cf14);
      return r0;
    }
    m11(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = false;
      let _1116_v0;
      _1116_v0 = _dafny.Seq.UnicodeFromString("cyjxdokyt");
      let _1117_v1;
      _1117_v1 = _dafny.MultiSet.fromElements(_1116_v0, _1116_v0, _1116_v0, _1116_v0, _1116_v0);
      let _1118_v2;
      _1118_v2 = new BigNumber(132);
      r1 = (!_dafny.Seq.contains((_this).f14, _module.__default.fm2((_this).f15, globalState))) && ((_dafny.MultiSet.fromElements(_1116_v0, _dafny.Seq.UnicodeFromString("rpqincvt"), _1116_v0, _1116_v0, _module.__default.fm40(_module.D4.create_DC14((_this).f15), _1118_v2, false, (_this).f15, globalState))).IsSubsetOf(_1117_v1));
      let _1119_v3;
      _1119_v3 = _module.D0.create_DC2(new BigNumber((_1116_v0).length), (_this).f15);
      let _pat_let_tv16 = _1117_v1;
      let _source18 = function (_source19) {
        if (_source19.is_DC1) {
          return _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber((_pat_let_tv16).cardinality())));
        } else if (_source19.is_DC2) {
          let _1120___mcc_h8 = (_source19).cf1;
          let _1121___mcc_h9 = (_source19).cf2;
          let _1122_cf2 = _1121___mcc_h9;
          let _1123_cf1 = _1120___mcc_h8;
          return _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(!(true),_1123_cf1));
        } else {
          let _1124___mcc_h10 = (_source19).cf0;
          let _1125_cf0 = _1124___mcc_h10;
          return _module.D8.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(_1125_cf0,new BigNumber(117)));
        }
      }(_1119_v3);
      if (_source18.is_DC24) {
        let _1126___mcc_h0 = (_source18).cf36;
        let _1127_cf36 = _1126___mcc_h0;
        let _1128_v4;
        _1128_v4 = _dafny.Set.fromElements(_1127_cf36, (_this).f15, true);
        let _1129_v5;
        _1129_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1128_v4).Union(_1128_v4)).length),_1127_cf36);
        let _1130_v6;
        _1130_v6 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1131_v7;
        _1131_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1130_v6,(((_1129_v5).contains(_1118_v2)) ? ((_1129_v5).get(_1118_v2)) : (_1127_cf36)));
        _1129_v5 = (_1129_v5).update(new BigNumber((_1131_v7).length), (_this).f15);
        let _1132_v8;
        _1132_v8 = _dafny.MultiSet.fromElements(_1127_cf36, false);
        (globalState).f4 = (((_this).fm38(_1118_v2, _1118_v2, (_dafny.ZERO).minus(_1118_v2), (_this).f15, globalState)).minus(_1118_v2)).isLessThanOrEqualTo(new BigNumber((_1132_v8).cardinality()));
        let _1133_v9;
        let _nw174 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
        _1133_v9 = _nw174;
        let _1134_v10;
        _1134_v10 = _module.D4.create_DC12(_dafny.Seq.of((_this).f15, _1127_cf36));
        let _1135_v11;
        _1135_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1127_cf36,(_this).f14);
        let _nw175 = Array((new BigNumber(24)).toNumber());
        _nw175[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_this).f14, (_this).f14);
        _nw175[(_dafny.ONE).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(3)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(4)).toNumber()] = (_1134_v10).dtor_cf18;
        _nw175[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_this).f14, (_this).f14);
        _nw175[(new BigNumber(6)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(7)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(8)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(9)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(10)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(11)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(12)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_module.__default.fm2((_this).f15, globalState));
        _nw175[(new BigNumber(14)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(15)).toNumber()] = _dafny.Seq.of((_this).f15, _1127_cf36);
        _nw175[(new BigNumber(16)).toNumber()] = _dafny.Seq.of((_this).f15, _1127_cf36);
        _nw175[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat((_this).f14, (_this).f14);
        _nw175[(new BigNumber(18)).toNumber()] = (((_1135_v11).contains(_1127_cf36)) ? ((_1135_v11).get(_1127_cf36)) : ((_this).f14));
        _nw175[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat((_this).f14, (_this).f14);
        _nw175[(new BigNumber(20)).toNumber()] = _dafny.Seq.of((_this).f15, (_this).f15);
        _nw175[(new BigNumber(21)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(22)).toNumber()] = (_this).f14;
        _nw175[(new BigNumber(23)).toNumber()] = (_this).f14;
        _1133_v9 = _nw175;
        (globalState).f4 = (!_dafny.Seq.contains(_1116_v0, _1130_v6)) || (_1127_cf36);
      } else if (_source18.is_DC25) {
        let _1136___mcc_h1 = (_source18).cf37;
        let _1137___mcc_h2 = (_source18).cf38;
        let _1138___mcc_h3 = (_source18).cf39;
        let _1139___mcc_h4 = (_source18).cf40;
        let _1140_cf40 = _1139___mcc_h4;
        let _1141_cf39 = _1138___mcc_h3;
        let _1142_cf38 = _1137___mcc_h2;
        let _1143_cf37 = _1136___mcc_h1;
        let _1144_v12;
        let _nw176 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1144_v12 = _nw176;
        let _1145_v13;
        _1145_v13 = new _dafny.CodePoint('l'.codePointAt(0));
        let _index153 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1144_v12).length));
        (_1144_v12)[_index153] = _1145_v13;
        let _index154 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1144_v12).length));
        let _rhs111 = (_1118_v2).minus(_1142_cf38);
        let _rhs112 = (_module.__default.safeDivisionInt(new BigNumber(421), _1141_cf39)).minus(_1118_v2);
        let _rhs113 = _module.__default.fm2((_this).f15, globalState);
        let _rhs114 = _1145_v13;
        let _rhs115 = _1145_v13;
        let _lhs68 = _1144_v12;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1144_v12).length));
        _1142_cf38 = _rhs111;
        _1118_v2 = _rhs112;
        _1143_cf37 = _rhs113;
        r0 = _rhs114;
        _lhs68[_lhs69] = _rhs115;
        let _1146_v14;
        let _nw177 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1146_v14 = _nw177;
        let _index155 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_1146_v14).length));
        (_1146_v14)[_index155] = (_1142_cf38).plus(_1118_v2);
        let _index156 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_1146_v14).length));
        (_1146_v14)[_index156] = _1142_cf38;
        let _1147_v15;
        let _1148_v16;
        let _1149_v17;
        let _1150_v18;
        let _out89;
        let _out90;
        let _out91;
        let _out92;
        let _outcollector23 = _module.__default.m0((_this).f15, (((_this).f15) ? (_1119_v3) : (_1119_v3)), _dafny.Seq.contains(_1116_v0, _1145_v13), globalState);
        _out89 = _outcollector23[0];
        _out90 = _outcollector23[1];
        _out91 = _outcollector23[2];
        _out92 = _outcollector23[3];
        _1147_v15 = _out89;
        _1148_v16 = _out90;
        _1149_v17 = _out91;
        _1150_v18 = _out92;
        if (!((_1148_v16) && (!(_1150_v18)))) {
          let _1151_v19;
          _1151_v19 = _module.D4.create_DC14(true);
          let _1152_v20;
          _1152_v20 = _module.D4.create_DC12((_this).f14);
          let _1153_v21;
          let _nw178 = Array((new BigNumber(16)).toNumber());
          _nw178[(_dafny.ZERO).toNumber()] = (_this).f14;
          _nw178[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_1151_v19).dtor_cf24, _1148_v16, _1150_v18, !(!(_1150_v18)), _1148_v16);
          _nw178[(new BigNumber(2)).toNumber()] = (_1152_v20).dtor_cf18;
          _nw178[(new BigNumber(3)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(4)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(5)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(6)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(7)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(8)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm2(_1150_v18, globalState)), (_this).f14);
          _nw178[(new BigNumber(10)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1143_cf37), (_this).f14);
          _nw178[(new BigNumber(12)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(13)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(14)).toNumber()] = (_this).f14;
          _nw178[(new BigNumber(15)).toNumber()] = (_this).f14;
          _1153_v21 = _nw178;
          let _1154_v22;
          _1154_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1149_v17,_1143_cf37);
          let _1155_v23;
          _1155_v23 = _dafny.Map.Empty.slice().updateUnsafe((((_1154_v22).contains(_1118_v2)) ? ((_1154_v22).get(_1118_v2)) : ((_this).f15)),(_this).f15);
          let _index157 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1153_v21).length));
          (_1153_v21)[_index157] = _dafny.Seq.of((_this).f15, !(!((((_1155_v23).contains((_this).f15)) ? ((_1155_v23).get((_this).f15)) : (_1143_cf37)))), _1143_cf37);
          let _1156_v24;
          _1156_v24 = _dafny.MultiSet.fromElements((_this).f15);
          let _1157_v25;
          _1157_v25 = _1156_v24;
          let _index158 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1153_v21).length));
          (_1153_v21)[_index158] = _dafny.Seq.Concat(_module.__default.fm43(_1157_v25, _module.D7.create_DC22((_this).f15), globalState), _dafny.Seq.Concat((_this).f14, (_this).f14));
          let _1158_v26;
          _1158_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(575),_1118_v2);
          let _1159_v27;
          _1159_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm49(_1143_cf37, _1148_v16, _1148_v16, _module.__default.fm39(new BigNumber((_1158_v26).length), (_this).f15, _1141_cf39, _1116_v0, globalState), globalState),_1148_v16);
          _1159_v27 = _1159_v27;
          let _1160_v28;
          let _nw179 = Array((new BigNumber(14)).toNumber()).fill(false);
          _1160_v28 = _nw179;
          _1160_v28 = _1160_v28;
          _1160_v28 = _1160_v28;
          let _1161_v30;
          _1161_v30 = _module.D6.create_DC17(function () {
  let _coll47 = new _dafny.Set();
  for (const _compr_47 of _dafny.IntegerRange(new BigNumber(908), new BigNumber(-420))) {
    let _1162_v29 = _compr_47;
    if (((new BigNumber(908)).isLessThanOrEqualTo(_1162_v29)) && ((_1162_v29).isLessThan(new BigNumber(-420)))) {
      _coll47.add((_1162_v29).multipliedBy(new BigNumber(-300)));
    }
  }
  return _coll47;
}());
          let _1163_v31;
          _1163_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1161_v30,_module.__default.safeModuloInt(_1141_cf39, _1142_cf38));
          _1163_v31 = (_1163_v31).update(_1161_v30, new BigNumber(664));
        } else {
          _1149_v17 = _module.__default.safeModuloInt(_1141_cf39, _1142_cf38);
          let _1164_v32;
          _1164_v32 = _dafny.MultiSet.fromElements((_this).f15, _1143_cf37);
          let _1165_v33;
          _1165_v33 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),_1118_v2);
          let _1166_v34;
          _1166_v34 = _dafny.MultiSet.fromElements(_1142_cf38);
          let _1167_v35;
          let _nw180 = Array((new BigNumber(20)).toNumber());
          _nw180[(_dafny.ZERO).toNumber()] = (_1164_v32).IsSubsetOf(_1164_v32);
          _nw180[(_dafny.ONE).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm39(_1141_cf39, (_this).f15, (_dafny.ZERO).minus(_1149_v17), _1116_v0, globalState),_1149_v17)).equals(_1165_v33);
          _nw180[(new BigNumber(2)).toNumber()] = (!(true)) === (_1143_cf37);
          _nw180[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(231)), function (_1168_i0) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }), _1116_v0);
          _nw180[(new BigNumber(4)).toNumber()] = (_this).f15;
          _nw180[(new BigNumber(5)).toNumber()] = ((_1143_cf37) ? (_1150_v18) : (_1150_v18));
          _nw180[(new BigNumber(6)).toNumber()] = _1150_v18;
          _nw180[(new BigNumber(7)).toNumber()] = _1143_cf37;
          _nw180[(new BigNumber(8)).toNumber()] = ((_1143_cf37) ? (false) : ((_this).f15));
          _nw180[(new BigNumber(9)).toNumber()] = (_this).f15;
          _nw180[(new BigNumber(10)).toNumber()] = (_this).f15;
          _nw180[(new BigNumber(11)).toNumber()] = (_this).f15;
          _nw180[(new BigNumber(12)).toNumber()] = ((_this).f14)[_module.__default.safeIndex(_1141_cf39, new BigNumber(((_this).f14).length))];
          _nw180[(new BigNumber(13)).toNumber()] = _1148_v16;
          _nw180[(new BigNumber(14)).toNumber()] = (_this).f15;
          _nw180[(new BigNumber(15)).toNumber()] = false;
          _nw180[(new BigNumber(16)).toNumber()] = _1150_v18;
          _nw180[(new BigNumber(17)).toNumber()] = _1143_cf37;
          _nw180[(new BigNumber(18)).toNumber()] = ((((_1166_v34).contains(new BigNumber((_1147_v15).length))) ? ((_1166_v34).get(new BigNumber((_1147_v15).length))) : (_1141_cf39))).isLessThanOrEqualTo(_1142_cf38);
          _nw180[(new BigNumber(19)).toNumber()] = _1148_v16;
          _1167_v35 = _nw180;
          let _index159 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1167_v35).length));
          (_1167_v35)[_index159] = (!((_this).f15)) || (_1143_cf37);
          let _index160 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1167_v35).length));
          (_1167_v35)[_index160] = _1148_v16;
          let _1169_v36;
          _1169_v36 = _dafny.MultiSet.fromElements(_1167_v35);
          let _index161 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1167_v35).length));
          (_1167_v35)[_index161] = (_1169_v36).IsSubsetOf(_dafny.MultiSet.fromElements(_1167_v35, _1167_v35));
          let _1170_v37;
          _1170_v37 = _dafny.Map.Empty.slice().updateUnsafe((_1140_cf40).update((_1167_v35)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_1167_v35).length))], new BigNumber(293)),_1116_v0);
          let _1171_v38;
          _1171_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1148_v16,(_1141_cf39).isEqualTo(_1149_v17));
          let _index162 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_1146_v14).length));
          let _rhs116 = ((_1170_v37).Merge(_1170_v37)).Merge(_1170_v37);
          let _rhs117 = (_1147_v15)[_module.__default.safeIndex((_dafny.ZERO).minus(_1118_v2), new BigNumber((_1147_v15).length))];
          let _rhs118 = !((((_1171_v38).contains(_module.__default.fm2((_1167_v35)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_1167_v35).length))], globalState))) ? ((_1171_v38).get(_module.__default.fm2((_1167_v35)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_1167_v35).length))], globalState))) : (false)));
          let _rhs119 = _1167_v35;
          let _lhs70 = _1146_v14;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_1146_v14).length));
          let _lhs72 = globalState;
          _1170_v37 = _rhs116;
          _lhs70[_lhs71] = _rhs117;
          _lhs72.f4 = _rhs118;
          _1167_v35 = _rhs119;
          let _index163 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1167_v35).length));
          (_1167_v35)[_index163] = _1143_cf37;
        }
      } else if (_source18.is_DC26) {
        let _1172___mcc_h5 = (_source18).cf41;
        let _1173_cf41 = _1172___mcc_h5;
        let _1174_v39;
        let _nw181 = new _module.C0();
        _nw181.__ctor();
        _1174_v39 = _nw181;
        let _1175_v40;
        _1175_v40 = _dafny.Seq.of((_this).f15);
        let _1176_v41;
        _1176_v41 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1118_v2),_1116_v0);
        let _1177_v42;
        _1177_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1173_cf41,new BigNumber(((((_1176_v41).contains(_1118_v2)) ? ((_1176_v41).get(_1118_v2)) : (_1116_v0))).length));
        let _1178_v43;
        let _nw182 = Array((new BigNumber(15)).toNumber());
        _nw182[(_dafny.ZERO).toNumber()] = (_1118_v2).isLessThan(new BigNumber(((_this).f16).length));
        _nw182[(_dafny.ONE).toNumber()] = (_this).f15;
        _nw182[(new BigNumber(2)).toNumber()] = (_1118_v2).isLessThanOrEqualTo((((_1177_v42).contains((_this).f15)) ? ((_1177_v42).get((_this).f15)) : (_1118_v2)));
        _nw182[(new BigNumber(3)).toNumber()] = _1173_cf41;
        _nw182[(new BigNumber(4)).toNumber()] = false;
        _nw182[(new BigNumber(5)).toNumber()] = (_this).f15;
        _nw182[(new BigNumber(6)).toNumber()] = (_this).f15;
        _nw182[(new BigNumber(7)).toNumber()] = true;
        _nw182[(new BigNumber(8)).toNumber()] = _1173_cf41;
        _nw182[(new BigNumber(9)).toNumber()] = _1173_cf41;
        _nw182[(new BigNumber(10)).toNumber()] = !(_1118_v2).isEqualTo((_dafny.ZERO).minus(_1118_v2));
        _nw182[(new BigNumber(11)).toNumber()] = _1173_cf41;
        _nw182[(new BigNumber(12)).toNumber()] = (_this).f15;
        _nw182[(new BigNumber(13)).toNumber()] = _dafny.Seq.IsPrefixOf(_1116_v0, _module.__default.fm40(_module.__default.fm50((_this).f15, _1173_cf41, globalState), _1118_v2, _1173_cf41, (_this).f15, globalState));
        _nw182[(new BigNumber(14)).toNumber()] = _1173_cf41;
        _1178_v43 = _nw182;
        let _index164 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1178_v43).length));
        (_1178_v43)[_index164] = (_this).f15;
        let _1179_v44;
        let _init32 = ((_1180_v2) => function (_1181_i1) {
          return (_1181_i1).minus(_1180_v2);
        })(_1118_v2);
        let _nw183 = Array((new BigNumber(14)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw183.length); _i0_32++) {
          _nw183[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1179_v44 = _nw183;
        let _index165 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1178_v43).length));
        let _rhs120 = _dafny.Seq.Concat(_1175_v40, (_this).f14);
        let _rhs121 = (_this).f15;
        let _rhs122 = _1179_v44;
        let _rhs123 = _1118_v2;
        let _lhs73 = _1178_v43;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1178_v43).length));
        _1175_v40 = _rhs120;
        _lhs73[_lhs74] = _rhs121;
        _1179_v44 = _rhs122;
        _1118_v2 = _rhs123;
        let _1182_v45;
        _1182_v45 = _dafny.Set.fromElements(_1118_v2, _1118_v2);
        let _1183_v46;
        _1183_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1182_v45);
        let _1184_v47;
        _1184_v47 = new _dafny.CodePoint('l'.codePointAt(0));
        _1183_v46 = _module.__default.fm51(_1173_cf41, new _dafny.CodePoint('f'.codePointAt(0)), (_this).f15, _1184_v47, globalState);
        let _1185_v48;
        _1185_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v2,new BigNumber((_1116_v0).length));
        let _1186_v49;
        _1186_v49 = _module.D11.create_DC32((_dafny.ZERO).minus((((_1185_v48).contains(_1118_v2)) ? ((_1185_v48).get(_1118_v2)) : (_1118_v2))), _1118_v2);
        let _1187_v50;
        _1187_v50 = _dafny.Seq.of(_1175_v40);
        let _pat_let_tv17 = _1116_v0;
        let _1188_v51;
        let _nw184 = Array((new BigNumber(20)).toNumber());
        _nw184[(_dafny.ZERO).toNumber()] = _1186_v49;
        _nw184[(_dafny.ONE).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(2)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(3)).toNumber()] = _module.D11.create_DC32(new BigNumber(((_1187_v50)[_module.__default.safeIndex(_1118_v2, new BigNumber((_1187_v50).length))]).length), _1118_v2);
        _nw184[(new BigNumber(4)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(5)).toNumber()] = function (_pat_let16_0) {
          return function (_1189_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_1190_dt__update_hcf46_h0) {
                return _module.D11.create_DC32(_1190_dt__update_hcf46_h0, (_1189_dt__update__tmp_h0).dtor_cf47);
              }(_pat_let17_0);
            }(new BigNumber((_pat_let_tv17).length));
          }(_pat_let16_0);
        }(_1186_v49);
        _nw184[(new BigNumber(6)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(7)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(8)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(9)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(10)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(11)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(12)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(13)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(14)).toNumber()] = _module.D11.create_DC32(_1118_v2, new BigNumber((_1116_v0).length));
        _nw184[(new BigNumber(15)).toNumber()] = (((_this).f15) ? (_1186_v49) : (_1186_v49));
        _nw184[(new BigNumber(16)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(17)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(18)).toNumber()] = _1186_v49;
        _nw184[(new BigNumber(19)).toNumber()] = _1186_v49;
        _1188_v51 = _nw184;
        _1188_v51 = _1188_v51;
      } else if (_source18.is_DC23) {
        let _1191___mcc_h6 = (_source18).cf35;
        let _1192_cf35 = _1191___mcc_h6;
        let _1193_v52;
        _1193_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v2,(_this).f15);
        let _1194_v53;
        _1194_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D12.create_DC37(_1116_v0, (_this).f15)).dtor_cf59).length),(((_1193_v52).contains(new BigNumber(653))) ? ((_1193_v52).get(new BigNumber(653))) : ((_this).f15)));
        _1194_v53 = function () {
          let _coll48 = new _dafny.Map();
          for (const _compr_48 of _dafny.IntegerRange(new BigNumber(-743), new BigNumber(427))) {
            let _1195_v54 = _compr_48;
            if (((new BigNumber(-743)).isLessThanOrEqualTo(_1195_v54)) && ((_1195_v54).isLessThan(new BigNumber(427)))) {
              _coll48.push([(_1195_v54).multipliedBy(new BigNumber(567)),(_this).f15]);
            }
          }
          return _coll48;
        }();
        let _1196_v55;
        let _init33 = function (_1197_i2) {
          return (((_this).f15) ? ((_this).f15) : ((_this).f15));
        };
        let _nw185 = Array((new BigNumber(16)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw185.length); _i0_33++) {
          _nw185[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1196_v55 = _nw185;
        let _1198_v56;
        _1198_v56 = _dafny.MultiSet.fromElements((_this).f15);
        let _index166 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_1196_v55).length));
        (_1196_v55)[_index166] = (_1198_v56).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).f15, (_this).f15));
        let _1199_v57;
        _1199_v57 = _1198_v56;
        let _1200_v58;
        _1200_v58 = _module.D7.create_DC22((_this).f15);
        let _1201_v59;
        _1201_v59 = _dafny.Seq.of(_module.__default.fm43(_1199_v57, _1200_v58, globalState), (((_this).f15) ? ((_this).f14) : ((_this).f14)), (_this).f14, _dafny.Seq.Concat((_this).f14, (_this).f14));
        let _1202_v60;
        _1202_v60 = _dafny.Set.fromElements(_1118_v2, ((_this).f16)[_module.__default.safeIndex(new BigNumber(847), new BigNumber(((_this).f16).length))], _1118_v2, new BigNumber((_dafny.MultiSet.fromElements(_1193_v52)).cardinality()), (_dafny.ZERO).minus(_1118_v2));
        let _index167 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_1196_v55).length));
        let _rhs124 = _module.__default.fm2((_this).f15, globalState);
        let _rhs125 = _1201_v59;
        let _rhs126 = (_1202_v60).Union(_dafny.Set.fromElements(new BigNumber((_1116_v0).length), _1118_v2));
        let _lhs75 = _1196_v55;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_1196_v55).length));
        _lhs75[_lhs76] = _rhs124;
        _1201_v59 = _rhs125;
        _1202_v60 = _rhs126;
        let _1203_v61;
        _1203_v61 = _dafny.Seq.of(_1196_v55, _1196_v55, _1196_v55, _1196_v55);
        _1118_v2 = (_dafny.ZERO).minus((new BigNumber((_1203_v61).length)).multipliedBy(_1118_v2));
        _1193_v52 = _1194_v53;
      } else {
        let _1204___mcc_h7 = (_source18).cf42;
        let _1205_cf42 = _1204___mcc_h7;
        let _1206_v62;
        _1206_v62 = _dafny.MultiSet.fromElements((_this).f15);
        let _1207_v63;
        _1207_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber((_1116_v0).length));
        let _1208_v64;
        _1208_v64 = _module.D8.create_DC25((_this).f15, new BigNumber(698), new BigNumber((_1206_v62).cardinality()), (_1207_v63).update((_this).f15, new BigNumber(220)));
        _1208_v64 = _1208_v64;
        let _1209_v65;
        _1209_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v2,(_1118_v2).plus(_1118_v2));
        _1209_v65 = (_1209_v65).update((_this).fm5(!((_this).f15), globalState), _1118_v2);
        let _1210_v66;
        let _nw186 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
        _1210_v66 = _nw186;
        let _1211_v67;
        let _nw187 = Array((new BigNumber(17)).toNumber()).fill([]);
        _1211_v67 = _nw187;
        let _1212_v68;
        let _nw188 = Array((new BigNumber(9)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = false;
        _nw188[(_dafny.ONE).toNumber()] = (_this).f15;
        _nw188[(new BigNumber(2)).toNumber()] = (_this).f15;
        _nw188[(new BigNumber(3)).toNumber()] = (_this).f15;
        _nw188[(new BigNumber(4)).toNumber()] = (_this).f15;
        _nw188[(new BigNumber(5)).toNumber()] = true;
        _nw188[(new BigNumber(6)).toNumber()] = true;
        _nw188[(new BigNumber(7)).toNumber()] = false;
        _nw188[(new BigNumber(8)).toNumber()] = (_this).f15;
        _1212_v68 = _nw188;
        let _index168 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1211_v67).length));
        (_1211_v67)[_index168] = _1212_v68;
        let _1213_v69;
        let _nw189 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _1213_v69 = _nw189;
        let _index169 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_1213_v69).length));
        (_1213_v69)[_index169] = _1118_v2;
        let _1214_v70;
        _1214_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v2,(_this).f15);
        let _1215_v71;
        _1215_v71 = _dafny.Set.fromElements(_1214_v70);
        let _1216_v72;
        _1216_v72 = _module.D4.create_DC12((_this).f14);
        let _1217_v73;
        _1217_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v2,_1216_v72);
        let _1218_v74;
        _1218_v74 = _dafny.Set.fromElements(_1217_v73, _1217_v73, _1217_v73, _1217_v73, _dafny.Map.Empty.slice().updateUnsafe(_1118_v2,_1216_v72));
        let _1219_v75;
        _1219_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
        let _1220_v76;
        let _nw190 = new _module.C1();
        _nw190.__ctor((_this).f15, (_this).f15, ((_this).f13).update((_this).f15, _1219_v75), (_this).f14, _this.f12);
        _1220_v76 = _nw190;
        let _1221_v77;
        _1221_v77 = _module.D3.create_DC10((_this).f15, _1118_v2, _1220_v76);
        let _1222_v78;
        _1222_v78 = _dafny.Set.fromElements((_this).f15, (_this).f15, (_1221_v77).dtor_cf14);
        let _1223_v79;
        _1223_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1222_v78);
        let _1224_v80;
        _1224_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1222_v78,_1223_v79);
        let _1225_v81;
        _1225_v81 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1226_v82;
        _1226_v82 = _module.D4.create_DC14((_this).f15);
        let _1227_v83;
        _1227_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1225_v81,_module.__default.fm40(_1226_v82, new BigNumber(((_this).f16).length), (_this).f15, (_this).f15, globalState));
        let _index170 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1211_v67).length));
        let _index171 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_1213_v69).length));
        let _rhs127 = (((_dafny.Set.fromElements(_1217_v73)).IsDisjointFrom(_1218_v74)) ? (_1210_v66) : (_1210_v66));
        let _rhs128 = !((_this).f15);
        let _rhs129 = _1212_v68;
        let _rhs130 = new BigNumber(((((_1224_v80).contains(_1222_v78)) ? ((_1224_v80).get(_1222_v78)) : (_1223_v79))).length);
        let _rhs131 = _module.__default.fm52(_1227_v83, !((_this).f15), _module.__default.fm2((_this).f15, globalState), globalState);
        let _lhs77 = _1211_v67;
        let _lhs78 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1211_v67).length));
        let _lhs79 = _1213_v69;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_1213_v69).length));
        _1210_v66 = _rhs127;
        r1 = _rhs128;
        _lhs77[_lhs78] = _rhs129;
        _lhs79[_lhs80] = _rhs130;
        _1215_v71 = _rhs131;
        r1 = false;
      }
      _1118_v2 = _1118_v2;
      r1 = !((new BigNumber(((_this).f16).length)).isLessThan(new BigNumber(((_this).f16).length)));
      let _1228_v84;
      _1228_v84 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_1118_v2);
      let _1229_v85;
      _1229_v85 = _module.D8.create_DC23(_1228_v84);
      let _1230_v86;
      let _nw191 = new _module.C1();
      _nw191.__ctor((_this).f15, (_this).f15, (_this).f13, (_this).f14, _this.f12);
      _1230_v86 = _nw191;
      let _1231_v87;
      _1231_v87 = _module.D11.create_DC33((_this).f15, _1118_v2, false, _1230_v86, (_this).f15);
      let _pat_let_tv18 = _1116_v0;
      let _pat_let_tv19 = _1118_v2;
      let _pat_let_tv20 = _1118_v2;
      let _pat_let_tv21 = _1116_v0;
      let _pat_let_tv22 = _1118_v2;
      let _rhs132 = function (_source20) {
        if (_source20.is_DC24) {
          let _1232___mcc_h11 = (_source20).cf36;
          let _1233_cf36 = _1232___mcc_h11;
          return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv18,new BigNumber((_dafny.MultiSet.FromArray((_this).f14)).cardinality()))).length)).plus(_pat_let_tv19);
        } else if (_source20.is_DC25) {
          let _1234___mcc_h12 = (_source20).cf37;
          let _1235___mcc_h13 = (_source20).cf38;
          let _1236___mcc_h14 = (_source20).cf39;
          let _1237___mcc_h15 = (_source20).cf40;
          let _1238_cf40 = _1237___mcc_h15;
          let _1239_cf39 = _1236___mcc_h14;
          let _1240_cf38 = _1235___mcc_h13;
          let _1241_cf37 = _1234___mcc_h12;
          return _module.__default.safeDivisionInt(_pat_let_tv20, _1240_cf38);
        } else if (_source20.is_DC26) {
          let _1242___mcc_h16 = (_source20).cf41;
          let _1243_cf41 = _1242___mcc_h16;
          return new BigNumber(((_this).f14).length);
        } else if (_source20.is_DC23) {
          let _1244___mcc_h17 = (_source20).cf35;
          let _1245_cf35 = _1244___mcc_h17;
          return new BigNumber((_pat_let_tv21).length);
        } else {
          let _1246___mcc_h18 = (_source20).cf42;
          let _1247_cf42 = _1246___mcc_h18;
          return _pat_let_tv22;
        }
      }(_1229_v85);
      let _rhs133 = (((_this).f15) ? (_1118_v2) : ((_1231_v87).dtor_cf49));
      _1118_v2 = _rhs132;
      _1118_v2 = _rhs133;
      if (_1230_v86.f18) {
        _1118_v2 = new BigNumber(((_this).f14).length);
        let _1248_v88;
        let _nw192 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _1248_v88 = _nw192;
        let _index172 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1248_v88).length));
        (_1248_v88)[_index172] = _module.__default.safeDivisionInt(_1118_v2, _1118_v2);
        let _index173 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1248_v88).length));
        (_1248_v88)[_index173] = ((_module.__default.fm2((_1230_v86).f17, globalState)) ? (_1118_v2) : (_1118_v2));
        (globalState).f4 = _1230_v86.f18;
        let _1249_v89;
        _1249_v89 = _dafny.Map.Empty.slice().updateUnsafe((_1248_v88)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1248_v88).length))],new BigNumber(49));
        _1249_v89 = (_1249_v89).update(new BigNumber(-248), (_1248_v88)[_module.__default.safeIndex(new BigNumber(739), new BigNumber((_1248_v88).length))]);
        (_1230_v86).f18 = false;
      } else {
        (_1230_v86).f18 = _1230_v86.f18;
        let _1250_v90;
        _1250_v90 = _dafny.Seq.of(_dafny.Seq.update((_this).f14, _module.__default.safeIndex(_1118_v2, new BigNumber(((_this).f14).length)), true), _dafny.Seq.of((_this).f15));
        let _1251_v91;
        _1251_v91 = _module.D2.create_DC4(_1250_v90);
        _1251_v91 = ((_1230_v86.f18) ? (_1251_v91) : (_module.__default.fm53(new BigNumber(769), (_1230_v86).f17, globalState)));
        let _1252_v92;
        _1252_v92 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1253_v93;
        _1253_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1230_v86.f18,_1252_v92);
        _1253_v93 = ((_1253_v93).Merge(_1253_v93)).update(_1230_v86.f18, _1252_v92);
        let _1254_v94;
        _1254_v94 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1118_v2), new BigNumber(((_this).f16).length));
        let _1255_v95;
        _1255_v95 = _dafny.Seq.of(_1116_v0);
        _1118_v2 = (_1118_v2).minus(new BigNumber((_dafny.Seq.Concat((_1250_v90)[_module.__default.safeIndex(new BigNumber(((_this).fm10((_1230_v86).f17, _1254_v94, (_1230_v86).fm6((_1255_v95)[_module.__default.safeIndex(_1118_v2, new BigNumber((_1255_v95).length))], globalState), _1118_v2, globalState)).length), new BigNumber((_1250_v90).length))], (_this).f14)).length));
        let _1256_v96;
        let _nw193 = Array((new BigNumber(18)).toNumber()).fill([]);
        _1256_v96 = _nw193;
        let _1257_v97;
        let _nw194 = Array((new BigNumber(11)).toNumber());
        _nw194[(_dafny.ZERO).toNumber()] = (_1255_v95)[_module.__default.safeIndex(_1118_v2, new BigNumber((_1255_v95).length))];
        _nw194[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("nddakcmt");
        _nw194[(new BigNumber(2)).toNumber()] = _1116_v0;
        _nw194[(new BigNumber(3)).toNumber()] = _1116_v0;
        _nw194[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("bkkowjg");
        _nw194[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("t");
        _nw194[(new BigNumber(6)).toNumber()] = _1116_v0;
        _nw194[(new BigNumber(7)).toNumber()] = _1116_v0;
        _nw194[(new BigNumber(8)).toNumber()] = _1116_v0;
        _nw194[(new BigNumber(9)).toNumber()] = _1116_v0;
        _nw194[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("kw");
        _1257_v97 = _nw194;
        let _1258_v98;
        let _init34 = ((_1259_v86) => function (_1260_i3) {
          return (_1259_v86).f17;
        })(_1230_v86);
        let _nw195 = Array((new BigNumber(17)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw195.length); _i0_34++) {
          _nw195[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1258_v98 = _nw195;
        let _1261_v99;
        _1261_v99 = _dafny.Set.fromElements(_1258_v98, _1258_v98);
        let _1262_v100;
        _1262_v100 = _module.D12.create_DC36(_1257_v97, _1118_v2, _module.__default.fm47(globalState), _1261_v99);
        let _1263_v101;
        _1263_v101 = _1256_v96;
        let _rhs134 = _1230_v86.f18;
        let _rhs135 = (_1263_v101);
        let _rhs136 = _1230_v86.f18;
        let _rhs137 = (_dafny.ZERO).minus((_this).fm8(!((_this).f15), (_this).f15, globalState));
        let _rhs138 = _module.D12.create_DC36(_1257_v97, _1118_v2, _dafny.Map.Empty.slice().updateUnsafe(_1116_v0,(_this).f16), _dafny.Set.fromElements(_1258_v98));
        let _lhs81 = globalState;
        let _lhs82 = _1230_v86;
        _lhs81.f4 = _rhs134;
        _1256_v96 = _rhs135;
        _lhs82.f18 = _rhs136;
        _1118_v2 = _rhs137;
        _1262_v100 = _rhs138;
      }
      let _1264_v102;
      _1264_v102 = new _dafny.CodePoint('x'.codePointAt(0));
      r0 = _1264_v102;
      r1 = (_1230_v86).f17;
      return [r0, r1];
    }
    m12(p0, globalState) {
      let _this = this;
      let _1265_v0;
      _1265_v0 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
      let _1266_v1;
      _1266_v1 = _module.D3.create_DC8(_1265_v0);
      let _1267_v2;
      _1267_v2 = new BigNumber(-280);
      let _1268_v3;
      let _nw196 = Array((new BigNumber(14)).toNumber());
      _nw196[(_dafny.ZERO).toNumber()] = _module.__default.fm54((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-545)), function (_1269_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length)), globalState);
      _nw196[(_dafny.ONE).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(2)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(3)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(4)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(5)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(6)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(7)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(8)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(9)).toNumber()] = ((((_this).f14)[_module.__default.safeIndex(_1267_v2, new BigNumber(((_this).f14).length))]) ? (_module.D3.create_DC8(_1265_v0)) : (_1266_v1));
      _nw196[(new BigNumber(10)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(11)).toNumber()] = _1266_v1;
      _nw196[(new BigNumber(12)).toNumber()] = _module.__default.fm54(_1267_v2, globalState);
      _nw196[(new BigNumber(13)).toNumber()] = _1266_v1;
      _1268_v3 = _nw196;
      let _index174 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1268_v3).length));
      (_1268_v3)[_index174] = _1266_v1;
      let _index175 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1268_v3).length));
      let _rhs139 = _module.__default.fm54(_1267_v2, globalState);
      let _rhs140 = _1267_v2;
      let _rhs141 = p0;
      let _lhs83 = _1268_v3;
      let _lhs84 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1268_v3).length));
      let _lhs85 = globalState;
      _lhs83[_lhs84] = _rhs139;
      _1267_v2 = _rhs140;
      _lhs85.f4 = _rhs141;
      let _1270_v4;
      _1270_v4 = _dafny.MultiSet.fromElements(!(false), (_this).f15, p0, (_this).f15, true);
      let _1271_v5;
      _1271_v5 = _dafny.MultiSet.fromElements((_1270_v4).IsDisjointFrom(_1270_v4));
      _1267_v2 = (((_1271_v5).contains(!(_1267_v2).isEqualTo(_1267_v2))) ? ((_1271_v5).get(!(_1267_v2).isEqualTo(_1267_v2))) : (new BigNumber(-22)));
      let _1272_v6;
      _1272_v6 = new _dafny.CodePoint('b'.codePointAt(0));
      _1272_v6 = _1272_v6;
      let _1273_v7;
      _1273_v7 = _dafny.Set.fromElements(_1267_v2);
      let _1274_v8;
      _1274_v8 = _module.D6.create_DC19(_module.D6.create_DC17(_1273_v7));
      let _1275_v9;
      _1275_v9 = _module.D6.create_DC19(_module.D6.create_DC19(_1274_v8));
      _1275_v9 = _1275_v9;
      let _1276_v10;
      let _init35 = ((_1277_v2) => function (_1278_i1) {
        return (_1278_i1).multipliedBy(_1277_v2);
      })(_1267_v2);
      let _nw197 = Array((new BigNumber(23)).toNumber());
      for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw197.length); _i0_35++) {
        _nw197[_i0_35] = _init35(new BigNumber(_i0_35));
      }
      _1276_v10 = _nw197;
      let _1279_v11;
      _1279_v11 = _dafny.Seq.of(_1276_v10, _1276_v10, _1276_v10);
      _1267_v2 = new BigNumber((_1279_v11).length);
      let _1280_v12;
      _1280_v12 = _module.D4.create_DC13(p0, p0, _1267_v2, _1267_v2, _module.__default.fm2((_this).f15, globalState));
      let _1281_v13;
      _1281_v13 = _module.D7.create_DC22(p0);
      let _1282_v14;
      _1282_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1280_v12,(_1281_v13).dtor_cf34);
      let _1283_v15;
      _1283_v15 = _dafny.Seq.UnicodeFromString("n");
      let _index176 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_1276_v10).length));
      (_1276_v10)[_index176] = new BigNumber((_1283_v15).length);
      let _1284_v16;
      _1284_v16 = _dafny.Seq.of(_1282_v14, (_1282_v14).update(_module.__default.fm55(globalState), (_this).f15));
      let _1285_v18;
      _1285_v18 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("hxxlj")).length));
      let _pat_let_tv23 = _1285_v18;
      let _1286_v19;
      _1286_v19 = _dafny.Set.fromElements(_module.__default.fm55(globalState), function (_pat_let18_0) {
        return function (_1287_dt__update__tmp_h0) {
          return function (_pat_let19_0) {
            return function (_1288_dt__update_hcf22_h0) {
              return function (_pat_let20_0) {
                return function (_1289_dt__update_hcf19_h0) {
                  return _module.D4.create_DC13(_1289_dt__update_hcf19_h0, (_1287_dt__update__tmp_h0).dtor_cf20, (_1287_dt__update__tmp_h0).dtor_cf21, _1288_dt__update_hcf22_h0, (_1287_dt__update__tmp_h0).dtor_cf23);
                }(_pat_let20_0);
              }((_this).f15);
            }(_pat_let19_0);
          }(new BigNumber((_pat_let_tv23).cardinality()));
        }(_pat_let18_0);
      }(_module.D4.create_DC13(p0, (_this).f15, _1267_v2, _1267_v2, p0)));
      let _1290_v20;
      _1290_v20 = _dafny.Seq.of(_1282_v14, _1282_v14, (_1284_v16)[_module.__default.safeIndex(new BigNumber((_module.__default.fm56(_1272_v6, globalState)).cardinality()), new BigNumber((_1284_v16).length))], function () {
        let _coll49 = new _dafny.Map();
        for (const _compr_49 of (_1286_v19).Elements) {
          let _1291_v17 = _compr_49;
          if ((_1286_v19).contains(_1291_v17)) {
            _coll49.push([_1291_v17,(_this).f15]);
          }
        }
        return _coll49;
      }(), _1282_v14);
      let _1292_v21;
      _1292_v21 = _dafny.Set.fromElements((_this).f15, p0, false);
      let _index177 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_1276_v10).length));
      let _rhs142 = (_1290_v20)[_module.__default.safeIndex((_1267_v2).multipliedBy(_1267_v2), new BigNumber((_1290_v20).length))];
      let _rhs143 = new BigNumber((_1292_v21).length);
      let _rhs144 = new BigNumber((_dafny.Set.fromElements((_1267_v2).isLessThanOrEqualTo(_1267_v2), _module.__default.fm2(p0, globalState), p0, !(p0))).length);
      let _lhs86 = _1276_v10;
      let _lhs87 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_1276_v10).length));
      _1282_v14 = _rhs142;
      _lhs86[_lhs87] = _rhs143;
      _1267_v2 = _rhs144;
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f12 = [];
      this._f15 = false;
      this._f13 = _dafny.Map.Empty;
      this._f14 = _dafny.Seq.of();
      this._f20 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T2, _module.T1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    __ctor(f20, f12, f15, f13, f14) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f12 = f12;
      (_this)._f15 = f15;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.UnicodeFromString("jpicrv")).length);
    };
    fm7(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber(46), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0))))).cardinality())), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15)).length)).minus(new BigNumber(((_this).f14).length))));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      let _source21 = _module.D7.create_DC22((_this).f15);
      if (_source21.is_DC21) {
        let _1293___mcc_h0 = (_source21).cf31;
        let _1294___mcc_h1 = (_source21).cf32;
        let _1295___mcc_h2 = (_source21).cf33;
        let _1296_cf33 = _1295___mcc_h2;
        let _1297_cf32 = _1294___mcc_h1;
        let _1298_cf31 = _1293___mcc_h0;
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),(_dafny.ZERO).minus(new BigNumber(((_this).f14).length)))).length);
      } else if (_source21.is_DC22) {
        let _1299___mcc_h3 = (_source21).cf34;
        let _1300_cf34 = _1299___mcc_h3;
        return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f15),(_this).f15)).length)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))).length)));
      } else {
        let _1301___mcc_h4 = (_source21).cf30;
        let _1302_cf30 = _1301___mcc_h4;
        return (_dafny.ZERO).minus((new BigNumber((_dafny.Seq.UnicodeFromString("pp")).length)).multipliedBy(new BigNumber(315)));
      }
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(547)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(394), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f20, (_this).f15)).length),(_this).f15)).length),new _dafny.CodePoint('w'.codePointAt(0)))).length), new BigNumber((function () {
        let _coll50 = new _dafny.Set();
        for (const _compr_50 of _dafny.IntegerRange(new BigNumber(117), new BigNumber(419))) {
          let _1303_v0 = _compr_50;
          if (((new BigNumber(117)).isLessThanOrEqualTo(_1303_v0)) && ((_1303_v0).isLessThan(new BigNumber(419)))) {
            _coll50.add((_1303_v0).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("yfbwyy")).length), new BigNumber(226))).length)));
          }
        }
        return _coll50;
      }()).length)), _dafny.Seq.of(new BigNumber(-612), new BigNumber(737), new BigNumber(-32))));
    };
    fm33(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))).Difference(_dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))))).Difference((_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).Union(_dafny.MultiSet.fromElements((new _dafny.CodePoint('w'.codePointAt(0))))));
    };
    fm34(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gm"), _dafny.Seq.UnicodeFromString("mcwxe"))).length)).plus(_module.__default.safeModuloInt(new BigNumber(-276), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-153)), function (_1304_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length)));
    };
    m3(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let r3 = false;
      let _1305_v0;
      _1305_v0 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f20),_this.f12);
      let _1306_v1;
      let _nw198 = new _module.C1();
      _nw198.__ctor((_this).f20, (_this).f20, (_this).f13, (_this).f14, (((_1305_v0).contains((_this).f20)) ? ((_1305_v0).get((_this).f20)) : (_this.f12)));
      _1306_v1 = _nw198;
      let _1307_v2;
      _1307_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1306_v1,((_1306_v1.f18) ? (_module.__default.fm2((_1306_v1).f17, globalState)) : ((_1306_v1).f17)));
      _1307_v2 = (_1307_v2).update(_1306_v1, (_1306_v1).f17);
      let _1308_v3;
      _1308_v3 = new BigNumber(369);
      let _hi6 = (_this).fm34((_1306_v1).f17, _1308_v3, (_this).f20, (_1306_v1).f17, globalState);
      for (let _1309_i0 = _1308_v3; _1309_i0.isLessThan(_hi6); _1309_i0 = _1309_i0.plus(_dafny.ONE)) {
        let _1310_v4;
        _1310_v4 = _module.D4.create_DC12((_this).f14);
        let _1311_v5;
        _1311_v5 = _dafny.Seq.of((_this).f14, (_this).f14, (_this).f14);
        let _1312_v6;
        let _nw199 = Array((new BigNumber(20)).toNumber());
        _nw199[(_dafny.ZERO).toNumber()] = (_this).f14;
        _nw199[(_dafny.ONE).toNumber()] = (_1310_v4).dtor_cf18;
        _nw199[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update((_this).f14, _module.__default.safeIndex(_1309_i0, new BigNumber(((_this).f14).length)), !(_1306_v1.f18)), _dafny.Seq.of(false, (_this).f20));
        _nw199[(new BigNumber(4)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(5)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(6)).toNumber()] = (_1311_v5)[_module.__default.safeIndex(new BigNumber(-846), new BigNumber((_1311_v5).length))];
        _nw199[(new BigNumber(7)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(8)).toNumber()] = (_1311_v5)[_module.__default.safeIndex(new BigNumber(833), new BigNumber((_1311_v5).length))];
        _nw199[(new BigNumber(9)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat((_this).f14, (_this).f14);
        _nw199[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat((_this).f14, _dafny.Seq.of(_1306_v1.f18, _1306_v1.f18, (_1306_v1).f17));
        _nw199[(new BigNumber(12)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(13)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(14)).toNumber()] = _dafny.Seq.update((_this).f14, _module.__default.safeIndex(_1309_i0, new BigNumber(((_this).f14).length)), (_1306_v1).f17);
        _nw199[(new BigNumber(15)).toNumber()] = _dafny.Seq.of((_1306_v1).f17);
        _nw199[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat((_this).f14, (_this).f14);
        _nw199[(new BigNumber(17)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(18)).toNumber()] = (_this).f14;
        _nw199[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_1306_v1.f18);
        _1312_v6 = _nw199;
        let _1313_v7;
        _1313_v7 = _dafny.Seq.of(_1309_i0, _1308_v3);
        let _1314_v8;
        let _nw200 = new _module.C1();
        _nw200.__ctor((_this).f15, _module.__default.fm2(_1306_v1.f18, globalState), (_this).f13, (_this).f14, _this.f12);
        _1314_v8 = _nw200;
        let _1315_v9;
        _1315_v9 = _dafny.Set.fromElements(_module.D3.create_DC10((_1306_v1).fm17(_1313_v7, new BigNumber(515), globalState), _1308_v3, _1314_v8));
        let _index178 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_1312_v6).length));
        (_1312_v6)[_index178] = _dafny.Seq.Concat(_dafny.Seq.of(_1306_v1.f18), (_1311_v5)[_module.__default.safeIndex(new BigNumber((_1315_v9).length), new BigNumber((_1311_v5).length))]);
        let _index179 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_1312_v6).length));
        (_1312_v6)[_index179] = (_1314_v8).f14;
        let _1316_v10;
        _1316_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1306_v1).f17,new BigNumber(91));
        let _1317_v11;
        _1317_v11 = _module.D8.create_DC25((_this).f15, _1308_v3, _1308_v3, _1316_v10);
        let _1318_v12;
        _1318_v12 = _dafny.Seq.UnicodeFromString("jas");
        let _1319_v13;
        _1319_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_dafny.Seq.Create(_module.__default.abs(new BigNumber(505)), function (_1320_i2) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }));
        let _1321_v14;
        _1321_v14 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1322_v15;
        _1322_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1309_i0,_1318_v12);
        let _1323_v16;
        let _nw201 = Array((new BigNumber(14)).toNumber());
        _nw201[(_dafny.ZERO).toNumber()] = _module.__default.fm35(_1308_v3, _1317_v11, globalState);
        _nw201[(_dafny.ONE).toNumber()] = _1318_v12;
        _nw201[(new BigNumber(2)).toNumber()] = _1318_v12;
        _nw201[(new BigNumber(3)).toNumber()] = _1318_v12;
        _nw201[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(809)), function (_1324_i1) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        });
        _nw201[(new BigNumber(5)).toNumber()] = _1318_v12;
        _nw201[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("cpq");
        _nw201[(new BigNumber(7)).toNumber()] = _1318_v12;
        _nw201[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1318_v12, _1318_v12);
        _nw201[(new BigNumber(9)).toNumber()] = (((_1319_v13).contains(false)) ? ((_1319_v13).get(false)) : (_1318_v12));
        _nw201[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_1318_v12, _module.__default.safeIndex(_1308_v3, new BigNumber((_1318_v12).length)), _1321_v14);
        _nw201[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jpvq"), (((_1322_v15).contains((_dafny.ZERO).minus(_1308_v3))) ? ((_1322_v15).get((_dafny.ZERO).minus(_1308_v3))) : (_dafny.Seq.UnicodeFromString("ytiar"))));
        _nw201[(new BigNumber(12)).toNumber()] = _1318_v12;
        _nw201[(new BigNumber(13)).toNumber()] = _1318_v12;
        _1323_v16 = _nw201;
        let _index180 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1323_v16).length));
        (_1323_v16)[_index180] = _1318_v12;
        let _index181 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1323_v16).length));
        let _rhs145 = _dafny.Seq.Concat(_1318_v12, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iw"), _1318_v12));
        let _rhs146 = _1306_v1.f18;
        let _lhs88 = _1323_v16;
        let _lhs89 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_1323_v16).length));
        let _lhs90 = globalState;
        _lhs88[_lhs89] = _rhs145;
        _lhs90.f4 = _rhs146;
        r1 = _1308_v3;
        let _1325_v17;
        let _nw202 = new _module.C1();
        _nw202.__ctor(_1306_v1.f18, _1306_v1.f18, (_this).f13, (_this).f14, _this.f12);
        _1325_v17 = _nw202;
      }
      let _1326_v18;
      _1326_v18 = _dafny.Seq.of(_1308_v3);
      let _1327_v19;
      _1327_v19 = _dafny.Seq.of(_1326_v18);
      let _1328_v20;
      _1328_v20 = _dafny.Set.fromElements((_1327_v19)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("cekuw")).length), new BigNumber((_1327_v19).length))]);
      let _1329_v21;
      _1329_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_dafny.Seq.update(_1326_v18, _module.__default.safeIndex((_dafny.ZERO).minus(_1308_v3), new BigNumber((_1326_v18).length)), new BigNumber((_1328_v20).length)));
      let _hi7 = (new BigNumber(((_this).f14).length)).plus(_1308_v3);
      for (let _1330_i3 = new BigNumber(((((_1329_v21).contains((_1306_v1).f17)) ? ((_1329_v21).get((_1306_v1).f17)) : (_1326_v18))).length); _1330_i3.isLessThan(_hi7); _1330_i3 = _1330_i3.plus(_dafny.ONE)) {
        let _1331_v22;
        _1331_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1308_v3,(_1306_v1).fm17(_dafny.Seq.update(_dafny.Seq.of(_1330_i3), _module.__default.safeIndex(_1330_i3, new BigNumber((_dafny.Seq.of(_1330_i3)).length)), (_dafny.ZERO).minus(_1308_v3)), _1330_i3, globalState));
        let _1332_v23;
        _1332_v23 = _dafny.Seq.UnicodeFromString("v");
        _1331_v22 = (_1331_v22).update(new BigNumber((_1331_v22).length), _dafny.Seq.IsProperPrefixOf(_1332_v23, _1332_v23));
        (globalState).f4 = (_this).f20;
        let _1333_v24;
        let _init36 = ((_1334_v18, _1335_v1, _1336_v21) => function (_1337_i4) {
          return (_1337_i4).plus((_1334_v18)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray((((_1336_v21).contains(_1335_v1.f18)) ? ((_1336_v21).get(_1335_v1.f18)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), ((_1338_v18) => function (_1339_i5) {
            return new BigNumber((_1338_v18).length);
          })(_1334_v18)))))).cardinality()), new BigNumber((_1334_v18).length))]);
        })(_1326_v18, _1306_v1, _1329_v21);
        let _nw203 = Array((new BigNumber(23)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw203.length); _i0_36++) {
          _nw203[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1333_v24 = _nw203;
        let _index182 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1333_v24).length));
        (_1333_v24)[_index182] = (_1306_v1).fm5(true, globalState);
        let _index183 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1333_v24).length));
        let _rhs147 = (_1308_v3).plus(_1330_i3);
        let _rhs148 = (_1308_v3).isEqualTo(_1308_v3);
        let _rhs149 = _1330_i3;
        let _lhs91 = _1333_v24;
        let _lhs92 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_1333_v24).length));
        r1 = _rhs147;
        r0 = _rhs148;
        _lhs91[_lhs92] = _rhs149;
        r0 = !((_1306_v1).f17) || (!(!((_1306_v1).f17)));
      }
      let _1340_v25;
      _1340_v25 = _dafny.Seq.UnicodeFromString("k");
      let _1341_v26;
      _1341_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1308_v3,_1340_v25);
      let _1342_v27;
      _1342_v27 = _dafny.Seq.of((((_1341_v26).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f14)).length))) ? ((_1341_v26).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f14)).length))) : (_1340_v25)), _1340_v25);
      let _1343_v28;
      _1343_v28 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).fm7((_dafny.ZERO).minus(new BigNumber((_1326_v18).length)), globalState)), (_1308_v3).plus(_1308_v3), new BigNumber((_1342_v27).length));
      let _1344_v29;
      _1344_v29 = _dafny.MultiSet.fromElements(new BigNumber((_1328_v20).length));
      let _1345_v30;
      _1345_v30 = new _dafny.CodePoint('b'.codePointAt(0));
      let _rhs150 = _1308_v3;
      let _rhs151 = _dafny.Set.fromElements(new BigNumber((_1344_v29).cardinality()), _module.__default.safeDivisionInt(_1308_v3, _1308_v3), _1308_v3, _1308_v3);
      let _rhs152 = (new BigNumber((_dafny.Seq.of(_1345_v30, new _dafny.CodePoint('c'.codePointAt(0)), _1345_v30, _1345_v30)).length)).minus(_1308_v3);
      let _rhs153 = _1340_v25;
      _1308_v3 = _rhs150;
      _1343_v28 = _rhs151;
      r1 = _rhs152;
      _1340_v25 = _rhs153;
      let _1346_v31;
      _1346_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(436),(_1306_v1).fm17(_1326_v18, new BigNumber((_1343_v28).length), globalState));
      let _1347_v32;
      _1347_v32 = _module.D7.create_DC21(_1346_v31, false, ((_this).f14)[_module.__default.safeIndex(_1308_v3, new BigNumber(((_this).f14).length))]);
      let _1348_v33;
      _1348_v33 = _dafny.MultiSet.fromElements(_1346_v31, (_1346_v31).Merge((_1347_v32).dtor_cf31), (_1346_v31).update(_1308_v3, !(false)), _1346_v31);
      _1348_v33 = _1348_v33;
      let _1349_v34;
      _1349_v34 = _dafny.Set.fromElements((_1306_v1).f17);
      let _rhs154 = (_1308_v3).isLessThanOrEqualTo(_1308_v3);
      let _rhs155 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1349_v34).Intersect(_dafny.Set.fromElements(false)),_1308_v3)).length);
      r0 = _rhs154;
      r1 = _rhs155;
      r0 = (_1344_v29).IsSubsetOf(_1344_v29);
      let _1350_v35;
      _1350_v35 = _module.D4.create_DC13(!((_this).f15), _1306_v1.f18, _1308_v3, _1308_v3, (_this).f20);
      let _pat_let_tv24 = _1346_v31;
      let _pat_let_tv25 = _1306_v1;
      let _pat_let_tv26 = _1326_v18;
      let _pat_let_tv27 = _1308_v3;
      r1 = function (_source22) {
        if (_source22.is_DC13) {
          let _1351___mcc_h0 = (_source22).cf19;
          let _1352___mcc_h1 = (_source22).cf20;
          let _1353___mcc_h2 = (_source22).cf21;
          let _1354___mcc_h3 = (_source22).cf22;
          let _1355___mcc_h4 = (_source22).cf23;
          let _1356_cf23 = _1355___mcc_h4;
          let _1357_cf22 = _1354___mcc_h3;
          let _1358_cf21 = _1353___mcc_h2;
          let _1359_cf20 = _1352___mcc_h1;
          let _1360_cf19 = _1351___mcc_h0;
          return new BigNumber((_pat_let_tv24).length);
        } else if (_source22.is_DC14) {
          let _1361___mcc_h5 = (_source22).cf24;
          let _1362_cf24 = _1361___mcc_h5;
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_pat_let_tv25).f17)).length);
        } else {
          let _1363___mcc_h6 = (_source22).cf18;
          let _1364_cf18 = _1363___mcc_h6;
          return (new BigNumber((_pat_let_tv26).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(665)), ((_1365_v3) => function (_1366_i6) {
            return _1365_v3;
          })(_pat_let_tv27))).length));
        }
      }(_1350_v35);
      let _1367_v37;
      _1367_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1306_v1.f18,_1308_v3);
      let _1368_v38;
      _1368_v38 = _dafny.Set.fromElements(_1340_v25, _1340_v25, _1340_v25, _dafny.Seq.Create(_module.__default.abs(new BigNumber(58)), ((_1369_v30) => function (_1370_i7) {
        return _1369_v30;
      })(_1345_v30)), _module.__default.fm35(_1308_v3, _module.D8.create_DC25((_this).f15, _1308_v3, _1308_v3, _1367_v37), globalState));
      r2 = _module.__default.fm36(_1308_v3, _1308_v3, (((_1306_v1).f17) ? (function () {
        let _coll51 = new _dafny.Set();
        for (const _compr_51 of (_1342_v27).Elements) {
          let _1371_v36 = _compr_51;
          if (_dafny.Seq.contains(_1342_v27, _1371_v36)) {
            _coll51.add(_1371_v36);
          }
        }
        return _coll51;
      }()) : (_1368_v38)), globalState);
      r3 = !((_this).f15) || (!((_1306_v1).f17));
      return [r0, r1, r2, r3];
    }
    m9(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.Map.Empty;
      (_this).m10(globalState);
      let _1372_v0;
      _1372_v0 = new BigNumber(-546);
      let _1373_v1;
      _1373_v1 = _dafny.Set.fromElements(_1372_v0);
      let _1374_i0;
      _1374_i0 = _dafny.ZERO;
      L10: {
        while ((_1373_v1).IsSubsetOf(_1373_v1)) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1374_i0)) {
              break L10;
            }
            _1374_i0 = (_1374_i0).plus(_dafny.ONE);
            let _1375_v2;
            _1375_v2 = _dafny.Seq.UnicodeFromString("vakf");
            let _1376_v3;
            _1376_v3 = new _dafny.CodePoint('a'.codePointAt(0));
            let _rhs156 = _1375_v2;
            let _rhs157 = (_1372_v0).multipliedBy(_1372_v0);
            let _rhs158 = (_this).f20;
            let _rhs159 = _1376_v3;
            let _lhs93 = globalState;
            _1375_v2 = _rhs156;
            r0 = _rhs157;
            _lhs93.f4 = _rhs158;
            _1376_v3 = _rhs159;
            let _1377_v4;
            _1377_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(897),p1);
            let _1378_v5;
            let _nw204 = Array((new BigNumber(17)).toNumber());
            _nw204[(_dafny.ZERO).toNumber()] = (_this).f20;
            _nw204[(_dafny.ONE).toNumber()] = p0;
            _nw204[(new BigNumber(2)).toNumber()] = p1;
            _nw204[(new BigNumber(3)).toNumber()] = !(p1);
            _nw204[(new BigNumber(4)).toNumber()] = p0;
            _nw204[(new BigNumber(5)).toNumber()] = true;
            _nw204[(new BigNumber(6)).toNumber()] = (_this).f15;
            _nw204[(new BigNumber(7)).toNumber()] = p1;
            _nw204[(new BigNumber(8)).toNumber()] = (_this).f15;
            _nw204[(new BigNumber(9)).toNumber()] = !(p0);
            _nw204[(new BigNumber(10)).toNumber()] = (_this).f20;
            _nw204[(new BigNumber(11)).toNumber()] = (_this).f15;
            _nw204[(new BigNumber(12)).toNumber()] = p0;
            _nw204[(new BigNumber(13)).toNumber()] = p1;
            _nw204[(new BigNumber(14)).toNumber()] = p0;
            _nw204[(new BigNumber(15)).toNumber()] = _module.__default.fm2(p1, globalState);
            _nw204[(new BigNumber(16)).toNumber()] = (((_1377_v4).contains(_1372_v0)) ? ((_1377_v4).get(_1372_v0)) : ((_this).f20));
            _1378_v5 = _nw204;
            let _1379_v6;
            _1379_v6 = _module.D2.create_DC6(_1378_v5, _1378_v5);
            let _1380_v7;
            _1380_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1378_v5);
            let _pat_let_tv28 = _1378_v5;
            let _1381_v8;
            let _nw205 = Array((new BigNumber(29)).toNumber());
            _nw205[(_dafny.ZERO).toNumber()] = _1378_v5;
            _nw205[(_dafny.ONE).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(2)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(3)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(4)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(5)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(6)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(7)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(8)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(9)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(10)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(11)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(12)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(13)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(14)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(15)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(16)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(17)).toNumber()] = (((_this).f20) ? (_1378_v5) : (_1378_v5));
            _nw205[(new BigNumber(18)).toNumber()] = (_1379_v6).dtor_cf9;
            _nw205[(new BigNumber(19)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(20)).toNumber()] = (((_1380_v7).contains(true)) ? ((_1380_v7).get(true)) : (_1378_v5));
            _nw205[(new BigNumber(21)).toNumber()] = (function (_pat_let21_0) {
              return function (_1382_dt__update__tmp_h0) {
                return function (_pat_let22_0) {
                  return function (_1383_dt__update_hcf8_h0) {
                    return _module.D2.create_DC6(_1383_dt__update_hcf8_h0, (_1382_dt__update__tmp_h0).dtor_cf9);
                  }(_pat_let22_0);
                }(_pat_let_tv28);
              }(_pat_let21_0);
            }(_module.D2.create_DC6(_1378_v5, _1378_v5))).dtor_cf9;
            _nw205[(new BigNumber(22)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(23)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(24)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(25)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(26)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(27)).toNumber()] = _1378_v5;
            _nw205[(new BigNumber(28)).toNumber()] = _1378_v5;
            _1381_v8 = _nw205;
            let _index184 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1378_v5).length));
            (_1378_v5)[_index184] = (_this).f15;
            let _index185 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1378_v5).length));
            let _rhs160 = _1381_v8;
            let _rhs161 = true;
            let _rhs162 = !((_dafny.ZERO).minus(new BigNumber(-910))).isEqualTo(_1372_v0);
            let _rhs163 = _1372_v0;
            let _rhs164 = _1372_v0;
            let _lhs94 = _1378_v5;
            let _lhs95 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_1378_v5).length));
            let _lhs96 = globalState;
            _1381_v8 = _rhs160;
            _lhs94[_lhs95] = _rhs161;
            _lhs96.f4 = _rhs162;
            _1372_v0 = _rhs163;
            r0 = _rhs164;
            r0 = (_1372_v0).minus(_1372_v0);
            let _1384_v9;
            _1384_v9 = _dafny.MultiSet.fromElements((_1378_v5)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_1378_v5).length))], (_this).f15, p0, !(p0));
            r1 = (_1372_v0).minus(new BigNumber((_1384_v9).cardinality()));
          }
        }
      }
      let _1385_v10;
      _1385_v10 = _dafny.MultiSet.fromElements((_this).f20, p0, p1);
      let _hi8 = new BigNumber((_1385_v10).cardinality());
      for (let _1386_i1 = _1372_v0; _1386_i1.isLessThan(_hi8); _1386_i1 = _1386_i1.plus(_dafny.ONE)) {
        let _1387_v11;
        _1387_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f20);
        if ((((_1387_v11).contains(p2)) ? ((_1387_v11).get(p2)) : (((_this).f15) && (p0)))) {
          let _1388_v12;
          _1388_v12 = _dafny.Map.Empty.slice().updateUnsafe(((!((_this).f20)) ? ((_this).f15) : (p0)),_dafny.Seq.UnicodeFromString("omhw"));
          let _1389_v13;
          _1389_v13 = new _dafny.CodePoint('m'.codePointAt(0));
          let _rhs165 = _1388_v12;
          let _rhs166 = (new BigNumber(164)).plus(new BigNumber((_1385_v10).cardinality()));
          let _rhs167 = _1389_v13;
          _1388_v12 = _rhs165;
          _1372_v0 = _rhs166;
          _1389_v13 = _rhs167;
          let _1390_v14;
          _1390_v14 = _module.D4.create_DC12(_dafny.Seq.of((_this).f15));
          let _1391_v15;
          _1391_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_1372_v0);
          let _1392_v16;
          _1392_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1386_i1,new BigNumber((_1391_v15).length));
          let _1393_v17;
          _1393_v17 = _dafny.Seq.of(_1392_v16);
          let _1394_v18;
          _1394_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1386_i1),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ymee"), _module.__default.safeIndex(new BigNumber((_1393_v17).length), new BigNumber((_dafny.Seq.UnicodeFromString("ymee")).length)), _1389_v13)).length)));
          let _1395_v19;
          let _nw206 = new _module.C1();
          _nw206.__ctor((_this).f20, _module.__default.fm2(p0, globalState), ((((_this).f14)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((p2).length)), new BigNumber(((_this).f14).length))]) ? ((_this).f13) : ((_this).f13)), _dafny.Seq.update((_1390_v14).dtor_cf18, _module.__default.safeIndex(new BigNumber(((_1394_v18).update(_1386_i1, new BigNumber((_1373_v1).length))).length), new BigNumber(((_1390_v14).dtor_cf18).length)), p1), _this.f12);
          _1395_v19 = _nw206;
          (_1395_v19).f18 = p0;
          let _1396_v20;
          let _init37 = ((_1397_v19) => function (_1398_i2) {
            return (_1397_v19).f17;
          })(_1395_v19);
          let _nw207 = Array((new BigNumber(22)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw207.length); _i0_37++) {
            _nw207[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1396_v20 = _nw207;
          let _1399_v21;
          _1399_v21 = _dafny.Seq.of(_1396_v20);
          let _1400_v22;
          _1400_v22 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_1399_v21, _1399_v21), _module.__default.safeIndex((_this).fm7(new BigNumber(708), globalState), new BigNumber((_dafny.Seq.Concat(_1399_v21, _1399_v21)).length)), _1396_v20),_1386_i1);
          _1400_v22 = (_1400_v22).update(_1399_v21, (_dafny.ZERO).minus((_1386_i1).plus(_1372_v0)));
          let _1401_v23;
          _1401_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,p0);
          let _1402_v24;
          _1402_v24 = _module.D3.create_DC8(_1401_v23);
          let _1403_v25;
          _1403_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1372_v0,true);
          _1402_v24 = _module.D3.create_DC8((_1401_v23).update((((_1403_v25).contains(_1372_v0)) ? ((_1403_v25).get(_1372_v0)) : ((_this).f20)), false));
        } else {
          let _1404_v26;
          _1404_v26 = _module.D4.create_DC12((_this).f14);
          let _1405_v27;
          _1405_v27 = _dafny.MultiSet.fromElements(_1386_i1, new BigNumber((_1373_v1).length), _1386_i1);
          let _1406_v28;
          _1406_v28 = new _dafny.CodePoint('v'.codePointAt(0));
          let _1407_v29;
          let _nw208 = Array((new BigNumber(24)).toNumber());
          _nw208[(_dafny.ZERO).toNumber()] = _1404_v26;
          _nw208[(_dafny.ONE).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(2)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(3)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(4)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(5)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(6)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(7)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(8)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(9)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(10)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(11)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(12)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(13)).toNumber()] = _module.__default.fm37((_this).f14, _1405_v27, new BigNumber(843), new _dafny.CodePoint('p'.codePointAt(0)), globalState);
          _nw208[(new BigNumber(14)).toNumber()] = function (_pat_let23_0) {
            return function (_1408_dt__update__tmp_h1) {
              return function (_pat_let24_0) {
                return function (_1409_dt__update_hcf18_h0) {
                  return _module.D4.create_DC12(_1409_dt__update_hcf18_h0);
                }(_pat_let24_0);
              }((_this).f14);
            }(_pat_let23_0);
          }(_1404_v26);
          _nw208[(new BigNumber(15)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(16)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(17)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(18)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(19)).toNumber()] = _module.__default.fm37((_this).f14, _1405_v27, _1386_i1, _1406_v28, globalState);
          _nw208[(new BigNumber(20)).toNumber()] = function (_pat_let25_0) {
            return function (_1410_dt__update__tmp_h2) {
              return function (_pat_let26_0) {
                return function (_1411_dt__update_hcf18_h1) {
                  return _module.D4.create_DC12(_1411_dt__update_hcf18_h1);
                }(_pat_let26_0);
              }((_this).f14);
            }(_pat_let25_0);
          }(_1404_v26);
          _nw208[(new BigNumber(21)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(22)).toNumber()] = _1404_v26;
          _nw208[(new BigNumber(23)).toNumber()] = _1404_v26;
          _1407_v29 = _nw208;
          let _1412_v30;
          let _nw209 = Array((new BigNumber(6)).toNumber());
          _nw209[(_dafny.ZERO).toNumber()] = (_this).f15;
          _nw209[(_dafny.ONE).toNumber()] = p1;
          _nw209[(new BigNumber(2)).toNumber()] = p1;
          _nw209[(new BigNumber(3)).toNumber()] = p1;
          _nw209[(new BigNumber(4)).toNumber()] = (_this).f20;
          _nw209[(new BigNumber(5)).toNumber()] = (_this).f20;
          _1412_v30 = _nw209;
          let _1413_v31;
          _1413_v31 = _dafny.MultiSet.fromElements(_1412_v30, _1412_v30);
          let _rhs168 = _1386_i1;
          let _rhs169 = _1407_v29;
          let _rhs170 = _module.__default.fm2(!(_1413_v31).contains(_1412_v30), globalState);
          r0 = _rhs168;
          _1407_v29 = _rhs169;
          r2 = _rhs170;
          let _1414_v32;
          _1414_v32 = _dafny.Seq.of(((_module.D0.create_DC2(_1386_i1, p0)).dtor_cf1).isLessThan(_1372_v0), !(!(!((_this).f15))) || (!(p1)));
          _1414_v32 = _dafny.Seq.Concat((_this).f14, (_this).f14);
          let _1415_v33;
          _1415_v33 = _dafny.Seq.of(p2, p2, _dafny.Seq.Concat(p2, p2), p2);
          _1415_v33 = _dafny.Seq.Concat(_1415_v33, _dafny.Seq.Concat(_1415_v33, _dafny.Seq.of(p2, p2)));
          let _index186 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1412_v30).length));
          (_1412_v30)[_index186] = p0;
          let _1416_v34;
          _1416_v34 = _dafny.Set.fromElements((_this).f15, (_this).f20, (_this).f20);
          let _index187 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1412_v30).length));
          (_1412_v30)[_index187] = (_1386_i1).isLessThan(new BigNumber((_1416_v34).length));
          (_this).m10(globalState);
        }
        let _1417_v35;
        let _nw210 = new _module.C2();
        _nw210.__ctor(_dafny.Seq.of(new BigNumber(78)), (_this).f20, (_this).f13, (_this).f14, _this.f12);
        _1417_v35 = _nw210;
        let _1418_v36;
        _1418_v36 = _dafny.Seq.of(_1417_v35);
        let _1419_v37;
        let _nw211 = Array((new BigNumber(12)).toNumber());
        _nw211[(_dafny.ZERO).toNumber()] = (((_this).f20) ? (_1417_v35) : (_1417_v35));
        _nw211[(_dafny.ONE).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(2)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(3)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(4)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(5)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(6)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(7)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(8)).toNumber()] = (_1418_v36)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_1418_v36).length))];
        _nw211[(new BigNumber(9)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(10)).toNumber()] = _1417_v35;
        _nw211[(new BigNumber(11)).toNumber()] = _1417_v35;
        _1419_v37 = _nw211;
        let _index188 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1419_v37).length));
        (_1419_v37)[_index188] = _1417_v35;
        let _1420_v38;
        _1420_v38 = _dafny.Set.fromElements((_this).f15, false, p1, p0);
        let _1421_v39;
        _1421_v39 = _dafny.Seq.of(_1386_i1, new BigNumber((_1420_v38).length), _1386_i1, new BigNumber(((_this).f14).length));
        let _index189 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1419_v37).length));
        let _nw212 = new _module.C2();
        _nw212.__ctor(_1421_v39, (_this).f15, (_this).f13, (_this).f14, _this.f12);
        (_1419_v37)[_index189] = _nw212;
        let _1422_v40;
        let _nw213 = new _module.C0();
        _nw213.__ctor();
        _1422_v40 = _nw213;
        let _1423_v41;
        _1423_v41 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(219)), function (_1424_i3) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }));
        _1423_v41 = _dafny.Seq.Concat(_1423_v41, _1423_v41);
      }
      r0 = _module.__default.safeModuloInt(_1372_v0, new BigNumber((_module.__default.fm57(p2, (_this).f15, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1385_v10).cardinality()),(_this).f15), globalState)).cardinality()));
      let _1425_v42;
      _1425_v42 = new _dafny.CodePoint('r'.codePointAt(0));
      _1425_v42 = _1425_v42;
      let _1426_v43;
      let _nw214 = new _module.C1();
      _nw214.__ctor((_this).f15, (_this).f15, (_this).f13, (_this).f14, _this.f12);
      _1426_v43 = _nw214;
      let _1427_v44;
      _1427_v44 = _module.D3.create_DC10(p0, _1372_v0, _1426_v43);
      _1427_v44 = _1427_v44;
      r0 = _1372_v0;
      r1 = _1372_v0;
      r2 = (_this).f20;
      let _1428_v46;
      _1428_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1372_v0,p0);
      r3 = ((function () {
        let _coll52 = new _dafny.Map();
        for (const _compr_52 of _dafny.IntegerRange(new BigNumber(455), new BigNumber(775))) {
          let _1429_v45 = _compr_52;
          if (((new BigNumber(455)).isLessThanOrEqualTo(_1429_v45)) && ((_1429_v45).isLessThan(new BigNumber(775)))) {
            _coll52.push([_module.__default.safeModuloInt(_1429_v45, _1372_v0),(_module.D7.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1385_v10).cardinality()),(_this).f15), (_this).f15, p1)).dtor_cf32]);
          }
        }
        return _coll52;
      }()).Merge(_1428_v46)).Merge(_1428_v46);
      return [r0, r1, r2, r3];
    }
    m10(globalState) {
      let _this = this;
      let _1430_v0;
      let _nw215 = Array((new BigNumber(23)).toNumber()).fill(false);
      _1430_v0 = _nw215;
      let _1431_v1;
      _1431_v1 = new BigNumber(414);
      let _1432_v2;
      _1432_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1430_v0,_1431_v1);
      _1432_v2 = _1432_v2;
      _1431_v1 = _1431_v1;
      let _1433_v3;
      _1433_v3 = _module.D8.create_DC24((_this).f15);
      let _pat_let_tv29 = _1431_v1;
      let _pat_let_tv30 = _1431_v1;
      _1431_v1 = function (_source23) {
        if (_source23.is_DC24) {
          let _1434___mcc_h0 = (_source23).cf36;
          let _1435_cf36 = _1434___mcc_h0;
          return _pat_let_tv29;
        } else if (_source23.is_DC25) {
          let _1436___mcc_h1 = (_source23).cf37;
          let _1437___mcc_h2 = (_source23).cf38;
          let _1438___mcc_h3 = (_source23).cf39;
          let _1439___mcc_h4 = (_source23).cf40;
          let _1440_cf40 = _1439___mcc_h4;
          let _1441_cf39 = _1438___mcc_h3;
          let _1442_cf38 = _1437___mcc_h2;
          let _1443_cf37 = _1436___mcc_h1;
          return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_module.D3.create_DC9(new BigNumber(906), new BigNumber((_dafny.Set.fromElements(_1443_cf37)).length))), _dafny.Seq.of(_module.D3.create_DC9(_1442_cf38, _1441_cf39)))).length);
        } else if (_source23.is_DC26) {
          let _1444___mcc_h5 = (_source23).cf41;
          let _1445_cf41 = _1444___mcc_h5;
          return (_pat_let_tv30).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1445_cf41,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15)).length))).length));
        } else if (_source23.is_DC23) {
          let _1446___mcc_h6 = (_source23).cf35;
          let _1447_cf35 = _1446___mcc_h6;
          return new BigNumber((_dafny.Seq.UnicodeFromString("tajrmdyiq")).length);
        } else {
          let _1448___mcc_h7 = (_source23).cf42;
          let _1449_cf42 = _1448___mcc_h7;
          return new BigNumber(((_dafny.Set.fromElements((_this).f20, true, (_this).f20, (_this).f15, (_this).f15)).Difference(_dafny.Set.fromElements((_this).f20, (_this).f20))).length);
        }
      }(_1433_v3);
      let _hi9 = _1431_v1;
      for (let _1450_i0 = _1431_v1; _1450_i0.isLessThan(_hi9); _1450_i0 = _1450_i0.plus(_dafny.ONE)) {
        let _1451_v4;
        _1451_v4 = new _dafny.CodePoint('k'.codePointAt(0));
        _1451_v4 = _1451_v4;
        let _1452_v5;
        _1452_v5 = _dafny.Seq.of(_1431_v1);
        (globalState).f4 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1452_v5, _dafny.Seq.update(_1452_v5, _module.__default.safeIndex(new BigNumber(-57), new BigNumber((_1452_v5).length)), _1431_v1)), _1452_v5);
        _1432_v2 = _1432_v2;
        let _1453_v6;
        _1453_v6 = _dafny.Set.fromElements(_1450_i0, _1431_v1);
        let _1454_v7;
        _1454_v7 = _dafny.Seq.of((_this).f15, (_this).f15, (_1453_v6).IsProperSubsetOf(_1453_v6), (_this).f20, (((_this).f20) ? (_module.__default.fm2((_this).f15, globalState)) : ((_this).f15)));
        _1454_v7 = _dafny.Seq.Concat((_this).f14, _1454_v7);
      }
      let _1455_v8;
      _1455_v8 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1431_v1),(_this).f20);
      let _index190 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1430_v0).length));
      (_1430_v0)[_index190] = (((_1455_v8).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_1457_i1) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length))) ? ((_1455_v8).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_1456_i1) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length))) : (true));
      let _index191 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1430_v0).length));
      (_1430_v0)[_index191] = false;
      (globalState).f4 = (_module.__default.safeModuloInt(new BigNumber(407), new BigNumber(562))).isLessThan(_1431_v1);
      return;
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [_module.T4];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return function () {
        let _coll53 = new _dafny.Set();
        for (const _compr_53 of ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(495),false)).Merge(function () {
          let _coll54 = new _dafny.Map();
          for (const _compr_54 of _dafny.IntegerRange(new BigNumber(914), new BigNumber(952))) {
            let _1458_v0 = _compr_54;
            if (((new BigNumber(914)).isLessThanOrEqualTo(_1458_v0)) && ((_1458_v0).isLessThan(new BigNumber(952)))) {
              _coll54.push([_module.__default.safeDivisionInt(_1458_v0, new BigNumber(-486)),true]);
            }
          }
          return _coll54;
        }())).Keys.Elements) {
          let _1459_v1 = _compr_53;
          if (((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(495),false)).Merge(function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of _dafny.IntegerRange(new BigNumber(914), new BigNumber(952))) {
              let _1458_v0 = _compr_55;
              if (((new BigNumber(914)).isLessThanOrEqualTo(_1458_v0)) && ((_1458_v0).isLessThan(new BigNumber(952)))) {
                _coll55.push([_module.__default.safeDivisionInt(_1458_v0, new BigNumber(-486)),true]);
              }
            }
            return _coll55;
          }())).contains(_1459_v1)) {
            _coll53.add((_1459_v1).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("s")).length)));
          }
        }
        return _coll53;
      }();
    };
    fm26(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.D5.create_DC15(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(false))).length)))).dtor_cf25;
    };
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _1460_v0;
      _1460_v0 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1461_v1;
      _1461_v1 = _dafny.Set.fromElements(_1460_v0);
      let _1462_v2;
      _1462_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1461_v1).length),new BigNumber((_dafny.MultiSet.fromElements(_1460_v0)).cardinality()));
      let _1463_v3;
      _1463_v3 = false;
      let _1464_v4;
      _1464_v4 = _module.D0.create_DC2(p2, !(_1463_v3));
      if (((((_1462_v2).contains(new BigNumber(148))) ? ((_1462_v2).get(new BigNumber(148))) : ((_1464_v4).dtor_cf1))).isEqualTo((((_1462_v2).contains(p2)) ? ((_1462_v2).get(p2)) : (p2)))) {
        (globalState).f4 = _1463_v3;
        let _1465_v5;
        _1465_v5 = _dafny.Seq.of(p1);
        r0 = (p2).isLessThan(_module.__default.safeModuloInt(p2, new BigNumber((_1465_v5).length)));
        let _1466_v6;
        _1466_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1463_v3,p2);
        let _1467_v7;
        _1467_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1466_v6);
        let _1468_v8;
        _1468_v8 = _module.D3.create_DC9((_module.__default.fm0(_1463_v3, _1467_v7, _1463_v3, globalState)).multipliedBy((_dafny.ZERO).minus(p2)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_1469_i0) {
  return new _dafny.CodePoint('a'.codePointAt(0));
})).length));
        let _source24 = _1468_v8;
        if (_source24.is_DC9) {
          let _1470___mcc_h0 = (_source24).cf12;
          let _1471___mcc_h1 = (_source24).cf13;
          let _1472_cf13 = _1471___mcc_h1;
          let _1473_cf12 = _1470___mcc_h0;
          (globalState).f4 = _module.__default.fm2(!(_1463_v3), globalState);
          _1472_cf13 = _module.__default.fm0(true, _1467_v7, !(_1463_v3), globalState);
          let _1474_v9;
          _1474_v9 = _module.D6.create_DC18(p2);
          _1474_v9 = _module.D6.create_DC18(_module.__default.safeDivisionInt(_1473_cf12, _1472_cf13));
          (globalState).f4 = _1463_v3;
        } else if (_source24.is_DC10) {
          let _1475___mcc_h2 = (_source24).cf14;
          let _1476___mcc_h3 = (_source24).cf15;
          let _1477___mcc_h4 = (_source24).cf16;
          let _1478_cf16 = _1477___mcc_h4;
          let _1479_cf15 = _1476___mcc_h3;
          let _1480_cf14 = _1475___mcc_h2;
          _1480_cf14 = ((_1478_cf16).f14)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1479_cf15, _1479_cf15)), new BigNumber(((_1478_cf16).f14).length))];
          let _1481_v10;
          let _init38 = ((_1482_v0, _1483_p1) => function (_1484_i1) {
            return _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), ((_1485_v0) => function (_1486_i2) {
              return _1485_v0;
            })(_1482_v0)), _module.__default.safeIndex(new BigNumber((_1483_p1).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), ((_1487_v0) => function (_1488_i2) {
              return _1487_v0;
            })(_1482_v0))).length)), new _dafny.CodePoint('f'.codePointAt(0)));
          })(_1460_v0, p1);
          let _nw216 = Array((new BigNumber(12)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw216.length); _i0_38++) {
            _nw216[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1481_v10 = _nw216;
          let _index192 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_1481_v10).length));
          (_1481_v10)[_index192] = p1;
          let _index193 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_1481_v10).length));
          (_1481_v10)[_index193] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uqnkr"), p1);
          let _1489_v11;
          let _nw217 = new _module.C1();
          _nw217.__ctor(_1480_cf14, _1463_v3, (_1478_cf16).f13, ((_1463_v3) ? ((_1478_cf16).f14) : ((_1478_cf16).f14)), _1478_cf16.f12);
          _1489_v11 = _nw217;
          _1489_v11 = _1489_v11;
          let _1490_v12;
          let _1491_v13;
          let _1492_v14;
          let _out93;
          let _out94;
          let _out95;
          let _outcollector24 = (_this).m8(globalState);
          _out93 = _outcollector24[0];
          _out94 = _outcollector24[1];
          _out95 = _outcollector24[2];
          _1490_v12 = _out93;
          _1491_v13 = _out94;
          _1492_v14 = _out95;
        } else if (_source24.is_DC8) {
          let _1493___mcc_h5 = (_source24).cf11;
          let _1494_cf11 = _1493___mcc_h5;
          let _1495_v15;
          _1495_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1463_v3);
          let _1496_v16;
          _1496_v16 = _dafny.MultiSet.fromElements(p2, p2);
          let _1497_v17;
          _1497_v17 = _dafny.Seq.of(_1496_v16, _1496_v16, _1496_v16, _1496_v16, (_1496_v16).update(new BigNumber(955), _module.__default.abs(p2)));
          let _1498_v18;
          _1498_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1495_v15,new BigNumber(((_1497_v17)[_module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_1497_v17).length))]).cardinality()));
          let _1499_v19;
          let _nw218 = Array((new BigNumber(19)).toNumber());
          _nw218[(_dafny.ZERO).toNumber()] = p2;
          _nw218[(_dafny.ONE).toNumber()] = p2;
          _nw218[(new BigNumber(2)).toNumber()] = new BigNumber(((p0).Intersect(_module.__default.fm27(globalState))).length);
          _nw218[(new BigNumber(3)).toNumber()] = p2;
          _nw218[(new BigNumber(4)).toNumber()] = p2;
          _nw218[(new BigNumber(5)).toNumber()] = p2;
          _nw218[(new BigNumber(6)).toNumber()] = p2;
          _nw218[(new BigNumber(7)).toNumber()] = (((_1466_v6).contains(_1463_v3)) ? ((_1466_v6).get(_1463_v3)) : (p2));
          _nw218[(new BigNumber(8)).toNumber()] = p2;
          _nw218[(new BigNumber(9)).toNumber()] = (p2).minus(p2);
          _nw218[(new BigNumber(10)).toNumber()] = (new BigNumber(656)).plus(_module.__default.fm0(_1463_v3, _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(p2)),_1466_v6), _1463_v3, globalState));
          _nw218[(new BigNumber(11)).toNumber()] = p2;
          _nw218[(new BigNumber(12)).toNumber()] = p2;
          _nw218[(new BigNumber(13)).toNumber()] = p2;
          _nw218[(new BigNumber(14)).toNumber()] = p2;
          _nw218[(new BigNumber(15)).toNumber()] = (((_1498_v18).contains(_1495_v15)) ? ((_1498_v18).get(_1495_v15)) : (p2));
          _nw218[(new BigNumber(16)).toNumber()] = p2;
          _nw218[(new BigNumber(17)).toNumber()] = new BigNumber((_1494_cf11).length);
          _nw218[(new BigNumber(18)).toNumber()] = p2;
          _1499_v19 = _nw218;
          let _index194 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1499_v19).length));
          (_1499_v19)[_index194] = p2;
          let _index195 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1499_v19).length));
          (_1499_v19)[_index195] = p2;
          let _1500_v20;
          _1500_v20 = _dafny.Seq.of(_1463_v3);
          let _index196 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1499_v19).length));
          let _rhs171 = _1463_v3;
          let _rhs172 = (((_1499_v19)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_1499_v19).length))]).minus(p2)).minus((new BigNumber((_1500_v20).length)).multipliedBy(p2));
          let _lhs97 = _1499_v19;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1499_v19).length));
          _1463_v3 = _rhs171;
          _lhs97[_lhs98] = _rhs172;
          let _1501_v21;
          _1501_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1463_v3,_1494_cf11);
          let _1502_v22;
          let _nw219 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1502_v22 = _nw219;
          let _1503_v23;
          let _nw220 = new _module.C1();
          _nw220.__ctor(_1463_v3, _1463_v3, _1501_v21, _dafny.Seq.of(_1463_v3, !(_1463_v3)), _1502_v22);
          _1503_v23 = _nw220;
          let _1504_v24;
          let _nw221 = new _module.C0();
          _nw221.__ctor();
          _1504_v24 = _nw221;
          let _1505_v25;
          _1505_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v24,_module.__default.fm2(true, globalState));
          let _1506_v26;
          let _nw222 = Array((new BigNumber(15)).toNumber());
          _nw222[(_dafny.ZERO).toNumber()] = _1503_v23.f18;
          _nw222[(_dafny.ONE).toNumber()] = _1463_v3;
          _nw222[(new BigNumber(2)).toNumber()] = _1503_v23.f18;
          _nw222[(new BigNumber(3)).toNumber()] = (((_1505_v25).contains(_1504_v24)) ? ((_1505_v25).get(_1504_v24)) : (_1503_v23.f18));
          _nw222[(new BigNumber(4)).toNumber()] = !((_1503_v23).f17);
          _nw222[(new BigNumber(5)).toNumber()] = _1503_v23.f18;
          _nw222[(new BigNumber(6)).toNumber()] = _1463_v3;
          _nw222[(new BigNumber(7)).toNumber()] = _1503_v23.f18;
          _nw222[(new BigNumber(8)).toNumber()] = _1463_v3;
          _nw222[(new BigNumber(9)).toNumber()] = true;
          _nw222[(new BigNumber(10)).toNumber()] = (_1503_v23).f17;
          _nw222[(new BigNumber(11)).toNumber()] = _1503_v23.f18;
          _nw222[(new BigNumber(12)).toNumber()] = (_1503_v23).f17;
          _nw222[(new BigNumber(13)).toNumber()] = (_1503_v23).f17;
          _nw222[(new BigNumber(14)).toNumber()] = (_1503_v23).f17;
          _1506_v26 = _nw222;
          let _1507_v27;
          _1507_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1506_v26,_1503_v23.f18);
          let _index197 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1499_v19).length));
          (_1499_v19)[_index197] = _module.__default.fm0((((_1500_v20)[_module.__default.safeIndex(new BigNumber((_1507_v27).length), new BigNumber((_1500_v20).length))]) ? (_1463_v3) : ((_1503_v23).f17)), _1467_v7, true, globalState);
        } else {
          let _1508___mcc_h6 = (_source24).cf17;
          let _1509_cf17 = _1508___mcc_h6;
          let _1510_v28;
          _1510_v28 = new BigNumber(-781);
          _1510_v28 = (_dafny.ZERO).minus(p2);
          _1510_v28 = new BigNumber((_module.__default.fm27(globalState)).length);
          let _1511_v29;
          _1511_v29 = _dafny.Set.fromElements(p1);
          (globalState).f4 = !(((_1511_v29).Union(_1511_v29)).IsSubsetOf(_1511_v29));
          let _1512_v30;
          let _nw223 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1512_v30 = _nw223;
          let _1513_v31;
          _1513_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1512_v30,_1510_v28);
          let _1514_v32;
          _1514_v32 = _dafny.MultiSet.fromElements(p2);
          let _1515_v33;
          let _nw224 = Array((new BigNumber(23)).toNumber()).fill(_module.D5.Default());
          _1515_v33 = _nw224;
          let _1516_v34;
          _1516_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_1513_v31);
          let _rhs173 = (_dafny.Map.Empty.slice().updateUnsafe(_1512_v30,_1510_v28)).Merge((((_1516_v34).contains(p2)) ? ((_1516_v34).get(p2)) : (_1513_v31)));
          let _rhs174 = (((new BigNumber(-977)).isLessThanOrEqualTo(_1510_v28)) ? (_1514_v32) : (_1514_v32));
          let _rhs175 = (((p2).isLessThan(p2)) ? (_1515_v33) : (_1515_v33));
          _1513_v31 = _rhs173;
          _1514_v32 = _rhs174;
          _1515_v33 = _rhs175;
        }
        let _1517_v35;
        let _nw225 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1517_v35 = _nw225;
        let _index198 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_1517_v35).length));
        (_1517_v35)[_index198] = _module.__default.fm28(p1, _1463_v3, p2, globalState);
        let _1518_v36;
        _1518_v36 = _dafny.MultiSet.fromElements(p2);
        let _index199 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_1517_v35).length));
        (_1517_v35)[_index199] = _1518_v36;
        (globalState).f4 = _1463_v3;
      } else {
        let _1519_v37;
        _1519_v37 = _dafny.Map.Empty.slice().updateUnsafe(!(_1463_v3),_1463_v3);
        let _1520_v38;
        let _nw226 = Array((new BigNumber(15)).toNumber());
        _nw226[(_dafny.ZERO).toNumber()] = _1463_v3;
        _nw226[(_dafny.ONE).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(2)).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(3)).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(4)).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(5)).toNumber()] = !(!(_1463_v3));
        _nw226[(new BigNumber(6)).toNumber()] = (_1463_v3) && (_1463_v3);
        _nw226[(new BigNumber(7)).toNumber()] = !(_1519_v37).contains(true);
        _nw226[(new BigNumber(8)).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(9)).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsPrefixOf(p1, p1);
        _nw226[(new BigNumber(11)).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(12)).toNumber()] = _1463_v3;
        _nw226[(new BigNumber(13)).toNumber()] = !(_1463_v3) || (_1463_v3);
        _nw226[(new BigNumber(14)).toNumber()] = _1463_v3;
        _1520_v38 = _nw226;
        let _index200 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length));
        (_1520_v38)[_index200] = _1463_v3;
        let _1521_v39;
        _1521_v39 = _dafny.MultiSet.fromElements(_1464_v4);
        let _index201 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length));
        let _rhs176 = !((((_1519_v37).contains(_1463_v3)) ? ((_1519_v37).get(_1463_v3)) : (false))) || (true);
        let _rhs177 = _1463_v3;
        let _rhs178 = _1521_v39;
        let _lhs99 = globalState;
        let _lhs100 = _1520_v38;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length));
        _lhs99.f4 = _rhs176;
        _lhs100[_lhs101] = _rhs177;
        _1521_v39 = _rhs178;
        let _index202 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length));
        (_1520_v38)[_index202] = (_1520_v38)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length))];
        let _index203 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length));
        (_1520_v38)[_index203] = !(_1463_v3) || ((_1520_v38)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length))]);
        let _1522_v40;
        let _nw227 = Array((new BigNumber(17)).toNumber()).fill(_module.D2.Default());
        _1522_v40 = _nw227;
        let _1523_v41;
        _1523_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1522_v40,p2);
        let _1524_v42;
        _1524_v42 = _dafny.Seq.of((((_1523_v41).contains(_1522_v40)) ? ((_1523_v41).get(_1522_v40)) : (p2)));
        _1524_v42 = _dafny.Seq.Concat(_1524_v42, _dafny.Seq.of(new BigNumber(77)));
        let _index204 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length));
        (_1520_v38)[_index204] = (_1520_v38)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1520_v38).length))];
      }
      let _1525_v43;
      _1525_v43 = new BigNumber(584);
      let _1526_v44;
      _1526_v44 = _dafny.Seq.of(_1463_v3);
      _1525_v43 = ((_1525_v43).plus(new BigNumber((_1526_v44).length))).plus(_1525_v43);
      _1525_v43 = (_dafny.ZERO).minus((p2).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), function (_1527_i3) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length), p2)));
      let _1528_v45;
      _1528_v45 = _module.D0.create_DC1();
      _1528_v45 = _1528_v45;
      let _hi10 = (p2).multipliedBy(_1525_v43);
      for (let _1529_i4 = (p2).multipliedBy(p2); _1529_i4.isLessThan(_hi10); _1529_i4 = _1529_i4.plus(_dafny.ONE)) {
        _1525_v43 = (_module.__default.safeModuloInt(p2, _1525_v43)).plus(_1529_i4);
        _1525_v43 = ((new BigNumber(375)).minus((_dafny.ZERO).minus(new BigNumber(-426)))).plus(new BigNumber(-117));
        let _1530_v46;
        let _init39 = ((_1531_v3) => function (_1532_i5) {
          return (_dafny.MultiSet.fromElements(_1531_v3, _1531_v3)).equals(_dafny.MultiSet.fromElements(!(_1531_v3)));
        })(_1463_v3);
        let _nw228 = Array((new BigNumber(6)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw228.length); _i0_39++) {
          _nw228[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1530_v46 = _nw228;
        _1530_v46 = _1530_v46;
        _1463_v3 = _1463_v3;
      }
      let _1533_v47;
      let _init40 = ((_1534_p2) => function (_1535_i6) {
        return _module.__default.safeDivisionInt(_1535_i6, _1534_p2);
      })(p2);
      let _nw229 = Array((new BigNumber(18)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw229.length); _i0_40++) {
        _nw229[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1533_v47 = _nw229;
      let _index205 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1533_v47).length));
      (_1533_v47)[_index205] = new BigNumber((p0).length);
      let _1536_v48;
      _1536_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1460_v0,_1463_v3);
      let _index206 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1533_v47).length));
      (_1533_v47)[_index206] = (new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((p1).length)), _1460_v0)).length)).multipliedBy(new BigNumber((_1536_v48).length));
      r0 = _1463_v3;
      return r0;
    }
    m6(globalState) {
      let _this = this;
      let r0 = false;
      let _1537_v0;
      _1537_v0 = false;
      let _1538_v1;
      _1538_v1 = _dafny.Seq.of(_1537_v0);
      let _1539_v2;
      _1539_v2 = _dafny.Seq.of(_1538_v1);
      let _1540_v3;
      _1540_v3 = _module.D2.create_DC4(_1539_v2);
      let _1541_v4;
      _1541_v4 = _module.D2.create_DC7(_1540_v3);
      let _1542_v5;
      _1542_v5 = _module.D2.create_DC7(_1541_v4);
      let _1543_v6;
      _1543_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1542_v5);
      _1543_v6 = (_1543_v6).update(_1537_v0, _1542_v5);
      if (_1537_v0) {
        let _1544_v7;
        let _nw230 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
        _1544_v7 = _nw230;
        let _1545_v8;
        let _nw231 = Array((new BigNumber(6)).toNumber());
        _nw231[(_dafny.ZERO).toNumber()] = _1537_v0;
        _nw231[(_dafny.ONE).toNumber()] = _1537_v0;
        _nw231[(new BigNumber(2)).toNumber()] = _1537_v0;
        _nw231[(new BigNumber(3)).toNumber()] = (_module.__default.fm2(_1537_v0, globalState)) || (_1537_v0);
        _nw231[(new BigNumber(4)).toNumber()] = _1537_v0;
        _nw231[(new BigNumber(5)).toNumber()] = (_1537_v0) && (_1537_v0);
        _1545_v8 = _nw231;
        let _1546_v9;
        _1546_v9 = new BigNumber(774);
        let _1547_v10;
        _1547_v10 = _module.D2.create_DC6(_1545_v8, _1545_v8);
        let _1548_v11;
        _1548_v11 = _module.D3.create_DC9(_1546_v9, new BigNumber((_1538_v1).length));
        let _rhs179 = _1544_v7;
        let _rhs180 = (_1547_v10).dtor_cf8;
        let _rhs181 = (_dafny.ZERO).minus((_1548_v11).dtor_cf12);
        _1544_v7 = _rhs179;
        _1545_v8 = _rhs180;
        _1546_v9 = _rhs181;
        let _1549_v12;
        _1549_v12 = _dafny.Seq.UnicodeFromString("gcj");
        let _1550_v13;
        _1550_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1537_v0);
        let _1551_v14;
        _1551_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1550_v13);
        let _1552_v15;
        let _nw232 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1552_v15 = _nw232;
        let _1553_v16;
        let _nw233 = new _module.C1();
        _nw233.__ctor(_1537_v0, _1537_v0, _1551_v14, _1538_v1, _1552_v15);
        _1553_v16 = _nw233;
        let _1554_v17;
        _1554_v17 = _module.D3.create_DC10(true, new BigNumber(194), _1553_v16);
        let _1555_v18;
        _1555_v18 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1556_v19;
        let _nw234 = Array((new BigNumber(9)).toNumber());
        _nw234[(_dafny.ZERO).toNumber()] = _1549_v12;
        _nw234[(_dafny.ONE).toNumber()] = _1549_v12;
        _nw234[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm29((_1554_v17).dtor_cf14, _1537_v0, _1537_v0, _1546_v9, globalState), _1549_v12);
        _nw234[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_1549_v12, _module.__default.safeIndex(_1546_v9, new BigNumber((_1549_v12).length)), _1555_v18);
        _nw234[(new BigNumber(4)).toNumber()] = _1549_v12;
        _nw234[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1549_v12, _1549_v12);
        _nw234[(new BigNumber(6)).toNumber()] = _1549_v12;
        _nw234[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1549_v12, _1549_v12);
        _nw234[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pysu"), _dafny.Seq.UnicodeFromString("ndnibmluf"));
        _1556_v19 = _nw234;
        let _index207 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_1556_v19).length));
        (_1556_v19)[_index207] = _dafny.Seq.UnicodeFromString("pxcb");
        let _index208 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_1556_v19).length));
        (_1556_v19)[_index208] = _module.__default.fm29(_1537_v0, _module.__default.fm2(true, globalState), _1537_v0, _1546_v9, globalState);
        let _1557_v20;
        _1557_v20 = _dafny.MultiSet.fromElements(_1537_v0, _1537_v0, true);
        _1546_v9 = _module.__default.safeModuloInt((_1546_v9).minus(new BigNumber((_1557_v20).cardinality())), _1546_v9);
        if ((_1546_v9).isLessThanOrEqualTo(_1546_v9)) {
          let _1558_v21;
          let _nw235 = new _module.C0();
          _nw235.__ctor();
          _1558_v21 = _nw235;
          _1545_v8 = _1545_v8;
          let _1559_v22;
          let _nw236 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _1559_v22 = _nw236;
          let _index209 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_1559_v22).length));
          (_1559_v22)[_index209] = _1546_v9;
          let _1560_v23;
          _1560_v23 = _dafny.Set.fromElements(_1537_v0);
          let _index210 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_1559_v22).length));
          (_1559_v22)[_index210] = (_1546_v9).plus(_module.__default.safeModuloInt(new BigNumber((_1560_v23).length), _1546_v9));
          let _1561_v24;
          _1561_v24 = _dafny.Seq.of(_1545_v8, _1545_v8);
          let _index211 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_1559_v22).length));
          (_1559_v22)[_index211] = new BigNumber((_dafny.Seq.Concat(_1561_v24, _dafny.Seq.Concat(_dafny.Seq.of(_1545_v8, _1545_v8), _1561_v24))).length);
          _1555_v18 = _1555_v18;
        } else {
          let _1562_v25;
          _1562_v25 = _dafny.Seq.of(((_1537_v0) ? (_1546_v9) : ((_1553_v16).fm5(_1537_v0, globalState))), _1546_v9);
          _1546_v9 = (_1562_v25)[_module.__default.safeIndex((_1546_v9).plus(new BigNumber(((_1556_v19)[_module.__default.safeIndex(new BigNumber(325), new BigNumber((_1556_v19).length))]).length)), new BigNumber((_1562_v25).length))];
          let _1563_v26;
          let _nw237 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _1563_v26 = _nw237;
          let _1564_v27;
          _1564_v27 = _module.D2.create_DC5(new BigNumber((_1549_v12).length), _1537_v0, !(_1537_v0));
          let _index212 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_1563_v26).length));
          (_1563_v26)[_index212] = ((_1564_v27).dtor_cf5).multipliedBy(new BigNumber((function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of (_dafny.Seq.of(_1546_v9)).Elements) {
              let _1565_v28 = _compr_56;
              if (_dafny.Seq.contains(_dafny.Seq.of(_1546_v9), _1565_v28)) {
                _coll56.push([(_1565_v28).multipliedBy(_1546_v9),_dafny.MultiSet.fromElements(new BigNumber(283))]);
              }
            }
            return _coll56;
          }()).length));
          let _index213 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_1563_v26).length));
          (_1563_v26)[_index213] = new BigNumber(973);
          let _1566_v29;
          let _nw238 = new _module.C0();
          _nw238.__ctor();
          _1566_v29 = _nw238;
          let _index214 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_1563_v26).length));
          (_1563_v26)[_index214] = (_1563_v26)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_1563_v26).length))];
          let _1567_v30;
          _1567_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1546_v9,((_1537_v0) ? (_1545_v8) : (_1545_v8)));
          _1567_v30 = ((_1567_v30).Merge(_1567_v30)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1556_v19)[_module.__default.safeIndex(new BigNumber(325), new BigNumber((_1556_v19).length))]).length),_1545_v8)).update(new BigNumber((_dafny.Seq.update(_1562_v25, _module.__default.safeIndex(new BigNumber((_1562_v25).length), new BigNumber((_1562_v25).length)), _1546_v9)).length), _1545_v8));
        }
        let _1568_v31;
        _1568_v31 = _dafny.Set.fromElements(_1546_v9, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_1537_v0),_1550_v13)).length), _1546_v9, _1546_v9);
        let _1569_v32;
        _1569_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1568_v31,new BigNumber((_1549_v12).length));
        let _1570_v33;
        _1570_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,(_1569_v32).update(_dafny.Set.fromElements(new BigNumber(180)), (_dafny.ZERO).minus(_1546_v9)));
        _1570_v33 = (_1570_v33).update(_1537_v0, _1569_v32);
      } else {
        let _1571_v34;
        _1571_v34 = _dafny.Seq.UnicodeFromString("kbusqdbk");
        let _1572_v35;
        _1572_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1571_v34,new BigNumber(993));
        let _1573_v36;
        let _nw239 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1573_v36 = _nw239;
        let _1574_v37;
        _1574_v37 = _module.D2.create_DC6(_1573_v36, _1573_v36);
        let _1575_v38;
        _1575_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1572_v35).Merge(_1572_v35),_1574_v37);
        _1575_v38 = (_1575_v38).update(_1572_v35, _module.D2.create_DC6(_1573_v36, _1573_v36));
        let _1576_v39;
        _1576_v39 = new BigNumber(153);
        let _1577_v40;
        _1577_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1576_v39,_1537_v0);
        r0 = (((_1577_v40).contains(_1576_v39)) ? ((_1577_v40).get(_1576_v39)) : (_1537_v0));
        if (((new BigNumber((_1571_v34).length)).plus(_1576_v39)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(299)), ((_1578_v39) => function (_1579_i0) {
          return (_dafny.ZERO).minus(_1578_v39);
        })(_1576_v39))).length))) {
          let _1580_v41;
          _1580_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1576_v39);
          let _1581_v42;
          _1581_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1576_v39,_1580_v41);
          let _1582_v43;
          _1582_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1537_v0);
          let _1583_v44;
          _1583_v44 = _module.D3.create_DC8(_1582_v43);
          let _1584_v45;
          _1584_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,(_1583_v44).dtor_cf11);
          let _1585_v46;
          _1585_v46 = _dafny.MultiSet.fromElements(new BigNumber((_1580_v41).length), _1576_v39);
          let _1586_v47;
          _1586_v47 = _dafny.Seq.of(_1576_v39, (_dafny.ZERO).minus(_1576_v39));
          let _1587_v48;
          let _nw240 = Array((new BigNumber(15)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = _module.__default.fm28(_1571_v34, _1537_v0, _1576_v39, globalState);
          _nw240[(_dafny.ONE).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(2)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(3)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(4)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(5)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_1586_v47);
          _nw240[(new BigNumber(7)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(8)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber(-746), _1576_v39, _1576_v39);
          _nw240[(new BigNumber(10)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(11)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(12)).toNumber()] = _1585_v46;
          _nw240[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_1576_v39);
          _nw240[(new BigNumber(14)).toNumber()] = _1585_v46;
          _1587_v48 = _nw240;
          let _1588_v49;
          let _nw241 = new _module.C1();
          _nw241.__ctor(_1537_v0, _1537_v0, _1584_v45, _1538_v1, _1587_v48);
          _1588_v49 = _nw241;
          let _1589_v50;
          _1589_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1588_v49,(((_1582_v43).contains((_1588_v49).f17)) ? ((_1582_v43).get((_1588_v49).f17)) : (_1588_v49.f18)));
          let _1590_v51;
          _1590_v51 = _dafny.Set.fromElements(_module.__default.fm0(_1537_v0, (_1581_v42).update(new BigNumber((_1589_v50).length), _1580_v41), (_1588_v49).f17, globalState), (_dafny.ZERO).minus((_dafny.ZERO).minus(_1576_v39)), _1576_v39, _1576_v39);
          let _rhs182 = ((function () {
            let _coll57 = new _dafny.Set();
            for (const _compr_57 of (_1577_v40).Keys.Elements) {
              let _1591_v52 = _compr_57;
              if ((_1577_v40).contains(_1591_v52)) {
                _coll57.add(_module.__default.safeDivisionInt(_1591_v52, new BigNumber((_dafny.MultiSet.fromElements(true, !(true))).cardinality())));
              }
            }
            return _coll57;
          }()).Intersect(_1590_v51)).Difference(_1590_v51);
          let _rhs183 = _1573_v36;
          let _rhs184 = _1573_v36;
          let _rhs185 = (_1576_v39).multipliedBy((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-947), _1576_v39)));
          _1590_v51 = _rhs182;
          _1573_v36 = _rhs183;
          _1573_v36 = _rhs184;
          _1576_v39 = _rhs185;
          (globalState).f4 = (_1588_v49).f17;
          let _1592_v53;
          _1592_v53 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1593_v54;
          let _nw242 = Array((new BigNumber(24)).toNumber());
          _nw242[(_dafny.ZERO).toNumber()] = _1571_v34;
          _nw242[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("di"), _dafny.Seq.update(_1571_v34, _module.__default.safeIndex(new BigNumber(-842), new BigNumber((_1571_v34).length)), _1592_v53)), _module.__default.safeIndex(_1576_v39, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("di"), _dafny.Seq.update(_1571_v34, _module.__default.safeIndex(new BigNumber(-842), new BigNumber((_1571_v34).length)), _1592_v53))).length)), _1592_v53);
          _nw242[(new BigNumber(2)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(3)).toNumber()] = _module.__default.fm29((_1588_v49).fm17(_1586_v47, (_dafny.ZERO).minus((_1588_v49).fm5(_1588_v49.f18, globalState)), globalState), (_1588_v49).f17, _1588_v49.f18, new BigNumber(305), globalState);
          _nw242[(new BigNumber(4)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(5)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("qafisv");
          _nw242[(new BigNumber(7)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("qxdoh");
          _nw242[(new BigNumber(9)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("bnfm");
          _nw242[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_1571_v34, _module.__default.safeIndex(_1576_v39, new BigNumber((_1571_v34).length)), _1592_v53);
          _nw242[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("lwfvvvjpx");
          _nw242[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1571_v34, _1571_v34);
          _nw242[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1571_v34, _dafny.Seq.UnicodeFromString("pawwyk"));
          _nw242[(new BigNumber(15)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(16)).toNumber()] = _module.__default.fm29(_1588_v49.f18, _1537_v0, _1537_v0, (_module.D3.create_DC9(_1576_v39, _1576_v39)).dtor_cf12, globalState);
          _nw242[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_1571_v34, _1571_v34);
          _nw242[(new BigNumber(18)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1571_v34, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-652)), ((_1594_v53) => function (_1595_i1) {
            return _1594_v53;
          })(_1592_v53)));
          _nw242[(new BigNumber(20)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(21)).toNumber()] = _1571_v34;
          _nw242[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1571_v34, _dafny.Seq.UnicodeFromString("uqvycdvnb"));
          _nw242[(new BigNumber(23)).toNumber()] = _1571_v34;
          _1593_v54 = _nw242;
          let _index215 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1593_v54).length));
          (_1593_v54)[_index215] = _1571_v34;
          let _1596_v55;
          _1596_v55 = _dafny.Seq.of(_1571_v34);
          let _index216 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1593_v54).length));
          (_1593_v54)[_index216] = _dafny.Seq.Concat((_1596_v55)[_module.__default.safeIndex(_1576_v39, new BigNumber((_1596_v55).length))], _1571_v34);
          let _1597_v56;
          _1597_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray((_1593_v54)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_1593_v54).length))]),(_1576_v39).multipliedBy(_1576_v39));
          let _1598_v57;
          _1598_v57 = _dafny.MultiSet.fromElements(_1592_v53);
          let _1599_v58;
          _1599_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1576_v39,_1576_v39);
          _1597_v56 = (_1597_v56).update((_1598_v57).Difference((_1598_v57).update(_1592_v53, _module.__default.abs(new BigNumber(-271)))), (((_1599_v58).contains(new BigNumber((_1571_v34).length))) ? ((_1599_v58).get(new BigNumber((_1571_v34).length))) : (_1576_v39)));
          _1576_v39 = _1576_v39;
        } else {
          (globalState).f4 = (_1537_v0) && (_1537_v0);
          (globalState).f4 = ((!(_1537_v0) || (_1537_v0)) ? (_1537_v0) : (_1537_v0));
          let _1600_v59;
          let _nw243 = new _module.C0();
          _nw243.__ctor();
          _1600_v59 = _nw243;
          (globalState).f4 = _1537_v0;
          let _1601_v60;
          _1601_v60 = _dafny.Seq.of(_1571_v34);
          let _1602_v61;
          _1602_v61 = _dafny.Seq.of(_1576_v39, _1576_v39);
          let _1603_v62;
          _1603_v62 = _dafny.Map.Empty.slice().updateUnsafe(true,!(_1537_v0));
          let _1604_v63;
          _1604_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1603_v62);
          let _1605_v64;
          let _nw244 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1605_v64 = _nw244;
          let _1606_v65;
          let _nw245 = new _module.C1();
          _nw245.__ctor((new BigNumber(((_1601_v60)[_module.__default.safeIndex(_1576_v39, new BigNumber((_1601_v60).length))]).length)).isEqualTo(_1576_v39), _dafny.Seq.contains(_1602_v61, (((_1572_v35).contains(_1571_v34)) ? ((_1572_v35).get(_1571_v34)) : (_1576_v39))), _1604_v63, (_1539_v2)[_module.__default.safeIndex(_1576_v39, new BigNumber((_1539_v2).length))], _1605_v64);
          _1606_v65 = _nw245;
          _1606_v65 = _1606_v65;
        }
        (globalState).f4 = _1537_v0;
        let _1607_v66;
        let _nw246 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
        _1607_v66 = _nw246;
        let _1608_v67;
        _1608_v67 = _dafny.Set.fromElements(_1576_v39, _1576_v39);
        let _index217 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1607_v66).length));
        (_1607_v66)[_index217] = (_1608_v67).Intersect(_1608_v67);
        let _index218 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1607_v66).length));
        (_1607_v66)[_index218] = _1608_v67;
      }
      let _1609_v68;
      _1609_v68 = new _dafny.CodePoint('b'.codePointAt(0));
      _1609_v68 = _1609_v68;
      let _1610_v69;
      _1610_v69 = new BigNumber(475);
      let _hi11 = _1610_v69;
      for (let _1611_i2 = _1610_v69; _1611_i2.isLessThan(_hi11); _1611_i2 = _1611_i2.plus(_dafny.ONE)) {
        let _1612_v70;
        _1612_v70 = _dafny.Seq.UnicodeFromString("gqbu");
        _1612_v70 = _dafny.Seq.Concat(_1612_v70, _1612_v70);
        let _1613_v71;
        _1613_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1611_i2,_1610_v69);
        let _1614_v72;
        _1614_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,(((_1613_v71).contains(_1610_v69)) ? ((_1613_v71).get(_1610_v69)) : (_1611_i2)));
        _1614_v72 = (_1614_v72).update(_1537_v0, _1611_i2);
        let _1615_v73;
        let _nw247 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
        _1615_v73 = _nw247;
        _1615_v73 = _1615_v73;
        let _1616_v74;
        _1616_v74 = _dafny.MultiSet.fromElements(_1610_v69);
        let _1617_v76;
        _1617_v76 = _dafny.MultiSet.fromElements(_1609_v68);
        if ((_1616_v74).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1610_v69, _1610_v69, new BigNumber((_1612_v70).length), _1610_v69, new BigNumber((function () {
          let _coll58 = new _dafny.Map();
          for (const _compr_58 of ((_1617_v76).update(_1609_v68, _module.__default.abs(_1611_i2))).Elements) {
            let _1618_v75 = _compr_58;
            if (((_1617_v76).update(_1609_v68, _module.__default.abs(_1611_i2))).contains(_1618_v75)) {
              _coll58.push([_1618_v75,_1537_v0]);
            }
          }
          return _coll58;
        }()).length)))) {
          let _1619_v77;
          _1619_v77 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),(_1610_v69).minus(_1611_i2));
          let _1620_v78;
          _1620_v78 = _module.D6.create_DC19(_module.__default.fm30(_1610_v69, _1610_v69, globalState));
          let _rhs186 = _1619_v77;
          let _rhs187 = _1620_v78;
          let _rhs188 = ((((_1614_v72).contains(_1537_v0)) ? ((_1614_v72).get(_1537_v0)) : (_1611_i2))).isLessThanOrEqualTo(_1611_i2);
          let _rhs189 = _1610_v69;
          let _rhs190 = (_1610_v69).multipliedBy(_1611_i2);
          _1619_v77 = _rhs186;
          _1620_v78 = _rhs187;
          r0 = _rhs188;
          _1610_v69 = _rhs189;
          _1610_v69 = _rhs190;
          (globalState).f4 = (_1538_v1)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1538_v1).length))];
          (globalState).f4 = _1537_v0;
          r0 = true;
          _1614_v72 = (_1614_v72).update(!(_1537_v0) || (_module.__default.fm2(_1537_v0, globalState)), _1610_v69);
        } else {
          _1610_v69 = _1610_v69;
          let _1621_v79;
          _1621_v79 = _dafny.Set.fromElements(_1537_v0);
          _1610_v69 = (new BigNumber(((_1621_v79).Union(_1621_v79)).length)).plus(_1610_v69);
          let _1622_v80;
          let _1623_v81;
          let _1624_v82;
          let _1625_v83;
          let _out96;
          let _out97;
          let _out98;
          let _out99;
          let _outcollector25 = _module.__default.m0((_dafny.Set.fromElements(false, _module.__default.fm2(_1537_v0, globalState))).IsProperSubsetOf(_1621_v79), _module.__default.fm31(_1611_i2, globalState), true, globalState);
          _out96 = _outcollector25[0];
          _out97 = _outcollector25[1];
          _out98 = _outcollector25[2];
          _out99 = _outcollector25[3];
          _1622_v80 = _out96;
          _1623_v81 = _out97;
          _1624_v82 = _out98;
          _1625_v83 = _out99;
          let _1626_v84;
          _1626_v84 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1612_v70).length),_1614_v72);
          _1624_v82 = (_dafny.ZERO).minus(_module.__default.fm0(_1623_v81, _1626_v84, _1537_v0, globalState));
          let _1627_v85;
          _1627_v85 = _module.D7.create_DC20(_dafny.MultiSet.fromElements(_1611_i2, _1624_v82, _1610_v69, new BigNumber(747)));
          let _1628_v86;
          _1628_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1625_v83);
          let _1629_v87;
          _1629_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1625_v83,_1628_v86);
          let _1630_v88;
          let _nw248 = Array((new BigNumber(18)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1630_v88 = _nw248;
          let _1631_v89;
          let _nw249 = new _module.C1();
          _nw249.__ctor(_1625_v83, ((_1627_v85).dtor_cf30).IsSubsetOf(_module.__default.fm28(_dafny.Seq.UnicodeFromString("fkodyaqi"), false, _1610_v69, globalState)), _1629_v87, _1538_v1, _1630_v88);
          _1631_v89 = _nw249;
        }
      }
      let _1632_v90;
      let _nw250 = Array((new BigNumber(3)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1632_v90 = _nw250;
      let _1633_v91;
      _1633_v91 = _dafny.Seq.UnicodeFromString("ckfjn");
      let _1634_v92;
      _1634_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1633_v91);
      let _index219 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1632_v90).length));
      (_1632_v90)[_index219] = _module.__default.fm28((((_1634_v92).contains(true)) ? ((_1634_v92).get(true)) : (_1633_v91)), true, _1610_v69, globalState);
      let _1635_v93;
      _1635_v93 = _dafny.Seq.of(_1610_v69);
      let _1636_v94;
      _1636_v94 = _dafny.Seq.of((_dafny.ZERO).minus((_1635_v93)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1635_v93).length))]), _1610_v69, _1610_v69, _1610_v69);
      let _1637_v95;
      _1637_v95 = _dafny.MultiSet.fromElements(new BigNumber((_1636_v94).length), _1610_v69);
      let _index220 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1632_v90).length));
      (_1632_v90)[_index220] = _1637_v95;
      if (_1537_v0) {
        let _1638_v96;
        _1638_v96 = _dafny.MultiSet.fromElements(true, _1537_v0);
        let _1639_v97;
        _1639_v97 = _dafny.MultiSet.fromElements(_1638_v96, _1638_v96);
        let _1640_v98;
        _1640_v98 = _dafny.Set.fromElements(false);
        let _1641_v99;
        let _nw251 = Array((new BigNumber(16)).toNumber());
        _nw251[(_dafny.ZERO).toNumber()] = _1610_v69;
        _nw251[(_dafny.ONE).toNumber()] = ((_1537_v0) ? (_1610_v69) : (_1610_v69));
        _nw251[(new BigNumber(2)).toNumber()] = _1610_v69;
        _nw251[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(472), _1610_v69);
        _nw251[(new BigNumber(4)).toNumber()] = _1610_v69;
        _nw251[(new BigNumber(5)).toNumber()] = new BigNumber(-742);
        _nw251[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(_1610_v69, _1610_v69);
        _nw251[(new BigNumber(7)).toNumber()] = _1610_v69;
        _nw251[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(-773), (((_1639_v97).contains(_1638_v96)) ? ((_1639_v97).get(_1638_v96)) : (new BigNumber((_1640_v98).length))));
        _nw251[(new BigNumber(9)).toNumber()] = _1610_v69;
        _nw251[(new BigNumber(10)).toNumber()] = new BigNumber((_1638_v96).cardinality());
        _nw251[(new BigNumber(11)).toNumber()] = _1610_v69;
        _nw251[(new BigNumber(12)).toNumber()] = _1610_v69;
        _nw251[(new BigNumber(13)).toNumber()] = new BigNumber(674);
        _nw251[(new BigNumber(14)).toNumber()] = _1610_v69;
        _nw251[(new BigNumber(15)).toNumber()] = _1610_v69;
        _1641_v99 = _nw251;
        _1641_v99 = _1641_v99;
        (globalState).f4 = _module.__default.fm2((_1610_v69).isEqualTo(_1610_v69), globalState);
        _1610_v69 = new BigNumber(-413);
        let _1642_v100;
        _1642_v100 = _dafny.Map.Empty.slice().updateUnsafe((_1538_v1)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1538_v1).length))],_1537_v0);
        _1610_v69 = new BigNumber((_1642_v100).length);
        let _1643_v101;
        _1643_v101 = _module.D3.create_DC8((_1642_v100).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1537_v0)).update(_1537_v0, _1537_v0)));
        let _source25 = _1643_v101;
        if (_source25.is_DC9) {
          let _1644___mcc_h0 = (_source25).cf12;
          let _1645___mcc_h1 = (_source25).cf13;
          let _1646_cf13 = _1645___mcc_h1;
          let _1647_cf12 = _1644___mcc_h0;
          let _1648_v102;
          _1648_v102 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-38)), ((_1649_v68) => function (_1650_i3) {
            return _1649_v68;
          })(_1609_v68)));
          let _index221 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1641_v99).length));
          (_1641_v99)[_index221] = new BigNumber(881);
          let _1651_v103;
          _1651_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1647_cf12,true);
          let _1652_v105;
          _1652_v105 = _dafny.Seq.of(_1651_v103, ((!(_1537_v0)) ? (function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_1651_v103).Keys.Elements) {
              let _1653_v104 = _compr_59;
              if ((_1651_v103).contains(_1653_v104)) {
                _coll59.push([(_1653_v104).multipliedBy(_1647_cf12),_1537_v0]);
              }
            }
            return _coll59;
          }()) : (_dafny.Map.Empty.slice().updateUnsafe(_1646_cf13,_1537_v0))));
          let _index222 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1641_v99).length));
          let _rhs191 = _1648_v102;
          let _rhs192 = _1610_v69;
          let _rhs193 = (_1538_v1)[_module.__default.safeIndex((_1636_v94)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1636_v94).length))], new BigNumber((_1538_v1).length))];
          let _rhs194 = new BigNumber((_1652_v105).length);
          let _lhs102 = _1641_v99;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_1641_v99).length));
          let _lhs104 = globalState;
          _1648_v102 = _rhs191;
          _lhs102[_lhs103] = _rhs192;
          _lhs104.f4 = _rhs193;
          _1610_v69 = _rhs194;
          let _1654_v106;
          let _nw252 = new _module.C1();
          _nw252.__ctor(_dafny.Seq.IsProperPrefixOf(_1635_v93, _1636_v94), true, _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1642_v100), _dafny.Seq.Concat(_1538_v1, _1538_v1), _1632_v90);
          _1654_v106 = _nw252;
          let _1655_v107;
          let _init41 = ((_1656_v99) => function (_1657_i4) {
            return _module.__default.safeDivisionInt(_1657_i4, (_1656_v99)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1656_v99).length))]);
          })(_1641_v99);
          let _nw253 = Array((new BigNumber(14)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw253.length); _i0_41++) {
            _nw253[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1655_v107 = _nw253;
          let _1658_v108;
          _1658_v108 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1646_cf13), _1646_cf13, new BigNumber((_1633_v91).length), new BigNumber((_1633_v91).length), _1610_v69);
          let _nw254 = Array((new BigNumber(7)).toNumber());
          _nw254[(_dafny.ZERO).toNumber()] = ((_dafny.ZERO).minus(_1610_v69)).multipliedBy(new BigNumber((_1636_v94).length));
          _nw254[(_dafny.ONE).toNumber()] = new BigNumber((_1658_v108).length);
          _nw254[(new BigNumber(2)).toNumber()] = _1610_v69;
          _nw254[(new BigNumber(3)).toNumber()] = _1610_v69;
          _nw254[(new BigNumber(4)).toNumber()] = _1646_cf13;
          _nw254[(new BigNumber(5)).toNumber()] = _1646_cf13;
          _nw254[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((_1641_v99)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1641_v99).length))], (_1641_v99)[_module.__default.safeIndex(new BigNumber(382), new BigNumber((_1641_v99).length))]);
          _1655_v107 = _nw254;
          _1610_v69 = _1647_cf12;
        } else if (_source25.is_DC10) {
          let _1659___mcc_h2 = (_source25).cf14;
          let _1660___mcc_h3 = (_source25).cf15;
          let _1661___mcc_h4 = (_source25).cf16;
          let _1662_cf16 = _1661___mcc_h4;
          let _1663_cf15 = _1660___mcc_h3;
          let _1664_cf14 = _1659___mcc_h2;
          _1537_v0 = _1664_cf14;
          let _1665_v109;
          let _nw255 = Array((new BigNumber(23)).toNumber());
          _nw255[(_dafny.ZERO).toNumber()] = _1641_v99;
          _nw255[(_dafny.ONE).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(2)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(3)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(4)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(5)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(6)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(7)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(8)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(9)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(10)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(11)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(12)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(13)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(14)).toNumber()] = ((_module.__default.fm2(_1537_v0, globalState)) ? (_1641_v99) : (_1641_v99));
          _nw255[(new BigNumber(15)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(16)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(17)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(18)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(19)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(20)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(21)).toNumber()] = _1641_v99;
          _nw255[(new BigNumber(22)).toNumber()] = _1641_v99;
          _1665_v109 = _nw255;
          let _1666_v110;
          let _nw256 = Array((new BigNumber(17)).toNumber()).fill(false);
          _1666_v110 = _nw256;
          let _index223 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1666_v110).length));
          (_1666_v110)[_index223] = !((_1610_v69).isLessThan(_1663_cf15));
          let _index224 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1641_v99).length));
          (_1641_v99)[_index224] = _module.__default.safeDivisionInt(_1663_cf15, _1610_v69);
          let _1667_v112;
          _1667_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1663_cf15,new BigNumber((function () {
            let _coll60 = new _dafny.Map();
            for (const _compr_60 of _dafny.IntegerRange(new BigNumber(409), new BigNumber(735))) {
              let _1668_v111 = _compr_60;
              if (((new BigNumber(409)).isLessThanOrEqualTo(_1668_v111)) && ((_1668_v111).isLessThan(new BigNumber(735)))) {
                _coll60.push([_module.__default.safeDivisionInt(_1668_v111, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1664_cf14,_1610_v69)).length)),_1537_v0]);
              }
            }
            return _coll60;
          }()).length));
          let _index225 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1666_v110).length));
          let _index226 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1641_v99).length));
          let _rhs195 = _1665_v109;
          let _rhs196 = _1663_cf15;
          let _rhs197 = (_1610_v69).isLessThanOrEqualTo(_1610_v69);
          let _rhs198 = (new BigNumber((_1667_v112).length)).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_1663_cf15).plus(_1663_cf15))));
          let _rhs199 = ((new BigNumber(704)).minus(_1663_cf15)).minus(_1663_cf15);
          let _lhs105 = _1666_v110;
          let _lhs106 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1666_v110).length));
          let _lhs107 = _1641_v99;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1641_v99).length));
          _1665_v109 = _rhs195;
          _1663_cf15 = _rhs196;
          _lhs105[_lhs106] = _rhs197;
          _lhs107[_lhs108] = _rhs198;
          _1663_cf15 = _rhs199;
          _1640_v98 = _1640_v98;
          let _index227 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_1641_v99).length));
          (_1641_v99)[_index227] = _1610_v69;
        } else if (_source25.is_DC8) {
          let _1669___mcc_h5 = (_source25).cf11;
          let _1670_cf11 = _1669___mcc_h5;
          let _1671_v113;
          let _init42 = ((_1672_v69) => function (_1673_i5) {
            return _module.D0.create_DC2(_1672_v69, true);
          })(_1610_v69);
          let _nw257 = Array((new BigNumber(13)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw257.length); _i0_42++) {
            _nw257[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1671_v113 = _nw257;
          let _1674_v114;
          _1674_v114 = _module.D0.create_DC2(_1610_v69, _1537_v0);
          let _index228 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1671_v113).length));
          (_1671_v113)[_index228] = _1674_v114;
          let _pat_let_tv31 = _1537_v0;
          let _pat_let_tv32 = globalState;
          let _pat_let_tv33 = _1537_v0;
          let _index229 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_1671_v113).length));
          (_1671_v113)[_index229] = function (_pat_let27_0) {
            return function (_1677_dt__update__tmp_h1) {
              return function (_pat_let30_0) {
                return function (_1678_dt__update_hcf2_h0) {
                  return _module.D0.create_DC2((_1677_dt__update__tmp_h1).dtor_cf1, _1678_dt__update_hcf2_h0);
                }(_pat_let30_0);
              }((false) === (_pat_let_tv33));
            }(_pat_let27_0);
          }(function (_pat_let28_0) {
            return function (_1675_dt__update__tmp_h0) {
              return function (_pat_let29_0) {
                return function (_1676_dt__update_hcf1_h0) {
                  return _module.D0.create_DC2(_1676_dt__update_hcf1_h0, (_1675_dt__update__tmp_h0).dtor_cf2);
                }(_pat_let29_0);
              }(new BigNumber((_dafny.Seq.of(_module.__default.fm2(_pat_let_tv31, _pat_let_tv32))).length));
            }(_pat_let28_0);
          }(_1674_v114));
          let _1679_v115;
          let _1680_v116;
          let _1681_v117;
          let _out100;
          let _out101;
          let _out102;
          let _outcollector26 = (_this).m8(globalState);
          _out100 = _outcollector26[0];
          _out101 = _outcollector26[1];
          _out102 = _outcollector26[2];
          _1679_v115 = _out100;
          _1680_v116 = _out101;
          _1681_v117 = _out102;
          _1640_v98 = ((_1640_v98).Intersect(_1640_v98)).Union(_1640_v98);
          let _1682_v118;
          _1682_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1679_v115,_1610_v69);
          let _1683_v119;
          _1683_v119 = _dafny.Seq.of(_1682_v118);
          let _1684_v120;
          _1684_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1610_v69,_1537_v0);
          let _1685_v121;
          _1685_v121 = _module.D8.create_DC23(_1682_v118);
          let _1686_v122;
          let _nw258 = Array((new BigNumber(24)).toNumber());
          _nw258[(_dafny.ZERO).toNumber()] = _1682_v118;
          _nw258[(_dafny.ONE).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(2)).toNumber()] = (_1683_v119)[_module.__default.safeIndex(new BigNumber((_1684_v120).length), new BigNumber((_1683_v119).length))];
          _nw258[(new BigNumber(3)).toNumber()] = (_1682_v118).Merge(_1682_v118);
          _nw258[(new BigNumber(4)).toNumber()] = (_1682_v118).update(_1679_v115, new BigNumber((_1636_v94).length));
          _nw258[(new BigNumber(5)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(6)).toNumber()] = (_1682_v118).update(!(_1537_v0), new BigNumber(191));
          _nw258[(new BigNumber(7)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(8)).toNumber()] = (_1683_v119)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1683_v119).length))];
          _nw258[(new BigNumber(9)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(10)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(11)).toNumber()] = (_1682_v118).Merge(_1682_v118);
          _nw258[(new BigNumber(12)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(13)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(14)).toNumber()] = (_1685_v121).dtor_cf35;
          _nw258[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1679_v115,new BigNumber(-276));
          _nw258[(new BigNumber(16)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(17)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1679_v115,_1610_v69)).update(_1537_v0, new BigNumber(858));
          _nw258[(new BigNumber(18)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(19)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(20)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(21)).toNumber()] = _1682_v118;
          _nw258[(new BigNumber(22)).toNumber()] = (_1682_v118).update(_1679_v115, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-849)), ((_1687_v93) => function (_1688_i6) {
            return new BigNumber((_1687_v93).length);
          })(_1635_v93))).length));
          _nw258[(new BigNumber(23)).toNumber()] = (_1682_v118).Merge(_1682_v118);
          _1686_v122 = _nw258;
          _1686_v122 = _1686_v122;
        } else {
          let _1689___mcc_h6 = (_source25).cf17;
          let _1690_cf17 = _1689___mcc_h6;
          _1610_v69 = _1610_v69;
          (globalState).f4 = _1537_v0;
          _1537_v0 = (new BigNumber((_1633_v91).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ydrbvsno"), _1633_v91)).length));
          let _1691_v123;
          let _nw259 = new _module.C0();
          _nw259.__ctor();
          _1691_v123 = _nw259;
        }
      } else {
        _1609_v68 = _1609_v68;
        let _1692_v124;
        _1692_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1537_v0);
        let _1693_v125;
        _1693_v125 = _module.D3.create_DC11(_module.D3.create_DC8(_1692_v124));
        let _source26 = _1693_v125;
        if (_source26.is_DC9) {
          let _1694___mcc_h7 = (_source26).cf12;
          let _1695___mcc_h8 = (_source26).cf13;
          let _1696_cf13 = _1695___mcc_h8;
          let _1697_cf12 = _1694___mcc_h7;
          let _1698_v126;
          let _nw260 = Array((new BigNumber(22)).toNumber());
          _nw260[(_dafny.ZERO).toNumber()] = _1633_v91;
          _nw260[(_dafny.ONE).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(2)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(3)).toNumber()] = (((_1634_v92).contains(_1537_v0)) ? ((_1634_v92).get(_1537_v0)) : (_1633_v91));
          _nw260[(new BigNumber(4)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("xunrvxhf");
          _nw260[(new BigNumber(6)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(7)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(8)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(9)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(10)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(623)), ((_1699_v68) => function (_1700_i7) {
            return _1699_v68;
          })(_1609_v68));
          _nw260[(new BigNumber(12)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_1633_v91, _module.__default.safeIndex(_1697_cf12, new BigNumber((_1633_v91).length)), _1609_v68);
          _nw260[(new BigNumber(14)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(15)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(16)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("b");
          _nw260[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), ((_1701_v68) => function (_1702_i8) {
            return _1701_v68;
          })(_1609_v68)), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1696_cf13,_1537_v0)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), ((_1703_v68) => function (_1704_i8) {
            return _1703_v68;
          })(_1609_v68))).length)), _1609_v68);
          _nw260[(new BigNumber(19)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(20)).toNumber()] = _1633_v91;
          _nw260[(new BigNumber(21)).toNumber()] = _1633_v91;
          _1698_v126 = _nw260;
          let _1705_v127;
          _1705_v127 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1696_cf13),_1698_v126);
          _1705_v127 = (_1705_v127).update(_1696_cf13, _1698_v126);
          let _1706_v128;
          _1706_v128 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(331),(_1637_v95).IsSubsetOf(_1637_v95));
          _1706_v128 = (_1706_v128).update(_1697_cf12, (_1610_v69).isLessThan(new BigNumber(399)));
          _1610_v69 = _1610_v69;
          let _rhs200 = _1610_v69;
          let _rhs201 = (_1539_v2)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1539_v2).length))];
          _1610_v69 = _rhs200;
          _1538_v1 = _rhs201;
        } else if (_source26.is_DC10) {
          let _1707___mcc_h9 = (_source26).cf14;
          let _1708___mcc_h10 = (_source26).cf15;
          let _1709___mcc_h11 = (_source26).cf16;
          let _1710_cf16 = _1709___mcc_h11;
          let _1711_cf15 = _1708___mcc_h10;
          let _1712_cf14 = _1707___mcc_h9;
          _1711_cf15 = new BigNumber(-590);
          let _1713_v129;
          let _nw261 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _1713_v129 = _nw261;
          let _1714_v130;
          _1714_v130 = _module.D8.create_DC24(_1712_cf14);
          let _rhs202 = (_1714_v130).dtor_cf36;
          let _rhs203 = _1713_v129;
          let _lhs109 = globalState;
          _lhs109.f4 = _rhs202;
          _1713_v129 = _rhs203;
          _1610_v69 = _dafny.ZERO;
          _1633_v91 = _1633_v91;
        } else if (_source26.is_DC8) {
          let _1715___mcc_h12 = (_source26).cf11;
          let _1716_cf11 = _1715___mcc_h12;
          _1610_v69 = (((_1610_v69).isLessThan(_1610_v69)) ? (new BigNumber((_1633_v91).length)) : (new BigNumber(962)));
          let _1717_v131;
          let _nw262 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1717_v131 = _nw262;
          _1717_v131 = _1717_v131;
          let _1718_v132;
          _1718_v132 = _dafny.Map.Empty.slice().updateUnsafe(_1537_v0,_1716_cf11);
          let _1719_v133;
          let _nw263 = new _module.C1();
          _nw263.__ctor(_1537_v0, _1537_v0, _1718_v132, _1538_v1, _1632_v90);
          _1719_v133 = _nw263;
          let _1720_v134;
          let _init43 = function (_1721_i9) {
            return (_1721_i9).plus(new BigNumber(((_module.D4.create_DC12(_dafny.Seq.of(false))).dtor_cf18).length));
          };
          let _nw264 = Array((new BigNumber(24)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw264.length); _i0_43++) {
            _nw264[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1720_v134 = _nw264;
          let _index230 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_1720_v134).length));
          (_1720_v134)[_index230] = (_1636_v94)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1636_v94).length))];
          let _index231 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_1720_v134).length));
          (_1720_v134)[_index231] = (_1610_v69).multipliedBy(_1610_v69);
        } else {
          let _1722___mcc_h13 = (_source26).cf17;
          let _1723_cf17 = _1722___mcc_h13;
          let _1724_v135;
          _1724_v135 = _dafny.Set.fromElements(_1610_v69, _1610_v69);
          let _1725_v136;
          _1725_v136 = _module.D6.create_DC17(_1724_v135);
          let _1726_v137;
          _1726_v137 = _module.D6.create_DC17((_1725_v136).dtor_cf27);
          let _1727_v138;
          _1727_v138 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), ((_1728_v69) => function (_1729_i10) {
            return _1728_v69;
          })(_1610_v69)));
          let _1730_v139;
          _1730_v139 = _dafny.Map.Empty.slice().updateUnsafe(_1726_v137,_dafny.Seq.Concat(_1727_v138, _1727_v138));
          _1730_v139 = (_1730_v139).update(_module.D6.create_DC17(_1724_v135), _dafny.Seq.Concat(_dafny.Seq.of(_1635_v93, _dafny.Seq.of(_1610_v69), _1635_v93), _dafny.Seq.Create(_module.__default.abs(new BigNumber(473)), ((_1731_v69) => function (_1732_i11) {
            return _dafny.Seq.update(_dafny.Seq.of(_1731_v69, _1731_v69), _module.__default.safeIndex((_module.D3.create_DC9(_1731_v69, _1731_v69)).dtor_cf12, new BigNumber((_dafny.Seq.of(_1731_v69, _1731_v69)).length)), (_dafny.ZERO).minus(_1731_v69));
          })(_1610_v69))));
          _1635_v93 = _dafny.Seq.update((_1727_v138)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1727_v138).length))], _module.__default.safeIndex(_1610_v69, new BigNumber(((_1727_v138)[_module.__default.safeIndex(_1610_v69, new BigNumber((_1727_v138).length))]).length)), new BigNumber(-147));
          let _1733_v140;
          let _nw265 = new _module.C0();
          _nw265.__ctor();
          _1733_v140 = _nw265;
          _1610_v69 = new BigNumber(149);
        }
        let _1734_v141;
        let _nw266 = Array((new BigNumber(23)).toNumber()).fill(false);
        _1734_v141 = _nw266;
        let _index232 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length));
        (_1734_v141)[_index232] = false;
        let _index233 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length));
        (_1734_v141)[_index233] = _1537_v0;
        if (((_1734_v141)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length))]) && (!(_1610_v69).isEqualTo((_dafny.ZERO).minus(_1610_v69)))) {
          _1633_v91 = _module.__default.fm29(_1537_v0, _1537_v0, !(_1537_v0), new BigNumber(-320), globalState);
          let _1735_v142;
          _1735_v142 = _dafny.Set.fromElements(_1610_v69);
          let _1736_v143;
          _1736_v143 = _dafny.MultiSet.fromElements(_1537_v0, (_1735_v142).IsProperSubsetOf(_1735_v142), _1537_v0, (_1734_v141)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length))]);
          _1736_v143 = _1736_v143;
          let _1737_v144;
          let _nw267 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _1737_v144 = _nw267;
          _1737_v144 = _1737_v144;
          let _1738_v145;
          _1738_v145 = _dafny.Map.Empty.slice().updateUnsafe(_1610_v69,_1610_v69);
          _1738_v145 = _1738_v145;
          let _1739_v146;
          _1739_v146 = _dafny.MultiSet.fromElements(_1737_v144, _1737_v144, _1737_v144, _1737_v144);
          _1739_v146 = _1739_v146;
        } else {
          _1610_v69 = _1610_v69;
          let _1740_v147;
          _1740_v147 = _dafny.MultiSet.fromElements((_1734_v141)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length))]);
          let _1741_v148;
          _1741_v148 = _dafny.Map.Empty.slice().updateUnsafe(_1609_v68,_1537_v0);
          let _index234 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1632_v90).length));
          (_1632_v90)[_index234] = _dafny.MultiSet.fromElements((((_1740_v147).contains((_1734_v141)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length))])) ? ((_1740_v147).get((_1734_v141)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length))])) : (new BigNumber((_dafny.Seq.of(_1610_v69, new BigNumber((_1741_v148).length))).length))));
          let _index235 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length));
          (_1734_v141)[_index235] = _1537_v0;
          _1610_v69 = _1610_v69;
          (globalState).f4 = (_1537_v0) || ((_1734_v141)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_1734_v141).length))]);
        }
        _1537_v0 = _1537_v0;
      }
      r0 = _1537_v0;
      return r0;
    }
    m8(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Seq.of();
      let _1742_v0;
      _1742_v0 = true;
      let _1743_v1;
      let _nw268 = Array((new BigNumber(6)).toNumber());
      _nw268[(_dafny.ZERO).toNumber()] = !(!(_1742_v0));
      _nw268[(_dafny.ONE).toNumber()] = false;
      _nw268[(new BigNumber(2)).toNumber()] = _1742_v0;
      _nw268[(new BigNumber(3)).toNumber()] = _1742_v0;
      _nw268[(new BigNumber(4)).toNumber()] = _1742_v0;
      _nw268[(new BigNumber(5)).toNumber()] = _1742_v0;
      _1743_v1 = _nw268;
      let _1744_v2;
      let _nw269 = Array((new BigNumber(15)).toNumber());
      _nw269[(_dafny.ZERO).toNumber()] = _1743_v1;
      _nw269[(_dafny.ONE).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(2)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(3)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(4)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(5)).toNumber()] = ((_1742_v0) ? (_1743_v1) : (_1743_v1));
      _nw269[(new BigNumber(6)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(7)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(8)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(9)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(10)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(11)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(12)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(13)).toNumber()] = _1743_v1;
      _nw269[(new BigNumber(14)).toNumber()] = _1743_v1;
      _1744_v2 = _nw269;
      _1744_v2 = _1744_v2;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1743_v1).length))) {
        let _1745_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1745_i0)) && ((_1745_i0).isLessThan(new BigNumber((_1743_v1).length))))) {
          (_1743_v1)[(_1745_i0)] = !(!(_1742_v0));
        }
      }
      let _1746_v3;
      let _init44 = function (_1747_i2) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(73)), function (_1748_i3) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        });
      };
      let _nw270 = Array((new BigNumber(10)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw270.length); _i0_44++) {
        _nw270[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1746_v3 = _nw270;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1746_v3).length))) {
        let _1749_i1 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1749_i1)) && ((_1749_i1).isLessThan(new BigNumber((_1746_v3).length))))) {
          (_1746_v3)[(_1749_i1)] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vaderto"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(411)), function (_1750_i4) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          }));
        }
      }
      let _1751_v4;
      _1751_v4 = _dafny.Seq.UnicodeFromString("gfbewqfba");
      let _1752_v5;
      _1752_v5 = new BigNumber(571);
      let _1753_v6;
      let _nw271 = Array((new BigNumber(8)).toNumber());
      _nw271[(_dafny.ZERO).toNumber()] = new BigNumber(867);
      _nw271[(_dafny.ONE).toNumber()] = _1752_v5;
      _nw271[(new BigNumber(2)).toNumber()] = _1752_v5;
      _nw271[(new BigNumber(3)).toNumber()] = _1752_v5;
      _nw271[(new BigNumber(4)).toNumber()] = _1752_v5;
      _nw271[(new BigNumber(5)).toNumber()] = _1752_v5;
      _nw271[(new BigNumber(6)).toNumber()] = _1752_v5;
      _nw271[(new BigNumber(7)).toNumber()] = _1752_v5;
      _1753_v6 = _nw271;
      let _1754_v7;
      _1754_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1751_v4).length),_1753_v6);
      let _1755_v8;
      _1755_v8 = _dafny.Set.fromElements(new BigNumber((_1754_v7).length));
      let _1756_v10;
      _1756_v10 = _dafny.Seq.of(_1742_v0, _1742_v0);
      let _1757_v11;
      _1757_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1752_v5);
      let _1758_v12;
      _1758_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1756_v10).length),_1757_v11);
      let _1759_v13;
      _1759_v13 = _module.D4.create_DC13(_module.__default.fm2(_1742_v0, globalState), (_1755_v8).IsSubsetOf(function () {
  let _coll61 = new _dafny.Set();
  for (const _compr_61 of _dafny.IntegerRange(new BigNumber(453), new BigNumber(964))) {
    let _1760_v9 = _compr_61;
    if (((new BigNumber(453)).isLessThanOrEqualTo(_1760_v9)) && ((_1760_v9).isLessThan(new BigNumber(964)))) {
      _coll61.add(_module.__default.safeModuloInt(_1760_v9, _1752_v5));
    }
  }
  return _coll61;
}()), _module.__default.fm0(_1742_v0, _1758_v12, _1742_v0, globalState), _module.__default.safeModuloInt(_1752_v5, _1752_v5), true);
      let _source27 = _1759_v13;
      if (_source27.is_DC13) {
        let _1761___mcc_h0 = (_source27).cf19;
        let _1762___mcc_h1 = (_source27).cf20;
        let _1763___mcc_h2 = (_source27).cf21;
        let _1764___mcc_h3 = (_source27).cf22;
        let _1765___mcc_h4 = (_source27).cf23;
        let _1766_cf23 = _1765___mcc_h4;
        let _1767_cf22 = _1764___mcc_h3;
        let _1768_cf21 = _1763___mcc_h2;
        let _1769_cf20 = _1762___mcc_h1;
        let _1770_cf19 = _1761___mcc_h0;
        let _index236 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1753_v6).length));
        (_1753_v6)[_index236] = new BigNumber(608);
        let _1771_v14;
        _1771_v14 = new _dafny.CodePoint('e'.codePointAt(0));
        let _1772_v15;
        _1772_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1768_cf21,_1771_v14);
        let _index237 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1753_v6).length));
        (_1753_v6)[_index237] = (_1752_v5).plus((_dafny.ZERO).minus(new BigNumber((_module.__default.fm32(_1770_cf19, _1771_v14, new BigNumber((_1772_v15).length), new BigNumber(322), globalState)).length)));
        let _1773_v16;
        _1773_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1770_cf19);
        let _1774_v17;
        _1774_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1773_v16);
        let _1775_v18;
        let _init45 = ((_1776_cf22) => function (_1777_i5) {
          return _dafny.MultiSet.fromElements(_1776_cf22);
        })(_1767_cf22);
        let _nw272 = Array((new BigNumber(14)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw272.length); _i0_45++) {
          _nw272[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1775_v18 = _nw272;
        let _1778_v19;
        let _nw273 = new _module.C1();
        _nw273.__ctor(_1766_cf23, _module.__default.fm2(_1770_cf19, globalState), _1774_v17, _1756_v10, _1775_v18);
        _1778_v19 = _nw273;
        let _1779_v20;
        let _init46 = ((_1780_v14) => function (_1781_i6) {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),_1780_v14);
        })(_1771_v14);
        let _nw274 = Array((new BigNumber(28)).toNumber());
        for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw274.length); _i0_46++) {
          _nw274[_i0_46] = _init46(new BigNumber(_i0_46));
        }
        _1779_v20 = _nw274;
        let _1782_v21;
        _1782_v21 = _dafny.Seq.of(_1778_v19, _1778_v19, _1778_v19);
        let _index238 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1753_v6).length));
        let _rhs204 = _1768_cf21;
        let _rhs205 = ((((_1766_cf23) ? (_1770_cf19) : (_1742_v0))) ? (((_1769_cf20) ? (_1778_v19) : (_1778_v19))) : (((_1766_cf23) ? (_1778_v19) : ((_1782_v21)[_module.__default.safeIndex(_1768_cf21, new BigNumber((_1782_v21).length))]))));
        let _rhs206 = !(_module.__default.fm2(_1742_v0, globalState));
        let _rhs207 = _1779_v20;
        let _lhs110 = _1753_v6;
        let _lhs111 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_1753_v6).length));
        _lhs110[_lhs111] = _rhs204;
        _1778_v19 = _rhs205;
        _1770_cf19 = _rhs206;
        _1779_v20 = _rhs207;
        let _1783_v22;
        _1783_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1743_v1,_1770_cf19);
        let _1784_v23;
        let _nw275 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1784_v23 = _nw275;
        let _1785_v24;
        let _nw276 = Array((new BigNumber(18)).toNumber());
        _nw276[(_dafny.ZERO).toNumber()] = _1784_v23;
        _nw276[(_dafny.ONE).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(2)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(3)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(4)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(5)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(6)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(7)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(8)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(9)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(10)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(11)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(12)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(13)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(14)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(15)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(16)).toNumber()] = _1784_v23;
        _nw276[(new BigNumber(17)).toNumber()] = _1784_v23;
        _1785_v24 = _nw276;
        let _1786_v25;
        _1786_v25 = _1785_v24;
        let _rhs208 = _1783_v22;
        let _rhs209 = _dafny.Seq.Concat(_1751_v4, _1751_v4);
        let _rhs210 = (_1786_v25);
        _1783_v22 = _rhs208;
        _1751_v4 = _rhs209;
        _1785_v24 = _rhs210;
        let _1787_v26;
        _1787_v26 = _1778_v19.f12;
        let _1788_v27;
        let _nw277 = new _module.C3();
        _nw277.__ctor(true, (_1787_v26), (_1755_v8).contains(new BigNumber(((_1778_v19).f14).length)), _1774_v17, (_1778_v19).f14);
        _1788_v27 = _nw277;
        let _1789_v28;
        _1789_v28 = _module.D14.create_DC40(_1788_v27);
        _1788_v27 = (_1789_v28).dtor_cf63;
      } else if (_source27.is_DC14) {
        let _1790___mcc_h5 = (_source27).cf24;
        let _1791_cf24 = _1790___mcc_h5;
        let _1792_v29;
        _1792_v29 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1793_v30;
        _1793_v30 = _dafny.Seq.of(_1751_v4, _1751_v4, _1751_v4);
        let _rhs211 = false;
        let _rhs212 = ((!_dafny.Seq.contains(_1751_v4, _1792_v29)) ? (_1751_v4) : (_dafny.Seq.Concat((_1793_v30)[_module.__default.safeIndex(new BigNumber(536), new BigNumber((_1793_v30).length))], _dafny.Seq.UnicodeFromString("kmbesbw"))));
        _1742_v0 = _rhs211;
        _1751_v4 = _rhs212;
        let _1794_v31;
        _1794_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1791_cf24);
        let _1795_v32;
        _1795_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1791_cf24,_1794_v31);
        let _1796_v33;
        let _nw278 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1796_v33 = _nw278;
        let _1797_v34;
        let _nw279 = new _module.C1();
        _nw279.__ctor(_1791_cf24, _1791_cf24, (_1795_v32).Merge(_1795_v32), _module.__default.fm32(_1742_v0, new _dafny.CodePoint('o'.codePointAt(0)), _1752_v5, _1752_v5, globalState), _1796_v33);
        _1797_v34 = _nw279;
        r2 = _dafny.Seq.of((_1797_v34).f17, _1742_v0);
        let _1798_v35;
        let _nw280 = new _module.C0();
        _nw280.__ctor();
        _1798_v35 = _nw280;
      } else {
        let _1799___mcc_h6 = (_source27).cf18;
        let _1800_cf18 = _1799___mcc_h6;
        _1742_v0 = _1742_v0;
        let _1801_v36;
        _1801_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1752_v5,_1752_v5);
        _1752_v5 = (((_1801_v36).contains(new BigNumber(723))) ? ((_1801_v36).get(new BigNumber(723))) : ((_dafny.ZERO).minus(_1752_v5)));
        _1752_v5 = ((_dafny.ZERO).minus(_module.__default.fm0(_1742_v0, _1758_v12, _1742_v0, globalState))).minus(_1752_v5);
        let _1802_v37;
        _1802_v37 = _dafny.MultiSet.fromElements(_1742_v0);
        _1757_v11 = (_1757_v11).Merge(_module.__default.fm45(false, new BigNumber((_1802_v37).cardinality()), (_dafny.ZERO).minus((_dafny.ZERO).minus(_1752_v5)), globalState));
      }
      _1742_v0 = _1742_v0;
      let _1803_v38;
      _1803_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1752_v5,_1742_v0);
      let _1804_v39;
      _1804_v39 = _dafny.Seq.of(new BigNumber(403));
      let _1805_v40;
      _1805_v40 = _module.D7.create_DC21((_1803_v38).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1752_v5,_1742_v0)), _module.__default.fm2(_1742_v0, globalState), !((_1752_v5).isLessThanOrEqualTo((_1804_v39)[_module.__default.safeIndex(_1752_v5, new BigNumber((_1804_v39).length))])));
      let _source28 = _1805_v40;
      if (_source28.is_DC21) {
        let _1806___mcc_h7 = (_source28).cf31;
        let _1807___mcc_h8 = (_source28).cf32;
        let _1808___mcc_h9 = (_source28).cf33;
        let _1809_cf33 = _1808___mcc_h9;
        let _1810_cf32 = _1807___mcc_h8;
        let _1811_cf31 = _1806___mcc_h7;
        if (_1742_v0) {
          _1742_v0 = _module.__default.fm2(_1809_cf33, globalState);
          let _1812_v41;
          _1812_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1752_v5,_1751_v4);
          let _1813_v42;
          _1813_v42 = _dafny.MultiSet.fromElements((((_1812_v41).contains(new BigNumber(21))) ? ((_1812_v41).get(new BigNumber(21))) : (_1751_v4)));
          let _1814_v43;
          _1814_v43 = _dafny.MultiSet.fromElements(_1752_v5, _1752_v5);
          let _1815_v44;
          _1815_v44 = _dafny.Map.Empty.slice().updateUnsafe(true,_1809_cf33);
          let _1816_v45;
          _1816_v45 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1817_v46;
          let _nw281 = Array((new BigNumber(29)).toNumber());
          _nw281[(_dafny.ZERO).toNumber()] = _1814_v43;
          _nw281[(_dafny.ONE).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(2)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(3)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(4)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(5)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(6)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(7)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(8)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(9)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(10)).toNumber()] = ((_1814_v43).update(new BigNumber((_1803_v38).length), _module.__default.abs(new BigNumber((_1815_v44).length)))).update(_1752_v5, _module.__default.abs(new BigNumber(-474)));
          _nw281[(new BigNumber(11)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(12)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_1814_v43).cardinality()));
          _nw281[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_1752_v5)).length));
          _nw281[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements(_1752_v5, _1752_v5);
          _nw281[(new BigNumber(16)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(17)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(18)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(19)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(20)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(21)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(22)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(23)).toNumber()] = _module.__default.fm56(_1816_v45, globalState);
          _nw281[(new BigNumber(24)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(25)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(26)).toNumber()] = _dafny.MultiSet.fromElements(_1752_v5);
          _nw281[(new BigNumber(27)).toNumber()] = _1814_v43;
          _nw281[(new BigNumber(28)).toNumber()] = _1814_v43;
          _1817_v46 = _nw281;
          let _1818_v47;
          let _nw282 = new _module.C1();
          _nw282.__ctor(_1809_cf33, _1742_v0, _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_dafny.Map.Empty.slice().updateUnsafe(_1809_cf33,_1810_cf32)), _dafny.Seq.of(_1809_cf33), _1817_v46);
          _1818_v47 = _nw282;
          let _1819_v48;
          _1819_v48 = _dafny.Map.Empty.slice().updateUnsafe(_1818_v47,_1809_cf33);
          (globalState).f4 = (new BigNumber(((_1813_v42).Union(_1813_v42)).cardinality())).isEqualTo(new BigNumber((_1819_v48).length));
          let _index239 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index239] = _1810_cf32;
          let _index240 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index240] = (_1818_v47).f17;
          let _rhs213 = ((_1752_v5).multipliedBy(_1752_v5)).multipliedBy(_1752_v5);
          let _rhs214 = (_1752_v5).minus(new BigNumber((_dafny.Seq.Concat(_1756_v10, _1756_v10)).length));
          _1752_v5 = _rhs213;
          _1752_v5 = _rhs214;
          let _index241 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1746_v3).length));
          (_1746_v3)[_index241] = _1751_v4;
          let _index242 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_1746_v3).length));
          (_1746_v3)[_index242] = _1751_v4;
        } else {
          let _1820_v49;
          _1820_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm58(new BigNumber((_1756_v10).length), globalState)).length),new BigNumber((_1804_v39).length));
          _1820_v49 = (_1820_v49).update(_1752_v5, _1752_v5);
          let _1821_v50;
          _1821_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1751_v4,_module.__default.fm0(_1810_cf32, _dafny.Map.Empty.slice().updateUnsafe((((_1820_v49).contains(_1752_v5)) ? ((_1820_v49).get(_1752_v5)) : (_1752_v5)),_1757_v11), false, globalState));
          _1821_v50 = (_1821_v50).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(773)), function (_1822_i7) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }), (_dafny.ZERO).minus(new BigNumber(-880)));
          let _1823_v51;
          _1823_v51 = _module.D2.create_DC5(new BigNumber((_1751_v4).length), ((_1809_cf33) ? (_1809_cf33) : (_1742_v0)), _1809_cf33);
          _1823_v51 = _1823_v51;
          let _init47 = ((_1824_v10, _1825_v39) => function (_1826_i8) {
            return (_1824_v10)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1825_v39)).cardinality()), new BigNumber((_1824_v10).length))];
          })(_1756_v10, _1804_v39);
          let _nw283 = Array((new BigNumber(3)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw283.length); _i0_47++) {
            _nw283[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1743_v1 = _nw283;
          _1752_v5 = _1752_v5;
        }
        let _1827_v52;
        _1827_v52 = _dafny.MultiSet.fromElements(_1752_v5, _1752_v5);
        let _1828_v53;
        _1828_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1752_v5,_1827_v52);
        _1757_v11 = (_1757_v11).update((_1752_v5).isLessThanOrEqualTo(_1752_v5), _module.__default.safeDivisionInt(_1752_v5, new BigNumber((_1828_v53).length)));
        r0 = _1809_cf33;
        if (_1809_cf33) {
          _1752_v5 = _1752_v5;
          let _1829_v54;
          _1829_v54 = _dafny.Map.Empty.slice().updateUnsafe(!(_1742_v0),_1809_cf33);
          let _1830_v55;
          _1830_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1809_cf33,_1829_v54);
          let _1831_v56;
          let _nw284 = Array((new BigNumber(4)).toNumber());
          _nw284[(_dafny.ZERO).toNumber()] = _1827_v52;
          _nw284[(_dafny.ONE).toNumber()] = _1827_v52;
          _nw284[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_1752_v5, _1752_v5, (_dafny.ZERO).minus(_1752_v5));
          _nw284[(new BigNumber(3)).toNumber()] = _1827_v52;
          _1831_v56 = _nw284;
          let _1832_v57;
          let _nw285 = new _module.C1();
          _nw285.__ctor(_1742_v0, _1742_v0, _1830_v55, _1756_v10, _1831_v56);
          _1832_v57 = _nw285;
          let _1833_v58;
          _1833_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1757_v11).length),_1832_v57);
          let _1834_v59;
          _1834_v59 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("lfqfoua"));
          let _1835_v60;
          _1835_v60 = new _dafny.CodePoint('w'.codePointAt(0));
          (globalState).f4 = (!_dafny.areEqual(_1804_v39, _dafny.Seq.of(_1752_v5, new BigNumber((_1833_v58).length)))) || ((_1752_v5).isLessThanOrEqualTo((((_1834_v59).contains(_dafny.Seq.update(_1751_v4, _module.__default.safeIndex(_1752_v5, new BigNumber((_1751_v4).length)), _1835_v60))) ? ((_1834_v59).get(_dafny.Seq.update(_1751_v4, _module.__default.safeIndex(_1752_v5, new BigNumber((_1751_v4).length)), _1835_v60))) : (_1752_v5))));
          let _1836_v61;
          _1836_v61 = _module.D4.create_DC14(_1809_cf33);
          _1742_v0 = (_1836_v61).dtor_cf24;
          let _index243 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index243] = (_1742_v0) || (_1810_cf32);
          let _1837_v62;
          _1837_v62 = _dafny.Set.fromElements(_1835_v60);
          let _index244 = _module.__default.safeIndex(new BigNumber(600), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index244] = _module.__default.fm2((_1837_v62).IsSubsetOf(_1837_v62), globalState);
          let _1838_v63;
          _1838_v63 = _module.D4.create_DC12(_1756_v10);
          r0 = (_module.__default.fm44(_1838_v63, true, globalState)).dtor_cf6;
        } else {
          let _1839_v64;
          _1839_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1810_cf32,_1810_cf32);
          _1839_v64 = _module.__default.fm59(_1827_v52, _1810_cf32, ((_1810_cf32) ? (_1809_cf33) : (_1810_cf32)), ((_1810_cf32) ? (_1752_v5) : (new BigNumber((function () {
            let _coll62 = new _dafny.Set();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(835), new BigNumber(884))) {
              let _1840_v65 = _compr_62;
              if (((new BigNumber(835)).isLessThanOrEqualTo(_1840_v65)) && ((_1840_v65).isLessThan(new BigNumber(884)))) {
                _coll62.add(_module.__default.safeDivisionInt(_1840_v65, _1752_v5));
              }
            }
            return _coll62;
          }()).length))), globalState);
          (globalState).f4 = (_1839_v64).equals(_dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1742_v0));
          let _index245 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index245] = _1809_cf33;
          let _1841_v66;
          _1841_v66 = _dafny.Seq.of(_1755_v8, _1755_v8);
          let _1842_v67;
          _1842_v67 = _dafny.Seq.of(_1841_v66, _1841_v66, _1841_v66, _1841_v66, _1841_v66);
          let _index246 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index246] = _dafny.Seq.IsProperPrefixOf((_1842_v67)[_module.__default.safeIndex(new BigNumber(916), new BigNumber((_1842_v67).length))], _1841_v66);
          _1756_v10 = _dafny.Seq.of((_1756_v10)[_module.__default.safeIndex(_1752_v5, new BigNumber((_1756_v10).length))]);
          let _1843_v68;
          _1843_v68 = _dafny.Set.fromElements(_1809_cf33, true, _1810_cf32);
          let _1844_v69;
          _1844_v69 = _module.D7.create_DC20(_1827_v52);
          _1843_v68 = _dafny.Set.fromElements((_1743_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_1743_v1).length))], (_1827_v52).IsSubsetOf((_1844_v69).dtor_cf30), (_1756_v10)[_module.__default.safeIndex(new BigNumber((_1756_v10).length), new BigNumber((_1756_v10).length))]);
        }
      } else if (_source28.is_DC22) {
        let _1845___mcc_h10 = (_source28).cf34;
        let _1846_cf34 = _1845___mcc_h10;
        _1846_cf34 = _1846_cf34;
        let _1847_v70;
        _1847_v70 = new _dafny.CodePoint('m'.codePointAt(0));
        let _1848_v71;
        _1848_v71 = _dafny.Map.Empty.slice().updateUnsafe(((!(_1742_v0)) ? (_1847_v70) : (_1847_v70)),_1743_v1);
        _1848_v71 = (_1848_v71).update((_1751_v4)[_module.__default.safeIndex(_1752_v5, new BigNumber((_1751_v4).length))], _1743_v1);
        let _1849_v72;
        _1849_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1846_cf34,_dafny.Map.Empty.slice().updateUnsafe(_1846_cf34,_1742_v0));
        let _1850_v73;
        _1850_v73 = _dafny.MultiSet.fromElements(_1846_cf34, _1846_cf34);
        let _1851_v74;
        _1851_v74 = _dafny.MultiSet.fromElements(new BigNumber((_1850_v73).cardinality()));
        let _1852_v75;
        _1852_v75 = _dafny.Set.fromElements(_1846_cf34);
        let _1853_v76;
        let _nw286 = Array((new BigNumber(21)).toNumber());
        _nw286[(_dafny.ZERO).toNumber()] = _1851_v74;
        _nw286[(_dafny.ONE).toNumber()] = (_1851_v74).update(_1752_v5, _module.__default.abs(new BigNumber(269)));
        _nw286[(new BigNumber(2)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(3)).toNumber()] = (_1851_v74).update(new BigNumber((_1852_v75).length), _module.__default.abs(_1752_v5));
        _nw286[(new BigNumber(4)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(5)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(6)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(7)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(8)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(9)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(10)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(11)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_1752_v5, _1752_v5);
        _nw286[(new BigNumber(13)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(14)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(15)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(16)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(17)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(18)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(19)).toNumber()] = _1851_v74;
        _nw286[(new BigNumber(20)).toNumber()] = _1851_v74;
        _1853_v76 = _nw286;
        let _1854_v77;
        let _nw287 = new _module.C1();
        _nw287.__ctor(_1742_v0, _1742_v0, _1849_v72, _dafny.Seq.of(_1846_cf34), _1853_v76);
        _1854_v77 = _nw287;
        let _pat_let_tv34 = _1752_v5;
        let _source29 = function (_pat_let31_0) {
          return function (_1855_dt__update__tmp_h0) {
            return function (_pat_let32_0) {
              return function (_1856_dt__update_hcf49_h0) {
                return _module.D11.create_DC33((_1855_dt__update__tmp_h0).dtor_cf48, _1856_dt__update_hcf49_h0, (_1855_dt__update__tmp_h0).dtor_cf50, (_1855_dt__update__tmp_h0).dtor_cf51, (_1855_dt__update__tmp_h0).dtor_cf52);
              }(_pat_let32_0);
            }(_pat_let_tv34);
          }(_pat_let31_0);
        }(_module.D11.create_DC33(_1846_cf34, _1752_v5, !(true), _1854_v77, (_1854_v77).f17));
        if (_source29.is_DC31) {
          (_1854_v77).f18 = _1742_v0;
          let _index247 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1753_v6).length));
          (_1753_v6)[_index247] = _1752_v5;
          let _index248 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_1753_v6).length));
          (_1753_v6)[_index248] = (_dafny.ZERO).minus(_1752_v5);
          _1752_v5 = (_1753_v6)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_1753_v6).length))];
          let _index249 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index249] = ((_1742_v0) ? (true) : ((_1854_v77).f17));
          let _index250 = _module.__default.safeIndex(new BigNumber(17), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index250] = !(false);
        } else if (_source29.is_DC32) {
          let _1857___mcc_h12 = (_source29).cf46;
          let _1858___mcc_h13 = (_source29).cf47;
          let _1859_cf47 = _1858___mcc_h13;
          let _1860_cf46 = _1857___mcc_h12;
          _1752_v5 = (_dafny.ZERO).minus(_1752_v5);
          _1753_v6 = _1753_v6;
          let _index251 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1744_v2).length));
          (_1744_v2)[_index251] = _1743_v1;
          let _index252 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_1744_v2).length));
          (_1744_v2)[_index252] = _1743_v1;
          (globalState).f4 = (_1854_v77).f17;
        } else if (_source29.is_DC33) {
          let _1861___mcc_h14 = (_source29).cf48;
          let _1862___mcc_h15 = (_source29).cf49;
          let _1863___mcc_h16 = (_source29).cf50;
          let _1864___mcc_h17 = (_source29).cf51;
          let _1865___mcc_h18 = (_source29).cf52;
          let _1866_cf52 = _1865___mcc_h18;
          let _1867_cf51 = _1864___mcc_h17;
          let _1868_cf50 = _1863___mcc_h16;
          let _1869_cf49 = _1862___mcc_h15;
          let _1870_cf48 = _1861___mcc_h14;
          let _1871_v78;
          let _nw288 = new _module.C3();
          _nw288.__ctor(_1870_cf48, _1853_v76, (_1854_v77).f17, _module.__default.fm60(_1869_cf49, globalState), _1756_v10);
          _1871_v78 = _nw288;
          let _1872_v79;
          _1872_v79 = _module.D14.create_DC40(_1871_v78);
          let _1873_v80;
          _1873_v80 = _module.D14.create_DC43((((_1867_cf51).f17) ? (_module.D14.create_DC42()) : (_1872_v79)));
          let _rhs215 = _1854_v77.f18;
          let _rhs216 = _1752_v5;
          let _rhs217 = _1873_v80;
          let _rhs218 = _1752_v5;
          let _lhs112 = _1854_v77;
          _lhs112.f18 = _rhs215;
          _1752_v5 = _rhs216;
          _1873_v80 = _rhs217;
          _1752_v5 = _rhs218;
          _1752_v5 = _module.__default.safeDivisionInt(_1869_cf49, _module.__default.safeModuloInt(new BigNumber(428), new BigNumber((function () {
            let _coll63 = new _dafny.Map();
            for (const _compr_63 of _dafny.IntegerRange(new BigNumber(-568), new BigNumber(-612))) {
              let _1874_v81 = _compr_63;
              if (((new BigNumber(-568)).isLessThanOrEqualTo(_1874_v81)) && ((_1874_v81).isLessThan(new BigNumber(-612)))) {
                _coll63.push([_module.__default.safeModuloInt(_1874_v81, _1869_cf49),new BigNumber((_dafny.Seq.of(_1867_cf51.f18)).length)]);
              }
            }
            return _coll63;
          }()).length)));
          let _index253 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index253] = !(_1854_v77.f18);
          let _1875_v82;
          _1875_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1870_cf48,(_1867_cf51).f17);
          let _1876_v83;
          _1876_v83 = _1850_v73;
          let _1877_v84;
          let _nw289 = new _module.C2();
          _nw289.__ctor(_dafny.Seq.of(_1869_cf49), _1742_v0, ((_1871_v78).f13).update(_1867_cf51.f18, _1875_v82), _module.__default.fm43(_dafny.MultiSet.FromArray((_1871_v78).f14), _module.D7.create_DC22((_1867_cf51).f17), globalState), _1871_v78.f12);
          _1877_v84 = _nw289;
          let _1878_v85;
          _1878_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1866_cf52,_1877_v84);
          let _index254 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1743_v1).length));
          let _rhs219 = (_1867_cf51).f17;
          let _rhs220 = (_module.__default.safeDivisionInt(new BigNumber((_1878_v85).length), _1869_cf49)).isLessThanOrEqualTo(_1869_cf49);
          let _rhs221 = _1751_v4;
          let _rhs222 = _1743_v1;
          let _rhs223 = _1743_v1;
          let _lhs113 = _1743_v1;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_1743_v1).length));
          let _lhs115 = globalState;
          _lhs113[_lhs114] = _rhs219;
          _lhs115.f4 = _rhs220;
          _1751_v4 = _rhs221;
          _1743_v1 = _rhs222;
          _1743_v1 = _rhs223;
          _1752_v5 = _1869_cf49;
        } else {
          let _1879___mcc_h19 = (_source29).cf45;
          let _1880_cf45 = _1879___mcc_h19;
          let _1881_v86;
          _1881_v86 = _dafny.Map.Empty.slice().updateUnsafe(((!(_1854_v77.f18)) ? (_dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1854_v77.f18)) : (_dafny.Map.Empty.slice().updateUnsafe((_1854_v77).f17,!((_1854_v77).f17)))),((_1854_v77.f18) ? (_1751_v4) : (_dafny.Seq.UnicodeFromString("qfjvynxmm"))));
          let _index255 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1753_v6).length));
          (_1753_v6)[_index255] = _1752_v5;
          let _index256 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1753_v6).length));
          let _rhs224 = _module.__default.fm61(globalState);
          let _rhs225 = ((new BigNumber(613)).multipliedBy(_1752_v5)).minus(_1752_v5);
          let _lhs116 = _1753_v6;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1753_v6).length));
          _1881_v86 = _rhs224;
          _lhs116[_lhs117] = _rhs225;
          let _1882_v87;
          _1882_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1743_v1,(_1854_v77).f17);
          _1742_v0 = !(_1882_v87).equals((_1882_v87).update(_1743_v1, false));
          let _1883_v88;
          let _1884_v89;
          let _1885_v90;
          let _1886_v91;
          let _out103;
          let _out104;
          let _out105;
          let _out106;
          let _outcollector27 = (_1854_v77).m3(globalState);
          _out103 = _outcollector27[0];
          _out104 = _outcollector27[1];
          _out105 = _outcollector27[2];
          _out106 = _outcollector27[3];
          _1883_v88 = _out103;
          _1884_v89 = _out104;
          _1885_v90 = _out105;
          _1886_v91 = _out106;
          let _index257 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1753_v6).length));
          (_1753_v6)[_index257] = (_1753_v6)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_1753_v6).length))];
        }
        _1851_v74 = _1851_v74;
      } else {
        let _1887___mcc_h11 = (_source28).cf30;
        let _1888_cf30 = _1887___mcc_h11;
        let _1889_v92;
        let _init48 = function (_1890_i9) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        };
        let _nw290 = Array((new BigNumber(16)).toNumber());
        for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw290.length); _i0_48++) {
          _nw290[_i0_48] = _init48(new BigNumber(_i0_48));
        }
        _1889_v92 = _nw290;
        let _1891_v93;
        _1891_v93 = _dafny.Set.fromElements(_1889_v92);
        let _rhs226 = _dafny.Seq.UnicodeFromString("ahyfwsh");
        let _rhs227 = _module.__default.fm0(((_1742_v0) ? (_1742_v0) : (false)), (_1758_v12).update(new BigNumber((_1891_v93).length), _1757_v11), _1742_v0, globalState);
        _1751_v4 = _rhs226;
        _1752_v5 = _rhs227;
        if (!(_1742_v0)) {
          _1752_v5 = _1752_v5;
          let _index258 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          (_1753_v6)[_index258] = new BigNumber(958);
          let _index259 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1753_v6).length));
          (_1753_v6)[_index259] = _1752_v5;
          let _1892_v94;
          _1892_v94 = _module.D8.create_DC26(_1742_v0);
          let _1893_v95;
          _1893_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1743_v1,_1892_v94);
          let _1894_v96;
          _1894_v96 = _module.D12.create_DC34(_1893_v95);
          let _1895_v97;
          _1895_v97 = _dafny.MultiSet.fromElements(_module.D12.create_DC38(_1894_v96));
          let _index260 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          let _index261 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1753_v6).length));
          let _rhs228 = _1752_v5;
          let _rhs229 = ((_1895_v97).Difference(_1895_v97)).IsSubsetOf(_1895_v97);
          let _rhs230 = new BigNumber((_module.__default.fm27(globalState)).length);
          let _rhs231 = ((_1752_v5).minus((_dafny.ZERO).minus(_1752_v5))).multipliedBy(_1752_v5);
          let _rhs232 = _1742_v0;
          let _lhs118 = _1753_v6;
          let _lhs119 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          let _lhs120 = globalState;
          let _lhs121 = _1753_v6;
          let _lhs122 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1753_v6).length));
          let _lhs123 = globalState;
          _lhs118[_lhs119] = _rhs228;
          _lhs120.f4 = _rhs229;
          _1752_v5 = _rhs230;
          _lhs121[_lhs122] = _rhs231;
          _lhs123.f4 = _rhs232;
          let _index262 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1743_v1).length));
          (_1743_v1)[_index262] = true;
          let _1896_v98;
          _1896_v98 = _dafny.Map.Empty.slice().updateUnsafe((_1753_v6)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length))],_1752_v5);
          let _1897_v99;
          _1897_v99 = new _dafny.CodePoint('j'.codePointAt(0));
          let _index263 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1743_v1).length));
          let _index264 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          let _rhs233 = _1742_v0;
          let _rhs234 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(658)), function (_1898_i10) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }), _module.__default.safeIndex((((_1896_v98).contains((_1753_v6)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length))])) ? ((_1896_v98).get((_1753_v6)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length))])) : ((_1753_v6)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length))])), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(658)), function (_1899_i10) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          })).length)), _1897_v99), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("adwid"), _dafny.Seq.update(_1751_v4, _module.__default.safeIndex((_1753_v6)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length))], new BigNumber((_1751_v4).length)), _1897_v99)))).length);
          let _lhs124 = _1743_v1;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1743_v1).length));
          let _lhs126 = _1753_v6;
          let _lhs127 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          _lhs124[_lhs125] = _rhs233;
          _lhs126[_lhs127] = _rhs234;
          let _1900_v100;
          let _nw291 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
          _1900_v100 = _nw291;
          let _1901_v101;
          _1901_v101 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_module.__default.fm0(false, _1758_v12, _1742_v0, globalState), _1752_v5),_1900_v100);
          _1901_v101 = (_1901_v101).update((_1753_v6)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length))], _1900_v100);
          let _1902_v102;
          let _nw292 = Array((new BigNumber(2)).toNumber()).fill([]);
          _1902_v102 = _nw292;
          let _1903_v103;
          _1903_v103 = _1902_v102;
          let _1904_v104;
          let _nw293 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _1904_v104 = _nw293;
          let _1905_v105;
          _1905_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1751_v4,new BigNumber((_1751_v4).length));
          let _index265 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          let _index266 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1743_v1).length));
          let _index267 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          let _rhs235 = _1903_v103;
          let _rhs236 = ((((_1905_v105).contains(_dafny.Seq.update(_1751_v4, _module.__default.safeIndex(_1752_v5, new BigNumber((_1751_v4).length)), _1897_v99))) ? ((_1905_v105).get(_dafny.Seq.update(_1751_v4, _module.__default.safeIndex(_1752_v5, new BigNumber((_1751_v4).length)), _1897_v99))) : (_module.__default.fm0((_1743_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1743_v1).length))], _1758_v12, (_1756_v10)[_module.__default.safeIndex(_module.__default.fm0(_1742_v0, _1758_v12, false, globalState), new BigNumber((_1756_v10).length))], globalState)))).plus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_1752_v5, new BigNumber((_1757_v11).length))));
          let _rhs237 = (_1743_v1)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_1743_v1).length))];
          let _rhs238 = _1904_v104;
          let _rhs239 = (_1804_v39)[_module.__default.safeIndex(_1752_v5, new BigNumber((_1804_v39).length))];
          let _lhs128 = _1753_v6;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          let _lhs130 = _1743_v1;
          let _lhs131 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1743_v1).length));
          let _lhs132 = _1753_v6;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1753_v6).length));
          _1903_v103 = _rhs235;
          _lhs128[_lhs129] = _rhs236;
          _lhs130[_lhs131] = _rhs237;
          _1904_v104 = _rhs238;
          _lhs132[_lhs133] = _rhs239;
        } else {
          _1752_v5 = _1752_v5;
          (globalState).f4 = false;
          let _1906_v106;
          _1906_v106 = _dafny.Map.Empty.slice().updateUnsafe(_1752_v5,_1804_v39);
          let _1907_v107;
          _1907_v107 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1906_v106);
          let _index268 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1753_v6).length));
          (_1753_v6)[_index268] = new BigNumber(((((_1907_v107).contains(_1742_v0)) ? ((_1907_v107).get(_1742_v0)) : (_1906_v106))).length);
          let _index269 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1753_v6).length));
          let _rhs240 = _1742_v0;
          let _rhs241 = _module.__default.fm0(_dafny.Seq.IsProperPrefixOf(_1751_v4, _1751_v4), _1758_v12, _1742_v0, globalState);
          let _rhs242 = (_1752_v5).isLessThan(((_1742_v0) ? ((_dafny.ZERO).minus(_1752_v5)) : (_1752_v5)));
          let _lhs134 = _1753_v6;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1753_v6).length));
          let _lhs136 = globalState;
          r0 = _rhs240;
          _lhs134[_lhs135] = _rhs241;
          _lhs136.f4 = _rhs242;
          _1752_v5 = (_1753_v6)[_module.__default.safeIndex(new BigNumber(272), new BigNumber((_1753_v6).length))];
          _1752_v5 = (_1753_v6)[_module.__default.safeIndex(new BigNumber(272), new BigNumber((_1753_v6).length))];
        }
        let _index270 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length));
        (_1753_v6)[_index270] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("n")).length)), _1752_v5));
        let _index271 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length));
        (_1753_v6)[_index271] = ((_1742_v0) ? (new BigNumber((_1751_v4).length)) : (_1752_v5));
        if (true) {
          let _1908_v108;
          _1908_v108 = _dafny.Seq.of(_1743_v1, _1743_v1);
          (globalState).f4 = !_dafny.areEqual(_dafny.Seq.Concat(_1908_v108, _1908_v108), _1908_v108);
          let _1909_v109;
          _1909_v109 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1756_v10).length),_1804_v39);
          let _1910_v110;
          let _nw294 = Array((new BigNumber(6)).toNumber());
          _nw294[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1804_v39, _module.__default.safeIndex(_1752_v5, new BigNumber((_1804_v39).length)), new BigNumber(341));
          _nw294[(_dafny.ONE).toNumber()] = _1804_v39;
          _nw294[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(new BigNumber(68));
          _nw294[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_1753_v6)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length))], _1752_v5), (((_1909_v109).contains((_1753_v6)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length))])) ? ((_1909_v109).get((_1753_v6)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length))])) : (_1804_v39)));
          _nw294[(new BigNumber(4)).toNumber()] = _1804_v39;
          _nw294[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(126)), function (_1911_i11) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-106)), function (_1912_i12) {
              return new _dafny.CodePoint('n'.codePointAt(0));
            })).length);
          });
          _1910_v110 = _nw294;
          let _index272 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1910_v110).length));
          (_1910_v110)[_index272] = _module.__default.fm1(_1742_v0, _1742_v0, globalState);
          let _1913_v111;
          _1913_v111 = _module.D0.create_DC1();
          let _1914_v113;
          _1914_v113 = _dafny.Map.Empty.slice().updateUnsafe((_1753_v6)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length))],_1743_v1);
          let _1915_v114;
          _1915_v114 = _dafny.MultiSet.fromElements(_1914_v113, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1751_v4).length),_1743_v1));
          let _index273 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1910_v110).length));
          let _index274 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length));
          let _index275 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length));
          let _rhs243 = _dafny.Seq.Concat(_1804_v39, _dafny.Seq.Create(_module.__default.abs(new BigNumber(594)), function (_1916_i13) {
            return new BigNumber(635);
          }));
          let _rhs244 = ((new BigNumber(-415)).multipliedBy(_1752_v5)).minus((_1753_v6)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length))]);
          let _rhs245 = _module.__default.safeModuloInt(_module.__default.fm0(!(_1742_v0), function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of _dafny.IntegerRange(new BigNumber(940), new BigNumber(896))) {
              let _1917_v112 = _compr_64;
              if (((new BigNumber(940)).isLessThanOrEqualTo(_1917_v112)) && ((_1917_v112).isLessThan(new BigNumber(896)))) {
                _coll64.push([_module.__default.safeDivisionInt(_1917_v112, _1752_v5),(_module.D8.create_DC25(_1742_v0, (_1753_v6)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC20(_1888_cf30),_1752_v5)).length), _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,(_1753_v6)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length))]))).dtor_cf40]);
              }
            }
            return _coll64;
          }(), _1742_v0, globalState), (((_1915_v114).contains(_1914_v113)) ? ((_1915_v114).get(_1914_v113)) : (_1752_v5)));
          let _rhs246 = _1913_v111;
          let _lhs137 = _1910_v110;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_1910_v110).length));
          let _lhs139 = _1753_v6;
          let _lhs140 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length));
          let _lhs141 = _1753_v6;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(943), new BigNumber((_1753_v6).length));
          _lhs137[_lhs138] = _rhs243;
          _lhs139[_lhs140] = _rhs244;
          _lhs141[_lhs142] = _rhs245;
          _1913_v111 = _rhs246;
          let _init49 = ((_1918_v0) => function (_1919_i14) {
            return _1918_v0;
          })(_1742_v0);
          let _nw295 = Array((new BigNumber(5)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw295.length); _i0_49++) {
            _nw295[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1743_v1 = _nw295;
          let _1920_v115;
          _1920_v115 = _dafny.Set.fromElements(_1751_v4);
          (globalState).f4 = (_1920_v115).IsSubsetOf(_1920_v115);
          let _1921_v116;
          _1921_v116 = _dafny.MultiSet.fromElements(_1742_v0, _1742_v0, _1742_v0);
          let _1922_v117;
          _1922_v117 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_1742_v0, globalState),true);
          let _1923_v118;
          _1923_v118 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_1742_v0, globalState),_1922_v117);
          let _1924_v119;
          let _nw296 = Array((new BigNumber(28)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1924_v119 = _nw296;
          let _1925_v120;
          let _nw297 = new _module.C1();
          _nw297.__ctor(_1742_v0, (_1921_v116).IsProperSubsetOf(_1921_v116), _1923_v118, _1756_v10, _1924_v119);
          _1925_v120 = _nw297;
        } else {
          _1804_v39 = _1804_v39;
          let _1926_v121;
          _1926_v121 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v0,_1742_v0);
          let _1927_v122;
          _1927_v122 = _dafny.Map.Empty.slice().updateUnsafe(!(_1742_v0),_1926_v121);
          let _1928_v123;
          let _init50 = ((_1929_cf30) => function (_1930_i15) {
            return _1929_cf30;
          })(_1888_cf30);
          let _nw298 = Array((new BigNumber(29)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw298.length); _i0_50++) {
            _nw298[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1928_v123 = _nw298;
          let _1931_v124;
          let _nw299 = new _module.C1();
          _nw299.__ctor(!(_1742_v0), _1742_v0, _1927_v122, _1756_v10, _1928_v123);
          _1931_v124 = _nw299;
          _1931_v124 = _1931_v124;
          _1742_v0 = _1742_v0;
          let _1932_v125;
          _1932_v125 = new _dafny.CodePoint('o'.codePointAt(0));
          let _index276 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_1889_v92).length));
          (_1889_v92)[_index276] = _1932_v125;
          let _index277 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_1889_v92).length));
          (_1889_v92)[_index277] = _1932_v125;
          _1751_v4 = _dafny.Seq.UnicodeFromString("nugcjll");
        }
      }
      r0 = _1742_v0;
      r1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(254)), function (_1933_i16) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(600)), function (_1934_i17) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        });
      });
      r2 = _dafny.Seq.Concat(_1756_v10, _1756_v10);
      return [r0, r1, r2];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f12 = [];
      this._f13 = _dafny.Map.Empty;
      this._f14 = _dafny.Seq.of();
      this._f19 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1, _module.T4];
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
    __ctor(f19, f12, f13, f14) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      return;
    }
    fm5(p0, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(944), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_module.D2.create_DC4(_dafny.Seq.of((_this).f14)), _module.D2.create_DC4(_dafny.Seq.of((_this).f14))),(_this).f14)).length)));
    };
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(408)), _dafny.Seq.of(new BigNumber(436), new BigNumber((_dafny.Seq.of(new BigNumber(((_this).f14).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-25),new BigNumber(366))).length))).length), (_dafny.ZERO).minus(new BigNumber((function () {
        let _coll65 = new _dafny.Map();
        for (const _compr_65 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-832)), function (_1935_i0) {
          return new BigNumber(-887);
        })).Elements) {
          let _1936_v0 = _compr_65;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-832)), function (_1935_i0) {
            return new BigNumber(-887);
          }), _1936_v0)) {
            _coll65.push([(_1936_v0).minus(new BigNumber(435)),new BigNumber(150)]);
          }
        }
        return _coll65;
      }()).length)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), function (_1937_i1) {
        return (_dafny.ZERO).minus(new BigNumber(-155));
      }));
    };
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(new BigNumber(293))).Intersect(_dafny.Set.fromElements(new BigNumber(-796)))).Difference((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(-378), new BigNumber(984), new BigNumber((_dafny.Seq.UnicodeFromString("wnblljb")).length))).length))).Difference(_dafny.Set.fromElements(new BigNumber(-459))));
    };
    m3(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let r3 = false;
      let _1938_v0;
      _1938_v0 = new BigNumber(208);
      let _1939_v1;
      _1939_v1 = _dafny.MultiSet.fromElements(_1938_v0);
      let _1940_v2;
      _1940_v2 = _module.__default.fm20((_this).f19, _1938_v0, globalState);
      _1939_v1 = function (_source30) {
        let _1941___mcc_h0 = _source30;
        let _1942_cf3 = _1941___mcc_h0;
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xh")).length));
      }(_1940_v2);
      let _1943_v3;
      _1943_v3 = _module.D0.create_DC1();
      let _source31 = _1943_v3;
      if (_source31.is_DC1) {
        let _1944_v4;
        let _init51 = function (_1945_i0) {
          return (_this).f19;
        };
        let _nw300 = Array((new BigNumber(21)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw300.length); _i0_51++) {
          _nw300[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _1944_v4 = _nw300;
        let _index278 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1944_v4).length));
        (_1944_v4)[_index278] = (_this).f19;
        let _1946_v5;
        _1946_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
        let _1947_v6;
        _1947_v6 = _dafny.MultiSet.fromElements(_1946_v5, _1946_v5);
        let _index279 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1944_v4).length));
        (_1944_v4)[_index279] = (_1947_v6).IsSubsetOf(((_module.__default.fm2((_this).f19, globalState)) ? ((_1947_v6).update(_1946_v5, _module.__default.abs(new BigNumber(493)))) : (_1947_v6)));
        let _1948_v7;
        _1948_v7 = _dafny.Seq.UnicodeFromString("sacmgtow");
        _1948_v7 = _dafny.Seq.update(_module.__default.fm21((_this).f19, globalState), _module.__default.safeIndex(_1938_v0, new BigNumber((_module.__default.fm21((_this).f19, globalState)).length)), new _dafny.CodePoint('m'.codePointAt(0)));
        let _1949_v8;
        let _nw301 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _1949_v8 = _nw301;
        let _index280 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1949_v8).length));
        (_1949_v8)[_index280] = (_1938_v0).minus(_1938_v0);
        let _index281 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1949_v8).length));
        (_1949_v8)[_index281] = _1938_v0;
        let _1950_v9;
        _1950_v9 = _module.D0.create_DC2(new BigNumber(363), (_this).f19);
        _1950_v9 = _1950_v9;
      } else if (_source31.is_DC2) {
        let _1951___mcc_h1 = (_source31).cf1;
        let _1952___mcc_h2 = (_source31).cf2;
        let _1953_cf2 = _1952___mcc_h2;
        let _1954_cf1 = _1951___mcc_h1;
        let _1955_v10;
        _1955_v10 = _dafny.Map.Empty.slice().updateUnsafe(false,_1953_cf2);
        let _1956_v11;
        _1956_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1955_v10).length),_1953_cf2);
        let _1957_v12;
        _1957_v12 = _dafny.MultiSet.fromElements(_1953_cf2);
        let _1958_v13;
        _1958_v13 = _dafny.Seq.UnicodeFromString("jsmors");
        let _rhs247 = _1956_v11;
        let _rhs248 = !(((_this).f14)[_module.__default.safeIndex(_1938_v0, new BigNumber(((_this).f14).length))]);
        let _rhs249 = (_module.__default.safeModuloInt(_1938_v0, (_dafny.ZERO).minus(_1938_v0))).isLessThanOrEqualTo((((_1957_v12).contains(true)) ? ((_1957_v12).get(true)) : (new BigNumber((_dafny.Seq.of(new BigNumber((_1958_v13).length))).length))));
        _1956_v11 = _rhs247;
        _1953_cf2 = _rhs248;
        r3 = _rhs249;
        r0 = _1953_cf2;
        let _1959_v14;
        let _nw302 = new _module.C0();
        _nw302.__ctor();
        _1959_v14 = _nw302;
        let _1960_v15;
        _1960_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f19, _1953_cf2),_1959_v14);
        _1960_v15 = (_1960_v15).update((_1957_v12).Union(_1957_v12), _1959_v14);
        let _1961_v16;
        _1961_v16 = _module.D0.create_DC0(_1953_cf2);
        let _1962_v17;
        _1962_v17 = _module.D2.create_DC5(new BigNumber((_dafny.Seq.UnicodeFromString("hbajbykx")).length), (_this).f19, (_1961_v16).dtor_cf0);
        let _1963_v18;
        _1963_v18 = _module.D2.create_DC7(_1962_v17);
        _1963_v18 = _1963_v18;
      } else {
        let _1964___mcc_h3 = (_source31).cf0;
        let _1965_cf0 = _1964___mcc_h3;
        let _1966_v19;
        let _init52 = ((_1967_v0) => function (_1968_i1) {
          return (_1968_i1).minus(_1967_v0);
        })(_1938_v0);
        let _nw303 = Array((new BigNumber(11)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw303.length); _i0_52++) {
          _nw303[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1966_v19 = _nw303;
        _1966_v19 = _1966_v19;
        let _rhs250 = (_1938_v0).multipliedBy(_1938_v0);
        let _rhs251 = _1938_v0;
        _1938_v0 = _rhs250;
        r1 = _rhs251;
        let _1969_v20;
        let _nw304 = new _module.C0();
        _nw304.__ctor();
        _1969_v20 = _nw304;
        let _1970_v21;
        _1970_v21 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1971_v22;
        _1971_v22 = _dafny.Seq.UnicodeFromString("d");
        _1965_cf0 = !_dafny.Seq.contains(_1971_v22, _1970_v21);
      }
      let _1972_v23;
      _1972_v23 = _dafny.Seq.of((_this).f14, _dafny.Seq.of((_this).f19));
      let _1973_v24;
      _1973_v24 = _module.D2.create_DC4(_1972_v23);
      _1940_v2 = _module.__default.fm22(_1973_v24, globalState);
      if ((_this).f19) {
        _1938_v0 = _1938_v0;
        let _1974_v25;
        let _nw305 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _1974_v25 = _nw305;
        let _1975_v26;
        _1975_v26 = _dafny.MultiSet.fromElements(_1974_v25, _1974_v25);
        let _1976_v27;
        _1976_v27 = _dafny.Set.fromElements(_1938_v0);
        let _1977_v28;
        _1977_v28 = _dafny.Seq.UnicodeFromString("bpo");
        let _1978_v29;
        let _nw306 = Array((new BigNumber(22)).toNumber());
        _nw306[(_dafny.ZERO).toNumber()] = (((_this).f19) ? ((_this).f19) : ((_this).f19));
        _nw306[(_dafny.ONE).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(2)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(3)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(4)).toNumber()] = ((_this).f19) === (_module.__default.fm2((_this).f19, globalState));
        _nw306[(new BigNumber(5)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(6)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(7)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(8)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.of((_this).f19, (_this).f19, (_this).f19, (_this).f19, (_this).f19), (_this).f19);
        _nw306[(new BigNumber(9)).toNumber()] = !(false);
        _nw306[(new BigNumber(10)).toNumber()] = (_1975_v26).IsSubsetOf((_1975_v26).update(_1974_v25, _module.__default.abs(_1938_v0)));
        _nw306[(new BigNumber(11)).toNumber()] = !(new BigNumber(-711)).isEqualTo(_1938_v0);
        _nw306[(new BigNumber(12)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(13)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(14)).toNumber()] = ((_this).fm10(false, _1976_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(348)), function (_1979_i2) {
          return new BigNumber(-833);
        }), _1938_v0, globalState)).contains(_1938_v0);
        _nw306[(new BigNumber(15)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(16)).toNumber()] = _module.__default.fm2((_this).f19, globalState);
        _nw306[(new BigNumber(17)).toNumber()] = _dafny.areEqual(_1977_v28, _1977_v28);
        _nw306[(new BigNumber(18)).toNumber()] = false;
        _nw306[(new BigNumber(19)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(20)).toNumber()] = (_this).f19;
        _nw306[(new BigNumber(21)).toNumber()] = (true) === ((_this).f19);
        _1978_v29 = _nw306;
        let _index282 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length));
        (_1978_v29)[_index282] = (_this).f19;
        let _index283 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length));
        (_1978_v29)[_index283] = true;
        let _1980_v30;
        _1980_v30 = _dafny.Seq.of(_1938_v0);
        _1938_v0 = ((_dafny.ZERO).minus(_1938_v0)).plus((_1980_v30)[_module.__default.safeIndex(new BigNumber(-412), new BigNumber((_1980_v30).length))]);
        if ((_this).f19) {
          let _index284 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1974_v25).length));
          (_1974_v25)[_index284] = new BigNumber(((_1976_v27).Intersect(_1976_v27)).length);
          let _1981_v31;
          let _nw307 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1981_v31 = _nw307;
          let _1982_v32;
          let _nw308 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1982_v32 = _nw308;
          let _index285 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1981_v31).length));
          (_1981_v31)[_index285] = _1982_v32;
          let _1983_v33;
          _1983_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(615),(_this).f19);
          let _index286 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1974_v25).length));
          let _index287 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1981_v31).length));
          let _rhs252 = _1938_v0;
          let _rhs253 = _module.__default.safeDivisionInt(_1938_v0, _1938_v0);
          let _rhs254 = _module.__default.safeDivisionInt(new BigNumber(501), new BigNumber(576));
          let _rhs255 = _1982_v32;
          let _rhs256 = (_1980_v30)[_module.__default.safeIndex(new BigNumber((_1983_v33).length), new BigNumber((_1980_v30).length))];
          let _lhs143 = _1974_v25;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1974_v25).length));
          let _lhs145 = _1981_v31;
          let _lhs146 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1981_v31).length));
          _lhs143[_lhs144] = _rhs252;
          _1938_v0 = _rhs253;
          r1 = _rhs254;
          _lhs145[_lhs146] = _rhs255;
          r1 = _rhs256;
          let _1984_v34;
          let _nw309 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _1984_v34 = _nw309;
          _1984_v34 = _1984_v34;
          _1980_v30 = _dafny.Seq.Concat(_1980_v30, _1980_v30);
          let _1985_v35;
          _1985_v35 = _module.D0.create_DC0((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))]);
          let _index288 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length));
          let _index289 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length));
          let _index290 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1974_v25).length));
          let _rhs257 = false;
          let _rhs258 = (_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))];
          let _rhs259 = ((_module.__default.fm2((_1985_v35).dtor_cf0, globalState)) ? (new BigNumber((_1977_v28).length)) : ((_1974_v25)[_module.__default.safeIndex(new BigNumber(596), new BigNumber((_1974_v25).length))]));
          let _lhs147 = _1978_v29;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length));
          let _lhs149 = _1978_v29;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length));
          let _lhs151 = _1974_v25;
          let _lhs152 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_1974_v25).length));
          _lhs147[_lhs148] = _rhs257;
          _lhs149[_lhs150] = _rhs258;
          _lhs151[_lhs152] = _rhs259;
          _1938_v0 = (_this).fm5((_this).f19, globalState);
        } else {
          (globalState).f4 = (_this).f19;
          (globalState).f4 = (_this).f19;
          let _rhs260 = (_dafny.ZERO).minus(_1938_v0);
          let _rhs261 = _1978_v29;
          r1 = _rhs260;
          _1978_v29 = _rhs261;
          r3 = !((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))]);
          let _1986_v36;
          let _nw310 = new _module.C0();
          _nw310.__ctor();
          _1986_v36 = _nw310;
        }
        let _1987_v37;
        _1987_v37 = new _dafny.CodePoint('r'.codePointAt(0));
        let _source32 = ((((_this).f14)[_module.__default.safeIndex(_1938_v0, new BigNumber(((_this).f14).length))]) ? (_1973_v24) : (_module.__default.fm23(_1938_v0, _1987_v37, globalState)));
        if (_source32.is_DC5) {
          let _1988___mcc_h4 = (_source32).cf5;
          let _1989___mcc_h5 = (_source32).cf6;
          let _1990___mcc_h6 = (_source32).cf7;
          let _1991_cf7 = _1990___mcc_h6;
          let _1992_cf6 = _1989___mcc_h5;
          let _1993_cf5 = _1988___mcc_h4;
          _1993_cf5 = new BigNumber(174);
          (globalState).f4 = _1991_cf7;
          let _1994_v38;
          _1994_v38 = _module.D0.create_DC2(_1993_cf5, true);
          let _1995_v39;
          let _1996_v40;
          let _1997_v41;
          let _1998_v42;
          let _out107;
          let _out108;
          let _out109;
          let _out110;
          let _outcollector28 = _module.__default.m0(false, _1994_v38, (_this).f19, globalState);
          _out107 = _outcollector28[0];
          _out108 = _outcollector28[1];
          _out109 = _outcollector28[2];
          _out110 = _outcollector28[3];
          _1995_v39 = _out107;
          _1996_v40 = _out108;
          _1997_v41 = _out109;
          _1998_v42 = _out110;
          let _1999_v43;
          _1999_v43 = _dafny.Set.fromElements(_1978_v29, _1978_v29, _1978_v29, _1978_v29);
          _1999_v43 = ((_1999_v43).Union(_1999_v43)).Difference(_1999_v43);
        } else if (_source32.is_DC6) {
          let _2000___mcc_h7 = (_source32).cf8;
          let _2001___mcc_h8 = (_source32).cf9;
          let _2002_cf9 = _2001___mcc_h8;
          let _2003_cf8 = _2000___mcc_h7;
          r1 = (_1938_v0).multipliedBy(_module.__default.safeModuloInt(_1938_v0, new BigNumber(((_this).f14).length)));
          let _2004_v44;
          _2004_v44 = _module.D2.create_DC5(_1938_v0, (_this).f19, (_this).f19);
          let _2005_v45;
          _2005_v45 = _module.D2.create_DC7(_2004_v44);
          _2005_v45 = _2005_v45;
          (globalState).f4 = (_1938_v0).isEqualTo((_1938_v0).multipliedBy((_1980_v30)[_module.__default.safeIndex(_1938_v0, new BigNumber((_1980_v30).length))]));
          r3 = (_this).f19;
        } else if (_source32.is_DC4) {
          let _2006___mcc_h9 = (_source32).cf4;
          let _2007_cf4 = _2006___mcc_h9;
          let _2008_v46;
          _2008_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1987_v37,(_this).f19);
          _2008_v46 = (_2008_v46).update(_1987_v37, _dafny.Seq.IsPrefixOf(_1980_v30, _1980_v30));
          let _index291 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1974_v25).length));
          (_1974_v25)[_index291] = _1938_v0;
          let _index292 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1974_v25).length));
          (_1974_v25)[_index292] = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(_1938_v0, _1938_v0)).plus(_1938_v0));
          let _2009_v47;
          _2009_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1980_v30);
          let _2010_v48;
          _2010_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,!((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))]));
          let _2011_v49;
          _2011_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2009_v47,(((_2010_v48).contains((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))])) ? ((_2010_v48).get((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))])) : ((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))])));
          _2011_v49 = (_2011_v49).update(_2009_v47, (_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))]);
          let _2012_v50;
          let _init53 = ((_2013_v37) => function (_2014_i3) {
            return _2013_v37;
          })(_1987_v37);
          let _nw311 = Array((new BigNumber(9)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw311.length); _i0_53++) {
            _nw311[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _2012_v50 = _nw311;
          let _index293 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_2012_v50).length));
          (_2012_v50)[_index293] = _1987_v37;
          let _index294 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_2012_v50).length));
          (_2012_v50)[_index294] = ((true) ? (_1987_v37) : (_1987_v37));
        } else {
          let _2015___mcc_h10 = (_source32).cf10;
          let _2016_cf10 = _2015___mcc_h10;
          let _2017_v51;
          _2017_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1938_v0,_1938_v0);
          r1 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f14).length),_1938_v0)).Merge(_2017_v51)).length);
          r0 = _module.__default.fm2(!((_this).f19) || ((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))]), globalState);
          _1938_v0 = (_1938_v0).multipliedBy(_1938_v0);
          let _2018_v52;
          _2018_v52 = _dafny.Set.fromElements(_module.__default.fm24((_1978_v29)[_module.__default.safeIndex(new BigNumber(88), new BigNumber((_1978_v29).length))], _dafny.MultiSet.FromArray(_1980_v30), new BigNumber(749), globalState));
          (globalState).f4 = (_2018_v52).IsProperSubsetOf(_2018_v52);
        }
      } else {
        let _2019_v53;
        let _nw312 = Array((new BigNumber(27)).toNumber()).fill(false);
        _2019_v53 = _nw312;
        let _index295 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_2019_v53).length));
        (_2019_v53)[_index295] = (_this).f19;
        let _index296 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_2019_v53).length));
        (_2019_v53)[_index296] = ((_this).f14)[_module.__default.safeIndex(_1938_v0, new BigNumber(((_this).f14).length))];
        let _2020_v54;
        _2020_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1938_v0,_2019_v53);
        let _2021_v55;
        let _nw313 = Array((new BigNumber(29)).toNumber());
        _nw313[(_dafny.ZERO).toNumber()] = _2019_v53;
        _nw313[(_dafny.ONE).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(2)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(3)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(4)).toNumber()] = (((_2019_v53)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_2019_v53).length))]) ? (_2019_v53) : (_2019_v53));
        _nw313[(new BigNumber(5)).toNumber()] = (((_2020_v54).contains(_1938_v0)) ? ((_2020_v54).get(_1938_v0)) : (_2019_v53));
        _nw313[(new BigNumber(6)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(7)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(8)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(9)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(10)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(11)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(12)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(13)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(14)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(15)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(16)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(17)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(18)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(19)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(20)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(21)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(22)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(23)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(24)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(25)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(26)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(27)).toNumber()] = _2019_v53;
        _nw313[(new BigNumber(28)).toNumber()] = _2019_v53;
        _2021_v55 = _nw313;
        _2021_v55 = _2021_v55;
        let _2022_v56;
        let _nw314 = new _module.C1();
        _nw314.__ctor((_2019_v53)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_2019_v53).length))], ((_2019_v53)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_2019_v53).length))]) || ((_2019_v53)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_2019_v53).length))]), (_this).f13, (_this).f14, _this.f12);
        _2022_v56 = _nw314;
        let _2023_v57;
        let _nw315 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _2023_v57 = _nw315;
        let _index297 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_2023_v57).length));
        (_2023_v57)[_index297] = (_1938_v0).multipliedBy(_1938_v0);
        let _index298 = _module.__default.safeIndex(new BigNumber(993), new BigNumber((_2023_v57).length));
        (_2023_v57)[_index298] = ((_2022_v56).fm5(!(false), globalState)).multipliedBy(_1938_v0);
        let _2024_v58;
        _2024_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_2022_v56).f17);
        r3 = (_2024_v58).equals((_2024_v58).update((_this).f14, (_this).f19));
      }
      let _2025_i4;
      _2025_i4 = _dafny.ZERO;
      L11: {
        while ((_this).f19) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2025_i4)) {
              break L11;
            }
            _2025_i4 = (_2025_i4).plus(_dafny.ONE);
            let _2026_v59;
            let _nw316 = Array((_dafny.ONE).toNumber());
            _nw316[(_dafny.ZERO).toNumber()] = !(!(true) || (true));
            _2026_v59 = _nw316;
            let _index299 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_2026_v59).length));
            (_2026_v59)[_index299] = (_this).f19;
            let _index300 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_2026_v59).length));
            (_2026_v59)[_index300] = (_this).f19;
            _1938_v0 = (new BigNumber(43)).minus(_1938_v0);
            let _2027_v60;
            _2027_v60 = new _dafny.CodePoint('s'.codePointAt(0));
            let _2028_v61;
            _2028_v61 = _dafny.Set.fromElements(_2027_v60);
            let _index301 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_2026_v59).length));
            (_2026_v59)[_index301] = (((false) ? (_2028_v61) : (_2028_v61))).contains(_2027_v60);
            let _2029_v62;
            _2029_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1938_v0,(_this).f13);
            let _2030_v63;
            _2030_v63 = _dafny.Map.Empty.slice().updateUnsafe((_2026_v59)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_2026_v59).length))],(_this).f19);
            let _2031_v64;
            _2031_v64 = _module.D3.create_DC8(_2030_v63);
            let _2032_v65;
            let _nw317 = new _module.C1();
            _nw317.__ctor((_2026_v59)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_2026_v59).length))], !(!((_2026_v59)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_2026_v59).length))]) || ((_this).f19)), ((((_2029_v62).contains(_1938_v0)) ? ((_2029_v62).get(_1938_v0)) : ((_this).f13))).Merge(((_this).f13).update((_this).f19, (_2031_v64).dtor_cf11)), (_this).f14, _this.f12);
            _2032_v65 = _nw317;
          }
        }
      }
      let _2033_v66;
      _2033_v66 = _dafny.Seq.UnicodeFromString("pyqhtx");
      let _2034_v67;
      _2034_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2033_v66,!(((_this).f14)[_module.__default.safeIndex(_1938_v0, new BigNumber(((_this).f14).length))]));
      r3 = (_2034_v67).contains(_dafny.Seq.UnicodeFromString("bqw"));
      r0 = (_module.D0.create_DC0((_this).f19)).dtor_cf0;
      r1 = _1938_v0;
      let _2035_v68;
      _2035_v68 = _dafny.Set.fromElements((_this).f14);
      r2 = _2035_v68;
      let _2036_v69;
      _2036_v69 = new _dafny.CodePoint('r'.codePointAt(0));
      let _2037_v70;
      _2037_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_2036_v69);
      r3 = (_module.__default.fm2((_this).f19, globalState)) || (!(_2037_v70).contains((_this).f19));
      return [r0, r1, r2, r3];
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _2038_v0;
      _2038_v0 = _module.D3.create_DC9(new BigNumber(956), p2);
      let _source33 = _2038_v0;
      if (_source33.is_DC9) {
        let _2039___mcc_h0 = (_source33).cf12;
        let _2040___mcc_h1 = (_source33).cf13;
        let _2041_cf13 = _2040___mcc_h1;
        let _2042_cf12 = _2039___mcc_h0;
        let _2043_v1;
        _2043_v1 = _module.D4.create_DC12((_this).f14);
        let _2044_v2;
        _2044_v2 = _dafny.MultiSet.fromElements((_2043_v1).dtor_cf18);
        let _2045_v3;
        _2045_v3 = _dafny.Seq.of(p1);
        let _2046_v4;
        _2046_v4 = _dafny.Map.Empty.slice().updateUnsafe(_2044_v2,_dafny.Seq.update(_dafny.Seq.Concat(_2045_v3, _2045_v3), _module.__default.safeIndex(_2042_cf12, new BigNumber((_dafny.Seq.Concat(_2045_v3, _2045_v3)).length)), p1));
        _2046_v4 = (_2046_v4).update(_2044_v2, _2045_v3);
        let _2047_v5;
        let _nw318 = Array((new BigNumber(11)).toNumber()).fill(false);
        _2047_v5 = _nw318;
        let _2048_v6;
        let _nw319 = new _module.C1();
        _nw319.__ctor((_module.D2.create_DC5(_2042_cf12, (_this).f19, (_this).f19)).dtor_cf6, (_this).f19, (_this).f13, (_this).f14, _this.f12);
        _2048_v6 = _nw319;
        let _index302 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_2047_v5).length));
        (_2047_v5)[_index302] = (function (_pat_let33_0) {
          return function (_2049_dt__update__tmp_h0) {
            return function (_pat_let34_0) {
              return function (_2050_dt__update_hcf14_h0) {
                return _module.D3.create_DC10(_2050_dt__update_hcf14_h0, (_2049_dt__update__tmp_h0).dtor_cf15, (_2049_dt__update__tmp_h0).dtor_cf16);
              }(_pat_let34_0);
            }((_this).f19);
          }(_pat_let33_0);
        }(_module.D3.create_DC10((_this).f19, _2042_cf12, _2048_v6))).dtor_cf14;
        let _2051_v7;
        _2051_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
        let _index303 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_2047_v5).length));
        (_2047_v5)[_index303] = !((_2051_v7).Merge(_2051_v7)).equals(_2051_v7);
        let _2052_v8;
        _2052_v8 = _dafny.Seq.UnicodeFromString("vbxdea");
        _2052_v8 = _2052_v8;
        (globalState).f4 = (p2).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber(649), _2042_cf12));
      } else if (_source33.is_DC10) {
        let _2053___mcc_h2 = (_source33).cf14;
        let _2054___mcc_h3 = (_source33).cf15;
        let _2055___mcc_h4 = (_source33).cf16;
        let _2056_cf16 = _2055___mcc_h4;
        let _2057_cf15 = _2054___mcc_h3;
        let _2058_cf14 = _2053___mcc_h2;
        let _2059_v9;
        _2059_v9 = _dafny.Seq.of(p2);
        r0 = _module.__default.fm2(_dafny.areEqual(_dafny.Seq.of(p2), _2059_v9), globalState);
        let _2060_v10;
        _2060_v10 = _dafny.Map.Empty.slice().updateUnsafe(((_2056_cf16).f14)[_module.__default.safeIndex(_2057_cf15, new BigNumber(((_2056_cf16).f14).length))],new BigNumber(351));
        _2060_v10 = (_2060_v10).update(!((_2057_cf15).isEqualTo(_2057_cf15)), _2057_cf15);
        let _2061_v11;
        _2061_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
        let _2062_v12;
        let _nw320 = new _module.C1();
        _nw320.__ctor(_2058_cf14, !(_2058_cf14) || (_2058_cf14), ((_2056_cf16).f13).Merge(((_2056_cf16).f13).update(_2058_cf14, _2061_v11)), (_this).f14, ((_2058_cf14) ? (_2056_cf16.f12) : (_this.f12)));
        _2062_v12 = _nw320;
        let _source34 = _module.__default.fm25(_2059_v9, globalState);
        if (_source34.is_DC9) {
          let _2063___mcc_h7 = (_source34).cf12;
          let _2064___mcc_h8 = (_source34).cf13;
          let _2065_cf13 = _2064___mcc_h8;
          let _2066_cf12 = _2063___mcc_h7;
          let _2067_v13;
          let _init54 = ((_2068_p2) => function (_2069_i0) {
            return _module.__default.safeDivisionInt(_2069_i0, _2068_p2);
          })(p2);
          let _nw321 = Array((new BigNumber(8)).toNumber());
          for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw321.length); _i0_54++) {
            _nw321[_i0_54] = _init54(new BigNumber(_i0_54));
          }
          _2067_v13 = _nw321;
          _2067_v13 = ((_2058_cf14) ? (_2067_v13) : (_2067_v13));
          (_2062_v12).f18 = !(!(_2062_v12.f18) || (_2062_v12.f18)) || ((_this).f19);
          let _index304 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2067_v13).length));
          (_2067_v13)[_index304] = _2066_cf12;
          let _2070_v14;
          _2070_v14 = _module.D5.create_DC15(_2059_v9);
          let _index305 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2067_v13).length));
          let _rhs262 = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(789)), ((_2071_p2) => function (_2072_i1) {
            return _2071_p2;
          })(p2)), _dafny.Seq.Concat((_2070_v14).dtor_cf25, _2059_v9)));
          let _rhs263 = (_2059_v9)[_module.__default.safeIndex(_2066_cf12, new BigNumber((_2059_v9).length))];
          let _rhs264 = _module.__default.safeModuloInt(_2065_cf13, _2065_cf13);
          let _rhs265 = _2058_cf14;
          let _lhs153 = globalState;
          let _lhs154 = _2067_v13;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_2067_v13).length));
          let _lhs156 = _2062_v12;
          _lhs153.f4 = _rhs262;
          _2057_cf15 = _rhs263;
          _lhs154[_lhs155] = _rhs264;
          _lhs156.f18 = _rhs265;
          let _2073_v15;
          let _nw322 = new _module.C0();
          _nw322.__ctor();
          _2073_v15 = _nw322;
        } else if (_source34.is_DC10) {
          let _2074___mcc_h9 = (_source34).cf14;
          let _2075___mcc_h10 = (_source34).cf15;
          let _2076___mcc_h11 = (_source34).cf16;
          let _2077_cf16 = _2076___mcc_h11;
          let _2078_cf15 = _2075___mcc_h10;
          let _2079_cf14 = _2074___mcc_h9;
          let _2080_v16;
          let _init55 = ((_2081_v12) => function (_2082_i2) {
            return !(_2081_v12.f18);
          })(_2062_v12);
          let _nw323 = Array((new BigNumber(2)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw323.length); _i0_55++) {
            _nw323[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _2080_v16 = _nw323;
          let _2083_v17;
          _2083_v17 = _dafny.Set.fromElements(_2080_v16, _2080_v16, _2080_v16, _2080_v16, _2080_v16);
          _2083_v17 = _2083_v17;
          let _2084_v18;
          _2084_v18 = _dafny.Map.Empty.slice().updateUnsafe((_2062_v12).f17,p1);
          _2084_v18 = (_2084_v18).update(_2079_cf14, _dafny.Seq.Concat(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_2085_i3) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })));
          let _2086_v19;
          let _nw324 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2086_v19 = _nw324;
          let _index306 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_2086_v19).length));
          (_2086_v19)[_index306] = _dafny.Seq.UnicodeFromString("su");
          let _index307 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_2086_v19).length));
          (_2086_v19)[_index307] = _dafny.Seq.Concat(p1, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yi"), _dafny.Seq.UnicodeFromString("rffepbepm")));
          (globalState).f4 = _module.__default.fm2((_2062_v12).f17, globalState);
        } else if (_source34.is_DC8) {
          let _2087___mcc_h12 = (_source34).cf11;
          let _2088_cf11 = _2087___mcc_h12;
          let _2089_v20;
          let _nw325 = Array((new BigNumber(15)).toNumber()).fill(false);
          _2089_v20 = _nw325;
          let _index308 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length));
          (_2089_v20)[_index308] = (_this).f19;
          let _index309 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length));
          let _rhs266 = (_2062_v12).f17;
          let _rhs267 = (_2062_v12).fm17(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(22)), ((_2090_cf15) => function (_2091_i4) {
            return _2090_cf15;
          })(_2057_cf15)), _2059_v9), new BigNumber((((_2058_cf14) ? (p1) : (p1))).length), globalState);
          let _lhs157 = _2089_v20;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length));
          _lhs157[_lhs158] = _rhs266;
          _2058_cf14 = _rhs267;
          let _2092_v21;
          let _nw326 = new _module.C0();
          _nw326.__ctor();
          _2092_v21 = _nw326;
          let _index310 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length));
          (_2089_v20)[_index310] = !((((_2088_cf11).contains(!((_2089_v20)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length))]))) ? ((_2088_cf11).get(!((_2089_v20)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length))]))) : (_2062_v12.f18))) || ((((_2089_v20)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length))]) ? ((_2089_v20)[_module.__default.safeIndex(new BigNumber(286), new BigNumber((_2089_v20).length))]) : (_2058_cf14)));
          _2057_cf15 = p2;
        } else {
          let _2093___mcc_h13 = (_source34).cf17;
          let _2094_cf17 = _2093___mcc_h13;
          r0 = _2062_v12.f18;
          let _2095_v22;
          _2095_v22 = _dafny.MultiSet.fromElements(p2, p2);
          let _2096_v23;
          _2096_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2057_cf15);
          let _2097_v24;
          _2097_v24 = _module.D6.create_DC17(_dafny.Set.fromElements(p2, new BigNumber(300), p2));
          let _2098_v25;
          _2098_v25 = _dafny.Set.fromElements(p2);
          let _pat_let_tv35 = _2098_v25;
          let _2099_v26;
          let _nw327 = Array((new BigNumber(24)).toNumber());
          _nw327[(_dafny.ZERO).toNumber()] = p2;
          _nw327[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(_2057_cf15, _2057_cf15);
          _nw327[(new BigNumber(2)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(3)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(4)).toNumber()] = (_this).fm5((_2062_v12).f17, globalState);
          _nw327[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw327[(new BigNumber(6)).toNumber()] = new BigNumber(-268);
          _nw327[(new BigNumber(7)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(8)).toNumber()] = (((_2060_v10).contains((_2062_v12).f17)) ? ((_2060_v10).get((_2062_v12).f17)) : ((((_2095_v22).contains(_2057_cf15)) ? ((_2095_v22).get(_2057_cf15)) : (p2))));
          _nw327[(new BigNumber(9)).toNumber()] = new BigNumber((_2096_v23).length);
          _nw327[(new BigNumber(10)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(11)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(12)).toNumber()] = p2;
          _nw327[(new BigNumber(13)).toNumber()] = new BigNumber(((_2056_cf16).fm6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(755)), function (_2100_i5) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }), globalState)).length);
          _nw327[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_2057_cf15);
          _nw327[(new BigNumber(15)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(16)).toNumber()] = (_2057_cf15).multipliedBy(_2057_cf15);
          _nw327[(new BigNumber(17)).toNumber()] = (_2057_cf15).minus(_2057_cf15);
          _nw327[(new BigNumber(18)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(19)).toNumber()] = p2;
          _nw327[(new BigNumber(20)).toNumber()] = p2;
          _nw327[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-289));
          _nw327[(new BigNumber(22)).toNumber()] = _2057_cf15;
          _nw327[(new BigNumber(23)).toNumber()] = (new BigNumber(((function (_pat_let35_0) {
            return function (_2101_dt__update__tmp_h1) {
              return function (_pat_let36_0) {
                return function (_2102_dt__update_hcf27_h0) {
                  return _module.D6.create_DC17(_2102_dt__update_hcf27_h0);
                }(_pat_let36_0);
              }(_pat_let_tv35);
            }(_pat_let35_0);
          }(_2097_v24)).dtor_cf27).length)).plus(new BigNumber(842));
          _2099_v26 = _nw327;
          let _2103_v27;
          _2103_v27 = _module.D3.create_DC10((_2062_v12).f17, _2057_cf15, _2056_cf16);
          let _index311 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_2099_v26).length));
          (_2099_v26)[_index311] = (_2103_v27).dtor_cf15;
          let _index312 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_2099_v26).length));
          (_2099_v26)[_index312] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_2057_cf15).minus(_2057_cf15)));
          let _2104_v28;
          _2104_v28 = _dafny.MultiSet.fromElements(_2058_cf14);
          let _2105_v29;
          _2105_v29 = _dafny.Seq.of(_2104_v28, _2104_v28, _2104_v28);
          let _2106_v30;
          _2106_v30 = _dafny.Map.Empty.slice().updateUnsafe(!((_2062_v12).f17),_2104_v28);
          let _2107_v31;
          _2107_v31 = _module.D4.create_DC13(_2062_v12.f18, (_2062_v12).f17, p2, (_2062_v12).fm5((_2062_v12).f17, globalState), _2062_v12.f18);
          let _2108_v32;
          let _nw328 = Array((new BigNumber(27)).toNumber());
          _nw328[(_dafny.ZERO).toNumber()] = _2104_v28;
          _nw328[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_2062_v12.f18, (_2062_v12).f17);
          _nw328[(new BigNumber(2)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(3)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(4)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(5)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(6)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(7)).toNumber()] = (_2105_v29)[_module.__default.safeIndex(_2057_cf15, new BigNumber((_2105_v29).length))];
          _nw328[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.fromElements(_2062_v12.f18)).Difference(_dafny.MultiSet.FromArray((_2056_cf16).f14));
          _nw328[(new BigNumber(9)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(10)).toNumber()] = (_2104_v28).Intersect(_2104_v28);
          _nw328[(new BigNumber(11)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_2058_cf14, (_2062_v12).f17);
          _nw328[(new BigNumber(13)).toNumber()] = (_2104_v28).Difference(_dafny.MultiSet.FromArray((_2056_cf16).f14));
          _nw328[(new BigNumber(14)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.FromArray((_2056_cf16).f14)).update(_2058_cf14, _module.__default.abs(_2057_cf15));
          _nw328[(new BigNumber(16)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(17)).toNumber()] = (((_2106_v30).contains(_2062_v12.f18)) ? ((_2106_v30).get(_2062_v12.f18)) : (_dafny.MultiSet.fromElements(_2062_v12.f18)));
          _nw328[(new BigNumber(18)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(19)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(20)).toNumber()] = (_2104_v28).Union(_dafny.MultiSet.fromElements(true, (_2062_v12).fm17(_2059_v9, _2057_cf15, globalState)));
          _nw328[(new BigNumber(21)).toNumber()] = (((_this).f19) ? (_2104_v28) : (_2104_v28));
          _nw328[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements(_2062_v12.f18, false);
          _nw328[(new BigNumber(23)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(24)).toNumber()] = _2104_v28;
          _nw328[(new BigNumber(25)).toNumber()] = _dafny.MultiSet.fromElements((_2107_v31).dtor_cf20);
          _nw328[(new BigNumber(26)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat((_this).f14, (_2056_cf16).f14));
          _2108_v32 = _nw328;
          let _index313 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_2108_v32).length));
          (_2108_v32)[_index313] = (_2104_v28).update((_2062_v12).f17, _module.__default.abs(p2));
          let _index314 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_2099_v26).length));
          let _index315 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_2108_v32).length));
          let _rhs268 = (_2099_v26)[_module.__default.safeIndex(new BigNumber(210), new BigNumber((_2099_v26).length))];
          let _rhs269 = true;
          let _rhs270 = (_dafny.ZERO).minus((new BigNumber(268)).plus(p2));
          let _rhs271 = (_2104_v28).Difference(_2104_v28);
          let _rhs272 = (_2062_v12).f17;
          let _lhs159 = _2099_v26;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_2099_v26).length));
          let _lhs161 = globalState;
          let _lhs162 = _2108_v32;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_2108_v32).length));
          let _lhs164 = globalState;
          _lhs159[_lhs160] = _rhs268;
          _lhs161.f4 = _rhs269;
          _2057_cf15 = _rhs270;
          _lhs162[_lhs163] = _rhs271;
          _lhs164.f4 = _rhs272;
          let _2109_v33;
          _2109_v33 = new _dafny.CodePoint('u'.codePointAt(0));
          _2109_v33 = _2109_v33;
        }
      } else if (_source33.is_DC8) {
        let _2110___mcc_h5 = (_source33).cf11;
        let _2111_cf11 = _2110___mcc_h5;
        if ((_this).f19) {
          let _2112_v34;
          let _init56 = function (_2113_i6) {
            return (_this).f19;
          };
          let _nw329 = Array((new BigNumber(11)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw329.length); _i0_56++) {
            _nw329[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2112_v34 = _nw329;
          let _index316 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2112_v34).length));
          (_2112_v34)[_index316] = (_this).f19;
          let _2114_v35;
          _2114_v35 = _module.D4.create_DC13(true, (_this).f19, p2, p2, (_this).f19);
          let _2115_v36;
          _2115_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f19);
          let _index317 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2112_v34).length));
          let _rhs273 = (_this).f19;
          let _rhs274 = (((_this).f19) ? ((_this).f19) : ((_this).f19));
          let _rhs275 = (((_2115_v36).contains(p2)) ? ((_2115_v36).get(p2)) : (_module.__default.fm2((_this).f19, globalState)));
          let _rhs276 = _module.D4.create_DC13((_this).f19, (_this).f19, (_dafny.ZERO).minus(p2), new BigNumber(894), !((_this).f19));
          let _lhs165 = _2112_v34;
          let _lhs166 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_2112_v34).length));
          let _lhs167 = globalState;
          r0 = _rhs273;
          _lhs165[_lhs166] = _rhs274;
          _lhs167.f4 = _rhs275;
          _2114_v35 = _rhs276;
          let _2116_v37;
          _2116_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Set.fromElements((_2112_v34)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_2112_v34).length))])).length));
          (globalState).f4 = !(((((_2116_v37).contains(p2)) ? ((_2116_v37).get(p2)) : (p2))).isLessThanOrEqualTo(p2));
          r0 = (_2112_v34)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_2112_v34).length))];
          let _2117_v38;
          let _nw330 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _2117_v38 = _nw330;
          let _2118_v39;
          _2118_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2117_v38,p2);
          _2118_v39 = (_2118_v39).update(_2117_v38, (p2).multipliedBy(p2));
          let _2119_v40;
          let _nw331 = Array((_dafny.ONE).toNumber()).fill([]);
          _2119_v40 = _nw331;
          let _index318 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2119_v40).length));
          (_2119_v40)[_index318] = _2112_v34;
          let _index319 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2119_v40).length));
          (_2119_v40)[_index319] = _2112_v34;
        } else {
          (globalState).f4 = (((_this).f14)[_module.__default.safeIndex(p2, new BigNumber(((_this).f14).length))]) && ((_this).f19);
          let _2120_v41;
          _2120_v41 = _dafny.Map.Empty.slice().updateUnsafe((_module.D5.create_DC16((_this).f19)).dtor_cf26,new BigNumber((p1).length));
          let _2121_v42;
          _2121_v42 = _dafny.MultiSet.fromElements(new BigNumber(((_2120_v41).update((_this).f19, p2)).length), p2, p2);
          _2121_v42 = _2121_v42;
          let _2122_v43;
          let _nw332 = new _module.C1();
          _nw332.__ctor((_this).f19, (_this).f19, _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19)), _dafny.Seq.update((_this).f14, _module.__default.safeIndex(p2, new BigNumber(((_this).f14).length)), !(_module.__default.fm2(true, globalState))), _this.f12);
          _2122_v43 = _nw332;
          let _2123_v44;
          _2123_v44 = new BigNumber(199);
          _2123_v44 = p2;
          let _2124_v45;
          let _nw333 = new _module.C4();
          _nw333.__ctor();
          _2124_v45 = _nw333;
          let _2125_v46;
          _2125_v46 = _dafny.Seq.of(_2124_v45, _2124_v45, _2124_v45);
          let _2126_v47;
          _2126_v47 = _dafny.Map.Empty.slice().updateUnsafe((_2122_v43).f17,((!((_2122_v43).f17)) ? ((_2125_v46)[_module.__default.safeIndex(p2, new BigNumber((_2125_v46).length))]) : (_2124_v45)));
          _2124_v45 = (((_2126_v47).contains((_this).f19)) ? ((_2126_v47).get((_this).f19)) : (_2124_v45));
        }
        let _2127_v48;
        _2127_v48 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p2));
        let _2128_v49;
        _2128_v49 = _dafny.Seq.of(_2127_v48, _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), ((_2129_p2) => function (_2130_i7) {
          return (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.update((_this).f14, _module.__default.safeIndex(_2129_p2, new BigNumber(((_this).f14).length)), (_this).f19), _module.__default.safeIndex(new BigNumber(960), new BigNumber((_dafny.Seq.update((_this).f14, _module.__default.safeIndex(_2129_p2, new BigNumber(((_this).f14).length)), (_this).f19)).length)), (_this).f19)).length))));
        })(p2))));
        _2128_v49 = _2128_v49;
        let _2131_v50;
        _2131_v50 = _module.D5.create_DC16((_this).f19);
        let _2132_v51;
        _2132_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2111_cf11,_2131_v50);
        _2132_v51 = (_2132_v51).update((_2111_cf11).Merge(_2111_cf11), _2131_v50);
        let _2133_v52;
        let _nw334 = Array((new BigNumber(24)).toNumber()).fill(false);
        _2133_v52 = _nw334;
        let _index320 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_2133_v52).length));
        (_2133_v52)[_index320] = true;
        let _index321 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_2133_v52).length));
        (_2133_v52)[_index321] = (_this).f19;
      } else {
        let _2134___mcc_h6 = (_source33).cf17;
        let _2135_cf17 = _2134___mcc_h6;
        let _2136_v53;
        let _nw335 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _2136_v53 = _nw335;
        let _index322 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length));
        (_2136_v53)[_index322] = (_dafny.ZERO).minus(new BigNumber((p1).length));
        let _2137_v54;
        _2137_v54 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f19),p2);
        let _2138_v55;
        _2138_v55 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2137_v54);
        let _index323 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length));
        (_2136_v53)[_index323] = _module.__default.fm0((_this).f19, (((_this).f19) ? (_module.__default.fm62(new BigNumber(355), globalState)) : (_2138_v55)), (_this).f19, globalState);
        let _2139_v56;
        _2139_v56 = _dafny.Seq.of(new BigNumber(-570));
        let _2140_v57;
        _2140_v57 = _dafny.Set.fromElements(_dafny.Seq.Concat(_module.__default.fm1((_this).f19, (_this).f19, globalState), _2139_v56), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-499), (_2136_v53)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length))], (_2136_v53)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length))], p2, p2), _dafny.Seq.of((_2136_v53)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length))], (_2136_v53)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length))], (_2136_v53)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length))])), _2139_v56);
        _2140_v57 = _2140_v57;
        let _2141_v58;
        _2141_v58 = _module.D0.create_DC2(new BigNumber(455), (_this).f19);
        let _2142_v59;
        let _2143_v60;
        let _2144_v61;
        let _2145_v62;
        let _out111;
        let _out112;
        let _out113;
        let _out114;
        let _outcollector29 = _module.__default.m0(!((_this).f19), (((_this).f19) ? (_2141_v58) : (_2141_v58)), true, globalState);
        _out111 = _outcollector29[0];
        _out112 = _outcollector29[1];
        _out113 = _outcollector29[2];
        _out114 = _outcollector29[3];
        _2142_v59 = _out111;
        _2143_v60 = _out112;
        _2144_v61 = _out113;
        _2145_v62 = _out114;
        let _index324 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length));
        let _index325 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length));
        let _rhs277 = _2144_v61;
        let _rhs278 = (_this).f19;
        let _rhs279 = _2144_v61;
        let _lhs168 = _2136_v53;
        let _lhs169 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length));
        let _lhs170 = _2136_v53;
        let _lhs171 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_2136_v53).length));
        _lhs168[_lhs169] = _rhs277;
        _2145_v62 = _rhs278;
        _lhs170[_lhs171] = _rhs279;
      }
      let _2146_i8;
      _2146_i8 = _dafny.ZERO;
      L12: {
        while ((p2).isLessThanOrEqualTo(p2)) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2146_i8)) {
              break L12;
            }
            _2146_i8 = (_2146_i8).plus(_dafny.ONE);
            r0 = !((p2).isEqualTo(p2));
            let _2147_v63;
            let _nw336 = Array((new BigNumber(2)).toNumber()).fill(_module.D12.Default());
            _2147_v63 = _nw336;
            let _2148_v64;
            _2148_v64 = _dafny.Set.fromElements(_2147_v63, _2147_v63, _2147_v63);
            _2148_v64 = (_2148_v64).Union(_2148_v64);
            let _2149_v65;
            _2149_v65 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f19);
            let _2150_v66;
            _2150_v66 = _module.D7.create_DC21(_2149_v65, !((_this).f19) || ((_this).f19), !(((_this).f19) === ((_this).f19)));
            _2150_v66 = _2150_v66;
            let _2151_v67;
            _2151_v67 = new BigNumber(443);
            _2151_v67 = new BigNumber(126);
          }
        }
      }
      let _2152_v68;
      _2152_v68 = new _dafny.CodePoint('u'.codePointAt(0));
      let _2153_v69;
      _2153_v69 = _dafny.Map.Empty.slice().updateUnsafe(_2152_v68,(_this).f19);
      let _2154_v70;
      _2154_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,new BigNumber((_2153_v69).length));
      let _2155_v71;
      _2155_v71 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2154_v70);
      let _2156_v72;
      _2156_v72 = _dafny.MultiSet.fromElements(p2);
      if (((_2156_v72).update(p2, _module.__default.abs(p2))).IsSubsetOf((_dafny.MultiSet.fromElements(_module.__default.fm0(!((_this).f19), _2155_v71, (_this).f19, globalState), p2, new BigNumber(899), new BigNumber((p1).length))).Union(_2156_v72))) {
        r0 = (_this).f19;
        let _2157_v73;
        let _init57 = function (_2158_i9) {
          return false;
        };
        let _nw337 = Array((new BigNumber(24)).toNumber());
        for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw337.length); _i0_57++) {
          _nw337[_i0_57] = _init57(new BigNumber(_i0_57));
        }
        _2157_v73 = _nw337;
        let _index326 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length));
        (_2157_v73)[_index326] = (_this).f19;
        let _index327 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length));
        (_2157_v73)[_index327] = ((_dafny.ZERO).minus(p2)).isEqualTo(p2);
        if ((_2157_v73)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length))]) {
          let _2159_v74;
          _2159_v74 = _module.D0.create_DC2(p2, _module.__default.fm2(false, globalState));
          let _2160_v75;
          _2160_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
          let _2161_v76;
          let _nw338 = new _module.C1();
          _nw338.__ctor((_2157_v73)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length))], (_2157_v73)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length))], ((_this).f13).update((_this).f19, _2160_v75), (_this).f14, _this.f12);
          _2161_v76 = _nw338;
          let _2162_v77;
          _2162_v77 = _dafny.Seq.of(_2161_v76);
          let _2163_v78;
          _2163_v78 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _2164_v79;
          _2164_v79 = _dafny.Set.fromElements((_2162_v77)[_module.__default.safeIndex(new BigNumber((_2163_v78).length), new BigNumber((_2162_v77).length))]);
          let _2165_v80;
          _2165_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(706),new BigNumber((_2164_v79).length));
          let _2166_v81;
          let _2167_v82;
          let _2168_v83;
          let _2169_v84;
          let _out115;
          let _out116;
          let _out117;
          let _out118;
          let _outcollector30 = _module.__default.m0((_this).f19, _2159_v74, !(_2163_v78).equals(_2165_v80), globalState);
          _out115 = _outcollector30[0];
          _out116 = _outcollector30[1];
          _out117 = _outcollector30[2];
          _out118 = _outcollector30[3];
          _2166_v81 = _out115;
          _2167_v82 = _out116;
          _2168_v83 = _out117;
          _2169_v84 = _out118;
          let _2170_v85;
          _2170_v85 = _dafny.Seq.UnicodeFromString("xaj");
          let _2171_v86;
          _2171_v86 = _module.D12.create_DC37(_dafny.Seq.UnicodeFromString("ruokmp"), _2169_v84);
          _2170_v85 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("unebn"), (function (_pat_let37_0) {
            return function (_2172_dt__update__tmp_h2) {
              return function (_pat_let38_0) {
                return function (_2173_dt__update_hcf60_h0) {
                  return _module.D12.create_DC37((_2172_dt__update__tmp_h2).dtor_cf59, _2173_dt__update_hcf60_h0);
                }(_pat_let38_0);
              }(true);
            }(_pat_let37_0);
          }(_2171_v86)).dtor_cf59);
          let _2174_v87;
          _2174_v87 = _dafny.Set.fromElements(_2157_v73);
          r0 = (_2174_v87).IsDisjointFrom((_2174_v87).Union(_2174_v87));
          _2166_v81 = _2166_v81;
          _2163_v78 = (_2163_v78).update(new BigNumber((_2154_v70).length), (_dafny.ZERO).minus(p2));
        } else {
          let _index328 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length));
          (_2157_v73)[_index328] = _module.__default.fm2(_module.__default.fm2(!((_2157_v73)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length))]), globalState), globalState);
          let _2175_v88;
          _2175_v88 = new BigNumber(350);
          _2175_v88 = (_this).fm5(!((_this).f19), globalState);
          let _2176_v89;
          let _nw339 = new _module.C3();
          _nw339.__ctor(true, _this.f12, (_2157_v73)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_2157_v73).length))], (_this).f13, (_this).f14);
          _2176_v89 = _nw339;
          (globalState).f4 = !(!(((_this).f14)[_module.__default.safeIndex(new BigNumber(22), new BigNumber(((_this).f14).length))]));
          let _2177_v90;
          _2177_v90 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _2178_v91;
          _2178_v91 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2175_v88,new BigNumber((p1).length)), _2177_v90, _2177_v90);
          let _2179_v92;
          let _nw340 = Array((new BigNumber(4)).toNumber());
          _2179_v92 = _nw340;
          let _2180_v93;
          let _nw341 = new _module.C1();
          _nw341.__ctor((_2176_v89).f20, true, (_this).f13, (_this).f14, _this.f12);
          _2180_v93 = _nw341;
          let _index329 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2179_v92).length));
          (_2179_v92)[_index329] = _2180_v93;
          let _2181_v94;
          _2181_v94 = _dafny.MultiSet.fromElements((_2156_v72).Union(_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(_2175_v88)))), _2156_v72);
          let _2182_v95;
          _2182_v95 = _dafny.Seq.of((_2180_v93).fm5(false, globalState), _2175_v88);
          let _index330 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2179_v92).length));
          let _rhs280 = ((_this).f14)[_module.__default.safeIndex(new BigNumber(260), new BigNumber(((_this).f14).length))];
          let _rhs281 = _dafny.Seq.of(_2177_v90, _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),_2175_v88), (_2177_v90).Merge(_2177_v90), _2177_v90, _2177_v90);
          let _rhs282 = _2180_v93;
          let _rhs283 = ((((_this).f19) || (false)) ? (_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_2182_v95), _2156_v72)) : (_2181_v94));
          let _rhs284 = _2157_v73;
          let _lhs172 = globalState;
          let _lhs173 = _2179_v92;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_2179_v92).length));
          _lhs172.f4 = _rhs280;
          _2178_v91 = _rhs281;
          _lhs173[_lhs174] = _rhs282;
          _2181_v94 = _rhs283;
          _2157_v73 = _rhs284;
        }
        let _2183_v96;
        _2183_v96 = new BigNumber(938);
        _2183_v96 = (new BigNumber((p1).length)).plus(p2);
        let _2184_v97;
        _2184_v97 = _dafny.Seq.UnicodeFromString("d");
        _2184_v97 = _dafny.Seq.Concat(_2184_v97, _dafny.Seq.Concat(p1, _2184_v97));
      } else {
        r0 = !(p2).isEqualTo(p2);
        let _2185_v98;
        _2185_v98 = new BigNumber(336);
        let _2186_v99;
        _2186_v99 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _2187_v100;
        _2187_v100 = _dafny.Seq.of((((_2186_v99).contains(p2)) ? ((_2186_v99).get(p2)) : (p2)), p2);
        _2185_v98 = (_2187_v100)[_module.__default.safeIndex((new BigNumber((p1).length)).plus(new BigNumber(((_this).f14).length)), new BigNumber((_2187_v100).length))];
        let _2188_v101;
        _2188_v101 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat((_this).f14, (_this).f14),p1);
        _2188_v101 = (_2188_v101).update(((!((_this).f19)) ? ((_this).f14) : ((_this).f14)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), ((_2189_v68) => function (_2190_i10) {
          return _2189_v68;
        })(_2152_v68)));
        _2185_v98 = p2;
        let _2191_v102;
        let _init58 = ((_2192_p1) => function (_2193_i11) {
          return _2192_p1;
        })(p1);
        let _nw342 = Array((new BigNumber(8)).toNumber());
        for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw342.length); _i0_58++) {
          _nw342[_i0_58] = _init58(new BigNumber(_i0_58));
        }
        _2191_v102 = _nw342;
        let _2194_v104;
        _2194_v104 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cpwcbgmj"));
        let _2195_v105;
        let _init59 = function (_2196_i12) {
          return (_this).f19;
        };
        let _nw343 = Array((new BigNumber(17)).toNumber());
        for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw343.length); _i0_59++) {
          _nw343[_i0_59] = _init59(new BigNumber(_i0_59));
        }
        _2195_v105 = _nw343;
        let _2197_v106;
        _2197_v106 = _dafny.Set.fromElements(_2195_v105);
        let _2198_v107;
        _2198_v107 = _module.D12.create_DC36(_2191_v102, _module.__default.fm0((_this).f19, (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-214),_2154_v70)).update(p2, _2154_v70), (_this).f19, globalState), function () {
  let _coll66 = new _dafny.Map();
  for (const _compr_66 of (_2194_v104).Elements) {
    let _2199_v103 = _compr_66;
    if (_dafny.Seq.contains(_2194_v104, _2199_v103)) {
      _coll66.push([_2199_v103,_2187_v100]);
    }
  }
  return _coll66;
}(), _2197_v106);
        let _2200_v108;
        _2200_v108 = _dafny.Map.Empty.slice().updateUnsafe(_2185_v98,_2198_v107);
        let _2201_v109;
        _2201_v109 = _dafny.MultiSet.fromElements((_this).f19);
        let _2202_v110;
        _2202_v110 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,false);
        let _2203_v111;
        let _nw344 = Array((new BigNumber(14)).toNumber());
        _nw344[(_dafny.ZERO).toNumber()] = _2185_v98;
        _nw344[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f19,p2)).length);
        _nw344[(new BigNumber(2)).toNumber()] = new BigNumber(-405);
        _nw344[(new BigNumber(3)).toNumber()] = _2185_v98;
        _nw344[(new BigNumber(4)).toNumber()] = (((_2154_v70).contains((_this).f19)) ? ((_2154_v70).get((_this).f19)) : (_2185_v98));
        _nw344[(new BigNumber(5)).toNumber()] = new BigNumber(((_this).f14).length);
        _nw344[(new BigNumber(6)).toNumber()] = (((_2186_v99).contains(_2185_v98)) ? ((_2186_v99).get(_2185_v98)) : (_2185_v98));
        _nw344[(new BigNumber(7)).toNumber()] = p2;
        _nw344[(new BigNumber(8)).toNumber()] = (((_2186_v99).contains(_2185_v98)) ? ((_2186_v99).get(_2185_v98)) : (new BigNumber((_2200_v108).length)));
        _nw344[(new BigNumber(9)).toNumber()] = (_2185_v98).minus((((_2201_v109).contains(false)) ? ((_2201_v109).get(false)) : (new BigNumber((p1).length))));
        _nw344[(new BigNumber(10)).toNumber()] = new BigNumber((_2202_v110).length);
        _nw344[(new BigNumber(11)).toNumber()] = p2;
        _nw344[(new BigNumber(12)).toNumber()] = (new BigNumber(((_this).f14).length)).plus(new BigNumber((_2187_v100).length));
        _nw344[(new BigNumber(13)).toNumber()] = p2;
        _2203_v111 = _nw344;
        let _2204_v112;
        _2204_v112 = _dafny.Map.Empty.slice().updateUnsafe(_2185_v98,_2195_v105);
        let _rhs285 = _2185_v98;
        let _rhs286 = _2185_v98;
        let _rhs287 = ((((_2201_v109).contains((_this).f19)) ? ((_2201_v109).get((_this).f19)) : (new BigNumber((_2204_v112).length)))).minus(new BigNumber((_2186_v99).length));
        let _rhs288 = _2203_v111;
        let _rhs289 = _2152_v68;
        _2185_v98 = _rhs285;
        _2185_v98 = _rhs286;
        _2185_v98 = _rhs287;
        _2203_v111 = _rhs288;
        _2152_v68 = _rhs289;
      }
      let _2205_i13;
      _2205_i13 = _dafny.ZERO;
      L13: {
        while (true) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2205_i13)) {
              break L13;
            }
            _2205_i13 = (_2205_i13).plus(_dafny.ONE);
            let _2206_v113;
            _2206_v113 = new BigNumber(997);
            _2206_v113 = _2206_v113;
            let _2207_v114;
            _2207_v114 = _module.D4.create_DC14((_this).f19);
            let _2208_v115;
            _2208_v115 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_module.__default.fm40(_2207_v114, new BigNumber(373), (_this).f19, (_this).f19, globalState));
            let _2209_v116;
            _2209_v116 = _dafny.Seq.of((_this).f14, (_this).f14, (_this).f14, (_this).f14, (_this).f14);
            let _2210_v117;
            _2210_v117 = _dafny.Set.fromElements((_2209_v116)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_2209_v116).length))]);
            _2208_v115 = (_2208_v115).update((p2).isLessThan(new BigNumber((_2210_v117).length)), p1);
            let _2211_v118;
            _2211_v118 = _dafny.Seq.of(_2152_v68);
            let _2212_v119;
            _2212_v119 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(((_2153_v69).contains(_2152_v68)) ? ((_2153_v69).get(_2152_v68)) : ((_this).f19)));
            _2211_v118 = _dafny.Seq.Concat(_module.__default.fm29((_this).f19, (_this).f19, (_this).f19, _2206_v113, globalState), (((_this).f19) ? (_module.__default.fm40(_2207_v114, p2, (((_2212_v119).contains((_this).f19)) ? ((_2212_v119).get((_this).f19)) : ((_this).f19)), (_this).f19, globalState)) : (_dafny.Seq.of(_2152_v68, _2152_v68, _2152_v68))));
            if ((((_this).f19) ? (!((_this).f19) || (false)) : (true))) {
              _2206_v113 = _2206_v113;
              let _2213_v120;
              _2213_v120 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jchc")).length)), p2, new BigNumber((_2153_v69).length));
              (globalState).f4 = ((_this).f19) === (_dafny.areEqual(_2213_v120, _dafny.Seq.of(p2)));
              (globalState).f4 = (_this).f19;
              let _2214_v121;
              let _nw345 = new _module.C1();
              _nw345.__ctor((_this).f19, (_this).f19, (_this).f13, _dafny.Seq.Concat((_this).f14, _dafny.Seq.of((_this).f19)), _this.f12);
              _2214_v121 = _nw345;
              let _2215_v122;
              let _nw346 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _2215_v122 = _nw346;
              let _index331 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_2215_v122).length));
              (_2215_v122)[_index331] = p2;
              let _index332 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_2215_v122).length));
              (_2215_v122)[_index332] = new BigNumber(((_2156_v72).Intersect(_2156_v72)).cardinality());
            } else {
              let _2216_v123;
              let _nw347 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
              _2216_v123 = _nw347;
              let _2217_v124;
              let _nw348 = Array((new BigNumber(6)).toNumber());
              _nw348[(_dafny.ZERO).toNumber()] = _2216_v123;
              _nw348[(_dafny.ONE).toNumber()] = _2216_v123;
              _nw348[(new BigNumber(2)).toNumber()] = _2216_v123;
              _nw348[(new BigNumber(3)).toNumber()] = _2216_v123;
              _nw348[(new BigNumber(4)).toNumber()] = _2216_v123;
              _nw348[(new BigNumber(5)).toNumber()] = _2216_v123;
              _2217_v124 = _nw348;
              let _index333 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2217_v124).length));
              (_2217_v124)[_index333] = _2216_v123;
              let _index334 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_2217_v124).length));
              (_2217_v124)[_index334] = _2216_v123;
              let _2218_v125;
              let _nw349 = new _module.C1();
              _nw349.__ctor((_this).f19, (_this).f19, (_this).f13, (_this).f14, _this.f12);
              _2218_v125 = _nw349;
              let _2219_v126;
              let _nw350 = Array((new BigNumber(15)).toNumber()).fill(false);
              _2219_v126 = _nw350;
              let _index335 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2219_v126).length));
              (_2219_v126)[_index335] = !_dafny.Seq.contains(p1, _2152_v68);
              let _2220_v127;
              let _nw351 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _2220_v127 = _nw351;
              let _index336 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_2220_v127).length));
              (_2220_v127)[_index336] = _2211_v118;
              let _index337 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2219_v126).length));
              let _index338 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_2220_v127).length));
              let _rhs290 = _module.__default.fm56(new _dafny.CodePoint('a'.codePointAt(0)), globalState);
              let _rhs291 = _module.__default.safeDivisionInt(_2206_v113, _2206_v113);
              let _rhs292 = (((_2218_v125).f17) ? ((_2218_v125).f17) : (!((_2218_v125).f17)));
              let _rhs293 = (_this).f19;
              let _rhs294 = p1;
              let _lhs175 = _2219_v126;
              let _lhs176 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2219_v126).length));
              let _lhs177 = _2220_v127;
              let _lhs178 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_2220_v127).length));
              _2156_v72 = _rhs290;
              _2206_v113 = _rhs291;
              r0 = _rhs292;
              _lhs175[_lhs176] = _rhs293;
              _lhs177[_lhs178] = _rhs294;
              let _2221_v128;
              let _nw352 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
              _2221_v128 = _nw352;
              let _index339 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_2221_v128).length));
              (_2221_v128)[_index339] = _module.__default.safeModuloInt(_2206_v113, _2206_v113);
              let _2222_v129;
              _2222_v129 = _dafny.Set.fromElements(_2219_v126);
              let _index340 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_2221_v128).length));
              let _rhs295 = _2206_v113;
              let _rhs296 = (new BigNumber(((_2222_v129).Union(_dafny.Set.fromElements(_2219_v126))).length)).isLessThanOrEqualTo(((!((_this).f19)) ? (_2206_v113) : (new BigNumber(981))));
              let _lhs179 = _2221_v128;
              let _lhs180 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_2221_v128).length));
              let _lhs181 = globalState;
              _lhs179[_lhs180] = _rhs295;
              _lhs181.f4 = _rhs296;
              let _index341 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_2219_v126).length));
              (_2219_v126)[_index341] = _module.__default.fm2(_2218_v125.f18, globalState);
            }
          }
        }
      }
      (globalState).f4 = true;
      let _2223_v130;
      _2223_v130 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f19);
      let _2224_v131;
      _2224_v131 = _module.D7.create_DC21(_2223_v130, (p2).isEqualTo(p2), (!((_this).f19)) && ((_this).f19));
      _2224_v131 = _2224_v131;
      r0 = (_this).f19;
      return r0;
    }
    m6(globalState) {
      let _this = this;
      let r0 = false;
      if (!((_this).f19) || ((_this).f19)) {
        let _2225_v0;
        _2225_v0 = _module.D8.create_DC26((_this).f19);
        let _source35 = function (_pat_let39_0) {
          return function (_2228_dt__update__tmp_h1) {
            return function (_pat_let42_0) {
              return function (_2229_dt__update_hcf41_h1) {
                return _module.D8.create_DC26(_2229_dt__update_hcf41_h1);
              }(_pat_let42_0);
            }((_this).f19);
          }(_pat_let39_0);
        }(function (_pat_let40_0) {
          return function (_2226_dt__update__tmp_h0) {
            return function (_pat_let41_0) {
              return function (_2227_dt__update_hcf41_h0) {
                return _module.D8.create_DC26(_2227_dt__update_hcf41_h0);
              }(_pat_let41_0);
            }((_this).f19);
          }(_pat_let40_0);
        }(_2225_v0));
        if (_source35.is_DC24) {
          let _2230___mcc_h0 = (_source35).cf36;
          let _2231_cf36 = _2230___mcc_h0;
          _2231_cf36 = !(_2231_cf36);
          let _2232_v1;
          _2232_v1 = _dafny.Seq.UnicodeFromString("pkgacn");
          _2232_v1 = _2232_v1;
          (globalState).f4 = _2231_cf36;
          let _2233_v2;
          _2233_v2 = _module.D0.create_DC0(!((_this).f19));
          r0 = (_2233_v2).dtor_cf0;
        } else if (_source35.is_DC25) {
          let _2234___mcc_h1 = (_source35).cf37;
          let _2235___mcc_h2 = (_source35).cf38;
          let _2236___mcc_h3 = (_source35).cf39;
          let _2237___mcc_h4 = (_source35).cf40;
          let _2238_cf40 = _2237___mcc_h4;
          let _2239_cf39 = _2236___mcc_h3;
          let _2240_cf38 = _2235___mcc_h2;
          let _2241_cf37 = _2234___mcc_h1;
          _2240_cf38 = _2240_cf38;
          (globalState).f4 = _module.__default.fm2(_module.__default.fm2((_this).f19, globalState), globalState);
          let _2242_v3;
          let _init60 = function (_2243_i0) {
            return true;
          };
          let _nw353 = Array((new BigNumber(11)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw353.length); _i0_60++) {
            _nw353[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _2242_v3 = _nw353;
          let _2244_v4;
          _2244_v4 = _dafny.Set.fromElements(_2242_v3, _2242_v3);
          _2239_cf39 = new BigNumber(((_2244_v4).Difference((_2244_v4).Union(_2244_v4))).length);
          (globalState).f4 = !(((_this).f19) && ((_this).f19));
        } else if (_source35.is_DC26) {
          let _2245___mcc_h5 = (_source35).cf41;
          let _2246_cf41 = _2245___mcc_h5;
          let _2247_v5;
          _2247_v5 = new BigNumber(824);
          _2247_v5 = _2247_v5;
          let _2248_v6;
          _2248_v6 = new _dafny.CodePoint('i'.codePointAt(0));
          _2248_v6 = _2248_v6;
          let _2249_v7;
          _2249_v7 = _dafny.Set.fromElements(_2246_cf41);
          let _2250_v8;
          _2250_v8 = _2249_v7;
          let _rhs297 = (_2249_v7).IsDisjointFrom((_2250_v8));
          let _rhs298 = _module.__default.safeDivisionInt(_2247_v5, _module.__default.safeModuloInt(_2247_v5, new BigNumber(41)));
          let _lhs182 = globalState;
          _lhs182.f4 = _rhs297;
          _2247_v5 = _rhs298;
          let _2251_v9;
          let _nw354 = new _module.C3();
          _nw354.__ctor(_2246_cf41, _this.f12, _2246_cf41, (_this).f13, (_this).f14);
          _2251_v9 = _nw354;
          let _2252_v10;
          _2252_v10 = _module.D14.create_DC40(_2251_v9);
          let _2253_v11;
          _2253_v11 = _dafny.Map.Empty.slice().updateUnsafe(true,_2251_v9);
          let _pat_let_tv36 = _2253_v11;
          let _pat_let_tv37 = _2251_v9;
          let _pat_let_tv38 = _2253_v11;
          _2252_v10 = function (_pat_let43_0) {
            return function (_2254_dt__update__tmp_h2) {
              return function (_pat_let44_0) {
                return function (_2255_dt__update_hcf63_h0) {
                  return _module.D14.create_DC40(_2255_dt__update_hcf63_h0);
                }(_pat_let44_0);
              }((((_pat_let_tv38).contains((_this).f19)) ? ((_pat_let_tv36).get((_this).f19)) : (_pat_let_tv37)));
            }(_pat_let43_0);
          }(_2252_v10);
        } else if (_source35.is_DC23) {
          let _2256___mcc_h6 = (_source35).cf35;
          let _2257_cf35 = _2256___mcc_h6;
          let _2258_v12;
          let _init61 = function (_2259_i1) {
            return (_this).f19;
          };
          let _nw355 = Array((new BigNumber(18)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw355.length); _i0_61++) {
            _nw355[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _2258_v12 = _nw355;
          let _index342 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_2258_v12).length));
          (_2258_v12)[_index342] = (((_this).f19) ? ((_this).f19) : ((_this).f19));
          let _index343 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_2258_v12).length));
          (_2258_v12)[_index343] = (_this).f19;
          let _2260_v13;
          let _nw356 = new _module.C0();
          _nw356.__ctor();
          _2260_v13 = _nw356;
          let _2261_v14;
          _2261_v14 = new BigNumber(-413);
          let _2262_v15;
          _2262_v15 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2261_v14));
          _2261_v14 = (_2261_v14).multipliedBy(new BigNumber((_2262_v15).cardinality()));
          let _2263_v16;
          _2263_v16 = _dafny.Seq.UnicodeFromString("mhlgooi");
          let _2264_v17;
          _2264_v17 = _dafny.Seq.of(_dafny.Seq.IsPrefixOf(_2263_v16, _2263_v16));
          _2264_v17 = (_this).f14;
        } else {
          let _2265___mcc_h7 = (_source35).cf42;
          let _2266_cf42 = _2265___mcc_h7;
          let _2267_v18;
          _2267_v18 = new _dafny.CodePoint('d'.codePointAt(0));
          _2267_v18 = _2267_v18;
          (globalState).f4 = (_this).f19;
          let _2268_v19;
          _2268_v19 = new BigNumber(-786);
          _2268_v19 = new BigNumber(756);
          let _2269_v20;
          let _nw357 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
          _2269_v20 = _nw357;
          let _2270_v21;
          _2270_v21 = _dafny.MultiSet.fromElements(_2268_v19);
          let _2271_v22;
          _2271_v22 = _dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(_2268_v19));
          let _2272_v23;
          _2272_v23 = _dafny.Set.fromElements(_2268_v19);
          let _2273_v24;
          _2273_v24 = _dafny.Seq.UnicodeFromString("gb");
          let _2274_v25;
          _2274_v25 = _dafny.MultiSet.fromElements((_this).f19, (_this).f19);
          let _2275_v26;
          _2275_v26 = _dafny.Seq.of(_2268_v19);
          let _2276_v27;
          _2276_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_2267_v18);
          let _2277_v28;
          _2277_v28 = _dafny.MultiSet.fromElements(_2276_v27, _2276_v27, _2276_v27, _2276_v27, _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_2267_v18));
          let _2278_v29;
          let _nw358 = Array((new BigNumber(27)).toNumber());
          _nw358[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements(_2268_v19)).length);
          _nw358[(_dafny.ONE).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(2)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(3)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(4)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(5)).toNumber()] = new BigNumber((_2270_v21).cardinality());
          _nw358[(new BigNumber(6)).toNumber()] = new BigNumber(618);
          _nw358[(new BigNumber(7)).toNumber()] = (((_2271_v22).contains(!((_this).f19))) ? ((_2271_v22).get(!((_this).f19))) : (_2268_v19));
          _nw358[(new BigNumber(8)).toNumber()] = new BigNumber(303);
          _nw358[(new BigNumber(9)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(10)).toNumber()] = new BigNumber((_2272_v23).length);
          _nw358[(new BigNumber(11)).toNumber()] = new BigNumber(851);
          _nw358[(new BigNumber(12)).toNumber()] = new BigNumber(582);
          _nw358[(new BigNumber(13)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(14)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(15)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(16)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(17)).toNumber()] = new BigNumber((_2273_v24).length);
          _nw358[(new BigNumber(18)).toNumber()] = (((_2274_v25).contains(!((_this).f19))) ? ((_2274_v25).get(!((_this).f19))) : (new BigNumber((_2275_v26).length)));
          _nw358[(new BigNumber(19)).toNumber()] = new BigNumber(((_2274_v25).update((_this).f19, _module.__default.abs((_dafny.ZERO).minus((((_2271_v22).contains(false)) ? ((_2271_v22).get(false)) : (new BigNumber((_2277_v28).cardinality()))))))).cardinality());
          _nw358[(new BigNumber(20)).toNumber()] = new BigNumber((_2275_v26).length);
          _nw358[(new BigNumber(21)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(22)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(23)).toNumber()] = new BigNumber(819);
          _nw358[(new BigNumber(24)).toNumber()] = new BigNumber((_2273_v24).length);
          _nw358[(new BigNumber(25)).toNumber()] = _2268_v19;
          _nw358[(new BigNumber(26)).toNumber()] = _2268_v19;
          _2278_v29 = _nw358;
          let _2279_v30;
          _2279_v30 = _dafny.Set.fromElements(_2278_v29);
          let _index344 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_2269_v20).length));
          (_2269_v20)[_index344] = (((_this).f19) ? (_2279_v30) : (_2279_v30));
          let _index345 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_2269_v20).length));
          (_2269_v20)[_index345] = (_2279_v30).Union(_2279_v30);
        }
        let _2280_v31;
        let _nw359 = new _module.C4();
        _nw359.__ctor();
        _2280_v31 = _nw359;
        let _2281_v32;
        _2281_v32 = new BigNumber(522);
        let _2282_v33;
        _2282_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2281_v32,_2281_v32);
        _2282_v33 = (_2282_v33).update((_dafny.ZERO).minus(_2281_v32), _2281_v32);
        let _2283_v34;
        _2283_v34 = new _dafny.CodePoint('t'.codePointAt(0));
        let _2284_v35;
        _2284_v35 = _dafny.Seq.UnicodeFromString("uonlgggk");
        let _2285_v36;
        _2285_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2283_v34,_2284_v35);
        let _2286_v37;
        _2286_v37 = _dafny.Set.fromElements(_2280_v31);
        let _2287_v38;
        _2287_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,true);
        let _2288_v39;
        _2288_v39 = _dafny.Seq.of(_2281_v32);
        let _2289_v40;
        _2289_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_2288_v39)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_2283_v34)).length), new BigNumber((_2288_v39).length))]);
        let _2290_v41;
        _2290_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f14).length),_2289_v40);
        let _2291_v42;
        _2291_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2281_v32,(_this).f19);
        let _2292_v43;
        _2292_v43 = _module.D7.create_DC21(_2291_v42, !((_this).f19), (_this).f19);
        let _2293_v44;
        _2293_v44 = _dafny.Set.fromElements(_2281_v32, (_dafny.ZERO).minus(_2281_v32));
        let _2294_v45;
        let _nw360 = Array((new BigNumber(21)).toNumber());
        _nw360[(_dafny.ZERO).toNumber()] = _2281_v32;
        _nw360[(_dafny.ONE).toNumber()] = new BigNumber((_2286_v37).length);
        _nw360[(new BigNumber(2)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(3)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(4)).toNumber()] = _module.__default.fm0((((_2287_v38).contains((_this).f19)) ? ((_2287_v38).get((_this).f19)) : ((_this).f19)), _2290_v41, (_this).f19, globalState);
        _nw360[(new BigNumber(5)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(6)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(7)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(8)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(9)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(10)).toNumber()] = _module.__default.fm0((_this).f19, _2290_v41, (_this).f19, globalState);
        _nw360[(new BigNumber(11)).toNumber()] = (((_2282_v33).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2292_v43).dtor_cf32,_2283_v34)).length))) ? ((_2282_v33).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_2292_v43).dtor_cf32,_2283_v34)).length))) : (_2281_v32));
        _nw360[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(_2281_v32);
        _nw360[(new BigNumber(13)).toNumber()] = new BigNumber((_2284_v35).length);
        _nw360[(new BigNumber(14)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(15)).toNumber()] = new BigNumber((_2293_v44).length);
        _nw360[(new BigNumber(16)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(17)).toNumber()] = new BigNumber(((_this).f14).length);
        _nw360[(new BigNumber(18)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(19)).toNumber()] = _2281_v32;
        _nw360[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.update(_2288_v39, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_2284_v35).length)), new BigNumber((_2288_v39).length)), new BigNumber((_2288_v39).length))).length);
        _2294_v45 = _nw360;
        let _2295_v46;
        _2295_v46 = _dafny.Set.fromElements(_2294_v45);
        let _2296_v47;
        _2296_v47 = _dafny.Map.Empty.slice().updateUnsafe((((_2285_v36).contains(_2283_v34)) ? ((_2285_v36).get(_2283_v34)) : (_2284_v35)),_2295_v46);
        _2296_v47 = (_2296_v47).update(_dafny.Seq.UnicodeFromString("hya"), (_2295_v46).Union(_2295_v46));
        _2281_v32 = (_module.__default.fm0((_this).f19, _2290_v41, true, globalState)).minus(new BigNumber((_2288_v39).length));
      } else {
        let _2297_v48;
        let _init62 = function (_2298_i2) {
          return _module.D8.create_DC24((_this).f19);
        };
        let _nw361 = Array((new BigNumber(4)).toNumber());
        for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw361.length); _i0_62++) {
          _nw361[_i0_62] = _init62(new BigNumber(_i0_62));
        }
        _2297_v48 = _nw361;
        let _2299_v49;
        _2299_v49 = _module.D8.create_DC24((_this).f19);
        let _index346 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_2297_v48).length));
        (_2297_v48)[_index346] = _2299_v49;
        let _2300_v50;
        let _nw362 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _2300_v50 = _nw362;
        let _index347 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_2297_v48).length));
        let _rhs299 = _module.D8.create_DC24((_this).f19);
        let _rhs300 = _2300_v50;
        let _lhs183 = _2297_v48;
        let _lhs184 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_2297_v48).length));
        _lhs183[_lhs184] = _rhs299;
        _2300_v50 = _rhs300;
        let _2301_v51;
        let _nw363 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
        _2301_v51 = _nw363;
        let _2302_v52;
        _2302_v52 = new BigNumber(815);
        let _2303_v53;
        _2303_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_2302_v52);
        let _2304_v54;
        _2304_v54 = _dafny.Set.fromElements((((_2303_v53).contains((_this).f19)) ? ((_2303_v53).get((_this).f19)) : (_2302_v52)));
        let _2305_v55;
        _2305_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_module.D6.create_DC17(_2304_v54));
        let _index348 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_2301_v51).length));
        (_2301_v51)[_index348] = _2305_v55;
        let _index349 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_2301_v51).length));
        (_2301_v51)[_index349] = (_2305_v55).Merge(_2305_v55);
        let _2306_v56;
        let _nw364 = new _module.C4();
        _nw364.__ctor();
        _2306_v56 = _nw364;
        let _2307_v57;
        let _init63 = ((_2308_v52) => function (_2309_i3) {
          return (_module.D4.create_DC13((_this).f19, (_this).f19, (_dafny.ZERO).minus(_2308_v52), _2308_v52, (_this).f19)).dtor_cf19;
        })(_2302_v52);
        let _nw365 = Array((new BigNumber(17)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw365.length); _i0_63++) {
          _nw365[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2307_v57 = _nw365;
        let _2310_v58;
        _2310_v58 = _dafny.Set.fromElements((_this).f19);
        let _index350 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2307_v57).length));
        (_2307_v57)[_index350] = (_2310_v58).IsSubsetOf(_2310_v58);
        let _index351 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_2307_v57).length));
        (_2307_v57)[_index351] = (_this).f19;
        let _2311_v59;
        _2311_v59 = _module.D8.create_DC25((_this).f19, _2302_v52, _2302_v52, _2303_v53);
        let _index352 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2300_v50).length));
        (_2300_v50)[_index352] = (_2311_v59).dtor_cf38;
        let _index353 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2300_v50).length));
        (_2300_v50)[_index353] = ((((_this).f19) && (true)) ? (_2302_v52) : (_2302_v52));
      }
      if ((_this).f19) {
        (globalState).f4 = (_this).f19;
        let _2312_v60;
        _2312_v60 = new BigNumber(855);
        let _2313_v61;
        _2313_v61 = _dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)));
        let _2314_v62;
        _2314_v62 = _dafny.Set.fromElements((new BigNumber(351)).multipliedBy(_2312_v60), (_2312_v60).multipliedBy(new BigNumber((_2313_v61).length)), _2312_v60);
        let _2315_v63;
        _2315_v63 = _dafny.Seq.of(_2312_v60, _2312_v60);
        _2314_v62 = (_this).fm10((_this).f19, (_dafny.Set.fromElements(new BigNumber(265), _2312_v60)).Intersect(_2314_v62), _dafny.Seq.Concat(_2315_v63, _2315_v63), _2312_v60, globalState);
        let _2316_v64;
        _2316_v64 = _dafny.Map.Empty.slice().updateUnsafe((_2312_v60).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f19,new BigNumber(489))).length)),_dafny.Seq.Concat((_this).f14, (_this).f14));
        _2316_v64 = (_2316_v64).update(_2312_v60, _dafny.Seq.Concat((_this).f14, (_this).f14));
        let _2317_v65;
        _2317_v65 = new _dafny.CodePoint('p'.codePointAt(0));
        _2317_v65 = _2317_v65;
        (globalState).f4 = (_this).f19;
      } else {
        let _2318_v66;
        _2318_v66 = _dafny.MultiSet.fromElements((_this).f19, (_this).f19);
        let _2319_v67;
        _2319_v67 = _module.D0.create_DC1();
        let _2320_v68;
        _2320_v68 = new BigNumber(421);
        let _source36 = (((_2318_v66).IsProperSubsetOf(_2318_v66)) ? (_2319_v67) : (_module.__default.fm63(_2320_v68, globalState)));
        if (_source36.is_DC1) {
          let _2321_v69;
          _2321_v69 = _dafny.Seq.of(!((_this).f19), !((_this).f19));
          let _2322_v70;
          _2322_v70 = _dafny.MultiSet.fromElements(_2320_v68);
          let _2323_v71;
          _2323_v71 = _dafny.Seq.UnicodeFromString("qwkupqk");
          let _2324_v72;
          _2324_v72 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("pqexl"), _2323_v71);
          let _2325_v73;
          _2325_v73 = _dafny.Seq.of((_2324_v72)[_module.__default.safeIndex(_2320_v68, new BigNumber((_2324_v72).length))], _2323_v71, (_2324_v72)[_module.__default.safeIndex((_dafny.ZERO).minus(_2320_v68), new BigNumber((_2324_v72).length))], _2323_v71);
          let _rhs301 = (_this).f19;
          let _rhs302 = _dafny.Seq.update((_this).f14, _module.__default.safeIndex(new BigNumber((_2322_v70).cardinality()), new BigNumber(((_this).f14).length)), (_this).f19);
          let _rhs303 = (_2321_v69)[_module.__default.safeIndex(_2320_v68, new BigNumber((_2321_v69).length))];
          let _rhs304 = new BigNumber(((_2324_v72)[_module.__default.safeIndex(_2320_v68, new BigNumber((_2324_v72).length))]).length);
          let _rhs305 = (((new BigNumber((_dafny.Seq.of((_this).f19)).length)).isLessThan(_2320_v68)) ? ((_this).f19) : ((_this).f19));
          let _lhs185 = globalState;
          let _lhs186 = globalState;
          let _lhs187 = globalState;
          _lhs185.f4 = _rhs301;
          _2321_v69 = _rhs302;
          _lhs186.f4 = _rhs303;
          _2320_v68 = _rhs304;
          _lhs187.f4 = _rhs305;
          let _2326_v74;
          _2326_v74 = _dafny.Set.fromElements(_2320_v68);
          let _2327_v75;
          _2327_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2326_v74,_2320_v68);
          let _2328_v76;
          _2328_v76 = _dafny.Seq.of(_2322_v70, _2322_v70);
          _2327_v75 = (_2327_v75).update(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_2328_v76).length)), _2320_v68), _2320_v68);
          let _2329_v77;
          let _nw366 = new _module.C3();
          _nw366.__ctor((_this).f19, _this.f12, (_this).f19, (_this).f13, _dafny.Seq.Concat(_2321_v69, _2321_v69));
          _2329_v77 = _nw366;
          let _2330_v78;
          let _nw367 = new _module.C0();
          _nw367.__ctor();
          _2330_v78 = _nw367;
        } else if (_source36.is_DC2) {
          let _2331___mcc_h8 = (_source36).cf1;
          let _2332___mcc_h9 = (_source36).cf2;
          let _2333_cf2 = _2332___mcc_h9;
          let _2334_cf1 = _2331___mcc_h8;
          let _2335_v79;
          let _init64 = ((_2336_v68) => function (_2337_i4) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_2336_v68,_dafny.Seq.of(new BigNumber(709)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2336_v68,_dafny.Seq.of(new BigNumber(144))));
          })(_2320_v68);
          let _nw368 = Array((new BigNumber(10)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw368.length); _i0_64++) {
            _nw368[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2335_v79 = _nw368;
          let _2338_v80;
          _2338_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_2334_cf1);
          let _2339_v81;
          _2339_v81 = _dafny.Map.Empty.slice().updateUnsafe((((_2338_v80).contains((_this).f19)) ? ((_2338_v80).get((_this).f19)) : (_2320_v68)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), function (_2340_i5) {
            return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-747)), function (_2341_i6) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            })).length);
          }));
          let _index354 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_2335_v79).length));
          (_2335_v79)[_index354] = _2339_v81;
          let _index355 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_2335_v79).length));
          (_2335_v79)[_index355] = (_2339_v81).Merge(_2339_v81);
          _2333_cf2 = _2333_cf2;
          let _2342_v82;
          _2342_v82 = _dafny.Seq.UnicodeFromString("veo");
          let _2343_v83;
          _2343_v83 = _dafny.Seq.of(_2334_cf1, _2320_v68);
          let _2344_v84;
          _2344_v84 = _dafny.Set.fromElements(_2333_cf2);
          let _2345_v85;
          _2345_v85 = _dafny.Set.fromElements((_2343_v83)[_module.__default.safeIndex(new BigNumber((_2344_v84).length), new BigNumber((_2343_v83).length))]);
          let _2346_v86;
          _2346_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2342_v82,(_dafny.ZERO).minus(new BigNumber((_2345_v85).length)));
          let _2347_v87;
          _2347_v87 = _module.D14.create_DC41(_2342_v82, _2334_cf1);
          _2346_v86 = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("jxkuwhg"),_2320_v68)).update((_2347_v87).dtor_cf64, ((_2333_cf2) ? (_2320_v68) : (new BigNumber((_2342_v82).length))));
          let _2348_v88;
          let _nw369 = new _module.C1();
          _nw369.__ctor((_2320_v68).isLessThanOrEqualTo(new BigNumber((_2342_v82).length)), _2333_cf2, (_this).f13, (_this).f14, _this.f12);
          _2348_v88 = _nw369;
        } else {
          let _2349___mcc_h10 = (_source36).cf0;
          let _2350_cf0 = _2349___mcc_h10;
          let _2351_v89;
          _2351_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_2320_v68);
          let _2352_v90;
          _2352_v90 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2351_v89).length),_dafny.Seq.UnicodeFromString("onysgsey")),(_this).f19);
          let _2353_v91;
          _2353_v91 = _dafny.Seq.UnicodeFromString("ghunfsbs");
          (globalState).f4 = (((_2352_v90).contains(_dafny.Map.Empty.slice().updateUnsafe(_2320_v68,_2353_v91))) ? ((_2352_v90).get(_dafny.Map.Empty.slice().updateUnsafe(_2320_v68,_2353_v91))) : ((((_this).f19) ? (_module.__default.fm2((_this).f19, globalState)) : (_module.__default.fm2((_this).f19, globalState)))));
          let _2354_v92;
          let _init65 = ((_2355_cf0) => function (_2356_i7) {
            return _2355_cf0;
          })(_2350_cf0);
          let _nw370 = Array((new BigNumber(26)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw370.length); _i0_65++) {
            _nw370[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2354_v92 = _nw370;
          let _2357_v93;
          _2357_v93 = _dafny.Set.fromElements(_2354_v92);
          let _2358_v94;
          _2358_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_dafny.Set.fromElements(_2354_v92));
          let _2359_v95;
          _2359_v95 = _dafny.Seq.of(new BigNumber(620), _2320_v68, _2320_v68, _2320_v68, _2320_v68);
          let _2360_v96;
          let _nw371 = new _module.C3();
          _nw371.__ctor(_2350_cf0, _this.f12, (_this).f19, (_this).f13, _dafny.Seq.of((_this).f19, (_this).f19));
          _2360_v96 = _nw371;
          let _2361_v97;
          _2361_v97 = _dafny.Map.Empty.slice().updateUnsafe(_2350_cf0,_2360_v96);
          let _2362_v98;
          let _nw372 = Array((new BigNumber(16)).toNumber());
          _nw372[(_dafny.ZERO).toNumber()] = (_2357_v93).IsProperSubsetOf((((_2358_v94).contains((_this).f14)) ? ((_2358_v94).get((_this).f14)) : (_dafny.Set.fromElements(_2354_v92))));
          _nw372[(_dafny.ONE).toNumber()] = (((_this).f19) ? ((_this).f19) : (_2350_cf0));
          _nw372[(new BigNumber(2)).toNumber()] = (_this).f19;
          _nw372[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_2359_v95, _2359_v95);
          _nw372[(new BigNumber(4)).toNumber()] = (_this).f19;
          _nw372[(new BigNumber(5)).toNumber()] = _2350_cf0;
          _nw372[(new BigNumber(6)).toNumber()] = _2350_cf0;
          _nw372[(new BigNumber(7)).toNumber()] = (_2320_v68).isLessThanOrEqualTo(_2320_v68);
          _nw372[(new BigNumber(8)).toNumber()] = (_2361_v97).contains((_this).f19);
          _nw372[(new BigNumber(9)).toNumber()] = _2350_cf0;
          _nw372[(new BigNumber(10)).toNumber()] = true;
          _nw372[(new BigNumber(11)).toNumber()] = (_this).f19;
          _nw372[(new BigNumber(12)).toNumber()] = _2350_cf0;
          _nw372[(new BigNumber(13)).toNumber()] = (_this).f19;
          _nw372[(new BigNumber(14)).toNumber()] = (_this).f19;
          _nw372[(new BigNumber(15)).toNumber()] = (_this).f19;
          _2362_v98 = _nw372;
          let _index356 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2362_v98).length));
          (_2362_v98)[_index356] = ((_this).f14)[_module.__default.safeIndex(_2320_v68, new BigNumber(((_this).f14).length))];
          let _index357 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2362_v98).length));
          (_2362_v98)[_index357] = (_this).f19;
          let _2363_v99;
          _2363_v99 = _dafny.Set.fromElements(_2318_v66, _2318_v66, _2318_v66);
          let _2364_v100;
          _2364_v100 = _module.D7.create_DC22(!((_2362_v98)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_2362_v98).length))]));
          let _2365_v101;
          _2365_v101 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(402),_2320_v68);
          _2363_v99 = (_module.__default.fm64(_2350_cf0, _2364_v100, (((_2365_v101).contains(_2320_v68)) ? ((_2365_v101).get(_2320_v68)) : (_2320_v68)), _2320_v68, globalState)).Union((_2363_v99).Union(_2363_v99));
          let _2366_v102;
          let _nw373 = Array((new BigNumber(2)).toNumber()).fill([]);
          _2366_v102 = _nw373;
          let _2367_v103;
          _2367_v103 = _dafny.Map.Empty.slice().updateUnsafe(_2320_v68,(_2364_v100).dtor_cf34);
          let _2368_v104;
          _2368_v104 = _module.D11.create_DC31();
          let _2369_v105;
          let _nw374 = new _module.C1();
          _nw374.__ctor((_this).f19, (_2362_v98)[_module.__default.safeIndex(new BigNumber(67), new BigNumber((_2362_v98).length))], (_this).f13, (_this).f14, _2360_v96.f12);
          _2369_v105 = _nw374;
          let _2370_v106;
          _2370_v106 = _dafny.Map.Empty.slice().updateUnsafe(_2369_v105,(_2369_v105).f17);
          let _2371_v107;
          let _nw375 = new _module.C2();
          _nw375.__ctor(_dafny.Seq.of(_2320_v68, _2320_v68), (((_2370_v106).contains(_2369_v105)) ? ((_2370_v106).get(_2369_v105)) : (_2369_v105.f18)), (_this).f13, _dafny.Seq.update((_this).f14, _module.__default.safeIndex(new BigNumber((_2351_v89).length), new BigNumber(((_this).f14).length)), (_2369_v105).f17), _2360_v96.f12);
          _2371_v107 = _nw375;
          let _2372_v108;
          _2372_v108 = _dafny.Map.Empty.slice().updateUnsafe(_2368_v104,_2371_v107);
          let _2373_v109;
          _2373_v109 = _dafny.MultiSet.fromElements(_2372_v108, (_2372_v108).update(_module.D11.create_DC31(), _2371_v107), _2372_v108, _2372_v108, _2372_v108);
          let _2374_v110;
          let _nw376 = Array((new BigNumber(21)).toNumber());
          _nw376[(_dafny.ZERO).toNumber()] = _2320_v68;
          _nw376[(_dafny.ONE).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(2)).toNumber()] = new BigNumber(247);
          _nw376[(new BigNumber(3)).toNumber()] = new BigNumber((_2367_v103).length);
          _nw376[(new BigNumber(4)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(5)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(6)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(7)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(8)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(9)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(10)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_2320_v68);
          _nw376[(new BigNumber(12)).toNumber()] = new BigNumber((_2373_v109).cardinality());
          _nw376[(new BigNumber(13)).toNumber()] = new BigNumber(((_this).f14).length);
          _nw376[(new BigNumber(14)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(15)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(16)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(17)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(18)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(19)).toNumber()] = _2320_v68;
          _nw376[(new BigNumber(20)).toNumber()] = new BigNumber(238);
          _2374_v110 = _nw376;
          let _index358 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_2366_v102).length));
          (_2366_v102)[_index358] = _2374_v110;
          let _2375_v111;
          _2375_v111 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index359 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2362_v98).length));
          let _index360 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_2366_v102).length));
          let _rhs306 = _dafny.Seq.update(_2353_v91, _module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber(619), _2320_v68), new BigNumber((_2353_v91).length)), _2375_v111);
          let _rhs307 = ((_this).f14)[_module.__default.safeIndex((_2371_v107).fm38((_dafny.ZERO).minus(_2320_v68), _2320_v68, _2320_v68, (_this).f19, globalState), new BigNumber(((_this).f14).length))];
          let _rhs308 = (((_2369_v105).f17) ? ((_2369_v105.f18) || (_2369_v105.f18)) : (((_2369_v105).f17) && (false)));
          let _rhs309 = (_2320_v68).multipliedBy(new BigNumber(682));
          let _rhs310 = _2374_v110;
          let _lhs188 = _2362_v98;
          let _lhs189 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2362_v98).length));
          let _lhs190 = _2366_v102;
          let _lhs191 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_2366_v102).length));
          _2353_v91 = _rhs306;
          _lhs188[_lhs189] = _rhs307;
          r0 = _rhs308;
          _2320_v68 = _rhs309;
          _lhs190[_lhs191] = _rhs310;
        }
        let _2376_v112;
        _2376_v112 = _dafny.Seq.UnicodeFromString("ip");
        let _2377_v113;
        _2377_v113 = _dafny.MultiSet.fromElements(_2376_v112);
        (globalState).f4 = (((_dafny.MultiSet.FromArray(_dafny.Seq.of(_2376_v112, _2376_v112))).IsProperSubsetOf(_2377_v113)) ? (!((_this).f19)) : ((_this).f19));
        let _2378_v114;
        let _nw377 = new _module.C3();
        _nw377.__ctor(true, _this.f12, (_this).f19, (_this).f13, (_this).f14);
        _2378_v114 = _nw377;
        let _2379_v115;
        let _init66 = ((_2380_v68) => function (_2381_i8) {
          return _module.__default.safeDivisionInt(_2381_i8, _2380_v68);
        })(_2320_v68);
        let _nw378 = Array((new BigNumber(21)).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw378.length); _i0_66++) {
          _nw378[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2379_v115 = _nw378;
        let _2382_v116;
        _2382_v116 = _dafny.MultiSet.fromElements(_2320_v68, _2320_v68);
        let _2383_v117;
        _2383_v117 = _dafny.Seq.of(new BigNumber((_2382_v116).cardinality()));
        let _index361 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_2379_v115).length));
        (_2379_v115)[_index361] = (((_2382_v116).contains(_2320_v68)) ? ((_2382_v116).get(_2320_v68)) : (new BigNumber((_2383_v117).length)));
        let _2384_v118;
        let _nw379 = Array((new BigNumber(13)).toNumber()).fill(false);
        _2384_v118 = _nw379;
        let _2385_v119;
        _2385_v119 = _dafny.Set.fromElements((_2378_v114).f20, false, (_this).f19, (_2378_v114).f20, (_2378_v114).f20);
        let _index362 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_2384_v118).length));
        (_2384_v118)[_index362] = (_2385_v119).IsSubsetOf(_2385_v119);
        let _2386_v120;
        _2386_v120 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2385_v119).length),_2385_v119);
        let _index363 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_2379_v115).length));
        (_2379_v115)[_index363] = _module.__default.safeDivisionInt(_2320_v68, new BigNumber((_2386_v120).length));
        let _2387_v121;
        _2387_v121 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2320_v68),(_this).f19);
        let _2388_v122;
        let _init67 = ((_2389_v112) => function (_2390_i9) {
          return _2389_v112;
        })(_2376_v112);
        let _nw380 = Array((new BigNumber(2)).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw380.length); _i0_67++) {
          _nw380[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2388_v122 = _nw380;
        let _2391_v123;
        _2391_v123 = _dafny.Map.Empty.slice().updateUnsafe(_2376_v112,_module.__default.fm1((_2378_v114).f20, (_2378_v114).f20, globalState));
        let _2392_v124;
        _2392_v124 = _module.D2.create_DC6(_2384_v118, _2384_v118);
        let _2393_v125;
        _2393_v125 = _dafny.Set.fromElements((_2392_v124).dtor_cf9, _2384_v118);
        let _2394_v126;
        _2394_v126 = _module.D12.create_DC36(_2388_v122, new BigNumber((_2383_v117).length), _2391_v123, _2393_v125);
        let _index364 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_2379_v115).length));
        let _index365 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_2384_v118).length));
        let _index366 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_2379_v115).length));
        let _rhs311 = _2320_v68;
        let _rhs312 = ((true) ? (_2320_v68) : (new BigNumber((_2387_v121).length)));
        let _rhs313 = _module.__default.fm2(!(true), globalState);
        let _rhs314 = (_2394_v126).dtor_cf56;
        let _lhs192 = _2379_v115;
        let _lhs193 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_2379_v115).length));
        let _lhs194 = _2384_v118;
        let _lhs195 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_2384_v118).length));
        let _lhs196 = _2379_v115;
        let _lhs197 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_2379_v115).length));
        _2320_v68 = _rhs311;
        _lhs192[_lhs193] = _rhs312;
        _lhs194[_lhs195] = _rhs313;
        _lhs196[_lhs197] = _rhs314;
        _2320_v68 = (((_this).f19) ? (_2320_v68) : ((_2379_v115)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_2379_v115).length))]));
      }
      r0 = !(!((_this).f19)) || ((_this).f19);
      let _2395_v127;
      _2395_v127 = new BigNumber(372);
      _2395_v127 = _2395_v127;
      let _2396_v128;
      _2396_v128 = _dafny.Seq.UnicodeFromString("wwb");
      let _2397_v129;
      _2397_v129 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm65(new BigNumber(-587), (_this).f19, globalState),_2396_v128);
      let _2398_v130;
      _2398_v130 = new _dafny.CodePoint('v'.codePointAt(0));
      _2396_v128 = (((_2397_v129).contains((_dafny.Map.Empty.slice().updateUnsafe(_2398_v130,_2398_v130)).update(_2398_v130, _2398_v130))) ? ((_2397_v129).get((_dafny.Map.Empty.slice().updateUnsafe(_2398_v130,_2398_v130)).update(_2398_v130, _2398_v130))) : (_dafny.Seq.Concat(_2396_v128, _2396_v128)));
      _2395_v127 = _2395_v127;
      let _2399_v131;
      _2399_v131 = _dafny.Seq.of(_2395_v127);
      let _2400_v132;
      let _nw381 = new _module.C1();
      _nw381.__ctor((_this).f19, (_this).f19, (_this).f13, _dafny.Seq.of((_this).f19), _this.f12);
      _2400_v132 = _nw381;
      let _2401_v133;
      _2401_v133 = _module.D3.create_DC10(false, new BigNumber((_2399_v131).length), _2400_v132);
      let _2402_v134;
      _2402_v134 = _dafny.MultiSet.fromElements(_2401_v133);
      r0 = !(!(_2402_v134).contains(_2401_v133)) || ((_this).f19);
      return r0;
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f12 = [];
      this._f16 = _dafny.Seq.of();
      this._f15 = false;
      this._f13 = _dafny.Map.Empty;
      this._f14 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T3, _module.T4, _module.T2, _module.T1, _module.T0];
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
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    __ctor(f16, f15, f13, f14, f12) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f15 = f15;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f12 = f12;
      return;
    }
    fm9(globalState) {
      let _this = this;
      if ((_this).f15) {
        return _dafny.MultiSet.fromElements((_this).f15, (_this).f15);
      } else {
        return _dafny.MultiSet.fromElements((_this).f15, (_this).f15, (_this).f15, (_this).f15);
      }
    };
    fm7(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("h"), _dafny.Seq.UnicodeFromString("qi"))).length)), _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_this).f15),_dafny.MultiSet.fromElements((_this).f15, (_this).f15))).length), new BigNumber(-561)));
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return ((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(354)), function (_2403_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length)).plus(new BigNumber(((_this).f14).length))).plus(((_module.D0.create_DC2(new BigNumber(-182), false)).dtor_cf1).multipliedBy(new BigNumber(-297)));
    };
    fm6(p0, globalState) {
      let _this = this;
      let _source37 = _module.D0.create_DC2((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_2404_i0) {
  return new _dafny.CodePoint('j'.codePointAt(0));
})).length))), (_this).f15);
      if (_source37.is_DC1) {
        return (_this).f16;
      } else if (_source37.is_DC2) {
        let _2405___mcc_h0 = (_source37).cf1;
        let _2406___mcc_h1 = (_source37).cf2;
        let _2407_cf2 = _2406___mcc_h1;
        let _2408_cf1 = _2405___mcc_h0;
        return (_this).f16;
      } else {
        let _2409___mcc_h2 = (_source37).cf0;
        let _2410_cf0 = _2409___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray((_this).f14)).cardinality()))), _dafny.Seq.update((_this).f16, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2410_cf0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15)).length))).length), new BigNumber(((_this).f16).length)), new BigNumber(17)));
      }
    };
    fm5(p0, globalState) {
      let _this = this;
      return new BigNumber(932);
    };
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC2(new BigNumber((_dafny.Seq.UnicodeFromString("l")).length), (_this).f15),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), function (_2411_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length)))).length)))).Union((_dafny.Set.fromElements(new BigNumber(832))).Union(_dafny.Set.fromElements(new BigNumber(254), new BigNumber(-852))));
    };
    fm11(globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("gitfodix");
    };
    fm12(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ydouas"), _dafny.Seq.UnicodeFromString("qlaj"));
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _2412_v0;
      _2412_v0 = _dafny.Set.fromElements((_this).f15);
      let _2413_v1;
      _2413_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2412_v0,(p0).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.fromElements((_this).f15)).cardinality())));
      _2413_v1 = (_2413_v1).update(_2412_v0, (_this).f15);
      let _2414_v2;
      _2414_v2 = _module.D0.create_DC2(p1, (_this).f15);
      let _2415_v3;
      _2415_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f15);
      let _2416_v4;
      _2416_v4 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(421)).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements((_this).f15, (_2414_v2).dtor_cf2)).length)),(((_2415_v3).contains((_dafny.ZERO).minus((_module.D0.create_DC2(new BigNumber(100), (_this).f15)).dtor_cf1))) ? ((_2415_v3).get((_dafny.ZERO).minus((_module.D0.create_DC2(new BigNumber(100), (_this).f15)).dtor_cf1))) : (!((_this).f15))));
      _2416_v4 = (_2416_v4).update(((_this).f15) || ((_this).f15), false);
      if ((_module.__default.fm13(p1, globalState)).equals(_2412_v0)) {
        let _2417_v5;
        _2417_v5 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(349)), function (_2418_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }),p2);
        _2417_v5 = (((false) ? (_2417_v5) : (_2417_v5))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
        let _2419_v6;
        let _nw382 = new _module.C0();
        _nw382.__ctor();
        _2419_v6 = _nw382;
        let _2420_v7;
        let _nw383 = Array((new BigNumber(18)).toNumber()).fill(false);
        _2420_v7 = _nw383;
        let _2421_v8;
        _2421_v8 = _dafny.MultiSet.fromElements(p2);
        let _2422_v9;
        _2422_v9 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f15),_2421_v8);
        let _index367 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_2420_v7).length));
        (_2420_v7)[_index367] = !(!((((_2422_v9).contains(false)) ? ((_2422_v9).get(false)) : (_2421_v8))).contains(p2));
        let _index368 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_2420_v7).length));
        (_2420_v7)[_index368] = (_this).f15;
        r0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(476), (p1).multipliedBy(p1)));
        let _2423_v10;
        _2423_v10 = _module.D2.create_DC6(_2420_v7, _2420_v7);
        let _2424_v11;
        _2424_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v10,new BigNumber(((_2416_v4).Merge(_2416_v4)).length));
        let _pat_let_tv39 = _2420_v7;
        _2424_v11 = (_2424_v11).update(function (_pat_let45_0) {
          return function (_2425_dt__update__tmp_h0) {
            return function (_pat_let46_0) {
              return function (_2426_dt__update_hcf8_h0) {
                return _module.D2.create_DC6(_2426_dt__update_hcf8_h0, (_2425_dt__update__tmp_h0).dtor_cf9);
              }(_pat_let46_0);
            }(_pat_let_tv39);
          }(_pat_let45_0);
        }(_2423_v10), _module.__default.safeModuloInt(p0, p0));
      } else {
        (globalState).f4 = (_this).f15;
        let _2427_v12;
        let _nw384 = new _module.C0();
        _nw384.__ctor();
        _2427_v12 = _nw384;
        r0 = p1;
        let _2428_v13;
        let _nw385 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _2428_v13 = _nw385;
        let _index369 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_2428_v13).length));
        (_2428_v13)[_index369] = p0;
        let _index370 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_2428_v13).length));
        (_2428_v13)[_index370] = (_dafny.ZERO).minus(((p0).minus(new BigNumber(496))).plus(p0));
        let _index371 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_2428_v13).length));
        (_2428_v13)[_index371] = new BigNumber(660);
      }
      let _hi12 = new BigNumber((_2412_v0).length);
      for (let _2429_i1 = p1; _2429_i1.isLessThan(_hi12); _2429_i1 = _2429_i1.plus(_dafny.ONE)) {
        _2416_v4 = (_2416_v4).update((_this).f15, !((_this).f15));
        let _2430_v14;
        let _nw386 = new _module.C0();
        _nw386.__ctor();
        _2430_v14 = _nw386;
        let _2431_v15;
        _2431_v15 = _dafny.Map.Empty.slice().updateUnsafe((_2429_i1).isLessThan(_2429_i1),new BigNumber(659));
        r0 = new BigNumber((_2431_v15).length);
        let _2432_v16;
        let _nw387 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _2432_v16 = _nw387;
        let _index372 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_2432_v16).length));
        (_2432_v16)[_index372] = _module.__default.safeModuloInt(p0, _2429_i1);
        let _index373 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_2432_v16).length));
        (_2432_v16)[_index373] = (_2429_i1).multipliedBy(p0);
      }
      if ((_this).f15) {
        (globalState).f4 = (_this).f15;
        r0 = _module.__default.safeDivisionInt((p1).plus(p1), new BigNumber(738));
        let _2433_v17;
        _2433_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        _2433_v17 = (_2433_v17).update(p1, _dafny.Seq.Concat(p2, p2));
        let _2434_v18;
        _2434_v18 = new _dafny.CodePoint('q'.codePointAt(0));
        let _2435_v19;
        _2435_v19 = _dafny.Seq.of(_dafny.Seq.Concat(p2, _dafny.Seq.update(p2, _module.__default.safeIndex(p1, new BigNumber((p2).length)), _2434_v18)), p2, _dafny.Seq.Concat(p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(371)), ((_2436_v18) => function (_2437_i2) {
          return _2436_v18;
        })(_2434_v18))), (_this).fm11(globalState));
        let _2438_v20;
        _2438_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2415_v3).length),_2435_v19);
        _2435_v19 = (((_2438_v20).contains((_dafny.ZERO).minus(((!((_this).f15)) ? (p0) : (p1))))) ? ((_2438_v20).get((_dafny.ZERO).minus(((!((_this).f15)) ? (p0) : (p1))))) : (_2435_v19));
        r0 = p0;
      } else {
        let _source38 = _module.__default.fm16(_module.__default.fm2(false, globalState), globalState);
        if (_source38.is_DC5) {
          let _2439___mcc_h0 = (_source38).cf5;
          let _2440___mcc_h1 = (_source38).cf6;
          let _2441___mcc_h2 = (_source38).cf7;
          let _2442_cf7 = _2441___mcc_h2;
          let _2443_cf6 = _2440___mcc_h1;
          let _2444_cf5 = _2439___mcc_h0;
          let _2445_v21;
          let _nw388 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _2445_v21 = _nw388;
          let _index374 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_2445_v21).length));
          (_2445_v21)[_index374] = _2416_v4;
          let _2446_v22;
          _2446_v22 = _dafny.Set.fromElements((_this).f14, (_this).f14);
          let _2447_v23;
          _2447_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2446_v22).length),p1);
          let _index375 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_2445_v21).length));
          let _rhs315 = new BigNumber(343);
          let _rhs316 = (((_2447_v23).contains(_2444_cf5)) ? ((_2447_v23).get(_2444_cf5)) : (p0));
          let _rhs317 = _2416_v4;
          let _rhs318 = p1;
          let _rhs319 = (_this).f15;
          let _lhs198 = _2445_v21;
          let _lhs199 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_2445_v21).length));
          let _lhs200 = globalState;
          r0 = _rhs315;
          _2444_cf5 = _rhs316;
          _lhs198[_lhs199] = _rhs317;
          _2444_cf5 = _rhs318;
          _lhs200.f4 = _rhs319;
          r0 = (_dafny.ZERO).minus(p1);
          let _2448_v24;
          let _nw389 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2448_v24 = _nw389;
          _2448_v24 = _2448_v24;
          _2444_cf5 = _2444_cf5;
        } else if (_source38.is_DC6) {
          let _2449___mcc_h3 = (_source38).cf8;
          let _2450___mcc_h4 = (_source38).cf9;
          let _2451_cf9 = _2450___mcc_h4;
          let _2452_cf8 = _2449___mcc_h3;
          let _2453_v25;
          _2453_v25 = _dafny.Seq.of((_this).f15);
          _2453_v25 = _2453_v25;
          (globalState).f4 = !((_module.__default.safeModuloInt((_dafny.ZERO).minus(p1), p1)).isLessThan(p1));
          let _2454_v26;
          _2454_v26 = _dafny.Seq.of(_2453_v25, (_this).f14, _2453_v25);
          let _2455_v27;
          _2455_v27 = _module.D2.create_DC4(_2454_v26);
          let _2456_v28;
          _2456_v28 = _module.D2.create_DC7(_2455_v27);
          let _2457_v29;
          _2457_v29 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f15, (_this).f15, (_this).f15, (_this).f15, (_this).f15),_2456_v28);
          _2457_v29 = (_2457_v29).update((_this).f14, _2456_v28);
          let _2458_v30;
          _2458_v30 = _dafny.Set.fromElements(p0);
          let _2459_v31;
          _2459_v31 = _dafny.Set.fromElements(new BigNumber((_2458_v30).length), p0);
          let _2460_v32;
          _2460_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(89), new BigNumber((_2459_v31).length)),_2452_cf8);
          _2451_cf9 = (((_2460_v32).contains((_dafny.ZERO).minus((p1).multipliedBy(new BigNumber(290))))) ? ((_2460_v32).get((_dafny.ZERO).minus((p1).multipliedBy(new BigNumber(290))))) : (_2451_cf9));
        } else if (_source38.is_DC4) {
          let _2461___mcc_h5 = (_source38).cf4;
          let _2462_cf4 = _2461___mcc_h5;
          let _2463_v33;
          let _init68 = ((_2464_v4) => function (_2465_i3) {
            return (_2465_i3).multipliedBy(new BigNumber((_2464_v4).length));
          })(_2416_v4);
          let _nw390 = Array((new BigNumber(7)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw390.length); _i0_68++) {
            _nw390[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _2463_v33 = _nw390;
          let _2466_v34;
          _2466_v34 = _module.D0.create_DC1();
          let _2467_v35;
          _2467_v35 = _dafny.Seq.of(p2);
          let _2468_v36;
          _2468_v36 = _module.D2.create_DC5(p0, (_this).f15, !((_this).f15));
          let _2469_v37;
          _2469_v37 = _module.D0.create_DC0((_this).f15);
          let _2470_v38;
          _2470_v38 = new _dafny.CodePoint('i'.codePointAt(0));
          let _rhs320 = _2463_v33;
          let _rhs321 = (((_2468_v36).dtor_cf7) ? (_2466_v34) : (_2466_v34));
          let _rhs322 = (_this).fm12(_2469_v37, _2470_v38, globalState);
          let _rhs323 = p0;
          let _rhs324 = new BigNumber((_dafny.Seq.of((_this).f15, (_this).f15)).length);
          _2463_v33 = _rhs320;
          _2466_v34 = _rhs321;
          _2467_v35 = _rhs322;
          r0 = _rhs323;
          r0 = _rhs324;
          _2416_v4 = (_2416_v4).update((_this).f15, ((_this).f14)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber(((_this).f14).length))]);
          (globalState).f4 = _module.__default.fm2((p0).isLessThan(p1), globalState);
          let _index376 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_2463_v33).length));
          (_2463_v33)[_index376] = new BigNumber((_2412_v0).length);
          let _index377 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_2463_v33).length));
          (_2463_v33)[_index377] = p0;
        } else {
          let _2471___mcc_h6 = (_source38).cf10;
          let _2472_cf10 = _2471___mcc_h6;
          let _2473_v39;
          _2473_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _2474_v40;
          _2474_v40 = _dafny.MultiSet.fromElements(_2473_v39);
          (globalState).f4 = !(_2474_v40).equals((_2474_v40).Difference(_2474_v40));
          let _2475_v41;
          _2475_v41 = _module.D2.create_DC5(p0, false, (_this).f15);
          let _2476_v42;
          let _nw391 = new _module.C1();
          _nw391.__ctor(true, (_2475_v41).dtor_cf6, (_this).f13, (_this).f14, _this.f12);
          _2476_v42 = _nw391;
          let _2477_v43;
          _2477_v43 = _dafny.MultiSet.fromElements(_2476_v42);
          (globalState).f4 = ((_2477_v43).Union(_2477_v43)).equals(_2477_v43);
          (globalState).f4 = !(_2412_v0).contains(!((_this).f15));
          let _2478_v44;
          let _init69 = function (_2479_i4) {
            return _module.D0.create_DC0((_this).f15);
          };
          let _nw392 = Array((new BigNumber(27)).toNumber());
          for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw392.length); _i0_69++) {
            _nw392[_i0_69] = _init69(new BigNumber(_i0_69));
          }
          _2478_v44 = _nw392;
          let _index378 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2478_v44).length));
          (_2478_v44)[_index378] = _module.__default.fm19(_dafny.Set.fromElements((_this).f15, (_this).f15), (_this).f15, new _dafny.CodePoint('l'.codePointAt(0)), p1, globalState);
          let _2480_v45;
          _2480_v45 = _module.D0.create_DC0((_this).f15);
          let _index379 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2478_v44).length));
          (_2478_v44)[_index379] = _2480_v45;
        }
        let _2481_v46;
        _2481_v46 = _dafny.Seq.UnicodeFromString("xl");
        _2481_v46 = _2481_v46;
        let _2482_v47;
        let _nw393 = new _module.C0();
        _nw393.__ctor();
        _2482_v47 = _nw393;
        let _2483_v48;
        let _nw394 = Array((new BigNumber(23)).toNumber()).fill(false);
        _2483_v48 = _nw394;
        let _2484_v49;
        let _nw395 = new _module.C4();
        _nw395.__ctor();
        _2484_v49 = _nw395;
        let _2485_v50;
        _2485_v50 = _dafny.Map.Empty.slice().updateUnsafe(_2484_v49,_dafny.Seq.UnicodeFromString("hku"));
        let _index380 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_2483_v48).length));
        (_2483_v48)[_index380] = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_2486_i5) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }), (((_2485_v50).contains(_2484_v49)) ? ((_2485_v50).get(_2484_v49)) : ((_this).fm11(globalState)))));
        let _2487_v51;
        let _nw396 = new _module.C1();
        _nw396.__ctor(!(!(true) || ((_this).f15)), (_this).f15, (_this).f13, _dafny.Seq.of((_this).f15, false), _this.f12);
        _2487_v51 = _nw396;
        let _2488_v52;
        _2488_v52 = _dafny.MultiSet.fromElements((_this).f15);
        let _2489_v53;
        _2489_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2488_v52,(_this).f15);
        let _2490_v54;
        _2490_v54 = _2482_v47;
        let _index381 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_2483_v48).length));
        let _rhs325 = (_2487_v51).f17;
        let _rhs326 = _module.__default.safeDivisionInt(p0, _module.__default.safeModuloInt(p0, new BigNumber((_2489_v53).length)));
        let _rhs327 = (_2490_v54);
        let _rhs328 = false;
        let _rhs329 = _2487_v51;
        let _lhs201 = globalState;
        let _lhs202 = _2483_v48;
        let _lhs203 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_2483_v48).length));
        _lhs201.f4 = _rhs325;
        r0 = _rhs326;
        _2482_v47 = _rhs327;
        _lhs202[_lhs203] = _rhs328;
        _2487_v51 = _rhs329;
        let _2491_v55;
        let _nw397 = new _module.C3();
        _nw397.__ctor((_2487_v51).f17, _this.f12, (_this).f15, (_this).f13, (((_2487_v51).f17) ? ((_this).f14) : ((_this).f14)));
        _2491_v55 = _nw397;
        _2481_v46 = p2;
      }
      let _2492_v56;
      let _nw398 = Array((_dafny.ONE).toNumber());
      _nw398[(_dafny.ZERO).toNumber()] = p2;
      _2492_v56 = _nw398;
      let _index382 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_2492_v56).length));
      (_2492_v56)[_index382] = p2;
      let _index383 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_2492_v56).length));
      (_2492_v56)[_index383] = p2;
      r0 = _module.__default.safeDivisionInt(p0, new BigNumber(((_2492_v56)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_2492_v56).length))]).length));
      return r0;
    }
    m3(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let r3 = false;
      let _2493_v0;
      _2493_v0 = _dafny.Seq.UnicodeFromString("bxffsdmo");
      let _2494_v1;
      _2494_v1 = new BigNumber(30);
      let _hi13 = (_this).fm5((_this).f15, globalState);
      for (let _2495_i0 = _module.__default.safeModuloInt(new BigNumber((_2493_v0).length), (_this).fm7(_2494_v1, globalState)); _2495_i0.isLessThan(_hi13); _2495_i0 = _2495_i0.plus(_dafny.ONE)) {
        let _2496_v2;
        _2496_v2 = _dafny.MultiSet.fromElements(_2494_v1, _2494_v1);
        let _2497_v3;
        _2497_v3 = _dafny.MultiSet.fromElements((((_2496_v2).contains(_2494_v1)) ? ((_2496_v2).get(_2494_v1)) : (_2494_v1)));
        let _2498_v4;
        _2498_v4 = _dafny.Set.fromElements(new BigNumber(228), new BigNumber((_2496_v2).cardinality()), _2495_i0, _2495_i0, _2495_i0);
        _2498_v4 = ((_2498_v4).Union(function () {
          let _coll67 = new _dafny.Set();
          for (const _compr_67 of _dafny.IntegerRange(new BigNumber(211), new BigNumber(-38))) {
            let _2499_v5 = _compr_67;
            if (((new BigNumber(211)).isLessThanOrEqualTo(_2499_v5)) && ((_2499_v5).isLessThan(new BigNumber(-38)))) {
              _coll67.add((_2499_v5).plus(_2494_v1));
            }
          }
          return _coll67;
        }())).Intersect(_2498_v4);
        let _2500_v6;
        let _nw399 = Array((new BigNumber(28)).toNumber()).fill(false);
        _2500_v6 = _nw399;
        let _2501_v7;
        _2501_v7 = _module.D2.create_DC6(_2500_v6, _2500_v6);
        let _pat_let_tv40 = _2500_v6;
        let _source39 = function (_pat_let47_0) {
          return function (_2502_dt__update__tmp_h0) {
            return function (_pat_let48_0) {
              return function (_2503_dt__update_hcf8_h0) {
                return _module.D2.create_DC6(_2503_dt__update_hcf8_h0, (_2502_dt__update__tmp_h0).dtor_cf9);
              }(_pat_let48_0);
            }(_pat_let_tv40);
          }(_pat_let47_0);
        }(_2501_v7);
        if (_source39.is_DC5) {
          let _2504___mcc_h0 = (_source39).cf5;
          let _2505___mcc_h1 = (_source39).cf6;
          let _2506___mcc_h2 = (_source39).cf7;
          let _2507_cf7 = _2506___mcc_h2;
          let _2508_cf6 = _2505___mcc_h1;
          let _2509_cf5 = _2504___mcc_h0;
          _2509_cf5 = _module.__default.safeDivisionInt(new BigNumber(168), _2509_cf5);
          _2508_cf6 = _2507_cf7;
          r1 = ((_dafny.ZERO).minus(_2494_v1)).multipliedBy((_this).fm5((_this).f15, globalState));
          _2493_v0 = _2493_v0;
        } else if (_source39.is_DC6) {
          let _2510___mcc_h3 = (_source39).cf8;
          let _2511___mcc_h4 = (_source39).cf9;
          let _2512_cf9 = _2511___mcc_h4;
          let _2513_cf8 = _2510___mcc_h3;
          r1 = _2495_i0;
          let _2514_v8;
          _2514_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber(297));
          let _2515_v9;
          _2515_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2494_v1,_2514_v8);
          let _2516_v10;
          _2516_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2495_i0,(_this).f15);
          let _2517_v11;
          _2517_v11 = _dafny.Set.fromElements((_this).f15, false);
          r1 = _module.__default.fm0((_this).f15, _2515_v9, (((_2516_v10).contains((_dafny.ZERO).minus(new BigNumber((_2517_v11).length)))) ? ((_2516_v10).get((_dafny.ZERO).minus(new BigNumber((_2517_v11).length)))) : ((_this).f15)), globalState);
          _2493_v0 = _2493_v0;
          let _index384 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_2513_cf8).length));
          (_2513_cf8)[_index384] = (_this).f15;
          let _index385 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_2513_cf8).length));
          (_2513_cf8)[_index385] = (_this).f15;
        } else if (_source39.is_DC4) {
          let _2518___mcc_h5 = (_source39).cf4;
          let _2519_cf4 = _2518___mcc_h5;
          let _2520_v12;
          let _init70 = ((_2521_v1) => function (_2522_i1) {
            return _module.__default.safeDivisionInt(_2522_i1, (_dafny.ZERO).minus(_2521_v1));
          })(_2494_v1);
          let _nw400 = Array((new BigNumber(9)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw400.length); _i0_70++) {
            _nw400[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2520_v12 = _nw400;
          let _2523_v13;
          _2523_v13 = _module.D12.create_DC35((_this).f15);
          let _2524_v14;
          let _nw401 = Array((new BigNumber(11)).toNumber());
          _nw401[(_dafny.ZERO).toNumber()] = _2500_v6;
          _nw401[(_dafny.ONE).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(2)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(3)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(4)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(5)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(6)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(7)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(8)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(9)).toNumber()] = _2500_v6;
          _nw401[(new BigNumber(10)).toNumber()] = _2500_v6;
          _2524_v14 = _nw401;
          let _index386 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_2524_v14).length));
          (_2524_v14)[_index386] = _2500_v6;
          let _2525_v15;
          let _init71 = function (_2526_i2) {
            return (_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0))));
          };
          let _nw402 = Array((new BigNumber(10)).toNumber());
          for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw402.length); _i0_71++) {
            _nw402[_i0_71] = _init71(new BigNumber(_i0_71));
          }
          _2525_v15 = _nw402;
          let _2527_v16;
          _2527_v16 = new _dafny.CodePoint('j'.codePointAt(0));
          let _2528_v17;
          _2528_v17 = _dafny.MultiSet.fromElements(_2527_v16, _2527_v16);
          let _2529_v18;
          _2529_v18 = _dafny.Seq.of(_2528_v17, _2528_v17);
          let _index387 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_2525_v15).length));
          (_2525_v15)[_index387] = (_2529_v18)[_module.__default.safeIndex(_2494_v1, new BigNumber((_2529_v18).length))];
          let _index388 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_2524_v14).length));
          let _index389 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_2525_v15).length));
          let _rhs330 = _2520_v12;
          let _rhs331 = _2523_v13;
          let _rhs332 = _2500_v6;
          let _rhs333 = !((_this).f15);
          let _rhs334 = _dafny.MultiSet.fromElements(_2527_v16, _2527_v16, _2527_v16, _2527_v16);
          let _lhs204 = _2524_v14;
          let _lhs205 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_2524_v14).length));
          let _lhs206 = globalState;
          let _lhs207 = _2525_v15;
          let _lhs208 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_2525_v15).length));
          _2520_v12 = _rhs330;
          _2523_v13 = _rhs331;
          _lhs204[_lhs205] = _rhs332;
          _lhs206.f4 = _rhs333;
          _lhs207[_lhs208] = _rhs334;
          r0 = (!((_this).f15) || ((_this).f15)) || ((_this).f15);
          let _2530_v19;
          _2530_v19 = _module.D0.create_DC2(_2494_v1, (_this).f15);
          let _2531_v20;
          _2531_v20 = _dafny.MultiSet.fromElements((_this).f15);
          let _2532_v21;
          _2532_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2531_v20,_2500_v6);
          let _2533_v22;
          let _2534_v23;
          let _2535_v24;
          let _2536_v25;
          let _out119;
          let _out120;
          let _out121;
          let _out122;
          let _outcollector31 = _module.__default.m0((_this).f15, _2530_v19, (_2495_i0).isEqualTo(new BigNumber((_2532_v21).length)), globalState);
          _out119 = _outcollector31[0];
          _out120 = _outcollector31[1];
          _out121 = _outcollector31[2];
          _out122 = _outcollector31[3];
          _2533_v22 = _out119;
          _2534_v23 = _out120;
          _2535_v24 = _out121;
          _2536_v25 = _out122;
          let _2537_v26;
          let _nw403 = new _module.C1();
          _nw403.__ctor(!_dafny.Seq.contains(_2493_v0, _2527_v16), (((_this).f15) ? (_2536_v25) : (_2534_v23)), ((_this).f13).Merge((_this).f13), (_this).f14, _this.f12);
          _2537_v26 = _nw403;
          _2537_v26 = _2537_v26;
        } else {
          let _2538___mcc_h6 = (_source39).cf10;
          let _2539_cf10 = _2538___mcc_h6;
          let _2540_v27;
          _2540_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,true);
          let _2541_v28;
          _2541_v28 = _module.D8.create_DC26((_this).f15);
          let _2542_v29;
          _2542_v29 = _dafny.Map.Empty.slice().updateUnsafe(_2540_v27,(_2541_v28).dtor_cf41);
          _2542_v29 = (_2542_v29).update(_2540_v27, (((_this).f15) ? ((_this).f15) : ((_this).f15)));
          let _2543_v30;
          _2543_v30 = _dafny.MultiSet.fromElements(true, (_this).f15, (_this).f15, (_this).f15, !((_this).f15));
          _2543_v30 = _2543_v30;
          let _2544_v31;
          _2544_v31 = _dafny.Set.fromElements((_this).f15);
          let _2545_v32;
          _2545_v32 = _2544_v31;
          r3 = ((((_this).f15) ? ((_2545_v32)) : (_dafny.Set.fromElements((_this).f15, false)))).IsSubsetOf(_2544_v31);
          let _index390 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_2500_v6).length));
          (_2500_v6)[_index390] = (_this).f15;
          let _index391 = _module.__default.safeIndex(new BigNumber(798), new BigNumber((_2500_v6).length));
          (_2500_v6)[_index391] = (_this).f15;
        }
        let _2546_v33;
        let _nw404 = new _module.C3();
        _nw404.__ctor(_module.__default.fm2((_this).f15, globalState), _this.f12, (_this).f15, (_this).f13, _dafny.Seq.Concat((_this).f14, (_this).f14));
        _2546_v33 = _nw404;
        r0 = (_2546_v33).f20;
      }
      (globalState).f4 = !(_module.__default.safeModuloInt((_this).fm5((_this).f15, globalState), new BigNumber(90))).isEqualTo(_module.__default.safeDivisionInt(_2494_v1, _2494_v1));
      let _2547_v34;
      let _init72 = ((_2548_v1) => function (_2549_i3) {
        return _dafny.Set.fromElements(_2548_v1, _2548_v1, new BigNumber(-781), _2548_v1);
      })(_2494_v1);
      let _nw405 = Array((new BigNumber(20)).toNumber());
      for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw405.length); _i0_72++) {
        _nw405[_i0_72] = _init72(new BigNumber(_i0_72));
      }
      _2547_v34 = _nw405;
      let _2550_v35;
      _2550_v35 = _dafny.Set.fromElements(_2494_v1);
      let _index392 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_2547_v34).length));
      (_2547_v34)[_index392] = _2550_v35;
      let _index393 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_2547_v34).length));
      (_2547_v34)[_index393] = _2550_v35;
      let _2551_v36;
      _2551_v36 = _dafny.Set.fromElements(_2493_v0);
      let _2552_v37;
      _2552_v37 = _dafny.MultiSet.fromElements(new BigNumber((_2551_v36).length));
      let _2553_v38;
      _2553_v38 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2554_v39;
      _2554_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2553_v38,(_this).f15);
      let _2555_i4;
      _2555_i4 = _dafny.ZERO;
      L14: {
        while (((_dafny.ZERO).minus(_2494_v1)).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus((((_2552_v37).contains(_2494_v1)) ? ((_2552_v37).get(_2494_v1)) : (new BigNumber((_2554_v39).length)))), new BigNumber((_2493_v0).length)))) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2555_i4)) {
              break L14;
            }
            _2555_i4 = (_2555_i4).plus(_dafny.ONE);
            r0 = !_dafny.Seq.contains((_this).f16, _2494_v1);
            let _2556_v40;
            _2556_v40 = _module.D2.create_DC5(_2494_v1, (_this).f15, (_this).f15);
            let _rhs335 = _2556_v40;
            let _rhs336 = _2553_v38;
            let _rhs337 = _module.__default.fm2((_this).f15, globalState);
            let _lhs209 = globalState;
            _2556_v40 = _rhs335;
            _2553_v38 = _rhs336;
            _lhs209.f4 = _rhs337;
            let _2557_v41;
            _2557_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(814),(_this).f15);
            let _2558_v42;
            _2558_v42 = _dafny.Map.Empty.slice().updateUnsafe((((((_2557_v41).contains(new BigNumber((_dafny.Seq.of(_2553_v38)).length))) ? ((_2557_v41).get(new BigNumber((_dafny.Seq.of(_2553_v38)).length))) : ((_this).f15))) ? ((_this).f15) : ((_this).f15)),(_this).f15);
            let _2559_v43;
            _2559_v43 = _dafny.Set.fromElements((_this).f15);
            let _2560_v44;
            _2560_v44 = _2559_v43;
            let _rhs338 = (_this).f15;
            let _rhs339 = (_this).f15;
            let _rhs340 = (_2494_v1).minus(new BigNumber(((_2560_v44)).length));
            let _rhs341 = ((_2558_v42).Merge(_2558_v42)).update(!(_2494_v1).isEqualTo(new BigNumber(548)), (_this).f15);
            let _lhs210 = globalState;
            let _lhs211 = globalState;
            _lhs210.f4 = _rhs338;
            _lhs211.f4 = _rhs339;
            _2494_v1 = _rhs340;
            _2558_v42 = _rhs341;
            let _2561_v45;
            _2561_v45 = _module.D17.create_DC46((_this).f13);
            let _2562_v46;
            let _nw406 = new _module.C5();
            _nw406.__ctor((_this).f15, _this.f12, ((_this).f13).Merge((_2561_v45).dtor_cf69), (_this).f14);
            _2562_v46 = _nw406;
            let _2563_v47;
            _2563_v47 = _dafny.MultiSet.fromElements(!(_dafny.Seq.IsProperPrefixOf((_this).f16, _dafny.Seq.of((_module.D11.create_DC32(_2494_v1, new BigNumber((_2493_v0).length))).dtor_cf46, _2494_v1))), false);
            let _2564_v48;
            _2564_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2494_v1,_2563_v47);
            let _rhs342 = _2494_v1;
            let _rhs343 = _2562_v46;
            let _rhs344 = ((((_this).f15) ? (_2563_v47) : ((((_2564_v48).contains((_this).fm8((_this).f15, (_this).f15, globalState))) ? ((_2564_v48).get((_this).fm8((_this).f15, (_this).f15, globalState))) : (_2563_v47))))).update((_this).f15, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_2493_v0).length))));
            let _rhs345 = (_this).f15;
            r1 = _rhs342;
            _2562_v46 = _rhs343;
            _2563_v47 = _rhs344;
            r0 = _rhs345;
          }
        }
      }
      let _2565_i5;
      _2565_i5 = _dafny.ZERO;
      L15: {
        while (true) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2565_i5)) {
              break L15;
            }
            _2565_i5 = (_2565_i5).plus(_dafny.ONE);
            let _rhs346 = !(new BigNumber(749)).isEqualTo((_2494_v1).multipliedBy(_2494_v1));
            let _rhs347 = !((_this).f15);
            let _lhs212 = globalState;
            let _lhs213 = globalState;
            _lhs212.f4 = _rhs346;
            _lhs213.f4 = _rhs347;
            let _2566_v49;
            let _init73 = ((_2567_v1) => function (_2568_i6) {
              return _module.__default.safeModuloInt(_2568_i6, _2567_v1);
            })(_2494_v1);
            let _nw407 = Array((new BigNumber(17)).toNumber());
            for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw407.length); _i0_73++) {
              _nw407[_i0_73] = _init73(new BigNumber(_i0_73));
            }
            _2566_v49 = _nw407;
            let _index394 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_2566_v49).length));
            (_2566_v49)[_index394] = _2494_v1;
            let _index395 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_2566_v49).length));
            (_2566_v49)[_index395] = _2494_v1;
            let _2569_v50;
            let _nw408 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2569_v50 = _nw408;
            let _index396 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_2569_v50).length));
            (_2569_v50)[_index396] = _dafny.Seq.UnicodeFromString("kga");
            let _index397 = _module.__default.safeIndex(new BigNumber(253), new BigNumber((_2569_v50).length));
            (_2569_v50)[_index397] = _2493_v0;
            _2494_v1 = (_dafny.ZERO).minus((_2566_v49)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_2566_v49).length))]);
          }
        }
      }
      _2494_v1 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat((_this).f16, (_this).f16)).length), ((_module.D11.create_DC32(_2494_v1, _2494_v1)).dtor_cf46).minus(new BigNumber((_2493_v0).length)));
      r0 = (_this).f15;
      let _2570_v51;
      _2570_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2494_v1);
      let _2571_v52;
      _2571_v52 = _module.D8.create_DC25((_this).f15, _2494_v1, _2494_v1, (_2570_v51).update((_this).f15, _2494_v1));
      let _pat_let_tv41 = _2493_v0;
      let _pat_let_tv42 = _2570_v51;
      let _pat_let_tv43 = _2494_v1;
      r1 = (function (_pat_let49_0) {
        return function (_2572_dt__update__tmp_h1) {
          return function (_pat_let50_0) {
            return function (_2573_dt__update_hcf39_h0) {
              return function (_pat_let51_0) {
                return function (_2574_dt__update_hcf40_h0) {
                  return function (_pat_let52_0) {
                    return function (_2575_dt__update_hcf38_h0) {
                      return _module.D8.create_DC25((_2572_dt__update__tmp_h1).dtor_cf37, _2575_dt__update_hcf38_h0, _2573_dt__update_hcf39_h0, _2574_dt__update_hcf40_h0);
                    }(_pat_let52_0);
                  }(_pat_let_tv43);
                }(_pat_let51_0);
              }(_pat_let_tv42);
            }(_pat_let50_0);
          }(new BigNumber((_pat_let_tv41).length));
        }(_pat_let49_0);
      }(_2571_v52)).dtor_cf39;
      r2 = _dafny.Set.fromElements((_this).f14, (_this).f14, _module.__default.fm32((_this).f15, _2553_v38, new BigNumber((_module.__default.fm20((_this).f15, _2494_v1, globalState)).cardinality()), (_dafny.ZERO).minus(new BigNumber(-843)), globalState), (_this).f14, (_this).f14);
      r3 = (((_2494_v1).isLessThan(new BigNumber(-187))) ? (_dafny.areEqual(_2493_v0, _2493_v0)) : ((_this).f15));
      return [r0, r1, r2, r3];
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let _2576_v0;
      _2576_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
      _2576_v0 = (_2576_v0).update((_this).f15, (_this).f15);
      let _2577_v1;
      _2577_v1 = _dafny.Seq.of(_dafny.Seq.update((_this).f16, _module.__default.safeIndex(p2, new BigNumber(((_this).f16).length)), p2));
      let _2578_v2;
      _2578_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2577_v1).length),(_this).f15);
      _2578_v2 = (_2578_v2).update(p2, (_this).f15);
      let _2579_v3;
      _2579_v3 = _dafny.Seq.of(true);
      let _2580_v4;
      _2580_v4 = new _dafny.CodePoint('n'.codePointAt(0));
      let _2581_v5;
      _2581_v5 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f15) === ((_this).f15),_2580_v4);
      let _2582_v6;
      _2582_v6 = new BigNumber(875);
      let _2583_v7;
      _2583_v7 = _dafny.MultiSet.fromElements((_this).f15, (_this).f15);
      let _rhs348 = _dafny.Seq.Concat((_this).f14, (_this).f14);
      let _rhs349 = _2581_v5;
      let _rhs350 = p2;
      let _rhs351 = ((_2583_v7).Intersect(_2583_v7)).IsProperSubsetOf(_dafny.MultiSet.FromArray((_this).f14));
      _2579_v3 = _rhs348;
      _2581_v5 = _rhs349;
      _2582_v6 = _rhs350;
      r0 = _rhs351;
      let _2584_v8;
      let _nw409 = new _module.C5();
      _nw409.__ctor((_this).f15, _this.f12, (_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2576_v0)).Merge((_this).f13), _2579_v3);
      _2584_v8 = _nw409;
      let _2585_v9;
      let _nw410 = new _module.C5();
      _nw410.__ctor((_this).f15, _this.f12, _dafny.Map.Empty.slice().updateUnsafe(!(false),_2576_v0), (_this).f14);
      _2585_v9 = _nw410;
      let _2586_v10;
      _2586_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2585_v9,p2);
      r0 = (_2586_v10).contains(_2585_v9);
      let _2587_v11;
      _2587_v11 = _2580_v4;
      let _2588_v12;
      _2588_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(631),_2582_v6);
      let _2589_v13;
      let _nw411 = Array((new BigNumber(28)).toNumber());
      _nw411[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
      _nw411[(_dafny.ONE).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
      _nw411[(new BigNumber(3)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(4)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(5)).toNumber()] = (_2587_v11);
      _nw411[(new BigNumber(6)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(7)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
      _nw411[(new BigNumber(9)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(10)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(11)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(12)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(13)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(14)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(15)).toNumber()] = ((!((_this).f15)) ? (new _dafny.CodePoint('i'.codePointAt(0))) : (_2580_v4));
      _nw411[(new BigNumber(16)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(17)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(18)).toNumber()] = _module.__default.fm39((((_2588_v12).contains(_2582_v6)) ? ((_2588_v12).get(_2582_v6)) : (_2582_v6)), (_this).f15, p2, p1, globalState);
      _nw411[(new BigNumber(19)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(20)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(21)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(22)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(23)).toNumber()] = (_2587_v11);
      _nw411[(new BigNumber(24)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(25)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(26)).toNumber()] = _2580_v4;
      _nw411[(new BigNumber(27)).toNumber()] = _2580_v4;
      _2589_v13 = _nw411;
      let _index398 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2589_v13).length));
      (_2589_v13)[_index398] = new _dafny.CodePoint('q'.codePointAt(0));
      let _index399 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_2589_v13).length));
      (_2589_v13)[_index399] = _2580_v4;
      r0 = !((_dafny.MultiSet.fromElements(new BigNumber((p1).length))).contains(p2)) || ((_this).f15);
      return r0;
    }
    m6(globalState) {
      let _this = this;
      let r0 = false;
      let _2590_v0;
      _2590_v0 = new BigNumber(614);
      let _2591_v1;
      let _nw412 = Array((new BigNumber(23)).toNumber());
      _nw412[(_dafny.ZERO).toNumber()] = (_this).f15;
      _nw412[(_dafny.ONE).toNumber()] = false;
      _nw412[(new BigNumber(2)).toNumber()] = false;
      _nw412[(new BigNumber(3)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(4)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(5)).toNumber()] = false;
      _nw412[(new BigNumber(6)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(7)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(8)).toNumber()] = !((_this).f15);
      _nw412[(new BigNumber(9)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(10)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(11)).toNumber()] = false;
      _nw412[(new BigNumber(12)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(13)).toNumber()] = false;
      _nw412[(new BigNumber(14)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(15)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(16)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(17)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(18)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(19)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(20)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(21)).toNumber()] = (_this).f15;
      _nw412[(new BigNumber(22)).toNumber()] = (_this).f15;
      _2591_v1 = _nw412;
      let _2592_v2;
      _2592_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(431),_2591_v1);
      let _2593_v3;
      _2593_v3 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f14)[_module.__default.safeIndex(new BigNumber(((_2592_v2).update(_2590_v0, _2591_v1)).length), new BigNumber(((_this).f14).length))],(_this).f15);
      let _hi14 = new BigNumber((_2593_v3).length);
      for (let _2594_i0 = _2590_v0; _2594_i0.isLessThan(_hi14); _2594_i0 = _2594_i0.plus(_dafny.ONE)) {
        let _2595_v4;
        let _nw413 = new _module.C1();
        _nw413.__ctor((_module.D5.create_DC16((_this).f15)).dtor_cf26, (_this).f15, ((_this).f13).Merge((_this).f13), _dafny.Seq.update(_dafny.Seq.of(_module.__default.fm2(!((_this).f15), globalState)), _module.__default.safeIndex((_dafny.ZERO).minus(_2594_i0), new BigNumber((_dafny.Seq.of(_module.__default.fm2(!((_this).f15), globalState))).length)), (_this).f15), _this.f12);
        _2595_v4 = _nw413;
        let _2596_v5;
        _2596_v5 = _module.D2.create_DC6(_2591_v1, (_module.D2.create_DC6(_2591_v1, _2591_v1)).dtor_cf9);
        let _2597_v6;
        let _nw414 = Array((new BigNumber(22)).toNumber());
        _nw414[(_dafny.ZERO).toNumber()] = _2591_v1;
        _nw414[(_dafny.ONE).toNumber()] = (_2596_v5).dtor_cf9;
        _nw414[(new BigNumber(2)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(3)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(4)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(5)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(6)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(7)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(8)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(9)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(10)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(11)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(12)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(13)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(14)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(15)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(16)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(17)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(18)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(19)).toNumber()] = _2591_v1;
        _nw414[(new BigNumber(20)).toNumber()] = (_2596_v5).dtor_cf8;
        _nw414[(new BigNumber(21)).toNumber()] = _2591_v1;
        _2597_v6 = _nw414;
        let _2598_v7;
        _2598_v7 = _dafny.Map.Empty.slice().updateUnsafe((_2595_v4).f17,_2597_v6);
        _2598_v7 = (_2598_v7).update((_this).f15, _2597_v6);
        if (!(!(false)) || ((_this).f15)) {
          let _2599_v8;
          let _nw415 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _2599_v8 = _nw415;
          _2599_v8 = _2599_v8;
          let _2600_v9;
          _2600_v9 = _dafny.Seq.UnicodeFromString("mfwv");
          let _2601_v10;
          _2601_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2591_v1,new BigNumber(((_this).fm6(_2600_v9, globalState)).length));
          _2601_v10 = (_2601_v10).update(_2591_v1, _2594_i0);
          let _2602_v11;
          let _2603_v12;
          let _2604_v13;
          let _2605_v14;
          let _out123;
          let _out124;
          let _out125;
          let _out126;
          let _outcollector32 = (_2595_v4).m3(globalState);
          _out123 = _outcollector32[0];
          _out124 = _outcollector32[1];
          _out125 = _outcollector32[2];
          _out126 = _outcollector32[3];
          _2602_v11 = _out123;
          _2603_v12 = _out124;
          _2604_v13 = _out125;
          _2605_v14 = _out126;
          let _2606_v15;
          _2606_v15 = _module.D0.create_DC2(_2590_v0, !(_2605_v14));
          let _2607_v16;
          _2607_v16 = new _dafny.CodePoint('s'.codePointAt(0));
          let _2608_v17;
          _2608_v17 = _dafny.MultiSet.fromElements(_2607_v16, _2607_v16, _2607_v16, _2607_v16);
          let _2609_v18;
          let _2610_v19;
          let _2611_v20;
          let _2612_v21;
          let _out127;
          let _out128;
          let _out129;
          let _out130;
          let _outcollector33 = _module.__default.m0((_this).f15, _2606_v15, (_2608_v17).IsSubsetOf(_2608_v17), globalState);
          _out127 = _outcollector33[0];
          _out128 = _outcollector33[1];
          _out129 = _outcollector33[2];
          _out130 = _outcollector33[3];
          _2609_v18 = _out127;
          _2610_v19 = _out128;
          _2611_v20 = _out129;
          _2612_v21 = _out130;
          _2590_v0 = _2590_v0;
        } else {
          let _2613_v22;
          _2613_v22 = _dafny.MultiSet.fromElements((_2595_v4).f17);
          _2613_v22 = _2613_v22;
          let _2614_v23;
          _2614_v23 = _dafny.Set.fromElements((_this).f14);
          let _2615_v24;
          _2615_v24 = _dafny.Seq.UnicodeFromString("ai");
          let _2616_v25;
          let _nw416 = new _module.C4();
          _nw416.__ctor();
          _2616_v25 = _nw416;
          let _2617_v26;
          let _nw417 = Array((new BigNumber(28)).toNumber());
          _nw417[(_dafny.ZERO).toNumber()] = _2616_v25;
          _nw417[(_dafny.ONE).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(2)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(3)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(4)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(5)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(6)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(7)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(8)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(9)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(10)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(11)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(12)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(13)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(14)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(15)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(16)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(17)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(18)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(19)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(20)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(21)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(22)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(23)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(24)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(25)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(26)).toNumber()] = _2616_v25;
          _nw417[(new BigNumber(27)).toNumber()] = _2616_v25;
          _2617_v26 = _nw417;
          let _2618_v27;
          _2618_v27 = _dafny.MultiSet.fromElements(_2617_v26);
          let _rhs352 = !(((new BigNumber((_2615_v24).length)).multipliedBy(_2594_i0)).isLessThanOrEqualTo((_2594_i0).plus(new BigNumber((_2618_v27).cardinality()))));
          let _rhs353 = (_dafny.ZERO).minus(new BigNumber(-885));
          let _rhs354 = (_this).f15;
          let _rhs355 = (_2614_v23).Union(_2614_v23);
          let _lhs214 = _2595_v4;
          let _lhs215 = globalState;
          _lhs214.f18 = _rhs352;
          _2590_v0 = _rhs353;
          _lhs215.f4 = _rhs354;
          _2614_v23 = _rhs355;
          let _2619_v28;
          _2619_v28 = _dafny.MultiSet.fromElements(new BigNumber(780));
          let _2620_v29;
          _2620_v29 = _dafny.Seq.of((_2595_v4.f18) && (true), !(false), ((((_2619_v28).contains(new BigNumber(852))) ? ((_2619_v28).get(new BigNumber(852))) : (_2590_v0))).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), function (_2621_i1) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })).length))));
          _2620_v29 = _dafny.Seq.Concat(_2620_v29, _dafny.Seq.of((_2595_v4).f17, _2595_v4.f18, false));
          let _2622_v30;
          let _init74 = ((_2623_v0) => function (_2624_i2) {
            return (_2624_i2).plus(_2623_v0);
          })(_2590_v0);
          let _nw418 = Array((new BigNumber(4)).toNumber());
          for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw418.length); _i0_74++) {
            _nw418[_i0_74] = _init74(new BigNumber(_i0_74));
          }
          _2622_v30 = _nw418;
          let _index400 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2622_v30).length));
          (_2622_v30)[_index400] = (_2590_v0).multipliedBy(_2594_i0);
          let _index401 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2622_v30).length));
          (_2622_v30)[_index401] = _2594_i0;
          let _2625_v31;
          _2625_v31 = _module.D11.create_DC31();
          let _2626_v32;
          let _nw419 = Array((new BigNumber(17)).toNumber());
          _nw419[(_dafny.ZERO).toNumber()] = _2625_v31;
          _nw419[(_dafny.ONE).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(2)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(3)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(4)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(5)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(6)).toNumber()] = _module.D11.create_DC31();
          _nw419[(new BigNumber(7)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(8)).toNumber()] = _module.D11.create_DC31();
          _nw419[(new BigNumber(9)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(10)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(11)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(12)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(13)).toNumber()] = _module.D11.create_DC31();
          _nw419[(new BigNumber(14)).toNumber()] = _2625_v31;
          _nw419[(new BigNumber(15)).toNumber()] = _module.D11.create_DC31();
          _nw419[(new BigNumber(16)).toNumber()] = _2625_v31;
          _2626_v32 = _nw419;
          let _index402 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2626_v32).length));
          (_2626_v32)[_index402] = _module.D11.create_DC31();
          let _2627_v33;
          _2627_v33 = _module.D12.create_DC37(_2615_v24, _2595_v4.f18);
          let _2628_v34;
          _2628_v34 = _module.D4.create_DC14(_2595_v4.f18);
          let _2629_v35;
          _2629_v35 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.__default.fm40(_2628_v34, (_2622_v30)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_2622_v30).length))], _2595_v4.f18, (_this).f15, globalState));
          let _2630_v36;
          let _nw420 = Array((new BigNumber(22)).toNumber());
          _nw420[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2615_v24, _2615_v24);
          _nw420[(_dafny.ONE).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(2)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("snwcxjku");
          _nw420[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("fhlw");
          _nw420[(new BigNumber(5)).toNumber()] = (_2627_v33).dtor_cf59;
          _nw420[(new BigNumber(6)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(7)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_2615_v24, _2615_v24);
          _nw420[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-67)), function (_2631_i3) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }), _2615_v24);
          _nw420[(new BigNumber(10)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(11)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(12)).toNumber()] = (((_2629_v35).contains(_2595_v4.f18)) ? ((_2629_v35).get(_2595_v4.f18)) : (_dafny.Seq.UnicodeFromString("bn")));
          _nw420[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(438)), function (_2632_i4) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          });
          _nw420[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_2615_v24, _dafny.Seq.UnicodeFromString("xykacskdp"));
          _nw420[(new BigNumber(15)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(71)), function (_2633_i5) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }), _2615_v24);
          _nw420[(new BigNumber(17)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("ftpjtp");
          _nw420[(new BigNumber(19)).toNumber()] = _2615_v24;
          _nw420[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_2615_v24, _2615_v24);
          _nw420[(new BigNumber(21)).toNumber()] = _2615_v24;
          _2630_v36 = _nw420;
          let _index403 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_2630_v36).length));
          (_2630_v36)[_index403] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("htrpkjrkt"), _2615_v24);
          let _2634_v37;
          _2634_v37 = _dafny.Set.fromElements((_this).f15);
          let _index404 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2626_v32).length));
          let _index405 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_2630_v36).length));
          let _rhs356 = ((_dafny.Set.fromElements(!((_this).f15))).Difference(_2634_v37)).IsProperSubsetOf(_dafny.Set.fromElements((_this).f15));
          let _rhs357 = _2615_v24;
          let _rhs358 = _2625_v31;
          let _rhs359 = _2615_v24;
          let _rhs360 = _2595_v4.f18;
          let _lhs216 = _2595_v4;
          let _lhs217 = _2626_v32;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_2626_v32).length));
          let _lhs219 = _2630_v36;
          let _lhs220 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_2630_v36).length));
          let _lhs221 = _2595_v4;
          _lhs216.f18 = _rhs356;
          _2615_v24 = _rhs357;
          _lhs217[_lhs218] = _rhs358;
          _lhs219[_lhs220] = _rhs359;
          _lhs221.f18 = _rhs360;
        }
        let _2635_v38;
        let _init75 = ((_2636_v0) => function (_2637_i6) {
          return (_2637_i6).plus(_2636_v0);
        })(_2590_v0);
        let _nw421 = Array((new BigNumber(20)).toNumber());
        for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw421.length); _i0_75++) {
          _nw421[_i0_75] = _init75(new BigNumber(_i0_75));
        }
        _2635_v38 = _nw421;
        let _index406 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_2635_v38).length));
        (_2635_v38)[_index406] = _2594_i0;
        let _index407 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_2635_v38).length));
        (_2635_v38)[_index407] = _2590_v0;
      }
      let _2638_v39;
      _2638_v39 = new _dafny.CodePoint('o'.codePointAt(0));
      let _2639_v40;
      _2639_v40 = _dafny.Map.Empty.slice().updateUnsafe(true,_2638_v39);
      let _index408 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
      (_2591_v1)[_index408] = !(!(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2638_v39)).equals(_2639_v40));
      let _2640_v41;
      _2640_v41 = _dafny.Seq.UnicodeFromString("jimfygn");
      let _2641_v42;
      _2641_v42 = _dafny.Seq.of(_2640_v41, _2640_v41, _2640_v41, _dafny.Seq.UnicodeFromString("twglwr"));
      let _2642_v43;
      _2642_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2641_v42);
      let _index409 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
      (_2591_v1)[_index409] = (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(931)), function (_2643_i7) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _2640_v41), (((_2642_v43).contains((_this).f15)) ? ((_2642_v43).get((_this).f15)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(86)), ((_2644_v41) => function (_2645_i8) {
        return _2644_v41;
      })(_2640_v41)))))).length)).isLessThanOrEqualTo(_2590_v0);
      let _2646_v44;
      _2646_v44 = _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), _2638_v39, _2638_v39, new _dafny.CodePoint('f'.codePointAt(0)));
      let _2647_v45;
      _2647_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-38),_2646_v44);
      _2647_v45 = _2647_v45;
      let _2648_v46;
      _2648_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2590_v0);
      let _2649_v47;
      _2649_v47 = _module.D8.create_DC23(_2648_v46);
      let _2650_v48;
      _2650_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2649_v47,false);
      _2650_v48 = (_2650_v48).update(_2649_v47, (_this).f15);
      if ((_this).f15) {
        if ((_this).f15) {
          _2638_v39 = _2638_v39;
          let _2651_v49;
          _2651_v49 = _dafny.Seq.of(!((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]), !(_2593_v3).equals(_2593_v3), (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]);
          let _2652_v50;
          let _init76 = ((_2653_v1) => function (_2654_i9) {
            return _dafny.Map.Empty.slice().updateUnsafe((_2653_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2653_v1).length))],(_this).f15);
          })(_2591_v1);
          let _nw422 = Array((new BigNumber(7)).toNumber());
          for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw422.length); _i0_76++) {
            _nw422[_i0_76] = _init76(new BigNumber(_i0_76));
          }
          _2652_v50 = _nw422;
          let _2655_v51;
          _2655_v51 = _2652_v50;
          let _2656_v52;
          _2656_v52 = _dafny.Seq.of(_2652_v50, _2652_v50);
          let _2657_v53;
          _2657_v53 = _module.D12.create_DC37(_dafny.Seq.Create(_module.__default.abs(new BigNumber(548)), ((_2658_v39) => function (_2659_i10) {
  return _2658_v39;
})(_2638_v39)), (_this).f15);
          let _2660_v54;
          let _nw423 = Array((new BigNumber(26)).toNumber());
          _nw423[(_dafny.ZERO).toNumber()] = _2652_v50;
          _nw423[(_dafny.ONE).toNumber()] = (((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]) ? (_2652_v50) : (_2652_v50));
          _nw423[(new BigNumber(2)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(3)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(4)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(5)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(6)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(7)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(8)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(9)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(10)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(11)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(12)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(13)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(14)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(15)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(16)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(17)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(18)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(19)).toNumber()] = (_2652_v50);
          _nw423[(new BigNumber(20)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(21)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(22)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(23)).toNumber()] = (_2656_v52)[_module.__default.safeIndex(new BigNumber(((_2657_v53).dtor_cf59).length), new BigNumber((_2656_v52).length))];
          _nw423[(new BigNumber(24)).toNumber()] = _2652_v50;
          _nw423[(new BigNumber(25)).toNumber()] = _2652_v50;
          _2660_v54 = _nw423;
          let _index410 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2660_v54).length));
          (_2660_v54)[_index410] = _2652_v50;
          let _2661_v55;
          _2661_v55 = _dafny.MultiSet.fromElements((_this).f15);
          let _2662_v56;
          _2662_v56 = _dafny.MultiSet.fromElements(new BigNumber((_2661_v55).cardinality()));
          let _2663_v57;
          _2663_v57 = _dafny.Map.Empty.slice().updateUnsafe(_2662_v56,_2638_v39);
          let _index411 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2660_v54).length));
          let _rhs361 = ((!(!(_2590_v0).isEqualTo(new BigNumber((_2640_v41).length)))) ? (_dafny.Seq.Concat((_this).f14, _2651_v49)) : ((_this).f14));
          let _rhs362 = _2652_v50;
          let _rhs363 = (new BigNumber((_2640_v41).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(419)), ((_2664_v1, _2665_v0) => function (_2666_i11) {
            return (_module.D4.create_DC13((_2664_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2664_v1).length))], (_this).f15, _2665_v0, _2665_v0, false)).dtor_cf21;
          })(_2591_v1, _2590_v0))).length));
          let _rhs364 = _2663_v57;
          let _lhs222 = _2660_v54;
          let _lhs223 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_2660_v54).length));
          _2651_v49 = _rhs361;
          _lhs222[_lhs223] = _rhs362;
          _2590_v0 = _rhs363;
          _2663_v57 = _rhs364;
          let _2667_v58;
          _2667_v58 = _dafny.MultiSet.fromElements(_2640_v41, _2640_v41);
          _2590_v0 = new BigNumber(((((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]) ? ((_2667_v58).Intersect(_2667_v58)) : ((_dafny.MultiSet.FromArray(_2641_v42)).Difference(_dafny.MultiSet.fromElements(_2640_v41))))).cardinality());
          let _2668_v59;
          let _nw424 = new _module.C4();
          _nw424.__ctor();
          _2668_v59 = _nw424;
          _2590_v0 = (_this).fm7((_2590_v0).plus(_2590_v0), globalState);
        } else {
          let _2669_v60;
          _2669_v60 = _dafny.MultiSet.fromElements(_2591_v1, (((_this).f15) ? (_2591_v1) : (_2591_v1)));
          let _2670_v61;
          _2670_v61 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_2591_v1, _2591_v1, _2591_v1), _2669_v60, (_2669_v60).Union(_2669_v60), _2669_v60);
          _2669_v60 = (_2670_v61)[_module.__default.safeIndex((_2590_v0).plus(_2590_v0), new BigNumber((_2670_v61).length))];
          let _2671_v62;
          let _nw425 = new _module.C4();
          _nw425.__ctor();
          _2671_v62 = _nw425;
          r0 = (_this).f15;
          (globalState).f4 = !(!((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))])) || ((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]);
          let _2672_v63;
          let _nw426 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _2672_v63 = _nw426;
          let _index412 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_2672_v63).length));
          (_2672_v63)[_index412] = _module.__default.safeDivisionInt(new BigNumber(-180), _2590_v0);
          let _index413 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_2672_v63).length));
          (_2672_v63)[_index413] = _2590_v0;
        }
        if ((_this).f15) {
          let _2673_v64;
          _2673_v64 = _module.D2.create_DC5(_2590_v0, (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]);
          let _pat_let_tv44 = _2648_v46;
          let _pat_let_tv45 = _2591_v1;
          let _pat_let_tv46 = _2591_v1;
          let _pat_let_tv47 = _2648_v46;
          let _pat_let_tv48 = _2591_v1;
          let _pat_let_tv49 = _2591_v1;
          let _2674_v65;
          let _nw427 = Array((new BigNumber(2)).toNumber());
          _nw427[(_dafny.ZERO).toNumber()] = function (_pat_let53_0) {
            return function (_2675_dt__update__tmp_h1) {
              return function (_pat_let54_0) {
                return function (_2676_dt__update_hcf5_h0) {
                  return function (_pat_let55_0) {
                    return function (_2677_dt__update_hcf6_h0) {
                      return _module.D2.create_DC5(_2676_dt__update_hcf5_h0, _2677_dt__update_hcf6_h0, (_2675_dt__update__tmp_h1).dtor_cf7);
                    }(_pat_let55_0);
                  }((_this).f15);
                }(_pat_let54_0);
              }((((_pat_let_tv47).contains((_pat_let_tv49)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_pat_let_tv48).length))])) ? ((_pat_let_tv44).get((_pat_let_tv46)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_pat_let_tv45).length))])) : (new BigNumber(-254))));
            }(_pat_let53_0);
          }(_module.D2.create_DC5(new BigNumber(910), !((_this).f15), _module.__default.fm2((_this).f15, globalState)));
          _nw427[(_dafny.ONE).toNumber()] = _2673_v64;
          _2674_v65 = _nw427;
          let _index414 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2674_v65).length));
          (_2674_v65)[_index414] = _2673_v64;
          let _index415 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_2674_v65).length));
          (_2674_v65)[_index415] = _2673_v64;
          let _2678_v66;
          let _init77 = ((_2679_v41) => function (_2680_i12) {
            return _2679_v41;
          })(_2640_v41);
          let _nw428 = Array((new BigNumber(18)).toNumber());
          for (let _i0_77 = 0; _i0_77 < new BigNumber(_nw428.length); _i0_77++) {
            _nw428[_i0_77] = _init77(new BigNumber(_i0_77));
          }
          _2678_v66 = _nw428;
          let _index416 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2678_v66).length));
          (_2678_v66)[_index416] = _dafny.Seq.update(_2640_v41, _module.__default.safeIndex(_2590_v0, new BigNumber((_2640_v41).length)), _2638_v39);
          let _index417 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_2678_v66).length));
          (_2678_v66)[_index417] = _dafny.Seq.Concat(_dafny.Seq.Concat(_2640_v41, _2640_v41), _dafny.Seq.Concat(_2640_v41, _2640_v41));
          _2590_v0 = (_dafny.ZERO).minus(_2590_v0);
          _2640_v41 = (_2678_v66)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_2678_v66).length))];
          let _index418 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
          let _rhs365 = (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))];
          let _rhs366 = !(true);
          let _rhs367 = !(_module.__default.fm2((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], globalState)) || ((_this).f15);
          let _lhs224 = globalState;
          let _lhs225 = globalState;
          let _lhs226 = _2591_v1;
          let _lhs227 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
          _lhs224.f4 = _rhs365;
          _lhs225.f4 = _rhs366;
          _lhs226[_lhs227] = _rhs367;
        } else {
          let _rhs368 = (_this).f15;
          let _rhs369 = _2590_v0;
          let _rhs370 = (_2590_v0).plus(_2590_v0);
          let _rhs371 = _dafny.Seq.update(_2641_v42, _module.__default.safeIndex((_this).fm8(!((_this).f15), (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], globalState), new BigNumber((_2641_v42).length)), _2640_v41);
          let _lhs228 = globalState;
          _lhs228.f4 = _rhs368;
          _2590_v0 = _rhs369;
          _2590_v0 = _rhs370;
          _2641_v42 = _rhs371;
          let _2681_v67;
          _2681_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2590_v0,_2648_v46);
          let _2682_v68;
          _2682_v68 = _module.D11.create_DC30(_2681_v67);
          _2682_v68 = _2682_v68;
          let _2683_v69;
          let _nw429 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _2683_v69 = _nw429;
          let _2684_v70;
          _2684_v70 = _dafny.Seq.of(_2683_v69, _2683_v69, _2683_v69, _2683_v69, _2683_v69);
          _2683_v69 = (_2684_v70)[_module.__default.safeIndex(_2590_v0, new BigNumber((_2684_v70).length))];
          let _2685_v71;
          _2685_v71 = _dafny.MultiSet.fromElements(_2590_v0);
          let _index419 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
          (_2591_v1)[_index419] = (_2685_v71).IsSubsetOf(_2685_v71);
          let _2686_v72;
          _2686_v72 = _dafny.Map.Empty.slice().updateUnsafe(_2590_v0,(_2590_v0).minus(_2590_v0));
          let _2687_v73;
          _2687_v73 = _dafny.Seq.of((_this).f14, (_this).f14, (_this).f14);
          let _2688_v74;
          _2688_v74 = _module.D2.create_DC4(_2687_v73);
          let _2689_v75;
          _2689_v75 = _dafny.MultiSet.fromElements((_this).f15, (_this).f15);
          let _index420 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
          let _index421 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
          let _rhs372 = (_2686_v72).update((_dafny.ZERO).minus(_2590_v0), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update((_this).f16, _module.__default.safeIndex(new BigNumber((_2648_v46).length), new BigNumber(((_this).f16).length)), _2590_v0))).cardinality()));
          let _rhs373 = _module.D2.create_DC4(_dafny.Seq.Concat(_dafny.Seq.of((_this).f14, (_this).f14), _2687_v73));
          let _rhs374 = !(_2689_v75).equals((_2689_v75).Union(_2689_v75));
          let _rhs375 = (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))];
          let _rhs376 = (((_this).f15) ? ((((_this).f15) ? (_2590_v0) : (new BigNumber(454)))) : ((_this).fm8(!(!((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))])), _module.__default.fm2(true, globalState), globalState)));
          let _lhs229 = _2591_v1;
          let _lhs230 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
          let _lhs231 = _2591_v1;
          let _lhs232 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
          _2686_v72 = _rhs372;
          _2688_v74 = _rhs373;
          _lhs229[_lhs230] = _rhs374;
          _lhs231[_lhs232] = _rhs375;
          _2590_v0 = _rhs376;
        }
        let _2690_v76;
        let _nw430 = Array((new BigNumber(5)).toNumber()).fill([]);
        _2690_v76 = _nw430;
        let _2691_v77;
        _2691_v77 = _2690_v76;
        let _2692_v78;
        let _nw431 = Array((new BigNumber(29)).toNumber());
        _nw431[(_dafny.ZERO).toNumber()] = _2690_v76;
        _nw431[(_dafny.ONE).toNumber()] = (((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]) ? (_2691_v77) : (_2691_v77));
        _nw431[(new BigNumber(2)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(3)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(4)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(5)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(6)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(7)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(8)).toNumber()] = _2690_v76;
        _nw431[(new BigNumber(9)).toNumber()] = _2690_v76;
        _nw431[(new BigNumber(10)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(11)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(12)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(13)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(14)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(15)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(16)).toNumber()] = _2690_v76;
        _nw431[(new BigNumber(17)).toNumber()] = _2690_v76;
        _nw431[(new BigNumber(18)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(19)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(20)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(21)).toNumber()] = _2690_v76;
        _nw431[(new BigNumber(22)).toNumber()] = _2690_v76;
        _nw431[(new BigNumber(23)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(24)).toNumber()] = (((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]) ? (_2690_v76) : (_2690_v76));
        _nw431[(new BigNumber(25)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(26)).toNumber()] = _2691_v77;
        _nw431[(new BigNumber(27)).toNumber()] = _2690_v76;
        _nw431[(new BigNumber(28)).toNumber()] = _2691_v77;
        _2692_v78 = _nw431;
        let _index422 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2692_v78).length));
        (_2692_v78)[_index422] = _2691_v77;
        let _index423 = _module.__default.safeIndex(new BigNumber(899), new BigNumber((_2692_v78).length));
        (_2692_v78)[_index423] = _2690_v76;
        if ((_this).f15) {
          _2648_v46 = (_2648_v46).update((_2590_v0).isLessThanOrEqualTo((_dafny.ZERO).minus((((_2648_v46).contains((_this).f15)) ? ((_2648_v46).get((_this).f15)) : (new BigNumber(372))))), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_2640_v41, _module.__default.safeIndex(_2590_v0, new BigNumber((_2640_v41).length)), new _dafny.CodePoint('n'.codePointAt(0)))).length), _2590_v0));
          let _2693_v79;
          _2693_v79 = _dafny.Set.fromElements((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], (_this).f15);
          let _2694_v80;
          let _nw432 = new _module.C1();
          _nw432.__ctor((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], (_this).f15, (_this).f13, (_this).f14, _this.f12);
          _2694_v80 = _nw432;
          let _2695_v81;
          _2695_v81 = _module.D3.create_DC10((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], new BigNumber((_2693_v79).length), _2694_v80);
          let _2696_v82;
          _2696_v82 = _module.D3.create_DC11(_2695_v81);
          let _pat_let_tv50 = _2695_v81;
          let _2697_v83;
          let _nw433 = Array((new BigNumber(19)).toNumber());
          _nw433[(_dafny.ZERO).toNumber()] = _2696_v82;
          _nw433[(_dafny.ONE).toNumber()] = _module.D3.create_DC11(_2695_v81);
          _nw433[(new BigNumber(2)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(3)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(4)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(5)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(6)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(7)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(8)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(9)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(10)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(11)).toNumber()] = _module.D3.create_DC11(_2695_v81);
          _nw433[(new BigNumber(12)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(13)).toNumber()] = function (_pat_let56_0) {
            return function (_2698_dt__update__tmp_h11) {
              return function (_pat_let57_0) {
                return function (_2699_dt__update_hcf17_h0) {
                  return _module.D3.create_DC11(_2699_dt__update_hcf17_h0);
                }(_pat_let57_0);
              }(_pat_let_tv50);
            }(_pat_let56_0);
          }(_2696_v82);
          _nw433[(new BigNumber(14)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(15)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(16)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(17)).toNumber()] = _2696_v82;
          _nw433[(new BigNumber(18)).toNumber()] = _2696_v82;
          _2697_v83 = _nw433;
          let _2700_v84;
          _2700_v84 = _2697_v83;
          let _2701_v85;
          _2701_v85 = _dafny.MultiSet.fromElements((_2700_v84));
          let _2702_v86;
          _2702_v86 = _dafny.MultiSet.fromElements(new BigNumber(((((_this).f15) ? (_2701_v85) : (_2701_v85))).cardinality()));
          _2702_v86 = (_2702_v86).Difference(_2702_v86);
          _2590_v0 = _module.__default.safeDivisionInt(_2590_v0, _2590_v0);
          let _2703_v87;
          let _nw434 = new _module.C4();
          _nw434.__ctor();
          _2703_v87 = _nw434;
          let _2704_v88;
          _2704_v88 = _dafny.Seq.of(_2591_v1);
          let _rhs377 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_2590_v0).minus(_2590_v0)), (new BigNumber(734)).plus((_dafny.ZERO).minus(new BigNumber((_2704_v88).length))));
          let _rhs378 = _2703_v87;
          let _rhs379 = (_this).f15;
          _2590_v0 = _rhs377;
          _2703_v87 = _rhs378;
          r0 = _rhs379;
          _2640_v41 = _2640_v41;
        } else {
          _2590_v0 = (_this).fm5((_this).f15, globalState);
          let _2705_v89;
          let _nw435 = new _module.C2();
          _nw435.__ctor((_this).f16, (_this).f15, ((_this).f13).update(false, _2593_v3), (_this).f14, _this.f12);
          _2705_v89 = _nw435;
          let _2706_v90;
          _2706_v90 = _dafny.Map.Empty.slice().updateUnsafe(_2590_v0,_dafny.Map.Empty.slice().updateUnsafe(_2705_v89,_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15)));
          let _2707_v91;
          _2707_v91 = _dafny.Map.Empty.slice().updateUnsafe(_2705_v89,_2593_v3);
          let _2708_v92;
          _2708_v92 = _dafny.Map.Empty.slice().updateUnsafe((((_2706_v90).contains(new BigNumber(-934))) ? ((_2706_v90).get(new BigNumber(-934))) : (_2707_v91)),(_this).fm5((_this).f15, globalState));
          _2590_v0 = _module.__default.safeModuloInt((new BigNumber((_2640_v41).length)).multipliedBy((((_2708_v92).contains(_dafny.Map.Empty.slice().updateUnsafe(_2705_v89,_2593_v3))) ? ((_2708_v92).get(_dafny.Map.Empty.slice().updateUnsafe(_2705_v89,_2593_v3))) : (_2590_v0))), _2590_v0);
          let _2709_v93;
          _2709_v93 = _dafny.Set.fromElements(_2590_v0, new BigNumber(450));
          let _2710_v94;
          _2710_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2709_v93,(_this).f15);
          _2710_v94 = (_2710_v94).update(_2709_v93, (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]);
          (globalState).f4 = (_this).f15;
          let _2711_v95;
          let _nw436 = new _module.C2();
          _nw436.__ctor(_dafny.Seq.update(_dafny.Seq.of(_2590_v0), _module.__default.safeIndex(_2590_v0, new BigNumber((_dafny.Seq.of(_2590_v0)).length)), _2590_v0), !((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]), (_this).f13, (_this).f14, _2705_v89.f12);
          _2711_v95 = _nw436;
        }
        _2640_v41 = _dafny.Seq.Concat(_2640_v41, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), ((_2712_v39) => function (_2713_i13) {
          return _2712_v39;
        })(_2638_v39)), _2640_v41));
      } else {
        (globalState).f4 = (_this).f15;
        let _2714_v96;
        _2714_v96 = _dafny.Map.Empty.slice().updateUnsafe(_2640_v41,(_this).f15);
        let _2715_v97;
        _2715_v97 = _module.D12.create_DC37(_2640_v41, (_this).f15);
        let _index424 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
        (_2591_v1)[_index424] = (((_2714_v96).contains(_2640_v41)) ? ((_2714_v96).get(_2640_v41)) : ((_2715_v97).dtor_cf60));
        if (!((_this).f15)) {
          let _2716_v98;
          let _nw437 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _2716_v98 = _nw437;
          let _2717_v99;
          let _nw438 = new _module.C1();
          _nw438.__ctor((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], (_this).f13, (_this).f14, _this.f12);
          _2717_v99 = _nw438;
          let _2718_v100;
          _2718_v100 = _module.D11.create_DC33((_this).f15, _2590_v0, (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], _2717_v99, (_2717_v99).f17);
          let _index425 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_2716_v98).length));
          (_2716_v98)[_index425] = (_2590_v0).multipliedBy((_dafny.ZERO).minus((_2718_v100).dtor_cf49));
          let _2719_v101;
          _2719_v101 = _module.D17.create_DC46((_this).f13);
          let _2720_v102;
          let _nw439 = new _module.C2();
          _nw439.__ctor((_this).f16, _2717_v99.f18, (_2719_v101).dtor_cf69, (_this).f14, _this.f12);
          _2720_v102 = _nw439;
          let _2721_v103;
          _2721_v103 = _dafny.Set.fromElements(_2590_v0);
          let _index426 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_2716_v98).length));
          let _rhs380 = new BigNumber(((_dafny.Set.fromElements(_2590_v0)).Intersect(_2721_v103)).length);
          let _rhs381 = ((_2717_v99.f18) ? (_2590_v0) : ((new BigNumber(620)).multipliedBy(_2590_v0)));
          let _rhs382 = _2720_v102;
          let _rhs383 = !(!(new BigNumber(58)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_2590_v0))));
          let _lhs233 = _2716_v98;
          let _lhs234 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_2716_v98).length));
          let _lhs235 = _2717_v99;
          _2590_v0 = _rhs380;
          _lhs233[_lhs234] = _rhs381;
          _2720_v102 = _rhs382;
          _lhs235.f18 = _rhs383;
          let _2722_v104;
          _2722_v104 = _module.D0.create_DC0(!(false));
          let _2723_v105;
          _2723_v105 = _module.D4.create_DC13((_2722_v104).dtor_cf0, !((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]), new BigNumber(519), (_this).fm5(_2717_v99.f18, globalState), (_2717_v99).f17);
          let _2724_v106;
          _2724_v106 = _dafny.Map.Empty.slice().updateUnsafe((_2723_v105).dtor_cf21,_2590_v0);
          let _2725_v107;
          _2725_v107 = _dafny.MultiSet.fromElements(_2590_v0, new BigNumber((_2640_v41).length));
          _2724_v106 = (_2724_v106).update(new BigNumber(((_2725_v107).Intersect(_2725_v107)).cardinality()), _2590_v0);
          let _2726_v108;
          let _nw440 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _2726_v108 = _nw440;
          _2726_v108 = _2726_v108;
          r0 = (_2717_v99).f17;
          let _2727_v109;
          let _nw441 = Array((new BigNumber(15)).toNumber()).fill([]);
          _2727_v109 = _nw441;
          let _2728_v110;
          _2728_v110 = _dafny.MultiSet.fromElements((_this).f15, false, (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], _2717_v99.f18, (_2717_v99).f17);
          let _2729_v111;
          let _nw442 = Array((new BigNumber(2)).toNumber());
          _nw442[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_2717_v99.f18);
          _nw442[(_dafny.ONE).toNumber()] = _2728_v110;
          _2729_v111 = _nw442;
          let _index427 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_2727_v109).length));
          (_2727_v109)[_index427] = _2729_v111;
          let _index428 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_2727_v109).length));
          (_2727_v109)[_index428] = _2729_v111;
        } else {
          let _2730_v112;
          let _nw443 = Array((new BigNumber(2)).toNumber()).fill(_module.D14.Default());
          _2730_v112 = _nw443;
          let _2731_v113;
          let _nw444 = new _module.C3();
          _nw444.__ctor((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], _this.f12, !((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]), (_this).f13, (_this).f14);
          _2731_v113 = _nw444;
          let _2732_v114;
          _2732_v114 = _module.D14.create_DC40(_2731_v113);
          let _index429 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_2730_v112).length));
          (_2730_v112)[_index429] = _2732_v114;
          let _index430 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_2730_v112).length));
          let _rhs384 = _2732_v114;
          let _rhs385 = _2640_v41;
          let _rhs386 = _2638_v39;
          let _rhs387 = (_dafny.ZERO).minus(_2590_v0);
          let _lhs236 = _2730_v112;
          let _lhs237 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_2730_v112).length));
          _lhs236[_lhs237] = _rhs384;
          _2640_v41 = _rhs385;
          _2638_v39 = _rhs386;
          _2590_v0 = _rhs387;
          _2590_v0 = ((_this).f16)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_2640_v41, _2640_v41)).length), new BigNumber(((_this).f16).length))];
          _2638_v39 = _2638_v39;
          (globalState).f4 = (_2590_v0).isLessThan(_2590_v0);
          let _2733_v115;
          let _init78 = ((_2734_v0) => function (_2735_i14) {
            return (_2735_i14).minus(_2734_v0);
          })(_2590_v0);
          let _nw445 = Array((new BigNumber(19)).toNumber());
          for (let _i0_78 = 0; _i0_78 < new BigNumber(_nw445.length); _i0_78++) {
            _nw445[_i0_78] = _init78(new BigNumber(_i0_78));
          }
          _2733_v115 = _nw445;
          let _index431 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_2733_v115).length));
          (_2733_v115)[_index431] = _module.__default.safeDivisionInt(_2590_v0, _2590_v0);
          let _2736_v116;
          _2736_v116 = _dafny.MultiSet.fromElements((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]);
          let _2737_v117;
          _2737_v117 = _dafny.MultiSet.fromElements(_2590_v0);
          let _index432 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_2733_v115).length));
          let _rhs388 = (_module.__default.fm20((_this).f15, _2590_v0, globalState)).IsSubsetOf(_2736_v116);
          let _rhs389 = _2590_v0;
          let _rhs390 = new BigNumber(((_2731_v113).f14).length);
          let _rhs391 = _module.__default.fm59(_2737_v117, (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], _module.__default.fm2((_2731_v113).f15, globalState), new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(_2590_v0, _2590_v0, _2590_v0, new BigNumber(841)))).Union(_2737_v117)).cardinality()), globalState);
          let _lhs238 = globalState;
          let _lhs239 = _2733_v115;
          let _lhs240 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_2733_v115).length));
          _lhs238.f4 = _rhs388;
          _lhs239[_lhs240] = _rhs389;
          _2590_v0 = _rhs390;
          _2593_v3 = _rhs391;
        }
        let _2738_v118;
        let _nw446 = new _module.C3();
        _nw446.__ctor((_this).f15, _this.f12, !((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]), _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_2593_v3), _module.__default.fm15((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], _2590_v0, _2590_v0, new BigNumber(466), globalState));
        _2738_v118 = _nw446;
        let _2739_v119;
        _2739_v119 = _dafny.Map.Empty.slice().updateUnsafe(_2738_v118,_2590_v0);
        let _2740_v120;
        let _nw447 = new _module.C2();
        _nw447.__ctor((_this).f16, !(_2739_v119).equals(_2739_v119), _dafny.Map.Empty.slice().updateUnsafe((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))],_2593_v3), (_this).f14, _this.f12);
        _2740_v120 = _nw447;
        _2740_v120 = _2740_v120;
        let _index433 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length));
        (_2591_v1)[_index433] = (_this).f15;
      }
      let _hi15 = new BigNumber(((_2593_v3).Merge(_2593_v3)).length);
      for (let _2741_i15 = _module.__default.safeModuloInt((_this).fm5((_this).f15, globalState), _2590_v0); _2741_i15.isLessThan(_hi15); _2741_i15 = _2741_i15.plus(_dafny.ONE)) {
        let _2742_v121;
        let _init79 = ((_2743_i15) => function (_2744_i16) {
          return (_dafny.MultiSet.fromElements((_this).f16, (_this).f16, (_this).f16)).Union(_dafny.MultiSet.fromElements((_this).f16, (_this).f16, _dafny.Seq.Create(_module.__default.abs(new BigNumber(402)), ((_2745_i15) => function (_2746_i17) {
            return _2745_i15;
          })(_2743_i15))));
        })(_2741_i15);
        let _nw448 = Array((new BigNumber(13)).toNumber());
        for (let _i0_79 = 0; _i0_79 < new BigNumber(_nw448.length); _i0_79++) {
          _nw448[_i0_79] = _init79(new BigNumber(_i0_79));
        }
        _2742_v121 = _nw448;
        let _index434 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2742_v121).length));
        (_2742_v121)[_index434] = _dafny.MultiSet.fromElements(_dafny.Seq.update((_this).f16, _module.__default.safeIndex((_this).fm7((_dafny.ZERO).minus(((_this).f16)[_module.__default.safeIndex(_2590_v0, new BigNumber(((_this).f16).length))]), globalState), new BigNumber(((_this).f16).length)), _2590_v0), (_this).f16);
        let _2747_v122;
        _2747_v122 = _dafny.Seq.of((_this).f16, (_this).f16);
        let _2748_v123;
        _2748_v123 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-238)), ((_2749_i15) => function (_2750_i18) {
          return _2749_i15;
        })(_2741_i15)), _dafny.Seq.update((_2747_v122)[_module.__default.safeIndex((_dafny.ZERO).minus(_2590_v0), new BigNumber((_2747_v122).length))], _module.__default.safeIndex(_2741_i15, new BigNumber(((_2747_v122)[_module.__default.safeIndex((_dafny.ZERO).minus(_2590_v0), new BigNumber((_2747_v122).length))]).length)), _2741_i15));
        let _2751_v124;
        _2751_v124 = _dafny.Map.Empty.slice().updateUnsafe((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))],_2748_v123);
        let _2752_v125;
        _2752_v125 = _dafny.Map.Empty.slice().updateUnsafe((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))],(((_2751_v124).contains(true)) ? ((_2751_v124).get(true)) : (_2748_v123)));
        let _index435 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2742_v121).length));
        (_2742_v121)[_index435] = (_2748_v123).Difference((((_2751_v124).contains(false)) ? ((_2751_v124).get(false)) : (_2748_v123)));
        let _2753_v126;
        let _init80 = ((_2754_v39) => function (_2755_i19) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), ((_2756_v39) => function (_2757_i20) {
            return _2756_v39;
          })(_2754_v39));
        })(_2638_v39);
        let _nw449 = Array((new BigNumber(14)).toNumber());
        for (let _i0_80 = 0; _i0_80 < new BigNumber(_nw449.length); _i0_80++) {
          _nw449[_i0_80] = _init80(new BigNumber(_i0_80));
        }
        _2753_v126 = _nw449;
        let _index436 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_2753_v126).length));
        (_2753_v126)[_index436] = _module.__default.fm40(function (_pat_let58_0) {
          return function (_2758_dt__update__tmp_h12) {
            return function (_pat_let59_0) {
              return function (_2759_dt__update_hcf24_h0) {
                return _module.D4.create_DC14(_2759_dt__update_hcf24_h0);
              }(_pat_let59_0);
            }((_this).f15);
          }(_pat_let58_0);
        }(_module.D4.create_DC14((_this).f15)), new BigNumber((_2650_v48).length), (_this).f15, (_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))], globalState);
        let _index437 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_2753_v126).length));
        (_2753_v126)[_index437] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_2760_v39) => function (_2761_i21) {
          return _2760_v39;
        })(_2638_v39)), _2640_v41);
        let _2762_v127;
        let _nw450 = Array((new BigNumber(5)).toNumber());
        _nw450[(_dafny.ZERO).toNumber()] = _2741_i15;
        _nw450[(_dafny.ONE).toNumber()] = _2741_i15;
        _nw450[(new BigNumber(2)).toNumber()] = _2741_i15;
        _nw450[(new BigNumber(3)).toNumber()] = _2741_i15;
        _nw450[(new BigNumber(4)).toNumber()] = _2590_v0;
        _2762_v127 = _nw450;
        let _2763_v128;
        _2763_v128 = _dafny.Seq.of(_2762_v127, _2762_v127);
        _2763_v128 = _dafny.Seq.of(_2762_v127);
        (globalState).f4 = (_this).f15;
      }
      r0 = !(((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))]) === ((_2591_v1)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_2591_v1).length))])) || ((_2590_v0).isLessThan(_2590_v0));
      return r0;
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
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _module.D0.Default();
      let r1 = _dafny.Map.Empty;
      let r2 = false;
      let _2764_v0;
      _2764_v0 = false;
      if (_2764_v0) {
        let _2765_v1;
        _2765_v1 = _dafny.Seq.UnicodeFromString("hlnnql");
        let _2766_v2;
        _2766_v2 = new BigNumber(-982);
        let _2767_v4;
        _2767_v4 = _module.D0.create_DC2(_module.__default.safeModuloInt(new BigNumber((_2765_v1).length), _2766_v2), (_module.D0.create_DC2(new BigNumber((function () {
  let _coll68 = new _dafny.Map();
  for (const _compr_68 of _dafny.IntegerRange(new BigNumber(603), new BigNumber(114))) {
    let _2768_v3 = _compr_68;
    if (((new BigNumber(603)).isLessThanOrEqualTo(_2768_v3)) && ((_2768_v3).isLessThan(new BigNumber(114)))) {
      _coll68.push([(_2768_v3).plus(new BigNumber((_dafny.MultiSet.fromElements(_2764_v0, _2764_v0)).cardinality())),_2764_v0]);
    }
  }
  return _coll68;
}()).length), _2764_v0)).dtor_cf2);
        let _source40 = _2767_v4;
        if (_source40.is_DC1) {
          let _2769_v5;
          _2769_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2764_v0,(p1)[_module.__default.safeIndex(_2766_v2, new BigNumber((p1).length))]);
          let _2770_v6;
          _2770_v6 = _dafny.MultiSet.fromElements(new BigNumber((_2769_v5).length));
          _2770_v6 = _2770_v6;
          let _2771_v7;
          let _2772_v8;
          let _2773_v9;
          let _2774_v10;
          let _out131;
          let _out132;
          let _out133;
          let _out134;
          let _outcollector34 = _module.__default.m0((_2767_v4).dtor_cf2, _2767_v4, _module.__default.fm2(_2764_v0, globalState), globalState);
          _out131 = _outcollector34[0];
          _out132 = _outcollector34[1];
          _out133 = _outcollector34[2];
          _out134 = _outcollector34[3];
          _2771_v7 = _out131;
          _2772_v8 = _out132;
          _2773_v9 = _out133;
          _2774_v10 = _out134;
          (globalState).f4 = _module.__default.fm2(!(_2764_v0), globalState);
          _2773_v9 = new BigNumber(-206);
        } else if (_source40.is_DC2) {
          let _2775___mcc_h0 = (_source40).cf1;
          let _2776___mcc_h1 = (_source40).cf2;
          let _2777_cf2 = _2776___mcc_h1;
          let _2778_cf1 = _2775___mcc_h0;
          let _2779_v11;
          let _nw451 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2779_v11 = _nw451;
          let _index438 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_2779_v11).length));
          (_2779_v11)[_index438] = _2778_cf1;
          let _2780_v12;
          _2780_v12 = _dafny.MultiSet.fromElements(_2767_v4);
          let _index439 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_2779_v11).length));
          (_2779_v11)[_index439] = (((_2780_v12).contains(_2767_v4)) ? ((_2780_v12).get(_2767_v4)) : (new BigNumber(863)));
          let _2781_v13;
          let _nw452 = Array((_dafny.ONE).toNumber());
          _nw452[(_dafny.ZERO).toNumber()] = _module.D0.create_DC2((_dafny.ZERO).minus((_dafny.ZERO).minus(_2778_cf1)), _2764_v0);
          _2781_v13 = _nw452;
          let _index440 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_2781_v13).length));
          (_2781_v13)[_index440] = _module.D0.create_DC2(new BigNumber(-257), _2764_v0);
          let _2782_v14;
          _2782_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2766_v2);
          let _2783_v15;
          _2783_v15 = _dafny.Map.Empty.slice().updateUnsafe(!(_2777_cf2),(((_2782_v14).contains(p0)) ? ((_2782_v14).get(p0)) : ((_2779_v11)[_module.__default.safeIndex(new BigNumber(677), new BigNumber((_2779_v11).length))])));
          let _2784_v16;
          _2784_v16 = _dafny.Seq.of(_2783_v15, _2783_v15, _2783_v15, _2783_v15);
          let _index441 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_2781_v13).length));
          (_2781_v13)[_index441] = _module.__default.fm3(_2784_v16, globalState);
          _2765_v1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(504)), ((_2785_p0) => function (_2786_i0) {
            return _2785_p0;
          })(p0));
          _2778_cf1 = (_2779_v11)[_module.__default.safeIndex(new BigNumber(677), new BigNumber((_2779_v11).length))];
        } else {
          let _2787___mcc_h2 = (_source40).cf0;
          let _2788_cf0 = _2787___mcc_h2;
          _2766_v2 = _2766_v2;
          _2766_v2 = _2766_v2;
          let _2789_v17;
          _2789_v17 = _dafny.Set.fromElements(_2764_v0, _2764_v0, _2764_v0);
          let _2790_v18;
          let _nw453 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _2790_v18 = _nw453;
          let _2791_v19;
          _2791_v19 = _dafny.Seq.of(true, _2764_v0);
          let _2792_v20;
          _2792_v20 = _module.D0.create_DC0(_2788_cf0);
          let _2793_v21;
          _2793_v21 = _dafny.Seq.of(_2792_v20);
          let _2794_v22;
          let _nw454 = Array((new BigNumber(21)).toNumber());
          _nw454[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2789_v17,new BigNumber((_dafny.Seq.of(_2764_v0, false, _2764_v0, _2764_v0, _2788_cf0)).length))).length);
          _nw454[(_dafny.ONE).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(2)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(_2764_v0)).cardinality())).plus(_2766_v2);
          _nw454[(new BigNumber(3)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(4)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(5)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2765_v1, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(516)), ((_2795_p0) => function (_2796_i1) {
            return _2795_p0;
          })(p0)), _module.__default.safeIndex(_2766_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(516)), ((_2797_p0) => function (_2798_i1) {
            return _2797_p0;
          })(p0))).length)), p0))).length);
          _nw454[(new BigNumber(7)).toNumber()] = ((_dafny.ZERO).minus(_2766_v2)).multipliedBy(_2766_v2);
          _nw454[(new BigNumber(8)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(9)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(10)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(680), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2764_v0,_2790_v18)).length));
          _nw454[(new BigNumber(12)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(13)).toNumber()] = new BigNumber(555);
          _nw454[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_2791_v19)).cardinality());
          _nw454[(new BigNumber(15)).toNumber()] = new BigNumber(-863);
          _nw454[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.of(true, true)).length);
          _nw454[(new BigNumber(17)).toNumber()] = ((_dafny.ZERO).minus(_2766_v2)).multipliedBy(new BigNumber((_2765_v1).length));
          _nw454[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.update(_2793_v21, _module.__default.safeIndex(_2766_v2, new BigNumber((_2793_v21).length)), _module.__default.fm4(p0, new BigNumber(-652), globalState))).length);
          _nw454[(new BigNumber(19)).toNumber()] = _2766_v2;
          _nw454[(new BigNumber(20)).toNumber()] = _2766_v2;
          _2794_v22 = _nw454;
          let _index442 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length));
          (_2794_v22)[_index442] = _2766_v2;
          let _index443 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length));
          (_2794_v22)[_index443] = ((_2764_v0) ? (_2766_v2) : (_2766_v2));
          let _2799_v23;
          _2799_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2788_cf0,_2788_cf0);
          let _2800_v24;
          _2800_v24 = _dafny.Map.Empty.slice().updateUnsafe(true,_2799_v23);
          let _2801_v25;
          _2801_v25 = _dafny.MultiSet.fromElements((_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))]);
          let _2802_v26;
          _2802_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2788_cf0,_2801_v25);
          let _2803_v27;
          let _nw455 = Array((new BigNumber(26)).toNumber());
          _nw455[(_dafny.ZERO).toNumber()] = _2801_v25;
          _nw455[(_dafny.ONE).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(2)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(3)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(4)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(5)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(6)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(7)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_2799_v23).length), (_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))]);
          _nw455[(new BigNumber(9)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(10)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(11)).toNumber()] = (((_2802_v26).contains(!(_2788_cf0))) ? ((_2802_v26).get(!(_2788_cf0))) : (_module.__default.fm28(_2765_v1, _2764_v0, _2766_v2, globalState)));
          _nw455[(new BigNumber(12)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(13)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(14)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(15)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.FromArray(p1);
          _nw455[(new BigNumber(17)).toNumber()] = (_2801_v25).update((_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))], _module.__default.abs((_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))]));
          _nw455[(new BigNumber(18)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(19)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(20)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(21)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements((_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))], _2766_v2, (_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))], (_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))]);
          _nw455[(new BigNumber(23)).toNumber()] = _2801_v25;
          _nw455[(new BigNumber(24)).toNumber()] = _dafny.MultiSet.fromElements((_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))]);
          _nw455[(new BigNumber(25)).toNumber()] = _2801_v25;
          _2803_v27 = _nw455;
          let _2804_v28;
          let _nw456 = new _module.C6();
          _nw456.__ctor(p1, !(_2788_cf0), (_2800_v24).Merge(_2800_v24), _dafny.Seq.of(_2764_v0, (_2791_v19)[_module.__default.safeIndex((_2794_v22)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_2794_v22).length))], new BigNumber((_2791_v19).length))], true, _2764_v0), _2803_v27);
          _2804_v28 = _nw456;
        }
        let _2805_v29;
        let _nw457 = new _module.C0();
        _nw457.__ctor();
        _2805_v29 = _nw457;
        _2764_v0 = !_dafny.Seq.contains(p1, _module.__default.safeDivisionInt(_2766_v2, _2766_v2));
        _2764_v0 = _2764_v0;
        let _2806_v30;
        _2806_v30 = _module.D2.create_DC5(new BigNumber((_dafny.Set.fromElements(_2764_v0)).length), _2764_v0, false);
        let _2807_v31;
        _2807_v31 = _dafny.Set.fromElements(_module.D2.create_DC5(_2766_v2, _2764_v0, _2764_v0), _2806_v30, _2806_v30);
        let _2808_v32;
        _2808_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2766_v2,_2764_v0);
        let _rhs392 = _2766_v2;
        let _rhs393 = _module.__default.safeModuloInt(new BigNumber(116), new BigNumber(((_2807_v31).Union(_2807_v31)).length));
        let _rhs394 = _2764_v0;
        let _rhs395 = (false) === ((((_2808_v32).contains(_2766_v2)) ? ((_2808_v32).get(_2766_v2)) : (_2764_v0)));
        let _rhs396 = ((_2766_v2).multipliedBy(_2766_v2)).isEqualTo(_2766_v2);
        let _lhs241 = globalState;
        let _lhs242 = globalState;
        let _lhs243 = globalState;
        _2766_v2 = _rhs392;
        _2766_v2 = _rhs393;
        _lhs241.f4 = _rhs394;
        _lhs242.f4 = _rhs395;
        _lhs243.f4 = _rhs396;
      } else {
        let _2809_v33;
        _2809_v33 = new BigNumber(-101);
        let _2810_v34;
        _2810_v34 = _dafny.MultiSet.fromElements(_2809_v33);
        let _2811_v35;
        _2811_v35 = _dafny.Seq.of(_2764_v0, _2764_v0);
        let _2812_v36;
        _2812_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2764_v0,_2764_v0);
        let _2813_v37;
        _2813_v37 = _dafny.Map.Empty.slice().updateUnsafe(false,_2812_v36);
        let _2814_v38;
        let _nw458 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
        _2814_v38 = _nw458;
        let _2815_v39;
        let _nw459 = new _module.C2();
        _nw459.__ctor(_dafny.Seq.of(_2809_v33), _2764_v0, _2813_v37, _2811_v35, _2814_v38);
        _2815_v39 = _nw459;
        let _2816_v40;
        _2816_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2815_v39,_2764_v0);
        let _2817_v41;
        _2817_v41 = _dafny.Seq.of(_dafny.Seq.of(_2764_v0), _2811_v35);
        let _2818_v42;
        let _nw460 = Array((new BigNumber(6)).toNumber());
        _nw460[(_dafny.ZERO).toNumber()] = new BigNumber((_2811_v35).length);
        _nw460[(_dafny.ONE).toNumber()] = new BigNumber(120);
        _nw460[(new BigNumber(2)).toNumber()] = _2809_v33;
        _nw460[(new BigNumber(3)).toNumber()] = new BigNumber((_module.__default.fm45(_2764_v0, _2809_v33, _2809_v33, globalState)).length);
        _nw460[(new BigNumber(4)).toNumber()] = _2809_v33;
        _nw460[(new BigNumber(5)).toNumber()] = new BigNumber(540);
        _2818_v42 = _nw460;
        let _2819_v43;
        _2819_v43 = _dafny.MultiSet.fromElements(_2818_v42, _2818_v42);
        let _2820_v44;
        let _nw461 = new _module.C3();
        _nw461.__ctor(true, _2814_v38, _2764_v0, _2813_v37, _2811_v35);
        _2820_v44 = _nw461;
        let _2821_v45;
        _2821_v45 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2809_v33),_2820_v44);
        let _2822_v46;
        _2822_v46 = _dafny.Set.fromElements(_2809_v33);
        let _2823_v47;
        let _nw462 = new _module.C1();
        _nw462.__ctor((_2820_v44).f15, (_2820_v44).f15, _2813_v37, (_2820_v44).f14, _2814_v38);
        _2823_v47 = _nw462;
        let _2824_v48;
        _2824_v48 = _dafny.Seq.of(_2823_v47);
        let _2825_v49;
        _2825_v49 = _module.D11.create_DC33(_2764_v0, (_dafny.ZERO).minus(new BigNumber((_2822_v46).length)), _2764_v0, (_2824_v48)[_module.__default.safeIndex(_2809_v33, new BigNumber((_2824_v48).length))], (_2823_v47).f17);
        let _2826_v50;
        let _nw463 = Array((new BigNumber(25)).toNumber());
        _nw463[(_dafny.ZERO).toNumber()] = (new BigNumber((_2810_v34).cardinality())).multipliedBy(_2809_v33);
        _nw463[(_dafny.ONE).toNumber()] = new BigNumber((_module.__default.fm20(_2764_v0, _2809_v33, globalState)).cardinality());
        _nw463[(new BigNumber(2)).toNumber()] = new BigNumber((_2811_v35).length);
        _nw463[(new BigNumber(3)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(4)).toNumber()] = new BigNumber(964);
        _nw463[(new BigNumber(5)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_2809_v33));
        _nw463[(new BigNumber(7)).toNumber()] = (_2809_v33).multipliedBy(new BigNumber((_2816_v40).length));
        _nw463[(new BigNumber(8)).toNumber()] = new BigNumber(732);
        _nw463[(new BigNumber(9)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(10)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(11)).toNumber()] = (_2809_v33).plus(_2809_v33);
        _nw463[(new BigNumber(12)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(13)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(14)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(15)).toNumber()] = new BigNumber(((_2817_v41)[_module.__default.safeIndex(new BigNumber((_2819_v43).cardinality()), new BigNumber((_2817_v41).length))]).length);
        _nw463[(new BigNumber(16)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(17)).toNumber()] = (p1)[_module.__default.safeIndex(_2809_v33, new BigNumber((p1).length))];
        _nw463[(new BigNumber(18)).toNumber()] = ((_2764_v0) ? (_2809_v33) : ((_dafny.ZERO).minus(_2809_v33)));
        _nw463[(new BigNumber(19)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(20)).toNumber()] = new BigNumber((((_2764_v0) ? (_2821_v45) : (_dafny.Map.Empty.slice().updateUnsafe(_2809_v33,_2820_v44)))).length);
        _nw463[(new BigNumber(21)).toNumber()] = _2809_v33;
        _nw463[(new BigNumber(22)).toNumber()] = (_2820_v44).fm7((_dafny.ZERO).minus(_2809_v33), globalState);
        _nw463[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(_2809_v33, _2809_v33);
        _nw463[(new BigNumber(24)).toNumber()] = (_2825_v49).dtor_cf49;
        _2826_v50 = _nw463;
        let _index444 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length));
        (_2818_v42)[_index444] = _2809_v33;
        let _index445 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length));
        (_2818_v42)[_index445] = _2809_v33;
        let _2827_v51;
        let _nw464 = new _module.C0();
        _nw464.__ctor();
        _2827_v51 = _nw464;
        let _2828_v52;
        _2828_v52 = _2827_v51;
        let _2829_v53;
        _2829_v53 = _dafny.Map.Empty.slice().updateUnsafe(false,_2828_v52);
        let _2830_v54;
        _2830_v54 = _dafny.Map.Empty.slice().updateUnsafe(_2829_v53,p0);
        _2830_v54 = _2830_v54;
        if (!(_dafny.Seq.IsPrefixOf(p1, p1))) {
          let _2831_v55;
          _2831_v55 = _dafny.Map.Empty.slice().updateUnsafe((_2818_v42)[_module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length))],_2809_v33);
          _2831_v55 = _dafny.Map.Empty.slice().updateUnsafe((_2818_v42)[_module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length))],(p1)[_module.__default.safeIndex((_2818_v42)[_module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length))], new BigNumber((p1).length))]);
          (_2823_v47).f18 = _2823_v47.f18;
          let _2832_v56;
          let _nw465 = new _module.C5();
          _nw465.__ctor(!((_2820_v44).f15), _2820_v44.f12, _2813_v37, (_2820_v44).f14);
          _2832_v56 = _nw465;
          let _2833_v57;
          _2833_v57 = _dafny.Map.Empty.slice().updateUnsafe((_2823_v47).f17,_2832_v56);
          let _2834_v58;
          _2834_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2826_v50,_2826_v50);
          _2833_v57 = (_2833_v57).update(!((_2834_v58).update(_2818_v42, _2826_v50)).equals(_2834_v58), _2832_v56);
          let _2835_v59;
          let _init81 = ((_2836_v46) => function (_2837_i2) {
            return _2836_v46;
          })(_2822_v46);
          let _nw466 = Array((new BigNumber(2)).toNumber());
          for (let _i0_81 = 0; _i0_81 < new BigNumber(_nw466.length); _i0_81++) {
            _nw466[_i0_81] = _init81(new BigNumber(_i0_81));
          }
          _2835_v59 = _nw466;
          let _index446 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_2835_v59).length));
          (_2835_v59)[_index446] = _2822_v46;
          let _index447 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_2835_v59).length));
          (_2835_v59)[_index447] = _2822_v46;
          let _2838_v60;
          let _init82 = ((_2839_v44) => function (_2840_i3) {
            return (_2839_v44).f15;
          })(_2820_v44);
          let _nw467 = Array((new BigNumber(13)).toNumber());
          for (let _i0_82 = 0; _i0_82 < new BigNumber(_nw467.length); _i0_82++) {
            _nw467[_i0_82] = _init82(new BigNumber(_i0_82));
          }
          _2838_v60 = _nw467;
          let _2841_v61;
          _2841_v61 = _dafny.Seq.of(_2838_v60);
          _2831_v55 = (_2831_v55).update(new BigNumber((_2841_v61).length), _module.__default.safeModuloInt(new BigNumber((_2822_v46).length), (_2818_v42)[_module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length))]));
        } else {
          _2812_v36 = (_2812_v36).update(false, (_2823_v47).f17);
          let _2842_v62;
          _2842_v62 = _dafny.Seq.of(_2826_v50);
          _2842_v62 = _2842_v62;
          _2809_v33 = (_2809_v33).multipliedBy(_2809_v33);
          let _2843_v63;
          let _init83 = ((_2844_v47) => function (_2845_i4) {
            return (_2844_v47).f17;
          })(_2823_v47);
          let _nw468 = Array((new BigNumber(15)).toNumber());
          for (let _i0_83 = 0; _i0_83 < new BigNumber(_nw468.length); _i0_83++) {
            _nw468[_i0_83] = _init83(new BigNumber(_i0_83));
          }
          _2843_v63 = _nw468;
          let _index448 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_2843_v63).length));
          (_2843_v63)[_index448] = (_2820_v44).f15;
          let _index449 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_2843_v63).length));
          (_2843_v63)[_index449] = false;
          let _2846_v64;
          let _nw469 = new _module.C5();
          _nw469.__ctor(false, _2814_v38, (_2820_v44).f13, _2811_v35);
          _2846_v64 = _nw469;
          let _2847_v65;
          _2847_v65 = _dafny.Set.fromElements(_2846_v64);
          _2847_v65 = _2847_v65;
        }
        let _index450 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length));
        (_2818_v42)[_index450] = (_2809_v33).minus(new BigNumber(-417));
        let _index451 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_2818_v42).length));
        (_2818_v42)[_index451] = new BigNumber(320);
      }
      r2 = true;
      let _2848_v66;
      let _init84 = ((_2849_v0) => function (_2850_i5) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("djqnvnmv")).length), (_dafny.ZERO).minus(new BigNumber(-161)), new BigNumber(775), new BigNumber((_dafny.Seq.of(true, _2849_v0)).length));
      })(_2764_v0);
      let _nw470 = Array((new BigNumber(4)).toNumber());
      for (let _i0_84 = 0; _i0_84 < new BigNumber(_nw470.length); _i0_84++) {
        _nw470[_i0_84] = _init84(new BigNumber(_i0_84));
      }
      _2848_v66 = _nw470;
      let _2851_v67;
      _2851_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2764_v0,true);
      let _2852_v68;
      _2852_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2764_v0,_2851_v67);
      let _2853_v69;
      _2853_v69 = _dafny.Seq.of(_2764_v0, _module.__default.fm2(_2764_v0, globalState), _2764_v0, _2764_v0, _2764_v0);
      let _2854_v70;
      let _nw471 = new _module.C3();
      _nw471.__ctor(_2764_v0, _2848_v66, _2764_v0, _2852_v68, _2853_v69);
      _2854_v70 = _nw471;
      if (false) {
        let _2855_v71;
        _2855_v71 = new BigNumber(-129);
        let _2856_v72;
        _2856_v72 = _dafny.Set.fromElements((_2854_v70).f20, (_2854_v70).f20, (_2853_v69)[_module.__default.safeIndex(_2855_v71, new BigNumber((_2853_v69).length))], (_2854_v70).f20);
        let _2857_v73;
        _2857_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2856_v72).length),(_2854_v70).f20);
        _2857_v73 = (_2857_v73).update((_dafny.ZERO).minus(_2855_v71), (_2854_v70).f20);
        r2 = !((_2854_v70).f20);
        (globalState).f4 = _2764_v0;
        if ((_2854_v70).f20) {
          _2764_v0 = !((_2853_v69)[_module.__default.safeIndex(_2855_v71, new BigNumber((_2853_v69).length))]);
          _2855_v71 = new BigNumber(313);
          let _2858_v74;
          _2858_v74 = _module.D0.create_DC0((_2854_v70).f20);
          let _2859_v75;
          _2859_v75 = _dafny.Seq.UnicodeFromString("yv");
          r2 = ((_2858_v74).dtor_cf0) || (!(new BigNumber((_2859_v75).length)).isEqualTo(_2855_v71));
          let _2860_v76;
          _2860_v76 = _dafny.Map.Empty.slice().updateUnsafe(_2854_v70,_2857_v73);
          _2855_v71 = ((_2764_v0) ? (_module.__default.safeModuloInt(_2855_v71, new BigNumber((_dafny.MultiSet.fromElements((_2854_v70).f20, (((_2857_v73).contains(_2855_v71)) ? ((_2857_v73).get(_2855_v71)) : ((_2854_v70).f20)))).cardinality()))) : (_module.__default.safeDivisionInt(new BigNumber(((((_2860_v76).contains(_2854_v70)) ? ((_2860_v76).get(_2854_v70)) : (_2857_v73))).length), _2855_v71)));
          let _2861_v77;
          _2861_v77 = new _dafny.CodePoint('n'.codePointAt(0));
          let _2862_v78;
          _2862_v78 = _dafny.Map.Empty.slice().updateUnsafe(!((_2854_v70).f20),_2861_v77);
          let _2863_v79;
          _2863_v79 = _dafny.MultiSet.fromElements(new BigNumber((_2862_v78).length));
          let _2864_v80;
          _2864_v80 = _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((p1).length)), _2863_v79, _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_2855_v71), new BigNumber(-650)), _module.__default.fm28(_dafny.Seq.UnicodeFromString("n"), (_2854_v70).f20, _2855_v71, globalState));
          let _rhs397 = (((_2863_v79).contains(_2855_v71)) ? ((_2863_v79).get(_2855_v71)) : (_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_2859_v75).length)), _2855_v71)));
          let _rhs398 = _2855_v71;
          let _rhs399 = new _dafny.CodePoint('s'.codePointAt(0));
          let _rhs400 = _2855_v71;
          let _rhs401 = (_2864_v80)[_module.__default.safeIndex(_module.__default.safeModuloInt(_2855_v71, _2855_v71), new BigNumber((_2864_v80).length))];
          _2855_v71 = _rhs397;
          _2855_v71 = _rhs398;
          _2861_v77 = _rhs399;
          _2855_v71 = _rhs400;
          _2863_v79 = _rhs401;
        } else {
          let _2865_v81;
          _2865_v81 = _dafny.Seq.UnicodeFromString("yhmrbo");
          let _2866_v82;
          _2866_v82 = _dafny.Set.fromElements(_2865_v81);
          let _2867_v83;
          let _nw472 = new _module.C1();
          _nw472.__ctor((_2854_v70).f20, (_2854_v70).f20, _2852_v68, _2853_v69, _2848_v66);
          _2867_v83 = _nw472;
          let _2868_v84;
          _2868_v84 = _module.D3.create_DC10((_2854_v70).f20, new BigNumber((_2857_v73).length), _2867_v83);
          let _2869_v85;
          _2869_v85 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_2855_v71);
          let _2870_v86;
          _2870_v86 = _dafny.Map.Empty.slice().updateUnsafe((_2854_v70).f20,_2855_v71);
          let _2871_v87;
          _2871_v87 = _dafny.MultiSet.fromElements((_2854_v70).f20);
          let _2872_v88;
          _2872_v88 = _dafny.Set.fromElements(new BigNumber(858), (((_2871_v87).contains(_2764_v0)) ? ((_2871_v87).get(_2764_v0)) : (_2855_v71)));
          let _2873_v89;
          _2873_v89 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2857_v73).length),_2855_v71);
          let _2874_v90;
          let _nw473 = Array((new BigNumber(22)).toNumber());
          _nw473[(_dafny.ZERO).toNumber()] = new BigNumber((_2865_v81).length);
          _nw473[(_dafny.ONE).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(2)).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(3)).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2853_v69, _module.__default.fm18(_2764_v0, _2855_v71, false, _2855_v71, globalState))).length);
          _nw473[(new BigNumber(5)).toNumber()] = new BigNumber((_2866_v82).length);
          _nw473[(new BigNumber(6)).toNumber()] = (_2868_v84).dtor_cf15;
          _nw473[(new BigNumber(7)).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(8)).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(9)).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(10)).toNumber()] = new BigNumber(318);
          _nw473[(new BigNumber(11)).toNumber()] = new BigNumber(263);
          _nw473[(new BigNumber(12)).toNumber()] = ((_2764_v0) ? (_2855_v71) : (_2855_v71));
          _nw473[(new BigNumber(13)).toNumber()] = (((_2869_v85).contains(p0)) ? ((_2869_v85).get(p0)) : (_2855_v71));
          _nw473[(new BigNumber(14)).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2855_v71,_2855_v71)).length);
          _nw473[(new BigNumber(16)).toNumber()] = _2855_v71;
          _nw473[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_2855_v71);
          _nw473[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2865_v81, _dafny.Seq.UnicodeFromString("sa"))).length);
          _nw473[(new BigNumber(19)).toNumber()] = (((_2870_v86).contains((_2854_v70).f20)) ? ((_2870_v86).get((_2854_v70).f20)) : (new BigNumber((_2872_v88).length)));
          _nw473[(new BigNumber(20)).toNumber()] = new BigNumber((_2873_v89).length);
          _nw473[(new BigNumber(21)).toNumber()] = (_2855_v71).plus(_2855_v71);
          _2874_v90 = _nw473;
          let _2875_v91;
          let _nw474 = new _module.C6();
          _nw474.__ctor(p1, _2764_v0, _2852_v68, _dafny.Seq.of((_2854_v70).f20), _2848_v66);
          _2875_v91 = _nw474;
          let _2876_v92;
          _2876_v92 = _dafny.Set.fromElements(_2875_v91, _2875_v91, _2875_v91);
          let _index452 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length));
          (_2874_v90)[_index452] = _module.__default.safeModuloInt(new BigNumber((_2876_v92).length), _2855_v71);
          let _2877_v93;
          let _nw475 = new _module.C5();
          _nw475.__ctor((((_2857_v73).contains((_dafny.ZERO).minus(_2855_v71))) ? ((_2857_v73).get((_dafny.ZERO).minus(_2855_v71))) : ((_2854_v70).f20)), _2848_v66, _2852_v68, (_2867_v83).f14);
          _2877_v93 = _nw475;
          let _index453 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length));
          let _rhs402 = (_2855_v71).isEqualTo(_2855_v71);
          let _rhs403 = (_dafny.ZERO).minus(_2855_v71);
          let _rhs404 = _2877_v93;
          let _rhs405 = _module.__default.fm66(globalState);
          let _rhs406 = false;
          let _lhs244 = globalState;
          let _lhs245 = _2874_v90;
          let _lhs246 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length));
          _lhs244.f4 = _rhs402;
          _lhs245[_lhs246] = _rhs403;
          _2877_v93 = _rhs404;
          _2873_v89 = _rhs405;
          r2 = _rhs406;
          let _index454 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length));
          (_2874_v90)[_index454] = _2855_v71;
          let _2878_v94;
          _2878_v94 = _module.D3.create_DC11(_module.D3.create_DC9((_2874_v90)[_module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length))], new BigNumber(18)));
          let _2879_v95;
          _2879_v95 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm13(_2855_v71, globalState),_2878_v94);
          let _2880_v96;
          _2880_v96 = _module.D8.create_DC25(_2764_v0, _2855_v71, _2855_v71, _dafny.Map.Empty.slice().updateUnsafe(!(_2764_v0),(_2874_v90)[_module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length))]));
          let _2881_v97;
          _2881_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2865_v81).length),(_2880_v96).dtor_cf40);
          let _2882_v98;
          _2882_v98 = _module.D3.create_DC9(_module.__default.fm0((_2854_v70).f20, _2881_v97, (_2854_v70).f20, globalState), new BigNumber((_2865_v81).length));
          _2879_v95 = (_2879_v95).update(_2856_v72, _module.D3.create_DC11(_2882_v98));
          _2870_v86 = _2870_v86;
          let _index455 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length));
          (_2874_v90)[_index455] = (_2855_v71).minus((_2874_v90)[_module.__default.safeIndex(new BigNumber(267), new BigNumber((_2874_v90).length))]);
        }
        let _2883_v100;
        _2883_v100 = _dafny.MultiSet.fromElements(function () {
          let _coll69 = new _dafny.Map();
          for (const _compr_69 of _dafny.IntegerRange(new BigNumber(430), new BigNumber(-595))) {
            let _2884_v99 = _compr_69;
            if (((new BigNumber(430)).isLessThanOrEqualTo(_2884_v99)) && ((_2884_v99).isLessThan(new BigNumber(-595)))) {
              _coll69.push([_module.__default.safeDivisionInt(_2884_v99, _2855_v71),_2855_v71]);
            }
          }
          return _coll69;
        }());
        let _2885_v102;
        _2885_v102 = _dafny.Seq.UnicodeFromString("bcsn");
        let _2886_v103;
        _2886_v103 = _dafny.Map.Empty.slice().updateUnsafe(((_2764_v0) ? (_2855_v71) : (_2855_v71)),(_2883_v100).update(function () {
          let _coll70 = new _dafny.Map();
          for (const _compr_70 of _dafny.IntegerRange(new BigNumber(430), new BigNumber(674))) {
            let _2887_v101 = _compr_70;
            if (((new BigNumber(430)).isLessThanOrEqualTo(_2887_v101)) && ((_2887_v101).isLessThan(new BigNumber(674)))) {
              _coll70.push([_module.__default.safeDivisionInt(_2887_v101, (_dafny.ZERO).minus(_2855_v71)),new BigNumber(887)]);
            }
          }
          return _coll70;
        }(), _module.__default.abs(new BigNumber((_2885_v102).length))));
        let _2888_v104;
        _2888_v104 = _dafny.Map.Empty.slice().updateUnsafe(_2855_v71,_2855_v71);
        _2886_v103 = (_2886_v103).update((_dafny.ZERO).minus(_2855_v71), _dafny.MultiSet.fromElements(_2888_v104));
      } else {
        let _2889_v105;
        _2889_v105 = new BigNumber(983);
        let _2890_v106;
        _2890_v106 = _dafny.Map.Empty.slice().updateUnsafe(_2764_v0,_2889_v105);
        let _2891_v107;
        _2891_v107 = _module.D8.create_DC24(true);
        let _2892_v108;
        let _nw476 = Array((new BigNumber(11)).toNumber());
        _nw476[(_dafny.ZERO).toNumber()] = _module.D8.create_DC24((_2854_v70).f20);
        _nw476[(_dafny.ONE).toNumber()] = _module.__default.fm67(_2889_v105, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2890_v106).length))), _2889_v105, globalState);
        _nw476[(new BigNumber(2)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(3)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(4)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(5)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(6)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(7)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(8)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(9)).toNumber()] = _2891_v107;
        _nw476[(new BigNumber(10)).toNumber()] = _2891_v107;
        _2892_v108 = _nw476;
        let _index456 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2892_v108).length));
        (_2892_v108)[_index456] = _2891_v107;
        let _index457 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2892_v108).length));
        (_2892_v108)[_index457] = _module.__default.fm67(_2889_v105, (_dafny.ZERO).minus(_2889_v105), _2889_v105, globalState);
        let _2893_v109;
        _2893_v109 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2889_v105),_dafny.Map.Empty.slice().updateUnsafe(_2764_v0,_2889_v105));
        _2889_v105 = _module.__default.safeModuloInt(_module.__default.fm0((_2854_v70).f20, _2893_v109, (_2854_v70).f20, globalState), _2889_v105);
        if (_module.__default.fm2(_2764_v0, globalState)) {
          let _2894_v110;
          _2894_v110 = _dafny.Set.fromElements(true);
          let _2895_v111;
          _2895_v111 = _dafny.MultiSet.fromElements((p1)[_module.__default.safeIndex(_2889_v105, new BigNumber((p1).length))]);
          let _rhs407 = (_dafny.MultiSet.FromArray(p1)).IsDisjointFrom(_2895_v111);
          let _rhs408 = _2894_v110;
          let _lhs247 = globalState;
          _lhs247.f4 = _rhs407;
          _2894_v110 = _rhs408;
          _2889_v105 = _2889_v105;
          let _2896_v112;
          let _nw477 = Array((new BigNumber(21)).toNumber());
          _nw477[(_dafny.ZERO).toNumber()] = (_2854_v70).f20;
          _nw477[(_dafny.ONE).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(2)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(3)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(4)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(5)).toNumber()] = _2764_v0;
          _nw477[(new BigNumber(6)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(7)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(8)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(9)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(10)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(11)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(12)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(13)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(14)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(15)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(16)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(17)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(18)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(19)).toNumber()] = (_2854_v70).f20;
          _nw477[(new BigNumber(20)).toNumber()] = (_2854_v70).f20;
          _2896_v112 = _nw477;
          let _2897_v113;
          let _nw478 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _2897_v113 = _nw478;
          let _2898_v114;
          _2898_v114 = _dafny.Map.Empty.slice().updateUnsafe(_2896_v112,_2897_v113);
          let _2899_v115;
          let _nw479 = new _module.C5();
          _nw479.__ctor((_2854_v70).f20, _2848_v66, _2852_v68, _2853_v69);
          _2899_v115 = _nw479;
          let _2900_v116;
          _2900_v116 = _dafny.Map.Empty.slice().updateUnsafe((((_2898_v114).contains(_2896_v112)) ? ((_2898_v114).get(_2896_v112)) : (_2897_v113)),_2899_v115);
          let _rhs409 = _2900_v116;
          let _rhs410 = _module.__default.fm2((_2854_v70).f20, globalState);
          _2900_v116 = _rhs409;
          r2 = _rhs410;
          let _index458 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_2896_v112).length));
          (_2896_v112)[_index458] = (_2854_v70).f20;
          let _2901_v117;
          let _nw480 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2901_v117 = _nw480;
          let _index459 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2901_v117).length));
          (_2901_v117)[_index459] = p0;
          let _2902_v118;
          _2902_v118 = _dafny.MultiSet.fromElements((_2854_v70).f20);
          let _2903_v119;
          _2903_v119 = _dafny.Seq.UnicodeFromString("rnbue");
          let _2904_v120;
          _2904_v120 = _module.D14.create_DC41(_2903_v119, _2889_v105);
          let _index460 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_2896_v112).length));
          let _index461 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2901_v117).length));
          let _rhs411 = (_2889_v105).isLessThan(((((_2902_v118).contains((_2854_v70).f20)) ? ((_2902_v118).get((_2854_v70).f20)) : (_2889_v105))).multipliedBy((_2904_v120).dtor_cf65));
          let _rhs412 = p0;
          let _lhs248 = _2896_v112;
          let _lhs249 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_2896_v112).length));
          let _lhs250 = _2901_v117;
          let _lhs251 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2901_v117).length));
          _lhs248[_lhs249] = _rhs411;
          _lhs250[_lhs251] = _rhs412;
          let _2905_v121;
          let _nw481 = new _module.C0();
          _nw481.__ctor();
          _2905_v121 = _nw481;
          let _2906_v122;
          _2906_v122 = _2905_v121;
          let _2907_v123;
          _2907_v123 = _dafny.Map.Empty.slice().updateUnsafe(_2906_v122,_dafny.Seq.update(p1, _module.__default.safeIndex(_2889_v105, new BigNumber((p1).length)), _2889_v105));
          let _2908_v124;
          _2908_v124 = _dafny.Seq.of(p1, p1);
          let _2909_v125;
          _2909_v125 = _module.D5.create_DC15(p1);
          let _2910_v126;
          _2910_v126 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_2764_v0, globalState),p1);
          let _2911_v127;
          let _nw482 = Array((new BigNumber(26)).toNumber());
          _nw482[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_2889_v105);
          _nw482[(_dafny.ONE).toNumber()] = p1;
          _nw482[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements((_2854_v70).f20)).cardinality()), new BigNumber((p1).length)), _2889_v105);
          _nw482[(new BigNumber(3)).toNumber()] = (((_2899_v115).f19) ? (_dafny.Seq.of(_2889_v105, _2889_v105)) : (p1));
          _nw482[(new BigNumber(4)).toNumber()] = p1;
          _nw482[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_2889_v105);
          _nw482[(new BigNumber(6)).toNumber()] = p1;
          _nw482[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(p1, _module.__default.safeIndex(_2889_v105, new BigNumber((p1).length)), _2889_v105), _module.__default.fm1((_2854_v70).f20, (_2854_v70).f20, globalState));
          _nw482[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(p1, _module.__default.safeIndex(_2889_v105, new BigNumber((p1).length)), _2889_v105);
          _nw482[(new BigNumber(9)).toNumber()] = (((_2907_v123).contains(_2906_v122)) ? ((_2907_v123).get(_2906_v122)) : ((_2908_v124)[_module.__default.safeIndex(_2889_v105, new BigNumber((_2908_v124).length))]));
          _nw482[(new BigNumber(10)).toNumber()] = p1;
          _nw482[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(p1, p1);
          _nw482[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), function (_2912_i6) {
            return new BigNumber(875);
          });
          _nw482[(new BigNumber(13)).toNumber()] = p1;
          _nw482[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_2889_v105, _2889_v105, (((_2902_v118).contains((_2899_v115).f19)) ? ((_2902_v118).get((_2899_v115).f19)) : (_2889_v105)), _2889_v105, _2889_v105);
          _nw482[(new BigNumber(15)).toNumber()] = p1;
          _nw482[(new BigNumber(16)).toNumber()] = p1;
          _nw482[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), ((_2913_v105) => function (_2914_i7) {
            return _2913_v105;
          })(_2889_v105));
          _nw482[(new BigNumber(18)).toNumber()] = _module.__default.fm1((_2896_v112)[_module.__default.safeIndex(new BigNumber(274), new BigNumber((_2896_v112).length))], !((_2854_v70).f20), globalState);
          _nw482[(new BigNumber(19)).toNumber()] = p1;
          _nw482[(new BigNumber(20)).toNumber()] = p1;
          _nw482[(new BigNumber(21)).toNumber()] = p1;
          _nw482[(new BigNumber(22)).toNumber()] = (_2909_v125).dtor_cf25;
          _nw482[(new BigNumber(23)).toNumber()] = p1;
          _nw482[(new BigNumber(24)).toNumber()] = (_2899_v115).fm6(_dafny.Seq.UnicodeFromString("morqadjvx"), globalState);
          _nw482[(new BigNumber(25)).toNumber()] = (((_2910_v126).contains(_module.__default.fm2((_2854_v70).f20, globalState))) ? ((_2910_v126).get(_module.__default.fm2((_2854_v70).f20, globalState))) : (_dafny.Seq.of(new BigNumber((_2890_v106).length))));
          _2911_v127 = _nw482;
          let _index462 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2911_v127).length));
          (_2911_v127)[_index462] = _dafny.Seq.Concat(p1, p1);
          let _2915_v128;
          _2915_v128 = _dafny.Map.Empty.slice().updateUnsafe(_2889_v105,new BigNumber((_2894_v110).length));
          let _index463 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2911_v127).length));
          (_2911_v127)[_index463] = _dafny.Seq.Concat(_dafny.Seq.of(_2889_v105, new BigNumber((_dafny.Seq.update(_2903_v119, _module.__default.safeIndex(_2889_v105, new BigNumber((_2903_v119).length)), p0)).length), new BigNumber((_2915_v128).length)), _dafny.Seq.Concat(p1, p1));
        } else {
          let _2916_v129;
          _2916_v129 = _module.D0.create_DC2(new BigNumber((p1).length), (_2854_v70).f20);
          let _2917_v130;
          let _nw483 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
          _2917_v130 = _nw483;
          let _2918_v131;
          _2918_v131 = _2917_v130;
          let _2919_v132;
          let _2920_v133;
          let _2921_v134;
          let _2922_v135;
          let _out135;
          let _out136;
          let _out137;
          let _out138;
          let _outcollector35 = _module.__default.m0((_2854_v70).f20, _2916_v129, (_2889_v105).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2918_v131,(_2854_v70).f20)).length)), globalState);
          _out135 = _outcollector35[0];
          _out136 = _outcollector35[1];
          _out137 = _outcollector35[2];
          _out138 = _outcollector35[3];
          _2919_v132 = _out135;
          _2920_v133 = _out136;
          _2921_v134 = _out137;
          _2922_v135 = _out138;
          let _2923_v136;
          _2923_v136 = _dafny.Seq.UnicodeFromString("nn");
          let _2924_v137;
          _2924_v137 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_2923_v136, _2923_v136),_dafny.Seq.of(!(!((_2854_v70).f20))));
          _2924_v137 = (_2924_v137).update((_2854_v70).f20, _2853_v69);
          let _2925_v138;
          let _nw484 = Array((new BigNumber(14)).toNumber());
          _nw484[(_dafny.ZERO).toNumber()] = p0;
          _nw484[(_dafny.ONE).toNumber()] = p0;
          _nw484[(new BigNumber(2)).toNumber()] = p0;
          _nw484[(new BigNumber(3)).toNumber()] = p0;
          _nw484[(new BigNumber(4)).toNumber()] = p0;
          _nw484[(new BigNumber(5)).toNumber()] = p0;
          _nw484[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
          _nw484[(new BigNumber(7)).toNumber()] = p0;
          _nw484[(new BigNumber(8)).toNumber()] = p0;
          _nw484[(new BigNumber(9)).toNumber()] = p0;
          _nw484[(new BigNumber(10)).toNumber()] = p0;
          _nw484[(new BigNumber(11)).toNumber()] = p0;
          _nw484[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
          _nw484[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _2925_v138 = _nw484;
          let _index464 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2925_v138).length));
          (_2925_v138)[_index464] = p0;
          let _index465 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_2925_v138).length));
          (_2925_v138)[_index465] = p0;
          r2 = _2764_v0;
          let _2926_v139;
          let _nw485 = new _module.C1();
          _nw485.__ctor((_2854_v70).f20, (_2854_v70).f20, _2852_v68, _2853_v69, _2848_v66);
          _2926_v139 = _nw485;
          let _nw486 = new _module.C5();
          _nw486.__ctor((_2921_v134).isLessThanOrEqualTo(_2921_v134), _2926_v139.f12, (_2926_v139).f13, _dafny.Seq.Concat(_2853_v69, _module.__default.fm15((_2854_v70).f20, (p1)[_module.__default.safeIndex(_2921_v134, new BigNumber((p1).length))], _2889_v105, _2889_v105, globalState)));
          _2926_v139 = _nw486;
        }
        let _2927_v140;
        let _nw487 = new _module.C6();
        _nw487.__ctor(_dafny.Seq.Concat(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(523)), function (_2928_i8) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("fvbnjdfav")).length);
        })), false, (_2852_v68).Merge(_2852_v68), _2853_v69, _2848_v66);
        _2927_v140 = _nw487;
        if (!(_2764_v0)) {
          let _2929_v141;
          _2929_v141 = _module.D3.create_DC9(_2889_v105, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(708)), function (_2930_i9) {
  return new _dafny.CodePoint('g'.codePointAt(0));
})).length));
          let _2931_v142;
          let _nw488 = new _module.C0();
          _nw488.__ctor();
          _2931_v142 = _nw488;
          let _2932_v143;
          _2932_v143 = _dafny.Map.Empty.slice().updateUnsafe(_2931_v142,true);
          let _2933_v144;
          _2933_v144 = _module.D0.create_DC2((_2929_v141).dtor_cf12, (((_2932_v143).contains(_2931_v142)) ? ((_2932_v143).get(_2931_v142)) : (_2764_v0)));
          let _2934_v145;
          let _2935_v146;
          let _2936_v147;
          let _2937_v148;
          let _out139;
          let _out140;
          let _out141;
          let _out142;
          let _outcollector36 = _module.__default.m0(true, _2933_v144, _2764_v0, globalState);
          _out139 = _outcollector36[0];
          _out140 = _outcollector36[1];
          _out141 = _outcollector36[2];
          _out142 = _outcollector36[3];
          _2934_v145 = _out139;
          _2935_v146 = _out140;
          _2936_v147 = _out141;
          _2937_v148 = _out142;
          let _2938_v149;
          _2938_v149 = _dafny.Seq.of(_2890_v106);
          let _2939_v150;
          let _2940_v151;
          let _2941_v152;
          let _2942_v153;
          let _out143;
          let _out144;
          let _out145;
          let _out146;
          let _outcollector37 = _module.__default.m0(_2937_v148, _module.__default.fm3(_2938_v149, globalState), _2937_v148, globalState);
          _out143 = _outcollector37[0];
          _out144 = _outcollector37[1];
          _out145 = _outcollector37[2];
          _out146 = _outcollector37[3];
          _2939_v150 = _out143;
          _2940_v151 = _out144;
          _2941_v152 = _out145;
          _2942_v153 = _out146;
          let _2943_v154;
          _2943_v154 = _2848_v66;
          let _2944_v155;
          let _nw489 = new _module.C2();
          _nw489.__ctor(p1, false, _dafny.Map.Empty.slice().updateUnsafe(_2935_v146,_2851_v67), _2853_v69, (((_2854_v70).f20) ? (_2848_v66) : ((_2943_v154))));
          _2944_v155 = _nw489;
          _2944_v155 = _2944_v155;
          _2937_v148 = _2937_v148;
          let _2945_v156;
          let _init85 = ((_2946_v148) => function (_2947_i10) {
            return _2946_v148;
          })(_2937_v148);
          let _nw490 = Array((new BigNumber(5)).toNumber());
          for (let _i0_85 = 0; _i0_85 < new BigNumber(_nw490.length); _i0_85++) {
            _nw490[_i0_85] = _init85(new BigNumber(_i0_85));
          }
          _2945_v156 = _nw490;
          let _2948_v157;
          _2948_v157 = _dafny.MultiSet.fromElements(_2942_v153, _2940_v151);
          let _index466 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_2945_v156).length));
          (_2945_v156)[_index466] = (_2948_v157).IsSubsetOf(_2948_v157);
          let _index467 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_2945_v156).length));
          (_2945_v156)[_index467] = _2764_v0;
        } else {
          _2889_v105 = _2889_v105;
          _2889_v105 = (_dafny.ZERO).minus(_2889_v105);
          _2889_v105 = new BigNumber((_dafny.Seq.UnicodeFromString("ak")).length);
          _2853_v69 = _2853_v69;
          let _2949_v158;
          let _nw491 = Array((new BigNumber(10)).toNumber()).fill([]);
          _2949_v158 = _nw491;
          let _2950_v159;
          let _nw492 = Array((new BigNumber(29)).toNumber()).fill(false);
          _2950_v159 = _nw492;
          let _index468 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_2949_v158).length));
          (_2949_v158)[_index468] = _2950_v159;
          let _2951_v160;
          _2951_v160 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2854_v70).f20);
          let _2952_v161;
          _2952_v161 = _dafny.Map.Empty.slice().updateUnsafe(_2889_v105,_2764_v0);
          let _2953_v162;
          _2953_v162 = _module.D7.create_DC21(_2952_v161, _2764_v0, (((_2851_v67).contains(_2764_v0)) ? ((_2851_v67).get(_2764_v0)) : (true)));
          let _index469 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_2949_v158).length));
          let _nw493 = Array((new BigNumber(7)).toNumber());
          _nw493[(_dafny.ZERO).toNumber()] = (((_2854_v70).f20) ? ((_2854_v70).f20) : ((_2854_v70).f20));
          _nw493[(_dafny.ONE).toNumber()] = (_2889_v105).isEqualTo(_2889_v105);
          _nw493[(new BigNumber(2)).toNumber()] = (((_2951_v160).contains(new _dafny.CodePoint('o'.codePointAt(0)))) ? ((_2951_v160).get(new _dafny.CodePoint('o'.codePointAt(0)))) : ((_2854_v70).f20));
          _nw493[(new BigNumber(3)).toNumber()] = (_2854_v70).f20;
          _nw493[(new BigNumber(4)).toNumber()] = (_2953_v162).dtor_cf32;
          _nw493[(new BigNumber(5)).toNumber()] = !((((_2952_v161).contains((_2927_v140).fm7(_2889_v105, globalState))) ? ((_2952_v161).get((_2927_v140).fm7(_2889_v105, globalState))) : ((_2854_v70).f20)));
          _nw493[(new BigNumber(6)).toNumber()] = (_2854_v70).f20;
          (_2949_v158)[_index469] = _nw493;
        }
      }
      let _2954_v163;
      _2954_v163 = new BigNumber(519);
      let _2955_v164;
      _2955_v164 = _dafny.MultiSet.fromElements(_2954_v163);
      let _hi16 = (((_2955_v164).contains(_2954_v163)) ? ((_2955_v164).get(_2954_v163)) : (_2954_v163));
      for (let _2956_i11 = _2954_v163; _2956_i11.isLessThan(_hi16); _2956_i11 = _2956_i11.plus(_dafny.ONE)) {
        let _2957_v165;
        let _nw494 = new _module.C1();
        _nw494.__ctor((_2854_v70).f20, true, _2852_v68, _dafny.Seq.of((_2854_v70).f20, (_2854_v70).f20), _2848_v66);
        _2957_v165 = _nw494;
        let _2958_v166;
        let _nw495 = Array((new BigNumber(11)).toNumber()).fill(false);
        _2958_v166 = _nw495;
        let _2959_v167;
        _2959_v167 = _dafny.Map.Empty.slice().updateUnsafe(_2954_v163,true);
        let _2960_v168;
        _2960_v168 = _dafny.Map.Empty.slice().updateUnsafe(_2764_v0,new BigNumber((_2959_v167).length));
        let _2961_v169;
        _2961_v169 = _dafny.Set.fromElements(_2764_v0);
        let _2962_v170;
        _2962_v170 = _dafny.Map.Empty.slice().updateUnsafe(((_2764_v0) ? (new BigNumber((_2960_v168).length)) : (new BigNumber((_2961_v169).length))),_2958_v166);
        _2958_v166 = (((_2962_v170).contains((_2954_v163).multipliedBy(_2956_i11))) ? ((_2962_v170).get((_2954_v163).multipliedBy(_2956_i11))) : (_2958_v166));
        _2851_v67 = (_2851_v67).update(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_2957_v165.f18), _2853_v69), !(_2764_v0));
        let _2963_v171;
        _2963_v171 = _dafny.Map.Empty.slice().updateUnsafe((_2957_v165).f17,_2958_v166);
        _2963_v171 = (_2963_v171).update(!(!(_2764_v0)), _2958_v166);
      }
      if (_2764_v0) {
        let _2964_v172;
        _2964_v172 = _dafny.Seq.UnicodeFromString("sniuh");
        let _2965_v173;
        let _2966_v174;
        let _2967_v175;
        let _2968_v176;
        let _out147;
        let _out148;
        let _out149;
        let _out150;
        let _outcollector38 = (_2854_v70).m9((true) || (_2764_v0), false, _dafny.Seq.Concat(_2964_v172, _dafny.Seq.UnicodeFromString("pmgiy")), globalState);
        _out147 = _outcollector38[0];
        _out148 = _outcollector38[1];
        _out149 = _outcollector38[2];
        _out150 = _outcollector38[3];
        _2965_v173 = _out147;
        _2966_v174 = _out148;
        _2967_v175 = _out149;
        _2968_v176 = _out150;
        let _2969_v177;
        _2969_v177 = _dafny.Map.Empty.slice().updateUnsafe(_2966_v174,new BigNumber(474));
        _2969_v177 = (_2969_v177).update(new BigNumber((_dafny.Set.fromElements(_module.__default.fm2(_2764_v0, globalState))).length), _2954_v163);
        let _2970_v178;
        let _nw496 = new _module.C4();
        _nw496.__ctor();
        _2970_v178 = _nw496;
        (globalState).f4 = (_2966_v174).isLessThanOrEqualTo(new BigNumber((_2853_v69).length));
        let _2971_v179;
        _2971_v179 = _module.D11.create_DC32(_module.__default.safeModuloInt(_2965_v173, _2965_v173), _module.__default.safeModuloInt(_2954_v163, new BigNumber(922)));
        _2971_v179 = _module.D11.create_DC32(_2965_v173, _2966_v174);
      } else {
        _2954_v163 = _2954_v163;
        let _2972_v180;
        _2972_v180 = _dafny.Map.Empty.slice().updateUnsafe(_2954_v163,_2954_v163);
        let _2973_v181;
        _2973_v181 = _dafny.Set.fromElements(_2954_v163, (((_2955_v164).contains(new BigNumber((_2853_v69).length))) ? ((_2955_v164).get(new BigNumber((_2853_v69).length))) : (_2954_v163)), _2954_v163);
        let _2974_v182;
        _2974_v182 = _dafny.Seq.UnicodeFromString("jttuiim");
        let _2975_v183;
        _2975_v183 = _dafny.Map.Empty.slice().updateUnsafe(_2764_v0,_2974_v182);
        _2972_v180 = (_2972_v180).update(((_2764_v0) ? (_2954_v163) : (_2954_v163)), (new BigNumber((_2973_v181).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber(((((_2975_v183).contains((_2854_v70).f20)) ? ((_2975_v183).get((_2854_v70).f20)) : (_2974_v182))).length))));
        let _2976_v184;
        let _nw497 = new _module.C2();
        _nw497.__ctor(_dafny.Seq.of(new BigNumber(-849)), (_2854_v70).f20, _2852_v68, _2853_v69, _2848_v66);
        _2976_v184 = _nw497;
        let _2977_v185;
        _2977_v185 = _dafny.Set.fromElements(_2976_v184);
        let _2978_v186;
        _2978_v186 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-741), _2954_v163), p1, _dafny.Seq.of(new BigNumber((_2974_v182).length)), p1, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), ((_2979_v70) => function (_2980_i12) {
          return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_2979_v70).f20))).cardinality());
        })(_2854_v70)), _module.__default.safeIndex(_2954_v163, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), ((_2981_v70) => function (_2982_i12) {
          return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_2981_v70).f20))).cardinality());
        })(_2854_v70))).length)), new BigNumber((p1).length)));
        let _rhs413 = _module.__default.safeDivisionInt(_2954_v163, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_2954_v163, new BigNumber((_2978_v186).length))));
        let _rhs414 = _2954_v163;
        let _rhs415 = !(!((_2854_v70).f20));
        let _rhs416 = (_2977_v185).Intersect(_2977_v185);
        _2954_v163 = _rhs413;
        _2954_v163 = _rhs414;
        _2764_v0 = _rhs415;
        _2977_v185 = _rhs416;
        let _2983_v187;
        let _out151;
        _out151 = (_2976_v184).m4(_2954_v163, (_2954_v163).multipliedBy((p1)[_module.__default.safeIndex(_2954_v163, new BigNumber((p1).length))]), _2974_v182, _module.__default.fm56(p0, globalState), globalState);
        _2983_v187 = _out151;
        _2954_v163 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_2983_v187));
      }
      let _2984_v188;
      _2984_v188 = _module.D0.create_DC1();
      r0 = _2984_v188;
      r1 = (_2851_v67).Merge(_2851_v67);
      r2 = _2764_v0;
      return [r0, r1, r2];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _module.D0.Default();
      let _2985_v0;
      let _init86 = function (_2986_i1) {
        return false;
      };
      let _nw498 = Array((new BigNumber(12)).toNumber());
      for (let _i0_86 = 0; _i0_86 < new BigNumber(_nw498.length); _i0_86++) {
        _nw498[_i0_86] = _init86(new BigNumber(_i0_86));
      }
      _2985_v0 = _nw498;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2985_v0).length))) {
        let _2987_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2987_i0)) && ((_2987_i0).isLessThan(new BigNumber((_2985_v0).length))))) {
          (_2985_v0)[(_2987_i0)] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("cbrqflr"), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(630)), function (_2988_i2) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("m")));
        }
      }
      let _2989_v1;
      _2989_v1 = false;
      let _2990_v2;
      _2990_v2 = _dafny.Set.fromElements(((_2989_v1) ? (_2989_v1) : (_2989_v1)), true, _2989_v1);
      _2990_v2 = (_2990_v2).Intersect(_2990_v2);
      let _2991_v3;
      let _nw499 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _2991_v3 = _nw499;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2991_v3).length))) {
        let _2992_i3 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2992_i3)) && ((_2992_i3).isLessThan(new BigNumber((_2991_v3).length))))) {
          (_2991_v3)[(_2992_i3)] = (_2992_i3).plus(new BigNumber((((true) ? (_dafny.Seq.UnicodeFromString("ufueuh")) : (_dafny.Seq.UnicodeFromString("lbaoaupfj")))).length));
        }
      }
      let _2993_v4;
      _2993_v4 = _dafny.Seq.UnicodeFromString("nubb");
      let _hi17 = new BigNumber(-637);
      for (let _2994_i4 = new BigNumber((((_2989_v1) ? (_dafny.Seq.UnicodeFromString("ktxfhm")) : (_2993_v4))).length); _2994_i4.isLessThan(_hi17); _2994_i4 = _2994_i4.plus(_dafny.ONE)) {
        let _rhs417 = p0;
        let _rhs418 = _module.__default.fm2(_2989_v1, globalState);
        let _rhs419 = p0;
        let _rhs420 = _2989_v1;
        let _lhs252 = globalState;
        r1 = _rhs417;
        _2989_v1 = _rhs418;
        r0 = _rhs419;
        _lhs252.f4 = _rhs420;
        let _index470 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_2991_v3).length));
        (_2991_v3)[_index470] = _2994_i4;
        let _index471 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_2991_v3).length));
        (_2991_v3)[_index471] = p0;
        r0 = _2994_i4;
        let _2995_v5;
        let _nw500 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _2995_v5 = _nw500;
        _2995_v5 = _2995_v5;
      }
      let _2996_i5;
      _2996_i5 = _dafny.ZERO;
      L16: {
        while ((_2989_v1) || (_2989_v1)) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2996_i5)) {
              break L16;
            }
            _2996_i5 = (_2996_i5).plus(_dafny.ONE);
            r1 = new BigNumber(-328);
            if (_2989_v1) {
              let _2997_v6;
              _2997_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_2989_v1, globalState),p0);
              _2997_v6 = (_2997_v6).update(_2989_v1, p0);
              r1 = p0;
              let _2998_v7;
              _2998_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2989_v1,((_2989_v1) ? (_2989_v1) : (_2989_v1)));
              let _2999_v8;
              _2999_v8 = _dafny.Seq.of(p0);
              _2998_v7 = (_2998_v7).update(_dafny.Seq.contains(_2999_v8, p0), _2989_v1);
              let _3000_v9;
              _3000_v9 = _dafny.MultiSet.fromElements(_2989_v1, !(_2989_v1));
              _2989_v1 = (_3000_v9).IsProperSubsetOf(_dafny.MultiSet.fromElements(_2989_v1));
              let _3001_v10;
              let _nw501 = Array((new BigNumber(21)).toNumber()).fill(_dafny.MultiSet.Empty);
              _3001_v10 = _nw501;
              let _3002_v11;
              _3002_v11 = _dafny.MultiSet.fromElements(_2989_v1, _2989_v1, !(_2989_v1));
              let _index472 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_3001_v10).length));
              (_3001_v10)[_index472] = _3002_v11;
              let _3003_v12;
              _3003_v12 = _dafny.Seq.of(_2991_v3);
              let _index473 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_3001_v10).length));
              let _rhs421 = !(_dafny.Seq.IsProperPrefixOf(_3003_v12, _3003_v12));
              let _rhs422 = _3002_v11;
              let _rhs423 = _2993_v4;
              let _lhs253 = globalState;
              let _lhs254 = _3001_v10;
              let _lhs255 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_3001_v10).length));
              _lhs253.f4 = _rhs421;
              _lhs254[_lhs255] = _rhs422;
              _2993_v4 = _rhs423;
            } else {
              let _3004_v13;
              _3004_v13 = _dafny.Map.Empty.slice().updateUnsafe(_2989_v1,p0);
              let _3005_v14;
              _3005_v14 = _dafny.Map.Empty.slice().updateUnsafe(_2989_v1,_3004_v13);
              let _3006_v15;
              _3006_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_3005_v14).contains(_2989_v1)) ? ((_3005_v14).get(_2989_v1)) : (_3004_v13)));
              let _3007_v16;
              _3007_v16 = _dafny.Seq.of(_module.__default.fm0(_2989_v1, _3006_v15, _2989_v1, globalState));
              let _index474 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2991_v3).length));
              (_2991_v3)[_index474] = (((_3004_v13).contains(false)) ? ((_3004_v13).get(false)) : (new BigNumber((_3007_v16).length)));
              let _3008_v17;
              _3008_v17 = _dafny.Seq.of(true);
              let _index475 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_2991_v3).length));
              (_2991_v3)[_index475] = (p0).plus(new BigNumber((_dafny.Seq.Concat(_3008_v17, _3008_v17)).length));
              (globalState).f4 = _module.__default.fm2(!(!(!(_2989_v1))), globalState);
              let _3009_v18;
              let _init87 = function (_3010_i6) {
                return new _dafny.CodePoint('g'.codePointAt(0));
              };
              let _nw502 = Array((new BigNumber(21)).toNumber());
              for (let _i0_87 = 0; _i0_87 < new BigNumber(_nw502.length); _i0_87++) {
                _nw502[_i0_87] = _init87(new BigNumber(_i0_87));
              }
              _3009_v18 = _nw502;
              let _3011_v19;
              _3011_v19 = _module.D20.create_DC52(_3009_v18);
              let _3012_v20;
              _3012_v20 = _dafny.Set.fromElements(_3009_v18, _3009_v18, (_3011_v19).dtor_cf76, _3009_v18, _3009_v18);
              r1 = new BigNumber(((_3012_v20).Intersect(_3012_v20)).length);
              (globalState).f4 = !((_dafny.ZERO).minus(p0)).isEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0));
              let _3013_v21;
              _3013_v21 = _dafny.Set.fromElements(p0);
              let _rhs424 = !(_2989_v1);
              let _rhs425 = _2989_v1;
              let _rhs426 = (_3013_v21).IsProperSubsetOf(_3013_v21);
              let _rhs427 = (((_3004_v13).contains(_2989_v1)) ? ((_3004_v13).get(_2989_v1)) : (p0));
              let _lhs256 = globalState;
              let _lhs257 = globalState;
              let _lhs258 = globalState;
              _lhs256.f4 = _rhs424;
              _lhs257.f4 = _rhs425;
              _lhs258.f4 = _rhs426;
              r1 = _rhs427;
            }
            let _3014_v22;
            _3014_v22 = _dafny.MultiSet.fromElements(((_2989_v1) ? ((_dafny.ZERO).minus(p0)) : (new BigNumber(-108))));
            let _rhs428 = (_3014_v22).Intersect(_3014_v22);
            let _rhs429 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
            let _rhs430 = _2991_v3;
            _3014_v22 = _rhs428;
            r1 = _rhs429;
            _2991_v3 = _rhs430;
            let _3015_v23;
            _3015_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2989_v1,_2989_v1);
            r1 = new BigNumber((((_2989_v1) ? (_3015_v23) : (_dafny.Map.Empty.slice().updateUnsafe(_2989_v1,_2989_v1)))).length);
          }
        }
      }
      let _3016_v24;
      _3016_v24 = _dafny.Seq.of(p0);
      let _3017_v25;
      _3017_v25 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(true,_2989_v1));
      let _3018_v26;
      _3018_v26 = _dafny.Seq.of(true, _2989_v1);
      let _3019_v27;
      _3019_v27 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
      let _3020_v28;
      let _init88 = ((_3021_v26) => function (_3022_i8) {
        return _dafny.MultiSet.fromElements(new BigNumber(((_module.D4.create_DC12(_3021_v26)).dtor_cf18).length));
      })(_3018_v26);
      let _nw503 = Array((new BigNumber(20)).toNumber());
      for (let _i0_88 = 0; _i0_88 < new BigNumber(_nw503.length); _i0_88++) {
        _nw503[_i0_88] = _init88(new BigNumber(_i0_88));
      }
      _3020_v28 = _nw503;
      let _3023_v29;
      let _nw504 = new _module.C5();
      _nw504.__ctor(_2989_v1, _3020_v28, _3017_v25, _dafny.Seq.of(_2989_v1, _2989_v1));
      _3023_v29 = _nw504;
      let _3024_v30;
      _3024_v30 = _dafny.Map.Empty.slice().updateUnsafe(_3023_v29,_3019_v27);
      let _3025_v31;
      let _nw505 = Array((new BigNumber(7)).toNumber());
      _nw505[(_dafny.ZERO).toNumber()] = _3019_v27;
      _nw505[(_dafny.ONE).toNumber()] = _3019_v27;
      _nw505[(new BigNumber(2)).toNumber()] = _3019_v27;
      _nw505[(new BigNumber(3)).toNumber()] = _3019_v27;
      _nw505[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(p0, (_dafny.ZERO).minus(new BigNumber(-973)), p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_3018_v26).length),_2989_v1)).length), new BigNumber((_dafny.Seq.of(_2989_v1, _2989_v1, _2989_v1)).length));
      _nw505[(new BigNumber(5)).toNumber()] = (((_3024_v30).contains(_3023_v29)) ? ((_3024_v30).get(_3023_v29)) : (_3019_v27));
      _nw505[(new BigNumber(6)).toNumber()] = _3019_v27;
      _3025_v31 = _nw505;
      let _3026_v32;
      let _nw506 = new _module.C2();
      _nw506.__ctor(_dafny.Seq.Concat(_3016_v24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-379)), ((_3027_p0) => function (_3028_i7) {
        return _3027_p0;
      })(p0))), _2989_v1, _3017_v25, _dafny.Seq.update(_3018_v26, _module.__default.safeIndex(p0, new BigNumber((_3018_v26).length)), false), _3025_v31);
      _3026_v32 = _nw506;
      let _3029_v33;
      _3029_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(739));
      r0 = ((_2989_v1) ? (new BigNumber(206)) : ((((_3029_v33).contains(p0)) ? ((_3029_v33).get(p0)) : (p0))));
      let _3030_v34;
      _3030_v34 = _module.D8.create_DC26(_2989_v1);
      let _pat_let_tv51 = p0;
      let _pat_let_tv52 = p0;
      let _pat_let_tv53 = p0;
      let _pat_let_tv54 = p0;
      let _pat_let_tv55 = p0;
      let _pat_let_tv56 = globalState;
      r1 = function (_source41) {
        if (_source41.is_DC24) {
          let _3031___mcc_h0 = (_source41).cf36;
          let _3032_cf36 = _3031___mcc_h0;
          return _pat_let_tv51;
        } else if (_source41.is_DC25) {
          let _3033___mcc_h1 = (_source41).cf37;
          let _3034___mcc_h2 = (_source41).cf38;
          let _3035___mcc_h3 = (_source41).cf39;
          let _3036___mcc_h4 = (_source41).cf40;
          let _3037_cf40 = _3036___mcc_h4;
          let _3038_cf39 = _3035___mcc_h3;
          let _3039_cf38 = _3034___mcc_h2;
          let _3040_cf37 = _3033___mcc_h1;
          return _pat_let_tv52;
        } else if (_source41.is_DC26) {
          let _3041___mcc_h5 = (_source41).cf41;
          let _3042_cf41 = _3041___mcc_h5;
          return (_dafny.ZERO).minus(_pat_let_tv53);
        } else if (_source41.is_DC23) {
          let _3043___mcc_h6 = (_source41).cf35;
          let _3044_cf35 = _3043___mcc_h6;
          return (_dafny.ZERO).minus(_pat_let_tv54);
        } else {
          let _3045___mcc_h7 = (_source41).cf42;
          let _3046_cf42 = _3045___mcc_h7;
          return (_dafny.ZERO).minus(_pat_let_tv55);
        }
      }(function (_pat_let60_0) {
        return function (_3047_dt__update__tmp_h0) {
          return function (_pat_let61_0) {
            return function (_3048_dt__update_hcf41_h0) {
              return _module.D8.create_DC26(_3048_dt__update_hcf41_h0);
            }(_pat_let61_0);
          }(_module.__default.fm2(true, _pat_let_tv56));
        }(_pat_let60_0);
      }(_3030_v34));
      let _3049_v35;
      _3049_v35 = _module.D0.create_DC1();
      r2 = _3049_v35;
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
