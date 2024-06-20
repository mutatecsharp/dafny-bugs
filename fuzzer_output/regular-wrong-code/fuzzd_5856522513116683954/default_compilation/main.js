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
      return (_dafny.MultiSet.fromElements(false)).equals(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(false))));
    };
    static fm6(globalState) {
      return (new BigNumber((_dafny.Set.fromElements(false, true)).length)).isLessThan(new BigNumber(743));
    };
    static fm9(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.areEqual(_dafny.Seq.UnicodeFromString("gyeq"), _dafny.Seq.UnicodeFromString("gvla")), (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(192), new BigNumber(920)))).equals(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(606))).length))).length))));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements((new BigNumber(-374)).plus(new BigNumber(980)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(252)), function (_0_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length), (new BigNumber((_dafny.Set.fromElements(true)).length)).multipliedBy(new BigNumber(99)));
    };
    static fm11(p0, globalState) {
      return new BigNumber((function (_source0) {
        if (_source0.is_DC23) {
          let _1___mcc_h0 = (_source0).cf35;
          let _2___mcc_h1 = (_source0).cf36;
          let _3___mcc_h2 = (_source0).cf37;
          let _4_cf37 = _3___mcc_h2;
          let _5_cf36 = _2___mcc_h1;
          let _6_cf35 = _1___mcc_h0;
          return (_dafny.Map.Empty.slice().updateUnsafe(_5_cf36,_5_cf36)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_6_cf35));
        } else if (_source0.is_DC24) {
          return (_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),false));
        } else {
          let _7___mcc_h3 = (_source0).cf34;
          let _8_cf34 = _7___mcc_h3;
          return (_module.D16.create_DC39(_dafny.Map.Empty.slice().updateUnsafe(false,false))).dtor_cf57;
        }
      }(_module.D8.create_DC24())).length);
    };
    static fm12(p0, p1, p2, globalState) {
      return _module.D0.create_DC2((_module.D8.create_DC23(true, true, new BigNumber((_dafny.Set.fromElements(new BigNumber(800))).length))).dtor_cf36);
    };
    static fm15(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("ftpspqpu")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(740)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-788))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(673))));
    };
    static fm16(globalState) {
      let _source1 = _module.D8.create_DC24();
      if (_source1.is_DC23) {
        let _9___mcc_h0 = (_source1).cf35;
        let _10___mcc_h1 = (_source1).cf36;
        let _11___mcc_h2 = (_source1).cf37;
        let _12_cf37 = _11___mcc_h2;
        let _13_cf36 = _10___mcc_h1;
        let _14_cf35 = _9___mcc_h0;
        return _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(676)), function (_15_i0) {
          return new BigNumber(263);
        }));
      } else if (_source1.is_DC24) {
        return _dafny.MultiSet.fromElements(_dafny.ONE);
      } else {
        let _16___mcc_h3 = (_source1).cf34;
        let _17_cf34 = _16___mcc_h3;
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(905)));
      }
    };
    static fm17(p0, globalState) {
      let _source2 = _module.D7.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(946),true));
      if (_source2.is_DC19) {
        let _18___mcc_h0 = (_source2).cf27;
        let _19___mcc_h1 = (_source2).cf28;
        let _20___mcc_h2 = (_source2).cf29;
        let _21___mcc_h3 = (_source2).cf30;
        let _22_cf30 = _21___mcc_h3;
        let _23_cf29 = _20___mcc_h2;
        let _24_cf28 = _19___mcc_h1;
        let _25_cf27 = _18___mcc_h0;
        return _module.D0.create_DC1(_23_cf29, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("whpnil")), _dafny.Seq.UnicodeFromString("xdqneud"));
      } else if (_source2.is_DC20) {
        let _26___mcc_h4 = (_source2).cf31;
        let _27___mcc_h5 = (_source2).cf32;
        let _28_cf32 = _27___mcc_h5;
        let _29_cf31 = _26___mcc_h4;
        if (_28_cf32) {
          return _module.D0.create_DC1(_28_cf32, _dafny.Seq.of(_29_cf31), (_module.D7.create_DC20(_29_cf31, _28_cf32)).dtor_cf31);
        } else {
          return _module.D0.create_DC1(true, _dafny.Seq.of(_29_cf31, _29_cf31, _29_cf31, _29_cf31, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-869)), function (_30_i0) {
  return new _dafny.CodePoint('j'.codePointAt(0));
})), _29_cf31);
        }
      } else if (_source2.is_DC18) {
        let _31___mcc_h6 = (_source2).cf26;
        let _32_cf26 = _31___mcc_h6;
        return _module.D0.create_DC1(true, _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_33_i1) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})), _dafny.Seq.UnicodeFromString("narkygd"));
      } else {
        let _34___mcc_h7 = (_source2).cf33;
        let _35_cf33 = _34___mcc_h7;
        return _module.D0.create_DC1(!(false), _dafny.Seq.Create(_module.__default.abs(new BigNumber(558)), function (_36_i2) {
  return _dafny.Seq.UnicodeFromString("u");
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_37_i3) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}));
      }
    };
    static fm18(p0, p1, globalState) {
      return _dafny.Seq.Concat(((!(false)) ? (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(256)), function (_38_i0) {
        return _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("hwq")).length));
      }), _dafny.Seq.of(_module.D0.create_DC0(new BigNumber(694))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(559)), function (_39_i1) {
        return _module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length));
      }))) : (_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(865)), function (_40_i2) {
        return _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("dnmlywcha")).length));
      })))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-414)), function (_41_i3) {
        return _dafny.Seq.of(_module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements(false, false)).length)), _module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("bgl")).length)));
      }));
    };
    static fm19(p0, p1, globalState) {
      return !(false) || (!(!(true)));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(583)), function (_42_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("h"));
    };
    static fm22(p0, p1, p2, globalState) {
      return new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)))), _module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))), _dafny.Seq.of(_module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)))))))).Difference(_dafny.MultiSet.fromElements(_module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)))), _module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))), _module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0)))), _module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)))), _module.D14.create_DC35(_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))))))).cardinality());
    };
    static fm23(p0, p1, p2, p3, globalState) {
      let _source3 = _module.D9.create_DC27(!(false), new BigNumber((_dafny.Seq.of(new BigNumber(-433), new BigNumber(600))).length), false);
      if (_source3.is_DC26) {
        let _43___mcc_h0 = (_source3).cf39;
        let _44___mcc_h1 = (_source3).cf40;
        let _45___mcc_h2 = (_source3).cf41;
        let _46_cf41 = _45___mcc_h2;
        let _47_cf40 = _44___mcc_h1;
        let _48_cf39 = _43___mcc_h0;
        return _47_cf40;
      } else if (_source3.is_DC27) {
        let _49___mcc_h3 = (_source3).cf42;
        let _50___mcc_h4 = (_source3).cf43;
        let _51___mcc_h5 = (_source3).cf44;
        let _52_cf44 = _51___mcc_h5;
        let _53_cf43 = _50___mcc_h4;
        let _54_cf42 = _49___mcc_h3;
        return new _dafny.CodePoint('a'.codePointAt(0));
      } else {
        let _55___mcc_h6 = (_source3).cf38;
        let _56_cf38 = _55___mcc_h6;
        return new _dafny.CodePoint('e'.codePointAt(0));
      }
    };
    static fm24(globalState) {
      return (((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(true,false))) : (_dafny.Map.Empty.slice().updateUnsafe((_module.D9.create_DC26(true, new _dafny.CodePoint('a'.codePointAt(0)), true)).dtor_cf40,_dafny.Map.Empty.slice().updateUnsafe(!(false),false))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(true,true)));
    };
    static fm25(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-739), new BigNumber(362), new BigNumber((_dafny.Seq.of(false)).length))).length))).Union(_dafny.MultiSet.fromElements(new BigNumber(600), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(612)), function (_57_i0) {
        return new BigNumber(801);
      }))).cardinality()), new BigNumber(708), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(true), !(true))).length)), new BigNumber((_dafny.Seq.UnicodeFromString("hbr")).length)))).Intersect(((_module.D12.create_DC31(_dafny.MultiSet.fromElements(new BigNumber(108)))).dtor_cf48).Union(_dafny.MultiSet.fromElements(new BigNumber(417))));
    };
    static fm26(p0, p1, globalState) {
      let _source4 = _module.D4.create_DC12(new BigNumber(642));
      if (_source4.is_DC11) {
        return _module.D1.create_DC5(!(false), _dafny.Seq.UnicodeFromString("hknqq"), _dafny.Seq.of(!(false), false));
      } else if (_source4.is_DC12) {
        let _58___mcc_h0 = (_source4).cf18;
        let _59_cf18 = _58___mcc_h0;
        return _module.D1.create_DC5(true, _dafny.Seq.UnicodeFromString("tdv"), _dafny.Seq.of(true));
      } else {
        let _60___mcc_h1 = (_source4).cf17;
        let _61_cf17 = _60___mcc_h1;
        return _module.D1.create_DC5(true, _dafny.Seq.UnicodeFromString("spg"), _dafny.Seq.of(false));
      }
    };
    static fm27(p0, p1, p2, globalState) {
      return _dafny.Seq.of(!(!(true) || (false)));
    };
    static fm28(globalState) {
      if ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(984)), function (_62_i0) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length)).isLessThan(new BigNumber(497))) {
        if (!(false)) {
          return _module.D0.create_DC3(_module.D0.create_DC2(true));
        } else {
          return _module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC1(!(false), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(442)), function (_63_i1) {
  return new _dafny.CodePoint('p'.codePointAt(0));
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(958)), function (_64_i2) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})), _dafny.Seq.UnicodeFromString("fwyar"))));
        }
      } else {
        return _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber(296)));
      }
    };
    static fm29(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC7((new BigNumber((_dafny.Seq.UnicodeFromString("tfptdctot")).length)).isEqualTo(new BigNumber(855)), _dafny.Map.Empty.slice().updateUnsafe(false,_module.D0.create_DC3(_module.D0.create_DC0(new BigNumber(272)))), (new BigNumber(987)).multipliedBy(new BigNumber((_dafny.Seq.of(false, true, true, !(false), true)).length)), !(!(true) || (false)));
    };
    static fm30(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(93)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(25)), function (_65_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(753), new BigNumber(220))).length));
      }));
    };
    static fm31(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-89)), function (_66_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("mijsmoem"))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_module.D2.create_DC7(!(true), _dafny.Map.Empty.slice().updateUnsafe(false,_module.D0.create_DC3(_module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(430))).length)))), new BigNumber(710), true)).dtor_cf13)).length));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(990),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(false)).length))).length),false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),false));
    };
    static fm33(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(!(false), true, false, true)).Union((_dafny.Set.fromElements(false, !(false), !(true), true)).Difference((_dafny.Set.fromElements(false))));
    };
    static fm34(globalState) {
      if ((_dafny.Set.fromElements(new BigNumber(14), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, false, !(false))).cardinality())))).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("amok")).length)))) {
        return (_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(-406), new BigNumber((_dafny.Seq.of(false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(true), true)).length)))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((function () {
          let _coll0 = new _dafny.Map();
          for (const _compr_0 of _dafny.IntegerRange(new BigNumber(299), new BigNumber(729))) {
            let _67_v0 = _compr_0;
            if (((new BigNumber(299)).isLessThanOrEqualTo(_67_v0)) && ((_67_v0).isLessThan(new BigNumber(729)))) {
              _coll0.push([(_67_v0).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),false)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mwcb")).length))).length))).length)),new _dafny.CodePoint('u'.codePointAt(0))]);
            }
          }
          return _coll0;
        }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(777)), function (_68_i0) {
          return new BigNumber(962);
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(589)), function (_69_i1) {
          return new BigNumber(159);
        }), _dafny.Seq.of(new BigNumber((function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), function (_70_i2) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length);
          })).length))).Elements) {
            let _71_v1 = _compr_1;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), function (_70_i2) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length);
            })).length)), _71_v1)) {
              _coll1.push([_module.__default.safeDivisionInt(_71_v1, new BigNumber(672)),new BigNumber((_dafny.Seq.of(new BigNumber(228), new BigNumber((function () {
                let _coll2 = new _dafny.Set();
                for (const _compr_2 of _dafny.IntegerRange(new BigNumber(228), new BigNumber(681))) {
                  let _72_v2 = _compr_2;
                  if (((new BigNumber(228)).isLessThanOrEqualTo(_72_v2)) && ((_72_v2).isLessThan(new BigNumber(681)))) {
                    _coll2.add(_module.__default.safeModuloInt(_72_v2, new BigNumber(70)));
                  }
                }
                return _coll2;
              }()).length))).length)]);
            }
          }
          return _coll1;
        }()).length))));
      } else {
        return (_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("juf")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), function (_73_i3) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(919), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("uise")).length),new BigNumber(23))).length)))).cardinality())), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(!(true)))).length)), _dafny.Seq.of(new BigNumber(739), new BigNumber(490), new BigNumber((_dafny.Seq.of(new BigNumber(738))).length)))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(false), false)).length), new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(487)), function (_74_i4) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        })).length))));
      }
    };
    static fm35(p0, p1, p2, globalState) {
      return _module.D4.create_DC12(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-151), new BigNumber((_dafny.Set.fromElements(new BigNumber(788))).length)), _dafny.Seq.of(new BigNumber(-655)))).length));
    };
    static fm36(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),new BigNumber(238));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm38(globalState) {
      return _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), function (_75_i0) {
        return new BigNumber(-886);
      }));
    };
    static fm39(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(!(true), !(false))).Union(_dafny.MultiSet.fromElements(false, true, false))).Intersect((_dafny.MultiSet.fromElements(false)).Difference(_dafny.MultiSet.fromElements(true, false)));
    };
    static fm40(p0, p1, globalState) {
      return _module.D7.create_DC20(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-531)), function (_76_i0) {
  return new _dafny.CodePoint('p'.codePointAt(0));
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(874)), function (_77_i1) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})), !((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber(979))).length)).isLessThan((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, false)).length))))));
    };
    static fm41(p0, p1, globalState) {
      return _module.D9.create_DC26(!(_dafny.Set.fromElements(true, true)).contains(!(true)), new _dafny.CodePoint('s'.codePointAt(0)), !(!(!(true))));
    };
    static fm42(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_module.D0.create_DC1(false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-783)), function (_78_i0) {
  return _dafny.Seq.UnicodeFromString("jtdyx");
}), _dafny.Seq.UnicodeFromString("pj")), _module.D0.create_DC1(!(false), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("woy"), _dafny.Seq.UnicodeFromString("dwuebhwk")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(321)), function (_79_i1) {
  return new _dafny.CodePoint('r'.codePointAt(0));
})), _module.D0.create_DC1(false, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ue")), _dafny.Seq.UnicodeFromString("wo")))).Intersect(_dafny.MultiSet.fromElements(_module.D0.create_DC1(!(!(true)), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wg"), _dafny.Seq.UnicodeFromString("jwmnfd")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(475)), function (_80_i2) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})), _module.D0.create_DC1(false, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("jmo")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_81_i3) {
  return new _dafny.CodePoint('q'.codePointAt(0));
})), _module.D0.create_DC1(false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), function (_82_i4) {
  return _dafny.Seq.UnicodeFromString("boc");
}), _dafny.Seq.UnicodeFromString("cihhuxg"))));
    };
    static fm43(p0, globalState) {
      return ((_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false)), _dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(!(!(true))))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(!(true))))).Intersect((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true, true))).Union(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, false, true, false), _dafny.MultiSet.fromElements(true))));
    };
    static fm44(p0, p1, p2, globalState) {
      if (((false) ? (false) : (false))) {
        return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),!(!(true))), _dafny.Map.Empty.slice().updateUnsafe(true,false));
      } else {
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(false,!(false)), _dafny.Map.Empty.slice().updateUnsafe(false,false))));
      }
    };
    static Main(__noArgsParameter) {
      let _83_v0;
      let _nw0 = Array((new BigNumber(14)).toNumber()).fill(false);
      _83_v0 = _nw0;
      let _84_v1;
      let _nw1 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
      _84_v1 = _nw1;
      let _85_v2;
      _85_v2 = _dafny.Seq.UnicodeFromString("fno");
      let _86_v3;
      _86_v3 = new BigNumber(636);
      let _87_v4;
      _87_v4 = new _dafny.CodePoint('a'.codePointAt(0));
      let _88_v5;
      _88_v5 = _dafny.Set.fromElements((_85_v2)[_module.__default.safeIndex(_86_v3, new BigNumber((_85_v2).length))], _87_v4);
      let _89_v6;
      _89_v6 = _dafny.Map.Empty.slice().updateUnsafe(_88_v5,_86_v3);
      let _90_v7;
      _90_v7 = _dafny.Seq.of(_89_v6, _89_v6);
      let _91_v8;
      _91_v8 = false;
      let _92_v9;
      _92_v9 = _dafny.MultiSet.fromElements(_91_v8);
      let _93_v10;
      _93_v10 = _dafny.Map.Empty.slice().updateUnsafe(_91_v8,_91_v8);
      let _94_v11;
      _94_v11 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_91_v8,_91_v8), _93_v10, _93_v10);
      let _95_v12;
      _95_v12 = _dafny.Map.Empty.slice().updateUnsafe(_86_v3,_92_v9);
      let _96_v13;
      _96_v13 = _dafny.Set.fromElements(_91_v8);
      let _97_globalState;
      let _nw2 = new _module.GlobalState();
      _nw2.__ctor(new _dafny.CodePoint('b'.codePointAt(0)), _83_v0, _84_v1, true, (_90_v7)[_module.__default.safeIndex((((_92_v9).contains(true)) ? ((_92_v9).get(true)) : (new BigNumber((_94_v11).cardinality()))), new BigNumber((_90_v7).length))], true, new BigNumber(-209), (_95_v12).Merge(_95_v12), _dafny.Seq.Concat(_85_v2, _85_v2), true, true, _96_v13, _85_v2, false);
      _97_globalState = _nw2;
      _83_v0 = _83_v0;
      let _98_v14;
      let _nw3 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _98_v14 = _nw3;
      let _index0 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
      (_98_v14)[_index0] = _86_v3;
      let _index1 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
      (_98_v14)[_index1] = _86_v3;
      let _index2 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
      (_83_v0)[_index2] = !((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_85_v2,_85_v2)).length)));
      let _99_v16;
      _99_v16 = _dafny.Seq.of((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]);
      let _100_v17;
      _100_v17 = _dafny.Seq.of(new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_99_v16).Elements) {
          let _101_v15 = _compr_3;
          if (_dafny.Seq.contains(_99_v16, _101_v15)) {
            _coll3.push([(_101_v15).multipliedBy((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]),_91_v8]);
          }
        }
        return _coll3;
      }()).length));
      let _102_v18;
      _102_v18 = _dafny.MultiSet.fromElements(_99_v16);
      let _103_v19;
      _103_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(964),_86_v3);
      let _104_v20;
      _104_v20 = _dafny.MultiSet.fromElements((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], new BigNumber((_92_v9).cardinality()), (((_103_v19).contains(_86_v3)) ? ((_103_v19).get(_86_v3)) : ((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])));
      let _index3 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
      let _index4 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
      let _index5 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
      let _rhs0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]), _86_v3);
      let _rhs1 = ((_104_v20).Union(_104_v20)).IsSubsetOf(_dafny.MultiSet.fromElements((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], (((_102_v18).contains(_100_v17)) ? ((_102_v18).get(_100_v17)) : (_86_v3)), _86_v3, new BigNumber(-296)));
      let _rhs2 = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("qahncvdqm"), (_85_v2)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_93_v10).length)), new BigNumber((_85_v2).length))]);
      let _rhs3 = ((new BigNumber(145)).plus(new BigNumber(482))).isLessThan(_module.__default.safeModuloInt(_86_v3, (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]));
      let _rhs4 = (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))];
      let _lhs0 = _98_v14;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
      let _lhs2 = _97_globalState;
      let _lhs3 = _83_v0;
      let _lhs4 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
      let _lhs5 = _98_v14;
      let _lhs6 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
      _lhs0[_lhs1] = _rhs0;
      _91_v8 = _rhs1;
      _lhs2.f9 = _rhs2;
      _lhs3[_lhs4] = _rhs3;
      _lhs5[_lhs6] = _rhs4;
      let _105_v21;
      _105_v21 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_91_v8,_91_v8)).update((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))], (_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]),_98_v14);
      _98_v14 = (((_105_v21).contains(_93_v10)) ? ((_105_v21).get(_93_v10)) : (_98_v14));
      _86_v3 = (_dafny.ZERO).minus(new BigNumber(-954));
      let _106_v22;
      let _nw4 = new _module.C7();
      _nw4.__ctor();
      _106_v22 = _nw4;
      let _107_v23;
      let _nw5 = new _module.C5();
      _nw5.__ctor();
      _107_v23 = _nw5;
      _91_v8 = ((!_dafny.Seq.contains(_dafny.Seq.of(_107_v23), _107_v23)) ? (_91_v8) : ((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]));
      let _hi0 = new BigNumber((_module.__default.fm20((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], _91_v8, _91_v8, _86_v3, _97_globalState)).length);
      for (let _108_i0 = ((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]).multipliedBy(_86_v3); _108_i0.isLessThan(_hi0); _108_i0 = _108_i0.plus(_dafny.ONE)) {
        let _109_v24;
        let _nw6 = Array((new BigNumber(5)).toNumber());
        _nw6[(_dafny.ZERO).toNumber()] = _87_v4;
        _nw6[(_dafny.ONE).toNumber()] = _87_v4;
        _nw6[(new BigNumber(2)).toNumber()] = _87_v4;
        _nw6[(new BigNumber(3)).toNumber()] = _87_v4;
        _nw6[(new BigNumber(4)).toNumber()] = _87_v4;
        _109_v24 = _nw6;
        let _110_v25;
        _110_v25 = _dafny.Map.Empty.slice().updateUnsafe(_109_v24,_86_v3);
        _110_v25 = _dafny.Map.Empty.slice().updateUnsafe(_109_v24,_108_i0);
        _93_v10 = (_93_v10).update(_91_v8, (_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]);
        let _111_v26;
        let _nw7 = new _module.C1();
        _nw7.__ctor();
        _111_v26 = _nw7;
        _85_v2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), ((_112_v4) => function (_113_i1) {
          return _112_v4;
        })(_87_v4));
      }
      let _114_v27;
      _114_v27 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-90)), ((_115_v4) => function (_116_i2) {
        return _115_v4;
      })(_87_v4))).length));
      let _117_v28;
      _117_v28 = _module.D9.create_DC25(_114_v27);
      let _118_v29;
      _118_v29 = _dafny.Map.Empty.slice().updateUnsafe(_86_v3,_117_v28);
      let _rhs5 = new BigNumber((_118_v29).length);
      let _rhs6 = _85_v2;
      _86_v3 = _rhs5;
      _85_v2 = _rhs6;
      let _119_v30;
      _119_v30 = _dafny.Map.Empty.slice().updateUnsafe((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))],_85_v2);
      let _hi1 = new BigNumber(((((_119_v30).contains(new BigNumber(288))) ? ((_119_v30).get(new BigNumber(288))) : (_85_v2))).length);
      for (let _120_i3 = (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]; _120_i3.isLessThan(_hi1); _120_i3 = _120_i3.plus(_dafny.ONE)) {
        let _121_v31;
        _121_v31 = _dafny.MultiSet.fromElements(_87_v4, _87_v4);
        let _122_v32;
        _122_v32 = _module.D14.create_DC35(_121_v31);
        let _123_v33;
        _123_v33 = _dafny.Set.fromElements(new BigNumber((_85_v2).length));
        let _124_v34;
        _124_v34 = _dafny.Seq.of(_module.__default.fm19(new BigNumber((_123_v33).length), (_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))], _97_globalState));
        if (!((_121_v31).IsSubsetOf((_122_v32).dtor_cf53)) || (_dafny.Seq.IsPrefixOf(_124_v34, _124_v34))) {
          (_97_globalState).f5 = (_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))];
          _86_v3 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber(85), (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]), _module.__default.safeDivisionInt(new BigNumber((_85_v2).length), (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]));
          let _125_v35;
          _125_v35 = _dafny.Set.fromElements(_85_v2);
          let _126_v36;
          let _nw8 = new _module.C2();
          _nw8.__ctor();
          _126_v36 = _nw8;
          let _127_v37;
          _127_v37 = _module.D4.create_DC12(_module.__default.safeDivisionInt(new BigNumber(-645), new BigNumber((_92_v9).cardinality())));
          let _128_v38;
          _128_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(44),_125_v35);
          let _129_v39;
          _129_v39 = _dafny.Map.Empty.slice().updateUnsafe((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))],_125_v35);
          let _rhs7 = ((((_128_v38).contains((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])) ? ((_128_v38).get((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])) : (_125_v35))).Intersect((_125_v35).Union((((_129_v39).contains(_91_v8)) ? ((_129_v39).get(_91_v8)) : (_125_v35))));
          let _rhs8 = _126_v36;
          let _rhs9 = _127_v37;
          let _rhs10 = _91_v8;
          let _lhs7 = _97_globalState;
          _125_v35 = _rhs7;
          _126_v36 = _rhs8;
          _127_v37 = _rhs9;
          _lhs7.f5 = _rhs10;
          let _130_v40;
          let _nw9 = new _module.C1();
          _nw9.__ctor();
          _130_v40 = _nw9;
          let _131_v41;
          _131_v41 = _dafny.Map.Empty.slice().updateUnsafe(!((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]),(_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]);
          let _132_v42;
          _132_v42 = _dafny.Map.Empty.slice().updateUnsafe((_100_v17)[_module.__default.safeIndex((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], new BigNumber((_100_v17).length))],_130_v40);
          let _index6 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
          let _rhs11 = (((_132_v42).contains(_120_i3)) ? ((_132_v42).get(_120_i3)) : (_130_v40));
          let _rhs12 = (_dafny.MultiSet.fromElements(_86_v3)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_120_i3, (_99_v16)[_module.__default.safeIndex(_120_i3, new BigNumber((_99_v16).length))]));
          let _rhs13 = _dafny.Seq.Concat(_124_v34, _124_v34);
          let _rhs14 = _103_v19;
          let _rhs15 = (_dafny.Map.Empty.slice().updateUnsafe((((_93_v10).contains((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))])) ? ((_93_v10).get((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))])) : (_91_v8)),new BigNumber((_85_v2).length))).Merge((_131_v41).Merge(_131_v41));
          let _lhs8 = _83_v0;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
          _130_v40 = _rhs11;
          _lhs8[_lhs9] = _rhs12;
          _124_v34 = _rhs13;
          _103_v19 = _rhs14;
          _131_v41 = _rhs15;
          let _133_v43;
          let _nw10 = new _module.C6();
          _nw10.__ctor();
          _133_v43 = _nw10;
        } else {
          let _index7 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
          (_98_v14)[_index7] = (((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]).multipliedBy(_86_v3)).plus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_120_i3, (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])));
          let _134_v44;
          _134_v44 = _dafny.Map.Empty.slice().updateUnsafe(_98_v14,(_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]);
          let _135_v45;
          let _nw11 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
          _135_v45 = _nw11;
          let _136_v46;
          _136_v46 = _module.D1.create_DC5((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))], _85_v2, _124_v34);
          let _137_v47;
          _137_v47 = _dafny.MultiSet.fromElements(_module.D1.create_DC5((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), ((_138_v4) => function (_139_i4) {
  return _138_v4;
})(_87_v4)), _124_v34), _136_v46);
          let _index8 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_135_v45).length));
          (_135_v45)[_index8] = _137_v47;
          let _index9 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_135_v45).length));
          let _rhs16 = _134_v44;
          let _rhs17 = _dafny.MultiSet.fromElements(_136_v46, _136_v46, _module.__default.fm26(true, _91_v8, _97_globalState));
          let _rhs18 = _117_v28;
          let _lhs10 = _135_v45;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_135_v45).length));
          _134_v44 = _rhs16;
          _lhs10[_lhs11] = _rhs17;
          _117_v28 = _rhs18;
          (_97_globalState).f5 = !(_91_v8);
          let _140_v48;
          _140_v48 = _dafny.Map.Empty.slice().updateUnsafe((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))],_91_v8);
          let _141_v49;
          _141_v49 = _module.D7.create_DC18(_140_v48);
          let _pat_let_tv0 = _91_v8;
          _141_v49 = function (_pat_let0_0) {
            return function (_142_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_143_dt__update_hcf26_h0) {
                  return _module.D7.create_DC18(_143_dt__update_hcf26_h0);
                }(_pat_let1_0);
              }(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(506),_pat_let_tv0));
            }(_pat_let0_0);
          }(_141_v49);
          (_97_globalState).f13 = _module.__default.fm6(_97_globalState);
        }
        let _144_v50;
        let _nw12 = new _module.C3();
        _nw12.__ctor(!((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]), _120_i3);
        _144_v50 = _nw12;
        let _145_v51;
        _145_v51 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_module.__default.fm11(_91_v8, _97_globalState), new BigNumber((_124_v34).length)), _104_v20);
        (_107_v23).m3(_86_v3, _144_v50, (_144_v50).fm0(_86_v3, _87_v4, _97_globalState), (_145_v51)[_module.__default.safeIndex((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], new BigNumber((_145_v51).length))], _97_globalState);
        let _index10 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
        (_83_v0)[_index10] = (_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))];
        let _index11 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
        (_98_v14)[_index11] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_86_v3));
      }
      let _146_v52;
      _146_v52 = _dafny.Map.Empty.slice().updateUnsafe(_86_v3,(_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]);
      let _147_v53;
      _147_v53 = _module.D7.create_DC18(_146_v52);
      let _148_v54;
      _148_v54 = _dafny.Seq.of((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]);
      let _149_v55;
      _149_v55 = _dafny.Seq.of(_85_v2, _85_v2);
      let _150_v56;
      _150_v56 = _module.D0.create_DC1(_91_v8, _149_v55, _85_v2);
      let _151_v57;
      _151_v57 = _module.D0.create_DC3(_module.D0.create_DC3(_150_v56));
      let _pat_let_tv1 = _147_v53;
      let _pat_let_tv2 = _147_v53;
      let _pat_let_tv3 = _146_v52;
      let _pat_let_tv4 = _146_v52;
      let _pat_let_tv5 = _147_v53;
      let _rhs19 = _93_v10;
      let _rhs20 = (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))];
      let _rhs21 = (_148_v54)[_module.__default.safeIndex(_86_v3, new BigNumber((_148_v54).length))];
      let _rhs22 = _85_v2;
      let _rhs23 = function (_source5) {
        if (_source5.is_DC1) {
          let _152___mcc_h0 = (_source5).cf1;
          let _153___mcc_h1 = (_source5).cf2;
          let _154___mcc_h2 = (_source5).cf3;
          let _155_cf3 = _154___mcc_h2;
          let _156_cf2 = _153___mcc_h1;
          let _157_cf1 = _152___mcc_h0;
          return _pat_let_tv1;
        } else if (_source5.is_DC2) {
          let _158___mcc_h3 = (_source5).cf4;
          let _159_cf4 = _158___mcc_h3;
          return _pat_let_tv2;
        } else if (_source5.is_DC0) {
          let _160___mcc_h4 = (_source5).cf0;
          let _161_cf0 = _160___mcc_h4;
          let _162_dt__update__tmp_h1 = _module.D7.create_DC18(_pat_let_tv3);
          let _163_dt__update_hcf26_h1 = _pat_let_tv4;
          return _module.D7.create_DC18(_163_dt__update_hcf26_h1);
        } else {
          let _164___mcc_h5 = (_source5).cf5;
          let _165_cf5 = _164___mcc_h5;
          return _module.D7.create_DC18((_pat_let_tv5).dtor_cf26);
        }
      }(_151_v57);
      let _lhs12 = _97_globalState;
      _93_v10 = _rhs19;
      _86_v3 = _rhs20;
      _lhs12.f13 = _rhs21;
      _85_v2 = _rhs22;
      _147_v53 = _rhs23;
      let _166_v58;
      _166_v58 = _dafny.Map.Empty.slice().updateUnsafe(_91_v8,_86_v3);
      let _167_v59;
      _167_v59 = _module.D9.create_DC26((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))], _87_v4, _91_v8);
      let _168_v60;
      let _nw13 = Array((new BigNumber(29)).toNumber());
      _nw13[(_dafny.ZERO).toNumber()] = _87_v4;
      _nw13[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
      _nw13[(new BigNumber(2)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(3)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(4)).toNumber()] = _module.__default.fm23((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))], _dafny.Map.Empty.slice().updateUnsafe(_87_v4,_93_v10), false, (((_166_v58).contains(_91_v8)) ? ((_166_v58).get(_91_v8)) : ((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])), _97_globalState);
      _nw13[(new BigNumber(5)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(6)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(7)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(8)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(9)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(10)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(11)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(12)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(13)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(14)).toNumber()] = (_167_v59).dtor_cf40;
      _nw13[(new BigNumber(15)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(16)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(17)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(18)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(19)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(20)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(21)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(22)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(23)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(24)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(25)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(26)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(27)).toNumber()] = _87_v4;
      _nw13[(new BigNumber(28)).toNumber()] = _87_v4;
      _168_v60 = _nw13;
      let _index12 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_168_v60).length));
      (_168_v60)[_index12] = _87_v4;
      let _index13 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_168_v60).length));
      (_168_v60)[_index13] = _87_v4;
      let _169_v61;
      let _nw14 = new _module.C5();
      _nw14.__ctor();
      _169_v61 = _nw14;
      let _170_v62;
      _170_v62 = _dafny.Map.Empty.slice().updateUnsafe(false,_169_v61);
      let _171_v63;
      _171_v63 = _dafny.Set.fromElements(_170_v62, _170_v62);
      let _172_i5;
      _172_i5 = _dafny.ZERO;
      L0: {
        while (((((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]) ? (_171_v63) : (_171_v63))).IsDisjointFrom(_171_v63)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_172_i5)) {
              break L0;
            }
            _172_i5 = (_172_i5).plus(_dafny.ONE);
            _86_v3 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(357), _86_v3, _86_v3, (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])).length))))).plus((_module.D0.create_DC0((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])).dtor_cf0);
            (_97_globalState).f13 = (((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]) ? (_91_v8) : (_91_v8));
            let _173_v64;
            _173_v64 = _dafny.Map.Empty.slice().updateUnsafe(_86_v3,(((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]) ? (_98_v14) : (_98_v14)));
            let _174_v65;
            _174_v65 = _module.D11.create_DC30(false);
            let _175_v66;
            _175_v66 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(465)), ((_176_v60) => function (_177_i6) {
              return (_176_v60)[_module.__default.safeIndex(new BigNumber(873), new BigNumber((_176_v60).length))];
            })(_168_v60))).length)),_174_v65);
            let _178_v67;
            _178_v67 = _module.D0.create_DC2((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]);
            let _index14 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
            let _index15 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
            let _rhs24 = _module.__default.safeDivisionInt(((_91_v8) ? (new BigNumber((_175_v66).length)) : (_86_v3)), (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]);
            let _rhs25 = _dafny.Map.Empty.slice().updateUnsafe((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))],_98_v14);
            let _rhs26 = _169_v61;
            let _rhs27 = (_107_v23).fm5(_178_v67, _97_globalState);
            let _rhs28 = !(false);
            let _lhs13 = _98_v14;
            let _lhs14 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
            let _lhs15 = _98_v14;
            let _lhs16 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length));
            let _lhs17 = _97_globalState;
            _lhs13[_lhs14] = _rhs24;
            _173_v64 = _rhs25;
            _169_v61 = _rhs26;
            _lhs15[_lhs16] = _rhs27;
            _lhs17.f13 = _rhs28;
            let _179_v68;
            let _nw15 = new _module.C1();
            _nw15.__ctor();
            _179_v68 = _nw15;
            let _180_v69;
            _180_v69 = _dafny.Set.fromElements(_179_v68, _179_v68);
            (_97_globalState).f3 = ((_180_v69).Intersect((_dafny.Set.fromElements(_179_v68)))).equals(((_module.__default.fm6(_97_globalState)) ? (_180_v69) : (_180_v69)));
          }
        }
      }
      if (_91_v8) {
        (_97_globalState).f13 = !(_module.__default.fm2(((_91_v8) ? (new _dafny.CodePoint('d'.codePointAt(0))) : (_87_v4)), _151_v57, (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], _97_globalState));
        let _181_v70;
        _181_v70 = _dafny.Set.fromElements((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]);
        let _182_v71;
        _182_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]) === (true)),(_181_v70).Union(_181_v70));
        let _183_v72;
        _183_v72 = _module.D4.create_DC10(_181_v70);
        _182_v71 = (_182_v71).update((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))], (_183_v72).dtor_cf17);
        (_97_globalState).f13 = _91_v8;
        let _184_v73;
        _184_v73 = _dafny.Map.Empty.slice().updateUnsafe(_98_v14,_85_v2);
        let _185_v74;
        let _nw16 = Array((new BigNumber(10)).toNumber()).fill([]);
        _185_v74 = _nw16;
        let _186_v75;
        let _nw17 = Array((new BigNumber(26)).toNumber());
        _186_v75 = _nw17;
        let _index16 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_185_v74).length));
        (_185_v74)[_index16] = _186_v75;
        let _187_v76;
        _187_v76 = _dafny.Seq.of(_83_v0, _83_v0, _83_v0, _83_v0, _83_v0);
        let _index17 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_185_v74).length));
        let _rhs29 = ((_184_v73).Merge(_184_v73)).Merge(((false) ? ((_184_v73).update(_98_v14, _85_v2)) : (_184_v73)));
        let _rhs30 = false;
        let _rhs31 = _186_v75;
        let _rhs32 = _dafny.Seq.contains(_dafny.Seq.Concat(_148_v54, _148_v54), _dafny.Seq.contains(_148_v54, _91_v8));
        let _rhs33 = _dafny.Seq.Concat(_dafny.Seq.Concat(_187_v76, _187_v76), _187_v76);
        let _lhs18 = _97_globalState;
        let _lhs19 = _185_v74;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_185_v74).length));
        let _lhs21 = _97_globalState;
        _184_v73 = _rhs29;
        _lhs18.f9 = _rhs30;
        _lhs19[_lhs20] = _rhs31;
        _lhs21.f13 = _rhs32;
        _187_v76 = _rhs33;
        (_107_v23).m3((_86_v3).plus(_86_v3), _169_v61, (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], (_module.__default.fm16(_97_globalState)).Difference(_104_v20), _97_globalState);
      } else {
        _148_v54 = _dafny.Seq.of((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]);
        let _188_v77;
        _188_v77 = _dafny.Map.Empty.slice().updateUnsafe((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))],_169_v61);
        let _189_v78;
        let _nw18 = new _module.C2();
        _nw18.__ctor();
        _189_v78 = _nw18;
        let _190_v79;
        _190_v79 = _dafny.Map.Empty.slice().updateUnsafe((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))],_189_v78);
        (_107_v23).m3(_module.__default.safeModuloInt((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], _86_v3), (((_188_v77).contains(_86_v3)) ? ((_188_v77).get(_86_v3)) : (_169_v61)), new BigNumber((_dafny.MultiSet.fromElements(_190_v79)).cardinality()), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))],new BigNumber((_146_v52).length))).length), (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], new BigNumber((_85_v2).length), (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]), _97_globalState);
        let _191_v81;
        _191_v81 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_86_v3,true), _146_v52, function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-331), new BigNumber(-333))) {
            let _192_v80 = _compr_4;
            if (((new BigNumber(-331)).isLessThanOrEqualTo(_192_v80)) && ((_192_v80).isLessThan(new BigNumber(-333)))) {
              _coll4.push([_module.__default.safeModuloInt(_192_v80, (_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))]),(_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]]);
            }
          }
          return _coll4;
        }(), _146_v52, _dafny.Map.Empty.slice().updateUnsafe((_99_v16)[_module.__default.safeIndex((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], new BigNumber((_99_v16).length))],(_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]));
        _86_v3 = (_dafny.ZERO).minus(((!(_91_v8)) ? (new BigNumber((_191_v81).length)) : (new BigNumber((_96_v13).length))));
        let _193_v82;
        _193_v82 = _dafny.Map.Empty.slice().updateUnsafe(_86_v3,_dafny.Map.Empty.slice().updateUnsafe(_91_v8,_91_v8));
        let _194_v83;
        _194_v83 = _dafny.Seq.of((((_193_v82).contains(_86_v3)) ? ((_193_v82).get(_86_v3)) : (_93_v10)), _93_v10);
        let _195_v84;
        _195_v84 = _dafny.Seq.of(_194_v83, _194_v83);
        let _196_v85;
        let _nw19 = Array((new BigNumber(16)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = (_94_v11).update(_dafny.Map.Empty.slice().updateUnsafe((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))],(_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]), _module.__default.abs(new BigNumber(((_146_v52).update((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], _91_v8)).length)));
        _nw19[(_dafny.ONE).toNumber()] = _94_v11;
        _nw19[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_93_v10);
        _nw19[(new BigNumber(3)).toNumber()] = _94_v11;
        _nw19[(new BigNumber(4)).toNumber()] = _module.__default.fm44((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))], _86_v3, _91_v8, _97_globalState);
        _nw19[(new BigNumber(5)).toNumber()] = (_94_v11).Difference(_94_v11);
        _nw19[(new BigNumber(6)).toNumber()] = _94_v11;
        _nw19[(new BigNumber(7)).toNumber()] = _94_v11;
        _nw19[(new BigNumber(8)).toNumber()] = _94_v11;
        _nw19[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray((_195_v84)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_98_v14)[_module.__default.safeIndex(new BigNumber(484), new BigNumber((_98_v14).length))])), new BigNumber((_195_v84).length))]);
        _nw19[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))],false), _93_v10, _93_v10, _93_v10);
        _nw19[(new BigNumber(11)).toNumber()] = (((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]) ? (_94_v11) : (_94_v11));
        _nw19[(new BigNumber(12)).toNumber()] = _94_v11;
        _nw19[(new BigNumber(13)).toNumber()] = _94_v11;
        _nw19[(new BigNumber(14)).toNumber()] = (_94_v11).update(_93_v10, _module.__default.abs(_86_v3));
        _nw19[(new BigNumber(15)).toNumber()] = _94_v11;
        _196_v85 = _nw19;
        let _index18 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_196_v85).length));
        (_196_v85)[_index18] = (_94_v11).Union(_94_v11);
        let _index19 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_196_v85).length));
        let _index20 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
        let _rhs34 = _module.__default.safeModuloInt((((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]) ? (_86_v3) : (_86_v3)), _86_v3);
        let _rhs35 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))],_91_v8));
        let _rhs36 = !((_83_v0)[_module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length))]);
        let _rhs37 = _91_v8;
        let _lhs22 = _196_v85;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_196_v85).length));
        let _lhs24 = _83_v0;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(454), new BigNumber((_83_v0).length));
        let _lhs26 = _97_globalState;
        _86_v3 = _rhs34;
        _lhs22[_lhs23] = _rhs35;
        _lhs24[_lhs25] = _rhs36;
        _lhs26.f9 = _rhs37;
        _86_v3 = _86_v3;
      }
      let _197_v86;
      let _nw20 = new _module.C2();
      _nw20.__ctor();
      _197_v86 = _nw20;
      _197_v86 = _197_v86;
      let _198_v87;
      let _199_v88;
      let _out0;
      let _out1;
      let _outcollector0 = (_106_v22).m0(_87_v4, _97_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _198_v87 = _out0;
      _199_v88 = _out1;
      process.stdout.write(_dafny.toString((_83_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_85_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_86_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_87_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_v5).equals(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_89_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))),new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_90_v7, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))),new BigNumber(636)), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))),new BigNumber(636))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_91_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_92_v9).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_94_v11).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(false,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_95_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(636),_dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_96_v13).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_97_globalState).f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_97_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_globalState.f4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0))),new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_97_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_97_globalState).f7).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(636),_dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_97_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_97_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_97_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_97_globalState).f11).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_97_globalState).f12).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_97_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_98_v14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_99_v16, _dafny.Seq.of(new BigNumber(636), new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_100_v17, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v18).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(636), new BigNumber(636))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(964),new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v20).equals(_dafny.MultiSet.fromElements(new BigNumber(636), new BigNumber(636), new BigNumber(636), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_105_v21).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_114_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(90)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_117_v28).dtor_cf38).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(90)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_v29).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(954),_module.D9.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber(90)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v30).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(636),_dafny.Seq.UnicodeFromString("fno")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_v52).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_v53).dtor_cf26).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_148_v54, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_149_v55, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("fno"), _dafny.Seq.UnicodeFromString("fno")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_150_v56).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_150_v56).dtor_cf2, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("fno"), _dafny.Seq.UnicodeFromString("fno")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_150_v56).dtor_cf3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_151_v57).dtor_cf5).dtor_cf5).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((((_151_v57).dtor_cf5).dtor_cf5).dtor_cf2, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("fno"), _dafny.Seq.UnicodeFromString("fno")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((((_151_v57).dtor_cf5).dtor_cf5).dtor_cf3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v58).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_v59).dtor_cf39));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_v59).dtor_cf40));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_167_v59).dtor_cf41));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v60)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_170_v62).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_171_v63).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_172_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_198_v87));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_199_v88));
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
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D0(3);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + this.cf3.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.Seq.of(), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC5(cf7, cf8, cf9) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC4(cf6) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ", " + this.cf8.toVerbatimString(true) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false, _dafny.Seq.UnicodeFromString(""), _dafny.Seq.of());
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
    static create_DC7(cf11, cf12, cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC8(cf15) {
      let $dt = new D2(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false, _dafny.Map.Empty, _dafny.ZERO, false);
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
    static create_DC9(cf16) {
      let $dt = new D3(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf16) + ")";
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D4(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC10(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf17) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11();
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
    static create_DC14(cf20, cf21, cf22) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D5(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(false, false, []);
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
    static create_DC16(cf24) {
      let $dt = new D6(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf23) {
      let $dt = new D6(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC17(cf25) {
      let $dt = new D6(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(_dafny.ZERO);
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
    static create_DC19(cf27, cf28, cf29, cf30) {
      let $dt = new D7(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC20(cf31, cf32) {
      let $dt = new D7(1);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC18(cf26) {
      let $dt = new D7(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC21(cf33) {
      let $dt = new D7(3);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + this.cf31.toVerbatimString(true) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29 && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(_module.D1.Default(), _dafny.ZERO, false, []);
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
    static create_DC23(cf35, cf36, cf37) {
      let $dt = new D8(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC24() {
      let $dt = new D8(1);
      return $dt;
    }
    static create_DC22(cf34) {
      let $dt = new D8(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf35 === other.cf35 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(false, false, _dafny.ZERO);
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
    static create_DC26(cf39, cf40, cf41) {
      let $dt = new D9(0);
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC27(cf42, cf43, cf44) {
      let $dt = new D9(1);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC25(cf38) {
      let $dt = new D9(2);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf39 === other.cf39 && _dafny.areEqual(this.cf40, other.cf40) && this.cf41 === other.cf41;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf42 === other.cf42 && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC26(false, new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC28(cf45) {
      let $dt = new D10(0);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45;
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf47) {
      let $dt = new D11(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC29(cf46) {
      let $dt = new D11(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf46) + ")";
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
      return _module.D11.create_DC30(false);
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
    static create_DC32(cf49, cf50) {
      let $dt = new D12(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC31(cf48) {
      let $dt = new D12(1);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf49 === other.cf49 && this.cf50 === other.cf50;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC32(false, false);
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
    static create_DC34(cf52) {
      let $dt = new D13(0);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC33(cf51) {
      let $dt = new D13(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC34(_dafny.ZERO);
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
    static create_DC36(cf54) {
      let $dt = new D14(0);
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC35(cf53) {
      let $dt = new D14(1);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC37(cf55) {
      let $dt = new D14(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC36(_dafny.ZERO);
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
    static create_DC38(cf56) {
      let $dt = new D15(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56);
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
    static create_DC40(cf58, cf59, cf60, cf61) {
      let $dt = new D16(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC39(cf57) {
      let $dt = new D16(1);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf58 === other.cf58 && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60) && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC40(false, false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
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
      this.f2 = [];
      this.f3 = false;
      this.f4 = _dafny.Map.Empty;
      this.f5 = false;
      this.f9 = false;
      this.f13 = false;
      this._f0 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f1 = [];
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.Map.Empty;
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f10 = false;
      this._f11 = _dafny.Set.Empty;
      this._f12 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f18 = _dafny.MultiSet.Empty;
      this._f19 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f18, f19) {
      let _this = this;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      return;
    }
    fm21(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((_this).f19) {
        return _dafny.Seq.of(!((_this).f19));
      } else {
        return _dafny.Seq.of(false, (_this).f19);
      }
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    m8(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _200_v0;
      _200_v0 = false;
      let _201_v1;
      _201_v1 = new BigNumber(-189);
      let _202_v2;
      _202_v2 = _module.D0.create_DC2(_200_v0);
      let _203_v3;
      _203_v3 = _dafny.Seq.of(_202_v2, _module.D0.create_DC2(_200_v0), _202_v2, _202_v2, _202_v2);
      (globalState).f13 = ((_200_v0) ? (false) : (((_dafny.ZERO).minus(_201_v1)).isLessThanOrEqualTo(new BigNumber((_203_v3).length))));
      let _204_v4;
      _204_v4 = _dafny.Map.Empty.slice().updateUnsafe(_200_v0,_200_v0);
      _204_v4 = (_204_v4).update((new BigNumber(-972)).isLessThanOrEqualTo(_201_v1), _200_v0);
      if (!(_200_v0)) {
        let _205_v5;
        let _init0 = function (_206_i0) {
          return _module.__default.safeDivisionInt(_206_i0, new BigNumber(-167));
        };
        let _nw21 = Array((new BigNumber(4)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw21.length); _i0_0++) {
          _nw21[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _205_v5 = _nw21;
        let _207_v6;
        _207_v6 = _dafny.Map.Empty.slice().updateUnsafe(_201_v1,_205_v5);
        let _208_v7;
        _208_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(_201_v1, _200_v0, globalState),(((_207_v6).contains(_201_v1)) ? ((_207_v6).get(_201_v1)) : (_205_v5)));
        let _209_v8;
        _209_v8 = _module.D1.create_DC4(_208_v7);
        _208_v7 = (_209_v8).dtor_cf6;
        let _210_v9;
        _210_v9 = _dafny.Seq.UnicodeFromString("a");
        _210_v9 = _210_v9;
        let _211_v10;
        let _init1 = ((_212_v0, _213_v4) => function (_214_i1) {
          return ((((_213_v4).contains(_212_v0)) ? ((_213_v4).get(_212_v0)) : (true))) === (_212_v0);
        })(_200_v0, _204_v4);
        let _nw22 = Array((new BigNumber(6)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw22.length); _i0_1++) {
          _nw22[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _211_v10 = _nw22;
        let _215_v11;
        _215_v11 = _dafny.Seq.of(_200_v0);
        let _216_v12;
        _216_v12 = _dafny.Seq.of(_200_v0, _200_v0, (_215_v11)[_module.__default.safeIndex(_201_v1, new BigNumber((_215_v11).length))]);
        let _index21 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_211_v10).length));
        (_211_v10)[_index21] = (_216_v12)[_module.__default.safeIndex(_201_v1, new BigNumber((_216_v12).length))];
        let _index22 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_205_v5).length));
        (_205_v5)[_index22] = ((_dafny.ZERO).minus(new BigNumber((_module.__default.fm20(new BigNumber((_210_v9).length), _200_v0, _200_v0, _201_v1, globalState)).length))).minus(_201_v1);
        let _index23 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_211_v10).length));
        let _index24 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_205_v5).length));
        let _rhs38 = !(false);
        let _rhs39 = ((_201_v1).plus(_201_v1)).minus((_201_v1).plus(_201_v1));
        let _rhs40 = !(((_dafny.ZERO).minus(new BigNumber(-89))).multipliedBy(_201_v1)).isEqualTo(_module.__default.safeDivisionInt(_201_v1, _201_v1));
        let _lhs27 = _211_v10;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_211_v10).length));
        let _lhs29 = _205_v5;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_205_v5).length));
        let _lhs31 = globalState;
        _lhs27[_lhs28] = _rhs38;
        _lhs29[_lhs30] = _rhs39;
        _lhs31.f3 = _rhs40;
        (globalState).f3 = true;
        let _217_v13;
        _217_v13 = _dafny.MultiSet.fromElements(!(_200_v0), _200_v0, (_211_v10)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_211_v10).length))]);
        r0 = _module.__default.safeDivisionInt((((_217_v13).contains(_200_v0)) ? ((_217_v13).get(_200_v0)) : ((_205_v5)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_205_v5).length))])), (_205_v5)[_module.__default.safeIndex(new BigNumber(819), new BigNumber((_205_v5).length))]);
      } else {
        let _218_v14;
        let _nw23 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _218_v14 = _nw23;
        let _219_v15;
        _219_v15 = _dafny.Seq.of(_201_v1);
        let _index25 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length));
        (_218_v14)[_index25] = new BigNumber((_dafny.Seq.update(_219_v15, _module.__default.safeIndex(_201_v1, new BigNumber((_219_v15).length)), _201_v1)).length);
        let _index26 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length));
        (_218_v14)[_index26] = (_dafny.ZERO).minus(_201_v1);
        let _220_v16;
        let _nw24 = Array((new BigNumber(20)).toNumber()).fill([]);
        _220_v16 = _nw24;
        let _221_v17;
        let _nw25 = Array((new BigNumber(18)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _221_v17 = _nw25;
        let _222_v18;
        _222_v18 = _dafny.Seq.of(_221_v17, _221_v17, _221_v17, _221_v17, _221_v17);
        let _223_v19;
        let _nw26 = Array((new BigNumber(28)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _221_v17;
        _nw26[(_dafny.ONE).toNumber()] = _221_v17;
        _nw26[(new BigNumber(2)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(3)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(4)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(5)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(6)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(7)).toNumber()] = (_222_v18)[_module.__default.safeIndex(new BigNumber(-11), new BigNumber((_222_v18).length))];
        _nw26[(new BigNumber(8)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(9)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(10)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(11)).toNumber()] = ((_200_v0) ? (_221_v17) : (_221_v17));
        _nw26[(new BigNumber(12)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(13)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(14)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(15)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(16)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(17)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(18)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(19)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(20)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(21)).toNumber()] = ((_200_v0) ? (_221_v17) : (_221_v17));
        _nw26[(new BigNumber(22)).toNumber()] = ((_200_v0) ? (_221_v17) : (_221_v17));
        _nw26[(new BigNumber(23)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(24)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(25)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(26)).toNumber()] = _221_v17;
        _nw26[(new BigNumber(27)).toNumber()] = _221_v17;
        _223_v19 = _nw26;
        let _index27 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_223_v19).length));
        (_223_v19)[_index27] = _221_v17;
        let _index28 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_223_v19).length));
        let _rhs41 = _220_v16;
        let _rhs42 = _221_v17;
        let _rhs43 = _200_v0;
        let _lhs32 = _223_v19;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_223_v19).length));
        _220_v16 = _rhs41;
        _lhs32[_lhs33] = _rhs42;
        _200_v0 = _rhs43;
        let _224_v20;
        _224_v20 = _dafny.Seq.UnicodeFromString("ddwvkjgl");
        if (!_dafny.areEqual(_224_v20, _224_v20)) {
          let _225_v21;
          _225_v21 = _dafny.MultiSet.fromElements(_201_v1, _201_v1);
          let _226_v22;
          _226_v22 = _dafny.Map.Empty.slice().updateUnsafe((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))],(_225_v21).Intersect(_225_v21));
          _226_v22 = (_226_v22).update(((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))]).plus(_201_v1), (_dafny.MultiSet.FromArray(_219_v15)).Intersect((_225_v21).update((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))], _module.__default.abs(_201_v1))));
          let _227_v23;
          _227_v23 = _dafny.Seq.of(_224_v20, _224_v20);
          let _228_v24;
          _228_v24 = _dafny.Seq.of((_227_v23)[_module.__default.safeIndex(_201_v1, new BigNumber((_227_v23).length))], _224_v20);
          let _229_v25;
          _229_v25 = _dafny.Seq.of(_200_v0);
          let _230_v26;
          _230_v26 = _module.D1.create_DC5(_200_v0, _224_v20, _229_v25);
          let _231_v27;
          _231_v27 = _dafny.Map.Empty.slice().updateUnsafe((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))],(_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))]);
          let _232_v28;
          _232_v28 = _dafny.Map.Empty.slice().updateUnsafe(_204_v4,_200_v0);
          let _233_v29;
          _233_v29 = _dafny.Set.fromElements(_200_v0);
          let _234_v30;
          _234_v30 = _dafny.MultiSet.fromElements(_233_v29);
          let _235_v31;
          let _nw27 = Array((new BigNumber(27)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _200_v0;
          _nw27[(_dafny.ONE).toNumber()] = (false) && (_200_v0);
          _nw27[(new BigNumber(2)).toNumber()] = true;
          _nw27[(new BigNumber(3)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(4)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(5)).toNumber()] = !(_200_v0);
          _nw27[(new BigNumber(6)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(7)).toNumber()] = _dafny.Seq.IsPrefixOf(_224_v20, _dafny.Seq.UnicodeFromString("rexhpd"));
          _nw27[(new BigNumber(8)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(9)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(10)).toNumber()] = !(_200_v0);
          _nw27[(new BigNumber(11)).toNumber()] = (new BigNumber(((_227_v23)[_module.__default.safeIndex((((_225_v21).contains((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))])) ? ((_225_v21).get((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))])) : ((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))])), new BigNumber((_227_v23).length))]).length)).isLessThan(_201_v1);
          _nw27[(new BigNumber(12)).toNumber()] = !(_200_v0);
          _nw27[(new BigNumber(13)).toNumber()] = true;
          _nw27[(new BigNumber(14)).toNumber()] = (_230_v26).dtor_cf7;
          _nw27[(new BigNumber(15)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(16)).toNumber()] = ((_231_v27).update((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))], (_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))])).equals(_231_v27);
          _nw27[(new BigNumber(17)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(18)).toNumber()] = (((((_232_v28).contains(_204_v4)) ? ((_232_v28).get(_204_v4)) : (_200_v0))) ? (true) : (!(!(_200_v0))));
          _nw27[(new BigNumber(19)).toNumber()] = true;
          _nw27[(new BigNumber(20)).toNumber()] = (_225_v21).IsProperSubsetOf(_225_v21);
          _nw27[(new BigNumber(21)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(22)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(23)).toNumber()] = ((((_234_v30).contains(_233_v29)) ? ((_234_v30).get(_233_v29)) : ((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))]))).isEqualTo(_201_v1);
          _nw27[(new BigNumber(24)).toNumber()] = _200_v0;
          _nw27[(new BigNumber(25)).toNumber()] = _dafny.areEqual(_229_v25, _229_v25);
          _nw27[(new BigNumber(26)).toNumber()] = _200_v0;
          _235_v31 = _nw27;
          let _index29 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_235_v31).length));
          (_235_v31)[_index29] = _dafny.Seq.IsPrefixOf(_224_v20, _224_v20);
          let _index30 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_235_v31).length));
          (_235_v31)[_index30] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("h"), _dafny.Seq.UnicodeFromString("kfkrjpla"));
          let _236_v32;
          _236_v32 = _module.D0.create_DC1(_200_v0, _227_v23, _224_v20);
          let _237_v33;
          _237_v33 = _dafny.MultiSet.fromElements(_236_v32, _236_v32, _236_v32, _236_v32);
          let _238_v34;
          let _nw28 = new _module.C0();
          _nw28.__ctor((_module.D2.create_DC6(_237_v33)).dtor_cf10, !(!(false)));
          _238_v34 = _nw28;
          _235_v31 = _235_v31;
          let _239_v35;
          let _init2 = ((_240_v20) => function (_241_i2) {
            return _dafny.Seq.Concat(_240_v20, _240_v20);
          })(_224_v20);
          let _nw29 = Array((new BigNumber(11)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw29.length); _i0_2++) {
            _nw29[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _239_v35 = _nw29;
          let _index31 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_239_v35).length));
          (_239_v35)[_index31] = _dafny.Seq.UnicodeFromString("mcihcac");
          let _242_v36;
          _242_v36 = _dafny.Map.Empty.slice().updateUnsafe((_238_v34).f19,(((_238_v34).f19) ? ((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))]) : (new BigNumber((_dafny.Set.fromElements(new BigNumber(584), (_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))])).length))));
          let _index32 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_239_v35).length));
          let _rhs44 = _dafny.Seq.Concat(_224_v20, _dafny.Seq.Concat(_224_v20, _224_v20));
          let _rhs45 = (_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))];
          let _rhs46 = _242_v36;
          let _lhs34 = _239_v35;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(664), new BigNumber((_239_v35).length));
          _lhs34[_lhs35] = _rhs44;
          r0 = _rhs45;
          _242_v36 = _rhs46;
        } else {
          (globalState).f3 = (_202_v2).dtor_cf4;
          let _243_v37;
          _243_v37 = new _dafny.CodePoint('s'.codePointAt(0));
          let _244_v38;
          let _nw30 = Array((new BigNumber(19)).toNumber()).fill(false);
          _244_v38 = _nw30;
          let _245_v39;
          _245_v39 = _dafny.MultiSet.fromElements(_201_v1, _201_v1, (_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))], (_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))]);
          let _246_v40;
          _246_v40 = _dafny.Seq.of(_224_v20);
          let _247_v41;
          _247_v41 = _module.D0.create_DC2(_200_v0);
          let _248_v42;
          _248_v42 = _module.D0.create_DC3(_247_v41);
          let _249_v43;
          _249_v43 = _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC1(_200_v0, _246_v40, _224_v20)).dtor_cf1,_248_v42);
          let _250_v44;
          _250_v44 = _module.D2.create_DC7(_200_v0, _249_v43, _201_v1, _200_v0);
          let _251_v45;
          _251_v45 = _dafny.Map.Empty.slice().updateUnsafe(_224_v20,_244_v38);
          let _252_v46;
          _252_v46 = _dafny.Set.fromElements((_dafny.ZERO).minus((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))]));
          let _index33 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length));
          let _rhs47 = _module.__default.fm22(_250_v44, _248_v42, _200_v0, globalState);
          let _rhs48 = _module.__default.fm22(_250_v44, _248_v42, _200_v0, globalState);
          let _rhs49 = _module.__default.fm23(_200_v0, _module.__default.fm24(globalState), _200_v0, (_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))], globalState);
          let _rhs50 = (((_251_v45).contains(_dafny.Seq.Concat(_224_v20, _224_v20))) ? ((_251_v45).get(_dafny.Seq.Concat(_224_v20, _224_v20))) : (_244_v38));
          let _rhs51 = (_245_v39).update((new BigNumber((_224_v20).length)).multipliedBy(_201_v1), _module.__default.abs(new BigNumber(((_252_v46).Difference(_252_v46)).length)));
          let _lhs36 = _218_v14;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length));
          _lhs36[_lhs37] = _rhs47;
          r0 = _rhs48;
          _243_v37 = _rhs49;
          _244_v38 = _rhs50;
          _245_v39 = _rhs51;
          let _253_v47;
          _253_v47 = _module.D0.create_DC1(_200_v0, _246_v40, _module.__default.fm20(_201_v1, !(_200_v0), _200_v0, new BigNumber((_224_v20).length), globalState));
          let _254_v48;
          let _nw31 = new _module.C0();
          _nw31.__ctor(_dafny.MultiSet.FromArray(_dafny.Seq.of(_253_v47, _253_v47)), _200_v0);
          _254_v48 = _nw31;
          let _index34 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length));
          (_218_v14)[_index34] = (_dafny.ZERO).minus((_201_v1).plus(_module.__default.safeDivisionInt((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))], (_dafny.ZERO).minus(_201_v1))));
          let _255_v49;
          let _nw32 = new _module.C0();
          _nw32.__ctor(((_254_v48).f18).update(_253_v47, _module.__default.abs(new BigNumber((_252_v46).length))), !((_254_v48).f19));
          _255_v49 = _nw32;
        }
        (globalState).f5 = _200_v0;
        let _256_v50;
        _256_v50 = _module.D0.create_DC0((_218_v14)[_module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length))]);
        let _257_v51;
        _257_v51 = _module.D0.create_DC3(_256_v50);
        let _258_v52;
        _258_v52 = _dafny.Map.Empty.slice().updateUnsafe(_200_v0,_257_v51);
        let _259_v53;
        _259_v53 = _module.D2.create_DC7(_200_v0, _258_v52, _201_v1, false);
        let _index35 = _module.__default.safeIndex(new BigNumber(755), new BigNumber((_218_v14).length));
        (_218_v14)[_index35] = _module.__default.fm22(_259_v53, _module.D0.create_DC3(_256_v50), !(_200_v0), globalState);
      }
      let _260_v54;
      _260_v54 = _dafny.Map.Empty.slice().updateUnsafe(_201_v1,_201_v1);
      let _261_v55;
      _261_v55 = _dafny.Seq.UnicodeFromString("tlaqlophn");
      _260_v54 = (_260_v54).update(new BigNumber((_261_v55).length), _201_v1);
      let _262_v56;
      _262_v56 = _dafny.Map.Empty.slice().updateUnsafe(_200_v0,_201_v1);
      let _263_v57;
      _263_v57 = new _dafny.CodePoint('g'.codePointAt(0));
      let _264_v58;
      _264_v58 = _dafny.MultiSet.fromElements(_263_v57);
      r0 = (_201_v1).plus((((_262_v56).contains(_200_v0)) ? ((_262_v56).get(_200_v0)) : ((((_264_v58).contains(_263_v57)) ? ((_264_v58).get(_263_v57)) : (_201_v1)))));
      let _265_i3;
      _265_i3 = _dafny.ZERO;
      L1: {
        while (_200_v0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_265_i3)) {
              break L1;
            }
            _265_i3 = (_265_i3).plus(_dafny.ONE);
            if ((_module.__default.fm19(_201_v1, _200_v0, globalState)) && (_200_v0)) {
              let _266_v59;
              let _init3 = ((_267_v0) => function (_268_i4) {
                return (_268_i4).minus(new BigNumber((_dafny.Seq.of(_267_v0)).length));
              })(_200_v0);
              let _nw33 = Array((new BigNumber(20)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw33.length); _i0_3++) {
                _nw33[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _266_v59 = _nw33;
              let _index36 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length));
              (_266_v59)[_index36] = (_dafny.ZERO).minus(_201_v1);
              let _269_v60;
              _269_v60 = _dafny.Map.Empty.slice().updateUnsafe(_201_v1,_200_v0);
              let _270_v61;
              _270_v61 = _module.D0.create_DC3(_module.D0.create_DC0(new BigNumber((_269_v60).length)));
              let _271_v62;
              _271_v62 = _module.D0.create_DC3(_270_v61);
              let _272_v63;
              _272_v63 = _dafny.Map.Empty.slice().updateUnsafe(_200_v0,_271_v62);
              let _273_v64;
              _273_v64 = _module.D2.create_DC7(_200_v0, _272_v63, _201_v1, _200_v0);
              let _274_v65;
              _274_v65 = _dafny.Seq.of(_200_v0);
              let _index37 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length));
              let _rhs52 = _module.__default.safeModuloInt(_201_v1, new BigNumber((_module.__default.fm20(new BigNumber((_262_v56).length), _200_v0, _200_v0, _module.__default.fm22(_273_v64, _271_v62, _200_v0, globalState), globalState)).length));
              let _rhs53 = !(_200_v0) || (_200_v0);
              let _rhs54 = new BigNumber((_274_v65).length);
              let _lhs38 = _266_v59;
              let _lhs39 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length));
              let _lhs40 = globalState;
              _lhs38[_lhs39] = _rhs52;
              _lhs40.f5 = _rhs53;
              _201_v1 = _rhs54;
              let _275_v66;
              _275_v66 = _dafny.MultiSet.fromElements((_266_v59)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length))]);
              let _276_v67;
              _276_v67 = _dafny.Seq.of((_266_v59)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length))], new BigNumber((_dafny.Seq.of((_266_v59)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length))])).length));
              let _rhs55 = ((_module.__default.fm25(_200_v0, _276_v67, globalState)).Intersect(_275_v66)).Union((_dafny.MultiSet.fromElements(new BigNumber(-802), _201_v1, _201_v1)).Union(_275_v66));
              let _rhs56 = _200_v0;
              let _lhs41 = globalState;
              _275_v66 = _rhs55;
              _lhs41.f13 = _rhs56;
              let _277_v68;
              _277_v68 = _module.D1.create_DC5(_200_v0, _261_v55, _dafny.Seq.update(_274_v65, _module.__default.safeIndex((_dafny.ZERO).minus(_201_v1), new BigNumber((_274_v65).length)), (_274_v65)[_module.__default.safeIndex(_201_v1, new BigNumber((_274_v65).length))]));
              let _pat_let_tv6 = _261_v55;
              let _278_v69;
              let _nw34 = Array((new BigNumber(5)).toNumber());
              _nw34[(_dafny.ZERO).toNumber()] = _module.__default.fm26(_200_v0, _200_v0, globalState);
              _nw34[(_dafny.ONE).toNumber()] = _277_v68;
              _nw34[(new BigNumber(2)).toNumber()] = _277_v68;
              _nw34[(new BigNumber(3)).toNumber()] = _277_v68;
              _nw34[(new BigNumber(4)).toNumber()] = function (_pat_let2_0) {
                return function (_279_dt__update__tmp_h0) {
                  return function (_pat_let3_0) {
                    return function (_280_dt__update_hcf8_h0) {
                      return _module.D1.create_DC5((_279_dt__update__tmp_h0).dtor_cf7, _280_dt__update_hcf8_h0, (_279_dt__update__tmp_h0).dtor_cf9);
                    }(_pat_let3_0);
                  }(_pat_let_tv6);
                }(_pat_let2_0);
              }(_module.__default.fm26(_200_v0, _200_v0, globalState));
              _278_v69 = _nw34;
              let _index38 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_278_v69).length));
              (_278_v69)[_index38] = _277_v68;
              let _index39 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_278_v69).length));
              (_278_v69)[_index39] = _277_v68;
              let _281_v70;
              _281_v70 = _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_275_v66).cardinality())), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-342)), ((_282_v1) => function (_283_i5) {
                return _282_v1;
              })(_201_v1)));
              let _index40 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length));
              let _index41 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length));
              let _rhs57 = _module.__default.safeDivisionInt((_266_v59)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length))], new BigNumber((_dafny.Seq.Concat(_module.__default.fm20((_dafny.ZERO).minus(new BigNumber((_276_v67).length)), !(_200_v0), true, (_266_v59)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length))], globalState), _261_v55)).length));
              let _rhs58 = (_dafny.Set.fromElements(_276_v67)).IsProperSubsetOf(_281_v70);
              let _rhs59 = ((_dafny.ZERO).minus(((_200_v0) ? ((_266_v59)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length))]) : (new BigNumber(273))))).plus((_201_v1).minus(_201_v1));
              let _rhs60 = _204_v4;
              let _lhs42 = _266_v59;
              let _lhs43 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length));
              let _lhs44 = _266_v59;
              let _lhs45 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length));
              _lhs42[_lhs43] = _rhs57;
              _200_v0 = _rhs58;
              _lhs44[_lhs45] = _rhs59;
              _204_v4 = _rhs60;
              r0 = (_276_v67)[_module.__default.safeIndex((_266_v59)[_module.__default.safeIndex(new BigNumber(961), new BigNumber((_266_v59).length))], new BigNumber((_276_v67).length))];
            } else {
              let _284_v71;
              _284_v71 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_200_v0,_263_v57)).length));
              let _285_v72;
              _285_v72 = _dafny.Seq.of(new BigNumber((_261_v55).length), _201_v1);
              let _286_v73;
              _286_v73 = _dafny.MultiSet.fromElements(_285_v72);
              let _287_v74;
              _287_v74 = _dafny.Seq.of(_260_v54, _260_v54, _260_v54, _260_v54);
              let _288_v75;
              let _nw35 = Array((new BigNumber(14)).toNumber());
              _nw35[(_dafny.ZERO).toNumber()] = ((_284_v71).update(_201_v1, _module.__default.abs(_201_v1))).IsSubsetOf(_284_v71);
              _nw35[(_dafny.ONE).toNumber()] = _200_v0;
              _nw35[(new BigNumber(2)).toNumber()] = false;
              _nw35[(new BigNumber(3)).toNumber()] = _200_v0;
              _nw35[(new BigNumber(4)).toNumber()] = (_200_v0) || (_200_v0);
              _nw35[(new BigNumber(5)).toNumber()] = !(new BigNumber((_286_v73).cardinality())).isEqualTo(_201_v1);
              _nw35[(new BigNumber(6)).toNumber()] = _200_v0;
              _nw35[(new BigNumber(7)).toNumber()] = _200_v0;
              _nw35[(new BigNumber(8)).toNumber()] = _200_v0;
              _nw35[(new BigNumber(9)).toNumber()] = !(true) || (false);
              _nw35[(new BigNumber(10)).toNumber()] = (false) || (_200_v0);
              _nw35[(new BigNumber(11)).toNumber()] = !_dafny.areEqual(_287_v74, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_201_v1,_201_v1)));
              _nw35[(new BigNumber(12)).toNumber()] = _200_v0;
              _nw35[(new BigNumber(13)).toNumber()] = _200_v0;
              _288_v75 = _nw35;
              let _index42 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_288_v75).length));
              (_288_v75)[_index42] = _200_v0;
              let _index43 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_288_v75).length));
              (_288_v75)[_index43] = (_dafny.Seq.IsPrefixOf(_261_v55, _261_v55)) || (_200_v0);
              let _289_v76;
              _289_v76 = _module.D0.create_DC1((_288_v75)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_288_v75).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), ((_290_v55) => function (_291_i6) {
  return _290_v55;
})(_261_v55)), _261_v55);
              let _292_v77;
              let _nw36 = new _module.C0();
              _nw36.__ctor(_dafny.MultiSet.fromElements(_289_v76, _289_v76, _289_v76), _200_v0);
              _292_v77 = _nw36;
              let _293_v78;
              let _nw37 = Array((new BigNumber(17)).toNumber());
              _nw37[(_dafny.ZERO).toNumber()] = _292_v77;
              _nw37[(_dafny.ONE).toNumber()] = _292_v77;
              _nw37[(new BigNumber(2)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(3)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(4)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(5)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(6)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(7)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(8)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(9)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(10)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(11)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(12)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(13)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(14)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(15)).toNumber()] = _292_v77;
              _nw37[(new BigNumber(16)).toNumber()] = _292_v77;
              _293_v78 = _nw37;
              let _294_v79;
              _294_v79 = _dafny.Map.Empty.slice().updateUnsafe(_200_v0,(((_288_v75)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_288_v75).length))]) ? (_293_v78) : (_293_v78)));
              _294_v79 = (_294_v79).update(false, _293_v78);
              _261_v55 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("phrjej"), _261_v55);
              let _295_v80;
              _295_v80 = _dafny.Map.Empty.slice().updateUnsafe((_292_v77).f19,_262_v56);
              let _296_v81;
              _296_v81 = _dafny.Seq.of(true);
              _262_v56 = (((_200_v0) ? (_262_v56) : (_dafny.Map.Empty.slice().updateUnsafe((_292_v77).f19,_201_v1)))).Merge((((_295_v80).contains((_296_v81)[_module.__default.safeIndex(_201_v1, new BigNumber((_296_v81).length))])) ? ((_295_v80).get((_296_v81)[_module.__default.safeIndex(_201_v1, new BigNumber((_296_v81).length))])) : (_262_v56)));
              let _index44 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_288_v75).length));
              (_288_v75)[_index44] = (((_204_v4).contains((_288_v75)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_288_v75).length))])) ? ((_204_v4).get((_288_v75)[_module.__default.safeIndex(new BigNumber(103), new BigNumber((_288_v75).length))])) : (!((_292_v77).f19)));
            }
            let _297_v83;
            _297_v83 = _dafny.MultiSet.fromElements(_201_v1);
            let _298_v84;
            _298_v84 = _dafny.Seq.of(_201_v1, new BigNumber((function () {
              let _coll5 = new _dafny.Map();
              for (const _compr_5 of (_297_v83).Elements) {
                let _299_v82 = _compr_5;
                if ((_297_v83).contains(_299_v82)) {
                  _coll5.push([_module.__default.safeModuloInt(_299_v82, (((_262_v56).contains(_200_v0)) ? ((_262_v56).get(_200_v0)) : (new BigNumber((_261_v55).length)))),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_201_v1)).length),_200_v0)).length)]);
                }
              }
              return _coll5;
            }()).length), _201_v1);
            r0 = (_298_v84)[_module.__default.safeIndex((_298_v84)[_module.__default.safeIndex(_201_v1, new BigNumber((_298_v84).length))], new BigNumber((_298_v84).length))];
            let _300_v86;
            _300_v86 = _dafny.Map.Empty.slice().updateUnsafe(_201_v1,_261_v55);
            let _301_v88;
            let _nw38 = Array((new BigNumber(11)).toNumber());
            _nw38[(_dafny.ZERO).toNumber()] = (function () {
              let _coll6 = new _dafny.Map();
              for (const _compr_6 of _dafny.IntegerRange(new BigNumber(955), new BigNumber(953))) {
                let _302_v85 = _compr_6;
                if (((new BigNumber(955)).isLessThanOrEqualTo(_302_v85)) && ((_302_v85).isLessThan(new BigNumber(953)))) {
                  _coll6.push([_module.__default.safeDivisionInt(_302_v85, new BigNumber(-853)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(346)), function (_303_i7) {
                    return new _dafny.CodePoint('i'.codePointAt(0));
                  })]);
                }
              }
              return _coll6;
            }()).Merge(_300_v86);
            _nw38[(_dafny.ONE).toNumber()] = _300_v86;
            _nw38[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_201_v1,_261_v55)).Merge(_300_v86);
            _nw38[(new BigNumber(3)).toNumber()] = ((_300_v86).update(new BigNumber(418), _module.__default.fm20(new BigNumber((_260_v54).length), _200_v0, _200_v0, new BigNumber((_dafny.Seq.UnicodeFromString("gj")).length), globalState))).Merge(_300_v86);
            _nw38[(new BigNumber(4)).toNumber()] = (_300_v86).Merge(_dafny.Map.Empty.slice().updateUnsafe(_201_v1,_261_v55));
            _nw38[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_201_v1,_261_v55);
            _nw38[(new BigNumber(6)).toNumber()] = _300_v86;
            _nw38[(new BigNumber(7)).toNumber()] = _300_v86;
            _nw38[(new BigNumber(8)).toNumber()] = (_300_v86).Merge(_300_v86);
            _nw38[(new BigNumber(9)).toNumber()] = _300_v86;
            _nw38[(new BigNumber(10)).toNumber()] = (function () {
              let _coll7 = new _dafny.Map();
              for (const _compr_7 of _dafny.IntegerRange(new BigNumber(81), new BigNumber(133))) {
                let _304_v87 = _compr_7;
                if (((new BigNumber(81)).isLessThanOrEqualTo(_304_v87)) && ((_304_v87).isLessThan(new BigNumber(133)))) {
                  _coll7.push([(_304_v87).minus(new BigNumber((_264_v58).cardinality())),_261_v55]);
                }
              }
              return _coll7;
            }()).update(_201_v1, _261_v55);
            _301_v88 = _nw38;
            let _index45 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_301_v88).length));
            (_301_v88)[_index45] = (_300_v86).Merge(_300_v86);
            let _305_v89;
            _305_v89 = _dafny.Seq.of(_300_v86);
            let _index46 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_301_v88).length));
            (_301_v88)[_index46] = ((_300_v86).Merge((_305_v89)[_module.__default.safeIndex(_201_v1, new BigNumber((_305_v89).length))])).Merge(_300_v86);
            _201_v1 = _201_v1;
          }
        }
      }
      let _306_v90;
      _306_v90 = _dafny.Seq.of(_201_v1, _201_v1, _201_v1);
      let _307_v91;
      _307_v91 = _module.D1.create_DC5(_200_v0, _261_v55, _module.__default.fm27(new BigNumber((_306_v90).length), _201_v1, _201_v1, globalState));
      let _pat_let_tv7 = _201_v1;
      let _pat_let_tv8 = _201_v1;
      r0 = function (_source6) {
        if (_source6.is_DC5) {
          let _308___mcc_h0 = (_source6).cf7;
          let _309___mcc_h1 = (_source6).cf8;
          let _310___mcc_h2 = (_source6).cf9;
          let _311_cf9 = _310___mcc_h2;
          let _312_cf8 = _309___mcc_h1;
          let _313_cf7 = _308___mcc_h0;
          return _pat_let_tv7;
        } else {
          let _314___mcc_h3 = (_source6).cf6;
          let _315_cf6 = _314___mcc_h3;
          return _pat_let_tv8;
        }
      }(_307_v91);
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
    fm0(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(226);
    };
    fm1(p0, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("fbtqkyvh");
    };
    fm14(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(354)), function (_316_i1) {
        return new BigNumber(-772);
      }),_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false)).length))))).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(944)), function (_317_i0) {
        return (_module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).dtor_cf0;
      }));
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      (globalState).f13 = !(p2) || (!(p2) || (p2));
      r1 = _module.__default.safeModuloInt(new BigNumber((_module.__default.fm15(p2, p2, p2, globalState)).length), ((p2) ? ((_this).fm0(new BigNumber((p0).length), new _dafny.CodePoint('p'.codePointAt(0)), globalState)) : (p3)));
      let _318_v0;
      _318_v0 = _dafny.Set.fromElements(p3);
      let _hi2 = new BigNumber(((_318_v0).Intersect(_318_v0)).length);
      for (let _319_i0 = p3; _319_i0.isLessThan(_hi2); _319_i0 = _319_i0.plus(_dafny.ONE)) {
        let _320_v1;
        let _nw39 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _320_v1 = _nw39;
        _320_v1 = _320_v1;
        r1 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(p3, _319_i0), p3);
        r1 = _319_i0;
        let _321_v2;
        let _nw40 = Array((new BigNumber(17)).toNumber()).fill([]);
        _321_v2 = _nw40;
        let _322_v3;
        let _nw41 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
        _322_v3 = _nw41;
        let _index47 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_321_v2).length));
        (_321_v2)[_index47] = ((p2) ? (_322_v3) : (_322_v3));
        let _index48 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_321_v2).length));
        (_321_v2)[_index48] = _322_v3;
      }
      let _323_v4;
      _323_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _324_v5;
      _324_v5 = _dafny.Seq.of((_this).fm14(_dafny.Set.fromElements(p3, new BigNumber((_323_v4).length)), !(p2), p1, globalState), false, true);
      let _325_v6;
      _325_v6 = new _dafny.CodePoint('a'.codePointAt(0));
      if ((_324_v5)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mkjhatikf"), p0), _module.__default.safeIndex(new BigNumber(631), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mkjhatikf"), p0)).length)), _325_v6)).length), new BigNumber((_324_v5).length))]) {
        let _326_v7;
        _326_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
        let _327_v8;
        let _nw42 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _327_v8 = _nw42;
        let _328_v9;
        _328_v9 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber(-712));
        let _329_v10;
        _329_v10 = _dafny.Map.Empty.slice().updateUnsafe(_327_v8,new BigNumber((_328_v9).length));
        let _330_v11;
        let _nw43 = Array((new BigNumber(12)).toNumber());
        _nw43[(_dafny.ZERO).toNumber()] = new BigNumber(-737);
        _nw43[(_dafny.ONE).toNumber()] = (((_326_v7).contains(p2)) ? ((_326_v7).get(p2)) : (p3));
        _nw43[(new BigNumber(2)).toNumber()] = p3;
        _nw43[(new BigNumber(3)).toNumber()] = p3;
        _nw43[(new BigNumber(4)).toNumber()] = p3;
        _nw43[(new BigNumber(5)).toNumber()] = p3;
        _nw43[(new BigNumber(6)).toNumber()] = p3;
        _nw43[(new BigNumber(7)).toNumber()] = p3;
        _nw43[(new BigNumber(8)).toNumber()] = p3;
        _nw43[(new BigNumber(9)).toNumber()] = p3;
        _nw43[(new BigNumber(10)).toNumber()] = new BigNumber((_329_v10).length);
        _nw43[(new BigNumber(11)).toNumber()] = p3;
        _330_v11 = _nw43;
        let _331_v12;
        _331_v12 = _dafny.Map.Empty.slice().updateUnsafe(p2,_327_v8);
        let _332_v13;
        _332_v13 = _dafny.Map.Empty.slice().updateUnsafe(p3,_325_v6);
        let _333_v15;
        _333_v15 = _dafny.MultiSet.fromElements((((_331_v12).contains((_this).fm14(function () {
          let _coll9 = new _dafny.Set();
          for (const _compr_9 of (_332_v13).Keys.Elements) {
            let _337_v14 = _compr_9;
            if ((_332_v13).contains(_337_v14)) {
              _coll9.add((_337_v14).multipliedBy(new BigNumber(290)));
            }
          }
          return _coll9;
        }(), p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(900)), ((_338_v6) => function (_339_i1) {
          return _338_v6;
        })(_325_v6)), globalState))) ? ((_331_v12).get((_this).fm14(function () {
          let _coll8 = new _dafny.Set();
          for (const _compr_8 of (_332_v13).Keys.Elements) {
            let _334_v14 = _compr_8;
            if ((_332_v13).contains(_334_v14)) {
              _coll8.add((_334_v14).multipliedBy(new BigNumber(290)));
            }
          }
          return _coll8;
        }(), p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(900)), ((_335_v6) => function (_336_i1) {
          return _335_v6;
        })(_325_v6)), globalState))) : (_330_v11)), _330_v11, _327_v8, _330_v11, _330_v11);
        let _340_v16;
        _340_v16 = _dafny.MultiSet.fromElements(p3);
        let _rhs61 = (_333_v15).Union(_333_v15);
        let _rhs62 = (_340_v16).Difference(_module.__default.fm16(globalState));
        _333_v15 = _rhs61;
        _340_v16 = _rhs62;
        let _341_v17;
        _341_v17 = _module.D0.create_DC3(_module.D0.create_DC0((_this).fm0(new BigNumber(144), new _dafny.CodePoint('w'.codePointAt(0)), globalState)));
        let _342_v18;
        _342_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p3, new BigNumber(698)),_341_v17);
        let _343_v20;
        _343_v20 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        _342_v18 = (_342_v18).Merge(function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of (_343_v20).Keys.Elements) {
            let _344_v19 = _compr_10;
            if ((_343_v20).contains(_344_v19)) {
              _coll10.push([_module.__default.safeDivisionInt(_344_v19, p3),_341_v17]);
            }
          }
          return _coll10;
        }());
        (globalState).f9 = p2;
        let _345_v21;
        let _nw44 = Array((new BigNumber(17)).toNumber()).fill(_module.D0.Default());
        _345_v21 = _nw44;
        let _index49 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_345_v21).length));
        (_345_v21)[_index49] = _module.__default.fm17(p2, globalState);
        let _346_v22;
        _346_v22 = _dafny.Map.Empty.slice().updateUnsafe(p2,_318_v0);
        let _347_v23;
        _347_v23 = _dafny.Seq.of(p3);
        let _348_v25;
        _348_v25 = _dafny.Seq.of((_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_346_v22).contains(p2)) ? ((_346_v22).get(p2)) : (function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_347_v23).Elements) {
            let _349_v24 = _compr_11;
            if (_dafny.Seq.contains(_347_v23, _349_v24)) {
              _coll11.add((_349_v24).plus(new BigNumber((_dafny.Seq.of(false)).length)));
            }
          }
          return _coll11;
        }()))).length),p3), globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-726)), ((_350_v6) => function (_351_i2) {
          return _350_v6;
        })(_325_v6)), p1);
        let _352_v26;
        _352_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm0(p3, _325_v6, globalState),_348_v25);
        let _353_v27;
        _353_v27 = _module.D0.create_DC1(p2, (((_352_v26).contains(p3)) ? ((_352_v26).get(p3)) : (_dafny.Seq.of(p1))), _dafny.Seq.Concat(p0, p0));
        let _index50 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_345_v21).length));
        (_345_v21)[_index50] = _353_v27;
        let _354_v28;
        let _init4 = ((_355_p2) => function (_356_i3) {
          return _355_p2;
        })(p2);
        let _nw45 = Array((new BigNumber(26)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw45.length); _i0_4++) {
          _nw45[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _354_v28 = _nw45;
        let _index51 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_354_v28).length));
        (_354_v28)[_index51] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(65)), function (_357_i4) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), p1);
        let _index52 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_354_v28).length));
        (_354_v28)[_index52] = !((p2) || (p2)) || ((p3).isLessThan(new BigNumber((_324_v5).length)));
      } else {
        let _rhs63 = p3;
        let _rhs64 = false;
        let _lhs46 = globalState;
        r1 = _rhs63;
        _lhs46.f3 = _rhs64;
        (globalState).f3 = p2;
        r1 = p3;
        let _358_v29;
        _358_v29 = _module.D0.create_DC0(p3);
        let _359_v30;
        _359_v30 = _dafny.Seq.of(_358_v29);
        let _360_v31;
        _360_v31 = _dafny.Seq.of(_359_v30);
        let _361_v32;
        _361_v32 = _dafny.MultiSet.fromElements(p3, p3);
        let _362_v33;
        _362_v33 = _dafny.Seq.of(_361_v32);
        _360_v31 = _module.__default.fm18(_325_v6, new BigNumber(((_362_v33)[_module.__default.safeIndex(p3, new BigNumber((_362_v33).length))]).cardinality()), globalState);
        r1 = p3;
      }
      let _363_v34;
      _363_v34 = _dafny.MultiSet.fromElements(_318_v0, _318_v0);
      let _hi3 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,(((_363_v34).contains(_318_v0)) ? ((_363_v34).get(_318_v0)) : (p3)))).length))).multipliedBy((_dafny.ZERO).minus(p3));
      for (let _364_i5 = p3; _364_i5.isLessThan(_hi3); _364_i5 = _364_i5.plus(_dafny.ONE)) {
        _325_v6 = _325_v6;
        _325_v6 = new _dafny.CodePoint('c'.codePointAt(0));
        let _365_v35;
        _365_v35 = _dafny.Seq.of(_364_i5);
        let _366_v36;
        _366_v36 = _dafny.Map.Empty.slice().updateUnsafe(_364_i5,new BigNumber((_365_v35).length));
        let _367_v37;
        _367_v37 = _dafny.Map.Empty.slice().updateUnsafe(_364_i5,new BigNumber((_366_v36).length));
        let _368_v38;
        _368_v38 = _dafny.Seq.of(p0, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(29)), ((_369_v6) => function (_370_i6) {
          return _369_v6;
        })(_325_v6)), (_this).fm1(_367_v37, globalState));
        let _371_v39;
        _371_v39 = _module.D0.create_DC1(p2, _368_v38, _dafny.Seq.UnicodeFromString("unnkkr"));
        let _372_v40;
        _372_v40 = _dafny.MultiSet.fromElements(_371_v39);
        let _373_v41;
        let _nw46 = new _module.C0();
        _nw46.__ctor((_372_v40).update(_371_v39, _module.__default.abs(_364_i5)), p2);
        _373_v41 = _nw46;
        let _374_v42;
        _374_v42 = _dafny.Seq.UnicodeFromString("qumwtaj");
        _374_v42 = _dafny.Seq.UnicodeFromString("juj");
      }
      let _hi4 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p3), p3);
      for (let _375_i7 = p3; _375_i7.isLessThan(_hi4); _375_i7 = _375_i7.plus(_dafny.ONE)) {
        let _376_v43;
        _376_v43 = _dafny.Seq.of(p3, _375_i7);
        let _377_v44;
        let _nw47 = Array((new BigNumber(4)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = p2;
        _nw47[(_dafny.ONE).toNumber()] = !(p2);
        _nw47[(new BigNumber(2)).toNumber()] = p2;
        _nw47[(new BigNumber(3)).toNumber()] = p2;
        _377_v44 = _nw47;
        let _378_v45;
        _378_v45 = _dafny.Map.Empty.slice().updateUnsafe((_376_v43)[_module.__default.safeIndex((_dafny.ZERO).minus(p3), new BigNumber((_376_v43).length))],_377_v44);
        let _379_v46;
        _379_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_376_v43).length),_324_v5);
        _378_v45 = (_378_v45).update(new BigNumber(((((_379_v46).contains(_375_i7)) ? ((_379_v46).get(_375_i7)) : (_324_v5))).length), _377_v44);
        _325_v6 = _325_v6;
        r1 = (_this).fm0(_375_i7, _325_v6, globalState);
        let _380_v47;
        _380_v47 = _module.D1.create_DC5(p2, p0, _dafny.Seq.Concat(_dafny.Seq.of(p2, p2), _324_v5));
        _380_v47 = _module.D1.create_DC5(p2, _dafny.Seq.Concat(p0, p0), _dafny.Seq.Concat(_324_v5, _324_v5));
      }
      let _381_v48;
      _381_v48 = _dafny.Map.Empty.slice().updateUnsafe(p3,!(p2) || (true));
      let _382_v49;
      _382_v49 = _dafny.Seq.of(p3);
      r0 = !((((_381_v48).contains((_382_v49)[_module.__default.safeIndex(p3, new BigNumber((_382_v49).length))])) ? ((_381_v48).get((_382_v49)[_module.__default.safeIndex(p3, new BigNumber((_382_v49).length))])) : (!(p2))));
      r1 = p3;
      r2 = p2;
      return [r0, r1, r2];
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.Seq.of();
      let r2 = false;
      let _383_v0;
      _383_v0 = new _dafny.CodePoint('b'.codePointAt(0));
      let _384_v1;
      _384_v1 = _dafny.MultiSet.fromElements(_383_v0, new _dafny.CodePoint('b'.codePointAt(0)), _383_v0, _383_v0);
      let _385_v2;
      _385_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_384_v1).cardinality()));
      let _386_v3;
      _386_v3 = _dafny.MultiSet.fromElements(_385_v2);
      let _387_v4;
      _387_v4 = _dafny.MultiSet.fromElements(!(p3), p3);
      let _388_v5;
      _388_v5 = _dafny.Seq.of(true);
      let _389_v6;
      let _nw48 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _389_v6 = _nw48;
      let _390_v7;
      _390_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_388_v5).length),_389_v6);
      let _391_v8;
      _391_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),p2);
      let _392_v9;
      _392_v9 = _dafny.Seq.UnicodeFromString("uwerfkjaf");
      let _393_v10;
      _393_v10 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_392_v9).length))),p3);
      let _394_v11;
      let _nw49 = Array((new BigNumber(25)).toNumber());
      _nw49[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw49[(_dafny.ONE).toNumber()] = (p2).minus(p2);
      _nw49[(new BigNumber(2)).toNumber()] = ((p3) ? (new BigNumber(435)) : (p1));
      _nw49[(new BigNumber(3)).toNumber()] = p2;
      _nw49[(new BigNumber(4)).toNumber()] = (p1).plus(new BigNumber(766));
      _nw49[(new BigNumber(5)).toNumber()] = new BigNumber((_386_v3).cardinality());
      _nw49[(new BigNumber(6)).toNumber()] = p2;
      _nw49[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_387_v4).cardinality()), p2);
      _nw49[(new BigNumber(8)).toNumber()] = (new BigNumber((_388_v5).length)).minus(new BigNumber((_390_v7).length));
      _nw49[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw49[(new BigNumber(10)).toNumber()] = (((_391_v8).contains((_dafny.ZERO).minus(p1))) ? ((_391_v8).get((_dafny.ZERO).minus(p1))) : (new BigNumber(970)));
      _nw49[(new BigNumber(11)).toNumber()] = p1;
      _nw49[(new BigNumber(12)).toNumber()] = (p2).multipliedBy(p1);
      _nw49[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_392_v9, _392_v9)).length);
      _nw49[(new BigNumber(14)).toNumber()] = new BigNumber(-260);
      _nw49[(new BigNumber(15)).toNumber()] = p2;
      _nw49[(new BigNumber(16)).toNumber()] = p2;
      _nw49[(new BigNumber(17)).toNumber()] = (((_385_v2).contains(p3)) ? ((_385_v2).get(p3)) : (p1));
      _nw49[(new BigNumber(18)).toNumber()] = p2;
      _nw49[(new BigNumber(19)).toNumber()] = new BigNumber(-272);
      _nw49[(new BigNumber(20)).toNumber()] = p2;
      _nw49[(new BigNumber(21)).toNumber()] = new BigNumber(156);
      _nw49[(new BigNumber(22)).toNumber()] = p1;
      _nw49[(new BigNumber(23)).toNumber()] = new BigNumber((_393_v10).length);
      _nw49[(new BigNumber(24)).toNumber()] = p2;
      _394_v11 = _nw49;
      let _init5 = ((_395_p3, _396_p1) => function (_397_i0) {
        return (_397_i0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_395_p3,_396_p1)).length));
      })(p3, p1);
      let _nw50 = Array((new BigNumber(8)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw50.length); _i0_5++) {
        _nw50[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _394_v11 = _nw50;
      let _398_i1;
      _398_i1 = _dafny.ZERO;
      L2: {
        while (p3) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_398_i1)) {
              break L2;
            }
            _398_i1 = (_398_i1).plus(_dafny.ONE);
            let _399_v12;
            _399_v12 = _dafny.Set.fromElements(p1);
            let _400_v13;
            _400_v13 = _dafny.Map.Empty.slice().updateUnsafe(_399_v12,_392_v9);
            let _401_v14;
            _401_v14 = _dafny.Set.fromElements(((false) ? (p1) : (new BigNumber(((((_400_v13).contains(_399_v12)) ? ((_400_v13).get(_399_v12)) : (_dafny.Seq.UnicodeFromString("j")))).length))), p1);
            _399_v12 = _399_v12;
            let _402_v15;
            _402_v15 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm28(globalState));
            let _403_v16;
            _403_v16 = _module.D2.create_DC7(p3, (_402_v15).Merge(_402_v15), p1, p3);
            let _source7 = _403_v16;
            if (_source7.is_DC7) {
              let _404___mcc_h0 = (_source7).cf11;
              let _405___mcc_h1 = (_source7).cf12;
              let _406___mcc_h2 = (_source7).cf13;
              let _407___mcc_h3 = (_source7).cf14;
              let _408_cf14 = _407___mcc_h3;
              let _409_cf13 = _406___mcc_h2;
              let _410_cf12 = _405___mcc_h1;
              let _411_cf11 = _404___mcc_h0;
              _409_cf13 = new BigNumber(655);
              let _412_v17;
              _412_v17 = _dafny.Seq.of(_392_v9, _392_v9);
              let _413_v18;
              _413_v18 = _module.D0.create_DC1(p3, _412_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(574)), ((_414_v0) => function (_415_i2) {
  return _414_v0;
})(_383_v0)));
              let _416_v19;
              _416_v19 = _dafny.MultiSet.fromElements(_413_v18, _413_v18);
              let _417_v20;
              let _nw51 = new _module.C0();
              _nw51.__ctor(_416_v19, _411_cf11);
              _417_v20 = _nw51;
              let _418_v21;
              _418_v21 = _dafny.Map.Empty.slice().updateUnsafe(_411_cf11,(_417_v20).f19);
              (globalState).f5 = (((_418_v21).contains(p3)) ? ((_418_v21).get(p3)) : (_411_cf11));
              let _419_v22;
              _419_v22 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
              let _420_v23;
              _420_v23 = _module.D0.create_DC2(true);
              let _421_v24;
              let _nw52 = Array((new BigNumber(26)).toNumber());
              _nw52[(_dafny.ZERO).toNumber()] = !(new BigNumber((_392_v9).length)).isEqualTo(_409_cf13);
              _nw52[(_dafny.ONE).toNumber()] = true;
              _nw52[(new BigNumber(2)).toNumber()] = ((_417_v20).f19) === (false);
              _nw52[(new BigNumber(3)).toNumber()] = _411_cf11;
              _nw52[(new BigNumber(4)).toNumber()] = (_393_v10).equals(_393_v10);
              _nw52[(new BigNumber(5)).toNumber()] = !((_417_v20).f19);
              _nw52[(new BigNumber(6)).toNumber()] = _408_cf14;
              _nw52[(new BigNumber(7)).toNumber()] = _408_cf14;
              _nw52[(new BigNumber(8)).toNumber()] = (_409_cf13).isLessThan(new BigNumber((p0).length));
              _nw52[(new BigNumber(9)).toNumber()] = !(_408_cf14);
              _nw52[(new BigNumber(10)).toNumber()] = p3;
              _nw52[(new BigNumber(11)).toNumber()] = _411_cf11;
              _nw52[(new BigNumber(12)).toNumber()] = (p0).IsDisjointFrom((((_419_v22).contains(_408_cf14)) ? ((_419_v22).get(_408_cf14)) : (p0)));
              _nw52[(new BigNumber(13)).toNumber()] = p3;
              _nw52[(new BigNumber(14)).toNumber()] = (_417_v20).f19;
              _nw52[(new BigNumber(15)).toNumber()] = (_411_cf11) === (_411_cf11);
              _nw52[(new BigNumber(16)).toNumber()] = (_417_v20).f19;
              _nw52[(new BigNumber(17)).toNumber()] = (!((_417_v20).f19)) === (_408_cf14);
              _nw52[(new BigNumber(18)).toNumber()] = _408_cf14;
              _nw52[(new BigNumber(19)).toNumber()] = p3;
              _nw52[(new BigNumber(20)).toNumber()] = _module.__default.fm19(_409_cf13, _411_cf11, globalState);
              _nw52[(new BigNumber(21)).toNumber()] = (_420_v23).dtor_cf4;
              _nw52[(new BigNumber(22)).toNumber()] = !(p3);
              _nw52[(new BigNumber(23)).toNumber()] = false;
              _nw52[(new BigNumber(24)).toNumber()] = !((_399_v12).IsSubsetOf(_401_v14));
              _nw52[(new BigNumber(25)).toNumber()] = (_417_v20).f19;
              _421_v24 = _nw52;
              let _index53 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_421_v24).length));
              (_421_v24)[_index53] = (_417_v20).f19;
              let _422_v25;
              _422_v25 = _dafny.Seq.of(p2, new BigNumber((_392_v9).length), p1, p2, _409_cf13);
              let _423_v26;
              _423_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(990),_392_v9);
              let _index54 = _module.__default.safeIndex(new BigNumber(271), new BigNumber((_421_v24).length));
              (_421_v24)[_index54] = (_this).fm14(_399_v12, _dafny.Seq.contains(_422_v25, new BigNumber((_392_v9).length)), (((_423_v26).contains(p1)) ? ((_423_v26).get(p1)) : (_392_v9)), globalState);
            } else if (_source7.is_DC6) {
              let _424___mcc_h4 = (_source7).cf10;
              let _425_cf10 = _424___mcc_h4;
              let _426_v27;
              _426_v27 = _dafny.Seq.of(new BigNumber(-738));
              let _index55 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_389_v6).length));
              (_389_v6)[_index55] = (p2).multipliedBy((_426_v27)[_module.__default.safeIndex(p2, new BigNumber((_426_v27).length))]);
              let _427_v28;
              _427_v28 = _module.D0.create_DC0(p1);
              let _index56 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_389_v6).length));
              let _rhs65 = _426_v27;
              let _rhs66 = !(p2).isEqualTo(p1);
              let _rhs67 = _module.__default.fm22(_403_v16, _module.D0.create_DC3(_427_v28), p3, globalState);
              let _lhs47 = _389_v6;
              let _lhs48 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_389_v6).length));
              _426_v27 = _rhs65;
              r2 = _rhs66;
              _lhs47[_lhs48] = _rhs67;
              let _428_v29;
              let _429_v30;
              let _430_v31;
              let _out2;
              let _out3;
              let _out4;
              let _outcollector1 = (_this).m6(_392_v9, _392_v9, p3, p2, globalState);
              _out2 = _outcollector1[0];
              _out3 = _outcollector1[1];
              _out4 = _outcollector1[2];
              _428_v29 = _out2;
              _429_v30 = _out3;
              _430_v31 = _out4;
              let _431_v32;
              _431_v32 = _module.D0.create_DC3(_module.D0.create_DC0((_389_v6)[_module.__default.safeIndex(new BigNumber(326), new BigNumber((_389_v6).length))]));
              let _pat_let_tv9 = _427_v28;
              let _index57 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_389_v6).length));
              (_389_v6)[_index57] = (_module.__default.fm22(_403_v16, function (_pat_let4_0) {
                return function (_432_dt__update__tmp_h0) {
                  return function (_pat_let5_0) {
                    return function (_433_dt__update_hcf5_h0) {
                      return _module.D0.create_DC3(_433_dt__update_hcf5_h0);
                    }(_pat_let5_0);
                  }(_pat_let_tv9);
                }(_pat_let4_0);
              }(_431_v32), _428_v29, globalState)).minus(new BigNumber((_388_v5).length));
              let _434_v33;
              _434_v33 = _dafny.Map.Empty.slice().updateUnsafe(_430_v31,_430_v31);
              let _435_v34;
              _435_v34 = _dafny.Map.Empty.slice().updateUnsafe(_383_v0,_434_v33);
              _383_v0 = _module.__default.fm23(_430_v31, _435_v34, _428_v29, new BigNumber(-87), globalState);
            } else {
              let _436___mcc_h5 = (_source7).cf15;
              let _437_cf15 = _436___mcc_h5;
              let _438_v35;
              let _nw53 = new _module.C1();
              _nw53.__ctor();
              _438_v35 = _nw53;
              let _439_v36;
              _439_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_383_v0);
              _439_v36 = (_439_v36).update((p1).minus(p1), _383_v0);
              (globalState).f9 = p3;
              let _440_v37;
              let _nw54 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _440_v37 = _nw54;
              let _index58 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_440_v37).length));
              (_440_v37)[_index58] = _dafny.Seq.Concat(_392_v9, _392_v9);
              let _441_v38;
              _441_v38 = _dafny.Map.Empty.slice().updateUnsafe(p3,_392_v9);
              let _442_v39;
              _442_v39 = _dafny.Seq.of(_dafny.Seq.of(_383_v0, _383_v0), _392_v9, (((_441_v38).contains(p3)) ? ((_441_v38).get(p3)) : (_392_v9)), _392_v9, _dafny.Seq.of(_383_v0));
              let _index59 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_440_v37).length));
              (_440_v37)[_index59] = _dafny.Seq.Concat((_442_v39)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).fm0(p2, _383_v0, globalState)), new BigNumber((_442_v39).length))], _392_v9);
            }
            let _443_v40;
            _443_v40 = _dafny.Seq.of(_392_v9);
            let _444_v41;
            _444_v41 = _module.D0.create_DC1(p3, _443_v40, _392_v9);
            let _445_v42;
            _445_v42 = _dafny.MultiSet.fromElements(_444_v41, _444_v41);
            let _446_v43;
            let _nw55 = new _module.C0();
            _nw55.__ctor(_445_v42, p3);
            _446_v43 = _nw55;
            let _447_v44;
            _447_v44 = _dafny.Seq.of(p1, p2);
            let _448_v45;
            _448_v45 = _dafny.Map.Empty.slice().updateUnsafe(_447_v44,p1);
            _448_v45 = (_448_v45).update(_447_v44, p2);
          }
        }
      }
      let _449_v46;
      _449_v46 = new BigNumber(-694);
      _449_v46 = new BigNumber(278);
      let _450_v47;
      _450_v47 = _dafny.Set.fromElements(p1);
      let _source8 = _module.D0.create_DC0(new BigNumber((_450_v47).length));
      if (_source8.is_DC1) {
        let _451___mcc_h6 = (_source8).cf1;
        let _452___mcc_h7 = (_source8).cf2;
        let _453___mcc_h8 = (_source8).cf3;
        let _454_cf3 = _453___mcc_h8;
        let _455_cf2 = _452___mcc_h7;
        let _456_cf1 = _451___mcc_h6;
        let _457_v48;
        let _nw56 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
        _457_v48 = _nw56;
        let _458_v49;
        _458_v49 = _dafny.Seq.of(p2, p1);
        let _459_v50;
        _459_v50 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(_456_cf1, _454_cf3, _388_v5),_458_v49);
        let _index60 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_457_v48).length));
        (_457_v48)[_index60] = _459_v50;
        let _460_v51;
        _460_v51 = _module.D0.create_DC1(p3, _455_cf2, _454_cf3);
        let _461_v52;
        _461_v52 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.D0.create_DC3(_460_v51));
        let _462_v53;
        _462_v53 = _module.D2.create_DC7(p3, _461_v52, new BigNumber((_dafny.Seq.of(_456_cf1)).length), _456_cf1);
        let _index61 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_394_v11).length));
        (_394_v11)[_index61] = ((_462_v53).dtor_cf13).multipliedBy(_449_v46);
        let _index62 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_457_v48).length));
        let _index63 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_394_v11).length));
        let _rhs68 = _459_v50;
        let _rhs69 = _454_cf3;
        let _rhs70 = _449_v46;
        let _rhs71 = p3;
        let _rhs72 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((new BigNumber((_450_v47).length)).minus(p1), p2));
        let _lhs49 = _457_v48;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_457_v48).length));
        let _lhs51 = _394_v11;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_394_v11).length));
        let _lhs53 = globalState;
        _lhs49[_lhs50] = _rhs68;
        _454_cf3 = _rhs69;
        _lhs51[_lhs52] = _rhs70;
        _lhs53.f13 = _rhs71;
        _449_v46 = _rhs72;
        let _index64 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_394_v11).length));
        (_394_v11)[_index64] = _module.__default.safeDivisionInt(_449_v46, _module.__default.safeModuloInt(p1, p1));
        let _463_v54;
        _463_v54 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Map.Empty.slice().updateUnsafe(_456_cf1,_389_v6));
        let _464_v55;
        _464_v55 = _dafny.Map.Empty.slice().updateUnsafe(p3,_389_v6);
        let _465_v56;
        _465_v56 = _module.D1.create_DC4((((_463_v54).contains(_456_cf1)) ? ((_463_v54).get(_456_cf1)) : (_464_v55)));
        let _466_v57;
        _466_v57 = _dafny.Map.Empty.slice().updateUnsafe(_465_v56,(_462_v53).dtor_cf13);
        _466_v57 = (_466_v57).update(_465_v56, new BigNumber((_455_cf2).length));
        let _index65 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_394_v11).length));
        let _rhs73 = (_385_v2).Merge(_385_v2);
        let _rhs74 = new BigNumber(-205);
        let _rhs75 = _dafny.Seq.Concat(_454_cf3, _392_v9);
        let _rhs76 = ((_456_cf1) ? (_388_v5) : (_dafny.Seq.Concat(_388_v5, _388_v5)));
        let _lhs54 = _394_v11;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(285), new BigNumber((_394_v11).length));
        _385_v2 = _rhs73;
        _lhs54[_lhs55] = _rhs74;
        _454_cf3 = _rhs75;
        r0 = _rhs76;
      } else if (_source8.is_DC2) {
        let _467___mcc_h9 = (_source8).cf4;
        let _468_cf4 = _467___mcc_h9;
        let _469_v58;
        let _nw57 = Array((new BigNumber(2)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = _468_cf4;
        _nw57[(_dafny.ONE).toNumber()] = _module.__default.fm19(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-373)), ((_470_v0) => function (_471_i3) {
          return _470_v0;
        })(_383_v0))).length), _468_cf4, globalState);
        _469_v58 = _nw57;
        let _472_v59;
        _472_v59 = _dafny.Seq.of(_392_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), ((_473_v0) => function (_474_i4) {
          return _473_v0;
        })(_383_v0)), _dafny.Seq.UnicodeFromString("t"));
        let _475_v60;
        _475_v60 = _dafny.Seq.of(new BigNumber((_472_v59).length));
        let _476_v61;
        _476_v61 = _dafny.Map.Empty.slice().updateUnsafe(_469_v58,_dafny.Seq.Concat(_475_v60, _475_v60));
        _449_v46 = (_dafny.ZERO).minus(new BigNumber((_476_v61).length));
        let _477_v62;
        let _nw58 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _477_v62 = _nw58;
        let _index66 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length));
        (_477_v62)[_index66] = _dafny.Seq.Concat(_392_v9, _392_v9);
        let _index67 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length));
        (_389_v6)[_index67] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_468_cf4,p3)).length), new BigNumber((p0).length));
        let _index68 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length));
        let _index69 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length));
        let _rhs77 = _394_v11;
        let _rhs78 = _392_v9;
        let _rhs79 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.safeDivisionInt(new BigNumber((_475_v60).length), new BigNumber((_385_v2).length))).minus(p2)));
        let _rhs80 = _dafny.Seq.UnicodeFromString("dwq");
        let _lhs56 = _477_v62;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length));
        let _lhs58 = _389_v6;
        let _lhs59 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length));
        _389_v6 = _rhs77;
        _lhs56[_lhs57] = _rhs78;
        _lhs58[_lhs59] = _rhs79;
        _392_v9 = _rhs80;
        let _478_v63;
        _478_v63 = _module.D0.create_DC0((p1).plus(p2));
        let _source9 = _478_v63;
        if (_source9.is_DC1) {
          let _479___mcc_h12 = (_source9).cf1;
          let _480___mcc_h13 = (_source9).cf2;
          let _481___mcc_h14 = (_source9).cf3;
          let _482_cf3 = _481___mcc_h14;
          let _483_cf2 = _480___mcc_h13;
          let _484_cf1 = _479___mcc_h12;
          (globalState).f9 = _484_cf1;
          let _485_v64;
          _485_v64 = _module.D0.create_DC1(_484_cf1, _472_v59, (_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))]);
          let _486_v65;
          _486_v65 = _module.D0.create_DC3(_485_v64);
          let _487_v66;
          _487_v66 = _module.D0.create_DC3(_485_v64);
          let _488_v67;
          _488_v67 = _module.D2.create_DC7(_468_cf4, _dafny.Map.Empty.slice().updateUnsafe(_484_cf1,_487_v66), new BigNumber(779), _484_cf1);
          _449_v46 = _module.__default.fm22(_488_v67, _487_v66, (_388_v5)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_449_v46)).cardinality()), new BigNumber((_388_v5).length))], globalState);
          let _489_v68;
          _489_v68 = _dafny.MultiSet.fromElements(_394_v11, _394_v11);
          let _490_v69;
          _490_v69 = _dafny.Seq.of((_489_v68).update(_394_v11, _module.__default.abs(_449_v46)));
          let _index70 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length));
          (_389_v6)[_index70] = (_dafny.ZERO).minus(new BigNumber(((_490_v69)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_490_v69).length))]).cardinality()));
          _469_v58 = _469_v58;
        } else if (_source9.is_DC2) {
          let _491___mcc_h15 = (_source9).cf4;
          let _492_cf4 = _491___mcc_h15;
          let _493_v70;
          _493_v70 = _module.D0.create_DC0(p1);
          let _494_v71;
          _494_v71 = _module.D0.create_DC3(_493_v70);
          let _495_v72;
          _495_v72 = _dafny.Map.Empty.slice().updateUnsafe(p3,_494_v71);
          let _496_v73;
          _496_v73 = _module.D2.create_DC7(p3, _495_v72, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(804)), ((_497_v5) => function (_498_i5) {
  return _497_v5;
})(_388_v5))).length)), p3);
          let _499_v74;
          _499_v74 = _dafny.Set.fromElements(_383_v0);
          let _pat_let_tv10 = p3;
          let _pat_let_tv11 = _495_v72;
          let _pat_let_tv12 = _388_v5;
          let _pat_let_tv13 = _492_cf4;
          let _pat_let_tv14 = p3;
          let _500_v77;
          let _nw59 = Array((new BigNumber(26)).toNumber());
          _nw59[(_dafny.ZERO).toNumber()] = _496_v73;
          _nw59[(_dafny.ONE).toNumber()] = _496_v73;
          _nw59[(new BigNumber(2)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(3)).toNumber()] = function (_pat_let6_0) {
            return function (_501_dt__update__tmp_h1) {
              return function (_pat_let7_0) {
                return function (_502_dt__update_hcf11_h0) {
                  return function (_pat_let8_0) {
                    return function (_503_dt__update_hcf12_h0) {
                      return _module.D2.create_DC7(_502_dt__update_hcf11_h0, _503_dt__update_hcf12_h0, (_501_dt__update__tmp_h1).dtor_cf13, (_501_dt__update__tmp_h1).dtor_cf14);
                    }(_pat_let8_0);
                  }(_pat_let_tv11);
                }(_pat_let7_0);
              }(_pat_let_tv10);
            }(_pat_let6_0);
          }(_496_v73);
          _nw59[(new BigNumber(4)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(5)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(6)).toNumber()] = function (_pat_let9_0) {
            return function (_504_dt__update__tmp_h2) {
              return function (_pat_let10_0) {
                return function (_505_dt__update_hcf13_h0) {
                  return function (_pat_let11_0) {
                    return function (_506_dt__update_hcf14_h0) {
                      return function (_pat_let12_0) {
                        return function (_507_dt__update_hcf11_h1) {
                          return _module.D2.create_DC7(_507_dt__update_hcf11_h1, (_504_dt__update__tmp_h2).dtor_cf12, _505_dt__update_hcf13_h0, _506_dt__update_hcf14_h0);
                        }(_pat_let12_0);
                      }(_pat_let_tv14);
                    }(_pat_let11_0);
                  }(_pat_let_tv13);
                }(_pat_let10_0);
              }(new BigNumber((_pat_let_tv12).length));
            }(_pat_let9_0);
          }(_496_v73);
          _nw59[(new BigNumber(7)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(8)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(9)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(10)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(11)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(12)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(13)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(14)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(15)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(16)).toNumber()] = _module.D2.create_DC7(p3, (_495_v72).update(p3, _494_v71), new BigNumber((_499_v74).length), _492_cf4);
          _nw59[(new BigNumber(17)).toNumber()] = _module.__default.fm29(new BigNumber((_385_v2).length), (_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))], (_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))], p1, globalState);
          _nw59[(new BigNumber(18)).toNumber()] = _module.D2.create_DC7((_388_v5)[_module.__default.safeIndex(_449_v46, new BigNumber((_388_v5).length))], _495_v72, (_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))], p3);
          _nw59[(new BigNumber(19)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(20)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(21)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(22)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(23)).toNumber()] = ((_492_cf4) ? (_496_v73) : (_module.D2.create_DC7((_this).fm14(function () {
  let _coll12 = new _dafny.Set();
  for (const _compr_12 of (function () {
    let _coll13 = new _dafny.Map();
    for (const _compr_13 of _dafny.IntegerRange(new BigNumber(288), new BigNumber(-460))) {
      let _508_v75 = _compr_13;
      if (((new BigNumber(288)).isLessThanOrEqualTo(_508_v75)) && ((_508_v75).isLessThan(new BigNumber(-460)))) {
        _coll13.push([(_508_v75).minus(p1),_492_cf4]);
      }
    }
    return _coll13;
  }()).Keys.Elements) {
    let _509_v76 = _compr_12;
    if ((function () {
      let _coll14 = new _dafny.Map();
      for (const _compr_14 of _dafny.IntegerRange(new BigNumber(288), new BigNumber(-460))) {
        let _508_v75 = _compr_14;
        if (((new BigNumber(288)).isLessThanOrEqualTo(_508_v75)) && ((_508_v75).isLessThan(new BigNumber(-460)))) {
          _coll14.push([(_508_v75).minus(p1),_492_cf4]);
        }
      }
      return _coll14;
    }()).contains(_509_v76)) {
      _coll12.add((_509_v76).minus(new BigNumber(-977)));
    }
  }
  return _coll12;
}(), false, _392_v9, globalState), _495_v72, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(globalState),_492_cf4)).length), p3)));
          _nw59[(new BigNumber(24)).toNumber()] = _496_v73;
          _nw59[(new BigNumber(25)).toNumber()] = _496_v73;
          _500_v77 = _nw59;
          let _index71 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_500_v77).length));
          (_500_v77)[_index71] = _module.D2.create_DC7(p3, _495_v72, (_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))], p3);
          let _index72 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_500_v77).length));
          (_500_v77)[_index72] = _module.__default.fm29((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_388_v5, _388_v5)).length)), (p2).plus((_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))]), _392_v9, _module.__default.safeDivisionInt((_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))], (_496_v73).dtor_cf13), globalState);
          let _510_v78;
          _510_v78 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gortewya"),!(_492_cf4));
          let _511_v79;
          _511_v79 = _module.D0.create_DC1(_492_cf4, _472_v59, (_472_v59)[_module.__default.safeIndex(_449_v46, new BigNumber((_472_v59).length))]);
          let _512_v80;
          _512_v80 = _dafny.Set.fromElements(p3);
          _510_v78 = (_510_v78).update((_511_v79).dtor_cf3, ((_512_v80)).IsSubsetOf(p0));
          _394_v11 = _394_v11;
          let _513_v81;
          _513_v81 = _dafny.MultiSet.fromElements(_module.D0.create_DC1(_492_cf4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-952)), ((_514_v9) => function (_515_i6) {
  return _514_v9;
})(_392_v9)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), ((_516_v0) => function (_517_i7) {
  return _516_v0;
})(_383_v0))));
          let _518_v82;
          _518_v82 = _dafny.Seq.of(_513_v81, _513_v81);
          let _519_v83;
          _519_v83 = _dafny.MultiSet.fromElements((_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))]);
          let _520_v84;
          _520_v84 = _dafny.Seq.of((_518_v82)[_module.__default.safeIndex(new BigNumber((_519_v83).cardinality()), new BigNumber((_518_v82).length))], _513_v81);
          let _521_v85;
          _521_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,_518_v82);
          _521_v85 = (_521_v85).update(((_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))]).plus(_449_v46), _518_v82);
        } else if (_source9.is_DC0) {
          let _522___mcc_h16 = (_source9).cf0;
          let _523_cf0 = _522___mcc_h16;
          let _nw60 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _394_v11 = _nw60;
          let _rhs81 = _383_v0;
          let _rhs82 = (_dafny.MultiSet.fromElements(_383_v0)).Difference((_384_v1).Union(_384_v1));
          let _rhs83 = new BigNumber((_dafny.Seq.Concat((_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))], (_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))])).length);
          let _rhs84 = (_dafny.ZERO).minus(_449_v46);
          let _rhs85 = (_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))];
          _383_v0 = _rhs81;
          _384_v1 = _rhs82;
          _449_v46 = _rhs83;
          _449_v46 = _rhs84;
          _449_v46 = _rhs85;
          let _524_v86;
          _524_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_393_v10).length),_472_v59);
          let _525_v87;
          _525_v87 = _module.D0.create_DC1(p3, _dafny.Seq.of(_392_v9), (_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(_449_v46,new BigNumber((_391_v8).length)), globalState));
          let _526_v88;
          _526_v88 = _dafny.MultiSet.fromElements(_module.D0.create_DC1(p3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(638)), ((_527_v9) => function (_528_i8) {
  return _527_v9;
})(_392_v9)), _dafny.Seq.UnicodeFromString("cv")), _module.D0.create_DC1(_468_cf4, _472_v59, (_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))]), _module.__default.fm17(p3, globalState), _module.D0.create_DC1(p3, (((_524_v86).contains(_449_v46)) ? ((_524_v86).get(_449_v46)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(105)), ((_529_v62) => function (_530_i9) {
  return (_529_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_529_v62).length))];
})(_477_v62)))), _392_v9), _525_v87);
          let _531_v89;
          let _nw61 = new _module.C0();
          _nw61.__ctor(_526_v88, _468_cf4);
          _531_v89 = _nw61;
          let _532_v90;
          _532_v90 = _dafny.Seq.of(_module.__default.fm17((_531_v89).f19, globalState));
          let _533_v91;
          let _nw62 = new _module.C0();
          _nw62.__ctor((_dafny.MultiSet.FromArray(_532_v90)).update(_module.__default.fm17((_531_v89).f19, globalState), _module.__default.abs(_449_v46)), _468_cf4);
          _533_v91 = _nw62;
        } else {
          let _534___mcc_h17 = (_source9).cf5;
          let _535_cf5 = _534___mcc_h17;
          let _536_v92;
          _536_v92 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm19(new BigNumber((_475_v60).length), p3, globalState));
          let _537_v93;
          _537_v93 = _module.D0.create_DC1(_468_cf4, _472_v59, _392_v9);
          let _538_v94;
          _538_v94 = _dafny.Map.Empty.slice().updateUnsafe(_468_cf4,_module.D0.create_DC3(_537_v93));
          let _539_v95;
          _539_v95 = _module.D2.create_DC7((((_393_v10).contains(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), ((_542_v6) => function (_543_i10) {
  return (_542_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_542_v6).length))];
})(_389_v6)))).cardinality()))) ? ((_393_v10).get(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), ((_540_v6) => function (_541_i10) {
  return (_540_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_540_v6).length))];
})(_389_v6)))).cardinality()))) : (p3)), _538_v94, _449_v46, !(_468_cf4));
          let _544_v96;
          _544_v96 = _dafny.Seq.of(_539_v95);
          let _545_v97;
          _545_v97 = _dafny.Map.Empty.slice().updateUnsafe(p2,_469_v58);
          let _546_v98;
          _546_v98 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.of(new BigNumber((_536_v92).length)), _module.__default.safeIndex(new BigNumber((_475_v60).length), new BigNumber((_dafny.Seq.of(new BigNumber((_536_v92).length))).length)), new BigNumber((_544_v96).length)),new BigNumber((_545_v97).length));
          let _547_v99;
          let _nw63 = Array((_dafny.ONE).toNumber()).fill([]);
          _547_v99 = _nw63;
          let _548_v100;
          let _nw64 = Array((new BigNumber(11)).toNumber()).fill([]);
          _548_v100 = _nw64;
          let _index73 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_547_v99).length));
          (_547_v99)[_index73] = _548_v100;
          let _index74 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_547_v99).length));
          let _rhs86 = _546_v98;
          let _rhs87 = _548_v100;
          let _lhs60 = _547_v99;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_547_v99).length));
          _546_v98 = _rhs86;
          _lhs60[_lhs61] = _rhs87;
          let _549_v101;
          _549_v101 = _dafny.Map.Empty.slice().updateUnsafe((_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))],_475_v60);
          let _550_v102;
          let _nw65 = Array((new BigNumber(10)).toNumber());
          _nw65[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_475_v60, _module.__default.safeIndex(new BigNumber(((_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))]).length), new BigNumber((_475_v60).length)), p1);
          _nw65[(_dafny.ONE).toNumber()] = _475_v60;
          _nw65[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_475_v60, _475_v60);
          _nw65[(new BigNumber(3)).toNumber()] = ((false) ? (_475_v60) : (_module.__default.fm30(globalState)));
          _nw65[(new BigNumber(4)).toNumber()] = (((_549_v101).contains(_module.__default.fm20(p1, p3, _468_cf4, _449_v46, globalState))) ? ((_549_v101).get(_module.__default.fm20(p1, p3, _468_cf4, _449_v46, globalState))) : (_475_v60));
          _nw65[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-514)), ((_551_v60) => function (_552_i11) {
            return new BigNumber((_551_v60).length);
          })(_475_v60));
          _nw65[(new BigNumber(6)).toNumber()] = _475_v60;
          _nw65[(new BigNumber(7)).toNumber()] = _475_v60;
          _nw65[(new BigNumber(8)).toNumber()] = _475_v60;
          _nw65[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_389_v6)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_389_v6).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), function (_553_i12) {
            return new BigNumber(-161);
          }));
          _550_v102 = _nw65;
          (globalState).f2 = _550_v102;
          let _554_v103;
          _554_v103 = _module.D1.create_DC5(!(p3), (_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))], _388_v5);
          let _rhs88 = (((_dafny.ZERO).minus(new BigNumber((_392_v9).length))).plus(p2)).multipliedBy(p1);
          let _rhs89 = (_554_v103).dtor_cf7;
          let _rhs90 = _468_cf4;
          let _lhs62 = globalState;
          let _lhs63 = globalState;
          _449_v46 = _rhs88;
          _lhs62.f13 = _rhs89;
          _lhs63.f5 = _rhs90;
          let _555_v104;
          _555_v104 = _dafny.Map.Empty.slice().updateUnsafe(p3,_394_v11);
          let _556_v105;
          _556_v105 = _module.D1.create_DC4(_555_v104);
          let _557_v106;
          _557_v106 = _dafny.Map.Empty.slice().updateUnsafe(_475_v60,_556_v105);
          let _558_v107;
          _558_v107 = _dafny.Map.Empty.slice().updateUnsafe(_557_v106,p3);
          _558_v107 = (_558_v107).update(_557_v106, p3);
        }
        let _559_v108;
        let _nw66 = Array((new BigNumber(8)).toNumber());
        _559_v108 = _nw66;
        let _560_v109;
        _560_v109 = _module.D0.create_DC1(_468_cf4, _472_v59, (_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))]);
        let _561_v110;
        _561_v110 = _module.D0.create_DC1(false, (_560_v109).dtor_cf2, (_477_v62)[_module.__default.safeIndex(new BigNumber(753), new BigNumber((_477_v62).length))]);
        let _562_v111;
        _562_v111 = _dafny.MultiSet.fromElements(_561_v110);
        let _563_v112;
        let _nw67 = new _module.C0();
        _nw67.__ctor(_562_v111, (_560_v109).dtor_cf1);
        _563_v112 = _nw67;
        let _564_v113;
        _564_v113 = _dafny.Map.Empty.slice().updateUnsafe(_559_v108,_563_v112);
        let _565_v114;
        _565_v114 = _dafny.Map.Empty.slice().updateUnsafe(_564_v113,p1);
        let _index75 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_469_v58).length));
        (_469_v58)[_index75] = !((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).fm0((((_565_v114).contains(_564_v113)) ? ((_565_v114).get(_564_v113)) : (new BigNumber(-858))), _383_v0, globalState)))).isEqualTo(p1);
        let _index76 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_469_v58).length));
        (_469_v58)[_index76] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ooskv"), _module.__default.fm20(_449_v46, _468_cf4, (_563_v112).f19, _449_v46, globalState)), _dafny.Seq.UnicodeFromString("xfq"));
      } else if (_source8.is_DC0) {
        let _566___mcc_h10 = (_source8).cf0;
        let _567_cf0 = _566___mcc_h10;
        if (p3) {
          let _568_v115;
          _568_v115 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), ((_569_v0) => function (_570_i13) {
            return _569_v0;
          })(_383_v0)), _dafny.Seq.update(_392_v9, _module.__default.safeIndex(_567_cf0, new BigNumber((_392_v9).length)), _383_v0));
          let _571_v116;
          _571_v116 = _module.D0.create_DC1(p3, _568_v115, _392_v9);
          let _572_v117;
          _572_v117 = _module.D0.create_DC3(_571_v116);
          let _573_v118;
          _573_v118 = _module.D0.create_DC3(_571_v116);
          let _574_v119;
          _574_v119 = _module.D2.create_DC7(false, _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(new BigNumber((_450_v47).length), p3, globalState),_573_v118), p2, p3);
          _392_v9 = (_this).fm1((_dafny.Map.Empty.slice().updateUnsafe(_449_v46,(_574_v119).dtor_cf13)).Merge(_module.__default.fm31(globalState)), globalState);
          let _575_v120;
          _575_v120 = _module.D0.create_DC1(p3, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), function (_576_i14) {
  return _dafny.Seq.UnicodeFromString("uert");
}), _module.__default.safeIndex(_449_v46, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), function (_577_i14) {
  return _dafny.Seq.UnicodeFromString("uert");
})).length)), _392_v9), _392_v9);
          let _578_v121;
          _578_v121 = _dafny.MultiSet.fromElements(_575_v120);
          let _579_v122;
          let _nw68 = Array((new BigNumber(14)).toNumber()).fill(false);
          _579_v122 = _nw68;
          let _580_v123;
          _580_v123 = _dafny.Map.Empty.slice().updateUnsafe(_579_v122,new BigNumber((_387_v4).cardinality()));
          let _581_v124;
          let _nw69 = new _module.C0();
          _nw69.__ctor((_578_v121).update(_module.D0.create_DC1(p3, _568_v115, _dafny.Seq.UnicodeFromString("xbt")), _module.__default.abs(new BigNumber((_580_v123).length))), p3);
          _581_v124 = _nw69;
          let _582_v125;
          let _583_v126;
          let _584_v127;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector2 = (_this).m6(_392_v9, _392_v9, p3, new BigNumber(163), globalState);
          _out5 = _outcollector2[0];
          _out6 = _outcollector2[1];
          _out7 = _outcollector2[2];
          _582_v125 = _out5;
          _583_v126 = _out6;
          _584_v127 = _out7;
          let _585_v128;
          let _586_v129;
          let _587_v130;
          let _out8;
          let _out9;
          let _out10;
          let _outcollector3 = (_this).m6(_392_v9, _392_v9, _582_v125, p2, globalState);
          _out8 = _outcollector3[0];
          _out9 = _outcollector3[1];
          _out10 = _outcollector3[2];
          _585_v128 = _out8;
          _586_v129 = _out9;
          _587_v130 = _out10;
          _583_v126 = (_module.__default.safeModuloInt(_567_cf0, (_dafny.ZERO).minus(_567_cf0))).plus(_567_cf0);
        } else {
          _449_v46 = p1;
          let _588_v131;
          let _nw70 = Array((new BigNumber(4)).toNumber()).fill(_dafny.MultiSet.Empty);
          _588_v131 = _nw70;
          let _rhs91 = _588_v131;
          let _rhs92 = p3;
          let _lhs64 = globalState;
          _588_v131 = _rhs91;
          _lhs64.f13 = _rhs92;
          (globalState).f3 = p3;
          let _589_v132;
          _589_v132 = _module.D1.create_DC5(p3, _392_v9, _388_v5);
          let _590_v133;
          _590_v133 = _dafny.Map.Empty.slice().updateUnsafe(_589_v132,_392_v9);
          _590_v133 = _590_v133;
          let _591_v134;
          _591_v134 = p0;
          let _592_v135;
          _592_v135 = _dafny.Seq.of(_392_v9);
          let _593_v136;
          _593_v136 = _module.D0.create_DC1(p3, _592_v135, _392_v9);
          let _594_v137;
          _594_v137 = _dafny.MultiSet.fromElements(_593_v136);
          let _595_v138;
          let _nw71 = new _module.C0();
          _nw71.__ctor(_594_v137, (((_393_v10).contains(new BigNumber((_387_v4).cardinality()))) ? ((_393_v10).get(new BigNumber((_387_v4).cardinality()))) : (p3)));
          _595_v138 = _nw71;
          let _596_v139;
          _596_v139 = _dafny.Seq.of(_595_v138);
          let _rhs93 = (p0).Union(p0);
          let _rhs94 = _389_v6;
          let _rhs95 = _dafny.Seq.Concat(_596_v139, _596_v139);
          _591_v134 = _rhs93;
          _394_v11 = _rhs94;
          _596_v139 = _rhs95;
        }
        let _597_v140;
        _597_v140 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("astkcoj"),p3);
        let _598_v142;
        _598_v142 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("v"),new BigNumber((_388_v5).length));
        _597_v140 = function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of ((_598_v142).update(_dafny.Seq.UnicodeFromString("rjqeb"), new BigNumber((_392_v9).length))).Keys.Elements) {
            let _599_v141 = _compr_15;
            if (((_598_v142).update(_dafny.Seq.UnicodeFromString("rjqeb"), new BigNumber((_392_v9).length))).contains(_599_v141)) {
              _coll15.push([_599_v141,p3]);
            }
          }
          return _coll15;
        }();
        let _600_v143;
        _600_v143 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
        let _601_v144;
        _601_v144 = _dafny.Set.fromElements((((_600_v143).contains((_this).fm14(_dafny.Set.fromElements(p1), p3, _392_v9, globalState))) ? ((_600_v143).get((_this).fm14(_dafny.Set.fromElements(p1), p3, _392_v9, globalState))) : (p3)), p3);
        _601_v144 = p0;
        let _602_v145;
        _602_v145 = _dafny.Seq.of(_392_v9, _392_v9);
        let _603_v146;
        _603_v146 = _module.D0.create_DC0(new BigNumber(((_module.D0.create_DC1(!(p3), _dafny.Seq.update(_602_v145, _module.__default.safeIndex(p1, new BigNumber((_602_v145).length)), _392_v9), _392_v9)).dtor_cf3).length));
        let _604_v147;
        _604_v147 = _module.D0.create_DC3(_603_v146);
        _604_v147 = _604_v147;
      } else {
        let _605___mcc_h11 = (_source8).cf5;
        let _606_cf5 = _605___mcc_h11;
        r2 = _dafny.Seq.IsProperPrefixOf(_392_v9, _dafny.Seq.Concat(_392_v9, _dafny.Seq.UnicodeFromString("ienvw")));
        _393_v10 = (_393_v10).update(_module.__default.safeModuloInt(p2, new BigNumber(-748)), (true) && (p3));
        let _607_v148;
        let _nw72 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _607_v148 = _nw72;
        let _608_v149;
        _608_v149 = _dafny.Map.Empty.slice().updateUnsafe(p3,_392_v9);
        let _index77 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_607_v148).length));
        (_607_v148)[_index77] = _dafny.Seq.Concat(_392_v9, (((_608_v149).contains(true)) ? ((_608_v149).get(true)) : (_392_v9)));
        let _index78 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_607_v148).length));
        (_607_v148)[_index78] = _module.__default.fm20(p2, p3, _dafny.Seq.IsProperPrefixOf(_392_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), ((_609_v0) => function (_610_i15) {
          return _609_v0;
        })(_383_v0))), p1, globalState);
        (globalState).f3 = p3;
      }
      let _611_v150;
      let _init6 = ((_612_p3, _613_v9) => function (_614_i16) {
        return (_module.D0.create_DC1(_612_p3, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("usbo")), _613_v9)).dtor_cf1;
      })(p3, _392_v9);
      let _nw73 = Array((new BigNumber(29)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw73.length); _i0_6++) {
        _nw73[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _611_v150 = _nw73;
      let _index79 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length));
      (_611_v150)[_index79] = p3;
      let _615_v151;
      _615_v151 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
      let _616_v152;
      _616_v152 = _dafny.Map.Empty.slice().updateUnsafe(_615_v151,p3);
      let _index80 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_611_v150).length));
      (_611_v150)[_index80] = (_616_v152).contains(_615_v151);
      let _617_v153;
      _617_v153 = _dafny.MultiSet.fromElements(new BigNumber(-21), (_dafny.ZERO).minus((_dafny.ZERO).minus(_449_v46)));
      let _618_v154;
      _618_v154 = _dafny.Seq.of(_617_v153);
      let _index81 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length));
      let _index82 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_611_v150).length));
      let _rhs96 = p3;
      let _rhs97 = (_this).fm14(_450_v47, p3, _392_v9, globalState);
      let _rhs98 = !_dafny.areEqual(_618_v154, ((p3) ? (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_391_v8).length), p2, p1), _617_v153)) : (_618_v154)));
      let _rhs99 = _449_v46;
      let _lhs65 = _611_v150;
      let _lhs66 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length));
      let _lhs67 = _611_v150;
      let _lhs68 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_611_v150).length));
      let _lhs69 = globalState;
      _lhs65[_lhs66] = _rhs96;
      _lhs67[_lhs68] = _rhs97;
      _lhs69.f5 = _rhs98;
      _449_v46 = _rhs99;
      let _619_v155;
      _619_v155 = _module.D0.create_DC0(p2);
      let _620_v156;
      _620_v156 = _module.D0.create_DC3(_module.D0.create_DC3(_619_v155));
      let _621_v157;
      _621_v157 = _dafny.Map.Empty.slice().updateUnsafe(p3,_620_v156);
      let _622_v158;
      _622_v158 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_module.__default.fm22(_module.D2.create_DC7((_611_v150)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length))], _621_v157, new BigNumber((_dafny.MultiSet.fromElements(p3, false)).cardinality()), p3), _module.__default.fm28(globalState), (_611_v150)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length))], globalState)), _391_v8);
      let _623_v159;
      _623_v159 = _dafny.Seq.of(_392_v9, _392_v9);
      let _624_v160;
      _624_v160 = _module.D0.create_DC1((_611_v150)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length))], _623_v159, _392_v9);
      let _625_v161;
      _625_v161 = _module.D2.create_DC6(_dafny.MultiSet.fromElements(_624_v160, _module.D0.create_DC1((_611_v150)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length))], _623_v159, _dafny.Seq.UnicodeFromString("bfwqt")), _624_v160));
      let _pat_let_tv15 = _622_v158;
      let _pat_let_tv16 = _622_v158;
      let _pat_let_tv17 = _622_v158;
      _622_v158 = function (_source10) {
        if (_source10.is_DC7) {
          let _626___mcc_h18 = (_source10).cf11;
          let _627___mcc_h19 = (_source10).cf12;
          let _628___mcc_h20 = (_source10).cf13;
          let _629___mcc_h21 = (_source10).cf14;
          let _630_cf14 = _629___mcc_h21;
          let _631_cf13 = _628___mcc_h20;
          let _632_cf12 = _627___mcc_h19;
          let _633_cf11 = _626___mcc_h18;
          return _pat_let_tv15;
        } else if (_source10.is_DC6) {
          let _634___mcc_h22 = (_source10).cf10;
          let _635_cf10 = _634___mcc_h22;
          return _pat_let_tv16;
        } else {
          let _636___mcc_h23 = (_source10).cf15;
          let _637_cf15 = _636___mcc_h23;
          return _pat_let_tv17;
        }
      }(_625_v161);
      r0 = _dafny.Seq.Concat(_388_v5, _dafny.Seq.Concat(_388_v5, _388_v5));
      r1 = _388_v5;
      r2 = (_611_v150)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_611_v150).length))];
      return [r0, r1, r2];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f16 = false;
      this._f17 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f16, f17) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_this).f17);
    };
    fm1(p0, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("vcy");
    };
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Seq.of((_this).f16, (_this).f16), _dafny.Seq.of((_this).f16), _dafny.Seq.of((_this).f16))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of((_this).f16), _dafny.Seq.of((_this).f16, (_this).f16, (_this).f16, (_this).f16, (_this).f16)))).IsDisjointFrom((function () {
        let _coll16 = new _dafny.Set();
        for (const _compr_16 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_638_i0) {
          return _dafny.Seq.of((_this).f16);
        })).Elements) {
          let _639_v0 = _compr_16;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_638_i0) {
            return _dafny.Seq.of((_this).f16);
          }), _639_v0)) {
            _coll16.add(_639_v0);
          }
        }
        return _coll16;
      }()).Intersect(_dafny.Set.fromElements(_dafny.Seq.of((_this).f16))));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _640_v0;
      let _nw74 = new _module.C2();
      _nw74.__ctor();
      _640_v0 = _nw74;
      let _641_v1;
      let _nw75 = Array((new BigNumber(26)).toNumber()).fill(false);
      _641_v1 = _nw75;
      _641_v1 = _641_v1;
      r1 = (_this).f17;
      (globalState).f3 = ((_this).f16) || ((_this).f16);
      let _642_v2;
      let _nw76 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _642_v2 = _nw76;
      let _643_v3;
      _643_v3 = _dafny.Map.Empty.slice().updateUnsafe(_642_v2,(_this).f16);
      let _644_v4;
      _644_v4 = _dafny.Seq.UnicodeFromString("eehtvwmi");
      let _645_v5;
      _645_v5 = _dafny.Seq.of((_this).f16);
      let _646_v6;
      _646_v6 = _module.D1.create_DC5((_this).f16, _dafny.Seq.UnicodeFromString("dr"), _645_v5);
      let _647_v7;
      _647_v7 = _dafny.Map.Empty.slice().updateUnsafe(_643_v3,new BigNumber(((((_this).f16) ? (_644_v4) : ((_646_v6).dtor_cf8))).length));
      _647_v7 = (_647_v7).update(_643_v3, (_this).fm0((_this).f17, p0, globalState));
      let _648_v8;
      let _nw77 = new _module.C1();
      _nw77.__ctor();
      _648_v8 = _nw77;
      r0 = (_this).f16;
      let _649_v9;
      _649_v9 = _dafny.Seq.of(_644_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(828)), ((_650_p0) => function (_651_i0) {
        return _650_p0;
      })(p0)));
      let _652_v10;
      _652_v10 = _module.D0.create_DC1(true, _649_v9, _644_v4);
      let _653_v11;
      _653_v11 = _dafny.MultiSet.fromElements((function (_pat_let13_0) {
        return function (_654_dt__update__tmp_h0) {
          return function (_pat_let14_0) {
            return function (_655_dt__update_hcf1_h0) {
              return _module.D0.create_DC1(_655_dt__update_hcf1_h0, (_654_dt__update__tmp_h0).dtor_cf2, (_654_dt__update__tmp_h0).dtor_cf3);
            }(_pat_let14_0);
          }((_this).f16);
        }(_pat_let13_0);
      }(_652_v10)).dtor_cf3);
      r1 = ((_this).f17).plus(new BigNumber((_653_v11).cardinality()));
      return [r0, r1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f14, f15) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of((_module.D0.create_DC0((_dafny.ZERO).minus((_this).f14))).dtor_cf0), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Set.fromElements((_this).f14)).Elements) {
          let _656_v0 = _compr_17;
          if ((_dafny.Set.fromElements((_this).f14)).contains(_656_v0)) {
            _coll17.push([_module.__default.safeModuloInt(_656_v0, new BigNumber(-958)),false]);
          }
        }
        return _coll17;
      }()).length), new BigNumber(832), (_this).f14), _dafny.Seq.of((_this).f14)));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return !((_dafny.Set.fromElements((_this).f14)).IsDisjointFrom(_dafny.Set.fromElements((_this).f14, new BigNumber((function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of (_dafny.Set.fromElements((_this).f14)).Elements) {
          let _657_v0 = _compr_18;
          if ((_dafny.Set.fromElements((_this).f14)).contains(_657_v0)) {
            _coll18.push([_module.__default.safeDivisionInt(_657_v0, (_this).f14),true]);
          }
        }
        return _coll18;
      }()).length), (_dafny.ZERO).minus((_this).f14), new BigNumber((_dafny.MultiSet.fromElements((_this).f14)).cardinality())))) || (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("axmcm"), _dafny.Seq.UnicodeFromString("elfoqgqbn")));
    };
    m5(p0, globalState) {
      let _this = this;
      let _658_v0;
      _658_v0 = false;
      let _659_i0;
      _659_i0 = _dafny.ZERO;
      L3: {
        while (_658_v0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_659_i0)) {
              break L3;
            }
            _659_i0 = (_659_i0).plus(_dafny.ONE);
            let _660_v1;
            _660_v1 = _module.D0.create_DC0(_module.__default.safeDivisionInt((_this).f14, (_this).f14));
            let _661_v2;
            _661_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,!(_658_v0));
            let _662_v3;
            _662_v3 = _dafny.Seq.of((((_661_v2).contains((_this).f14)) ? ((_661_v2).get((_this).f14)) : (_658_v0)), _658_v0);
            let _pat_let_tv18 = _662_v3;
            _660_v1 = function (_pat_let15_0) {
              return function (_663_dt__update__tmp_h0) {
                return function (_pat_let16_0) {
                  return function (_664_dt__update_hcf0_h0) {
                    return _module.D0.create_DC0(_664_dt__update_hcf0_h0);
                  }(_pat_let16_0);
                }(new BigNumber((_pat_let_tv18).length));
              }(_pat_let15_0);
            }(_660_v1);
            let _665_v4;
            _665_v4 = new BigNumber(665);
            _665_v4 = (_this).f14;
            let _666_v5;
            _666_v5 = _dafny.Map.Empty.slice().updateUnsafe(_658_v0,_658_v0);
            if (((!((_dafny.Set.fromElements((_this).f14)).IsDisjointFrom(_dafny.Set.fromElements((_dafny.ZERO).minus(_665_v4), _665_v4, new BigNumber((_666_v5).length), (_this).f14)))) ? (_658_v0) : (_658_v0))) {
              let _667_v6;
              _667_v6 = _dafny.Map.Empty.slice().updateUnsafe(_658_v0,_665_v4);
              let _668_v7;
              _668_v7 = _dafny.MultiSet.fromElements((_this).f14, _665_v4);
              let _669_v8;
              _669_v8 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_658_v0, _658_v0, _658_v0),(_this).f14);
              let _670_v9;
              _670_v9 = _dafny.Map.Empty.slice().updateUnsafe(_658_v0,_669_v8);
              let _671_v10;
              let _nw78 = Array((new BigNumber(26)).toNumber());
              _nw78[(_dafny.ZERO).toNumber()] = (((_667_v6).contains(_658_v0)) ? ((_667_v6).get(_658_v0)) : ((_this).f14));
              _nw78[(_dafny.ONE).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(2)).toNumber()] = (new BigNumber((_668_v7).cardinality())).minus((_this).f14);
              _nw78[(new BigNumber(3)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(4)).toNumber()] = (_660_v1).dtor_cf0;
              _nw78[(new BigNumber(5)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(6)).toNumber()] = new BigNumber(321);
              _nw78[(new BigNumber(7)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(8)).toNumber()] = new BigNumber(((((_670_v9).contains(false)) ? ((_670_v9).get(false)) : (_669_v8))).length);
              _nw78[(new BigNumber(9)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(_665_v4, new BigNumber(841));
              _nw78[(new BigNumber(11)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(12)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_module.__default.fm9(new BigNumber(-178), _662_v3, p0, globalState), _662_v3)).length);
              _nw78[(new BigNumber(14)).toNumber()] = _665_v4;
              _nw78[(new BigNumber(15)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(16)).toNumber()] = new BigNumber((_module.__default.fm10(_658_v0, new BigNumber(569), (_this).fm8(_665_v4, _658_v0, _658_v0, globalState), !(_658_v0), globalState)).length);
              _nw78[(new BigNumber(17)).toNumber()] = (new BigNumber((_668_v7).cardinality())).minus((_dafny.ZERO).minus(_module.__default.fm11(_658_v0, globalState)));
              _nw78[(new BigNumber(18)).toNumber()] = _665_v4;
              _nw78[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_665_v4), _665_v4);
              _nw78[(new BigNumber(20)).toNumber()] = _665_v4;
              _nw78[(new BigNumber(21)).toNumber()] = (_this).f14;
              _nw78[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_665_v4).minus(_665_v4));
              _nw78[(new BigNumber(23)).toNumber()] = ((_this).f14).plus((_this).f14);
              _nw78[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(_665_v4);
              _nw78[(new BigNumber(25)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("sdnbh")).length);
              _671_v10 = _nw78;
              let _index83 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_671_v10).length));
              (_671_v10)[_index83] = _665_v4;
              let _672_v11;
              _672_v11 = _dafny.MultiSet.fromElements(_module.__default.fm12(_658_v0, _658_v0, (_this).f14, globalState));
              let _index84 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_671_v10).length));
              (_671_v10)[_index84] = new BigNumber((function () {
                let _coll19 = new _dafny.Set();
                for (const _compr_19 of (_672_v11).Elements) {
                  let _673_v12 = _compr_19;
                  if ((_672_v11).contains(_673_v12)) {
                    _coll19.add(_673_v12);
                  }
                }
                return _coll19;
              }()).length);
              (globalState).f13 = _658_v0;
              (globalState).f5 = _658_v0;
              let _674_v13;
              let _nw79 = new _module.C2();
              _nw79.__ctor();
              _674_v13 = _nw79;
              let _675_v14;
              _675_v14 = _dafny.Set.fromElements(_658_v0, (_667_v6).equals(_module.__default.fm15(_658_v0, _658_v0, _658_v0, globalState)));
              let _676_v15;
              _676_v15 = _module.D1.create_DC5(_658_v0, p0, _dafny.Seq.of(true));
              let _677_v16;
              let _init7 = ((_678_v0) => function (_679_i1) {
                return _678_v0;
              })(_658_v0);
              let _nw80 = Array((new BigNumber(24)).toNumber());
              for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw80.length); _i0_7++) {
                _nw80[_i0_7] = _init7(new BigNumber(_i0_7));
              }
              _677_v16 = _nw80;
              let _index85 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_677_v16).length));
              (_677_v16)[_index85] = (_662_v3)[_module.__default.safeIndex(_665_v4, new BigNumber((_662_v3).length))];
              let _680_v17;
              _680_v17 = _dafny.Set.fromElements((_this).f14, _665_v4, (_this).f14, (_dafny.ZERO).minus((_this).f14));
              let _index86 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_677_v16).length));
              let _rhs100 = _675_v14;
              let _rhs101 = _676_v15;
              let _rhs102 = (_674_v13).fm14(_680_v17, _658_v0, p0, globalState);
              let _lhs70 = _677_v16;
              let _lhs71 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_677_v16).length));
              _675_v14 = _rhs100;
              _676_v15 = _rhs101;
              _lhs70[_lhs71] = _rhs102;
            } else {
              let _681_v18;
              let _nw81 = new _module.C2();
              _nw81.__ctor();
              _681_v18 = _nw81;
              _681_v18 = _681_v18;
              let _682_v19;
              let _nw82 = Array((new BigNumber(17)).toNumber()).fill(false);
              _682_v19 = _nw82;
              let _index87 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_682_v19).length));
              (_682_v19)[_index87] = _658_v0;
              let _683_v20;
              _683_v20 = _dafny.Seq.UnicodeFromString("rtdk");
              let _684_v21;
              _684_v21 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(_682_v19, _682_v19)).length), (_this).f14, (_this).f14, _665_v4, new BigNumber((_662_v3).length));
              let _index88 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_682_v19).length));
              let _rhs103 = _dafny.Seq.IsProperPrefixOf(_684_v21, _684_v21);
              let _rhs104 = p0;
              let _lhs72 = _682_v19;
              let _lhs73 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_682_v19).length));
              _lhs72[_lhs73] = _rhs103;
              _683_v20 = _rhs104;
              let _685_v23;
              _685_v23 = _dafny.MultiSet.fromElements(new BigNumber((function () {
                let _coll20 = new _dafny.Set();
                for (const _compr_20 of _dafny.IntegerRange(new BigNumber(720), new BigNumber(-698))) {
                  let _686_v22 = _compr_20;
                  if (((new BigNumber(720)).isLessThanOrEqualTo(_686_v22)) && ((_686_v22).isLessThan(new BigNumber(-698)))) {
                    _coll20.add(_module.__default.safeDivisionInt(_686_v22, (_this).f14));
                  }
                }
                return _coll20;
              }()).length));
              _685_v23 = _685_v23;
              let _687_v24;
              _687_v24 = _dafny.Seq.of(_683_v20);
              (globalState).f13 = !(_dafny.areEqual(_687_v24, _687_v24));
              let _688_v25;
              _688_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_682_v19)[_module.__default.safeIndex(new BigNumber(996), new BigNumber((_682_v19).length))], (_682_v19)[_module.__default.safeIndex(new BigNumber(996), new BigNumber((_682_v19).length))], (_682_v19)[_module.__default.safeIndex(new BigNumber(996), new BigNumber((_682_v19).length))]),_682_v19);
              _682_v19 = (((_688_v25).contains(_662_v3)) ? ((_688_v25).get(_662_v3)) : (_682_v19));
            }
            let _689_v26;
            _689_v26 = _dafny.Seq.UnicodeFromString("y");
            let _690_v27;
            _690_v27 = _dafny.Set.fromElements((_dafny.ZERO).minus(_665_v4), (_this).f14);
            let _691_v28;
            _691_v28 = _dafny.Map.Empty.slice().updateUnsafe(_658_v0,_665_v4);
            let _692_v29;
            _692_v29 = _dafny.Map.Empty.slice().updateUnsafe(_690_v27,new BigNumber((_691_v28).length));
            let _rhs105 = _dafny.Seq.Concat(_dafny.Seq.Concat(_689_v26, _689_v26), _689_v26);
            let _rhs106 = (_692_v29).Merge((_692_v29).Merge(function () {
              let _coll21 = new _dafny.Map();
              for (const _compr_21 of (_692_v29).Keys.Elements) {
                let _693_v30 = _compr_21;
                if ((_692_v29).contains(_693_v30)) {
                  _coll21.push([_693_v30,new BigNumber(615)]);
                }
              }
              return _coll21;
            }()));
            _689_v26 = _rhs105;
            _692_v29 = _rhs106;
          }
        }
      }
      let _694_v31;
      let _nw83 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
      _694_v31 = _nw83;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_694_v31).length))) {
        let _695_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_695_i2)) && ((_695_i2).isLessThan(new BigNumber((_694_v31).length))))) {
          (_694_v31)[(_695_i2)] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(213),new BigNumber(746))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(469),(_this).f14)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),(_this).f14)));
        }
      }
      let _hi5 = (_this).f14;
      for (let _696_i3 = new BigNumber((_module.__default.fm32(true, _658_v0, p0, _658_v0, globalState)).length); _696_i3.isLessThan(_hi5); _696_i3 = _696_i3.plus(_dafny.ONE)) {
        let _697_v32;
        _697_v32 = _dafny.Seq.of(_658_v0, _dafny.Seq.IsPrefixOf(p0, p0));
        if ((_697_v32)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_697_v32).length))]) {
          let _698_v33;
          _698_v33 = new BigNumber(412);
          _698_v33 = new BigNumber(491);
          let _699_v34;
          let _nw84 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _699_v34 = _nw84;
          let _index89 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_699_v34).length));
          (_699_v34)[_index89] = _696_i3;
          let _index90 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_699_v34).length));
          (_699_v34)[_index90] = _696_i3;
          let _700_v35;
          _700_v35 = _dafny.Set.fromElements(_module.__default.fm11(false, globalState));
          let _701_v36;
          _701_v36 = _module.D4.create_DC10(_700_v35);
          let _702_v37;
          _702_v37 = _dafny.Map.Empty.slice().updateUnsafe(!((_700_v35).IsDisjointFrom((_701_v36).dtor_cf17)),_696_i3);
          _702_v37 = (_702_v37).update(_658_v0, _698_v33);
          let _703_v38;
          let _nw85 = new _module.C2();
          _nw85.__ctor();
          _703_v38 = _nw85;
          _698_v33 = _696_i3;
        } else {
          let _704_v39;
          _704_v39 = new BigNumber(528);
          _704_v39 = (_dafny.ZERO).minus((((_this).f14).minus(_704_v39)).multipliedBy(_696_i3));
          (globalState).f5 = (_704_v39).isLessThanOrEqualTo((_this).f14);
          let _705_v40;
          _705_v40 = new _dafny.CodePoint('j'.codePointAt(0));
          let _706_v41;
          _706_v41 = _dafny.MultiSet.fromElements(_705_v40);
          _704_v39 = new BigNumber(((_706_v41).Union(_dafny.MultiSet.FromArray(p0))).cardinality());
          let _707_v42;
          let _init8 = ((_708_v0) => function (_709_i4) {
            return _708_v0;
          })(_658_v0);
          let _nw86 = Array((new BigNumber(27)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw86.length); _i0_8++) {
            _nw86[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _707_v42 = _nw86;
          let _710_v43;
          _710_v43 = _module.D5.create_DC13(_707_v42);
          _707_v42 = (_710_v43).dtor_cf19;
          let _711_v44;
          _711_v44 = _module.D0.create_DC2(_658_v0);
          let _712_v45;
          _712_v45 = _dafny.Seq.of(p0, p0);
          let _713_v46;
          _713_v46 = _module.D0.create_DC1(_module.__default.fm19(_696_i3, (_711_v44).dtor_cf4, globalState), _712_v45, p0);
          let _714_v47;
          _714_v47 = _dafny.MultiSet.fromElements(_713_v46);
          let _715_v48;
          _715_v48 = _dafny.Seq.of(_714_v47);
          let _716_v49;
          let _nw87 = new _module.C0();
          _nw87.__ctor((_715_v48)[_module.__default.safeIndex(_704_v39, new BigNumber((_715_v48).length))], false);
          _716_v49 = _nw87;
        }
        let _717_v50;
        let _nw88 = Array((new BigNumber(5)).toNumber()).fill([]);
        _717_v50 = _nw88;
        let _718_v51;
        let _nw89 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _718_v51 = _nw89;
        let _index91 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_717_v50).length));
        (_717_v50)[_index91] = _718_v51;
        let _719_v53;
        _719_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f14,p0)).Keys.Elements) {
            let _720_v52 = _compr_22;
            if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,p0)).contains(_720_v52)) {
              _coll22.add(_module.__default.safeDivisionInt(_720_v52, new BigNumber(275)));
            }
          }
          return _coll22;
        }()).length),new BigNumber(332));
        let _721_v54;
        _721_v54 = _dafny.MultiSet.fromElements(_696_i3, (_this).f14, new BigNumber((_719_v53).length), _696_i3);
        let _index92 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_717_v50).length));
        (_717_v50)[_index92] = (((_721_v54).IsSubsetOf(_dafny.MultiSet.fromElements(_696_i3))) ? (_718_v51) : (_718_v51));
        _719_v53 = (_719_v53).update(((_this).f14).minus((_this).f14), _module.__default.safeModuloInt(_696_i3, (_this).f14));
        let _722_v55;
        _722_v55 = new BigNumber(977);
        let _723_v56;
        _723_v56 = new _dafny.CodePoint('x'.codePointAt(0));
        _722_v55 = new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex((new BigNumber(902)).plus(_696_i3), new BigNumber((p0).length)), _723_v56)).length);
      }
      let _724_v57;
      _724_v57 = new _dafny.CodePoint('a'.codePointAt(0));
      let _725_v58;
      _725_v58 = _dafny.MultiSet.fromElements(_724_v57);
      let _726_v59;
      _726_v59 = _dafny.Seq.of(_658_v0);
      let _727_v60;
      _727_v60 = new BigNumber(976);
      let _728_v61;
      _728_v61 = _dafny.Seq.of((_this).f14, (_this).f14);
      let _729_v62;
      _729_v62 = _dafny.Map.Empty.slice().updateUnsafe(_658_v0,_727_v60);
      let _rhs107 = !((new BigNumber(-758)).isEqualTo(_module.__default.fm11(!(_658_v0), globalState)));
      let _rhs108 = ((_725_v58).Intersect(_dafny.MultiSet.fromElements(_724_v57))).Difference(_725_v58);
      let _rhs109 = _module.__default.fm27((_this).f14, _727_v60, _module.__default.safeModuloInt(new BigNumber((_728_v61).length), new BigNumber((_728_v61).length)), globalState);
      let _rhs110 = _658_v0;
      let _rhs111 = (new BigNumber((_729_v62).length)).multipliedBy((_this).f14);
      let _lhs74 = globalState;
      let _lhs75 = globalState;
      _lhs74.f13 = _rhs107;
      _725_v58 = _rhs108;
      _726_v59 = _rhs109;
      _lhs75.f13 = _rhs110;
      _727_v60 = _rhs111;
      let _730_v63;
      _730_v63 = _module.D0.create_DC1(_658_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), ((_731_v57) => function (_732_i5) {
  return _dafny.Seq.Create(_module.__default.abs(new BigNumber(104)), ((_733_v57) => function (_734_i6) {
    return _733_v57;
  })(_731_v57));
})(_724_v57)), p0);
      let _735_v64;
      _735_v64 = _dafny.MultiSet.fromElements(_730_v63, _730_v63, _module.D0.create_DC1(_658_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(67)), ((_736_p0) => function (_737_i7) {
  return _736_p0;
})(p0)), p0));
      let _738_v65;
      _738_v65 = _module.D2.create_DC6(_735_v64);
      let _pat_let_tv19 = _735_v64;
      let _pat_let_tv20 = _735_v64;
      let _source11 = function (_pat_let17_0) {
        return function (_739_dt__update__tmp_h1) {
          return function (_pat_let18_0) {
            return function (_740_dt__update_hcf10_h0) {
              return _module.D2.create_DC6(_740_dt__update_hcf10_h0);
            }(_pat_let18_0);
          }((_pat_let_tv19).Union(_pat_let_tv20));
        }(_pat_let17_0);
      }(_738_v65);
      if (_source11.is_DC7) {
        let _741___mcc_h0 = (_source11).cf11;
        let _742___mcc_h1 = (_source11).cf12;
        let _743___mcc_h2 = (_source11).cf13;
        let _744___mcc_h3 = (_source11).cf14;
        let _745_cf14 = _744___mcc_h3;
        let _746_cf13 = _743___mcc_h2;
        let _747_cf12 = _742___mcc_h1;
        let _748_cf11 = _741___mcc_h0;
        let _749_v66;
        _749_v66 = _dafny.Seq.of(_730_v63);
        let _750_v67;
        let _nw90 = new _module.C0();
        _nw90.__ctor((_735_v64).Union((_dafny.MultiSet.FromArray(_749_v66)).update(_730_v63, _module.__default.abs(_module.__default.fm11(_748_cf11, globalState)))), _748_cf11);
        _750_v67 = _nw90;
        _727_v60 = (_this).f14;
        let _751_v68;
        _751_v68 = _dafny.Map.Empty.slice().updateUnsafe(_746_cf13,_module.__default.safeDivisionInt((_dafny.ZERO).minus(_727_v60), new BigNumber(-627)));
        _751_v68 = (_751_v68).update(_727_v60, _727_v60);
        let _752_v69;
        _752_v69 = _module.D2.create_DC7(_658_v0, _747_cf12, _727_v60, _658_v0);
        let _753_v70;
        _753_v70 = _module.D0.create_DC2(!((_750_v67).f19));
        let _754_v71;
        _754_v71 = _module.D0.create_DC3(_module.D0.create_DC3(_753_v70));
        let _755_v72;
        _755_v72 = _dafny.Seq.of((_this).fm7(p0, _745_cf14, _748_cf11, _dafny.Map.Empty.slice().updateUnsafe(_745_cf14,_748_cf11), globalState));
        let _756_v73;
        let _nw91 = Array((new BigNumber(20)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = new BigNumber((_module.__default.fm10(_745_cf14, _746_cf13, _658_v0, (_750_v67).f19, globalState)).length);
        _nw91[(_dafny.ONE).toNumber()] = _727_v60;
        _nw91[(new BigNumber(2)).toNumber()] = _727_v60;
        _nw91[(new BigNumber(3)).toNumber()] = _727_v60;
        _nw91[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(((_745_cf14) ? (_746_cf13) : (_727_v60)));
        _nw91[(new BigNumber(5)).toNumber()] = _746_cf13;
        _nw91[(new BigNumber(6)).toNumber()] = (_746_cf13).multipliedBy(new BigNumber((_728_v61).length));
        _nw91[(new BigNumber(7)).toNumber()] = _module.__default.fm22(_752_v69, _754_v71, false, globalState);
        _nw91[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(405)), ((_757_v57) => function (_758_i8) {
          return _757_v57;
        })(_724_v57)))).length);
        _nw91[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt((_this).f14, _746_cf13);
        _nw91[(new BigNumber(10)).toNumber()] = new BigNumber((_755_v72).length);
        _nw91[(new BigNumber(11)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("sxb")).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_726_v59).length),(_750_v67).f19)).length));
        _nw91[(new BigNumber(12)).toNumber()] = _727_v60;
        _nw91[(new BigNumber(13)).toNumber()] = (_727_v60).minus((_this).f14);
        _nw91[(new BigNumber(14)).toNumber()] = _727_v60;
        _nw91[(new BigNumber(15)).toNumber()] = new BigNumber(884);
        _nw91[(new BigNumber(16)).toNumber()] = new BigNumber((_751_v68).length);
        _nw91[(new BigNumber(17)).toNumber()] = new BigNumber(945);
        _nw91[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(_746_cf13, _727_v60);
        _nw91[(new BigNumber(19)).toNumber()] = (_this).f14;
        _756_v73 = _nw91;
        let _index93 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_756_v73).length));
        (_756_v73)[_index93] = (_this).f14;
        let _759_v74;
        _759_v74 = _module.D1.create_DC4(_dafny.Map.Empty.slice().updateUnsafe(_745_cf14,_756_v73));
        let _760_v75;
        _760_v75 = _dafny.Map.Empty.slice().updateUnsafe(_745_cf14,_756_v73);
        let _761_v76;
        _761_v76 = _module.D6.create_DC15(_dafny.MultiSet.fromElements(_748_cf11, (_750_v67).f19));
        let _762_v77;
        _762_v77 = _module.D4.create_DC12(new BigNumber(((_761_v76).dtor_cf23).cardinality()));
        let _index94 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_756_v73).length));
        let _rhs112 = _module.__default.safeDivisionInt((_727_v60).minus(new BigNumber(42)), (_727_v60).multipliedBy((_this).f14));
        let _rhs113 = true;
        let _rhs114 = _module.D1.create_DC4(_760_v75);
        let _rhs115 = (_dafny.ZERO).minus((_762_v77).dtor_cf18);
        let _lhs76 = _756_v73;
        let _lhs77 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_756_v73).length));
        let _lhs78 = globalState;
        _lhs76[_lhs77] = _rhs112;
        _lhs78.f13 = _rhs113;
        _759_v74 = _rhs114;
        _727_v60 = _rhs115;
      } else if (_source11.is_DC6) {
        let _763___mcc_h4 = (_source11).cf10;
        let _764_cf10 = _763___mcc_h4;
        let _765_v78;
        let _nw92 = Array((new BigNumber(5)).toNumber()).fill(false);
        _765_v78 = _nw92;
        let _index95 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_765_v78).length));
        (_765_v78)[_index95] = false;
        let _766_v79;
        _766_v79 = _module.D1.create_DC5(_658_v0, p0, _726_v59);
        let _767_v80;
        _767_v80 = _dafny.Seq.of(_766_v79);
        let _768_v81;
        _768_v81 = _dafny.MultiSet.fromElements(_dafny.Seq.contains(_767_v80, _766_v79), _658_v0);
        let _769_v82;
        let _init9 = ((_770_p0) => function (_771_i9) {
          return _module.__default.safeModuloInt(_771_i9, new BigNumber((_770_p0).length));
        })(p0);
        let _nw93 = Array((new BigNumber(23)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw93.length); _i0_9++) {
          _nw93[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _769_v82 = _nw93;
        let _index96 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_769_v82).length));
        (_769_v82)[_index96] = (_this).f14;
        let _index97 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_765_v78).length));
        let _index98 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_769_v82).length));
        let _rhs116 = _658_v0;
        let _rhs117 = (_768_v81).Difference(_768_v81);
        let _rhs118 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(661)), function (_772_i10) {
          return (_this).f14;
        }), _728_v61);
        let _rhs119 = (_this).f14;
        let _rhs120 = !(_658_v0);
        let _lhs79 = _765_v78;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_765_v78).length));
        let _lhs81 = globalState;
        let _lhs82 = _769_v82;
        let _lhs83 = _module.__default.safeIndex(new BigNumber(217), new BigNumber((_769_v82).length));
        let _lhs84 = globalState;
        _lhs79[_lhs80] = _rhs116;
        _768_v81 = _rhs117;
        _lhs81.f13 = _rhs118;
        _lhs82[_lhs83] = _rhs119;
        _lhs84.f13 = _rhs120;
        (globalState).f5 = (_765_v78)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_765_v78).length))];
        let _773_v83;
        _773_v83 = _dafny.Seq.UnicodeFromString("umgbbbrln");
        _773_v83 = _dafny.Seq.Concat(p0, _773_v83);
        let _774_v84;
        _774_v84 = _module.D6.create_DC16((_this).f14);
        (globalState).f5 = (_727_v60).isLessThanOrEqualTo((_774_v84).dtor_cf24);
      } else {
        let _775___mcc_h5 = (_source11).cf15;
        let _776_cf15 = _775___mcc_h5;
        _724_v57 = ((!((_module.__default.fm11(_658_v0, globalState)).isLessThanOrEqualTo(_727_v60))) ? (_724_v57) : (_724_v57));
        (globalState).f13 = !(_658_v0);
        let _777_v85;
        let _nw94 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
        _777_v85 = _nw94;
        let _778_v86;
        let _init10 = function (_779_i11) {
          return true;
        };
        let _nw95 = Array((new BigNumber(16)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw95.length); _i0_10++) {
          _nw95[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _778_v86 = _nw95;
        let _index99 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_777_v85).length));
        (_777_v85)[_index99] = _dafny.Map.Empty.slice().updateUnsafe(_778_v86,_658_v0);
        let _780_v87;
        let _init11 = ((_781_v0) => function (_782_i12) {
          return (_dafny.Set.fromElements(_781_v0)).Intersect(_dafny.Set.fromElements(_781_v0));
        })(_658_v0);
        let _nw96 = Array((new BigNumber(12)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw96.length); _i0_11++) {
          _nw96[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _780_v87 = _nw96;
        let _783_v88;
        _783_v88 = _dafny.Set.fromElements(_658_v0, _658_v0, _658_v0, _658_v0, _658_v0);
        let _index100 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_780_v87).length));
        (_780_v87)[_index100] = _783_v88;
        let _784_v89;
        _784_v89 = _dafny.MultiSet.fromElements(_658_v0);
        let _785_v90;
        _785_v90 = _dafny.Map.Empty.slice().updateUnsafe(_778_v86,(_726_v59)[_module.__default.safeIndex((_this).f14, new BigNumber((_726_v59).length))]);
        let _786_v91;
        _786_v91 = _dafny.MultiSet.fromElements(_727_v60);
        let _787_v92;
        _787_v92 = _dafny.MultiSet.fromElements(new BigNumber((_786_v91).cardinality()));
        let _index101 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_777_v85).length));
        let _index102 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_780_v87).length));
        let _rhs121 = ((_658_v0) ? (_dafny.Seq.update(_dafny.Seq.Concat(_728_v61, _728_v61), _module.__default.safeIndex(new BigNumber(138), new BigNumber((_dafny.Seq.Concat(_728_v61, _728_v61)).length)), (_this).f14)) : (_728_v61));
        let _rhs122 = (_785_v90).Merge(_785_v90);
        let _rhs123 = ((_module.__default.fm33(_658_v0, _658_v0, (_726_v59)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,p0)).length), new BigNumber((_726_v59).length))], globalState)).Difference(_783_v88)).Union((_783_v88).Union(_dafny.Set.fromElements(false)));
        let _rhs124 = _dafny.MultiSet.FromArray(_726_v59);
        let _rhs125 = ((((_786_v91).update(_727_v60, _module.__default.abs(_727_v60))).IsSubsetOf(_787_v92)) ? (_778_v86) : (_778_v86));
        let _lhs85 = _777_v85;
        let _lhs86 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_777_v85).length));
        let _lhs87 = _780_v87;
        let _lhs88 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_780_v87).length));
        _728_v61 = _rhs121;
        _lhs85[_lhs86] = _rhs122;
        _lhs87[_lhs88] = _rhs123;
        _784_v89 = _rhs124;
        _778_v86 = _rhs125;
        let _index103 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_778_v86).length));
        (_778_v86)[_index103] = _658_v0;
        let _index104 = _module.__default.safeIndex(new BigNumber(250), new BigNumber((_778_v86).length));
        (_778_v86)[_index104] = _658_v0;
      }
      let _788_i13;
      _788_i13 = _dafny.ZERO;
      L4: {
        while (!(!(_658_v0) || (!(_658_v0) || (_658_v0)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_788_i13)) {
              break L4;
            }
            _788_i13 = (_788_i13).plus(_dafny.ONE);
            let _789_v93;
            _789_v93 = _dafny.Seq.UnicodeFromString("hglx");
            _789_v93 = _dafny.Seq.update(_789_v93, _module.__default.safeIndex(((_dafny.ZERO).minus((_this).f14)).plus((_this).f14), new BigNumber((_789_v93).length)), _724_v57);
            _727_v60 = new BigNumber((p0).length);
            let _790_v94;
            _790_v94 = _module.D1.create_DC5((_this).fm8((_this).f14, _658_v0, _658_v0, globalState), p0, _726_v59);
            let _791_v95;
            let _nw97 = new _module.C0();
            _nw97.__ctor(_735_v64, (_790_v94).dtor_cf7);
            _791_v95 = _nw97;
            _726_v59 = _dafny.Seq.Concat(_dafny.Seq.Concat(_726_v59, _dafny.Seq.of(_658_v0, _658_v0)), _726_v59);
          }
        }
      }
      return;
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

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return ((new BigNumber((function () {
        let _coll23 = new _dafny.Set();
        for (const _compr_23 of _dafny.IntegerRange(new BigNumber(947), new BigNumber(745))) {
          let _792_v0 = _compr_23;
          if (((new BigNumber(947)).isLessThanOrEqualTo(_792_v0)) && ((_792_v0).isLessThan(new BigNumber(745)))) {
            _coll23.add(_module.__default.safeModuloInt(_792_v0, new BigNumber((_dafny.Seq.of(false)).length)));
          }
        }
        return _coll23;
      }()).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).plus(new BigNumber(274));
    };
    fm1(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(714)), function (_793_i0) {
        return ((false) ? (new _dafny.CodePoint('e'.codePointAt(0))) : (new _dafny.CodePoint('c'.codePointAt(0))));
      });
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(825),new BigNumber((_dafny.Set.fromElements(false)).length))).length), new BigNumber(764)), _dafny.Seq.of(new BigNumber(-875), new BigNumber(578))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(856)), _dafny.Seq.of(new BigNumber(-773), new BigNumber((function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of _dafny.IntegerRange(new BigNumber(-249), new BigNumber(694))) {
          let _794_v0 = _compr_24;
          if (((new BigNumber(-249)).isLessThanOrEqualTo(_794_v0)) && ((_794_v0).isLessThan(new BigNumber(694)))) {
            _coll24.push([(_794_v0).multipliedBy(new BigNumber((function () {
              let _coll25 = new _dafny.Map();
              for (const _compr_25 of (_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(636))).cardinality()))).Elements) {
                let _795_v1 = _compr_25;
                if ((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(636))).cardinality()))).contains(_795_v1)) {
                  _coll25.push([_module.__default.safeDivisionInt(_795_v1, new BigNumber((_dafny.Seq.of(true)).length)),new BigNumber(-483)]);
                }
              }
              return _coll25;
            }()).length)),new BigNumber(365)]);
          }
        }
        return _coll24;
      }()).length))));
    };
    fm5(p0, globalState) {
      let _this = this;
      return new BigNumber(732);
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _796_v0;
      _796_v0 = true;
      let _797_v1;
      _797_v1 = _dafny.Seq.of(_module.__default.fm6(globalState), _796_v0, _796_v0);
      let _798_v2;
      _798_v2 = new BigNumber(-350);
      let _799_v3;
      _799_v3 = _dafny.Seq.of(_797_v1, _dafny.Seq.update(_797_v1, _module.__default.safeIndex(_798_v2, new BigNumber((_797_v1).length)), _796_v0));
      let _800_v4;
      _800_v4 = _dafny.Seq.UnicodeFromString("dnhvpnw");
      let _801_v5;
      _801_v5 = _dafny.Seq.of(_800_v4, _800_v4);
      let _802_i0;
      _802_i0 = _dafny.ZERO;
      L5: {
        while (_dafny.areEqual(_797_v1, _dafny.Seq.Concat((_799_v3)[_module.__default.safeIndex(new BigNumber((_801_v5).length), new BigNumber((_799_v3).length))], _dafny.Seq.of(_796_v0)))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_802_i0)) {
              break L5;
            }
            _802_i0 = (_802_i0).plus(_dafny.ONE);
            let _803_v6;
            let _nw98 = new _module.C2();
            _nw98.__ctor();
            _803_v6 = _nw98;
            _798_v2 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_798_v2), _798_v2);
            if (_796_v0) {
              _798_v2 = _798_v2;
              r1 = _798_v2;
              let _804_v8;
              _804_v8 = _dafny.MultiSet.fromElements(_798_v2);
              let _805_v9;
              _805_v9 = _dafny.Map.Empty.slice().updateUnsafe(_798_v2,_798_v2);
              let _806_v10;
              _806_v10 = _dafny.Seq.of(_798_v2, (_798_v2).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), ((_807_v2) => function (_808_i1) {
                return function () {
                  let _coll26 = new _dafny.Map();
                  for (const _compr_26 of (_dafny.MultiSet.fromElements(new BigNumber(-454), _807_v2, new BigNumber(576), _807_v2, _807_v2)).Elements) {
                    let _809_v7 = _compr_26;
                    if ((_dafny.MultiSet.fromElements(new BigNumber(-454), _807_v2, new BigNumber(576), _807_v2, _807_v2)).contains(_809_v7)) {
                      _coll26.push([_module.__default.safeDivisionInt(_809_v7, _807_v2),_807_v2]);
                    }
                  }
                  return _coll26;
                }();
              })(_798_v2)), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_804_v8).cardinality())), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(866)), ((_810_v2) => function (_811_i1) {
                return function () {
                  let _coll27 = new _dafny.Map();
                  for (const _compr_27 of (_dafny.MultiSet.fromElements(new BigNumber(-454), _810_v2, new BigNumber(576), _810_v2, _810_v2)).Elements) {
                    let _812_v7 = _compr_27;
                    if ((_dafny.MultiSet.fromElements(new BigNumber(-454), _810_v2, new BigNumber(576), _810_v2, _810_v2)).contains(_812_v7)) {
                      _coll27.push([_module.__default.safeDivisionInt(_812_v7, _810_v2),_810_v2]);
                    }
                  }
                  return _coll27;
                }();
              })(_798_v2))).length)), _805_v9)).length)), new BigNumber((_800_v4).length));
              let _rhs126 = _796_v0;
              let _rhs127 = (_798_v2).multipliedBy(_798_v2);
              let _rhs128 = _806_v10;
              r0 = _rhs126;
              _798_v2 = _rhs127;
              _806_v10 = _rhs128;
              let _813_v11;
              _813_v11 = new _dafny.CodePoint('b'.codePointAt(0));
              _813_v11 = _813_v11;
              let _814_v12;
              let _init12 = ((_815_v2) => function (_816_i2) {
                return _module.__default.safeModuloInt(_816_i2, _815_v2);
              })(_798_v2);
              let _nw99 = Array((new BigNumber(20)).toNumber());
              for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw99.length); _i0_12++) {
                _nw99[_i0_12] = _init12(new BigNumber(_i0_12));
              }
              _814_v12 = _nw99;
              let _817_v13;
              _817_v13 = _dafny.Map.Empty.slice().updateUnsafe(_814_v12,_dafny.Seq.Create(_module.__default.abs(new BigNumber(805)), ((_818_p0) => function (_819_i3) {
                return _818_p0;
              })(p0)));
              let _820_v14;
              _820_v14 = _module.D0.create_DC1(false, _801_v5, (((_817_v13).contains(_814_v12)) ? ((_817_v13).get(_814_v12)) : (_800_v4)));
              let _821_v15;
              _821_v15 = _dafny.MultiSet.fromElements(_820_v14, _820_v14);
              let _822_v16;
              _822_v16 = _module.D2.create_DC6(_821_v15);
              let _823_v17;
              let _nw100 = new _module.C0();
              _nw100.__ctor((_822_v16).dtor_cf10, _796_v0);
              _823_v17 = _nw100;
            } else {
              _796_v0 = (_796_v0) || (true);
              (globalState).f13 = _796_v0;
              (globalState).f13 = _796_v0;
              _797_v1 = _797_v1;
              (globalState).f9 = !(_796_v0);
            }
            let _824_v18;
            _824_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-708),_796_v0);
            let _825_v20;
            _825_v20 = _dafny.Set.fromElements(_796_v0);
            let _826_v21;
            _826_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_825_v20).length),_798_v2);
            let _827_v22;
            _827_v22 = _module.D7.create_DC18(function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of (_826_v21).Keys.Elements) {
    let _828_v19 = _compr_28;
    if ((_826_v21).contains(_828_v19)) {
      _coll28.push([(_828_v19).multipliedBy(new BigNumber(523)),true]);
    }
  }
  return _coll28;
}());
            _824_v18 = (_827_v22).dtor_cf26;
          }
        }
      }
      if (_796_v0) {
        let _829_v23;
        let _init13 = ((_830_p0) => function (_831_i4) {
          return _830_p0;
        })(p0);
        let _nw101 = Array((new BigNumber(12)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw101.length); _i0_13++) {
          _nw101[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _829_v23 = _nw101;
        let _index105 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_829_v23).length));
        (_829_v23)[_index105] = p0;
        let _index106 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_829_v23).length));
        (_829_v23)[_index106] = p0;
        let _832_v24;
        _832_v24 = _module.D0.create_DC3(_module.D0.create_DC1(_796_v0, _801_v5, _800_v4));
        let _833_v25;
        _833_v25 = _dafny.Map.Empty.slice().updateUnsafe(_796_v0,_832_v24);
        let _834_v26;
        _834_v26 = _module.D2.create_DC7(_796_v0, _833_v25, _798_v2, _796_v0);
        let _835_v27;
        _835_v27 = _module.D0.create_DC1(true, _801_v5, _800_v4);
        _798_v2 = ((_798_v2).minus(_module.__default.fm22(_834_v26, _module.D0.create_DC3(_835_v27), _796_v0, globalState))).plus(_module.__default.fm22(_module.D2.create_DC7(_796_v0, _833_v25, new BigNumber(168), _796_v0), _832_v24, _796_v0, globalState));
        let _836_v28;
        let _nw102 = Array((new BigNumber(27)).toNumber());
        _nw102[(_dafny.ZERO).toNumber()] = _796_v0;
        _nw102[(_dafny.ONE).toNumber()] = _796_v0;
        _nw102[(new BigNumber(2)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(3)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(4)).toNumber()] = !(_796_v0);
        _nw102[(new BigNumber(5)).toNumber()] = !(_796_v0);
        _nw102[(new BigNumber(6)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(7)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(8)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(9)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(10)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(11)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(12)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(13)).toNumber()] = !(_796_v0);
        _nw102[(new BigNumber(14)).toNumber()] = !(_796_v0);
        _nw102[(new BigNumber(15)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(16)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(17)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(18)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(19)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(20)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(21)).toNumber()] = false;
        _nw102[(new BigNumber(22)).toNumber()] = true;
        _nw102[(new BigNumber(23)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(24)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(25)).toNumber()] = _796_v0;
        _nw102[(new BigNumber(26)).toNumber()] = _796_v0;
        _836_v28 = _nw102;
        let _837_v29;
        let _nw103 = Array((new BigNumber(16)).toNumber());
        _nw103[(_dafny.ZERO).toNumber()] = _836_v28;
        _nw103[(_dafny.ONE).toNumber()] = _836_v28;
        _nw103[(new BigNumber(2)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(3)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(4)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(5)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(6)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(7)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(8)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(9)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(10)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(11)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(12)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(13)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(14)).toNumber()] = _836_v28;
        _nw103[(new BigNumber(15)).toNumber()] = _836_v28;
        _837_v29 = _nw103;
        _837_v29 = _837_v29;
        let _838_v30;
        let _nw104 = new _module.C3();
        _nw104.__ctor(_796_v0, _798_v2);
        _838_v30 = _nw104;
        let _839_v31;
        _839_v31 = _dafny.Set.fromElements(_796_v0);
        let _840_v32;
        _840_v32 = _dafny.MultiSet.fromElements(_798_v2);
        (_this).m3((_798_v2).plus(_798_v2), _838_v30, new BigNumber(((_839_v31).Union(_dafny.Set.fromElements(_796_v0, _796_v0, _796_v0))).length), (_840_v32).Union(_840_v32), globalState);
        let _841_v33;
        _841_v33 = _839_v31;
        let _842_v34;
        _842_v34 = _dafny.Set.fromElements(_841_v33);
        let _843_v35;
        _843_v35 = _dafny.Map.Empty.slice().updateUnsafe(_842_v34,_838_v30);
        (globalState).f3 = !(!(_843_v35).equals(_843_v35));
      } else {
        let _844_v36;
        _844_v36 = _dafny.MultiSet.fromElements(_796_v0);
        let _845_v37;
        _845_v37 = _dafny.Seq.of((((_844_v36).contains(_796_v0)) ? ((_844_v36).get(_796_v0)) : (_798_v2)), _798_v2, _798_v2, _798_v2);
        _798_v2 = _module.__default.safeModuloInt(_798_v2, ((((_844_v36).contains(_796_v0)) ? ((_844_v36).get(_796_v0)) : (new BigNumber((_845_v37).length)))).plus(_module.__default.fm11(_module.__default.fm6(globalState), globalState)));
        r0 = _796_v0;
        let _846_v38;
        _846_v38 = new _dafny.CodePoint('e'.codePointAt(0));
        let _847_v39;
        let _init14 = ((_848_v2) => function (_849_i5) {
          return (_849_i5).minus(_848_v2);
        })(_798_v2);
        let _nw105 = Array((new BigNumber(26)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw105.length); _i0_14++) {
          _nw105[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _847_v39 = _nw105;
        let _index107 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_847_v39).length));
        (_847_v39)[_index107] = new BigNumber((_dafny.Seq.Concat(_800_v4, _800_v4)).length);
        let _index108 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_847_v39).length));
        (_847_v39)[_index108] = (_845_v37)[_module.__default.safeIndex(_798_v2, new BigNumber((_845_v37).length))];
        let _850_v40;
        _850_v40 = _module.D0.create_DC2(_796_v0);
        let _851_v41;
        _851_v41 = _dafny.Seq.of(_850_v40);
        let _852_v42;
        _852_v42 = _dafny.Map.Empty.slice().updateUnsafe(_798_v2,!(false));
        let _index109 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_847_v39).length));
        let _index110 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_847_v39).length));
        let _rhs129 = _846_v38;
        let _rhs130 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ebekj"), (_801_v5)[_module.__default.safeIndex(_798_v2, new BigNumber((_801_v5).length))])).length);
        let _rhs131 = new BigNumber((_797_v1).length);
        let _rhs132 = (new BigNumber(-52)).minus(new BigNumber((_852_v42).length));
        let _rhs133 = _851_v41;
        let _lhs89 = _847_v39;
        let _lhs90 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_847_v39).length));
        let _lhs91 = _847_v39;
        let _lhs92 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_847_v39).length));
        _846_v38 = _rhs129;
        _lhs89[_lhs90] = _rhs130;
        _lhs91[_lhs92] = _rhs131;
        r1 = _rhs132;
        _851_v41 = _rhs133;
        r1 = _798_v2;
        let _853_v43;
        let _nw106 = new _module.C2();
        _nw106.__ctor();
        _853_v43 = _nw106;
      }
      if ((_796_v0) && (_796_v0)) {
        let _854_v44;
        let _nw107 = Array((new BigNumber(19)).toNumber());
        _nw107[(_dafny.ZERO).toNumber()] = p0;
        _nw107[(_dafny.ONE).toNumber()] = p0;
        _nw107[(new BigNumber(2)).toNumber()] = p0;
        _nw107[(new BigNumber(3)).toNumber()] = p0;
        _nw107[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
        _nw107[(new BigNumber(5)).toNumber()] = p0;
        _nw107[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw107[(new BigNumber(7)).toNumber()] = p0;
        _nw107[(new BigNumber(8)).toNumber()] = p0;
        _nw107[(new BigNumber(9)).toNumber()] = p0;
        _nw107[(new BigNumber(10)).toNumber()] = p0;
        _nw107[(new BigNumber(11)).toNumber()] = p0;
        _nw107[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw107[(new BigNumber(13)).toNumber()] = p0;
        _nw107[(new BigNumber(14)).toNumber()] = p0;
        _nw107[(new BigNumber(15)).toNumber()] = p0;
        _nw107[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
        _nw107[(new BigNumber(17)).toNumber()] = p0;
        _nw107[(new BigNumber(18)).toNumber()] = p0;
        _854_v44 = _nw107;
        let _index111 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_854_v44).length));
        (_854_v44)[_index111] = p0;
        let _index112 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_854_v44).length));
        (_854_v44)[_index112] = p0;
        let _855_v45;
        let _init15 = ((_856_v2) => function (_857_i6) {
          return _module.__default.safeDivisionInt(_857_i6, _856_v2);
        })(_798_v2);
        let _nw108 = Array((new BigNumber(28)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw108.length); _i0_15++) {
          _nw108[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _855_v45 = _nw108;
        let _858_v46;
        _858_v46 = _dafny.Set.fromElements(_796_v0, _module.__default.fm6(globalState), _796_v0, !(_796_v0), _796_v0);
        let _index113 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_855_v45).length));
        (_855_v45)[_index113] = new BigNumber((_858_v46).length);
        let _859_v47;
        _859_v47 = _dafny.Set.fromElements(_798_v2);
        let _860_v48;
        _860_v48 = _dafny.Map.Empty.slice().updateUnsafe(_798_v2,new BigNumber((_859_v47).length));
        let _index114 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_855_v45).length));
        (_855_v45)[_index114] = (new BigNumber((_860_v48).length)).multipliedBy(_798_v2);
        let _861_v49;
        let _nw109 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Set.Empty);
        _861_v49 = _nw109;
        let _862_v50;
        _862_v50 = _dafny.Set.fromElements(_855_v45);
        let _index115 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_861_v49).length));
        (_861_v49)[_index115] = _862_v50;
        let _index116 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_861_v49).length));
        (_861_v49)[_index116] = _862_v50;
        _798_v2 = (((_860_v48).contains((_dafny.ZERO).minus((_855_v45)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_855_v45).length))]))) ? ((_860_v48).get((_dafny.ZERO).minus((_855_v45)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_855_v45).length))]))) : ((_798_v2).minus(_798_v2)));
        r1 = (_855_v45)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_855_v45).length))];
      } else {
        (globalState).f9 = ((_796_v0) ? (_796_v0) : (_796_v0));
        let _863_v51;
        _863_v51 = _dafny.Set.fromElements(!(false), _796_v0, _796_v0);
        let _864_v52;
        _864_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_863_v51).length),_798_v2);
        _864_v52 = (_864_v52).update((new BigNumber(151)).multipliedBy(new BigNumber(147)), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_798_v2, _798_v2)));
        let _865_v53;
        _865_v53 = _dafny.Seq.of(new BigNumber((_800_v4).length));
        let _866_v54;
        _866_v54 = _dafny.MultiSet.fromElements(_865_v53, _dafny.Seq.of(_798_v2, _798_v2, _798_v2, _798_v2));
        let _867_v55;
        _867_v55 = _module.D1.create_DC5(_796_v0, _800_v4, _797_v1);
        let _868_v56;
        let _nw110 = Array((new BigNumber(24)).toNumber());
        _nw110[(_dafny.ZERO).toNumber()] = _800_v4;
        _nw110[(_dafny.ONE).toNumber()] = _800_v4;
        _nw110[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-339)), ((_869_p0) => function (_870_i7) {
          return _869_p0;
        })(p0));
        _nw110[(new BigNumber(3)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(932)), ((_871_p0) => function (_872_i8) {
          return _871_p0;
        })(p0));
        _nw110[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-895)), ((_873_p0) => function (_874_i9) {
          return _873_p0;
        })(p0));
        _nw110[(new BigNumber(6)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(7)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("rl");
        _nw110[(new BigNumber(9)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-2)), ((_875_p0) => function (_876_i10) {
          return _875_p0;
        })(p0));
        _nw110[(new BigNumber(11)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(12)).toNumber()] = _module.__default.fm20(new BigNumber((_865_v53).length), _796_v0, _796_v0, _798_v2, globalState);
        _nw110[(new BigNumber(13)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(14)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(15)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(16)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(17)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(18)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(19)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(20)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(21)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(22)).toNumber()] = _800_v4;
        _nw110[(new BigNumber(23)).toNumber()] = _800_v4;
        _868_v56 = _nw110;
        let _877_v57;
        _877_v57 = _module.D7.create_DC19(_867_v55, _798_v2, _796_v0, _868_v56);
        let _878_v58;
        _878_v58 = _module.D0.create_DC2(false);
        let _879_v59;
        _879_v59 = _dafny.Seq.of(_865_v53, _865_v53, _dafny.Seq.update(_dafny.Seq.update(_865_v53, _module.__default.safeIndex(_798_v2, new BigNumber((_865_v53).length)), new BigNumber(927)), _module.__default.safeIndex((_877_v57).dtor_cf28, new BigNumber((_dafny.Seq.update(_865_v53, _module.__default.safeIndex(_798_v2, new BigNumber((_865_v53).length)), new BigNumber(927))).length)), (_this).fm5(_878_v58, globalState)), _dafny.Seq.of(_798_v2), _865_v53);
        if (((_866_v54).Intersect(_module.__default.fm34(globalState))).IsDisjointFrom(_dafny.MultiSet.FromArray(_879_v59))) {
          let _880_v60;
          _880_v60 = _dafny.Map.Empty.slice().updateUnsafe(_796_v0,_798_v2);
          _880_v60 = (_880_v60).update(true, _798_v2);
          let _881_v61;
          _881_v61 = _module.D0.create_DC1(_796_v0, _801_v5, _800_v4);
          let _882_v62;
          _882_v62 = _dafny.MultiSet.fromElements(_881_v61);
          let _883_v63;
          let _nw111 = new _module.C0();
          _nw111.__ctor(_882_v62, _796_v0);
          _883_v63 = _nw111;
          let _884_v64;
          let _nw112 = Array((new BigNumber(27)).toNumber()).fill(false);
          _884_v64 = _nw112;
          let _885_v65;
          let _886_v66;
          let _887_v67;
          let _out11;
          let _out12;
          let _out13;
          let _outcollector4 = (_this).m4(_865_v53, _884_v64, globalState);
          _out11 = _outcollector4[0];
          _out12 = _outcollector4[1];
          _out13 = _outcollector4[2];
          _885_v65 = _out11;
          _886_v66 = _out12;
          _887_v67 = _out13;
          let _index117 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_884_v64).length));
          (_884_v64)[_index117] = ((_885_v65) ? (_796_v0) : (_796_v0));
          let _index118 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_884_v64).length));
          (_884_v64)[_index118] = _885_v65;
          let _888_v68;
          let _nw113 = Array((new BigNumber(21)).toNumber()).fill([]);
          _888_v68 = _nw113;
          let _index119 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_888_v68).length));
          (_888_v68)[_index119] = _884_v64;
          let _index120 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_888_v68).length));
          (_888_v68)[_index120] = _884_v64;
        } else {
          let _889_v69;
          _889_v69 = _dafny.MultiSet.fromElements(_798_v2);
          (globalState).f5 = !((_dafny.MultiSet.FromArray(_865_v53)).IsSubsetOf(_889_v69));
          (globalState).f13 = ((_dafny.ZERO).minus(_798_v2)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(-799)));
          let _890_v70;
          let _nw114 = Array((new BigNumber(12)).toNumber());
          _nw114[(_dafny.ZERO).toNumber()] = ((_796_v0) ? (new BigNumber((_dafny.Seq.update(_800_v4, _module.__default.safeIndex(_798_v2, new BigNumber((_800_v4).length)), p0)).length)) : (_798_v2));
          _nw114[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_797_v1, _dafny.Seq.of(_796_v0))).length);
          _nw114[(new BigNumber(2)).toNumber()] = (_798_v2).plus(_798_v2);
          _nw114[(new BigNumber(3)).toNumber()] = _798_v2;
          _nw114[(new BigNumber(4)).toNumber()] = _798_v2;
          _nw114[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_798_v2);
          _nw114[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_798_v2);
          _nw114[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(_798_v2, _798_v2);
          _nw114[(new BigNumber(8)).toNumber()] = _798_v2;
          _nw114[(new BigNumber(9)).toNumber()] = _798_v2;
          _nw114[(new BigNumber(10)).toNumber()] = (_798_v2).plus(_798_v2);
          _nw114[(new BigNumber(11)).toNumber()] = new BigNumber((_800_v4).length);
          _890_v70 = _nw114;
          let _index121 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_890_v70).length));
          (_890_v70)[_index121] = (_798_v2).plus(_798_v2);
          let _891_v71;
          let _nw115 = Array((new BigNumber(11)).toNumber()).fill(false);
          _891_v71 = _nw115;
          let _index122 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_891_v71).length));
          (_891_v71)[_index122] = !(!(!(_796_v0)));
          let _index123 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_890_v70).length));
          let _index124 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_891_v71).length));
          let _rhs134 = _798_v2;
          let _rhs135 = _796_v0;
          let _rhs136 = _796_v0;
          let _lhs93 = _890_v70;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_890_v70).length));
          let _lhs95 = _891_v71;
          let _lhs96 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_891_v71).length));
          let _lhs97 = globalState;
          _lhs93[_lhs94] = _rhs134;
          _lhs95[_lhs96] = _rhs135;
          _lhs97.f9 = _rhs136;
          let _892_v72;
          let _init16 = ((_893_v53) => function (_894_i11) {
            return _893_v53;
          })(_865_v53);
          let _nw116 = Array((new BigNumber(11)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw116.length); _i0_16++) {
            _nw116[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _892_v72 = _nw116;
          let _index125 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_892_v72).length));
          (_892_v72)[_index125] = _dafny.Seq.Concat(_865_v53, _865_v53);
          let _895_v73;
          _895_v73 = new _dafny.CodePoint('o'.codePointAt(0));
          let _index126 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_892_v72).length));
          let _rhs137 = _dafny.Seq.update(_dafny.Seq.Concat(_865_v53, ((true) ? (_865_v53) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(735)), ((_896_v71, _897_v2) => function (_898_i12) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_896_v71)[_module.__default.safeIndex(new BigNumber(83), new BigNumber((_896_v71).length))],_897_v2)).length);
          })(_891_v71, _798_v2))))), _module.__default.safeIndex((_dafny.ZERO).minus(_798_v2), new BigNumber((_dafny.Seq.Concat(_865_v53, ((true) ? (_865_v53) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(735)), ((_899_v71, _900_v2) => function (_901_i12) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_899_v71)[_module.__default.safeIndex(new BigNumber(83), new BigNumber((_899_v71).length))],_900_v2)).length);
          })(_891_v71, _798_v2)))))).length)), (_890_v70)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_890_v70).length))]);
          let _rhs138 = (_890_v70)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_890_v70).length))];
          let _rhs139 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_800_v4, _module.__default.safeIndex(_798_v2, new BigNumber((_800_v4).length)), p0), _800_v4), _dafny.Seq.UnicodeFromString("wje"));
          let _rhs140 = (_800_v4)[_module.__default.safeIndex(new BigNumber((_800_v4).length), new BigNumber((_800_v4).length))];
          let _lhs98 = _892_v72;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_892_v72).length));
          _lhs98[_lhs99] = _rhs137;
          _798_v2 = _rhs138;
          _800_v4 = _rhs139;
          _895_v73 = _rhs140;
          _895_v73 = p0;
        }
        let _902_v74;
        _902_v74 = _module.D0.create_DC1(_796_v0, _801_v5, _800_v4);
        let _903_v75;
        _903_v75 = _dafny.MultiSet.fromElements(_902_v74, _902_v74, _902_v74, _902_v74, _902_v74);
        let _904_v76;
        let _nw117 = new _module.C0();
        _nw117.__ctor(_903_v75, !(_796_v0));
        _904_v76 = _nw117;
        let _905_v77;
        _905_v77 = new _dafny.CodePoint('s'.codePointAt(0));
        _905_v77 = _905_v77;
      }
      let _906_v78;
      _906_v78 = _dafny.Set.fromElements(_796_v0, _796_v0, _796_v0);
      _906_v78 = (_906_v78).Intersect(_906_v78);
      let _907_v79;
      _907_v79 = _module.D0.create_DC2(_796_v0);
      let _908_v80;
      let _nw118 = new _module.C4();
      _nw118.__ctor(_798_v2, (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),(_this).fm5(_907_v79, globalState))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_798_v2)));
      _908_v80 = _nw118;
      let _909_i13;
      _909_i13 = _dafny.ZERO;
      L6: {
        while (_796_v0) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_909_i13)) {
              break L6;
            }
            _909_i13 = (_909_i13).plus(_dafny.ONE);
            let _910_v81;
            _910_v81 = _dafny.MultiSet.fromElements((_908_v80).f14, ((_908_v80).f14).multipliedBy(new BigNumber((_797_v1).length)));
            _910_v81 = _910_v81;
            _798_v2 = (_dafny.ZERO).minus(_module.__default.fm11(_796_v0, globalState));
            _798_v2 = (_908_v80).f14;
            (globalState).f5 = _796_v0;
          }
        }
      }
      r0 = (_796_v0) || (false);
      r1 = (_dafny.ZERO).minus((_798_v2).multipliedBy(new BigNumber((((_908_v80).f15).Merge((_908_v80).f15)).length)));
      return [r0, r1];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _911_v0;
      let _nw119 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _911_v0 = _nw119;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_911_v0).length))) {
        let _912_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_912_i0)) && ((_912_i0).isLessThan(new BigNumber((_911_v0).length))))) {
          (_911_v0)[(_912_i0)] = ((true) ? (new _dafny.CodePoint('y'.codePointAt(0))) : (new _dafny.CodePoint('a'.codePointAt(0))));
        }
      }
      let _913_v1;
      let _nw120 = Array((new BigNumber(28)).toNumber()).fill(false);
      _913_v1 = _nw120;
      let _914_v2;
      _914_v2 = true;
      let _index127 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
      (_913_v1)[_index127] = !(!(_914_v2));
      let _915_v3;
      _915_v3 = new BigNumber(-355);
      let _916_v4;
      _916_v4 = _dafny.Seq.UnicodeFromString("hlo");
      let _917_v5;
      _917_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(504),true);
      let _918_v6;
      _918_v6 = _module.D7.create_DC18(_917_v5);
      let _919_v7;
      _919_v7 = _module.D7.create_DC21(_918_v6);
      let _920_v8;
      _920_v8 = _dafny.Map.Empty.slice().updateUnsafe(_915_v3,_919_v7);
      let _921_v9;
      _921_v9 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_914_v2,_920_v8)).length));
      let _922_v10;
      _922_v10 = _dafny.Set.fromElements(_914_v2);
      let _index128 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
      let _rhs141 = !(!(_921_v9).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), ((_923_p0) => function (_924_i1) {
        return _923_p0;
      })(p0))).length)));
      let _rhs142 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))).multipliedBy(new BigNumber((((_914_v2) ? (_922_v10) : (_922_v10))).length));
      let _rhs143 = _914_v2;
      let _rhs144 = _916_v4;
      let _lhs100 = _913_v1;
      let _lhs101 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
      _lhs100[_lhs101] = _rhs141;
      _915_v3 = _rhs142;
      _914_v2 = _rhs143;
      _916_v4 = _rhs144;
      _915_v3 = (((p3).contains(p2)) ? ((p3).get(p2)) : (_module.__default.safeDivisionInt(p0, _915_v3)));
      (globalState).f3 = _914_v2;
      (globalState).f9 = !((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]) || (_914_v2);
      let _925_v11;
      _925_v11 = _module.D0.create_DC2(_914_v2);
      let _926_v12;
      _926_v12 = _module.D0.create_DC3(_925_v11);
      let _source12 = _926_v12;
      if (_source12.is_DC1) {
        let _927___mcc_h0 = (_source12).cf1;
        let _928___mcc_h1 = (_source12).cf2;
        let _929___mcc_h2 = (_source12).cf3;
        let _930_cf3 = _929___mcc_h2;
        let _931_cf2 = _928___mcc_h1;
        let _932_cf1 = _927___mcc_h0;
        _915_v3 = _module.__default.safeDivisionInt(_915_v3, (_dafny.ZERO).minus(p2));
        _915_v3 = ((_dafny.ZERO).minus(p0)).multipliedBy((p2).multipliedBy(p0));
        let _933_v13;
        _933_v13 = _module.D4.create_DC11();
        let _source13 = _933_v13;
        if (_source13.is_DC11) {
          let _934_v14;
          let _nw121 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _934_v14 = _nw121;
          let _index129 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_934_v14).length));
          (_934_v14)[_index129] = new BigNumber((_dafny.Seq.UnicodeFromString("lklrsh")).length);
          let _index130 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_934_v14).length));
          (_934_v14)[_index130] = p0;
          let _935_v15;
          _935_v15 = new _dafny.CodePoint('n'.codePointAt(0));
          let _936_v16;
          _936_v16 = _dafny.MultiSet.fromElements(_935_v15);
          _915_v3 = new BigNumber((_936_v16).cardinality());
          let _937_v17;
          _937_v17 = _dafny.Map.Empty.slice().updateUnsafe(_932_cf1,_933_v13);
          let _938_v18;
          let _nw122 = new _module.C3();
          _nw122.__ctor((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))], (((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]) ? (p2) : (new BigNumber((_937_v17).length))));
          _938_v18 = _nw122;
          _921_v9 = _module.__default.fm10((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))], p2, _932_cf1, (p2).isLessThanOrEqualTo(p2), globalState);
        } else if (_source13.is_DC12) {
          let _939___mcc_h6 = (_source13).cf18;
          let _940_cf18 = _939___mcc_h6;
          _921_v9 = _921_v9;
          let _941_v19;
          let _nw123 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _941_v19 = _nw123;
          let _index131 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_941_v19).length));
          (_941_v19)[_index131] = (_940_cf18).minus(new BigNumber(-548));
          let _index132 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_941_v19).length));
          (_941_v19)[_index132] = (_dafny.ZERO).minus(p0);
          (globalState).f13 = (_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))];
          let _942_v20;
          _942_v20 = _dafny.Seq.of(_932_cf1);
          let _943_v21;
          _943_v21 = _dafny.MultiSet.fromElements(_911_v0, _911_v0, _911_v0);
          let _944_v22;
          _944_v22 = _dafny.Map.Empty.slice().updateUnsafe(_942_v20,(_943_v21).Union(_943_v21));
          _944_v22 = (_944_v22).update(_942_v20, (((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]) ? (_943_v21) : (_943_v21)));
        } else {
          let _945___mcc_h7 = (_source13).cf17;
          let _946_cf17 = _945___mcc_h7;
          let _947_v23;
          _947_v23 = new _dafny.CodePoint('e'.codePointAt(0));
          let _948_v25;
          _948_v25 = _dafny.Seq.of(_946_cf17);
          _947_v23 = (_930_cf3)[_module.__default.safeIndex((((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]) ? (new BigNumber((function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (_dafny.MultiSet.FromArray(_948_v25)).Elements) {
              let _949_v24 = _compr_29;
              if ((_dafny.MultiSet.FromArray(_948_v25)).contains(_949_v24)) {
                _coll29.push([_949_v24,_915_v3]);
              }
            }
            return _coll29;
          }()).length)) : (p0)), new BigNumber((_930_cf3).length))];
          let _950_v26;
          _950_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(_932_cf1),new BigNumber(-362));
          let _951_v27;
          _951_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
          let _952_v28;
          let _nw124 = Array((new BigNumber(11)).toNumber());
          _nw124[(_dafny.ZERO).toNumber()] = p0;
          _nw124[(_dafny.ONE).toNumber()] = new BigNumber((_950_v26).length);
          _nw124[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw124[(new BigNumber(3)).toNumber()] = new BigNumber(578);
          _nw124[(new BigNumber(4)).toNumber()] = p0;
          _nw124[(new BigNumber(5)).toNumber()] = p0;
          _nw124[(new BigNumber(6)).toNumber()] = p2;
          _nw124[(new BigNumber(7)).toNumber()] = _915_v3;
          _nw124[(new BigNumber(8)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_915_v3,_915_v3)).Merge(_951_v27)).length);
          _nw124[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_915_v3);
          _nw124[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(p0, p2);
          _952_v28 = _nw124;
          let _index133 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_952_v28).length));
          (_952_v28)[_index133] = p0;
          let _953_v29;
          _953_v29 = _module.D0.create_DC0(p2);
          let _index134 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_952_v28).length));
          (_952_v28)[_index134] = (new BigNumber(-951)).multipliedBy(_915_v3);
          let _954_v30;
          _954_v30 = _dafny.MultiSet.fromElements(_913_v1);
          let _955_v31;
          _955_v31 = _dafny.Seq.of(!(_914_v2));
          let _956_v32;
          _956_v32 = _dafny.Map.Empty.slice().updateUnsafe(_955_v31,_915_v3);
          let _index135 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_952_v28).length));
          let _index136 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_952_v28).length));
          let _rhs145 = _module.__default.fm19((_dafny.ZERO).minus(_915_v3), _914_v2, globalState);
          let _rhs146 = p2;
          let _rhs147 = _module.D0.create_DC0((_dafny.ZERO).minus(p0));
          let _rhs148 = (((_956_v32).contains(_dafny.Seq.Concat(_955_v31, _955_v31))) ? ((_956_v32).get(_dafny.Seq.Concat(_955_v31, _955_v31))) : (new BigNumber((_951_v27).length)));
          let _rhs149 = _954_v30;
          let _lhs102 = globalState;
          let _lhs103 = _952_v28;
          let _lhs104 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_952_v28).length));
          let _lhs105 = _952_v28;
          let _lhs106 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_952_v28).length));
          _lhs102.f3 = _rhs145;
          _lhs103[_lhs104] = _rhs146;
          _953_v29 = _rhs147;
          _lhs105[_lhs106] = _rhs148;
          _954_v30 = _rhs149;
          let _957_v33;
          _957_v33 = _module.D0.create_DC2(_932_cf1);
          (globalState).f9 = (new BigNumber((_946_cf17).length)).isLessThan((_this).fm5(_957_v33, globalState));
          let _index137 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_952_v28).length));
          (_952_v28)[_index137] = _module.__default.fm11((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))], globalState);
        }
        let _index138 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
        (_913_v1)[_index138] = (_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))];
      } else if (_source12.is_DC2) {
        let _958___mcc_h3 = (_source12).cf4;
        let _959_cf4 = _958___mcc_h3;
        let _960_v35;
        _960_v35 = _dafny.Seq.of(_915_v3, p2);
        let _961_v36;
        _961_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_960_v35).length),p0);
        let _962_v37;
        let _nw125 = new _module.C4();
        _nw125.__ctor(new BigNumber(13), function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of ((p1).fm1(_961_v36, globalState)).Elements) {
            let _963_v34 = _compr_30;
            if (_dafny.Seq.contains((p1).fm1(_961_v36, globalState), _963_v34)) {
              _coll30.push([_963_v34,(_dafny.ZERO).minus(p2)]);
            }
          }
          return _coll30;
        }());
        _962_v37 = _nw125;
        _959_cf4 = !((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]);
        let _index139 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
        (_913_v1)[_index139] = _914_v2;
        _915_v3 = (_962_v37).f14;
      } else if (_source12.is_DC0) {
        let _964___mcc_h4 = (_source12).cf0;
        let _965_cf0 = _964___mcc_h4;
        let _966_v38;
        _966_v38 = new _dafny.CodePoint('j'.codePointAt(0));
        let _index140 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_911_v0).length));
        (_911_v0)[_index140] = _966_v38;
        let _index141 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_911_v0).length));
        (_911_v0)[_index141] = _966_v38;
        (globalState).f3 = (_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))];
        let _source14 = _module.D4.create_DC10((function () {
  let _coll31 = new _dafny.Set();
  for (const _compr_31 of _dafny.IntegerRange(new BigNumber(21), new BigNumber(-521))) {
    let _967_v39 = _compr_31;
    if (((new BigNumber(21)).isLessThanOrEqualTo(_967_v39)) && ((_967_v39).isLessThan(new BigNumber(-521)))) {
      _coll31.add(_module.__default.safeDivisionInt(_967_v39, _915_v3));
    }
  }
  return _coll31;
}()).Intersect(_921_v9));
        if (_source14.is_DC11) {
          let _968_v40;
          _968_v40 = _dafny.Seq.of((_dafny.ZERO).minus(p0), p0, new BigNumber((_922_v10).length));
          let _969_v41;
          _969_v41 = _dafny.Map.Empty.slice().updateUnsafe(_914_v2,_968_v40);
          let _970_v42;
          _970_v42 = _module.D6.create_DC16(new BigNumber((_968_v40).length));
          let _971_v43;
          _971_v43 = _module.D0.create_DC0(p0);
          let _972_v44;
          let _nw126 = Array((new BigNumber(21)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = _968_v40;
          _nw126[(_dafny.ONE).toNumber()] = (((_969_v41).contains(_914_v2)) ? ((_969_v41).get(_914_v2)) : (_968_v40));
          _nw126[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_973_i2) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vcqodo")).length));
          }), _968_v40);
          _nw126[(new BigNumber(3)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_968_v40, _module.__default.safeIndex(p0, new BigNumber((_968_v40).length)), new BigNumber((_dafny.MultiSet.fromElements(_914_v2, (_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))], _914_v2)).cardinality()));
          _nw126[(new BigNumber(5)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-667)), function (_974_i3) {
            return new BigNumber(8);
          });
          _nw126[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(268)), ((_975_cf0) => function (_976_i4) {
            return _975_cf0;
          })(_965_cf0));
          _nw126[(new BigNumber(8)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(9)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(10)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(11)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(12)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(13)).toNumber()] = _dafny.Seq.of((_970_v42).dtor_cf24, p2, _965_cf0);
          _nw126[(new BigNumber(14)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(15)).toNumber()] = ((_914_v2) ? (_968_v40) : (_968_v40));
          _nw126[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), ((_977_v3) => function (_978_i5) {
            return _977_v3;
          })(_915_v3)), _968_v40);
          _nw126[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(p0, _965_cf0, (_971_v43).dtor_cf0), _968_v40);
          _nw126[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_968_v40, _968_v40);
          _nw126[(new BigNumber(19)).toNumber()] = _968_v40;
          _nw126[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm30(globalState), _dafny.Seq.update(_968_v40, _module.__default.safeIndex(p0, new BigNumber((_968_v40).length)), new BigNumber((_module.__default.fm20(new BigNumber((_dafny.Seq.UnicodeFromString("ybsadm")).length), _914_v2, _914_v2, _915_v3, globalState)).length)));
          _972_v44 = _nw126;
          let _index142 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_972_v44).length));
          (_972_v44)[_index142] = _968_v40;
          let _index143 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_972_v44).length));
          (_972_v44)[_index143] = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(547), p0)).cardinality()), _965_cf0, _915_v3);
          let _979_v46;
          _979_v46 = _dafny.MultiSet.fromElements(_966_v38);
          let _980_v47;
          _980_v47 = _dafny.Seq.of(!(!((((_917_v5).contains(p2)) ? ((_917_v5).get(p2)) : ((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])))), _914_v2);
          let _rhs150 = !_dafny.Seq.contains(_916_v4, _module.__default.fm23(!(!(_914_v2)), function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of ((_979_v46).update((_911_v0)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_911_v0).length))], _module.__default.abs(p2))).Elements) {
              let _981_v45 = _compr_32;
              if (((_979_v46).update((_911_v0)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_911_v0).length))], _module.__default.abs(p2))).contains(_981_v45)) {
                _coll32.push([_981_v45,_dafny.Map.Empty.slice().updateUnsafe((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))],(_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])]);
              }
            }
            return _coll32;
          }(), _914_v2, p2, globalState));
          let _rhs151 = (new BigNumber((_980_v47).length)).plus(_915_v3);
          let _rhs152 = new BigNumber(831);
          let _lhs107 = globalState;
          _lhs107.f9 = _rhs150;
          _915_v3 = _rhs151;
          _915_v3 = _rhs152;
          (globalState).f13 = (_914_v2) && (!((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]));
          let _982_v48;
          _982_v48 = _module.D7.create_DC18(_917_v5);
          let _983_v49;
          _983_v49 = _dafny.Map.Empty.slice().updateUnsafe(_982_v48,p1);
          _915_v3 = new BigNumber((_dafny.Set.fromElements(!(_983_v49).contains(_982_v48), false)).length);
        } else if (_source14.is_DC12) {
          let _984___mcc_h8 = (_source14).cf18;
          let _985_cf18 = _984___mcc_h8;
          (globalState).f3 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_916_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), ((_986_v0) => function (_987_i6) {
            return (_986_v0)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_986_v0).length))];
          })(_911_v0))), _916_v4);
          _919_v7 = _module.D7.create_DC21(_918_v6);
          let _index144 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
          (_913_v1)[_index144] = (_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))];
          let _988_v50;
          _988_v50 = _module.D4.create_DC11();
          let _989_v51;
          let _nw127 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _989_v51 = _nw127;
          let _index145 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_989_v51).length));
          (_989_v51)[_index145] = _965_cf0;
          let _index146 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_989_v51).length));
          let _rhs153 = !((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]);
          let _rhs154 = _988_v50;
          let _rhs155 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(957)), ((_990_v38) => function (_991_i7) {
            return _990_v38;
          })(_966_v38)), _dafny.Seq.Concat(_dafny.Seq.update(_916_v4, _module.__default.safeIndex(_985_cf18, new BigNumber((_916_v4).length)), new _dafny.CodePoint('u'.codePointAt(0))), _dafny.Seq.UnicodeFromString("qwrmofy")))).length);
          let _lhs108 = globalState;
          let _lhs109 = _989_v51;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_989_v51).length));
          _lhs108.f9 = _rhs153;
          _988_v50 = _rhs154;
          _lhs109[_lhs110] = _rhs155;
        } else {
          let _992___mcc_h9 = (_source14).cf17;
          let _993_cf17 = _992___mcc_h9;
          let _994_v52;
          let _nw128 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _994_v52 = _nw128;
          _994_v52 = _994_v52;
          let _995_v53;
          _995_v53 = _dafny.Map.Empty.slice().updateUnsafe(_915_v3,_916_v4);
          let _996_v54;
          _996_v54 = _dafny.Map.Empty.slice().updateUnsafe(_913_v1,new BigNumber(198));
          let _997_v55;
          _997_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_996_v54).length),p2);
          _995_v53 = (_995_v53).update((_dafny.ZERO).minus(p0), (p1).fm1(_997_v55, globalState));
          let _998_v56;
          let _init17 = ((_999_v2, _1000_v12, _1001_v4, _1002_v38) => function (_1003_i8) {
            return _module.D2.create_DC8(_module.D2.create_DC7(_999_v2, _dafny.Map.Empty.slice().updateUnsafe(_999_v2,_1000_v12), new BigNumber((_dafny.Seq.update(_1001_v4, _module.__default.safeIndex(new BigNumber((_1001_v4).length), new BigNumber((_1001_v4).length)), _1002_v38)).length), _999_v2));
          })(_914_v2, _926_v12, _916_v4, _966_v38);
          let _nw129 = Array((new BigNumber(16)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw129.length); _i0_17++) {
            _nw129[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _998_v56 = _nw129;
          let _1004_v57;
          _1004_v57 = _dafny.Map.Empty.slice().updateUnsafe((_914_v2) === ((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]),_998_v56);
          _1004_v57 = (_1004_v57).update(((_914_v2) ? (_914_v2) : (_914_v2)), _998_v56);
          let _1005_v58;
          let _nw130 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _1005_v58 = _nw130;
          let _1006_v59;
          _1006_v59 = _dafny.Map.Empty.slice().updateUnsafe((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))],_1005_v58);
          _1005_v58 = (((_1006_v59).contains((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) ? ((_1006_v59).get((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) : (_1005_v58));
        }
        _965_cf0 = (_dafny.ZERO).minus(p0);
      } else {
        let _1007___mcc_h5 = (_source12).cf5;
        let _1008_cf5 = _1007___mcc_h5;
        let _1009_v60;
        _1009_v60 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1010_v61;
        _1010_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1009_v60,_916_v4);
        let _index147 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
        (_913_v1)[_index147] = (_1010_v61).contains(new _dafny.CodePoint('s'.codePointAt(0)));
        (globalState).f9 = _914_v2;
        let _1011_v62;
        _1011_v62 = _dafny.Seq.of(_916_v4);
        let _1012_v63;
        _1012_v63 = _dafny.Map.Empty.slice().updateUnsafe(_915_v3,_916_v4);
        let _1013_v64;
        _1013_v64 = _module.D0.create_DC1((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))], _1011_v62, (((_1012_v63).contains(p0)) ? ((_1012_v63).get(p0)) : (_916_v4)));
        let _1014_v65;
        _1014_v65 = _dafny.MultiSet.fromElements(_1013_v64);
        let _1015_v66;
        _1015_v66 = _module.D2.create_DC6(_1014_v65);
        let _pat_let_tv21 = _1014_v65;
        let _1016_v67;
        _1016_v67 = _module.D2.create_DC6((function (_pat_let19_0) {
  return function (_1017_dt__update__tmp_h0) {
    return function (_pat_let20_0) {
      return function (_1018_dt__update_hcf10_h0) {
        return _module.D2.create_DC6(_1018_dt__update_hcf10_h0);
      }(_pat_let20_0);
    }(_pat_let_tv21);
  }(_pat_let19_0);
}(_1015_v66)).dtor_cf10);
        let _source15 = _1015_v66;
        if (_source15.is_DC7) {
          let _1019___mcc_h10 = (_source15).cf11;
          let _1020___mcc_h11 = (_source15).cf12;
          let _1021___mcc_h12 = (_source15).cf13;
          let _1022___mcc_h13 = (_source15).cf14;
          let _1023_cf14 = _1022___mcc_h13;
          let _1024_cf13 = _1021___mcc_h12;
          let _1025_cf12 = _1020___mcc_h11;
          let _1026_cf11 = _1019___mcc_h10;
          let _1027_v68;
          let _init18 = ((_1028_cf13) => function (_1029_i9) {
            return (_1029_i9).minus(_1028_cf13);
          })(_1024_cf13);
          let _nw131 = Array((new BigNumber(29)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw131.length); _i0_18++) {
            _nw131[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _1027_v68 = _nw131;
          let _index148 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1027_v68).length));
          (_1027_v68)[_index148] = _1024_cf13;
          let _index149 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1027_v68).length));
          (_1027_v68)[_index149] = p2;
          let _1030_v69;
          _1030_v69 = _dafny.Map.Empty.slice().updateUnsafe(_913_v1,_1023_cf14);
          let _1031_v70;
          _1031_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1026_cf11,_1024_cf13);
          let _index150 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1027_v68).length));
          let _index151 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1027_v68).length));
          let _rhs156 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_913_v1,_914_v2)).Merge(_1030_v69)).length), _module.__default.safeDivisionInt(new BigNumber(-701), (((_1031_v70).contains(_914_v2)) ? ((_1031_v70).get(_914_v2)) : (new BigNumber(214))))));
          let _rhs157 = (_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))];
          let _rhs158 = new BigNumber((_916_v4).length);
          let _rhs159 = _915_v3;
          let _lhs111 = _1027_v68;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1027_v68).length));
          let _lhs113 = _1027_v68;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1027_v68).length));
          _1024_cf13 = _rhs156;
          _914_v2 = _rhs157;
          _lhs111[_lhs112] = _rhs158;
          _lhs113[_lhs114] = _rhs159;
          _913_v1 = _913_v1;
          let _1032_v71;
          _1032_v71 = _dafny.Seq.of(_1026_cf11);
          let _1033_v72;
          _1033_v72 = _dafny.Seq.of(_1032_v71);
          let _1034_v73;
          _1034_v73 = _module.D2.create_DC7((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))], _1025_cf12, _module.__default.fm11(_1023_cf14, globalState), _1026_cf11);
          let _pat_let_tv22 = _1025_cf12;
          _1024_cf13 = (new BigNumber(143)).plus(new BigNumber(((_1033_v72)[_module.__default.safeIndex((function (_pat_let21_0) {
            return function (_1035_dt__update__tmp_h1) {
              return function (_pat_let22_0) {
                return function (_1036_dt__update_hcf12_h0) {
                  return _module.D2.create_DC7((_1035_dt__update__tmp_h1).dtor_cf11, _1036_dt__update_hcf12_h0, (_1035_dt__update__tmp_h1).dtor_cf13, (_1035_dt__update__tmp_h1).dtor_cf14);
                }(_pat_let22_0);
              }(_pat_let_tv22);
            }(_pat_let21_0);
          }(_1034_v73)).dtor_cf13, new BigNumber((_1033_v72).length))]).length));
          let _1037_v74;
          let _nw132 = new _module.C1();
          _nw132.__ctor();
          _1037_v74 = _nw132;
        } else if (_source15.is_DC6) {
          let _1038___mcc_h14 = (_source15).cf10;
          let _1039_cf10 = _1038___mcc_h14;
          let _1040_v75;
          let _nw133 = new _module.C1();
          _nw133.__ctor();
          _1040_v75 = _nw133;
          (globalState).f13 = true;
          let _1041_v76;
          _1041_v76 = _module.D2.create_DC7(_914_v2, _dafny.Map.Empty.slice().updateUnsafe(_914_v2,_module.D0.create_DC3(_925_v11)), p0, _914_v2);
          let _1042_v77;
          _1042_v77 = _module.D2.create_DC8(_1041_v76);
          let _1043_v78;
          _1043_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1042_v77,p2);
          let _1044_v79;
          _1044_v79 = _dafny.Map.Empty.slice().updateUnsafe(_914_v2,_915_v3);
          let _1045_v80;
          _1045_v80 = _dafny.Map.Empty.slice().updateUnsafe((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))],(_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]);
          _1043_v78 = (_1043_v78).update(_module.D2.create_DC8(_1041_v76), (((_1044_v79).contains((((_1045_v80).contains((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) ? ((_1045_v80).get((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) : (_914_v2)))) ? ((_1044_v79).get((((_1045_v80).contains((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) ? ((_1045_v80).get((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) : (_914_v2)))) : (p2)));
          (globalState).f9 = _914_v2;
        } else {
          let _1046___mcc_h15 = (_source15).cf15;
          let _1047_cf15 = _1046___mcc_h15;
          let _1048_v81;
          _1048_v81 = _dafny.Map.Empty.slice().updateUnsafe(_914_v2,_926_v12);
          let _1049_v82;
          _1049_v82 = _module.D2.create_DC7(false, _1048_v81, (_dafny.ZERO).minus(p0), _914_v2);
          let _1050_v83;
          _1050_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1009_v60,_914_v2);
          _915_v3 = (_dafny.ZERO).minus(((_module.__default.fm22(_1049_v82, _926_v12, (((_1050_v83).contains(_1009_v60)) ? ((_1050_v83).get(_1009_v60)) : (_914_v2)), globalState)).multipliedBy(p2)).plus(p2));
          let _1051_v84;
          _1051_v84 = _dafny.Seq.of(_914_v2, _914_v2, _914_v2);
          let _1052_v85;
          _1052_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('k'.codePointAt(0)));
          let _1053_v86;
          _1053_v86 = _dafny.Map.Empty.slice().updateUnsafe((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))],(_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))]);
          let _1054_v87;
          _1054_v87 = _dafny.Map.Empty.slice().updateUnsafe(_914_v2,!((((_1053_v86).contains((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) ? ((_1053_v86).get((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))])) : (_module.__default.fm6(globalState)))));
          let _1055_v88;
          _1055_v88 = _dafny.Seq.of(p0, new BigNumber(-878), _915_v3, p0, new BigNumber((_1054_v87).length));
          let _1056_v89;
          let _nw134 = Array((new BigNumber(13)).toNumber());
          _nw134[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1051_v84).length));
          _nw134[(_dafny.ONE).toNumber()] = (_915_v3).minus(new BigNumber(-518));
          _nw134[(new BigNumber(2)).toNumber()] = _915_v3;
          _nw134[(new BigNumber(3)).toNumber()] = _915_v3;
          _nw134[(new BigNumber(4)).toNumber()] = new BigNumber((_1052_v85).length);
          _nw134[(new BigNumber(5)).toNumber()] = _915_v3;
          _nw134[(new BigNumber(6)).toNumber()] = (_1055_v88)[_module.__default.safeIndex(p2, new BigNumber((_1055_v88).length))];
          _nw134[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("uvsvmpmou")).length);
          _nw134[(new BigNumber(8)).toNumber()] = p0;
          _nw134[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-64));
          _nw134[(new BigNumber(10)).toNumber()] = p0;
          _nw134[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(p2, _915_v3)));
          _nw134[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(p0, p2);
          _1056_v89 = _nw134;
          let _index152 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
          let _rhs160 = true;
          let _rhs161 = _1056_v89;
          let _rhs162 = _914_v2;
          let _rhs163 = !_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(893)), ((_1057_v60) => function (_1058_i10) {
            return _1057_v60;
          })(_1009_v60)), _916_v4), _1009_v60);
          let _rhs164 = p2;
          let _lhs115 = globalState;
          let _lhs116 = globalState;
          let _lhs117 = _913_v1;
          let _lhs118 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length));
          _lhs115.f3 = _rhs160;
          _1056_v89 = _rhs161;
          _lhs116.f5 = _rhs162;
          _lhs117[_lhs118] = _rhs163;
          _915_v3 = _rhs164;
          (globalState).f3 = (_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))];
          let _1059_v90;
          let _nw135 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1059_v90 = _nw135;
          let _index153 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1059_v90).length));
          (_1059_v90)[_index153] = _1056_v89;
          let _index154 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_1059_v90).length));
          (_1059_v90)[_index154] = _1056_v89;
        }
        let _1060_v91;
        let _nw136 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _1060_v91 = _nw136;
        let _index155 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_1060_v91).length));
        (_1060_v91)[_index155] = _module.__default.fm11((_913_v1)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_913_v1).length))], globalState);
        let _index156 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_1060_v91).length));
        (_1060_v91)[_index156] = p0;
      }
      return;
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _1061_v0;
      _1061_v0 = _dafny.Seq.UnicodeFromString("scsqu");
      _1061_v0 = _1061_v0;
      let _1062_v1;
      _1062_v1 = false;
      r2 = (_this).fm5(_module.D0.create_DC2(_1062_v1), globalState);
      let _1063_v2;
      _1063_v2 = new BigNumber(931);
      let _hi6 = _1063_v2;
      for (let _1064_i0 = new BigNumber((_1061_v0).length); _1064_i0.isLessThan(_hi6); _1064_i0 = _1064_i0.plus(_dafny.ONE)) {
        r2 = (_1063_v2).multipliedBy(_1064_i0);
        let _1065_v3;
        _1065_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1063_v2,_module.__default.fm6(globalState));
        let _1066_v4;
        let _nw137 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
        _1066_v4 = _nw137;
        let _1067_v5;
        let _nw138 = Array((new BigNumber(14)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = _1062_v1;
        _nw138[(_dafny.ONE).toNumber()] = !(_1062_v1);
        _nw138[(new BigNumber(2)).toNumber()] = _1062_v1;
        _nw138[(new BigNumber(3)).toNumber()] = (_1062_v1) && (_1062_v1);
        _nw138[(new BigNumber(4)).toNumber()] = (_1063_v2).isLessThanOrEqualTo(_1063_v2);
        _nw138[(new BigNumber(5)).toNumber()] = false;
        _nw138[(new BigNumber(6)).toNumber()] = (_1063_v2).isLessThan(new BigNumber(807));
        _nw138[(new BigNumber(7)).toNumber()] = _1062_v1;
        _nw138[(new BigNumber(8)).toNumber()] = _1062_v1;
        _nw138[(new BigNumber(9)).toNumber()] = _1062_v1;
        _nw138[(new BigNumber(10)).toNumber()] = ((((_1065_v3).contains(new BigNumber(350))) ? ((_1065_v3).get(new BigNumber(350))) : (_1062_v1))) === (_1062_v1);
        _nw138[(new BigNumber(11)).toNumber()] = _1062_v1;
        _nw138[(new BigNumber(12)).toNumber()] = false;
        _nw138[(new BigNumber(13)).toNumber()] = false;
        _1067_v5 = _nw138;
        let _rhs165 = (function () {
          let _coll33 = new _dafny.Map();
          for (const _compr_33 of _dafny.IntegerRange(new BigNumber(791), new BigNumber(-974))) {
            let _1068_v6 = _compr_33;
            if (((new BigNumber(791)).isLessThanOrEqualTo(_1068_v6)) && ((_1068_v6).isLessThan(new BigNumber(-974)))) {
              _coll33.push([(_1068_v6).minus(_1063_v2),_1062_v1]);
            }
          }
          return _coll33;
        }()).update(_1064_i0, _1062_v1);
        let _rhs166 = _1064_i0;
        let _rhs167 = _1066_v4;
        let _rhs168 = ((!(!(!(_module.__default.fm6(globalState))))) ? (_1067_v5) : (p1));
        _1065_v3 = _rhs165;
        r1 = _rhs166;
        _1066_v4 = _rhs167;
        _1067_v5 = _rhs168;
        r2 = _1064_i0;
        let _1069_v7;
        let _nw139 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1069_v7 = _nw139;
        let _index157 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_1069_v7).length));
        (_1069_v7)[_index157] = _1063_v2;
        let _1070_v8;
        _1070_v8 = _dafny.Seq.of(_1062_v1);
        let _1071_v9;
        _1071_v9 = _module.D1.create_DC5(_1062_v1, _1061_v0, _1070_v8);
        let _1072_v10;
        _1072_v10 = _dafny.Seq.of(_1062_v1, !((_1071_v9).dtor_cf7));
        let _index158 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_1069_v7).length));
        (_1069_v7)[_index158] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_1062_v1, _1062_v1, _module.__default.fm6(globalState)), _1070_v8), _module.__default.safeIndex((_dafny.ZERO).minus(_1063_v2), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_1062_v1, _1062_v1, _module.__default.fm6(globalState)), _1070_v8)).length)), _1062_v1), _module.__default.safeIndex(_1063_v2, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_1062_v1, _1062_v1, _module.__default.fm6(globalState)), _1070_v8), _module.__default.safeIndex((_dafny.ZERO).minus(_1063_v2), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_1062_v1, _1062_v1, _module.__default.fm6(globalState)), _1070_v8)).length)), _1062_v1)).length)), _1062_v1), _1072_v10)).length);
      }
      let _1073_i1;
      _1073_i1 = _dafny.ZERO;
      L7: {
        while (!(_module.__default.fm6(globalState))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1073_i1)) {
              break L7;
            }
            _1073_i1 = (_1073_i1).plus(_dafny.ONE);
            let _1074_v11;
            _1074_v11 = _dafny.Seq.of(_1062_v1);
            let _1075_v12;
            _1075_v12 = new _dafny.CodePoint('c'.codePointAt(0));
            let _source16 = _module.__default.fm35(_dafny.Seq.update(_dafny.Seq.Concat(_1074_v11, _1074_v11), _module.__default.safeIndex(_1063_v2, new BigNumber((_dafny.Seq.Concat(_1074_v11, _1074_v11)).length)), _1062_v1), !(_1062_v1) || (_1062_v1), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(703)), ((_1076_v0, _1077_v2) => function (_1078_i2) {
              return (_1076_v0)[_module.__default.safeIndex(_1077_v2, new BigNumber((_1076_v0).length))];
            })(_1061_v0, _1063_v2)), _module.__default.safeIndex(_1063_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(703)), ((_1079_v0, _1080_v2) => function (_1081_i2) {
              return (_1079_v0)[_module.__default.safeIndex(_1080_v2, new BigNumber((_1079_v0).length))];
            })(_1061_v0, _1063_v2))).length)), _1075_v12), globalState);
            if (_source16.is_DC11) {
              _1061_v0 = _1061_v0;
              let _1082_v13;
              _1082_v13 = _dafny.MultiSet.fromElements(_1074_v11, _dafny.Seq.Concat(_1074_v11, _1074_v11), _1074_v11);
              let _rhs169 = (p0)[_module.__default.safeIndex(_1063_v2, new BigNumber((p0).length))];
              let _rhs170 = (_dafny.MultiSet.fromElements(_1074_v11, _1074_v11, _1074_v11)).Difference(_1082_v13);
              let _rhs171 = (_dafny.ZERO).minus(new BigNumber(-497));
              r2 = _rhs169;
              _1082_v13 = _rhs170;
              _1063_v2 = _rhs171;
              let _1083_v14;
              _1083_v14 = _dafny.MultiSet.fromElements(_1061_v0);
              let _rhs172 = _module.__default.safeDivisionInt(((_module.__default.fm6(globalState)) ? (_1063_v2) : ((_dafny.ZERO).minus(_1063_v2))), _1063_v2);
              let _rhs173 = _dafny.MultiSet.fromElements(_1061_v0);
              let _rhs174 = _1063_v2;
              let _rhs175 = _module.__default.fm6(globalState);
              let _rhs176 = !(!(_1062_v1));
              let _lhs119 = globalState;
              let _lhs120 = globalState;
              r2 = _rhs172;
              _1083_v14 = _rhs173;
              r1 = _rhs174;
              _lhs119.f3 = _rhs175;
              _lhs120.f3 = _rhs176;
              let _1084_v15;
              _1084_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v1,_1062_v1);
              let _1085_v16;
              _1085_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1075_v12,_1084_v15);
              let _1086_v17;
              let _nw140 = Array((new BigNumber(19)).toNumber());
              _nw140[(_dafny.ZERO).toNumber()] = _1075_v12;
              _nw140[(_dafny.ONE).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(2)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(3)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(4)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(5)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(6)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(7)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(8)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(9)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(10)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
              _nw140[(new BigNumber(12)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(13)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(14)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(15)).toNumber()] = _module.__default.fm23(_1062_v1, _1085_v16, _1062_v1, _1063_v2, globalState);
              _nw140[(new BigNumber(16)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(17)).toNumber()] = _1075_v12;
              _nw140[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
              _1086_v17 = _nw140;
              let _1087_v20;
              _1087_v20 = _dafny.Set.fromElements(_1063_v2);
              let _1088_v21;
              _1088_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1075_v12,function () {
                let _coll34 = new _dafny.Map();
                for (const _compr_34 of (_1087_v20).Elements) {
                  let _1089_v19 = _compr_34;
                  if ((_1087_v20).contains(_1089_v19)) {
                    _coll34.push([_module.__default.safeDivisionInt(_1089_v19, _1063_v2),_1062_v1]);
                  }
                }
                return _coll34;
              }());
              let _index159 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_1086_v17).length));
              (_1086_v17)[_index159] = _module.__default.fm23(_1062_v1, function () {
                let _coll35 = new _dafny.Map();
                for (const _compr_35 of (_1088_v21).Keys.Elements) {
                  let _1090_v18 = _compr_35;
                  if ((_1088_v21).contains(_1090_v18)) {
                    _coll35.push([_1090_v18,(_dafny.Map.Empty.slice().updateUnsafe(_1062_v1,_1062_v1)).update(_1062_v1, _1062_v1)]);
                  }
                }
                return _coll35;
              }(), _1062_v1, _1063_v2, globalState);
              let _index160 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_1086_v17).length));
              (_1086_v17)[_index160] = _1075_v12;
            } else if (_source16.is_DC12) {
              let _1091___mcc_h0 = (_source16).cf18;
              let _1092_cf18 = _1091___mcc_h0;
              (globalState).f3 = _1062_v1;
              r2 = _module.__default.fm11(_1062_v1, globalState);
              (globalState).f13 = true;
              _1062_v1 = _1062_v1;
            } else {
              let _1093___mcc_h1 = (_source16).cf17;
              let _1094_cf17 = _1093___mcc_h1;
              let _1095_v22;
              _1095_v22 = _dafny.Map.Empty.slice().updateUnsafe((_1094_cf17).Intersect(_1094_cf17),_dafny.Seq.UnicodeFromString("bgddjiqf"));
              _1095_v22 = (_1095_v22).update(_1094_cf17, _dafny.Seq.Concat(_1061_v0, _1061_v0));
              let _1096_v23;
              _1096_v23 = _dafny.Set.fromElements(_1062_v1);
              let _1097_v24;
              _1097_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1062_v1,_1074_v11)).length),(_1063_v2).multipliedBy(new BigNumber((_1096_v23).length)));
              _1097_v24 = _1097_v24;
              _1062_v1 = false;
              r2 = _1063_v2;
            }
            let _1098_v25;
            let _init19 = ((_1099_v2) => function (_1100_i3) {
              return (_1100_i3).plus(_1099_v2);
            })(_1063_v2);
            let _nw141 = Array((new BigNumber(3)).toNumber());
            for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw141.length); _i0_19++) {
              _nw141[_i0_19] = _init19(new BigNumber(_i0_19));
            }
            _1098_v25 = _nw141;
            let _1101_v26;
            _1101_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1063_v2,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm0(_1063_v2, _1075_v12, globalState),false)).length));
            let _index161 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1098_v25).length));
            (_1098_v25)[_index161] = new BigNumber((_1101_v26).length);
            let _index162 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1098_v25).length));
            (_1098_v25)[_index162] = new BigNumber((_1061_v0).length);
            let _1102_v27;
            _1102_v27 = _dafny.MultiSet.fromElements(new BigNumber(468));
            _1102_v27 = _dafny.MultiSet.fromElements((_1098_v25)[_module.__default.safeIndex(new BigNumber(903), new BigNumber((_1098_v25).length))], (_1098_v25)[_module.__default.safeIndex(new BigNumber(903), new BigNumber((_1098_v25).length))]);
            _1075_v12 = new _dafny.CodePoint('a'.codePointAt(0));
          }
        }
      }
      r2 = _1063_v2;
      let _1103_v28;
      _1103_v28 = _dafny.Set.fromElements((_1063_v2).isEqualTo(_1063_v2));
      _1103_v28 = _module.__default.fm33(_1062_v1, _1062_v1, _1062_v1, globalState);
      r0 = false;
      let _1104_v29;
      _1104_v29 = _module.D0.create_DC0(_1063_v2);
      let _1105_v30;
      _1105_v30 = _module.D0.create_DC3(_1104_v29);
      let _1106_v31;
      _1106_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1062_v1,_1105_v30);
      let _1107_v32;
      _1107_v32 = _module.D2.create_DC7(_1062_v1, _1106_v31, _1063_v2, false);
      r1 = _module.__default.fm22(_1107_v32, _module.D0.create_DC3(_1104_v29), (new BigNumber((p0).length)).isEqualTo(new BigNumber(229)), globalState);
      r2 = _1063_v2;
      return [r0, r1, r2];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-869))), (_module.D0.create_DC0(new BigNumber(211))).dtor_cf0)).multipliedBy(new BigNumber(-444));
    };
    fm1(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tiwttktph"), _dafny.Seq.UnicodeFromString("cd")), _dafny.Seq.UnicodeFromString("aprlnhebx"));
    };
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return false;
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _1108_v0;
      _1108_v0 = true;
      let _1109_v1;
      _1109_v1 = new BigNumber(619);
      let _1110_v2;
      let _nw142 = new _module.C3();
      _nw142.__ctor(_1108_v0, _1109_v1);
      _1110_v2 = _nw142;
      let _1111_v3;
      _1111_v3 = _dafny.Set.fromElements(_1110_v2, _1110_v2);
      let _1112_v4;
      _1112_v4 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(381)), ((_1113_p0) => function (_1114_i0) {
        return _1113_p0;
      })(p0))).length), new BigNumber((_1111_v3).length));
      if (((_1108_v0) ? ((_dafny.Set.fromElements(_1109_v1)).IsDisjointFrom(_1112_v4)) : (_1108_v0))) {
        let _1115_v5;
        _1115_v5 = _dafny.Seq.of(_1108_v0, _1108_v0, _1108_v0);
        let _1116_v6;
        let _1117_v7;
        let _1118_v8;
        let _1119_v9;
        let _out14;
        let _out15;
        let _out16;
        let _out17;
        let _outcollector5 = (_this).m2(_1115_v5, _1108_v0, globalState);
        _out14 = _outcollector5[0];
        _out15 = _outcollector5[1];
        _out16 = _outcollector5[2];
        _out17 = _outcollector5[3];
        _1116_v6 = _out14;
        _1117_v7 = _out15;
        _1118_v8 = _out16;
        _1119_v9 = _out17;
        let _1120_v10;
        let _nw143 = Array((new BigNumber(4)).toNumber()).fill(false);
        _1120_v10 = _nw143;
        let _index163 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length));
        (_1120_v10)[_index163] = _1108_v0;
        let _index164 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length));
        let _rhs177 = (_1109_v1).isLessThan(_1109_v1);
        let _rhs178 = (_1109_v1).multipliedBy(_1109_v1);
        let _rhs179 = _module.__default.fm19(_1109_v1, _1118_v8, globalState);
        let _lhs121 = globalState;
        let _lhs122 = _1120_v10;
        let _lhs123 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length));
        _lhs121.f5 = _rhs177;
        _1109_v1 = _rhs178;
        _lhs122[_lhs123] = _rhs179;
        let _1121_v11;
        _1121_v11 = _dafny.Seq.UnicodeFromString("mcaxr");
        let _1122_v12;
        _1122_v12 = _dafny.Seq.of(_1121_v11);
        let _1123_v13;
        _1123_v13 = _module.D0.create_DC1(_1118_v8, _1122_v12, _1121_v11);
        let _1124_v14;
        _1124_v14 = _module.D2.create_DC6(_dafny.MultiSet.fromElements(_1123_v13));
        let _source17 = _1124_v14;
        if (_source17.is_DC7) {
          let _1125___mcc_h0 = (_source17).cf11;
          let _1126___mcc_h1 = (_source17).cf12;
          let _1127___mcc_h2 = (_source17).cf13;
          let _1128___mcc_h3 = (_source17).cf14;
          let _1129_cf14 = _1128___mcc_h3;
          let _1130_cf13 = _1127___mcc_h2;
          let _1131_cf12 = _1126___mcc_h1;
          let _1132_cf11 = _1125___mcc_h0;
          let _1133_v15;
          let _init20 = ((_1134_p0, _1135_v11, _1136_cf14) => function (_1137_i1) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), ((_1138_p0) => function (_1139_i2) {
              return _1138_p0;
            })(_1134_p0)), (_module.D7.create_DC20(_1135_v11, _1136_cf14)).dtor_cf31);
          })(p0, _1121_v11, _1129_cf14);
          let _nw144 = Array((new BigNumber(24)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw144.length); _i0_20++) {
            _nw144[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _1133_v15 = _nw144;
          let _index165 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1133_v15).length));
          (_1133_v15)[_index165] = _1121_v11;
          let _1140_v16;
          _1140_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1132_cf11,_1130_cf13);
          let _index166 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_1133_v15).length));
          (_1133_v15)[_index166] = _module.__default.fm20(_1130_cf13, _1108_v0, (_1109_v1).isLessThan((((_1140_v16).contains(_1132_cf11)) ? ((_1140_v16).get(_1132_cf11)) : (new BigNumber((_dafny.Seq.UnicodeFromString("prnfmcqd")).length)))), _1109_v1, globalState);
          r1 = _1130_cf13;
          let _1141_v17;
          _1141_v17 = _dafny.Seq.of(_1115_v5);
          let _1142_v18;
          _1142_v18 = _dafny.Seq.of(_1109_v1, new BigNumber((_1141_v17).length), _1130_cf13, _1109_v1);
          let _1143_v19;
          _1143_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1109_v1,_1121_v11);
          _1130_cf13 = ((_this).fm0(_1109_v1, p0, globalState)).multipliedBy((_1142_v18)[_module.__default.safeIndex(new BigNumber(((((_1143_v19).contains(_1130_cf13)) ? ((_1143_v19).get(_1130_cf13)) : (_dafny.Seq.UnicodeFromString("uqbs")))).length), new BigNumber((_1142_v18).length))]);
          _1129_cf14 = !(!(_1109_v1).isEqualTo(new BigNumber(939)));
        } else if (_source17.is_DC6) {
          let _1144___mcc_h4 = (_source17).cf10;
          let _1145_cf10 = _1144___mcc_h4;
          let _1146_v20;
          let _nw145 = new _module.C4();
          _nw145.__ctor(_module.__default.safeModuloInt(new BigNumber(459), _1109_v1), _module.__default.fm36(globalState));
          _1146_v20 = _nw145;
          let _1147_v21;
          _1147_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1120_v10,_1118_v8);
          _1147_v21 = ((_dafny.Map.Empty.slice().updateUnsafe(_1120_v10,(_1120_v10)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length))])).Merge(_1147_v21)).Merge(_1147_v21);
          _1109_v1 = new BigNumber(999);
          let _1148_v22;
          _1148_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1118_v8,!_dafny.Seq.contains(_1115_v5, _1108_v0));
          _1148_v22 = (_1148_v22).update(_1118_v8, !(_1108_v0));
        } else {
          let _1149___mcc_h5 = (_source17).cf15;
          let _1150_cf15 = _1149___mcc_h5;
          r1 = _module.__default.safeModuloInt(_1109_v1, _1109_v1);
          (globalState).f9 = (_1120_v10)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length))];
          (globalState).f13 = ((_1112_v4).IsSubsetOf(_1112_v4)) === (!(!(_1118_v8) || (_1118_v8)));
          _1118_v8 = (((_1118_v8) ? (new BigNumber(850)) : (_1109_v1))).isLessThanOrEqualTo(_1109_v1);
        }
        let _1151_v23;
        let _nw146 = new _module.C2();
        _nw146.__ctor();
        _1151_v23 = _nw146;
        let _1152_v24;
        _1152_v24 = _dafny.MultiSet.fromElements(_1151_v23);
        (globalState).f9 = !((_1152_v24).IsDisjointFrom((_dafny.MultiSet.fromElements(_1151_v23)).update(_1151_v23, _module.__default.abs(new BigNumber((_1115_v5).length)))));
        if ((_1120_v10)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length))]) {
          _1109_v1 = _1109_v1;
          let _1153_v25;
          _1153_v25 = _dafny.MultiSet.fromElements(_module.D0.create_DC1(true, _1122_v12, _1121_v11));
          let _1154_v26;
          let _nw147 = new _module.C0();
          _nw147.__ctor((_1153_v25).Intersect(_1153_v25), _1118_v8);
          _1154_v26 = _nw147;
          let _1155_v27;
          let _1156_v28;
          let _1157_v29;
          let _1158_v30;
          let _out18;
          let _out19;
          let _out20;
          let _out21;
          let _outcollector6 = (_this).m2(_1115_v5, !(!(_1108_v0)), globalState);
          _out18 = _outcollector6[0];
          _out19 = _outcollector6[1];
          _out20 = _outcollector6[2];
          _out21 = _outcollector6[3];
          _1155_v27 = _out18;
          _1156_v28 = _out19;
          _1157_v29 = _out20;
          _1158_v30 = _out21;
          let _1159_v31;
          _1159_v31 = _module.D5.create_DC13(_1120_v10);
          let _1160_v32;
          _1160_v32 = _dafny.MultiSet.fromElements(_1159_v31);
          _1108_v0 = (_1160_v32).IsProperSubsetOf(_1160_v32);
          let _index167 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length));
          (_1120_v10)[_index167] = (_1112_v4).IsDisjointFrom(_1112_v4);
        } else {
          let _1161_v33;
          let _1162_v34;
          let _1163_v35;
          let _out22;
          let _out23;
          let _out24;
          let _outcollector7 = (_1151_v23).m6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-864)), function (_1164_i3) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }), _1121_v11, _1118_v8, _1109_v1, globalState);
          _out22 = _outcollector7[0];
          _out23 = _outcollector7[1];
          _out24 = _outcollector7[2];
          _1161_v33 = _out22;
          _1162_v34 = _out23;
          _1163_v35 = _out24;
          let _1165_v36;
          let _nw148 = new _module.C2();
          _nw148.__ctor();
          _1165_v36 = _nw148;
          let _1166_v37;
          _1166_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1162_v34,(_1120_v10)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length))]);
          let _1167_v38;
          _1167_v38 = _dafny.Seq.of(new BigNumber(-119), _1162_v34);
          let _1168_v39;
          _1168_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1163_v35,(_1167_v38)[_module.__default.safeIndex(_1162_v34, new BigNumber((_1167_v38).length))]);
          let _1169_v40;
          let _nw149 = Array((new BigNumber(25)).toNumber());
          _nw149[(_dafny.ZERO).toNumber()] = _1109_v1;
          _nw149[(_dafny.ONE).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(2)).toNumber()] = _1162_v34;
          _nw149[(new BigNumber(3)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(4)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1108_v0, _1161_v33)).length));
          _nw149[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), ((_1170_p0) => function (_1171_i4) {
            return _1170_p0;
          })(p0)), _1121_v11, _dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), ((_1172_p0) => function (_1173_i5) {
            return _1172_p0;
          })(p0)))).length);
          _nw149[(new BigNumber(7)).toNumber()] = _1162_v34;
          _nw149[(new BigNumber(8)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1118_v8,p0)).length);
          _nw149[(new BigNumber(10)).toNumber()] = _1162_v34;
          _nw149[(new BigNumber(11)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(12)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(13)).toNumber()] = new BigNumber((_1121_v11).length);
          _nw149[(new BigNumber(14)).toNumber()] = new BigNumber(36);
          _nw149[(new BigNumber(15)).toNumber()] = (_1165_v36).fm0(_1109_v1, p0, globalState);
          _nw149[(new BigNumber(16)).toNumber()] = (_1167_v38)[_module.__default.safeIndex(new BigNumber((_1168_v39).length), new BigNumber((_1167_v38).length))];
          _nw149[(new BigNumber(17)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(18)).toNumber()] = _1162_v34;
          _nw149[(new BigNumber(19)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(20)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(21)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(22)).toNumber()] = _1162_v34;
          _nw149[(new BigNumber(23)).toNumber()] = _1109_v1;
          _nw149[(new BigNumber(24)).toNumber()] = _1109_v1;
          _1169_v40 = _nw149;
          let _1174_v41;
          _1174_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1169_v40,(_dafny.ZERO).minus(_1109_v1));
          let _1175_v42;
          _1175_v42 = _module.D8.create_DC22(_1174_v41);
          let _1176_v43;
          _1176_v43 = _dafny.Set.fromElements(_1108_v0, _1163_v35, _1108_v0);
          _1166_v37 = (_1166_v37).update((_1167_v38)[_module.__default.safeIndex(new BigNumber(((_1175_v42).dtor_cf34).length), new BigNumber((_1167_v38).length))], (_1176_v43).IsProperSubsetOf(_1176_v43));
          _1121_v11 = _dafny.Seq.update((_1110_v2).fm1(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(553),_1162_v34), globalState), _module.__default.safeIndex(new BigNumber(233), new BigNumber(((_1110_v2).fm1(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(553),_1162_v34), globalState)).length)), ((!(_1118_v8)) ? (p0) : (p0)));
          _1161_v33 = (_1120_v10)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1120_v10).length))];
        }
      } else {
        let _source18 = _module.D7.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(_1109_v1,true));
        if (_source18.is_DC19) {
          let _1177___mcc_h6 = (_source18).cf27;
          let _1178___mcc_h7 = (_source18).cf28;
          let _1179___mcc_h8 = (_source18).cf29;
          let _1180___mcc_h9 = (_source18).cf30;
          let _1181_cf30 = _1180___mcc_h9;
          let _1182_cf29 = _1179___mcc_h8;
          let _1183_cf28 = _1178___mcc_h7;
          let _1184_cf27 = _1177___mcc_h6;
          let _1185_v44;
          _1185_v44 = _dafny.Seq.of(true);
          let _1186_v45;
          let _1187_v46;
          let _1188_v47;
          let _1189_v48;
          let _out25;
          let _out26;
          let _out27;
          let _out28;
          let _outcollector8 = (_this).m2(_dafny.Seq.Concat(_1185_v44, _1185_v44), (_1183_cf28).isLessThan(_1109_v1), globalState);
          _out25 = _outcollector8[0];
          _out26 = _outcollector8[1];
          _out27 = _outcollector8[2];
          _out28 = _outcollector8[3];
          _1186_v45 = _out25;
          _1187_v46 = _out26;
          _1188_v47 = _out27;
          _1189_v48 = _out28;
          _1109_v1 = _1183_cf28;
          let _1190_v49;
          _1190_v49 = _dafny.Map.Empty.slice().updateUnsafe(true,((_1182_cf29) ? (_1188_v47) : (_1188_v47)));
          let _1191_v50;
          _1191_v50 = _dafny.Seq.of(_1190_v49);
          let _1192_v51;
          _1192_v51 = _dafny.Seq.UnicodeFromString("tvnpxdv");
          let _1193_v52;
          _1193_v52 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("anlbx"), _1192_v51);
          let _1194_v53;
          _1194_v53 = _module.D0.create_DC1(_1188_v47, _dafny.Seq.update(_1193_v52, _module.__default.safeIndex(new BigNumber(516), new BigNumber((_1193_v52).length)), _1192_v51), _1192_v51);
          let _1195_v54;
          _1195_v54 = _dafny.MultiSet.fromElements(_1194_v53, _1194_v53);
          let _1196_v55;
          let _nw150 = new _module.C0();
          _nw150.__ctor(_1195_v54, _1108_v0);
          _1196_v55 = _nw150;
          let _1197_v56;
          _1197_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1196_v55,_1182_cf29);
          let _1198_v57;
          _1198_v57 = _dafny.Map.Empty.slice().updateUnsafe(!(_1182_cf29),new BigNumber(((_1197_v56).update(_1196_v55, (_1196_v55).f19)).length));
          _1190_v49 = (_1190_v49).Merge((_1191_v50)[_module.__default.safeIndex((((_1198_v57).contains(_1182_cf29)) ? ((_1198_v57).get(_1182_cf29)) : (_1183_cf28)), new BigNumber((_1191_v50).length))]);
          let _1199_v58;
          let _nw151 = new _module.C2();
          _nw151.__ctor();
          _1199_v58 = _nw151;
          _1199_v58 = _1199_v58;
        } else if (_source18.is_DC20) {
          let _1200___mcc_h10 = (_source18).cf31;
          let _1201___mcc_h11 = (_source18).cf32;
          let _1202_cf32 = _1201___mcc_h11;
          let _1203_cf31 = _1200___mcc_h10;
          let _1204_v59;
          _1204_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1109_v1,new BigNumber(207));
          let _1205_v60;
          _1205_v60 = _dafny.Seq.of(_1203_cf31);
          let _1206_v61;
          _1206_v61 = _module.D0.create_DC1(_1108_v0, _1205_v60, _1203_cf31);
          let _1207_v62;
          _1207_v62 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(new BigNumber((_dafny.Set.fromElements(_1202_cf32, _1108_v0, _1202_cf32, true)).length), _1108_v0, globalState),_module.D0.create_DC3(_1206_v61));
          let _1208_v63;
          _1208_v63 = _module.D2.create_DC7(_1108_v0, _1207_v62, new BigNumber((_1203_cf31).length), true);
          let _1209_v64;
          _1209_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1109_v1,new _dafny.CodePoint('t'.codePointAt(0)));
          let _1210_v65;
          _1210_v65 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1108_v0,_1202_cf32));
          let _1211_v67;
          _1211_v67 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1210_v65).cardinality()),_1202_cf32), function () {
            let _coll36 = new _dafny.Map();
            for (const _compr_36 of _dafny.IntegerRange(new BigNumber(713), new BigNumber(736))) {
              let _1212_v66 = _compr_36;
              if (((new BigNumber(713)).isLessThanOrEqualTo(_1212_v66)) && ((_1212_v66).isLessThan(new BigNumber(736)))) {
                _coll36.push([_module.__default.safeModuloInt(_1212_v66, _1109_v1),_1202_cf32]);
              }
            }
            return _coll36;
          }());
          let _1213_v68;
          _1213_v68 = _dafny.Set.fromElements(_1108_v0);
          let _1214_v69;
          let _nw152 = Array((new BigNumber(20)).toNumber());
          _nw152[(_dafny.ZERO).toNumber()] = _1109_v1;
          _nw152[(_dafny.ONE).toNumber()] = (_1109_v1).minus(_1109_v1);
          _nw152[(new BigNumber(2)).toNumber()] = (new BigNumber((_dafny.Set.fromElements(_1202_cf32)).length)).minus(new BigNumber(((_1204_v59).update(_1109_v1, _1109_v1)).length));
          _nw152[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1203_cf31).length), _1109_v1);
          _nw152[(new BigNumber(4)).toNumber()] = _1109_v1;
          _nw152[(new BigNumber(5)).toNumber()] = _1109_v1;
          _nw152[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_1109_v1, _1109_v1);
          _nw152[(new BigNumber(7)).toNumber()] = _1109_v1;
          _nw152[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_1109_v1);
          _nw152[(new BigNumber(9)).toNumber()] = _module.__default.fm22(_1208_v63, _module.D0.create_DC3(_1206_v61), _1202_cf32, globalState);
          _nw152[(new BigNumber(10)).toNumber()] = (_1109_v1).minus(_1109_v1);
          _nw152[(new BigNumber(11)).toNumber()] = _1109_v1;
          _nw152[(new BigNumber(12)).toNumber()] = _1109_v1;
          _nw152[(new BigNumber(13)).toNumber()] = _1109_v1;
          _nw152[(new BigNumber(14)).toNumber()] = new BigNumber(183);
          _nw152[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1109_v1, new BigNumber((_module.__default.fm37(!(_1202_cf32), _1109_v1, _1109_v1, (((_1209_v64).contains(_1109_v1)) ? ((_1209_v64).get(_1109_v1)) : (p0)), globalState)).length)));
          _nw152[(new BigNumber(16)).toNumber()] = _1109_v1;
          _nw152[(new BigNumber(17)).toNumber()] = (new BigNumber((_1211_v67).length)).minus(new BigNumber((_1203_cf31).length));
          _nw152[(new BigNumber(18)).toNumber()] = new BigNumber((_1213_v68).length);
          _nw152[(new BigNumber(19)).toNumber()] = (new BigNumber(342)).plus(new BigNumber(138));
          _1214_v69 = _nw152;
          let _index168 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_1214_v69).length));
          (_1214_v69)[_index168] = new BigNumber((_1204_v59).length);
          let _index169 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_1214_v69).length));
          (_1214_v69)[_index169] = new BigNumber(674);
          let _1215_v70;
          let _nw153 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1215_v70 = _nw153;
          _1215_v70 = _1215_v70;
          _1109_v1 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1202_cf32, _module.__default.fm6(globalState), _1202_cf32, !(_1108_v0), _1108_v0)).length));
          r1 = _1109_v1;
        } else if (_source18.is_DC18) {
          let _1216___mcc_h12 = (_source18).cf26;
          let _1217_cf26 = _1216___mcc_h12;
          let _1218_v71;
          let _init21 = ((_1219_v4) => function (_1220_i6) {
            return ((true) ? (_1219_v4) : (_1219_v4));
          })(_1112_v4);
          let _nw154 = Array((new BigNumber(14)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw154.length); _i0_21++) {
            _nw154[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _1218_v71 = _nw154;
          let _index170 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_1218_v71).length));
          (_1218_v71)[_index170] = ((_1108_v0) ? (_1112_v4) : (_1112_v4));
          let _index171 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_1218_v71).length));
          (_1218_v71)[_index171] = _1112_v4;
          let _1221_v72;
          let _init22 = ((_1222_v0) => function (_1223_i7) {
            return _module.__default.safeDivisionInt(_1223_i7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1222_v0,_1222_v0)).length));
          })(_1108_v0);
          let _nw155 = Array((_dafny.ONE).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw155.length); _i0_22++) {
            _nw155[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1221_v72 = _nw155;
          let _1224_v73;
          _1224_v73 = _dafny.Seq.UnicodeFromString("taqaliqve");
          let _index172 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_1221_v72).length));
          (_1221_v72)[_index172] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1224_v73).length), _1109_v1));
          let _index173 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_1221_v72).length));
          (_1221_v72)[_index173] = _1109_v1;
          let _1225_v74;
          let _nw156 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _1225_v74 = _nw156;
          let _1226_v75;
          _1226_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1221_v72,new BigNumber((_1224_v73).length));
          let _1227_v76;
          _1227_v76 = _module.D8.create_DC22((_1226_v75).update(_1221_v72, (_1221_v72)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_1221_v72).length))]));
          let _1228_v77;
          let _nw157 = Array((new BigNumber(3)).toNumber());
          _nw157[(_dafny.ZERO).toNumber()] = _1108_v0;
          _nw157[(_dafny.ONE).toNumber()] = _1108_v0;
          _nw157[(new BigNumber(2)).toNumber()] = _1108_v0;
          _1228_v77 = _nw157;
          let _index174 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1228_v77).length));
          (_1228_v77)[_index174] = _1108_v0;
          let _1229_v78;
          _1229_v78 = _dafny.Set.fromElements(_1108_v0, !(_1108_v0), _1108_v0, _1108_v0, _1108_v0);
          let _pat_let_tv23 = _1226_v75;
          let _index175 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_1221_v72).length));
          let _index176 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1228_v77).length));
          let _rhs180 = _1225_v74;
          let _rhs181 = !(_1108_v0);
          let _rhs182 = ((_this).fm0(_1109_v1, p0, globalState)).minus(_1109_v1);
          let _rhs183 = function (_pat_let23_0) {
            return function (_1230_dt__update__tmp_h0) {
              return function (_pat_let24_0) {
                return function (_1231_dt__update_hcf34_h0) {
                  return _module.D8.create_DC22(_1231_dt__update_hcf34_h0);
                }(_pat_let24_0);
              }(_pat_let_tv23);
            }(_pat_let23_0);
          }(_1227_v76);
          let _rhs184 = (_module.__default.fm33(_1108_v0, _1108_v0, _1108_v0, globalState)).IsDisjointFrom((_1229_v78).Intersect(_1229_v78));
          let _lhs124 = globalState;
          let _lhs125 = _1221_v72;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_1221_v72).length));
          let _lhs127 = _1228_v77;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_1228_v77).length));
          _1225_v74 = _rhs180;
          _lhs124.f3 = _rhs181;
          _lhs125[_lhs126] = _rhs182;
          _1227_v76 = _rhs183;
          _lhs127[_lhs128] = _rhs184;
          let _1232_v79;
          _1232_v79 = _module.D5.create_DC13(_1228_v77);
          _1228_v77 = (_1232_v79).dtor_cf19;
        } else {
          let _1233___mcc_h13 = (_source18).cf33;
          let _1234_cf33 = _1233___mcc_h13;
          let _1235_v80;
          _1235_v80 = _dafny.Seq.UnicodeFromString("pdxrtkj");
          r1 = _module.__default.safeDivisionInt(new BigNumber((_1235_v80).length), _module.__default.safeModuloInt(_1109_v1, _1109_v1));
          _1109_v1 = (_dafny.ZERO).minus(_1109_v1);
          let _1236_v81;
          let _nw158 = Array((new BigNumber(8)).toNumber()).fill([]);
          _1236_v81 = _nw158;
          let _1237_v82;
          let _init23 = ((_1238_v0) => function (_1239_i8) {
            return _1238_v0;
          })(_1108_v0);
          let _nw159 = Array((new BigNumber(13)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw159.length); _i0_23++) {
            _nw159[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1237_v82 = _nw159;
          let _index177 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1236_v81).length));
          (_1236_v81)[_index177] = _1237_v82;
          let _index178 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_1236_v81).length));
          (_1236_v81)[_index178] = _1237_v82;
          let _1240_v83;
          _1240_v83 = _module.D4.create_DC12(_1109_v1);
          let _1241_v84;
          _1241_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1240_v83,((_1108_v0) ? ((_dafny.ZERO).minus((_1110_v2).fm0(_1109_v1, p0, globalState))) : (_1109_v1)));
          r1 = (((_1241_v84).contains(_module.__default.fm35(_dafny.Seq.of(true), _1108_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_1244_p0) => function (_1245_i9) {
            return _1244_p0;
          })(p0)), globalState))) ? ((_1241_v84).get(_module.__default.fm35(_dafny.Seq.of(true), _1108_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_1242_p0) => function (_1243_i9) {
            return _1242_p0;
          })(p0)), globalState))) : (_module.__default.safeModuloInt(new BigNumber(279), _1109_v1)));
        }
        let _1246_v85;
        let _nw160 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _1246_v85 = _nw160;
        _1246_v85 = _1246_v85;
        let _1247_v86;
        _1247_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1109_v1,(false) === (_1108_v0));
        let _1248_v87;
        _1248_v87 = _dafny.Seq.UnicodeFromString("lgyduwq");
        let _rhs185 = _1247_v86;
        let _rhs186 = new BigNumber((_1248_v87).length);
        let _rhs187 = _1246_v85;
        _1247_v86 = _rhs185;
        r1 = _rhs186;
        _1246_v85 = _rhs187;
        let _1249_v88;
        _1249_v88 = _dafny.Map.Empty.slice().updateUnsafe((_1109_v1).isLessThanOrEqualTo(_1109_v1),_1108_v0);
        _1249_v88 = (_1249_v88).update(_1108_v0, false);
        if (_1108_v0) {
          let _1250_v89;
          _1250_v89 = _dafny.Seq.of(_1108_v0, _1108_v0, _1108_v0, _1108_v0, true);
          r1 = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1250_v89).length)))).length);
          let _1251_v90;
          _1251_v90 = _module.D1.create_DC5(_1108_v0, _1248_v87, _1250_v89);
          let _1252_v91;
          _1252_v91 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(760),_1251_v90));
          let _1253_v92;
          _1253_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1109_v1,_1251_v90);
          _1252_v91 = (_1252_v91).Union(_dafny.Set.fromElements((_1253_v92).update(_1109_v1, _1251_v90)));
          let _1254_v93;
          let _nw161 = new _module.C3();
          _nw161.__ctor((_dafny.Set.fromElements(new BigNumber((_1248_v87).length), new BigNumber((_dafny.Seq.UnicodeFromString("kenqbixrf")).length))).IsProperSubsetOf(_1112_v4), _1109_v1);
          _1254_v93 = _nw161;
          _1250_v89 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1250_v89, _1250_v89), _1250_v89);
          let _1255_v94;
          _1255_v94 = _dafny.Seq.of(_1109_v1);
          let _1256_v95;
          let _nw162 = Array((new BigNumber(29)).toNumber()).fill(_module.D0.Default());
          _1256_v95 = _nw162;
          let _1257_v96;
          _1257_v96 = _dafny.Set.fromElements((_module.D5.create_DC14((_1254_v93).fm13((_1254_v93).f17, _1255_v94, (_1254_v93).f17, _1108_v0, globalState), _1108_v0, _1256_v95)).dtor_cf20, _module.__default.fm19((_1254_v93).f17, (_1254_v93).f16, globalState));
          let _1258_v97;
          _1258_v97 = _dafny.Seq.of(_1257_v96, _1257_v96, _1257_v96, _1257_v96);
          _1257_v96 = ((!((_1254_v93).f17).isEqualTo(new BigNumber(-723))) ? ((_1257_v96).Union(_1257_v96)) : (((_1258_v97)[_module.__default.safeIndex(_1109_v1, new BigNumber((_1258_v97).length))]).Difference(_dafny.Set.fromElements((_1254_v93).f16))));
        } else {
          (globalState).f13 = _1108_v0;
          r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1109_v1));
          let _1259_v98;
          _1259_v98 = _dafny.Seq.of(_1108_v0, _1108_v0);
          r0 = _dafny.areEqual(_dafny.Seq.Concat(_1259_v98, _1259_v98), _dafny.Seq.Concat(_1259_v98, _1259_v98));
          r1 = _1109_v1;
          let _rhs188 = _dafny.Seq.UnicodeFromString("jmukv");
          let _rhs189 = _dafny.Seq.update(_dafny.Seq.update(_1248_v87, _module.__default.safeIndex(_1109_v1, new BigNumber((_1248_v87).length)), new _dafny.CodePoint('c'.codePointAt(0))), _module.__default.safeIndex(((_dafny.ZERO).minus(_1109_v1)).multipliedBy(_1109_v1), new BigNumber((_dafny.Seq.update(_1248_v87, _module.__default.safeIndex(_1109_v1, new BigNumber((_1248_v87).length)), new _dafny.CodePoint('c'.codePointAt(0)))).length)), p0);
          _1248_v87 = _rhs188;
          _1248_v87 = _rhs189;
        }
      }
      let _1260_v99;
      _1260_v99 = _module.D0.create_DC0(_1109_v1);
      let _1261_v100;
      _1261_v100 = _module.D0.create_DC3(_1260_v99);
      let _1262_v101;
      _1262_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1108_v0,_1261_v100);
      let _1263_v102;
      _1263_v102 = _module.D7.create_DC20(_dafny.Seq.UnicodeFromString("khrd"), _1108_v0);
      let _1264_v103;
      _1264_v103 = _module.D2.create_DC7(_1108_v0, _1262_v101, _1109_v1, (_1263_v102).dtor_cf32);
      let _1265_v104;
      let _nw163 = Array((new BigNumber(19)).toNumber());
      _nw163[(_dafny.ZERO).toNumber()] = _1108_v0;
      _nw163[(_dafny.ONE).toNumber()] = !(_1108_v0);
      _nw163[(new BigNumber(2)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(3)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(4)).toNumber()] = false;
      _nw163[(new BigNumber(5)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(6)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(7)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(8)).toNumber()] = !(_1108_v0);
      _nw163[(new BigNumber(9)).toNumber()] = false;
      _nw163[(new BigNumber(10)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(11)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(12)).toNumber()] = (_this).fm3(true, _1108_v0, new BigNumber(248), globalState);
      _nw163[(new BigNumber(13)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(14)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(15)).toNumber()] = (_1264_v103).dtor_cf11;
      _nw163[(new BigNumber(16)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(17)).toNumber()] = _1108_v0;
      _nw163[(new BigNumber(18)).toNumber()] = _1108_v0;
      _1265_v104 = _nw163;
      let _1266_v105;
      let _nw164 = Array((new BigNumber(9)).toNumber());
      _nw164[(_dafny.ZERO).toNumber()] = _1265_v104;
      _nw164[(_dafny.ONE).toNumber()] = _1265_v104;
      _nw164[(new BigNumber(2)).toNumber()] = _1265_v104;
      _nw164[(new BigNumber(3)).toNumber()] = _1265_v104;
      _nw164[(new BigNumber(4)).toNumber()] = _1265_v104;
      _nw164[(new BigNumber(5)).toNumber()] = _1265_v104;
      _nw164[(new BigNumber(6)).toNumber()] = _1265_v104;
      _nw164[(new BigNumber(7)).toNumber()] = _1265_v104;
      _nw164[(new BigNumber(8)).toNumber()] = _1265_v104;
      _1266_v105 = _nw164;
      let _1267_v106;
      _1267_v106 = _dafny.Seq.of(_1265_v104, _1265_v104);
      let _index179 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1266_v105).length));
      (_1266_v105)[_index179] = (_1267_v106)[_module.__default.safeIndex(new BigNumber(-808), new BigNumber((_1267_v106).length))];
      let _index180 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1266_v105).length));
      (_1266_v105)[_index180] = _1265_v104;
      let _index181 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
      (_1265_v104)[_index181] = _1108_v0;
      let _index182 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
      (_1265_v104)[_index182] = !(_1108_v0) || (((_1108_v0) ? (!(_1108_v0)) : (_1108_v0)));
      let _1268_v107;
      let _nw165 = new _module.C3();
      _nw165.__ctor((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))], _1109_v1);
      _1268_v107 = _nw165;
      let _1269_i10;
      _1269_i10 = _dafny.ZERO;
      L8: {
        while ((_1268_v107).f16) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1269_i10)) {
              break L8;
            }
            _1269_i10 = (_1269_i10).plus(_dafny.ONE);
            let _1270_v108;
            _1270_v108 = _dafny.Seq.UnicodeFromString("pjglis");
            _1270_v108 = ((_1108_v0) ? (_1270_v108) : (_1270_v108));
            let _1271_v109;
            _1271_v109 = _dafny.Seq.of(_1270_v108);
            let _1272_v110;
            _1272_v110 = _dafny.Seq.of(_module.D0.create_DC1((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))], _1271_v109, _1270_v108));
            let _1273_v111;
            _1273_v111 = _module.D0.create_DC1((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))], _1271_v109, _1270_v108);
            let _1274_v112;
            _1274_v112 = _dafny.MultiSet.fromElements(_1273_v111, _1273_v111, _1273_v111);
            let _pat_let_tv24 = _1274_v112;
            let _source19 = function (_pat_let25_0) {
              return function (_1275_dt__update__tmp_h1) {
                return function (_pat_let26_0) {
                  return function (_1276_dt__update_hcf15_h0) {
                    return _module.D2.create_DC8(_1276_dt__update_hcf15_h0);
                  }(_pat_let26_0);
                }(_module.D2.create_DC6(_pat_let_tv24));
              }(_pat_let25_0);
            }(_module.D2.create_DC8(_module.D2.create_DC6(_dafny.MultiSet.FromArray(_1272_v110))));
            if (_source19.is_DC7) {
              let _1277___mcc_h14 = (_source19).cf11;
              let _1278___mcc_h15 = (_source19).cf12;
              let _1279___mcc_h16 = (_source19).cf13;
              let _1280___mcc_h17 = (_source19).cf14;
              let _1281_cf14 = _1280___mcc_h17;
              let _1282_cf13 = _1279___mcc_h16;
              let _1283_cf12 = _1278___mcc_h15;
              let _1284_cf11 = _1277___mcc_h14;
              let _index183 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
              (_1265_v104)[_index183] = true;
              let _index184 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
              (_1265_v104)[_index184] = _1281_cf14;
              let _1285_v113;
              _1285_v113 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1268_v107).f17);
              let _1286_v114;
              _1286_v114 = _module.D9.create_DC25(_1285_v113);
              let _1287_v115;
              let _nw166 = new _module.C4();
              _nw166.__ctor(_1109_v1, (_1286_v114).dtor_cf38);
              _1287_v115 = _nw166;
              let _1288_v116;
              _1288_v116 = _dafny.Seq.of(_1281_cf14, true);
              let _index185 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
              (_1265_v104)[_index185] = (_1288_v116)[_module.__default.safeIndex((_1268_v107).f17, new BigNumber((_1288_v116).length))];
            } else if (_source19.is_DC6) {
              let _1289___mcc_h18 = (_source19).cf10;
              let _1290_cf10 = _1289___mcc_h18;
              let _1291_v117;
              _1291_v117 = _module.D9.create_DC27((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))], (_1268_v107).f17, (_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))]);
              let _1292_v118;
              _1292_v118 = _dafny.MultiSet.fromElements((_1268_v107).f16);
              let _rhs190 = new BigNumber(((((_1268_v107).f16) ? (_1292_v118) : (_1292_v118))).cardinality());
              let _rhs191 = _1291_v117;
              _1109_v1 = _rhs190;
              _1291_v117 = _rhs191;
              let _1293_v119;
              _1293_v119 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_module.__default.fm30(globalState)).length))).length));
              let _index186 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
              (_1265_v104)[_index186] = (_1293_v119).IsSubsetOf(_dafny.MultiSet.fromElements((_1268_v107).f17));
              _1270_v108 = _dafny.Seq.Concat(_1270_v108, _dafny.Seq.Concat(_1270_v108, _1270_v108));
              let _1294_v120;
              _1294_v120 = _dafny.Seq.of(_1268_v107);
              let _index187 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
              let _rhs192 = _dafny.Seq.IsProperPrefixOf(_1294_v120, _1294_v120);
              let _rhs193 = (_1268_v107).f17;
              let _lhs129 = _1265_v104;
              let _lhs130 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length));
              _lhs129[_lhs130] = _rhs192;
              r1 = _rhs193;
            } else {
              let _1295___mcc_h19 = (_source19).cf15;
              let _1296_cf15 = _1295___mcc_h19;
              let _1297_v121;
              _1297_v121 = _module.D5.create_DC13((_1266_v105)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_1266_v105).length))]);
              let _1298_v122;
              _1298_v122 = _module.D5.create_DC13((_1297_v121).dtor_cf19);
              let _index188 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1266_v105).length));
              (_1266_v105)[_index188] = (_1297_v121).dtor_cf19;
              (globalState).f5 = ((_1268_v107).f17).isLessThanOrEqualTo((_1268_v107).f17);
              let _1299_v123;
              _1299_v123 = _dafny.Map.Empty.slice().updateUnsafe((_1268_v107).f17,(_1268_v107).f17);
              r1 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1109_v1), new BigNumber(((_this).fm1(_1299_v123, globalState)).length))).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1270_v108).length)));
              let _1300_v124;
              _1300_v124 = _1266_v105;
              _1266_v105 = (_1300_v124);
            }
            let _1301_v125;
            _1301_v125 = _dafny.Set.fromElements((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))], _module.__default.fm19((_1268_v107).f17, _1108_v0, globalState));
            _1301_v125 = _1301_v125;
            let _1302_v126;
            let _nw167 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
            _1302_v126 = _nw167;
            let _1303_v127;
            _1303_v127 = _dafny.Map.Empty.slice().updateUnsafe(_1302_v126,(_1268_v107).f17);
            let _1304_v128;
            _1304_v128 = _module.D8.create_DC22((_1303_v127).update(_1302_v126, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), function (_1305_i11) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})).length)));
            _1304_v128 = _1304_v128;
          }
        }
      }
      let _1306_v129;
      let _nw168 = Array((new BigNumber(19)).toNumber());
      _1306_v129 = _nw168;
      let _1307_v130;
      let _nw169 = new _module.C2();
      _nw169.__ctor();
      _1307_v130 = _nw169;
      let _index189 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1306_v129).length));
      (_1306_v129)[_index189] = _1307_v130;
      let _1308_v131;
      let _nw170 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
      _1308_v131 = _nw170;
      let _index190 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1308_v131).length));
      (_1308_v131)[_index190] = _1112_v4;
      let _index191 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1306_v129).length));
      let _index192 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1308_v131).length));
      let _rhs194 = _1307_v130;
      let _rhs195 = _1112_v4;
      let _lhs131 = _1306_v129;
      let _lhs132 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1306_v129).length));
      let _lhs133 = _1308_v131;
      let _lhs134 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_1308_v131).length));
      _lhs131[_lhs132] = _rhs194;
      _lhs133[_lhs134] = _rhs195;
      let _1309_v132;
      _1309_v132 = _dafny.Seq.of(!(_module.__default.fm6(globalState)), (_1268_v107).f16, (((_1268_v107).f16) ? ((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))]) : (_1108_v0)));
      r0 = (_1309_v132)[_module.__default.safeIndex((_1268_v107).f17, new BigNumber((_1309_v132).length))];
      let _1310_v133;
      _1310_v133 = _dafny.Map.Empty.slice().updateUnsafe(true,_1309_v132);
      r1 = new BigNumber(((((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))]) ? (_dafny.Map.Empty.slice().updateUnsafe((_1265_v104)[_module.__default.safeIndex(new BigNumber(879), new BigNumber((_1265_v104).length))],_1309_v132)) : (_1310_v133))).length);
      return [r0, r1];
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.MultiSet.Empty;
      let r2 = false;
      let r3 = [];
      let _1311_v0;
      _1311_v0 = _dafny.Seq.UnicodeFromString("ypgvmngw");
      let _1312_v1;
      _1312_v1 = _dafny.MultiSet.fromElements(new BigNumber((_1311_v0).length));
      let _1313_v2;
      _1313_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1312_v1);
      let _1314_v3;
      _1314_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1313_v2,_1312_v1);
      let _1315_v4;
      _1315_v4 = new BigNumber(-81);
      let _1316_v5;
      _1316_v5 = _dafny.Seq.of(_1315_v4);
      _1314_v3 = (_1314_v3).update(_1313_v2, _dafny.MultiSet.FromArray(((p1) ? (_1316_v5) : (_1316_v5))));
      _1315_v4 = _module.__default.fm11((p1) && (true), globalState);
      let _1317_v6;
      let _nw171 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
      _1317_v6 = _nw171;
      let _1318_v7;
      _1318_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_1319_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }));
      let _index193 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_1317_v6).length));
      (_1317_v6)[_index193] = (_1318_v7).Merge(_1318_v7);
      let _1320_v8;
      _1320_v8 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1321_v9;
      _1321_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1316_v5,_1315_v4);
      let _1322_v10;
      _1322_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1315_v4,new BigNumber((_1312_v1).cardinality()));
      let _1323_v11;
      _1323_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1322_v10).length),_1315_v4);
      let _index194 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_1317_v6).length));
      let _rhs196 = (_dafny.ZERO).minus(_1315_v4);
      let _rhs197 = (_1318_v7).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_1311_v0));
      let _rhs198 = _1320_v8;
      let _rhs199 = _1321_v9;
      let _rhs200 = (_this).fm1(_1323_v11, globalState);
      let _lhs135 = _1317_v6;
      let _lhs136 = _module.__default.safeIndex(new BigNumber(73), new BigNumber((_1317_v6).length));
      _1315_v4 = _rhs196;
      _lhs135[_lhs136] = _rhs197;
      _1320_v8 = _rhs198;
      _1321_v9 = _rhs199;
      _1311_v0 = _rhs200;
      _1320_v8 = _1320_v8;
      let _1324_v12;
      _1324_v12 = _dafny.Set.fromElements(_1315_v4, new BigNumber((_dafny.Seq.UnicodeFromString("cdtw")).length));
      _1324_v12 = _1324_v12;
      if (p1) {
        _1311_v0 = _dafny.Seq.Concat(_1311_v0, _1311_v0);
        let _1325_v13;
        _1325_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1315_v4,p1);
        _1325_v13 = ((_1325_v13).Merge(_1325_v13)).Merge(_1325_v13);
        _1311_v0 = _1311_v0;
        let _1326_v14;
        _1326_v14 = _dafny.Seq.of(_1316_v5, _1316_v5);
        let _1327_v15;
        _1327_v15 = _dafny.Seq.of(_dafny.Seq.of(_1316_v5), _dafny.Seq.Concat(_1326_v14, _1326_v14), _module.__default.fm38(globalState));
        let _1328_v16;
        _1328_v16 = _module.D0.create_DC0(new BigNumber((p0).length));
        let _1329_v17;
        _1329_v17 = _module.D0.create_DC3(_1328_v16);
        let _1330_v18;
        _1330_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1329_v17);
        _1326_v14 = (_1327_v15)[_module.__default.safeIndex(_module.__default.fm22(_module.D2.create_DC7(p1, _1330_v18, _1315_v4, p1), _1329_v17, p1, globalState), new BigNumber((_1327_v15).length))];
        let _1331_v19;
        _1331_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
        let _1332_v20;
        let _nw172 = Array((new BigNumber(17)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = p1;
        _nw172[(_dafny.ONE).toNumber()] = (((_1331_v19).contains(p1)) ? ((_1331_v19).get(p1)) : (p1));
        _nw172[(new BigNumber(2)).toNumber()] = p1;
        _nw172[(new BigNumber(3)).toNumber()] = p1;
        _nw172[(new BigNumber(4)).toNumber()] = true;
        _nw172[(new BigNumber(5)).toNumber()] = true;
        _nw172[(new BigNumber(6)).toNumber()] = !(p1);
        _nw172[(new BigNumber(7)).toNumber()] = true;
        _nw172[(new BigNumber(8)).toNumber()] = (_1324_v12).IsProperSubsetOf(_1324_v12);
        _nw172[(new BigNumber(9)).toNumber()] = p1;
        _nw172[(new BigNumber(10)).toNumber()] = p1;
        _nw172[(new BigNumber(11)).toNumber()] = !((p1) === (p1));
        _nw172[(new BigNumber(12)).toNumber()] = p1;
        _nw172[(new BigNumber(13)).toNumber()] = !(_module.__default.fm19(_1315_v4, p1, globalState)) || (p1);
        _nw172[(new BigNumber(14)).toNumber()] = p1;
        _nw172[(new BigNumber(15)).toNumber()] = ((_dafny.ZERO).minus(_1315_v4)).isLessThan(_1315_v4);
        _nw172[(new BigNumber(16)).toNumber()] = !(p1);
        _1332_v20 = _nw172;
        let _index195 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1332_v20).length));
        (_1332_v20)[_index195] = false;
        let _index196 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_1332_v20).length));
        (_1332_v20)[_index196] = p1;
      } else {
        let _1333_v21;
        _1333_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1320_v8,_1315_v4);
        let _1334_v22;
        let _nw173 = new _module.C4();
        _nw173.__ctor((_1315_v4).multipliedBy(_1315_v4), _1333_v21);
        _1334_v22 = _nw173;
        let _1335_v23;
        _1335_v23 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus((_1334_v22).f14)),p0);
        (globalState).f5 = !_dafny.areEqual(p0, (((_1335_v23).contains((_1334_v22).f14)) ? ((_1335_v23).get((_1334_v22).f14)) : (_dafny.Seq.of(p1))));
        (globalState).f13 = p1;
        let _1336_v24;
        _1336_v24 = _module.D4.create_DC11();
        let _source20 = _1336_v24;
        if (_source20.is_DC11) {
          let _1337_v25;
          _1337_v25 = _module.D9.create_DC26(p1, _1320_v8, (p0)[_module.__default.safeIndex((_1334_v22).f14, new BigNumber((p0).length))]);
          let _1338_v26;
          _1338_v26 = _dafny.MultiSet.fromElements(p1, (_1337_v25).dtor_cf41);
          _1338_v26 = (_1338_v26).Union((_1338_v26).Intersect(_module.__default.fm39(p1, globalState)));
          let _1339_v27;
          _1339_v27 = _dafny.Set.fromElements(_dafny.Seq.Concat(_1311_v0, _1311_v0));
          _1339_v27 = _1339_v27;
          let _1340_v28;
          let _nw174 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1340_v28 = _nw174;
          let _1341_v29;
          _1341_v29 = _dafny.Seq.of(_1340_v28, _1340_v28);
          let _1342_v30;
          _1342_v30 = _dafny.Set.fromElements(p1, p1);
          let _1343_v31;
          _1343_v31 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1342_v30);
          let _1344_v32;
          _1344_v32 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1345_v33;
          _1345_v33 = _module.D8.create_DC23(p1, p1, new BigNumber((_1344_v32).length));
          let _1346_v34;
          _1346_v34 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1342_v30).length));
          let _1347_v35;
          let _nw175 = Array((new BigNumber(19)).toNumber());
          _nw175[(_dafny.ZERO).toNumber()] = p1;
          _nw175[(_dafny.ONE).toNumber()] = !(((_dafny.ZERO).minus((_dafny.ZERO).minus((_1334_v22).f14))).isLessThanOrEqualTo(new BigNumber((_1341_v29).length)));
          _nw175[(new BigNumber(2)).toNumber()] = p1;
          _nw175[(new BigNumber(3)).toNumber()] = !(p1);
          _nw175[(new BigNumber(4)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_1316_v5, _module.__default.safeIndex((_1334_v22).f14, new BigNumber((_1316_v5).length)), new BigNumber((_1343_v31).length)), _dafny.Seq.of(_1315_v4, (_1334_v22).f14));
          _nw175[(new BigNumber(5)).toNumber()] = (_1345_v33).dtor_cf36;
          _nw175[(new BigNumber(6)).toNumber()] = (_1346_v34).contains(p1);
          _nw175[(new BigNumber(7)).toNumber()] = (_1315_v4).isLessThanOrEqualTo(_1315_v4);
          _nw175[(new BigNumber(8)).toNumber()] = p1;
          _nw175[(new BigNumber(9)).toNumber()] = p1;
          _nw175[(new BigNumber(10)).toNumber()] = !(p1) || (true);
          _nw175[(new BigNumber(11)).toNumber()] = p1;
          _nw175[(new BigNumber(12)).toNumber()] = p1;
          _nw175[(new BigNumber(13)).toNumber()] = !(p1);
          _nw175[(new BigNumber(14)).toNumber()] = !(p1);
          _nw175[(new BigNumber(15)).toNumber()] = !(p1);
          _nw175[(new BigNumber(16)).toNumber()] = (p1) || (p1);
          _nw175[(new BigNumber(17)).toNumber()] = false;
          _nw175[(new BigNumber(18)).toNumber()] = p1;
          _1347_v35 = _nw175;
          let _index197 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_1347_v35).length));
          (_1347_v35)[_index197] = p1;
          let _index198 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_1347_v35).length));
          (_1347_v35)[_index198] = _module.__default.fm19((_1334_v22).f14, p1, globalState);
          let _1348_v36;
          _1348_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1316_v5)).cardinality()),_1341_v29);
          let _rhs201 = _dafny.areEqual(_1341_v29, (((_1348_v36).contains((_1334_v22).f14)) ? ((_1348_v36).get((_1334_v22).f14)) : (_dafny.Seq.of(_1340_v28))));
          let _rhs202 = _dafny.Seq.Concat(_1311_v0, _1311_v0);
          let _lhs137 = globalState;
          _lhs137.f3 = _rhs201;
          _1311_v0 = _rhs202;
        } else if (_source20.is_DC12) {
          let _1349___mcc_h0 = (_source20).cf18;
          let _1350_cf18 = _1349___mcc_h0;
          let _1351_v37;
          let _nw176 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1351_v37 = _nw176;
          let _1352_v38;
          _1352_v38 = _dafny.Seq.of(_1351_v37, _1351_v37, _1351_v37);
          _1315_v4 = new BigNumber((_1352_v38).length);
          let _index199 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1351_v37).length));
          (_1351_v37)[_index199] = _1350_cf18;
          let _index200 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1351_v37).length));
          (_1351_v37)[_index200] = _module.__default.fm22(_module.__default.fm29((_1316_v5)[_module.__default.safeIndex(new BigNumber(-759), new BigNumber((_1316_v5).length))], _1350_cf18, _dafny.Seq.UnicodeFromString("kkiqvf"), _1350_cf18, globalState), _module.D0.create_DC3(_module.D0.create_DC0(_1350_cf18)), p1, globalState);
          let _index201 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1351_v37).length));
          (_1351_v37)[_index201] = _module.__default.fm11(p1, globalState);
          let _index202 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1351_v37).length));
          (_1351_v37)[_index202] = _1350_cf18;
        } else {
          let _1353___mcc_h1 = (_source20).cf17;
          let _1354_cf17 = _1353___mcc_h1;
          (globalState).f3 = p1;
          let _1355_v39;
          let _init24 = ((_1356_cf17) => function (_1357_i1) {
            return (_1357_i1).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(317)), ((_1358_cf17) => function (_1359_i2) {
              return new BigNumber((_1358_cf17).length);
            })(_1356_cf17))).length));
          })(_1354_cf17);
          let _nw177 = Array((new BigNumber(9)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw177.length); _i0_24++) {
            _nw177[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1355_v39 = _nw177;
          let _index203 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1355_v39).length));
          (_1355_v39)[_index203] = (_dafny.ZERO).minus(((p1) ? ((_1334_v22).f14) : ((_1334_v22).f14)));
          let _index204 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1355_v39).length));
          (_1355_v39)[_index204] = (_1334_v22).f14;
          let _1360_v40;
          let _nw178 = Array((new BigNumber(3)).toNumber());
          _nw178[(_dafny.ZERO).toNumber()] = !(p1);
          _nw178[(_dafny.ONE).toNumber()] = _dafny.Seq.contains(_1311_v0, _1320_v8);
          _nw178[(new BigNumber(2)).toNumber()] = p1;
          _1360_v40 = _nw178;
          let _index205 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1360_v40).length));
          (_1360_v40)[_index205] = p1;
          let _index206 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1355_v39).length));
          let _index207 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1355_v39).length));
          let _index208 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1360_v40).length));
          let _rhs203 = (_1334_v22).f14;
          let _rhs204 = (new BigNumber((_dafny.Seq.Concat((((_1318_v7).contains(p1)) ? ((_1318_v7).get(p1)) : ((_this).fm1(_1323_v11, globalState))), _dafny.Seq.UnicodeFromString("wucl"))).length)).plus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()));
          let _rhs205 = p1;
          let _rhs206 = (_1315_v4).isEqualTo(_1315_v4);
          let _lhs138 = _1355_v39;
          let _lhs139 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1355_v39).length));
          let _lhs140 = _1355_v39;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1355_v39).length));
          let _lhs142 = _1360_v40;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1360_v40).length));
          _lhs138[_lhs139] = _rhs203;
          _lhs140[_lhs141] = _rhs204;
          _lhs142[_lhs143] = _rhs205;
          r2 = _rhs206;
          let _1361_v41;
          _1361_v41 = _dafny.Map.Empty.slice().updateUnsafe((_1360_v40)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1360_v40).length))],p1);
          let _1362_v42;
          _1362_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1320_v8,_1361_v41);
          let _1363_v43;
          let _nw179 = new _module.C4();
          _nw179.__ctor((_1334_v22).f14, _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm23((_1360_v40)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1360_v40).length))], _1362_v42, (_1360_v40)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1360_v40).length))], (_1334_v22).f14, globalState),new BigNumber(848)));
          _1363_v43 = _nw179;
          (globalState).f13 = (_1360_v40)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1360_v40).length))];
        }
        (_1334_v22).m5(((true) ? (_1311_v0) : (_1311_v0)), globalState);
      }
      let _1364_v44;
      _1364_v44 = _dafny.Seq.of(_1311_v0);
      let _1365_v45;
      _1365_v45 = _module.D0.create_DC1(p1, _1364_v44, _1311_v0);
      let _1366_v46;
      _1366_v46 = _dafny.Map.Empty.slice().updateUnsafe((_1365_v45).dtor_cf3,_1323_v11);
      r0 = (((_1366_v46).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(67)), ((_1369_v8) => function (_1370_i3) {
        return _1369_v8;
      })(_1320_v8)))) ? ((_1366_v46).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(67)), ((_1367_v8) => function (_1368_i3) {
        return _1367_v8;
      })(_1320_v8)))) : ((_module.__default.fm31(globalState)).Merge(_1323_v11)));
      r1 = _module.__default.fm25(p1, _module.__default.fm30(globalState), globalState);
      r2 = p1;
      let _1371_v47;
      let _nw180 = Array((new BigNumber(16)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1371_v47 = _nw180;
      r3 = _1371_v47;
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
    fm0(p0, p1, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.D0.create_DC0((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, true, false)).length))))).dtor_cf0));
    };
    fm1(p0, globalState) {
      let _this = this;
      return (_module.D0.create_DC1(false, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cfv"), _dafny.Seq.UnicodeFromString("lo")), _dafny.Seq.UnicodeFromString("lohm"))).dtor_cf3;
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _1372_v0;
      _1372_v0 = new BigNumber(57);
      let _1373_v1;
      _1373_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1372_v0,_1372_v0);
      _1373_v1 = (_1373_v1).update(_1372_v0, new BigNumber(31));
      let _1374_v2;
      _1374_v2 = false;
      let _1375_v3;
      _1375_v3 = _module.D0.create_DC0(new BigNumber(664));
      let _1376_v4;
      _1376_v4 = _module.D0.create_DC3(_1375_v3);
      let _pat_let_tv25 = _1375_v3;
      (globalState).f13 = _module.__default.fm2(p0, ((_1374_v2) ? (function (_pat_let27_0) {
        return function (_1377_dt__update__tmp_h0) {
          return function (_pat_let28_0) {
            return function (_1378_dt__update_hcf5_h0) {
              return _module.D0.create_DC3(_1378_dt__update_hcf5_h0);
            }(_pat_let28_0);
          }(_pat_let_tv25);
        }(_pat_let27_0);
      }(_1376_v4)) : (_1376_v4)), _module.__default.safeDivisionInt(_1372_v0, (_dafny.ZERO).minus(_1372_v0)), globalState);
      _1372_v0 = (_this).fm0((_this).fm0(_1372_v0, p0, globalState), p0, globalState);
      let _1379_v5;
      _1379_v5 = _dafny.Seq.UnicodeFromString("iqufbbgna");
      let _1380_v6;
      _1380_v6 = _dafny.Seq.of(_1379_v5, _1379_v5);
      let _1381_v7;
      _1381_v7 = _module.D0.create_DC1(_1374_v2, _1380_v6, ((_1374_v2) ? (_1379_v5) : (_1379_v5)));
      let _source21 = _1381_v7;
      if (_source21.is_DC1) {
        let _1382___mcc_h0 = (_source21).cf1;
        let _1383___mcc_h1 = (_source21).cf2;
        let _1384___mcc_h2 = (_source21).cf3;
        let _1385_cf3 = _1384___mcc_h2;
        let _1386_cf2 = _1383___mcc_h1;
        let _1387_cf1 = _1382___mcc_h0;
        let _1388_v8;
        let _nw181 = new _module.C5();
        _nw181.__ctor();
        _1388_v8 = _nw181;
        let _1389_v9;
        let _nw182 = new _module.C5();
        _nw182.__ctor();
        _1389_v9 = _nw182;
        let _1390_v10;
        let _init25 = ((_1391_v2) => function (_1392_i0) {
          return (_module.D0.create_DC2(_1391_v2)).dtor_cf4;
        })(_1374_v2);
        let _nw183 = Array((new BigNumber(13)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw183.length); _i0_25++) {
          _nw183[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _1390_v10 = _nw183;
        let _index209 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length));
        (_1390_v10)[_index209] = _1374_v2;
        let _1393_v11;
        _1393_v11 = _dafny.Seq.of(_1387_cf1, _1387_cf1, _1374_v2, _1374_v2);
        let _index210 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length));
        let _rhs207 = _1387_cf1;
        let _rhs208 = ((_1387_cf1) ? (_1389_v9) : (((_1387_cf1) ? (_1389_v9) : (_1389_v9))));
        let _rhs209 = _1374_v2;
        let _rhs210 = (_1374_v2) && ((new BigNumber((_1393_v11).length)).isLessThan(new BigNumber((_dafny.Seq.of(_1372_v0, new BigNumber(463), _1372_v0)).length)));
        let _lhs144 = globalState;
        let _lhs145 = _1390_v10;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length));
        _1374_v2 = _rhs207;
        _1389_v9 = _rhs208;
        _lhs144.f3 = _rhs209;
        _lhs145[_lhs146] = _rhs210;
        if (_1374_v2) {
          _1393_v11 = _dafny.Seq.update(_module.__default.fm27((_1372_v0).minus(_1372_v0), (_1372_v0).multipliedBy(_1372_v0), (_1372_v0).plus(_1372_v0), globalState), _module.__default.safeIndex(_1372_v0, new BigNumber((_module.__default.fm27((_1372_v0).minus(_1372_v0), (_1372_v0).multipliedBy(_1372_v0), (_1372_v0).plus(_1372_v0), globalState)).length)), (_1390_v10)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length))]);
          let _1394_v12;
          _1394_v12 = new _dafny.CodePoint('o'.codePointAt(0));
          _1394_v12 = p0;
          (globalState).f3 = !(_1374_v2);
          let _1395_v13;
          let _nw184 = new _module.C5();
          _nw184.__ctor();
          _1395_v13 = _nw184;
          let _1396_v14;
          _1396_v14 = _dafny.Set.fromElements(_1372_v0, _1372_v0);
          let _1397_v15;
          _1397_v15 = _dafny.Seq.of(new BigNumber((_1396_v14).length), new BigNumber(83));
          _1397_v15 = _dafny.Seq.Concat(_1397_v15, _1397_v15);
        } else {
          _1372_v0 = _1372_v0;
          r1 = _1372_v0;
          let _1398_v16;
          _1398_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1390_v10)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length))],_1376_v4);
          let _1399_v17;
          _1399_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1372_v0,_1374_v2);
          let _1400_v18;
          let _nw185 = new _module.C2();
          _nw185.__ctor();
          _1400_v18 = _nw185;
          let _1401_v19;
          _1401_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1399_v17).length),_1400_v18);
          let _1402_v20;
          _1402_v20 = _module.D2.create_DC7(true, _1398_v16, new BigNumber(((_1401_v19).update(_1372_v0, _1400_v18)).length), _1387_cf1);
          let _1403_v21;
          _1403_v21 = _module.D2.create_DC7((_1402_v20).dtor_cf14, _1398_v16, new BigNumber((_dafny.Seq.UnicodeFromString("vwxlge")).length), _1374_v2);
          _1372_v0 = _module.__default.fm22(_1403_v21, _1376_v4, true, globalState);
          let _1404_v22;
          _1404_v22 = _dafny.Set.fromElements(!(_1387_cf1), _1374_v2, _1374_v2);
          let _1405_v23;
          _1405_v23 = _dafny.Set.fromElements(_1404_v22);
          let _1406_v24;
          _1406_v24 = _module.D6.create_DC16(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-888)), ((_1407_p0) => function (_1408_i1) {
  return _1407_p0;
})(p0))).length));
          let _index211 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length));
          let _rhs211 = (_1405_v23).IsDisjointFrom(_dafny.Set.fromElements(_module.__default.fm33(_1387_cf1, _1387_cf1, (_1390_v10)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length))], globalState), _1404_v22, _1404_v22, _1404_v22, _1404_v22));
          let _rhs212 = (_1372_v0).multipliedBy(_1372_v0);
          let _rhs213 = _1387_cf1;
          let _rhs214 = (_1406_v24).dtor_cf24;
          let _rhs215 = (_1390_v10)[_module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length))];
          let _lhs147 = globalState;
          let _lhs148 = _1390_v10;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1390_v10).length));
          let _lhs150 = globalState;
          _lhs147.f5 = _rhs211;
          r1 = _rhs212;
          _lhs148[_lhs149] = _rhs213;
          r1 = _rhs214;
          _lhs150.f3 = _rhs215;
          r1 = (_1372_v0).plus(_1372_v0);
        }
        let _1409_v25;
        let _nw186 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1409_v25 = _nw186;
        r1 = new BigNumber((_dafny.Set.fromElements(_1409_v25, _1409_v25, _1409_v25, _1409_v25)).length);
      } else if (_source21.is_DC2) {
        let _1410___mcc_h3 = (_source21).cf4;
        let _1411_cf4 = _1410___mcc_h3;
        let _1412_v26;
        let _nw187 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _1412_v26 = _nw187;
        let _1413_v27;
        _1413_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1412_v26,new BigNumber(636));
        let _1414_v28;
        _1414_v28 = _module.D8.create_DC22(_1413_v27);
        let _1415_v29;
        _1415_v29 = _dafny.Seq.of(_1372_v0);
        let _pat_let_tv26 = _1412_v26;
        let _pat_let_tv27 = _1372_v0;
        let _rhs216 = function (_pat_let29_0) {
          return function (_1416_dt__update__tmp_h1) {
            return function (_pat_let30_0) {
              return function (_1417_dt__update_hcf34_h0) {
                return _module.D8.create_DC22(_1417_dt__update_hcf34_h0);
              }(_pat_let30_0);
            }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv26,_pat_let_tv27));
          }(_pat_let29_0);
        }(_1414_v28);
        let _rhs217 = (_dafny.ZERO).minus(((_1374_v2) ? (_1372_v0) : (_1372_v0)));
        let _rhs218 = (new BigNumber((((_1411_cf4) ? (_1415_v29) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), ((_1418_v0) => function (_1419_i2) {
          return _1418_v0;
        })(_1372_v0))))).length)).multipliedBy((_dafny.ZERO).minus(_1372_v0));
        _1414_v28 = _rhs216;
        r1 = _rhs217;
        r1 = _rhs218;
        let _1420_v30;
        _1420_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v2,new BigNumber((_dafny.Set.fromElements(_1411_cf4, _1374_v2, _module.__default.fm6(globalState))).length));
        _1373_v1 = (_1373_v1).update(new BigNumber((_1420_v30).length), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-311), _1372_v0)));
        let _1421_v32;
        let _nw188 = new _module.C4();
        _nw188.__ctor(_1372_v0, function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of (_dafny.MultiSet.fromElements(p0, p0)).Elements) {
            let _1422_v31 = _compr_37;
            if ((_dafny.MultiSet.fromElements(p0, p0)).contains(_1422_v31)) {
              _coll37.push([_1422_v31,_1372_v0]);
            }
          }
          return _coll37;
        }());
        _1421_v32 = _nw188;
        let _1423_v33;
        let _nw189 = new _module.C1();
        _nw189.__ctor();
        _1423_v33 = _nw189;
        let _nw190 = new _module.C1();
        _nw190.__ctor();
        _1423_v33 = _nw190;
      } else if (_source21.is_DC0) {
        let _1424___mcc_h4 = (_source21).cf0;
        let _1425_cf0 = _1424___mcc_h4;
        let _1426_v34;
        _1426_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(((_1374_v2) ? (_1374_v2) : (_1374_v2))),_1374_v2);
        _1426_v34 = (_1426_v34).update(_1374_v2, _1374_v2);
        (globalState).f9 = !(_1374_v2);
        _1425_cf0 = _1372_v0;
        let _1427_v35;
        _1427_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1379_v5,((_1374_v2) ? (!(_1374_v2)) : (_1374_v2)));
        _1427_v35 = (_1427_v35).update(_1379_v5, !((_module.__default.fm11(_1374_v2, globalState)).isEqualTo(_1372_v0)));
      } else {
        let _1428___mcc_h5 = (_source21).cf5;
        let _1429_cf5 = _1428___mcc_h5;
        let _1430_v36;
        _1430_v36 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p0));
        let _1431_v37;
        _1431_v37 = _dafny.Map.Empty.slice().updateUnsafe(((_1374_v2) ? (_1430_v36) : (_1430_v36)),_1372_v0);
        let _1432_v38;
        _1432_v38 = _dafny.Set.fromElements(!(_1374_v2));
        let _1433_v39;
        _1433_v39 = _dafny.Seq.of(_dafny.Set.fromElements(_1374_v2), _1432_v38, _dafny.Set.fromElements(false, _1374_v2));
        _1431_v37 = (_1431_v37).update(_1430_v36, new BigNumber(((_1433_v39)[_module.__default.safeIndex(_1372_v0, new BigNumber((_1433_v39).length))]).length));
        let _1434_v40;
        _1434_v40 = _module.D4.create_DC12(_1372_v0);
        let _source22 = _1434_v40;
        if (_source22.is_DC11) {
          r0 = (_1374_v2) || (_1374_v2);
          let _1435_v41;
          _1435_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v2,_1374_v2);
          let _1436_v42;
          _1436_v42 = _dafny.Set.fromElements(_1372_v0, _1372_v0, _1372_v0, (new BigNumber((_1379_v5).length)).plus(new BigNumber((_1435_v41).length)));
          _1436_v42 = _1436_v42;
          let _1437_v43;
          let _init26 = ((_1438_v42, _1439_v0) => function (_1440_i3) {
            return (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_1439_v0))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(_1438_v42, _1438_v42)));
          })(_1436_v42, _1372_v0);
          let _nw191 = Array((new BigNumber(29)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw191.length); _i0_26++) {
            _nw191[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1437_v43 = _nw191;
          let _1441_v44;
          _1441_v44 = _module.D8.create_DC23(_1374_v2, true, _1372_v0);
          let _index212 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_1437_v43).length));
          (_1437_v43)[_index212] = (_1441_v44).dtor_cf35;
          let _index213 = _module.__default.safeIndex(new BigNumber(569), new BigNumber((_1437_v43).length));
          (_1437_v43)[_index213] = _1374_v2;
          let _rhs219 = (_1437_v43)[_module.__default.safeIndex(new BigNumber(569), new BigNumber((_1437_v43).length))];
          let _rhs220 = _1379_v5;
          let _lhs151 = globalState;
          _lhs151.f9 = _rhs219;
          _1379_v5 = _rhs220;
        } else if (_source22.is_DC12) {
          let _1442___mcc_h6 = (_source22).cf18;
          let _1443_cf18 = _1442___mcc_h6;
          let _1444_v45;
          _1444_v45 = _dafny.Set.fromElements(_1443_cf18, new BigNumber((_1430_v36).cardinality()));
          let _1445_v46;
          _1445_v46 = _dafny.MultiSet.fromElements(_1444_v45);
          let _rhs221 = ((_1374_v2) ? (_1443_cf18) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(600)), ((_1446_p0) => function (_1447_i4) {
            return _1446_p0;
          })(p0))).length)));
          let _rhs222 = (_1445_v46).IsDisjointFrom(_1445_v46);
          let _lhs152 = globalState;
          _1443_cf18 = _rhs221;
          _lhs152.f13 = _rhs222;
          let _1448_v47;
          let _nw192 = new _module.C1();
          _nw192.__ctor();
          _1448_v47 = _nw192;
          let _1449_v48;
          _1449_v48 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1372_v0,_1443_cf18),_1444_v45);
          let _1450_v49;
          let _nw193 = Array((new BigNumber(19)).toNumber());
          _nw193[(_dafny.ZERO).toNumber()] = _1374_v2;
          _nw193[(_dafny.ONE).toNumber()] = !(_1374_v2);
          _nw193[(new BigNumber(2)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(3)).toNumber()] = (_1372_v0).isLessThanOrEqualTo(_1443_cf18);
          _nw193[(new BigNumber(4)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(5)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(6)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(7)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(8)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(9)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(10)).toNumber()] = (_1449_v48).contains(_1373_v1);
          _nw193[(new BigNumber(11)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(12)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(13)).toNumber()] = (_1372_v0).isLessThan(_1372_v0);
          _nw193[(new BigNumber(14)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(15)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(16)).toNumber()] = (_1374_v2) && (_1374_v2);
          _nw193[(new BigNumber(17)).toNumber()] = _1374_v2;
          _nw193[(new BigNumber(18)).toNumber()] = _1374_v2;
          _1450_v49 = _nw193;
          let _1451_v50;
          _1451_v50 = _dafny.MultiSet.fromElements(_1374_v2);
          let _1452_v51;
          _1452_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1443_cf18,_1448_v47);
          let _1453_v52;
          _1453_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v2,new BigNumber(815));
          let _rhs223 = (!(_1374_v2)) || ((new BigNumber(((_1451_v50).update(_1374_v2, _module.__default.abs(new BigNumber(-436)))).cardinality())).isEqualTo((_dafny.ZERO).minus(_1372_v0)));
          let _rhs224 = (((_1452_v51).contains(_1443_cf18)) ? ((_1452_v51).get(_1443_cf18)) : (_1448_v47));
          let _rhs225 = _1450_v49;
          let _rhs226 = (_1443_cf18).plus(new BigNumber((_1379_v5).length));
          let _rhs227 = (_this).fm0((((_1453_v52).contains(_1374_v2)) ? ((_1453_v52).get(_1374_v2)) : (_1443_cf18)), ((!(_1374_v2)) ? (p0) : (p0)), globalState);
          let _lhs153 = globalState;
          _lhs153.f13 = _rhs223;
          _1448_v47 = _rhs224;
          _1450_v49 = _rhs225;
          r1 = _rhs226;
          r1 = _rhs227;
          _1374_v2 = _1374_v2;
          _1379_v5 = _dafny.Seq.UnicodeFromString("uf");
        } else {
          let _1454___mcc_h7 = (_source22).cf17;
          let _1455_cf17 = _1454___mcc_h7;
          let _1456_v53;
          let _nw194 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1456_v53 = _nw194;
          let _1457_v54;
          _1457_v54 = _dafny.MultiSet.fromElements(_1374_v2);
          let _1458_v55;
          _1458_v55 = _module.D6.create_DC15(_1457_v54);
          let _1459_v56;
          _1459_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v2,((_1374_v2) ? (!(!(_1374_v2))) : (true)));
          let _rhs228 = _1456_v53;
          let _rhs229 = _1458_v55;
          let _rhs230 = _1372_v0;
          let _rhs231 = (((_1459_v56).contains(_1374_v2)) ? ((_1459_v56).get(_1374_v2)) : (_1374_v2));
          let _rhs232 = (_dafny.ZERO).minus(_module.__default.fm11((_1372_v0).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("lecq")).length)), globalState));
          let _lhs154 = globalState;
          _1456_v53 = _rhs228;
          _1458_v55 = _rhs229;
          r1 = _rhs230;
          _lhs154.f9 = _rhs231;
          _1372_v0 = _rhs232;
          let _index214 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1456_v53).length));
          (_1456_v53)[_index214] = _1374_v2;
          let _index215 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1456_v53).length));
          (_1456_v53)[_index215] = !(_module.__default.fm6(globalState)) || (_1374_v2);
          let _1460_v57;
          _1460_v57 = _dafny.Seq.of(_1374_v2, _1374_v2);
          let _1461_v58;
          _1461_v58 = _dafny.MultiSet.fromElements(p0);
          let _1462_v59;
          _1462_v59 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((new BigNumber((_1460_v57).length)).multipliedBy(new BigNumber((_1461_v58).cardinality()))), (_1372_v0).minus(_1372_v0), _1372_v0);
          let _index216 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1456_v53).length));
          let _index217 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1456_v53).length));
          let _rhs233 = _1374_v2;
          let _rhs234 = (_1372_v0).isLessThan(_1372_v0);
          let _rhs235 = (_1462_v59).Difference(_1462_v59);
          let _lhs155 = _1456_v53;
          let _lhs156 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1456_v53).length));
          let _lhs157 = _1456_v53;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1456_v53).length));
          _lhs155[_lhs156] = _rhs233;
          _lhs157[_lhs158] = _rhs234;
          _1462_v59 = _rhs235;
          let _1463_v60;
          _1463_v60 = _dafny.Map.Empty.slice().updateUnsafe(true,_1376_v4);
          let _1464_v61;
          _1464_v61 = _module.D2.create_DC7((_1456_v53)[_module.__default.safeIndex(new BigNumber(502), new BigNumber((_1456_v53).length))], _1463_v60, new BigNumber(229), _1374_v2);
          let _1465_v62;
          let _init27 = ((_1466_v0) => function (_1467_i5) {
            return (_1467_i5).plus(_1466_v0);
          })(_1372_v0);
          let _nw195 = Array((new BigNumber(4)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw195.length); _i0_27++) {
            _nw195[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1465_v62 = _nw195;
          let _1468_v63;
          _1468_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1465_v62,_1372_v0);
          let _pat_let_tv28 = _1375_v3;
          let _rhs236 = ((_1462_v59).Intersect(_1462_v59)).update(_module.__default.fm22(_1464_v61, _1376_v4, (_1456_v53)[_module.__default.safeIndex(new BigNumber(502), new BigNumber((_1456_v53).length))], globalState), _module.__default.abs((_dafny.ZERO).minus(_1372_v0)));
          let _rhs237 = _module.__default.fm22(_1464_v61, function (_pat_let31_0) {
            return function (_1469_dt__update__tmp_h2) {
              return function (_pat_let32_0) {
                return function (_1470_dt__update_hcf5_h1) {
                  return _module.D0.create_DC3(_1470_dt__update_hcf5_h1);
                }(_pat_let32_0);
              }(_pat_let_tv28);
            }(_pat_let31_0);
          }(_1376_v4), (_1460_v57)[_module.__default.safeIndex(_1372_v0, new BigNumber((_1460_v57).length))], globalState);
          let _rhs238 = !((_1468_v63).contains(_1465_v62));
          let _rhs239 = _module.__default.fm19(_1372_v0, !(_1374_v2) || (_1374_v2), globalState);
          let _lhs159 = globalState;
          let _lhs160 = globalState;
          _1462_v59 = _rhs236;
          _1372_v0 = _rhs237;
          _lhs159.f5 = _rhs238;
          _lhs160.f13 = _rhs239;
          let _1471_v64;
          _1471_v64 = _module.D8.create_DC24();
          _1471_v64 = _1471_v64;
        }
        let _1472_v65;
        let _nw196 = new _module.C6();
        _nw196.__ctor();
        _1472_v65 = _nw196;
        let _1473_v66;
        _1473_v66 = _dafny.Seq.of(_1472_v65);
        let _1474_v67;
        _1474_v67 = _module.D11.create_DC29(_dafny.MultiSet.FromArray(_1473_v66));
        let _1475_v68;
        let _nw197 = Array((new BigNumber(6)).toNumber());
        _nw197[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(_1372_v0, new BigNumber(940));
        _nw197[(_dafny.ONE).toNumber()] = _1372_v0;
        _nw197[(new BigNumber(2)).toNumber()] = (_this).fm0(_1372_v0, p0, globalState);
        _nw197[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(587), _1372_v0);
        _nw197[(new BigNumber(4)).toNumber()] = _1372_v0;
        _nw197[(new BigNumber(5)).toNumber()] = (_1372_v0).plus(new BigNumber(((_1474_v67).dtor_cf46).cardinality()));
        _1475_v68 = _nw197;
        _1475_v68 = _1475_v68;
        let _1476_v69;
        _1476_v69 = _dafny.Seq.of(_1372_v0);
        let _1477_v70;
        _1477_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1476_v69).length),_1379_v5);
        _1477_v70 = (_1477_v70).Merge(_1477_v70);
      }
      let _1478_v71;
      _1478_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v2,_1374_v2);
      let _1479_v72;
      _1479_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v2,_1478_v71);
      let _1480_v73;
      _1480_v73 = _dafny.Seq.of(_1372_v0, _1372_v0);
      let _1481_v74;
      _1481_v74 = _dafny.Seq.of(_1374_v2);
      let _1482_v75;
      let _init28 = ((_1483_v5) => function (_1484_i6) {
        return _1483_v5;
      })(_1379_v5);
      let _nw198 = Array((new BigNumber(25)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw198.length); _i0_28++) {
        _nw198[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _1482_v75 = _nw198;
      let _1485_v76;
      _1485_v76 = _module.D7.create_DC19(_module.D1.create_DC5(_1374_v2, _1379_v5, _1481_v74), new BigNumber((_1379_v5).length), _1374_v2, _1482_v75);
      let _1486_v77;
      _1486_v77 = _dafny.Seq.of(_1480_v73);
      let _1487_v78;
      _1487_v78 = _dafny.Seq.of((_1486_v77)[_module.__default.safeIndex(_1372_v0, new BigNumber((_1486_v77).length))], _1480_v73);
      let _1488_v79;
      let _nw199 = Array((new BigNumber(15)).toNumber());
      _nw199[(_dafny.ZERO).toNumber()] = _1372_v0;
      _nw199[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1372_v0);
      _nw199[(new BigNumber(2)).toNumber()] = _1372_v0;
      _nw199[(new BigNumber(3)).toNumber()] = _1372_v0;
      _nw199[(new BigNumber(4)).toNumber()] = new BigNumber((((((_1479_v72).contains(_1374_v2)) ? ((_1479_v72).get(_1374_v2)) : (_dafny.Map.Empty.slice().updateUnsafe(_1374_v2,_1374_v2)))).Merge(_1478_v71)).length);
      _nw199[(new BigNumber(5)).toNumber()] = _1372_v0;
      _nw199[(new BigNumber(6)).toNumber()] = _1372_v0;
      _nw199[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_1372_v0);
      _nw199[(new BigNumber(8)).toNumber()] = _1372_v0;
      _nw199[(new BigNumber(9)).toNumber()] = new BigNumber((_1480_v73).length);
      _nw199[(new BigNumber(10)).toNumber()] = ((!(_1374_v2)) ? ((_1485_v76).dtor_cf28) : (new BigNumber((_1486_v77).length)));
      _nw199[(new BigNumber(11)).toNumber()] = _1372_v0;
      _nw199[(new BigNumber(12)).toNumber()] = _1372_v0;
      _nw199[(new BigNumber(13)).toNumber()] = (_1372_v0).multipliedBy(_1372_v0);
      _nw199[(new BigNumber(14)).toNumber()] = _1372_v0;
      _1488_v79 = _nw199;
      let _index218 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_1488_v79).length));
      (_1488_v79)[_index218] = (new BigNumber(460)).plus(_1372_v0);
      let _index219 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_1488_v79).length));
      (_1488_v79)[_index219] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1372_v0));
      let _1489_i7;
      _1489_i7 = _dafny.ZERO;
      L9: {
        while (_1374_v2) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1489_i7)) {
              break L9;
            }
            _1489_i7 = (_1489_i7).plus(_dafny.ONE);
            let _1490_v80;
            _1490_v80 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1379_v5).length));
            let _1491_v81;
            let _nw200 = new _module.C4();
            _nw200.__ctor((_1488_v79)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_1488_v79).length))], _1490_v80);
            _1491_v81 = _nw200;
            let _1492_v82;
            _1492_v82 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1491_v81);
            _1492_v82 = ((_1492_v82).Merge(_1492_v82)).Merge((_1492_v82).Merge(_1492_v82));
            (globalState).f13 = _1374_v2;
            let _1493_v83;
            _1493_v83 = _dafny.Map.Empty.slice().updateUnsafe(!(_1374_v2),_1376_v4);
            let _1494_v86;
            _1494_v86 = _dafny.MultiSet.fromElements(p0, p0, p0);
            let _1495_v87;
            _1495_v87 = _module.D9.create_DC27(!(_1374_v2), new BigNumber((_1379_v5).length), _1374_v2);
            let _1496_v88;
            _1496_v88 = _module.D2.create_DC7(_1374_v2, _1493_v83, new BigNumber((function () {
  let _coll38 = new _dafny.Map();
  for (const _compr_38 of (function () {
    let _coll39 = new _dafny.Map();
    for (const _compr_39 of (_1494_v86).Elements) {
      let _1497_v85 = _compr_39;
      if ((_1494_v86).contains(_1497_v85)) {
        _coll39.push([_1497_v85,(_1491_v81).f14]);
      }
    }
    return _coll39;
  }()).Keys.Elements) {
    let _1498_v84 = _compr_38;
    if ((function () {
      let _coll40 = new _dafny.Map();
      for (const _compr_40 of (_1494_v86).Elements) {
        let _1497_v85 = _compr_40;
        if ((_1494_v86).contains(_1497_v85)) {
          _coll40.push([_1497_v85,(_1491_v81).f14]);
        }
      }
      return _coll40;
    }()).contains(_1498_v84)) {
      _coll38.push([_1498_v84,_dafny.Set.fromElements(_1374_v2)]);
    }
  }
  return _coll38;
}()).length), !((_1495_v87).dtor_cf42));
            r1 = _module.__default.fm22(_1496_v88, _1376_v4, _1374_v2, globalState);
            let _1499_v89;
            let _nw201 = Array((new BigNumber(25)).toNumber()).fill([]);
            _1499_v89 = _nw201;
            let _1500_v90;
            let _nw202 = new _module.C1();
            _nw202.__ctor();
            _1500_v90 = _nw202;
            let _1501_v91;
            let _nw203 = Array((new BigNumber(8)).toNumber());
            _nw203[(_dafny.ZERO).toNumber()] = _1500_v90;
            _nw203[(_dafny.ONE).toNumber()] = _1500_v90;
            _nw203[(new BigNumber(2)).toNumber()] = _1500_v90;
            _nw203[(new BigNumber(3)).toNumber()] = _1500_v90;
            _nw203[(new BigNumber(4)).toNumber()] = _1500_v90;
            _nw203[(new BigNumber(5)).toNumber()] = _1500_v90;
            _nw203[(new BigNumber(6)).toNumber()] = _1500_v90;
            _nw203[(new BigNumber(7)).toNumber()] = _1500_v90;
            _1501_v91 = _nw203;
            let _1502_v92;
            _1502_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1501_v91,(_1491_v81).f14);
            let _rhs240 = _1499_v89;
            let _rhs241 = (((_1502_v92).contains(_1501_v91)) ? ((_1502_v92).get(_1501_v91)) : ((_1372_v0).multipliedBy(new BigNumber(448))));
            _1499_v89 = _rhs240;
            _1372_v0 = _rhs241;
          }
        }
      }
      let _1503_v93;
      _1503_v93 = _dafny.MultiSet.fromElements(_1374_v2);
      let _1504_v94;
      _1504_v94 = _dafny.MultiSet.fromElements((_1488_v79)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_1488_v79).length))], new BigNumber((_1503_v93).cardinality()));
      r0 = (_1504_v94).IsSubsetOf((_dafny.MultiSet.fromElements(_1372_v0, (_1488_v79)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_1488_v79).length))])).update(_1372_v0, _module.__default.abs((_1488_v79)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_1488_v79).length))])));
      r1 = (_1488_v79)[_module.__default.safeIndex(new BigNumber(320), new BigNumber((_1488_v79).length))];
      return [r0, r1];
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _1505_v0;
      _1505_v0 = _dafny.MultiSet.fromElements(p0, p0, true);
      let _1506_v1;
      _1506_v1 = _module.D6.create_DC16(((p0) ? (new BigNumber((_1505_v0).cardinality())) : (new BigNumber(489))));
      let _source23 = _1506_v1;
      if (_source23.is_DC16) {
        let _1507___mcc_h0 = (_source23).cf24;
        let _1508_cf24 = _1507___mcc_h0;
        let _1509_v2;
        _1509_v2 = _module.D0.create_DC0(_1508_cf24);
        let _1510_v3;
        _1510_v3 = _module.D0.create_DC3(_1509_v2);
        let _1511_v4;
        _1511_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1510_v3);
        let _1512_v5;
        _1512_v5 = _module.D2.create_DC7(p0, _1511_v4, _1508_cf24, p0);
        let _1513_v6;
        _1513_v6 = _dafny.Seq.of(true);
        let _1514_v7;
        _1514_v7 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1515_v8;
        let _nw204 = Array((new BigNumber(22)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = p0;
        _nw204[(_dafny.ONE).toNumber()] = p0;
        _nw204[(new BigNumber(2)).toNumber()] = _module.__default.fm2(_1514_v7, _1510_v3, _1508_cf24, globalState);
        _nw204[(new BigNumber(3)).toNumber()] = p0;
        _nw204[(new BigNumber(4)).toNumber()] = p0;
        _nw204[(new BigNumber(5)).toNumber()] = p0;
        _nw204[(new BigNumber(6)).toNumber()] = p0;
        _nw204[(new BigNumber(7)).toNumber()] = p0;
        _nw204[(new BigNumber(8)).toNumber()] = p0;
        _nw204[(new BigNumber(9)).toNumber()] = !(p0);
        _nw204[(new BigNumber(10)).toNumber()] = p0;
        _nw204[(new BigNumber(11)).toNumber()] = p0;
        _nw204[(new BigNumber(12)).toNumber()] = true;
        _nw204[(new BigNumber(13)).toNumber()] = !(false);
        _nw204[(new BigNumber(14)).toNumber()] = p0;
        _nw204[(new BigNumber(15)).toNumber()] = p0;
        _nw204[(new BigNumber(16)).toNumber()] = p0;
        _nw204[(new BigNumber(17)).toNumber()] = p0;
        _nw204[(new BigNumber(18)).toNumber()] = !(p0);
        _nw204[(new BigNumber(19)).toNumber()] = p0;
        _nw204[(new BigNumber(20)).toNumber()] = p0;
        _nw204[(new BigNumber(21)).toNumber()] = p0;
        _1515_v8 = _nw204;
        let _1516_v9;
        _1516_v9 = _dafny.Seq.of(_1515_v8, _1515_v8);
        let _1517_v10;
        _1517_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-271)), ((_1518_v7) => function (_1519_i0) {
          return _1518_v7;
        })(_1514_v7)),p0);
        let _1520_v12;
        _1520_v12 = _dafny.Seq.of(new BigNumber(601), _1508_cf24);
        let _1521_v13;
        _1521_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1508_cf24,_1508_cf24);
        let _1522_v14;
        _1522_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1523_v15;
        _1523_v15 = _dafny.Set.fromElements(true);
        let _1524_v16;
        let _nw205 = Array((new BigNumber(24)).toNumber());
        _nw205[(_dafny.ZERO).toNumber()] = _module.__default.fm22(_1512_v5, _1510_v3, p0, globalState);
        _nw205[(_dafny.ONE).toNumber()] = (_1508_cf24).plus((_dafny.ZERO).minus(_1508_cf24));
        _nw205[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_1508_cf24, (_dafny.ZERO).minus(_1508_cf24));
        _nw205[(new BigNumber(3)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(4)).toNumber()] = _module.__default.fm22(_module.D2.create_DC7(p0, _1511_v4, _1508_cf24, p0), _1510_v3, p0, globalState);
        _nw205[(new BigNumber(5)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(6)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(7)).toNumber()] = _module.__default.fm11(p0, globalState);
        _nw205[(new BigNumber(8)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(9)).toNumber()] = new BigNumber((_1513_v6).length);
        _nw205[(new BigNumber(10)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(11)).toNumber()] = new BigNumber((_1516_v9).length);
        _nw205[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_1513_v6).length), new BigNumber((function () {
          let _coll41 = new _dafny.Set();
          for (const _compr_41 of (_1517_v10).Keys.Elements) {
            let _1525_v11 = _compr_41;
            if ((_1517_v10).contains(_1525_v11)) {
              _coll41.add(_1525_v11);
            }
          }
          return _coll41;
        }()).length)), _1520_v12)).length);
        _nw205[(new BigNumber(13)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(14)).toNumber()] = (new BigNumber((_1521_v13).length)).multipliedBy(new BigNumber((_1522_v14).length));
        _nw205[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((p1).length))).length);
        _nw205[(new BigNumber(16)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(17)).toNumber()] = new BigNumber((_1523_v15).length);
        _nw205[(new BigNumber(18)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(19)).toNumber()] = _1508_cf24;
        _nw205[(new BigNumber(20)).toNumber()] = (_1508_cf24).minus(_1508_cf24);
        _nw205[(new BigNumber(21)).toNumber()] = _module.__default.fm22(_1512_v5, _1510_v3, p0, globalState);
        _nw205[(new BigNumber(22)).toNumber()] = (new BigNumber((_1505_v0).cardinality())).multipliedBy(_1508_cf24);
        _nw205[(new BigNumber(23)).toNumber()] = _1508_cf24;
        _1524_v16 = _nw205;
        let _rhs242 = (_module.__default.safeDivisionInt(_1508_cf24, _1508_cf24)).multipliedBy(_1508_cf24);
        let _rhs243 = _1524_v16;
        _1508_cf24 = _rhs242;
        _1524_v16 = _rhs243;
        let _1526_v17;
        let _nw206 = new _module.C5();
        _nw206.__ctor();
        _1526_v17 = _nw206;
        let _1527_v18;
        _1527_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1526_v17,_1524_v16);
        _1527_v18 = (_1527_v18).update(_1526_v17, _1524_v16);
        let _1528_v19;
        _1528_v19 = _dafny.Seq.of(p1, p1, p1, p1);
        _1508_cf24 = _module.__default.safeModuloInt(new BigNumber(54), new BigNumber(((_1528_v19)[_module.__default.safeIndex((((_1505_v0).contains(p0)) ? ((_1505_v0).get(p0)) : (_1508_cf24)), new BigNumber((_1528_v19).length))]).length));
        let _1529_v20;
        _1529_v20 = _module.D8.create_DC24();
        _1529_v20 = _1529_v20;
      } else if (_source23.is_DC15) {
        let _1530___mcc_h1 = (_source23).cf23;
        let _1531_cf23 = _1530___mcc_h1;
        let _1532_v21;
        _1532_v21 = new BigNumber(-864);
        (globalState).f3 = (_1532_v21).isLessThanOrEqualTo(_1532_v21);
        r2 = _1532_v21;
        if (((p0) ? (true) : ((_1532_v21).isEqualTo(_1532_v21)))) {
          _1532_v21 = _1532_v21;
          (globalState).f3 = (_1532_v21).isLessThanOrEqualTo(_1532_v21);
          let _1533_v22;
          _1533_v22 = _dafny.MultiSet.fromElements(_1532_v21, _1532_v21, _1532_v21);
          let _1534_v23;
          _1534_v23 = _module.D12.create_DC31(_1533_v22);
          let _1535_v24;
          let _nw207 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _1535_v24 = _nw207;
          let _1536_v25;
          _1536_v25 = _dafny.MultiSet.fromElements(_1535_v24, _1535_v24);
          let _1537_v26;
          _1537_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1535_v24,_1535_v24);
          let _rhs244 = p0;
          let _rhs245 = (_1534_v23).dtor_cf48;
          let _rhs246 = p0;
          let _rhs247 = (_dafny.ZERO).minus((((_1536_v25).contains(((p0) ? ((((_1537_v26).contains(_1535_v24)) ? ((_1537_v26).get(_1535_v24)) : (_1535_v24))) : (_1535_v24)))) ? ((_1536_v25).get(((p0) ? ((((_1537_v26).contains(_1535_v24)) ? ((_1537_v26).get(_1535_v24)) : (_1535_v24))) : (_1535_v24)))) : (_1532_v21)));
          let _rhs248 = _module.__default.safeDivisionInt(_1532_v21, _1532_v21);
          let _lhs161 = globalState;
          let _lhs162 = globalState;
          _lhs161.f3 = _rhs244;
          _1533_v22 = _rhs245;
          _lhs162.f3 = _rhs246;
          _1532_v21 = _rhs247;
          _1532_v21 = _rhs248;
          r2 = _1532_v21;
          let _1538_v27;
          let _nw208 = Array((new BigNumber(7)).toNumber()).fill(false);
          _1538_v27 = _nw208;
          _1538_v27 = _1538_v27;
        } else {
          (globalState).f9 = (!(p0)) || ((p0) || (_module.__default.fm19(new BigNumber((p1).length), p0, globalState)));
          let _1539_v28;
          _1539_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1532_v21).multipliedBy((_dafny.ZERO).minus(_1532_v21)),_1532_v21);
          _1539_v28 = _1539_v28;
          (globalState).f13 = (new BigNumber(303)).isLessThan(_1532_v21);
          r2 = _1532_v21;
          let _1540_v29;
          let _nw209 = Array((new BigNumber(26)).toNumber()).fill(false);
          _1540_v29 = _nw209;
          let _1541_v30;
          let _nw210 = Array((new BigNumber(17)).toNumber());
          _nw210[(_dafny.ZERO).toNumber()] = _1540_v29;
          _nw210[(_dafny.ONE).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(2)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(3)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(4)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(5)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(6)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(7)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(8)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(9)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(10)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(11)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(12)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(13)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(14)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(15)).toNumber()] = _1540_v29;
          _nw210[(new BigNumber(16)).toNumber()] = _1540_v29;
          _1541_v30 = _nw210;
          let _index220 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_1541_v30).length));
          (_1541_v30)[_index220] = _1540_v29;
          let _index221 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_1541_v30).length));
          (_1541_v30)[_index221] = _1540_v29;
        }
        let _1542_v31;
        _1542_v31 = _dafny.Seq.of(p0);
        let _1543_v32;
        _1543_v32 = _dafny.Seq.of(p1);
        let _1544_v33;
        _1544_v33 = _module.D0.create_DC1(!(p0), _1543_v32, p1);
        let _1545_v34;
        _1545_v34 = _module.D0.create_DC3(_1544_v33);
        let _1546_v35;
        _1546_v35 = _dafny.Map.Empty.slice().updateUnsafe((_1542_v31)[_module.__default.safeIndex(_1532_v21, new BigNumber((_1542_v31).length))],_1545_v34);
        let _1547_v36;
        _1547_v36 = _module.D2.create_DC7(false, _1546_v35, _1532_v21, p0);
        let _1548_v37;
        _1548_v37 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1549_v38;
        _1549_v38 = _dafny.MultiSet.fromElements(_1548_v37);
        let _1550_v39;
        _1550_v39 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(!(p0)));
        let _1551_v40;
        _1551_v40 = _dafny.Seq.of(_1532_v21);
        let _1552_v41;
        _1552_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _1553_v42;
        _1553_v42 = _dafny.Seq.of(_1552_v41);
        let _1554_v43;
        let _nw211 = Array((new BigNumber(27)).toNumber());
        _nw211[(_dafny.ZERO).toNumber()] = _1532_v21;
        _nw211[(_dafny.ONE).toNumber()] = (_1532_v21).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(819)), function (_1555_i1) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        })).length));
        _nw211[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm22(_1547_v36, _1545_v34, p0, globalState));
        _nw211[(new BigNumber(3)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(4)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(5)).toNumber()] = (((_1549_v38).contains(_1548_v37)) ? ((_1549_v38).get(_1548_v37)) : (_1532_v21));
        _nw211[(new BigNumber(6)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(7)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(8)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(9)).toNumber()] = (((_1550_v39).contains(_1531_cf23)) ? ((_1550_v39).get(_1531_cf23)) : (_1532_v21));
        _nw211[(new BigNumber(10)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(11)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(12)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(13)).toNumber()] = (_1551_v40)[_module.__default.safeIndex(_1532_v21, new BigNumber((_1551_v40).length))];
        _nw211[(new BigNumber(14)).toNumber()] = ((p0) ? (_1532_v21) : (_1532_v21));
        _nw211[(new BigNumber(15)).toNumber()] = new BigNumber(542);
        _nw211[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("da"))).length);
        _nw211[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_1532_v21).plus(new BigNumber(471)));
        _nw211[(new BigNumber(18)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(19)).toNumber()] = ((p0) ? (_1532_v21) : (_1532_v21));
        _nw211[(new BigNumber(20)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1553_v42).length), (_dafny.ZERO).minus(_1532_v21));
        _nw211[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_1532_v21).plus(_1532_v21));
        _nw211[(new BigNumber(23)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(24)).toNumber()] = (_1532_v21).plus(_1532_v21);
        _nw211[(new BigNumber(25)).toNumber()] = _1532_v21;
        _nw211[(new BigNumber(26)).toNumber()] = _module.__default.safeDivisionInt(_1532_v21, (_dafny.ZERO).minus(_1532_v21));
        _1554_v43 = _nw211;
        let _1556_v44;
        let _nw212 = new _module.C2();
        _nw212.__ctor();
        _1556_v44 = _nw212;
        let _1557_v45;
        let _nw213 = Array((new BigNumber(29)).toNumber());
        _1557_v45 = _nw213;
        let _1558_v46;
        _1558_v46 = _module.D0.create_DC1(p0, _1543_v32, p1);
        let _1559_v47;
        _1559_v47 = _dafny.Seq.of(_1558_v46);
        let _1560_v48;
        let _nw214 = new _module.C0();
        _nw214.__ctor((_dafny.MultiSet.FromArray(_1559_v47)).update(_1558_v46, _module.__default.abs(_1532_v21)), p0);
        _1560_v48 = _nw214;
        let _index222 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_1557_v45).length));
        (_1557_v45)[_index222] = _1560_v48;
        let _1561_v49;
        _1561_v49 = _dafny.Set.fromElements((_1560_v48).f19);
        let _index223 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_1557_v45).length));
        let _rhs249 = _module.__default.fm19(((p0) ? (_1532_v21) : (new BigNumber((_1561_v49).length))), (_1560_v48).f19, globalState);
        let _rhs250 = (_module.D13.create_DC33(_1554_v43)).dtor_cf51;
        let _rhs251 = _1556_v44;
        let _rhs252 = _1560_v48;
        let _rhs253 = p0;
        let _lhs163 = globalState;
        let _lhs164 = _1557_v45;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_1557_v45).length));
        let _lhs166 = globalState;
        _lhs163.f3 = _rhs249;
        _1554_v43 = _rhs250;
        _1556_v44 = _rhs251;
        _lhs164[_lhs165] = _rhs252;
        _lhs166.f5 = _rhs253;
      } else {
        let _1562___mcc_h2 = (_source23).cf25;
        let _1563_cf25 = _1562___mcc_h2;
        let _1564_v50;
        _1564_v50 = new BigNumber(377);
        let _1565_v51;
        _1565_v51 = _dafny.Set.fromElements(p0, p0);
        let _1566_v52;
        _1566_v52 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(p0));
        let _1567_v53;
        _1567_v53 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1564_v50);
        let _1568_v54;
        _1568_v54 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p0,_1564_v50), _1567_v53, _1567_v53, _1567_v53);
        let _1569_v55;
        _1569_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1564_v50,p0);
        let _1570_v56;
        _1570_v56 = _dafny.Seq.of(_1564_v50, _1564_v50);
        let _1571_v57;
        let _nw215 = Array((new BigNumber(15)).toNumber());
        _nw215[(_dafny.ZERO).toNumber()] = (_1564_v50).minus(_1564_v50);
        _nw215[(_dafny.ONE).toNumber()] = _1564_v50;
        _nw215[(new BigNumber(2)).toNumber()] = (_1564_v50).multipliedBy(_1564_v50);
        _nw215[(new BigNumber(3)).toNumber()] = _1564_v50;
        _nw215[(new BigNumber(4)).toNumber()] = new BigNumber(((_1565_v51).Difference(_module.__default.fm33(p0, p0, !(p0), globalState))).length);
        _nw215[(new BigNumber(5)).toNumber()] = new BigNumber((_1566_v52).length);
        _nw215[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((p1).length));
        _nw215[(new BigNumber(7)).toNumber()] = _module.__default.fm11(p0, globalState);
        _nw215[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p1, p1)).length);
        _nw215[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(_1564_v50, _1564_v50);
        _nw215[(new BigNumber(10)).toNumber()] = (new BigNumber(124)).multipliedBy(_1564_v50);
        _nw215[(new BigNumber(11)).toNumber()] = new BigNumber(((_1568_v54).Union(_1568_v54)).length);
        _nw215[(new BigNumber(12)).toNumber()] = new BigNumber(((_1569_v55).Merge(_1569_v55)).length);
        _nw215[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.of(_1570_v56, _1570_v56, _1570_v56, _1570_v56)).length);
        _nw215[(new BigNumber(14)).toNumber()] = _1564_v50;
        _1571_v57 = _nw215;
        _1571_v57 = _1571_v57;
        r3 = p0;
        _1564_v50 = (_dafny.ZERO).minus(_1564_v50);
        let _1572_v58;
        _1572_v58 = _dafny.Seq.UnicodeFromString("vcvnthiyb");
        _1572_v58 = p1;
      }
      if (p0) {
        let _1573_v59;
        let _nw216 = Array((new BigNumber(3)).toNumber()).fill(false);
        _1573_v59 = _nw216;
        let _1574_v60;
        _1574_v60 = _dafny.Set.fromElements(_1573_v59, _1573_v59, _1573_v59, _1573_v59, _1573_v59);
        _1574_v60 = _1574_v60;
        let _index224 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
        (_1573_v59)[_index224] = false;
        let _1575_v61;
        _1575_v61 = new BigNumber(330);
        let _1576_v62;
        _1576_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-713)), ((_1577_v61) => function (_1578_i2) {
          return _1577_v61;
        })(_1575_v61))).length),_1575_v61);
        let _1579_v63;
        let _nw217 = Array((new BigNumber(24)).toNumber());
        _nw217[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_1575_v61);
        _nw217[(_dafny.ONE).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(2)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(3)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_1575_v61);
        _nw217[(new BigNumber(5)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(6)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(7)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(8)).toNumber()] = new BigNumber(440);
        _nw217[(new BigNumber(9)).toNumber()] = (((_1576_v62).contains(_1575_v61)) ? ((_1576_v62).get(_1575_v61)) : (_1575_v61));
        _nw217[(new BigNumber(10)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(11)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(12)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(13)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(14)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(15)).toNumber()] = new BigNumber(688);
        _nw217[(new BigNumber(16)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(17)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(18)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(19)).toNumber()] = _module.__default.fm11(p0, globalState);
        _nw217[(new BigNumber(20)).toNumber()] = (((_1505_v0).contains(p0)) ? ((_1505_v0).get(p0)) : (_1575_v61));
        _nw217[(new BigNumber(21)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(22)).toNumber()] = _1575_v61;
        _nw217[(new BigNumber(23)).toNumber()] = new BigNumber((_1574_v60).length);
        _1579_v63 = _nw217;
        let _1580_v64;
        _1580_v64 = _dafny.Seq.of(_1579_v63);
        let _1581_v65;
        _1581_v65 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1582_v66;
        _1582_v66 = _dafny.Seq.of(p0, p0, p0);
        let _1583_v67;
        let _nw218 = Array((new BigNumber(14)).toNumber());
        _nw218[(_dafny.ZERO).toNumber()] = (_1575_v61).plus(new BigNumber(233));
        _nw218[(_dafny.ONE).toNumber()] = (_1575_v61).multipliedBy(new BigNumber((_1580_v64).length));
        _nw218[(new BigNumber(2)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(3)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(4)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(5)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(6)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(7)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(8)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(9)).toNumber()] = _1575_v61;
        _nw218[(new BigNumber(10)).toNumber()] = new BigNumber((_1581_v65).length);
        _nw218[(new BigNumber(11)).toNumber()] = (_1575_v61).plus(_1575_v61);
        _nw218[(new BigNumber(12)).toNumber()] = (_1575_v61).multipliedBy(new BigNumber((_1582_v66).length));
        _nw218[(new BigNumber(13)).toNumber()] = _1575_v61;
        _1583_v67 = _nw218;
        let _index225 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length));
        (_1583_v67)[_index225] = _1575_v61;
        let _1584_v68;
        _1584_v68 = _dafny.Seq.of(_1575_v61, _1575_v61, _1575_v61, _1575_v61);
        let _1585_v69;
        let _nw219 = new _module.C5();
        _nw219.__ctor();
        _1585_v69 = _nw219;
        let _1586_v70;
        _1586_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1585_v69);
        let _1587_v71;
        _1587_v71 = _dafny.Seq.of((_1586_v70).update(false, _1585_v69), _1586_v70, _1586_v70, _1586_v70);
        let _1588_v72;
        _1588_v72 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(679)), function (_1589_i3) {
          return (_module.D4.create_DC12(new BigNumber(725))).dtor_cf18;
        }), _dafny.Seq.update(_1584_v68, _module.__default.safeIndex(_1575_v61, new BigNumber((_1584_v68).length)), new BigNumber(173)), _1584_v68, (((_1582_v66)[_module.__default.safeIndex(new BigNumber((_1587_v71).length), new BigNumber((_1582_v66).length))]) ? (_1584_v68) : (_1584_v68)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(178)), ((_1590_v62) => function (_1591_i4) {
          return new BigNumber((_1590_v62).length);
        })(_1576_v62)), _module.__default.fm30(globalState)));
        let _1592_v73;
        let _nw220 = new _module.C2();
        _nw220.__ctor();
        _1592_v73 = _nw220;
        let _1593_v74;
        _1593_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1592_v73,_1584_v68);
        let _1594_v75;
        _1594_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1583_v67,new BigNumber((_1593_v74).length));
        let _1595_v76;
        _1595_v76 = _dafny.MultiSet.fromElements(_1575_v61);
        let _index226 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
        let _index227 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length));
        let _rhs254 = _1575_v61;
        let _rhs255 = p0;
        let _rhs256 = (((_1594_v75).contains(_1583_v67)) ? ((_1594_v75).get(_1583_v67)) : ((new BigNumber((_1595_v76).cardinality())).multipliedBy(new BigNumber(623))));
        let _rhs257 = new BigNumber((_1581_v65).length);
        let _rhs258 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), ((_1596_v68) => function (_1597_i5) {
          return _dafny.Seq.Concat(_1596_v68, _1596_v68);
        })(_1584_v68));
        let _lhs167 = _1573_v59;
        let _lhs168 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
        let _lhs169 = _1583_v67;
        let _lhs170 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length));
        r2 = _rhs254;
        _lhs167[_lhs168] = _rhs255;
        r2 = _rhs256;
        _lhs169[_lhs170] = _rhs257;
        _1588_v72 = _rhs258;
        let _index228 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
        (_1573_v59)[_index228] = !(p0);
        if (!(false)) {
          let _index229 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
          (_1573_v59)[_index229] = (_1582_v66)[_module.__default.safeIndex(_1575_v61, new BigNumber((_1582_v66).length))];
          _1575_v61 = new BigNumber(653);
          r1 = p0;
          let _1598_v77;
          let _nw221 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
          _1598_v77 = _nw221;
          let _1599_v78;
          _1599_v78 = _dafny.Seq.of(_1598_v77, _1598_v77, _1598_v77);
          let _1600_v79;
          let _nw222 = Array((new BigNumber(27)).toNumber());
          _nw222[(_dafny.ZERO).toNumber()] = _1598_v77;
          _nw222[(_dafny.ONE).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(2)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(3)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(4)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(5)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(6)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(7)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(8)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(9)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(10)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(11)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(12)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(13)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(14)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(15)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(16)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(17)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(18)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(19)).toNumber()] = (_1599_v78)[_module.__default.safeIndex((_1583_v67)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length))], new BigNumber((_1599_v78).length))];
          _nw222[(new BigNumber(20)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(21)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(22)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(23)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(24)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(25)).toNumber()] = _1598_v77;
          _nw222[(new BigNumber(26)).toNumber()] = _1598_v77;
          _1600_v79 = _nw222;
          let _index230 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_1600_v79).length));
          (_1600_v79)[_index230] = _1598_v77;
          let _index231 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_1600_v79).length));
          let _index232 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
          let _rhs259 = _1598_v77;
          let _rhs260 = ((!(p0)) ? (((_1505_v0).update(p0, _module.__default.abs(_1575_v61))).IsSubsetOf(_1505_v0)) : (p0));
          let _lhs171 = _1600_v79;
          let _lhs172 = _module.__default.safeIndex(new BigNumber(493), new BigNumber((_1600_v79).length));
          let _lhs173 = _1573_v59;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
          _lhs171[_lhs172] = _rhs259;
          _lhs173[_lhs174] = _rhs260;
          r2 = _1575_v61;
        } else {
          let _1601_v80;
          _1601_v80 = _dafny.Set.fromElements(p0);
          r0 = ((_1601_v80).Intersect(_module.__default.fm33(p0, (_1573_v59)[_module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length))], (_1573_v59)[_module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length))], globalState))).Union(_1601_v80);
          let _1602_v81;
          _1602_v81 = _module.D0.create_DC0(new BigNumber((_1582_v66).length));
          let _1603_v82;
          let _nw223 = Array((new BigNumber(4)).toNumber());
          _nw223[(_dafny.ZERO).toNumber()] = _1602_v81;
          _nw223[(_dafny.ONE).toNumber()] = _1602_v81;
          _nw223[(new BigNumber(2)).toNumber()] = _module.D0.create_DC0(new BigNumber((p1).length));
          _nw223[(new BigNumber(3)).toNumber()] = _1602_v81;
          _1603_v82 = _nw223;
          let _1604_v83;
          _1604_v83 = _module.D7.create_DC20(p1, _module.__default.fm2(new _dafny.CodePoint('u'.codePointAt(0)), _module.__default.fm28(globalState), (_1583_v67)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length))], globalState));
          let _1605_v84;
          _1605_v84 = _dafny.Seq.of(p1);
          let _1606_v85;
          _1606_v85 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("bjx"),_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(416)), ((_1607_p1) => function (_1608_i6) {
            return _1607_p1;
          })(p1)), _1605_v84));
          let _1609_v86;
          _1609_v86 = _dafny.Set.fromElements(_1575_v61, _1575_v61);
          let _1610_v87;
          _1610_v87 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1609_v86);
          let _1611_v88;
          _1611_v88 = _module.D4.create_DC10((((_1610_v87).contains(p1)) ? ((_1610_v87).get(p1)) : (_1609_v86)));
          let _index233 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length));
          let _rhs261 = new BigNumber(((((_1606_v85).contains(p1)) ? ((_1606_v85).get(p1)) : (_1605_v84))).length);
          let _rhs262 = _1603_v82;
          let _rhs263 = (_1575_v61).multipliedBy((_dafny.ZERO).minus(_1575_v61));
          let _rhs264 = _module.__default.fm40(_1575_v61, _1611_v88, globalState);
          let _rhs265 = (_1575_v61).plus(_1575_v61);
          let _lhs175 = _1583_v67;
          let _lhs176 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length));
          _lhs175[_lhs176] = _rhs261;
          _1603_v82 = _rhs262;
          _1575_v61 = _rhs263;
          _1604_v83 = _rhs264;
          r2 = _rhs265;
          let _1612_v89;
          let _nw224 = new _module.C3();
          _nw224.__ctor((_1575_v61).isLessThanOrEqualTo(_1575_v61), (_1583_v67)[_module.__default.safeIndex(new BigNumber(756), new BigNumber((_1583_v67).length))]);
          _1612_v89 = _nw224;
          let _1613_v90;
          _1613_v90 = _dafny.Set.fromElements(_1604_v83);
          let _1614_v91;
          _1614_v91 = _dafny.Seq.of(_module.D7.create_DC20(_dafny.Seq.Create(_module.__default.abs(new BigNumber(965)), function (_1615_i7) {
  return new _dafny.CodePoint('g'.codePointAt(0));
}), p0));
          let _index234 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
          (_1573_v59)[_index234] = (((_1573_v59)[_module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length))]) ? ((_1613_v90).IsDisjointFrom(function () {
            let _coll42 = new _dafny.Set();
            for (const _compr_42 of (_1614_v91).Elements) {
              let _1616_v92 = _compr_42;
              if (_dafny.Seq.contains(_1614_v91, _1616_v92)) {
                _coll42.add(_1616_v92);
              }
            }
            return _coll42;
          }())) : (p0));
          let _1617_v93;
          let _nw225 = Array((new BigNumber(11)).toNumber()).fill([]);
          _1617_v93 = _nw225;
          _1617_v93 = _1617_v93;
        }
        let _1618_v94;
        _1618_v94 = _module.D13.create_DC33(_1579_v63);
        let _1619_v95;
        let _nw226 = Array((new BigNumber(11)).toNumber());
        _nw226[(_dafny.ZERO).toNumber()] = _1583_v67;
        _nw226[(_dafny.ONE).toNumber()] = _1583_v67;
        _nw226[(new BigNumber(2)).toNumber()] = _1583_v67;
        _nw226[(new BigNumber(3)).toNumber()] = ((p0) ? (_1583_v67) : (_1579_v63));
        _nw226[(new BigNumber(4)).toNumber()] = (_1580_v64)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-168)), new BigNumber((_1580_v64).length))];
        _nw226[(new BigNumber(5)).toNumber()] = (_1618_v94).dtor_cf51;
        _nw226[(new BigNumber(6)).toNumber()] = _1579_v63;
        _nw226[(new BigNumber(7)).toNumber()] = _1579_v63;
        _nw226[(new BigNumber(8)).toNumber()] = _1579_v63;
        _nw226[(new BigNumber(9)).toNumber()] = _1579_v63;
        _nw226[(new BigNumber(10)).toNumber()] = _1579_v63;
        _1619_v95 = _nw226;
        let _index235 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1619_v95).length));
        (_1619_v95)[_index235] = _1583_v67;
        let _1620_v96;
        _1620_v96 = _module.D5.create_DC13(_1573_v59);
        let _index236 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1619_v95).length));
        let _index237 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
        let _rhs266 = !(p0);
        let _rhs267 = _1583_v67;
        let _rhs268 = p0;
        let _rhs269 = ((p0) ? (p0) : (p0));
        let _rhs270 = (_1620_v96).dtor_cf19;
        let _lhs177 = globalState;
        let _lhs178 = _1619_v95;
        let _lhs179 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_1619_v95).length));
        let _lhs180 = globalState;
        let _lhs181 = _1573_v59;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1573_v59).length));
        _lhs177.f13 = _rhs266;
        _lhs178[_lhs179] = _rhs267;
        _lhs180.f9 = _rhs268;
        _lhs181[_lhs182] = _rhs269;
        _1573_v59 = _rhs270;
      } else {
        let _1621_v97;
        _1621_v97 = _dafny.Set.fromElements(!(false), p0);
        r0 = (_1621_v97).Intersect(_dafny.Set.fromElements(p0, p0, p0));
        let _1622_v98;
        let _nw227 = Array((new BigNumber(18)).toNumber()).fill(false);
        _1622_v98 = _nw227;
        let _index238 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1622_v98).length));
        (_1622_v98)[_index238] = true;
        let _index239 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1622_v98).length));
        (_1622_v98)[_index239] = p0;
        let _1623_v99;
        _1623_v99 = new BigNumber(564);
        r3 = (new BigNumber(836)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_1623_v99, _1623_v99));
        (globalState).f9 = !(p0);
        let _1624_v100;
        _1624_v100 = _dafny.Seq.UnicodeFromString("bspwfx");
        _1624_v100 = _dafny.Seq.Concat(p1, _dafny.Seq.Concat(p1, _1624_v100));
      }
      (globalState).f9 = _module.__default.fm19(new BigNumber((_dafny.Seq.UnicodeFromString("gb")).length), !(!(p0)), globalState);
      if (p0) {
        let _1625_v101;
        _1625_v101 = new _dafny.CodePoint('s'.codePointAt(0));
        _1625_v101 = _1625_v101;
        let _1626_v102;
        _1626_v102 = _dafny.Seq.UnicodeFromString("xmcve");
        _1626_v102 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-191)), ((_1627_v101) => function (_1628_i8) {
          return _1627_v101;
        })(_1625_v101));
        let _1629_v103;
        _1629_v103 = _dafny.Seq.of(p0);
        let _1630_v104;
        _1630_v104 = _dafny.Set.fromElements(_1629_v103, _1629_v103);
        r2 = new BigNumber(((_1630_v104).Union(_1630_v104)).length);
        let _1631_v105;
        _1631_v105 = new BigNumber(252);
        r2 = (_dafny.ZERO).minus(_1631_v105);
        let _1632_v106;
        _1632_v106 = _module.D9.create_DC26(p0, _1625_v101, p0);
        let _1633_v107;
        _1633_v107 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1634_v108;
        _1634_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1625_v101,_1633_v107);
        let _pat_let_tv29 = p0;
        let _pat_let_tv30 = p0;
        let _pat_let_tv31 = _1634_v108;
        let _pat_let_tv32 = p0;
        let _pat_let_tv33 = _1631_v105;
        let _pat_let_tv34 = globalState;
        let _1635_v109;
        let _nw228 = Array((new BigNumber(9)).toNumber());
        _nw228[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _nw228[(_dafny.ONE).toNumber()] = _1625_v101;
        _nw228[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw228[(new BigNumber(3)).toNumber()] = _1625_v101;
        _nw228[(new BigNumber(4)).toNumber()] = _1625_v101;
        _nw228[(new BigNumber(5)).toNumber()] = (function (_pat_let33_0) {
          return function (_1636_dt__update__tmp_h0) {
            return function (_pat_let34_0) {
              return function (_1637_dt__update_hcf41_h0) {
                return function (_pat_let35_0) {
                  return function (_1638_dt__update_hcf40_h0) {
                    return _module.D9.create_DC26((_1636_dt__update__tmp_h0).dtor_cf39, _1638_dt__update_hcf40_h0, _1637_dt__update_hcf41_h0);
                  }(_pat_let35_0);
                }(_module.__default.fm23(_pat_let_tv30, _pat_let_tv31, _pat_let_tv32, _pat_let_tv33, _pat_let_tv34));
              }(_pat_let34_0);
            }(_pat_let_tv29);
          }(_pat_let33_0);
        }(_1632_v106)).dtor_cf40;
        _nw228[(new BigNumber(6)).toNumber()] = _1625_v101;
        _nw228[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
        _nw228[(new BigNumber(8)).toNumber()] = _1625_v101;
        _1635_v109 = _nw228;
        let _index240 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1635_v109).length));
        (_1635_v109)[_index240] = new _dafny.CodePoint('w'.codePointAt(0));
        let _1639_v110;
        _1639_v110 = _dafny.Set.fromElements(true);
        let _1640_v111;
        _1640_v111 = _module.D7.create_DC20(_1626_v102, p0);
        let _index241 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1635_v109).length));
        let _rhs271 = (_module.__default.fm41(p0, true, globalState)).dtor_cf39;
        let _rhs272 = _module.__default.fm23((_1639_v110).IsProperSubsetOf(_dafny.Set.fromElements(p0, p0)), _1634_v108, (_1640_v111).dtor_cf32, _1631_v105, globalState);
        let _lhs183 = globalState;
        let _lhs184 = _1635_v109;
        let _lhs185 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_1635_v109).length));
        _lhs183.f5 = _rhs271;
        _lhs184[_lhs185] = _rhs272;
      } else {
        let _1641_v112;
        _1641_v112 = new BigNumber(579);
        r2 = _1641_v112;
        if (p0) {
          let _1642_v113;
          let _nw229 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _1642_v113 = _nw229;
          let _1643_v114;
          _1643_v114 = _dafny.Map.Empty.slice().updateUnsafe(_1642_v113,p0);
          _1643_v114 = _1643_v114;
          let _1644_v115;
          let _nw230 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1644_v115 = _nw230;
          let _1645_v116;
          _1645_v116 = _dafny.Seq.of(p0);
          let _index242 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1644_v115).length));
          (_1644_v115)[_index242] = !_dafny.Seq.contains(_1645_v116, p0);
          let _index243 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1644_v115).length));
          let _rhs273 = p0;
          let _rhs274 = _1641_v112;
          let _rhs275 = _module.__default.safeDivisionInt(new BigNumber((p1).length), _1641_v112);
          let _lhs186 = _1644_v115;
          let _lhs187 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1644_v115).length));
          _lhs186[_lhs187] = _rhs273;
          r2 = _rhs274;
          r2 = _rhs275;
          let _1646_v117;
          let _nw231 = Array((new BigNumber(18)).toNumber());
          _1646_v117 = _nw231;
          let _1647_v118;
          let _nw232 = new _module.C2();
          _nw232.__ctor();
          _1647_v118 = _nw232;
          let _index244 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1646_v117).length));
          (_1646_v117)[_index244] = _1647_v118;
          let _index245 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_1646_v117).length));
          (_1646_v117)[_index245] = _1647_v118;
          let _1648_v119;
          _1648_v119 = _dafny.Map.Empty.slice().updateUnsafe(!((_1644_v115)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_1644_v115).length))]),_1642_v113);
          let _1649_v120;
          _1649_v120 = _dafny.Set.fromElements((_1641_v112).minus(_1641_v112), new BigNumber(38), new BigNumber((_1648_v119).length));
          _1649_v120 = _1649_v120;
          (globalState).f9 = (_1644_v115)[_module.__default.safeIndex(new BigNumber(694), new BigNumber((_1644_v115).length))];
        } else {
          let _1650_v121;
          _1650_v121 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1651_v122;
          _1651_v122 = _dafny.Map.Empty.slice().updateUnsafe(_1650_v121,_1641_v112);
          let _1652_v123;
          let _nw233 = new _module.C4();
          _nw233.__ctor(_1641_v112, (_1651_v122).Merge(_1651_v122));
          _1652_v123 = _nw233;
          let _1653_v124;
          _1653_v124 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_1641_v112);
          let _1654_v125;
          _1654_v125 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), _1650_v121);
          let _1655_v126;
          let _nw234 = Array((new BigNumber(17)).toNumber());
          _nw234[(_dafny.ZERO).toNumber()] = (_1652_v123).f14;
          _nw234[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update((_this).fm1(_1653_v124, globalState), _module.__default.safeIndex(_1641_v112, new BigNumber(((_this).fm1(_1653_v124, globalState)).length)), _1650_v121)).length);
          _nw234[(new BigNumber(2)).toNumber()] = (_1652_v123).f14;
          _nw234[(new BigNumber(3)).toNumber()] = (_1652_v123).f14;
          _nw234[(new BigNumber(4)).toNumber()] = new BigNumber(((_1654_v125).update(_1650_v121, _module.__default.abs((_1652_v123).f14))).cardinality());
          _nw234[(new BigNumber(5)).toNumber()] = (_1652_v123).f14;
          _nw234[(new BigNumber(6)).toNumber()] = _1641_v112;
          _nw234[(new BigNumber(7)).toNumber()] = _1641_v112;
          _nw234[(new BigNumber(8)).toNumber()] = (_1652_v123).f14;
          _nw234[(new BigNumber(9)).toNumber()] = (_1652_v123).f14;
          _nw234[(new BigNumber(10)).toNumber()] = new BigNumber(-354);
          _nw234[(new BigNumber(11)).toNumber()] = (_1652_v123).f14;
          _nw234[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(148)), ((_1656_v123) => function (_1657_i9) {
            return (_1656_v123).f14;
          })(_1652_v123))).length);
          _nw234[(new BigNumber(13)).toNumber()] = _1641_v112;
          _nw234[(new BigNumber(14)).toNumber()] = new BigNumber((p1).length);
          _nw234[(new BigNumber(15)).toNumber()] = _1641_v112;
          _nw234[(new BigNumber(16)).toNumber()] = (_1652_v123).f14;
          _1655_v126 = _nw234;
          let _1658_v127;
          _1658_v127 = _dafny.Seq.of(_1655_v126);
          let _1659_v128;
          _1659_v128 = _dafny.Set.fromElements(!(p0));
          let _1660_v129;
          _1660_v129 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p0, true),(_1659_v128).IsDisjointFrom(_1659_v128));
          let _rhs276 = _module.__default.safeModuloInt((_1652_v123).f14, (_1652_v123).f14);
          let _rhs277 = (((_1660_v129).contains(_1505_v0)) ? ((_1660_v129).get(_1505_v0)) : (p0));
          let _rhs278 = _1658_v127;
          let _lhs188 = globalState;
          r2 = _rhs276;
          _lhs188.f9 = _rhs277;
          _1658_v127 = _rhs278;
          let _1661_v130;
          _1661_v130 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _1661_v130 = _1661_v130;
          let _1662_v131;
          _1662_v131 = _dafny.Seq.of(p0, p0);
          let _1663_v132;
          let _nw235 = new _module.C0();
          _nw235.__ctor(_module.__default.fm42(p0, new BigNumber((_1662_v131).length), (_1652_v123).f15, _dafny.Seq.of(p0), globalState), (!(p0)) && (true));
          _1663_v132 = _nw235;
          _1663_v132 = _1663_v132;
          r1 = p0;
        }
        let _1664_v133;
        let _nw236 = Array((new BigNumber(22)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1664_v133 = _nw236;
        let _1665_v134;
        _1665_v134 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1666_v135;
        _1666_v135 = _dafny.Seq.of(_1665_v134);
        let _1667_v136;
        _1667_v136 = _dafny.MultiSet.fromElements(_1665_v134, (_1666_v135)[_module.__default.safeIndex(_1641_v112, new BigNumber((_1666_v135).length))]);
        let _index246 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_1664_v133).length));
        (_1664_v133)[_index246] = _1667_v136;
        let _index247 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_1664_v133).length));
        (_1664_v133)[_index247] = (_1667_v136).Union(_dafny.MultiSet.fromElements(_1665_v134, _1665_v134));
        let _1668_v137;
        let _nw237 = new _module.C1();
        _nw237.__ctor();
        _1668_v137 = _nw237;
        let _1669_v138;
        let _nw238 = Array((new BigNumber(9)).toNumber()).fill([]);
        _1669_v138 = _nw238;
        _1669_v138 = _1669_v138;
      }
      let _1670_v139;
      let _nw239 = Array((new BigNumber(10)).toNumber()).fill(false);
      _1670_v139 = _nw239;
      let _index248 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length));
      (_1670_v139)[_index248] = !(p0);
      let _1671_v140;
      _1671_v140 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1672_v141;
      _1672_v141 = new BigNumber(18);
      let _1673_v142;
      _1673_v142 = _dafny.Seq.of(p0, p0, p0, false, true);
      let _1674_v143;
      _1674_v143 = _module.D0.create_DC2(!((_1673_v142)[_module.__default.safeIndex(_1672_v141, new BigNumber((_1673_v142).length))]));
      let _1675_v144;
      _1675_v144 = _module.D0.create_DC3(_1674_v143);
      let _index249 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length));
      (_1670_v139)[_index249] = (((_1671_v140).contains(((p0) ? (p0) : (p0)))) ? ((_1671_v140).get(((p0) ? (p0) : (p0)))) : (_module.__default.fm2((p1)[_module.__default.safeIndex(_1672_v141, new BigNumber((p1).length))], _1675_v144, _1672_v141, globalState)));
      if ((_module.__default.safeDivisionInt(_1672_v141, new BigNumber((_dafny.Seq.update(_1673_v142, _module.__default.safeIndex(_1672_v141, new BigNumber((_1673_v142).length)), p0)).length))).isLessThanOrEqualTo(new BigNumber(210))) {
        let _1676_v145;
        _1676_v145 = _dafny.Seq.of(_1670_v139, _1670_v139);
        let _1677_v146;
        _1677_v146 = _dafny.Map.Empty.slice().updateUnsafe((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))],_1670_v139);
        let _1678_v147;
        let _nw240 = Array((new BigNumber(26)).toNumber());
        _nw240[(_dafny.ZERO).toNumber()] = _1670_v139;
        _nw240[(_dafny.ONE).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(2)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(3)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(4)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(5)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(6)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(7)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(8)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(9)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(10)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(11)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(12)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(13)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(14)).toNumber()] = (_1676_v145)[_module.__default.safeIndex(_1672_v141, new BigNumber((_1676_v145).length))];
        _nw240[(new BigNumber(15)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(16)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(17)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(18)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(19)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(20)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(21)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(22)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(23)).toNumber()] = _1670_v139;
        _nw240[(new BigNumber(24)).toNumber()] = (((_1677_v146).contains((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))])) ? ((_1677_v146).get((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))])) : (_1670_v139));
        _nw240[(new BigNumber(25)).toNumber()] = _1670_v139;
        _1678_v147 = _nw240;
        let _index250 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length));
        (_1678_v147)[_index250] = _1670_v139;
        let _index251 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length));
        (_1678_v147)[_index251] = (_module.D5.create_DC13(_1670_v139)).dtor_cf19;
        let _1679_v148;
        _1679_v148 = _dafny.Seq.of(_1672_v141);
        let _1680_v149;
        _1680_v149 = _dafny.Seq.of(_1676_v145);
        let _1681_v150;
        let _nw241 = Array((new BigNumber(24)).toNumber());
        _nw241[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1676_v145, _module.__default.safeIndex(_1672_v141, new BigNumber((_1676_v145).length)), (_1678_v147)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length))]);
        _nw241[(_dafny.ONE).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(2)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(3)).toNumber()] = _dafny.Seq.of((_1678_v147)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length))]);
        _nw241[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1676_v145, _1676_v145);
        _nw241[(new BigNumber(5)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(6)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(7)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_1676_v145, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1679_v148).length)), new BigNumber((_1676_v145).length)), _1670_v139);
        _nw241[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1676_v145, _1676_v145);
        _nw241[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1676_v145, _1676_v145);
        _nw241[(new BigNumber(11)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_1678_v147)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length))]), _dafny.Seq.of((_1678_v147)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length))], _1670_v139));
        _nw241[(new BigNumber(13)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(14)).toNumber()] = _dafny.Seq.of((_1678_v147)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length))], (_1678_v147)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length))], _1670_v139);
        _nw241[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1676_v145, _1676_v145);
        _nw241[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat((_1680_v149)[_module.__default.safeIndex(_1672_v141, new BigNumber((_1680_v149).length))], _1676_v145);
        _nw241[(new BigNumber(17)).toNumber()] = _dafny.Seq.of((_1678_v147)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1678_v147).length))]);
        _nw241[(new BigNumber(18)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_1670_v139, _1670_v139);
        _nw241[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_1670_v139, _1670_v139), _module.__default.safeIndex(_1672_v141, new BigNumber((_dafny.Seq.of(_1670_v139, _1670_v139)).length)), _1670_v139);
        _nw241[(new BigNumber(21)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(22)).toNumber()] = _1676_v145;
        _nw241[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(_1670_v139, _1670_v139);
        _1681_v150 = _nw241;
        let _index252 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1681_v150).length));
        (_1681_v150)[_index252] = _1676_v145;
        let _1682_v151;
        _1682_v151 = _dafny.Map.Empty.slice().updateUnsafe((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))],_dafny.Seq.Create(_module.__default.abs(new BigNumber(791)), function (_1683_i11) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }));
        let _1684_v152;
        _1684_v152 = _dafny.Seq.of(p1, p1);
        let _1685_v153;
        _1685_v153 = new _dafny.CodePoint('e'.codePointAt(0));
        let _1686_v154;
        _1686_v154 = _module.D0.create_DC1(false, _dafny.Seq.update(_1684_v152, _module.__default.safeIndex(_1672_v141, new BigNumber((_1684_v152).length)), _dafny.Seq.update(p1, _module.__default.safeIndex(_1672_v141, new BigNumber((p1).length)), new _dafny.CodePoint('f'.codePointAt(0)))), _dafny.Seq.update(p1, _module.__default.safeIndex(_1672_v141, new BigNumber((p1).length)), _1685_v153));
        let _pat_let_tv35 = _1670_v139;
        let _pat_let_tv36 = _1670_v139;
        let _pat_let_tv37 = _1684_v152;
        let _1687_v155;
        let _nw242 = Array((new BigNumber(26)).toNumber());
        _nw242[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), ((_1688_p1, _1689_v141) => function (_1690_i10) {
          return (_1688_p1)[_module.__default.safeIndex((_dafny.ZERO).minus(_1689_v141), new BigNumber((_1688_p1).length))];
        })(p1, _1672_v141));
        _nw242[(_dafny.ONE).toNumber()] = (((_1682_v151).contains(p0)) ? ((_1682_v151).get(p0)) : (_dafny.Seq.UnicodeFromString("sfabqetxt")));
        _nw242[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(p1, p1);
        _nw242[(new BigNumber(3)).toNumber()] = p1;
        _nw242[(new BigNumber(4)).toNumber()] = p1;
        _nw242[(new BigNumber(5)).toNumber()] = p1;
        _nw242[(new BigNumber(6)).toNumber()] = p1;
        _nw242[(new BigNumber(7)).toNumber()] = p1;
        _nw242[(new BigNumber(8)).toNumber()] = p1;
        _nw242[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(p1, p1);
        _nw242[(new BigNumber(10)).toNumber()] = p1;
        _nw242[(new BigNumber(11)).toNumber()] = p1;
        _nw242[(new BigNumber(12)).toNumber()] = p1;
        _nw242[(new BigNumber(13)).toNumber()] = p1;
        _nw242[(new BigNumber(14)).toNumber()] = p1;
        _nw242[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("uatwgx"));
        _nw242[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(p1, (function (_pat_let36_0) {
          return function (_1691_dt__update__tmp_h1) {
            return function (_pat_let37_0) {
              return function (_1692_dt__update_hcf1_h0) {
                return function (_pat_let38_0) {
                  return function (_1693_dt__update_hcf2_h0) {
                    return _module.D0.create_DC1(_1692_dt__update_hcf1_h0, _1693_dt__update_hcf2_h0, (_1691_dt__update__tmp_h1).dtor_cf3);
                  }(_pat_let38_0);
                }(_pat_let_tv37);
              }(_pat_let37_0);
            }((_pat_let_tv36)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_pat_let_tv35).length))]);
          }(_pat_let36_0);
        }(_1686_v154)).dtor_cf3);
        _nw242[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), ((_1694_v153) => function (_1695_i12) {
          return _1694_v153;
        })(_1685_v153));
        _nw242[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("wocuatbi");
        _nw242[(new BigNumber(19)).toNumber()] = p1;
        _nw242[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(p1, p1);
        _nw242[(new BigNumber(21)).toNumber()] = (_module.D1.create_DC5(p0, p1, _dafny.Seq.update(_dafny.Seq.of((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], false), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(984)), ((_1696_v153) => function (_1697_i13) {
  return _1696_v153;
})(_1685_v153))).length), new BigNumber((_dafny.Seq.of((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], false)).length)), (_1673_v142)[_module.__default.safeIndex(_1672_v141, new BigNumber((_1673_v142).length))]))).dtor_cf8;
        _nw242[(new BigNumber(22)).toNumber()] = p1;
        _nw242[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(p1, p1);
        _nw242[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("jp");
        _nw242[(new BigNumber(25)).toNumber()] = _dafny.Seq.UnicodeFromString("scmvf");
        _1687_v155 = _nw242;
        let _index253 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_1687_v155).length));
        (_1687_v155)[_index253] = p1;
        let _1698_v156;
        _1698_v156 = _dafny.MultiSet.fromElements(_1672_v141);
        let _index254 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1681_v150).length));
        let _index255 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_1687_v155).length));
        let _rhs279 = (_dafny.MultiSet.fromElements(_1672_v141, _1672_v141, _1672_v141, new BigNumber((p1).length))).IsDisjointFrom(_1698_v156);
        let _rhs280 = (_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))];
        let _rhs281 = _dafny.Seq.Concat(_1676_v145, _1676_v145);
        let _rhs282 = p1;
        let _lhs189 = globalState;
        let _lhs190 = globalState;
        let _lhs191 = _1681_v150;
        let _lhs192 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1681_v150).length));
        let _lhs193 = _1687_v155;
        let _lhs194 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_1687_v155).length));
        _lhs189.f9 = _rhs279;
        _lhs190.f13 = _rhs280;
        _lhs191[_lhs192] = _rhs281;
        _lhs193[_lhs194] = _rhs282;
        let _1699_v157;
        _1699_v157 = _dafny.Set.fromElements((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))]);
        let _1700_v158;
        _1700_v158 = _dafny.Seq.of(_1699_v157);
        let _1701_v159;
        _1701_v159 = _dafny.MultiSet.fromElements(_1699_v157);
        _1673_v142 = (((_dafny.MultiSet.FromArray(_1700_v158)).IsProperSubsetOf(_1701_v159)) ? (_1673_v142) : (_1673_v142));
        let _1702_v160;
        let _nw243 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _1702_v160 = _nw243;
        _1702_v160 = _1702_v160;
        let _index256 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_1687_v155).length));
        (_1687_v155)[_index256] = p1;
      } else {
        r2 = (_dafny.ZERO).minus((_1672_v141).minus(_1672_v141));
        if ((((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(540)), function (_1703_i14) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })).length)).isLessThanOrEqualTo(_1672_v141)) ? ((((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))]) ? ((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))]) : ((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))]))) : ((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))]))) {
          let _1704_v161;
          _1704_v161 = _dafny.Seq.of(p1);
          let _1705_v162;
          _1705_v162 = _module.D0.create_DC1((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], _1704_v161, _dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_1706_i15) {
  return new _dafny.CodePoint('o'.codePointAt(0));
}));
          let _1707_v163;
          _1707_v163 = _dafny.MultiSet.fromElements(_1705_v162);
          let _1708_v164;
          let _nw244 = new _module.C0();
          _nw244.__ctor(_1707_v163, p0);
          _1708_v164 = _nw244;
          let _1709_v165;
          _1709_v165 = _dafny.Seq.UnicodeFromString("hrsxsmked");
          let _1710_v166;
          _1710_v166 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v141,false);
          let _1711_v167;
          _1711_v167 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v141,new BigNumber((_1673_v142).length));
          let _1712_v168;
          _1712_v168 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(_1672_v141, new BigNumber((_1710_v166).length)), new BigNumber((_1711_v167).length), _1672_v141);
          let _1713_v169;
          _1713_v169 = _module.D1.create_DC5((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], _1709_v165, _1673_v142);
          let _rhs283 = _dafny.Seq.Concat(_1709_v165, (_1713_v169).dtor_cf8);
          let _rhs284 = (_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))];
          let _rhs285 = _1712_v168;
          let _lhs195 = globalState;
          _1709_v165 = _rhs283;
          _lhs195.f9 = _rhs284;
          _1712_v168 = _rhs285;
          let _1714_v170;
          _1714_v170 = _dafny.Seq.of(_1672_v141, (((_1708_v164).f19) ? (_1672_v141) : (_1672_v141)), _1672_v141);
          _1714_v170 = _dafny.Seq.of(new BigNumber((_1671_v140).length), _1672_v141, _1672_v141, new BigNumber((_1671_v140).length));
          _1672_v141 = _1672_v141;
          let _1715_v171;
          let _nw245 = new _module.C5();
          _nw245.__ctor();
          _1715_v171 = _nw245;
          let _1716_v172;
          _1716_v172 = _dafny.Set.fromElements(_1715_v171);
          let _1717_v173;
          _1717_v173 = _module.D13.create_DC34(_module.__default.safeModuloInt(_1672_v141, new BigNumber((_1716_v172).length)));
          _1717_v173 = _1717_v173;
        } else {
          let _1718_v174;
          let _nw246 = new _module.C1();
          _nw246.__ctor();
          _1718_v174 = _nw246;
          let _1719_v175;
          let _nw247 = new _module.C6();
          _nw247.__ctor();
          _1719_v175 = _nw247;
          let _1720_v176;
          _1720_v176 = _dafny.Seq.of(_1672_v141);
          let _1721_v177;
          _1721_v177 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1720_v176);
          let _1722_v178;
          _1722_v178 = _dafny.Seq.of(_1721_v177);
          let _1723_v179;
          _1723_v179 = _module.D11.create_DC29((_dafny.MultiSet.fromElements(_1719_v175)).update(_1719_v175, _module.__default.abs(new BigNumber(((_1722_v178)[_module.__default.safeIndex(_1672_v141, new BigNumber((_1722_v178).length))]).length))));
          _1723_v179 = _1723_v179;
          (globalState).f9 = p0;
          let _1724_v180;
          _1724_v180 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(836)), function (_1725_i16) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }),_1672_v141);
          let _1726_v181;
          let _nw248 = new _module.C3();
          _nw248.__ctor((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], new BigNumber((_1505_v0).cardinality()));
          _1726_v181 = _nw248;
          let _1727_v182;
          _1727_v182 = _dafny.MultiSet.fromElements(_1726_v181);
          _1724_v180 = (_1724_v180).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), function (_1728_i17) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), new BigNumber((_1727_v182).cardinality()));
          let _1729_v183;
          let _nw249 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1729_v183 = _nw249;
          let _1730_v184;
          _1730_v184 = new _dafny.CodePoint('b'.codePointAt(0));
          let _index257 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1729_v183).length));
          (_1729_v183)[_index257] = _1730_v184;
          let _index258 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1729_v183).length));
          (_1729_v183)[_index258] = (p1)[_module.__default.safeIndex((new BigNumber((_1673_v142).length)).minus(_1672_v141), new BigNumber((p1).length))];
        }
        let _1731_v185;
        let _nw250 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1731_v185 = _nw250;
        let _index259 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1731_v185).length));
        (_1731_v185)[_index259] = new _dafny.CodePoint('c'.codePointAt(0));
        let _1732_v186;
        _1732_v186 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(_1672_v141, new BigNumber((p1).length))],_1671_v140);
        let _1733_v187;
        _1733_v187 = _dafny.Seq.of(_1672_v141, _1672_v141);
        let _index260 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length));
        let _index261 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1731_v185).length));
        let _rhs286 = (_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))];
        let _rhs287 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("wkpov"), p1);
        let _rhs288 = _module.__default.fm23(_dafny.Seq.IsPrefixOf(p1, p1), _1732_v186, (_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], (_1733_v187)[_module.__default.safeIndex(_1672_v141, new BigNumber((_1733_v187).length))], globalState);
        let _rhs289 = new BigNumber(986);
        let _rhs290 = _1672_v141;
        let _lhs196 = globalState;
        let _lhs197 = _1670_v139;
        let _lhs198 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length));
        let _lhs199 = _1731_v185;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_1731_v185).length));
        _lhs196.f3 = _rhs286;
        _lhs197[_lhs198] = _rhs287;
        _lhs199[_lhs200] = _rhs288;
        r2 = _rhs289;
        r2 = _rhs290;
        (globalState).f5 = true;
        let _1734_v188;
        _1734_v188 = _dafny.Set.fromElements(_1505_v0, _1505_v0);
        let _1735_v189;
        _1735_v189 = _dafny.MultiSet.fromElements(new BigNumber(759));
        let _1736_v190;
        _1736_v190 = _dafny.Map.Empty.slice().updateUnsafe((_1734_v188).IsDisjointFrom(_module.__default.fm43((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], globalState)),_1735_v189);
        let _1737_v191;
        _1737_v191 = _dafny.Map.Empty.slice().updateUnsafe(_1670_v139,_dafny.Seq.of(false, p0));
        _1736_v190 = (_1736_v190).update(_module.__default.fm2(_module.__default.fm23((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], _1732_v186, (_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))], new BigNumber((_1737_v191).length), globalState), _module.D0.create_DC3(_1674_v143), _1672_v141, globalState), _1735_v189);
      }
      let _1738_v192;
      _1738_v192 = _dafny.Set.fromElements(_module.__default.fm6(globalState), (_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))]);
      r0 = _1738_v192;
      r1 = !((_1670_v139)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1670_v139).length))]);
      r2 = _1672_v141;
      r3 = p0;
      return [r0, r1, r2, r3];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
