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
      return (true) === ((!(false)) && (!(true)));
    };
    static fm1(globalState) {
      return _dafny.Seq.UnicodeFromString("sxninyjx");
    };
    static fm2(globalState) {
      return new _dafny.CodePoint('s'.codePointAt(0));
    };
    static fm3(globalState) {
      return new BigNumber(332);
    };
    static fm4(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC3(_dafny.MultiSet.fromElements(false, false)), _module.D1.create_DC3(_dafny.MultiSet.fromElements(true))), _dafny.Seq.of(_module.D1.create_DC3(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(false))))), _module.D1.create_DC3(_dafny.MultiSet.fromElements(true, false)), _module.D1.create_DC3(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)))));
    };
    static fm16(p0, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of(true, !(!(true))), _dafny.Seq.of(true))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(false, false, true), _dafny.Seq.of(false)));
    };
    static fm17(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(960)), function (_0_i0) {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("lheve")).length)).multipliedBy(new BigNumber((_dafny.Seq.of(false)).length));
      });
    };
    static fm18(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("psg")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), function (_1_i0) {
        return _dafny.Seq.UnicodeFromString("wufuwmsuw");
      }), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("pfu"))));
    };
    static fm19(p0, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber(146)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(false, !(false))).cardinality())));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(!(!(!(false))), true), _dafny.Seq.of(false, false));
    };
    static fm21(p0, p1, globalState) {
      return _module.D4.create_DC12(true, !(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-894),_dafny.Seq.of(new BigNumber(283)))).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(328),_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(831)), function (_2_i0) {
  return new _dafny.CodePoint('p'.codePointAt(0));
})).length)))), ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(878), new BigNumber((_dafny.Seq.UnicodeFromString("du")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((function () {
  let _coll0 = new _dafny.Set();
  for (const _compr_0 of (_dafny.Seq.of(new BigNumber(241))).Elements) {
    let _3_v0 = _compr_0;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(241)), _3_v0)) {
      _coll0.add((_3_v0).plus(new BigNumber(-99)));
    }
  }
  return _coll0;
}()).length))).cardinality()))).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pnlbg")).length)));
    };
    static fm22(p0, p1, p2, globalState) {
      return _module.D8.create_DC20(_dafny.Seq.Concat(_dafny.Seq.of(!(!(false))), _dafny.Seq.of(true)));
    };
    static fm23(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_module.D11.create_DC31(false, new BigNumber(948), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-910)), function (_4_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length), true)).dtor_cf56,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(!(!(true))))));
    };
    static fm24(p0, p1, globalState) {
      return _module.D8.create_DC23();
    };
    static fm25(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.MultiSet.FromArray((_module.D14.create_DC38(_dafny.Seq.of(_module.D8.create_DC23()))).dtor_cf70)).Elements) {
          let _5_v0 = _compr_1;
          if ((_dafny.MultiSet.FromArray((_module.D14.create_DC38(_dafny.Seq.of(_module.D8.create_DC23()))).dtor_cf70)).contains(_5_v0)) {
            _coll1.add(_5_v0);
          }
        }
        return _coll1;
      }()).Union(_dafny.Set.fromElements(_module.D8.create_DC23())), _dafny.Set.fromElements(_module.D8.create_DC23()), (_dafny.Set.fromElements(_module.D8.create_DC23(), _module.D8.create_DC23(), _module.D8.create_DC23())).Union(_dafny.Set.fromElements(_module.D8.create_DC23(), _module.D8.create_DC23(), _module.D8.create_DC23())));
    };
    static fm26(globalState) {
      if ((_dafny.Set.fromElements(_dafny.Seq.of(true))).contains(_dafny.Seq.of(!(true)))) {
        return _module.D8.create_DC21(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(324))).length))).length), true, !(true), true, new BigNumber(194));
      } else {
        return _module.D8.create_DC21(new BigNumber((_dafny.Set.fromElements(new BigNumber(-729), (_dafny.ZERO).minus(new BigNumber(-109)))).length), false, true, true, new BigNumber(-579));
      }
    };
    static fm27(globalState) {
      return (function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)))).Elements) {
          let _6_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0))), _6_v0)) {
            _coll2.add(_6_v0);
          }
        }
        return _coll2;
      }()).Union((_dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0)))));
    };
    static fm28(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(708), new BigNumber((_dafny.Seq.UnicodeFromString("bnk")).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(true), true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(802),false)).length))).length), new BigNumber(475)), function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.Set.fromElements(new BigNumber(-250))).Elements) {
          let _7_v0 = _compr_3;
          if ((_dafny.Set.fromElements(new BigNumber(-250))).contains(_7_v0)) {
            _coll3.add((_7_v0).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)))).length))).length)));
          }
        }
        return _coll3;
      }())).Union((_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(769)), function (_8_i0) {
        return (_dafny.ZERO).minus(new BigNumber((function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(223), new BigNumber(890))) {
            let _9_v1 = _compr_4;
            if (((new BigNumber(223)).isLessThanOrEqualTo(_9_v1)) && ((_9_v1).isLessThan(new BigNumber(890)))) {
              _coll4.add(_module.__default.safeDivisionInt(_9_v1, new BigNumber(291)));
            }
          }
          return _coll4;
        }()).length));
      })).length)))).Union(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-278), new BigNumber((_dafny.Seq.of(new BigNumber(436))).length)), _dafny.Set.fromElements(new BigNumber((function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, !(false))).length),true)).Keys.Elements) {
            let _10_v2 = _compr_6;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, !(false))).length),true)).contains(_10_v2)) {
              _coll6.push([(_10_v2).multipliedBy(new BigNumber(15)),true]);
            }
          }
          return _coll6;
        }()).length)))).Elements) {
          let _11_v3 = _compr_5;
          if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-278), new BigNumber((_dafny.Seq.of(new BigNumber(436))).length)), _dafny.Set.fromElements(new BigNumber((function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, !(false))).length),true)).Keys.Elements) {
              let _10_v2 = _compr_7;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, !(false))).length),true)).contains(_10_v2)) {
                _coll7.push([(_10_v2).multipliedBy(new BigNumber(15)),true]);
              }
            }
            return _coll7;
          }()).length)))).contains(_11_v3)) {
            _coll5.add(_11_v3);
          }
        }
        return _coll5;
      }()));
    };
    static fm29(p0, p1, p2, globalState) {
      let _source0 = _module.D9.create_DC27(false, _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).length)));
      if (_source0.is_DC26) {
        let _12___mcc_h0 = (_source0).cf45;
        let _13___mcc_h1 = (_source0).cf46;
        let _14_cf46 = _13___mcc_h1;
        let _15_cf45 = _12___mcc_h0;
        return _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("efkgpwsh")).length)), function () {
          let _coll8 = new _dafny.Set();
          for (const _compr_8 of _dafny.IntegerRange(new BigNumber(809), new BigNumber(957))) {
            let _16_v0 = _compr_8;
            if (((new BigNumber(809)).isLessThanOrEqualTo(_16_v0)) && ((_16_v0).isLessThan(new BigNumber(957)))) {
              _coll8.add(_module.__default.safeModuloInt(_16_v0, (_14_cf46).f26));
            }
          }
          return _coll8;
        }());
      } else if (_source0.is_DC27) {
        let _17___mcc_h2 = (_source0).cf47;
        let _18___mcc_h3 = (_source0).cf48;
        let _19_cf48 = _18___mcc_h3;
        let _20_cf47 = _17___mcc_h2;
        return _dafny.Seq.Concat(_dafny.Seq.of(_19_cf48), _dafny.Seq.of(_19_cf48));
      } else {
        let _21___mcc_h4 = (_source0).cf44;
        let _22_cf44 = _21___mcc_h4;
        return _dafny.Seq.Concat(_dafny.Seq.of(function () {
          let _coll9 = new _dafny.Set();
          for (const _compr_9 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), function (_23_i0) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length))).Elements) {
            let _24_v1 = _compr_9;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), function (_23_i0) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            })).length)), _24_v1)) {
              _coll9.add((_24_v1).minus(new BigNumber(-706)));
            }
          }
          return _coll9;
        }()), _dafny.Seq.of(function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of _dafny.IntegerRange(new BigNumber(836), new BigNumber(-360))) {
            let _25_v2 = _compr_10;
            if (((new BigNumber(836)).isLessThanOrEqualTo(_25_v2)) && ((_25_v2).isLessThan(new BigNumber(-360)))) {
              _coll10.add(_module.__default.safeModuloInt(_25_v2, new BigNumber(13)));
            }
          }
          return _coll10;
        }()));
      }
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.FromArray(((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-865)), function (_26_i0) {
        return _module.D1.create_DC5(new BigNumber(320), true);
      })) : (_dafny.Seq.of(_module.D1.create_DC5(new BigNumber((_dafny.Seq.UnicodeFromString("dwmmy")).length), false), _module.D1.create_DC5(new BigNumber(301), true)))))).Difference(_dafny.MultiSet.fromElements(_module.D1.create_DC5(new BigNumber(-624), false), _module.D1.create_DC5(new BigNumber((_dafny.Seq.of(true)).length), false)));
    };
    static fm31(p0, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(154))).Difference(function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(530),false)).Keys.Elements) {
          let _27_v0 = _compr_11;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(530),false)).contains(_27_v0)) {
            _coll11.add(_module.__default.safeDivisionInt(_27_v0, new BigNumber((function () {
              let _coll12 = new _dafny.Set();
              for (const _compr_12 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_dafny.Seq.UnicodeFromString("bnkgmf"))).Keys.Elements) {
                let _28_v1 = _compr_12;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_dafny.Seq.UnicodeFromString("bnkgmf"))).contains(_28_v1)) {
                  _coll12.add(_28_v1);
                }
              }
              return _coll12;
            }()).length)));
          }
        }
        return _coll11;
      }())).Union(_dafny.Set.fromElements(new BigNumber(881)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      if (!(!(!(false) || (true)))) {
        return _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ehocjrafq")).length)));
      } else {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), function (_29_i0) {
          return _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, (_module.D8.create_DC21(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(387))).cardinality()), true, true, !(true), new BigNumber(990))).dtor_cf36, true)).cardinality()));
        });
      }
    };
    static fm33(p0, globalState) {
      return _module.D1.create_DC4(!(new BigNumber(-221)).isEqualTo((_module.D11.create_DC30(false, new BigNumber(-315))).dtor_cf52), (new BigNumber(-347)).plus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())), _dafny.Map.Empty.slice().updateUnsafe(false,!(true)));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC5(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), !(false));
    };
    static fm35(p0, p1, globalState) {
      return function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of (_dafny.Seq.Concat(_dafny.Seq.of(function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(-241), new BigNumber(396))) {
            let _30_v0 = _compr_14;
            if (((new BigNumber(-241)).isLessThanOrEqualTo(_30_v0)) && ((_30_v0).isLessThan(new BigNumber(396)))) {
              _coll14.push([_module.__default.safeDivisionInt(_30_v0, new BigNumber((_dafny.Seq.UnicodeFromString("xwfabehw")).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), function (_31_i0) {
                return new _dafny.CodePoint('l'.codePointAt(0));
              })).length))]);
            }
          }
          return _coll14;
        }()), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber(-954)), function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-481), new BigNumber(89))) {
            let _32_v1 = _compr_15;
            if (((new BigNumber(-481)).isLessThanOrEqualTo(_32_v1)) && ((_32_v1).isLessThan(new BigNumber(89)))) {
              _coll15.push([(_32_v1).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber(741))).length)),new BigNumber(-324)]);
            }
          }
          return _coll15;
        }()))).Elements) {
          let _33_v2 = _compr_13;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-241), new BigNumber(396))) {
              let _30_v0 = _compr_16;
              if (((new BigNumber(-241)).isLessThanOrEqualTo(_30_v0)) && ((_30_v0).isLessThan(new BigNumber(396)))) {
                _coll16.push([_module.__default.safeDivisionInt(_30_v0, new BigNumber((_dafny.Seq.UnicodeFromString("xwfabehw")).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), function (_31_i0) {
                  return new _dafny.CodePoint('l'.codePointAt(0));
                })).length))]);
              }
            }
            return _coll16;
          }()), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber(-954)), function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of _dafny.IntegerRange(new BigNumber(-481), new BigNumber(89))) {
              let _32_v1 = _compr_17;
              if (((new BigNumber(-481)).isLessThanOrEqualTo(_32_v1)) && ((_32_v1).isLessThan(new BigNumber(89)))) {
                _coll17.push([(_32_v1).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber(741))).length)),new BigNumber(-324)]);
              }
            }
            return _coll17;
          }())), _33_v2)) {
            _coll13.add(_33_v2);
          }
        }
        return _coll13;
      }();
    };
    static fm36(p0, p1, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).Union(_dafny.MultiSet.fromElements(true, !(false)))).Difference(_dafny.MultiSet.fromElements(!(false)));
    };
    static fm37(p0, p1, p2, globalState) {
      return _module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(893)), function (_34_i0) {
  return new _dafny.CodePoint('r'.codePointAt(0));
}));
    };
    static fm38(p0, p1, globalState) {
      if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false),true)).contains(_dafny.Set.fromElements(false))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(318)), function (_35_i0) {
          return new BigNumber(412);
        })).length),!(true))).Merge(function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of _dafny.IntegerRange(new BigNumber(918), new BigNumber(-354))) {
            let _36_v0 = _compr_18;
            if (((new BigNumber(918)).isLessThanOrEqualTo(_36_v0)) && ((_36_v0).isLessThan(new BigNumber(-354)))) {
              _coll18.push([_module.__default.safeModuloInt(_36_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), function (_37_i1) {
                return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(232)), function (_38_i2) {
                  return new _dafny.CodePoint('i'.codePointAt(0));
                })).length);
              })).length)),!(true)]);
            }
          }
          return _coll18;
        }());
      } else {
        return (_module.D15.create_DC40(_dafny.Map.Empty.slice().updateUnsafe((_module.D4.create_DC12(true, false, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-207))).length), new BigNumber(640))).cardinality()))).dtor_cf23,true))).dtor_cf73;
      }
    };
    static fm39(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("usfsg")).length)), (_module.D16.create_DC44(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-290))).length)))).dtor_cf82, (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(352), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(4))).cardinality())), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(884), new BigNumber(762))).length)))).cardinality()), (_dafny.ZERO).minus(new BigNumber(-430)))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(-57), (_dafny.ZERO).minus(new BigNumber(-119)))), (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-953),new BigNumber(886))).length),_module.D15.create_DC41(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(483))).length))).length), _dafny.Seq.of(new BigNumber(296)), _dafny.Seq.of(true), _dafny.Seq.UnicodeFromString("gp")))).length), new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(true), _dafny.Set.fromElements(false, (_module.D6.create_DC16(true, true)).dtor_cf30), _dafny.Set.fromElements(false, true), _dafny.Set.fromElements(false))).length), new BigNumber(580), new BigNumber(128), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(816), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-119)))).length), new BigNumber(131))).cardinality())))).Union(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(937),new BigNumber(-278))).Keys.Elements) {
          let _39_v0 = _compr_19;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(937),new BigNumber(-278))).contains(_39_v0)) {
            _coll19.add((_39_v0).minus(new BigNumber(313)));
          }
        }
        return _coll19;
      }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(204),false)).length), new BigNumber(707), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(!(!(true)))).length)), new BigNumber((_dafny.Seq.UnicodeFromString("k")).length))));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)))).Union((_dafny.MultiSet.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)))).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)))));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _40_v0;
      _40_v0 = false;
      if (_40_v0) {
        let _41_v1;
        _41_v1 = new BigNumber(9);
        let _42_v2;
        _42_v2 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.of(_module.__default.fm3(globalState)), _dafny.Seq.of(_41_v1, p2)),_40_v0);
        _41_v1 = new BigNumber((_42_v2).length);
        let _43_v3;
        _43_v3 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_41_v1, new BigNumber(469)),!(_40_v0));
        let _44_v4;
        _44_v4 = _dafny.Seq.of(new BigNumber((_43_v3).length), _41_v1, (_dafny.ZERO).minus(_41_v1));
        let _45_v5;
        _45_v5 = _dafny.Set.fromElements(_44_v4, _dafny.Seq.of(_41_v1), _44_v4);
        _45_v5 = _45_v5;
        let _46_v6;
        _46_v6 = _dafny.MultiSet.fromElements(p1, p1);
        let _47_v7;
        _47_v7 = new _dafny.CodePoint('m'.codePointAt(0));
        _46_v6 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), function (_48_i0) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        }), _dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((p1).length)), _47_v7));
        let _49_v8;
        let _init0 = ((_50_v0) => function (_51_i1) {
          return _50_v0;
        })(_40_v0);
        let _nw0 = Array((new BigNumber(24)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _49_v8 = _nw0;
        let _index0 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_49_v8).length));
        (_49_v8)[_index0] = !(_module.__default.fm0(_40_v0, _40_v0, globalState));
        let _index1 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_49_v8).length));
        let _rhs0 = _40_v0;
        let _rhs1 = true;
        let _lhs0 = globalState;
        let _lhs1 = _49_v8;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_49_v8).length));
        _lhs0.f2 = _rhs0;
        _lhs1[_lhs2] = _rhs1;
        let _index2 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_49_v8).length));
        (_49_v8)[_index2] = _module.__default.fm0(_dafny.Seq.IsPrefixOf(p1, p1), true, globalState);
      } else {
        let _52_v9;
        _52_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-2)), function (_53_i2) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }),p2);
        if (!(_52_v9).equals(_52_v9)) {
          let _54_v10;
          _54_v10 = new BigNumber(844);
          let _55_v11;
          _55_v11 = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,_40_v0);
          let _56_v12;
          _56_v12 = _module.D1.create_DC4(_40_v0, p2, (_55_v11).update(_40_v0, _40_v0));
          let _57_v13;
          _57_v13 = _dafny.MultiSet.fromElements((_56_v12).dtor_cf7, _40_v0);
          _54_v10 = (_module.__default.safeDivisionInt(p0, new BigNumber(323))).minus(_module.__default.safeModuloInt(new BigNumber((_57_v13).cardinality()), p0));
          let _58_v14;
          _58_v14 = _dafny.Set.fromElements(_40_v0);
          let _59_v15;
          _59_v15 = _dafny.Seq.of(new BigNumber((p1).length));
          let _60_v16;
          let _nw1 = new _module.C8();
          _nw1.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(909)), function (_61_i3) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }), _58_v14, _40_v0, _dafny.Seq.update(_dafny.Seq.Concat(_59_v15, _59_v15), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat(_59_v15, _59_v15)).length)), p3));
          _60_v16 = _nw1;
          let _62_v17;
          let _init1 = ((_63_v0) => function (_64_i4) {
            return _63_v0;
          })(_40_v0);
          let _nw2 = Array((new BigNumber(18)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw2.length); _i0_1++) {
            _nw2[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _62_v17 = _nw2;
          let _index3 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_62_v17).length));
          (_62_v17)[_index3] = _40_v0;
          let _index4 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_62_v17).length));
          (_62_v17)[_index4] = _40_v0;
          let _65_v18;
          let _nw3 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _65_v18 = _nw3;
          let _index5 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_65_v18).length));
          (_65_v18)[_index5] = p0;
          let _index6 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_65_v18).length));
          (_65_v18)[_index6] = (_module.__default.fm3(globalState)).minus(p3);
          let _66_v19;
          let _nw4 = new _module.C2();
          _nw4.__ctor(new BigNumber((_59_v15).length));
          _66_v19 = _nw4;
        } else {
          let _67_v20;
          let _nw5 = Array((new BigNumber(18)).toNumber()).fill(_module.D2.Default());
          _67_v20 = _nw5;
          let _68_v21;
          _68_v21 = _67_v20;
          _68_v21 = ((_40_v0) ? (_67_v20) : (_67_v20));
          let _69_v22;
          _69_v22 = _dafny.Seq.of(new BigNumber(672));
          let _70_v23;
          let _init2 = ((_71_p1) => function (_72_i6) {
            return _module.__default.safeDivisionInt(_72_i6, new BigNumber((_71_p1).length));
          })(p1);
          let _nw6 = Array((new BigNumber(28)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
            _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _70_v23 = _nw6;
          let _73_v24;
          _73_v24 = _dafny.Seq.of(_70_v23);
          let _74_v25;
          let _nw7 = Array((new BigNumber(13)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _69_v22;
          _nw7[(_dafny.ONE).toNumber()] = _69_v22;
          _nw7[(new BigNumber(2)).toNumber()] = _69_v22;
          _nw7[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(102)), ((_75_p0) => function (_76_i5) {
            return _75_p0;
          })(p0));
          _nw7[(new BigNumber(4)).toNumber()] = _69_v22;
          _nw7[(new BigNumber(5)).toNumber()] = _69_v22;
          _nw7[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(p2, new BigNumber((_73_v24).length), p2, p2, p0);
          _nw7[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_69_v22, _module.__default.fm17(globalState));
          _nw7[(new BigNumber(8)).toNumber()] = _module.__default.fm17(globalState);
          _nw7[(new BigNumber(9)).toNumber()] = _69_v22;
          _nw7[(new BigNumber(10)).toNumber()] = _69_v22;
          _nw7[(new BigNumber(11)).toNumber()] = _69_v22;
          _nw7[(new BigNumber(12)).toNumber()] = _69_v22;
          _74_v25 = _nw7;
          let _index7 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_74_v25).length));
          (_74_v25)[_index7] = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_40_v0)).length));
          let _77_v26;
          _77_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(_40_v0),_69_v22);
          let _78_v27;
          _78_v27 = _dafny.Seq.of((((_77_v26).contains(_40_v0)) ? ((_77_v26).get(_40_v0)) : (_dafny.Seq.of(p2, p0, p2))));
          let _index8 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_74_v25).length));
          (_74_v25)[_index8] = _dafny.Seq.Concat(_module.__default.fm17(globalState), (_78_v27)[_module.__default.safeIndex(p3, new BigNumber((_78_v27).length))]);
          let _79_v28;
          _79_v28 = new BigNumber(907);
          _79_v28 = p2;
          let _80_v29;
          _80_v29 = _dafny.MultiSet.fromElements(_79_v28);
          (globalState).f2 = ((_80_v29).Difference(_dafny.MultiSet.fromElements(p3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(327)), ((_81_p2, _82_v0) => function (_83_i7) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_81_p2,_82_v0)).length);
          })(p2, _40_v0))).length)))).IsSubsetOf(_dafny.MultiSet.FromArray((_74_v25)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_74_v25).length))]));
          _79_v28 = p0;
        }
        let _84_v30;
        let _init3 = ((_85_v0) => function (_86_i8) {
          return !(false) || (_85_v0);
        })(_40_v0);
        let _nw8 = Array((new BigNumber(28)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw8.length); _i0_3++) {
          _nw8[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _84_v30 = _nw8;
        let _index9 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_84_v30).length));
        (_84_v30)[_index9] = _40_v0;
        let _index10 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_84_v30).length));
        (_84_v30)[_index10] = _40_v0;
        let _87_v31;
        let _nw9 = new _module.C5();
        _nw9.__ctor(_40_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), ((_88_p0) => function (_89_i9) {
          return _88_p0;
        })(p0)));
        _87_v31 = _nw9;
        let _90_v32;
        _90_v32 = new _dafny.CodePoint('m'.codePointAt(0));
        let _91_v33;
        _91_v33 = _dafny.Map.Empty.slice().updateUnsafe(_87_v31,_90_v32);
        (globalState).f2 = !(_91_v33).contains(_87_v31);
        let _92_v34;
        let _nw10 = new _module.C5();
        _nw10.__ctor(!((_84_v30)[_module.__default.safeIndex(new BigNumber(37), new BigNumber((_84_v30).length))]) || (_87_v31.f16), _dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), ((_93_p0) => function (_94_i10) {
          return _93_p0;
        })(p0)));
        _92_v34 = _nw10;
        _90_v32 = _90_v32;
      }
      let _95_v35;
      _95_v35 = new BigNumber(-579);
      let _96_v36;
      _96_v36 = _module.D4.create_DC12(true, false, p2);
      let _pat_let_tv0 = p3;
      let _pat_let_tv1 = p3;
      let _pat_let_tv2 = _40_v0;
      let _pat_let_tv3 = _40_v0;
      let _pat_let_tv4 = _40_v0;
      let _pat_let_tv5 = _40_v0;
      let _pat_let_tv6 = _40_v0;
      _95_v35 = function (_source1) {
        if (_source1.is_DC12) {
          let _97___mcc_h0 = (_source1).cf21;
          let _98___mcc_h1 = (_source1).cf22;
          let _99___mcc_h2 = (_source1).cf23;
          let _100_cf23 = _99___mcc_h2;
          let _101_cf22 = _98___mcc_h1;
          let _102_cf21 = _97___mcc_h0;
          return _pat_let_tv0;
        } else if (_source1.is_DC13) {
          let _103___mcc_h3 = (_source1).cf24;
          let _104___mcc_h4 = (_source1).cf25;
          let _105___mcc_h5 = (_source1).cf26;
          let _106_cf26 = _105___mcc_h5;
          let _107_cf25 = _104___mcc_h4;
          let _108_cf24 = _103___mcc_h3;
          return (_dafny.ZERO).minus(_pat_let_tv1);
        } else {
          let _109___mcc_h6 = (_source1).cf20;
          let _110_cf20 = _109___mcc_h6;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv2, !(_pat_let_tv3), _pat_let_tv4), _dafny.Seq.of(_pat_let_tv5, _pat_let_tv6))).length));
        }
      }(_96_v36);
      let _111_v37;
      _111_v37 = new _dafny.CodePoint('j'.codePointAt(0));
      let _112_v38;
      _112_v38 = _dafny.Seq.of(_40_v0);
      let _113_v39;
      _113_v39 = _dafny.Map.Empty.slice().updateUnsafe(_111_v37,_112_v38);
      let _114_v40;
      _114_v40 = _dafny.Seq.of(_112_v38);
      _113_v39 = (_113_v39).update(_111_v37, _dafny.Seq.Concat(_112_v38, (_114_v40)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("dsxwg")).length), new BigNumber((_114_v40).length))]));
      let _115_v41;
      _115_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
      let _116_v42;
      _116_v42 = _dafny.Seq.of(_95_v35, p3);
      let _117_v43;
      let _nw11 = new _module.C3();
      _nw11.__ctor(p1, _40_v0, _116_v42);
      _117_v43 = _nw11;
      let _118_v44;
      _118_v44 = _module.D12.create_DC35(_40_v0, _40_v0, _115_v41, true, _117_v43);
      let _119_v45;
      let _nw12 = Array((new BigNumber(12)).toNumber());
      _nw12[(_dafny.ZERO).toNumber()] = _40_v0;
      _nw12[(_dafny.ONE).toNumber()] = (_118_v44).dtor_cf64;
      _nw12[(new BigNumber(2)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(3)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(4)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(5)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(6)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(7)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(8)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(9)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(10)).toNumber()] = _40_v0;
      _nw12[(new BigNumber(11)).toNumber()] = false;
      _119_v45 = _nw12;
      let _120_v46;
      _120_v46 = _module.D0.create_DC2(new BigNumber((_dafny.MultiSet.FromArray(_112_v38)).cardinality()), p0, p3, _119_v45);
      let _121_v47;
      let _nw13 = new _module.C6();
      _nw13.__ctor(function (_pat_let0_0) {
        return function (_122_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_123_dt__update_hcf3_h0) {
              return _module.D0.create_DC2((_122_dt__update__tmp_h0).dtor_cf2, _123_dt__update_hcf3_h0, (_122_dt__update__tmp_h0).dtor_cf4, (_122_dt__update__tmp_h0).dtor_cf5);
            }(_pat_let1_0);
          }(new BigNumber(-516));
        }(_pat_let0_0);
      }(_120_v46));
      _121_v47 = _nw13;
      let _124_v48;
      _124_v48 = _dafny.Set.fromElements(_121_v47);
      let _hi0 = new BigNumber((_124_v48).length);
      for (let _125_i11 = _95_v35; _125_i11.isLessThan(_hi0); _125_i11 = _125_i11.plus(_dafny.ONE)) {
        let _126_v49;
        let _nw14 = Array((new BigNumber(5)).toNumber());
        _nw14[(_dafny.ZERO).toNumber()] = _119_v45;
        _nw14[(_dafny.ONE).toNumber()] = _119_v45;
        _nw14[(new BigNumber(2)).toNumber()] = _119_v45;
        _nw14[(new BigNumber(3)).toNumber()] = ((_40_v0) ? (_119_v45) : (_119_v45));
        _nw14[(new BigNumber(4)).toNumber()] = _119_v45;
        _126_v49 = _nw14;
        let _index11 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_126_v49).length));
        (_126_v49)[_index11] = _119_v45;
        let _127_v50;
        _127_v50 = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,new BigNumber(260));
        let _128_v51;
        let _nw15 = Array((new BigNumber(3)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = _116_v42;
        _nw15[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), ((_129_i11) => function (_130_i12) {
          return _129_i11;
        })(_125_i11));
        _nw15[(new BigNumber(2)).toNumber()] = _116_v42;
        _128_v51 = _nw15;
        let _131_v52;
        _131_v52 = _module.D2.create_DC8(_127_v50, _128_v51, _40_v0, _40_v0, _125_i11);
        let _132_v53;
        _132_v53 = _module.D6.create_DC17(_131_v52);
        let _index12 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_126_v49).length));
        let _rhs2 = _119_v45;
        let _rhs3 = _132_v53;
        let _lhs3 = _126_v49;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_126_v49).length));
        _lhs3[_lhs4] = _rhs2;
        _132_v53 = _rhs3;
        _111_v37 = new _dafny.CodePoint('v'.codePointAt(0));
        let _133_v54;
        let _init4 = ((_134_v37, _135_p1) => function (_136_i13) {
          return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), ((_137_v37) => function (_138_i14) {
            return _137_v37;
          })(_134_v37)), _135_p1);
        })(_111_v37, p1);
        let _nw16 = Array((new BigNumber(28)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw16.length); _i0_4++) {
          _nw16[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _133_v54 = _nw16;
        let _index13 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_133_v54).length));
        (_133_v54)[_index13] = p1;
        let _index14 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_133_v54).length));
        (_133_v54)[_index14] = (_117_v43).f24;
        let _rhs4 = !(_40_v0);
        let _rhs5 = (_112_v38)[_module.__default.safeIndex(new BigNumber((_116_v42).length), new BigNumber((_112_v38).length))];
        let _lhs5 = globalState;
        let _lhs6 = globalState;
        _lhs5.f2 = _rhs4;
        _lhs6.f2 = _rhs5;
      }
      let _139_v55;
      _139_v55 = _dafny.Set.fromElements(_module.__default.fm0(_40_v0, _40_v0, globalState), _40_v0);
      let _140_v56;
      _140_v56 = _dafny.Seq.of(_139_v55, _139_v55, _139_v55);
      if (!(((_140_v56)[_module.__default.safeIndex((_116_v42)[_module.__default.safeIndex(_95_v35, new BigNumber((_116_v42).length))], new BigNumber((_140_v56).length))]).IsSubsetOf(_139_v55)) || (false)) {
        let _141_v57;
        _141_v57 = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,p3);
        let _142_v58;
        let _nw17 = new _module.C2();
        _nw17.__ctor(new BigNumber((_dafny.MultiSet.fromElements(_40_v0)).cardinality()));
        _142_v58 = _nw17;
        let _143_v59;
        _143_v59 = _dafny.Map.Empty.slice().updateUnsafe(p2,_142_v58);
        let _144_v60;
        _144_v60 = _dafny.Seq.of(_143_v59, _143_v59);
        let _145_v61;
        let _nw18 = new _module.C5();
        _nw18.__ctor(_40_v0, (_121_v47).fm6(_95_v35, _40_v0, _111_v37, p1, globalState));
        _145_v61 = _nw18;
        let _146_v62;
        _146_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,_40_v0);
        let _147_v63;
        _147_v63 = _dafny.Map.Empty.slice().updateUnsafe(_145_v61,_dafny.Map.Empty.slice().updateUnsafe(_40_v0,new BigNumber((_146_v62).length)));
        let _148_v64;
        let _nw19 = Array((new BigNumber(23)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _141_v57;
        _nw19[(_dafny.ONE).toNumber()] = (_141_v57).update(_40_v0, new BigNumber(194));
        _nw19[(new BigNumber(2)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(3)).toNumber()] = (_141_v57).Merge(_141_v57);
        _nw19[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,p2);
        _nw19[(new BigNumber(5)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(6)).toNumber()] = (_141_v57).Merge(_141_v57);
        _nw19[(new BigNumber(7)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(8)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(9)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(10)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(11)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(12)).toNumber()] = (_141_v57).Merge(_141_v57);
        _nw19[(new BigNumber(13)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,new BigNumber((_116_v42).length));
        _nw19[(new BigNumber(15)).toNumber()] = (_141_v57).Merge(_141_v57);
        _nw19[(new BigNumber(16)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(17)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(18)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(19)).toNumber()] = (_141_v57).Merge((_141_v57).update(_40_v0, new BigNumber((_144_v60).length)));
        _nw19[(new BigNumber(20)).toNumber()] = (((_147_v63).contains(_145_v61)) ? ((_147_v63).get(_145_v61)) : (_141_v57));
        _nw19[(new BigNumber(21)).toNumber()] = _141_v57;
        _nw19[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_145_v61.f16,p3);
        _148_v64 = _nw19;
        let _149_v65;
        _149_v65 = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,(_141_v57).update(_40_v0, (_142_v58).f25));
        let _index15 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_148_v64).length));
        (_148_v64)[_index15] = (((_149_v65).contains(_40_v0)) ? ((_149_v65).get(_40_v0)) : (_141_v57));
        let _150_v66;
        let _nw20 = new _module.C0();
        _nw20.__ctor(p3, _119_v45);
        _150_v66 = _nw20;
        let _151_v67;
        let _nw21 = new _module.C1();
        _nw21.__ctor(((_142_v58).f25).minus(new BigNumber(903)), new BigNumber((_dafny.Set.fromElements(_dafny.Seq.update(_112_v38, _module.__default.safeIndex(p0, new BigNumber((_112_v38).length)), false))).length), (_112_v38)[_module.__default.safeIndex((_module.D9.create_DC26((_142_v58).f25, _150_v66)).dtor_cf45, new BigNumber((_112_v38).length))], _116_v42, _95_v35);
        _151_v67 = _nw21;
        let _index16 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_148_v64).length));
        let _rhs6 = _141_v57;
        let _rhs7 = (_142_v58).f25;
        let _rhs8 = _151_v67;
        let _lhs7 = _148_v64;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_148_v64).length));
        _lhs7[_lhs8] = _rhs6;
        _95_v35 = _rhs7;
        _151_v67 = _rhs8;
        let _152_v68;
        let _nw22 = Array((new BigNumber(9)).toNumber());
        _nw22[(_dafny.ZERO).toNumber()] = (_151_v67).fm13(_151_v67.f16, globalState);
        _nw22[(_dafny.ONE).toNumber()] = p3;
        _nw22[(new BigNumber(2)).toNumber()] = (_151_v67).f28;
        _nw22[(new BigNumber(3)).toNumber()] = (_142_v58).f25;
        _nw22[(new BigNumber(4)).toNumber()] = p0;
        _nw22[(new BigNumber(5)).toNumber()] = p2;
        _nw22[(new BigNumber(6)).toNumber()] = p2;
        _nw22[(new BigNumber(7)).toNumber()] = new BigNumber((p1).length);
        _nw22[(new BigNumber(8)).toNumber()] = (_142_v58).f25;
        _152_v68 = _nw22;
        let _153_v69;
        let _nw23 = Array((new BigNumber(13)).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = ((_151_v67.f16) ? (_152_v68) : (_152_v68));
        _nw23[(_dafny.ONE).toNumber()] = _152_v68;
        _nw23[(new BigNumber(2)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(3)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(4)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(5)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(6)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(7)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(8)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(9)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(10)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(11)).toNumber()] = _152_v68;
        _nw23[(new BigNumber(12)).toNumber()] = _152_v68;
        _153_v69 = _nw23;
        let _index17 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_153_v69).length));
        (_153_v69)[_index17] = _152_v68;
        let _index18 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_153_v69).length));
        (_153_v69)[_index18] = _152_v68;
        if (_145_v61.f16) {
          _95_v35 = _95_v35;
          (globalState).f5 = _40_v0;
          let _154_v70;
          let _nw24 = Array((new BigNumber(27)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_150_v66).f26);
          _nw24[(_dafny.ONE).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(2)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(3)).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-381)), ((_155_v67) => function (_156_i15) {
            return (_155_v67).f28;
          })(_151_v67));
          _nw24[(new BigNumber(5)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(6)).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), ((_157_v61) => function (_158_i16) {
            return new BigNumber(((_157_v61).f17).length);
          })(_145_v61));
          _nw24[(new BigNumber(8)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(9)).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(10)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(11)).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(12)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(13)).toNumber()] = _116_v42;
          _nw24[(new BigNumber(14)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(15)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(503)), ((_159_v67) => function (_160_i17) {
            return (_159_v67).f28;
          })(_151_v67)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(503)), ((_161_v67) => function (_162_i17) {
            return (_161_v67).f28;
          })(_151_v67))).length)), new BigNumber(-749));
          _nw24[(new BigNumber(17)).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(p2);
          _nw24[(new BigNumber(19)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(20)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(21)).toNumber()] = _116_v42;
          _nw24[(new BigNumber(22)).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(23)).toNumber()] = (_151_v67).f17;
          _nw24[(new BigNumber(24)).toNumber()] = (_145_v61).f17;
          _nw24[(new BigNumber(25)).toNumber()] = _116_v42;
          _nw24[(new BigNumber(26)).toNumber()] = (_151_v67).f17;
          _154_v70 = _nw24;
          let _163_v71;
          _163_v71 = _module.D2.create_DC8(_dafny.Map.Empty.slice().updateUnsafe(_151_v67.f16,_95_v35), _154_v70, _40_v0, _145_v61.f16, _95_v35);
          let _164_v72;
          _164_v72 = _module.D6.create_DC17(_163_v71);
          let _pat_let_tv7 = _163_v71;
          let _165_v73;
          _165_v73 = _dafny.MultiSet.fromElements(function (_pat_let2_0) {
            return function (_166_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_167_dt__update_hcf31_h0) {
                  return _module.D6.create_DC17(_167_dt__update_hcf31_h0);
                }(_pat_let3_0);
              }(_pat_let_tv7);
            }(_pat_let2_0);
          }(_164_v72), _164_v72, _164_v72);
          let _168_v74;
          _168_v74 = _dafny.Seq.of(_165_v73);
          let _169_v75;
          _169_v75 = _168_v74;
          let _170_v76;
          _170_v76 = _dafny.Set.fromElements(_169_v75);
          let _171_v77;
          _171_v77 = _dafny.MultiSet.fromElements((_170_v76).IsSubsetOf(_170_v76));
          let _index19 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_119_v45).length));
          (_119_v45)[_index19] = _151_v67.f16;
          let _172_v78;
          let _nw25 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
          _172_v78 = _nw25;
          let _index20 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_172_v78).length));
          (_172_v78)[_index20] = ((_151_v67.f16) ? (_112_v38) : (_112_v38));
          let _173_v79;
          _173_v79 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),_151_v67);
          let _174_v80;
          _174_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_173_v79).length),_150_v66);
          let _index21 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_119_v45).length));
          let _index22 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_172_v78).length));
          let _rhs9 = _145_v61.f16;
          let _rhs10 = _171_v77;
          let _rhs11 = true;
          let _rhs12 = (_114_v40)[_module.__default.safeIndex(_95_v35, new BigNumber((_114_v40).length))];
          let _rhs13 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.fm3(globalState), new BigNumber((_174_v80).length)));
          let _lhs9 = _145_v61;
          let _lhs10 = _119_v45;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_119_v45).length));
          let _lhs12 = _172_v78;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_172_v78).length));
          _lhs9.f16 = _rhs9;
          _171_v77 = _rhs10;
          _lhs10[_lhs11] = _rhs11;
          _lhs12[_lhs13] = _rhs12;
          _95_v35 = _rhs13;
          (_150_v66).f27 = _150_v66.f27;
          (globalState).f5 = (_40_v0) || ((_119_v45)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_119_v45).length))]);
        } else {
          (_145_v61).f16 = !(_145_v61.f16);
          let _175_v81;
          _175_v81 = _module.D12.create_DC34(_119_v45, p3, _152_v68, _111_v37);
          let _176_v82;
          _176_v82 = _dafny.Seq.of(_150_v66.f27);
          let _177_v83;
          let _nw26 = Array((new BigNumber(10)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = _150_v66.f27;
          _nw26[(_dafny.ONE).toNumber()] = _150_v66.f27;
          _nw26[(new BigNumber(2)).toNumber()] = (_175_v81).dtor_cf59;
          _nw26[(new BigNumber(3)).toNumber()] = _150_v66.f27;
          _nw26[(new BigNumber(4)).toNumber()] = _150_v66.f27;
          _nw26[(new BigNumber(5)).toNumber()] = _150_v66.f27;
          _nw26[(new BigNumber(6)).toNumber()] = _150_v66.f27;
          _nw26[(new BigNumber(7)).toNumber()] = _150_v66.f27;
          _nw26[(new BigNumber(8)).toNumber()] = (_176_v82)[_module.__default.safeIndex((_142_v58).f25, new BigNumber((_176_v82).length))];
          _nw26[(new BigNumber(9)).toNumber()] = _119_v45;
          _177_v83 = _nw26;
          let _178_v84;
          _178_v84 = _dafny.Map.Empty.slice().updateUnsafe(_111_v37,_177_v83);
          _178_v84 = (_178_v84).update(_111_v37, _177_v83);
          let _179_v85;
          _179_v85 = _dafny.MultiSet.fromElements(_145_v61.f16);
          (globalState).f2 = (_179_v85).IsDisjointFrom(_179_v85);
          let _180_v86;
          let _nw27 = new _module.C7();
          _nw27.__ctor(_152_v68, _151_v67.f16, (_145_v61).f17);
          _180_v86 = _nw27;
          _180_v86 = _180_v86;
          let _181_v87;
          let _nw28 = new _module.C5();
          _nw28.__ctor(_151_v67.f16, (_145_v61).f17);
          _181_v87 = _nw28;
          let _182_v88;
          _182_v88 = _dafny.Map.Empty.slice().updateUnsafe(_181_v87,_40_v0);
          let _183_v89;
          _183_v89 = _dafny.Map.Empty.slice().updateUnsafe(_145_v61.f16,_145_v61.f16);
          let _184_v90;
          _184_v90 = _module.D1.create_DC4(_145_v61.f16, new BigNumber((_182_v88).length), ((_145_v61.f16) ? (_183_v89) : (_183_v89)));
          let _rhs14 = ((_145_v61.f16) ? (_dafny.areEqual(_112_v38, _dafny.Seq.of(_151_v67.f16, _145_v61.f16))) : (_40_v0));
          let _rhs15 = (((((_121_v47).f21).dtor_cf2).isLessThanOrEqualTo(p3)) ? (_184_v90) : (_184_v90));
          let _lhs14 = _151_v67;
          _lhs14.f16 = _rhs14;
          _184_v90 = _rhs15;
        }
        let _index23 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_119_v45).length));
        (_119_v45)[_index23] = _151_v67.f16;
        let _index24 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_119_v45).length));
        (_119_v45)[_index24] = _40_v0;
        (globalState).f2 = !(_40_v0);
      } else {
        let _185_v91;
        _185_v91 = _dafny.MultiSet.fromElements(_40_v0, _40_v0, _40_v0);
        let _186_v92;
        _186_v92 = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_116_v42).length),_40_v0)).length));
        let _187_v93;
        _187_v93 = _dafny.MultiSet.fromElements(_185_v91, _module.__default.fm36(new BigNumber((_186_v92).length), _40_v0, globalState), _dafny.MultiSet.fromElements(_40_v0));
        (globalState).f2 = !((_187_v93).Intersect(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_40_v0)))).contains((_185_v91).update(_40_v0, _module.__default.abs(_95_v35)));
        _112_v38 = _module.__default.fm20(_40_v0, _40_v0, !(_40_v0), _module.__default.safeModuloInt(new BigNumber(577), p2), globalState);
        (globalState).f5 = (_40_v0) === (_40_v0);
        let _188_v94;
        let _nw29 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _188_v94 = _nw29;
        let _index25 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_188_v94).length));
        (_188_v94)[_index25] = _111_v37;
        let _index26 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_188_v94).length));
        (_188_v94)[_index26] = _111_v37;
        _40_v0 = (_112_v38)[_module.__default.safeIndex(((_40_v0) ? ((((_115_v41).contains(new BigNumber((p1).length))) ? ((_115_v41).get(new BigNumber((p1).length))) : (new BigNumber(487)))) : (p0)), new BigNumber((_112_v38).length))];
      }
      if ((((!(_40_v0)) && (_40_v0)) ? (_dafny.Seq.contains(_112_v38, _40_v0)) : (!(!_dafny.areEqual(p1, _dafny.Seq.UnicodeFromString("ykfd")))))) {
        (globalState).f5 = !(p2).isEqualTo(p0);
        _95_v35 = _95_v35;
        let _189_v95;
        _189_v95 = _dafny.Seq.UnicodeFromString("vwqu");
        _189_v95 = _189_v95;
        let _190_v96;
        let _nw30 = new _module.C1();
        _nw30.__ctor(new BigNumber(527), ((_40_v0) ? (p3) : (p0)), false, _116_v42, p0);
        _190_v96 = _nw30;
        let _191_v97;
        _191_v97 = _dafny.MultiSet.fromElements(_111_v37, _111_v37);
        let _rhs16 = _module.__default.safeModuloInt(new BigNumber((((_40_v0) ? (_191_v97) : ((_module.__default.fm40(p3, _40_v0, _139_v55, new BigNumber(582), globalState)).update(_111_v37, _module.__default.abs(p0))))).cardinality()), _module.__default.safeModuloInt(p2, p2));
        let _rhs17 = _190_v96;
        _95_v35 = _rhs16;
        _190_v96 = _rhs17;
        let _192_v98;
        _192_v98 = _dafny.MultiSet.fromElements(_112_v38);
        let _index27 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_119_v45).length));
        (_119_v45)[_index27] = (_dafny.MultiSet.fromElements(_112_v38)).IsProperSubsetOf(_192_v98);
        let _193_v99;
        _193_v99 = _dafny.Set.fromElements((_117_v43).fm11(_112_v38, _40_v0, (_dafny.ZERO).minus(new BigNumber(-325)), _40_v0, globalState), (_117_v43).f24, _dafny.Seq.UnicodeFromString("vgwwcrgf"));
        let _194_v100;
        _194_v100 = _dafny.MultiSet.fromElements(p0, p3);
        let _index28 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_119_v45).length));
        (_119_v45)[_index28] = (_194_v100).IsSubsetOf(_194_v100);
        let _195_v101;
        let _nw31 = new _module.C8();
        _nw31.__ctor(_189_v95, _139_v55, _40_v0, _dafny.Seq.update(_116_v42, _module.__default.safeIndex(p2, new BigNumber((_116_v42).length)), (_190_v96).f30));
        _195_v101 = _nw31;
        let _196_v102;
        _196_v102 = _dafny.Set.fromElements(_195_v101, _195_v101);
        let _197_v103;
        _197_v103 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-662)), ((_198_v37) => function (_199_i18) {
          return _198_v37;
        })(_111_v37)),new BigNumber((_dafny.Seq.of(!(_40_v0), false, _40_v0)).length));
        let _index29 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_119_v45).length));
        let _index30 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_119_v45).length));
        let _rhs18 = (_196_v102).IsDisjointFrom(_196_v102);
        let _rhs19 = function () {
          let _coll20 = new _dafny.Set();
          for (const _compr_20 of (_197_v103).Keys.Elements) {
            let _200_v104 = _compr_20;
            if ((_197_v103).contains(_200_v104)) {
              _coll20.add(_200_v104);
            }
          }
          return _coll20;
        }();
        let _rhs20 = !(_40_v0);
        let _rhs21 = _module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p2));
        let _lhs15 = _119_v45;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_119_v45).length));
        let _lhs17 = _119_v45;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_119_v45).length));
        _lhs15[_lhs16] = _rhs18;
        _193_v99 = _rhs19;
        _lhs17[_lhs18] = _rhs20;
        _95_v35 = _rhs21;
      } else {
        let _201_v105;
        let _nw32 = new _module.C1();
        _nw32.__ctor(_95_v35, p3, _40_v0, _116_v42, p2);
        _201_v105 = _nw32;
        let _202_v106;
        _202_v106 = _dafny.Map.Empty.slice().updateUnsafe(_201_v105,_119_v45);
        let _203_v107;
        _203_v107 = _dafny.Map.Empty.slice().updateUnsafe(((_40_v0) ? (_119_v45) : ((((_202_v106).contains(_201_v105)) ? ((_202_v106).get(_201_v105)) : (_119_v45)))),_201_v105.f16);
        _203_v107 = (_203_v107).Merge(_203_v107);
        let _204_v108;
        let _init5 = ((_205_v105) => function (_206_i19) {
          return _module.D1.create_DC5((_205_v105).f28, _205_v105.f16);
        })(_201_v105);
        let _nw33 = Array((new BigNumber(13)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw33.length); _i0_5++) {
          _nw33[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _204_v108 = _nw33;
        let _index31 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_204_v108).length));
        (_204_v108)[_index31] = _module.__default.fm34(new BigNumber(13), new BigNumber(795), _139_v55, new BigNumber((_112_v38).length), globalState);
        let _207_v109;
        _207_v109 = _module.D1.create_DC5((_201_v105).f28, true);
        let _pat_let_tv8 = _40_v0;
        let _pat_let_tv9 = _40_v0;
        let _index32 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_204_v108).length));
        let _rhs22 = function (_pat_let4_0) {
          return function (_208_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_209_dt__update_hcf11_h0) {
                return _module.D1.create_DC5((_208_dt__update__tmp_h2).dtor_cf10, _209_dt__update_hcf11_h0);
              }(_pat_let5_0);
            }(((_pat_let_tv9) ? (true) : (_pat_let_tv8)));
          }(_pat_let4_0);
        }(_207_v109);
        let _rhs23 = (p0).isLessThan(p2);
        let _lhs19 = _204_v108;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_204_v108).length));
        _lhs19[_lhs20] = _rhs22;
        _40_v0 = _rhs23;
        let _210_v110;
        let _nw34 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _210_v110 = _nw34;
        let _index33 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_210_v110).length));
        (_210_v110)[_index33] = _111_v37;
        let _index34 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_210_v110).length));
        (_210_v110)[_index34] = _111_v37;
        let _211_v111;
        _211_v111 = _module.D2.create_DC7(_139_v55);
        let _index35 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_119_v45).length));
        (_119_v45)[_index35] = _201_v105.f16;
        let _212_v113;
        _212_v113 = _dafny.MultiSet.fromElements((_210_v110)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_210_v110).length))], (_210_v110)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_210_v110).length))], _111_v37);
        let _213_v114;
        _213_v114 = _module.D9.create_DC25(function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of (_212_v113).Elements) {
    let _214_v112 = _compr_21;
    if ((_212_v113).contains(_214_v112)) {
      _coll21.push([_214_v112,true]);
    }
  }
  return _coll21;
}());
        let _index36 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_119_v45).length));
        let _rhs24 = _211_v111;
        let _rhs25 = ((_117_v43).f24)[_module.__default.safeIndex(_95_v35, new BigNumber(((_117_v43).f24).length))];
        let _rhs26 = ((_201_v105.f16) ? (true) : (_201_v105.f16));
        let _rhs27 = !(_40_v0);
        let _rhs28 = _module.__default.fm38(_213_v114, (_201_v105).f28, globalState);
        let _lhs21 = globalState;
        let _lhs22 = _119_v45;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_119_v45).length));
        let _lhs24 = _201_v105;
        let _lhs25 = globalState;
        _211_v111 = _rhs24;
        _lhs21.f15 = _rhs25;
        _lhs22[_lhs23] = _rhs26;
        _lhs24.f16 = _rhs27;
        _lhs25.f8 = _rhs28;
        (_201_v105).f16 = !(_40_v0) || (!(true));
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _215_v0;
      _215_v0 = _dafny.Seq.UnicodeFromString("dg");
      let _216_v1;
      _216_v1 = new BigNumber(879);
      let _217_v2;
      _217_v2 = new _dafny.CodePoint('t'.codePointAt(0));
      let _218_v3;
      let _nw35 = Array((new BigNumber(12)).toNumber());
      _nw35[(_dafny.ZERO).toNumber()] = _215_v0;
      _nw35[(_dafny.ONE).toNumber()] = _215_v0;
      _nw35[(new BigNumber(2)).toNumber()] = _215_v0;
      _nw35[(new BigNumber(3)).toNumber()] = _215_v0;
      _nw35[(new BigNumber(4)).toNumber()] = _215_v0;
      _nw35[(new BigNumber(5)).toNumber()] = _215_v0;
      _nw35[(new BigNumber(6)).toNumber()] = _215_v0;
      _nw35[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("k");
      _nw35[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_215_v0, _module.__default.safeIndex(_216_v1, new BigNumber((_215_v0).length)), _217_v2);
      _nw35[(new BigNumber(9)).toNumber()] = _215_v0;
      _nw35[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_215_v0, _module.__default.safeIndex(_216_v1, new BigNumber((_215_v0).length)), _217_v2);
      _nw35[(new BigNumber(11)).toNumber()] = _215_v0;
      _218_v3 = _nw35;
      let _219_v4;
      let _nw36 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
      _219_v4 = _nw36;
      let _220_v5;
      let _init6 = function (_221_i0) {
        return true;
      };
      let _nw37 = Array((new BigNumber(26)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw37.length); _i0_6++) {
        _nw37[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _220_v5 = _nw37;
      let _222_v6;
      _222_v6 = _dafny.Seq.of(_220_v5);
      let _223_v7;
      _223_v7 = true;
      let _224_v8;
      _224_v8 = _dafny.Map.Empty.slice().updateUnsafe(_223_v7,_223_v7);
      let _225_v9;
      _225_v9 = _dafny.Map.Empty.slice().updateUnsafe(_216_v1,(((_224_v8).contains(false)) ? ((_224_v8).get(false)) : (_223_v7)));
      let _226_v10;
      _226_v10 = _dafny.Seq.of(_223_v7);
      let _227_v11;
      _227_v11 = _dafny.Map.Empty.slice().updateUnsafe(_226_v10,_223_v7);
      let _228_v12;
      _228_v12 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_216_v1, new BigNumber((_215_v0).length), _216_v1, _216_v1, _216_v1));
      let _229_v13;
      _229_v13 = _dafny.Seq.of(_216_v1, _216_v1);
      let _230_globalState;
      let _nw38 = new _module.GlobalState();
      _nw38.__ctor(_218_v3, _219_v4, false, new BigNumber(526), (_222_v6)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_222_v6).length))], true, true, new BigNumber(72), _225_v9, (_227_v11).Merge(_227_v11), _228_v12, false, _229_v13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(156)), ((_231_v2) => function (_232_i1) {
        return _231_v2;
      })(_217_v2)), true, new _dafny.CodePoint('w'.codePointAt(0)));
      _230_globalState = _nw38;
      let _233_v14;
      _233_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm0(_223_v7, _223_v7, _230_globalState)),_220_v5);
      _216_v1 = new BigNumber((((_233_v14).Merge(_233_v14)).update(_223_v7, _220_v5)).length);
      let _234_v15;
      let _init7 = ((_235_v1) => function (_236_i2) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_235_v1,new BigNumber(-637))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_235_v1,_235_v1));
      })(_216_v1);
      let _nw39 = Array((new BigNumber(22)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw39.length); _i0_7++) {
        _nw39[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _234_v15 = _nw39;
      let _237_v16;
      _237_v16 = _dafny.Map.Empty.slice().updateUnsafe(_216_v1,_216_v1);
      let _238_v17;
      _238_v17 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),new BigNumber((_215_v0).length));
      let _239_v18;
      _239_v18 = _dafny.Map.Empty.slice().updateUnsafe((((_237_v16).contains(_216_v1)) ? ((_237_v16).get(_216_v1)) : (new BigNumber((_module.__default.fm1(_230_globalState)).length))),new BigNumber((_238_v17).length));
      let _index37 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_234_v15).length));
      (_234_v15)[_index37] = _239_v18;
      let _index38 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_234_v15).length));
      (_234_v15)[_index38] = _239_v18;
      _216_v1 = (_216_v1).multipliedBy(_216_v1);
      let _240_v19;
      let _nw40 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _240_v19 = _nw40;
      let _241_v20;
      _241_v20 = _dafny.Map.Empty.slice().updateUnsafe(_240_v19,_239_v18);
      let _242_v21;
      let _nw41 = Array((new BigNumber(13)).toNumber());
      _nw41[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_226_v10, _226_v10)).length);
      _nw41[(_dafny.ONE).toNumber()] = (_216_v1).minus(new BigNumber(948));
      _nw41[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_216_v1);
      _nw41[(new BigNumber(3)).toNumber()] = _216_v1;
      _nw41[(new BigNumber(4)).toNumber()] = _216_v1;
      _nw41[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_241_v20).Merge(_241_v20)).length));
      _nw41[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_223_v7, _223_v7)).length);
      _nw41[(new BigNumber(7)).toNumber()] = _216_v1;
      _nw41[(new BigNumber(8)).toNumber()] = _216_v1;
      _nw41[(new BigNumber(9)).toNumber()] = _216_v1;
      _nw41[(new BigNumber(10)).toNumber()] = _216_v1;
      _nw41[(new BigNumber(11)).toNumber()] = _216_v1;
      _nw41[(new BigNumber(12)).toNumber()] = _216_v1;
      _242_v21 = _nw41;
      _240_v19 = _240_v19;
      let _243_i3;
      _243_i3 = _dafny.ZERO;
      L0: {
        while ((_223_v7) && (_223_v7)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_243_i3)) {
              break L0;
            }
            _243_i3 = (_243_i3).plus(_dafny.ONE);
            _216_v1 = _216_v1;
            _module.__default.m0(_216_v1, _dafny.Seq.update(_215_v0, _module.__default.safeIndex(_216_v1, new BigNumber((_215_v0).length)), _217_v2), new BigNumber((_dafny.Seq.of(_216_v1, (_dafny.ZERO).minus((_dafny.ZERO).minus(_216_v1)))).length), (_216_v1).minus(new BigNumber(132)), _230_globalState);
            let _244_v22;
            _244_v22 = _dafny.Map.Empty.slice().updateUnsafe(_222_v6,_223_v7);
            _244_v22 = (_244_v22).update(_dafny.Seq.Concat(_222_v6, _222_v6), (_226_v10)[_module.__default.safeIndex(_216_v1, new BigNumber((_226_v10).length))]);
            _module.__default.m0(_216_v1, _215_v0, (_216_v1).plus(_216_v1), _216_v1, _230_globalState);
          }
        }
      }
      _module.__default.m0(_216_v1, _dafny.Seq.UnicodeFromString("casqliqs"), _module.__default.safeDivisionInt(_216_v1, new BigNumber((_215_v0).length)), new BigNumber((function () {
        let _coll22 = new _dafny.Map();
        for (const _compr_22 of _dafny.IntegerRange(new BigNumber(327), new BigNumber(229))) {
          let _245_v23 = _compr_22;
          if (((new BigNumber(327)).isLessThanOrEqualTo(_245_v23)) && ((_245_v23).isLessThan(new BigNumber(229)))) {
            _coll22.push([_module.__default.safeModuloInt(_245_v23, _216_v1),_223_v7]);
          }
        }
        return _coll22;
      }()).length), _230_globalState);
      _216_v1 = _216_v1;
      let _246_i4;
      _246_i4 = _dafny.ZERO;
      L1: {
        while (_223_v7) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_246_i4)) {
              break L1;
            }
            _246_i4 = (_246_i4).plus(_dafny.ONE);
            let _247_v24;
            _247_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wqs"),_216_v1);
            let _248_v26;
            _248_v26 = _dafny.Map.Empty.slice().updateUnsafe(_215_v0,_dafny.Set.fromElements(_223_v7));
            let _249_v28;
            _249_v28 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-396)), ((_250_v2) => function (_251_i5) {
              return _250_v2;
            })(_217_v2)));
            let _252_v29;
            _252_v29 = _dafny.Seq.of(function () {
              let _coll23 = new _dafny.Map();
              for (const _compr_23 of (_248_v26).Keys.Elements) {
                let _253_v25 = _compr_23;
                if ((_248_v26).contains(_253_v25)) {
                  _coll23.push([_253_v25,_216_v1]);
                }
              }
              return _coll23;
            }(), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rsb"),_216_v1), (function () {
              let _coll24 = new _dafny.Map();
              for (const _compr_24 of (_dafny.Seq.update(_249_v28, _module.__default.safeIndex((_dafny.ZERO).minus(_216_v1), new BigNumber((_249_v28).length)), _215_v0)).Elements) {
                let _254_v27 = _compr_24;
                if (_dafny.Seq.contains(_dafny.Seq.update(_249_v28, _module.__default.safeIndex((_dafny.ZERO).minus(_216_v1), new BigNumber((_249_v28).length)), _215_v0), _254_v27)) {
                  _coll24.push([_254_v27,_216_v1]);
                }
              }
              return _coll24;
            }()).update(_215_v0, new BigNumber((_215_v0).length)));
            _247_v24 = ((_252_v29)[_module.__default.safeIndex(_216_v1, new BigNumber((_252_v29).length))]).Merge(_247_v24);
            let _nw42 = Array((new BigNumber(22)).toNumber()).fill(false);
            _220_v5 = _nw42;
            _216_v1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_216_v1));
            (_230_globalState).f2 = (new BigNumber((_239_v18).length)).isLessThanOrEqualTo(_216_v1);
          }
        }
      }
      let _index39 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length));
      (_218_v3)[_index39] = _215_v0;
      let _255_v30;
      _255_v30 = _dafny.Map.Empty.slice().updateUnsafe(_223_v7,_217_v2);
      let _256_v31;
      _256_v31 = _dafny.Map.Empty.slice().updateUnsafe(!(_223_v7),_255_v30);
      let _index40 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
      (_240_v19)[_index40] = (((((_224_v8).contains(_223_v7)) ? ((_224_v8).get(_223_v7)) : (_223_v7))) ? (new BigNumber((_256_v31).length)) : (_216_v1));
      let _257_v32;
      _257_v32 = _dafny.MultiSet.fromElements(_217_v2, _217_v2, _module.__default.fm2(_230_globalState));
      let _index41 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length));
      let _index42 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
      let _rhs29 = (_module.D0.create_DC0(_module.__default.fm1(_230_globalState))).dtor_cf0;
      let _rhs30 = new BigNumber((_257_v32).cardinality());
      let _rhs31 = _220_v5;
      let _lhs26 = _218_v3;
      let _lhs27 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length));
      let _lhs28 = _240_v19;
      let _lhs29 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
      _lhs26[_lhs27] = _rhs29;
      _lhs28[_lhs29] = _rhs30;
      _220_v5 = _rhs31;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_220_v5).length))) {
        let _258_i6 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_258_i6)) && ((_258_i6).isLessThan(new BigNumber((_220_v5).length))))) {
          (_220_v5)[(_258_i6)] = (_dafny.MultiSet.fromElements(_223_v7)).IsDisjointFrom((_dafny.MultiSet.fromElements(_223_v7)).Intersect(_dafny.MultiSet.fromElements(_223_v7)));
        }
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_220_v5).length))) {
        let _259_i7 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_259_i7)) && ((_259_i7).isLessThan(new BigNumber((_220_v5).length))))) {
          (_220_v5)[(_259_i7)] = !(_223_v7);
        }
      }
      if (_223_v7) {
        let _260_v34;
        _260_v34 = _dafny.Map.Empty.slice().updateUnsafe(_223_v7,function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of _dafny.IntegerRange(new BigNumber(333), new BigNumber(468))) {
            let _261_v33 = _compr_25;
            if (((new BigNumber(333)).isLessThanOrEqualTo(_261_v33)) && ((_261_v33).isLessThan(new BigNumber(468)))) {
              _coll25.push([_module.__default.safeDivisionInt(_261_v33, _216_v1),_223_v7]);
            }
          }
          return _coll25;
        }());
        _216_v1 = ((_223_v7) ? (new BigNumber(((_260_v34).Merge(_260_v34)).length)) : (_module.__default.fm3(_230_globalState)));
        let _262_v35;
        _262_v35 = _module.D0.create_DC1(_module.__default.fm1(_230_globalState));
        let _source2 = _262_v35;
        if (_source2.is_DC1) {
          let _263___mcc_h0 = (_source2).cf1;
          let _264_cf1 = _263___mcc_h0;
          let _265_v36;
          _265_v36 = _dafny.MultiSet.fromElements(_223_v7, _223_v7, _223_v7);
          let _index43 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_219_v4).length));
          (_219_v4)[_index43] = _265_v36;
          let _index44 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_220_v5).length));
          (_220_v5)[_index44] = _223_v7;
          let _266_v37;
          _266_v37 = _dafny.Seq.of(_265_v36);
          let _267_v38;
          _267_v38 = _module.D1.create_DC3((_266_v37)[_module.__default.safeIndex(_216_v1, new BigNumber((_266_v37).length))]);
          let _268_v39;
          _268_v39 = _module.D1.create_DC4(_223_v7, _216_v1, _224_v8);
          let _pat_let_tv10 = _265_v36;
          let _index45 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_219_v4).length));
          let _index46 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_220_v5).length));
          let _rhs32 = (function (_pat_let6_0) {
            return function (_269_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_270_dt__update_hcf6_h0) {
                  return _module.D1.create_DC3(_270_dt__update_hcf6_h0);
                }(_pat_let7_0);
              }(_pat_let_tv10);
            }(_pat_let6_0);
          }(_267_v38)).dtor_cf6;
          let _rhs33 = (_268_v39).dtor_cf7;
          let _rhs34 = !(_223_v7);
          let _rhs35 = !(!(!(_223_v7)));
          let _lhs30 = _219_v4;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_219_v4).length));
          let _lhs32 = _230_globalState;
          let _lhs33 = _220_v5;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(635), new BigNumber((_220_v5).length));
          let _lhs35 = _230_globalState;
          _lhs30[_lhs31] = _rhs32;
          _lhs32.f2 = _rhs33;
          _lhs33[_lhs34] = _rhs34;
          _lhs35.f5 = _rhs35;
          _216_v1 = (new BigNumber((_215_v0).length)).minus(_216_v1);
          let _271_v41;
          _271_v41 = _dafny.Seq.of(function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(-26), new BigNumber(365))) {
              let _272_v40 = _compr_26;
              if (((new BigNumber(-26)).isLessThanOrEqualTo(_272_v40)) && ((_272_v40).isLessThan(new BigNumber(365)))) {
                _coll26.push([(_272_v40).multipliedBy(_216_v1),false]);
              }
            }
            return _coll26;
          }(), (_225_v9).Merge(_225_v9));
          _271_v41 = (((_220_v5)[_module.__default.safeIndex(new BigNumber(635), new BigNumber((_220_v5).length))]) ? (_dafny.Seq.update(_271_v41, _module.__default.safeIndex(_216_v1, new BigNumber((_271_v41).length)), _225_v9)) : (_dafny.Seq.Concat(_271_v41, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_216_v1,(_220_v5)[_module.__default.safeIndex(new BigNumber(635), new BigNumber((_220_v5).length))]), _225_v9, _225_v9))));
          _225_v9 = (_225_v9).update((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], !_dafny.areEqual(_264_cf1, _215_v0));
        } else if (_source2.is_DC2) {
          let _273___mcc_h1 = (_source2).cf2;
          let _274___mcc_h2 = (_source2).cf3;
          let _275___mcc_h3 = (_source2).cf4;
          let _276___mcc_h4 = (_source2).cf5;
          let _277_cf5 = _276___mcc_h4;
          let _278_cf4 = _275___mcc_h3;
          let _279_cf3 = _274___mcc_h2;
          let _280_cf2 = _273___mcc_h1;
          _277_cf5 = _220_v5;
          _module.__default.m0(_278_cf4, _dafny.Seq.Concat(_215_v0, (_218_v3)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length))]), _280_cf2, (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], _230_globalState);
          let _281_v42;
          _281_v42 = _dafny.MultiSet.fromElements(_223_v7);
          let _282_v43;
          _282_v43 = _dafny.Seq.of(_module.D1.create_DC3(_281_v42));
          let _index47 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_220_v5).length));
          (_220_v5)[_index47] = !_dafny.areEqual(_module.__default.fm4(_223_v7, new BigNumber(-367), _230_globalState), _282_v43);
          let _index48 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _index49 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_220_v5).length));
          let _index50 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _rhs36 = _278_cf4;
          let _rhs37 = _223_v7;
          let _rhs38 = _223_v7;
          let _rhs39 = _216_v1;
          let _rhs40 = _279_cf3;
          let _lhs36 = _240_v19;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _lhs38 = _220_v5;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_220_v5).length));
          let _lhs40 = _230_globalState;
          let _lhs41 = _240_v19;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          _lhs36[_lhs37] = _rhs36;
          _lhs38[_lhs39] = _rhs37;
          _lhs40.f5 = _rhs38;
          _279_cf3 = _rhs39;
          _lhs41[_lhs42] = _rhs40;
          (_230_globalState).f2 = (_220_v5)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_220_v5).length))];
        } else {
          let _283___mcc_h5 = (_source2).cf0;
          let _284_cf0 = _283___mcc_h5;
          let _285_v44;
          _285_v44 = _dafny.MultiSet.fromElements(_223_v7, false);
          let _index51 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_219_v4).length));
          (_219_v4)[_index51] = (_285_v44).update(_223_v7, _module.__default.abs(_216_v1));
          let _index52 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_219_v4).length));
          let _rhs41 = _285_v44;
          let _rhs42 = _223_v7;
          let _lhs43 = _219_v4;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_219_v4).length));
          let _lhs45 = _230_globalState;
          _lhs43[_lhs44] = _rhs41;
          _lhs45.f5 = _rhs42;
          _module.__default.m0((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], _dafny.Seq.UnicodeFromString("wxynt"), (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], _module.__default.safeModuloInt((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], new BigNumber(255)), _230_globalState);
          let _286_v45;
          _286_v45 = _dafny.MultiSet.fromElements(new BigNumber(((_219_v4)[_module.__default.safeIndex(new BigNumber(245), new BigNumber((_219_v4).length))]).cardinality()));
          _module.__default.m0(_216_v1, (_218_v3)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length))], new BigNumber((_286_v45).cardinality()), (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], _230_globalState);
          let _index53 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_220_v5).length));
          (_220_v5)[_index53] = _223_v7;
          let _index54 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_220_v5).length));
          (_220_v5)[_index54] = _dafny.areEqual(_284_cf0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-422)), ((_287_v2) => function (_288_i8) {
            return _287_v2;
          })(_217_v2)));
        }
        let _index55 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
        (_240_v19)[_index55] = _216_v1;
        let _289_v46;
        let _nw43 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
        _289_v46 = _nw43;
        let _290_v47;
        _290_v47 = _dafny.Seq.of(_dafny.Seq.of(_223_v7, _223_v7, _223_v7, _223_v7, _223_v7), _226_v10, _226_v10, _dafny.Seq.update(_226_v10, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_223_v7)).length), new BigNumber((_226_v10).length)), _223_v7), _dafny.Seq.of(false));
        let _index56 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_289_v46).length));
        (_289_v46)[_index56] = _290_v47;
        let _index57 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_289_v46).length));
        let _index58 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
        let _rhs43 = _290_v47;
        let _rhs44 = new BigNumber((_dafny.Seq.Concat(_229_v13, _229_v13)).length);
        let _rhs45 = _223_v7;
        let _lhs46 = _289_v46;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_289_v46).length));
        let _lhs48 = _240_v19;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
        let _lhs50 = _230_globalState;
        _lhs46[_lhs47] = _rhs43;
        _lhs48[_lhs49] = _rhs44;
        _lhs50.f5 = _rhs45;
        let _291_v48;
        let _nw44 = Array((new BigNumber(27)).toNumber()).fill([]);
        _291_v48 = _nw44;
        let _index59 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_291_v48).length));
        (_291_v48)[_index59] = _220_v5;
        let _index60 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_291_v48).length));
        (_291_v48)[_index60] = _220_v5;
      } else {
        let _292_v49;
        let _nw45 = new _module.C7();
        _nw45.__ctor(_242_v21, (!(_223_v7)) === (_223_v7), _dafny.Seq.update(_dafny.Seq.of((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], _216_v1), _module.__default.safeIndex(new BigNumber((_215_v0).length), new BigNumber((_dafny.Seq.of((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], _216_v1)).length)), new BigNumber((_dafny.Seq.UnicodeFromString("xse")).length)));
        _292_v49 = _nw45;
        let _293_v50;
        let _nw46 = new _module.C5();
        _nw46.__ctor(_223_v7, _229_v13);
        _293_v50 = _nw46;
        _216_v1 = new BigNumber(-693);
        let _index61 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
        let _rhs46 = !(_223_v7);
        let _rhs47 = (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))];
        let _lhs51 = _230_globalState;
        let _lhs52 = _240_v19;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
        _lhs51.f5 = _rhs46;
        _lhs52[_lhs53] = _rhs47;
        let _294_v51;
        let _nw47 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
        _294_v51 = _nw47;
        let _index62 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
        let _rhs48 = ((_module.__default.fm0(_223_v7, _223_v7, _230_globalState)) ? (_294_v51) : (_294_v51));
        let _rhs49 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(_223_v7, _223_v7, _223_v7, _223_v7, _223_v7), _dafny.Seq.of(_223_v7)), _226_v10);
        let _rhs50 = _module.__default.fm3(_230_globalState);
        let _lhs54 = _240_v19;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
        _294_v51 = _rhs48;
        _223_v7 = _rhs49;
        _lhs54[_lhs55] = _rhs50;
      }
      let _index63 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
      (_240_v19)[_index63] = (_dafny.ZERO).minus(new BigNumber((_226_v10).length));
      let _295_v52;
      _295_v52 = _dafny.MultiSet.fromElements(_223_v7);
      let _296_v53;
      let _nw48 = new _module.C7();
      _nw48.__ctor(_242_v21, ((_295_v52).update(_223_v7, _module.__default.abs(_216_v1))).equals(_295_v52), _dafny.Seq.Create(_module.__default.abs(new BigNumber(384)), ((_297_v19) => function (_298_i9) {
        return (_297_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_297_v19).length))];
      })(_240_v19)));
      _296_v53 = _nw48;
      let _299_v54;
      _299_v54 = _module.D3.create_DC10(_223_v7);
      _299_v54 = _299_v54;
      let _300_v55;
      _300_v55 = _dafny.Set.fromElements(_223_v7, !(_223_v7));
      let _hi1 = (((_295_v52).contains(_223_v7)) ? ((_295_v52).get(_223_v7)) : (new BigNumber((_300_v55).length)));
      for (let _301_i10 = _216_v1; _301_i10.isLessThan(_hi1); _301_i10 = _301_i10.plus(_dafny.ONE)) {
        if (_223_v7) {
          let _302_v56;
          let _nw49 = Array((new BigNumber(13)).toNumber()).fill(_module.D12.Default());
          _302_v56 = _nw49;
          let _303_v57;
          let _nw50 = new _module.C8();
          _nw50.__ctor(_215_v0, _300_v55, _223_v7, _dafny.Seq.of(new BigNumber(529), _216_v1));
          _303_v57 = _nw50;
          let _304_v58;
          _304_v58 = _dafny.Map.Empty.slice().updateUnsafe(_303_v57,_223_v7);
          let _305_v59;
          _305_v59 = _dafny.Map.Empty.slice().updateUnsafe(_304_v58,_302_v56);
          let _306_v60;
          let _nw51 = Array((new BigNumber(8)).toNumber());
          _nw51[(_dafny.ZERO).toNumber()] = _302_v56;
          _nw51[(_dafny.ONE).toNumber()] = _302_v56;
          _nw51[(new BigNumber(2)).toNumber()] = _302_v56;
          _nw51[(new BigNumber(3)).toNumber()] = (((_305_v59).contains(_dafny.Map.Empty.slice().updateUnsafe(_303_v57,_223_v7))) ? ((_305_v59).get(_dafny.Map.Empty.slice().updateUnsafe(_303_v57,_223_v7))) : (_302_v56));
          _nw51[(new BigNumber(4)).toNumber()] = _302_v56;
          _nw51[(new BigNumber(5)).toNumber()] = _302_v56;
          _nw51[(new BigNumber(6)).toNumber()] = _302_v56;
          _nw51[(new BigNumber(7)).toNumber()] = _302_v56;
          _306_v60 = _nw51;
          let _index64 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_306_v60).length));
          (_306_v60)[_index64] = _302_v56;
          let _307_v61;
          _307_v61 = _dafny.Map.Empty.slice().updateUnsafe(_223_v7,(_dafny.ZERO).minus((_dafny.ZERO).minus((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))])));
          let _308_v62;
          _308_v62 = _dafny.Map.Empty.slice().updateUnsafe((_216_v1).multipliedBy(_216_v1),_dafny.Seq.Concat(_226_v10, _226_v10));
          let _309_v63;
          _309_v63 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true), _295_v52, _295_v52);
          let _310_v64;
          let _nw52 = new _module.C1();
          _nw52.__ctor((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], _216_v1, _module.__default.fm0(_223_v7, _223_v7, _230_globalState), _229_v13, _216_v1);
          _310_v64 = _nw52;
          let _311_v65;
          let _nw53 = new _module.C1();
          _nw53.__ctor(new BigNumber((_307_v61).length), new BigNumber((_dafny.MultiSet.fromElements(_310_v64, _310_v64)).cardinality()), _223_v7, (_310_v64).f17, (_310_v64).f28);
          _311_v65 = _nw53;
          let _312_v66;
          _312_v66 = _dafny.Map.Empty.slice().updateUnsafe((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))],_311_v65);
          let _313_v67;
          let _nw54 = new _module.C3();
          _nw54.__ctor(_215_v0, (_module.D1.create_DC5(new BigNumber((_303_v57.f19).length), _311_v65.f16)).dtor_cf11, (_311_v65).f17);
          _313_v67 = _nw54;
          let _314_v68;
          _314_v68 = _dafny.Map.Empty.slice().updateUnsafe(_313_v67,new BigNumber((_module.__default.fm31(_310_v64.f16, _230_globalState)).length));
          let _index65 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_306_v60).length));
          let _index66 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _rhs51 = _302_v56;
          let _rhs52 = _dafny.Map.Empty.slice().updateUnsafe((_309_v63).contains(_295_v52),(_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))]);
          let _rhs53 = (_301_i10).minus(((_223_v7) ? ((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))]) : (new BigNumber((_312_v66).length))));
          let _rhs54 = _308_v62;
          let _rhs55 = (_module.__default.safeDivisionInt(new BigNumber((_229_v13).length), new BigNumber(((_295_v52).update(false, _module.__default.abs((((_314_v68).contains(_313_v67)) ? ((_314_v68).get(_313_v67)) : ((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))]))))).cardinality()))).multipliedBy(_module.__default.safeDivisionInt(_301_i10, _301_i10));
          let _lhs56 = _306_v60;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_306_v60).length));
          let _lhs58 = _240_v19;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _lhs60 = _240_v19;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          _lhs56[_lhs57] = _rhs51;
          _307_v61 = _rhs52;
          _lhs58[_lhs59] = _rhs53;
          _308_v62 = _rhs54;
          _lhs60[_lhs61] = _rhs55;
          let _315_v69;
          _315_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_217_v2,_296_v53),(_dafny.ZERO).minus((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))]));
          let _316_v70;
          _316_v70 = _dafny.Map.Empty.slice().updateUnsafe(_301_i10,_296_v53);
          let _rhs56 = _315_v69;
          let _rhs57 = !((_316_v70).Merge(_316_v70)).equals((_316_v70).Merge(_316_v70));
          _315_v69 = _rhs56;
          _223_v7 = _rhs57;
          let _317_v71;
          _317_v71 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), ((_318_v8) => function (_319_i11) {
            return new BigNumber((_318_v8).length);
          })(_224_v8)), (_311_v65).f17, (_311_v65).f17, (_311_v65).f17, (_310_v64).f17);
          let _320_v72;
          let _nw55 = new _module.C1();
          _nw55.__ctor(_216_v1, _216_v1, false, _dafny.Seq.Concat((_317_v71)[_module.__default.safeIndex((_310_v64).f28, new BigNumber((_317_v71).length))], _dafny.Seq.of(_216_v1, new BigNumber(427))), _301_i10);
          _320_v72 = _nw55;
          let _321_v75;
          _321_v75 = _dafny.Map.Empty.slice().updateUnsafe(_300_v55,(_module.__default.fm26(_230_globalState)).dtor_cf35);
          let _322_v76;
          _322_v76 = _dafny.Set.fromElements(new BigNumber((function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of (function () {
              let _coll28 = new _dafny.Map();
              for (const _compr_28 of (_321_v75).Keys.Elements) {
                let _323_v74 = _compr_28;
                if ((_321_v75).contains(_323_v74)) {
                  _coll28.push([_323_v74,_217_v2]);
                }
              }
              return _coll28;
            }()).Keys.Elements) {
              let _324_v73 = _compr_27;
              if ((function () {
                let _coll29 = new _dafny.Map();
                for (const _compr_29 of (_321_v75).Keys.Elements) {
                  let _323_v74 = _compr_29;
                  if ((_321_v75).contains(_323_v74)) {
                    _coll29.push([_323_v74,_217_v2]);
                  }
                }
                return _coll29;
              }()).contains(_324_v73)) {
                _coll27.push([_324_v73,_310_v64.f16]);
              }
            }
            return _coll27;
          }()).length), new BigNumber(-809));
          let _325_v77;
          _325_v77 = _dafny.Set.fromElements(_301_i10, new BigNumber((_322_v76).length), (_320_v72).f30, _module.__default.safeModuloInt((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], new BigNumber((_dafny.Seq.UnicodeFromString("d")).length)));
          _322_v76 = function () {
            let _coll30 = new _dafny.Set();
            for (const _compr_30 of _dafny.IntegerRange(new BigNumber(-708), new BigNumber(256))) {
              let _326_v78 = _compr_30;
              if (((new BigNumber(-708)).isLessThanOrEqualTo(_326_v78)) && ((_326_v78).isLessThan(new BigNumber(256)))) {
                _coll30.add((_326_v78).minus((_320_v72).f30));
              }
            }
            return _coll30;
          }();
          let _327_v79;
          let _nw56 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _327_v79 = _nw56;
          let _328_v81;
          _328_v81 = _dafny.Seq.of(_226_v10);
          let _index68 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _index69 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _rhs58 = (_325_v77).Intersect(function () {
            let _coll31 = new _dafny.Set();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(76), new BigNumber(-151))) {
              let _329_v80 = _compr_31;
              if (((new BigNumber(76)).isLessThanOrEqualTo(_329_v80)) && ((_329_v80).isLessThan(new BigNumber(-151)))) {
                _coll31.add((_329_v80).multipliedBy((_320_v72).f30));
              }
            }
            return _coll31;
          }());
          let _rhs59 = ((_320_v72).f30).plus((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))]);
          let _rhs60 = _dafny.Seq.Concat((_328_v81)[_module.__default.safeIndex(new BigNumber((_226_v10).length), new BigNumber((_328_v81).length))], _226_v10);
          let _rhs61 = _327_v79;
          let _rhs62 = _301_i10;
          let _lhs62 = _240_v19;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _lhs64 = _240_v19;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          _322_v76 = _rhs58;
          _lhs62[_lhs63] = _rhs59;
          _226_v10 = _rhs60;
          _327_v79 = _rhs61;
          _lhs64[_lhs65] = _rhs62;
        } else {
          _218_v3 = _218_v3;
          let _330_v82;
          let _nw57 = new _module.C3();
          _nw57.__ctor(_dafny.Seq.Concat((_218_v3)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length))], _module.__default.fm1(_230_globalState)), _223_v7, _229_v13);
          _330_v82 = _nw57;
          (_230_globalState).f5 = _223_v7;
          let _331_v83;
          let _nw58 = new _module.C0();
          _nw58.__ctor((_dafny.ZERO).minus((_dafny.ZERO).minus(_216_v1)), _220_v5);
          _331_v83 = _nw58;
          let _332_v84;
          let _nw59 = new _module.C3();
          _nw59.__ctor((_330_v82).f24, _223_v7, _dafny.Seq.Concat(_dafny.Seq.update(_229_v13, _module.__default.safeIndex((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], new BigNumber((_229_v13).length)), _216_v1), _229_v13));
          _332_v84 = _nw59;
          _332_v84 = _330_v82;
        }
        _223_v7 = true;
        let _nw60 = Array((new BigNumber(2)).toNumber());
        _nw60[(_dafny.ZERO).toNumber()] = true;
        _nw60[(_dafny.ONE).toNumber()] = _223_v7;
        _220_v5 = _nw60;
        if (_223_v7) {
          let _333_v85;
          let _nw61 = new _module.C5();
          _nw61.__ctor(_223_v7, _229_v13);
          _333_v85 = _nw61;
          let _334_v86;
          _334_v86 = _dafny.Map.Empty.slice().updateUnsafe(_333_v85,_301_i10);
          let _335_v87;
          _335_v87 = _dafny.MultiSet.fromElements(_334_v86);
          let _336_v88;
          _336_v88 = _dafny.Map.Empty.slice().updateUnsafe(_223_v7,(((_335_v87).contains(_dafny.Map.Empty.slice().updateUnsafe(_333_v85,_216_v1))) ? ((_335_v87).get(_dafny.Map.Empty.slice().updateUnsafe(_333_v85,_216_v1))) : (_301_i10)));
          _336_v88 = (_336_v88).update(false, _216_v1);
          let _337_v89;
          let _nw62 = new _module.C1();
          _nw62.__ctor(new BigNumber(55), new BigNumber(((_218_v3)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length))]).length), _223_v7, (_333_v85).f17, (_dafny.ZERO).minus((_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))]));
          _337_v89 = _nw62;
          let _338_v90;
          _338_v90 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_230_globalState),_337_v89);
          let _rhs63 = ((((_333_v85.f16) ? (!(_337_v89.f16)) : (_223_v7))) ? (_module.__default.fm0(_223_v7, _333_v85.f16, _230_globalState)) : (_333_v85.f16));
          let _rhs64 = (((_338_v90).contains(new _dafny.CodePoint('v'.codePointAt(0)))) ? ((_338_v90).get(new _dafny.CodePoint('v'.codePointAt(0)))) : (_337_v89));
          let _rhs65 = _295_v52;
          let _lhs66 = _230_globalState;
          _lhs66.f2 = _rhs63;
          _337_v89 = _rhs64;
          _295_v52 = _rhs65;
          let _339_v91;
          _339_v91 = _module.D11.create_DC29(_217_v2);
          _339_v91 = _339_v91;
          let _340_v92;
          _340_v92 = _dafny.Set.fromElements((_337_v89).f28, (_337_v89).f28);
          let _index70 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          (_240_v19)[_index70] = new BigNumber(((_340_v92).Intersect(_340_v92)).length);
          let _index71 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          (_240_v19)[_index71] = _module.__default.fm3(_230_globalState);
        } else {
          let _index72 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length));
          (_218_v3)[_index72] = (_218_v3)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length))];
          _224_v8 = (_224_v8).update(_223_v7, !(_223_v7));
          let _index73 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _rhs66 = (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))];
          let _rhs67 = _223_v7;
          let _rhs68 = (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))];
          let _rhs69 = _301_i10;
          let _lhs67 = _230_globalState;
          let _lhs68 = _240_v19;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          _216_v1 = _rhs66;
          _lhs67.f5 = _rhs67;
          _216_v1 = _rhs68;
          _lhs68[_lhs69] = _rhs69;
          let _341_v93;
          let _nw63 = new _module.C4();
          _nw63.__ctor(new BigNumber(578), _module.__default.fm0((((_224_v8).contains(_223_v7)) ? ((_224_v8).get(_223_v7)) : (_223_v7)), _223_v7, _230_globalState));
          _341_v93 = _nw63;
          let _342_v94;
          _342_v94 = _dafny.Seq.of(_226_v10);
          let _rhs70 = _dafny.areEqual(_226_v10, _dafny.Seq.Concat(_226_v10, (_342_v94)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_342_v94).length))]));
          let _rhs71 = _module.__default.fm2(_230_globalState);
          let _rhs72 = _341_v93;
          _223_v7 = _rhs70;
          _217_v2 = _rhs71;
          _341_v93 = _rhs72;
          let _343_v95;
          let _nw64 = new _module.C3();
          _nw64.__ctor((_218_v3)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_218_v3).length))], _223_v7, _dafny.Seq.of(new BigNumber(-573), (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))]));
          _343_v95 = _nw64;
          let _344_v96;
          let _init8 = ((_345_v2) => function (_346_i12) {
            return _345_v2;
          })(_217_v2);
          let _nw65 = Array((new BigNumber(8)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw65.length); _i0_8++) {
            _nw65[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _344_v96 = _nw65;
          let _index74 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_344_v96).length));
          (_344_v96)[_index74] = _217_v2;
          let _347_v97;
          let _nw66 = new _module.C8();
          _nw66.__ctor(_215_v0, _300_v55, _223_v7, _dafny.Seq.of((_229_v13)[_module.__default.safeIndex(new BigNumber(814), new BigNumber((_229_v13).length))], new BigNumber((_module.__default.fm17(_230_globalState)).length), (_240_v19)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length))], (_229_v13)[_module.__default.safeIndex(_301_i10, new BigNumber((_229_v13).length))]));
          _347_v97 = _nw66;
          let _348_v98;
          _348_v98 = _dafny.Map.Empty.slice().updateUnsafe(_347_v97,_223_v7);
          let _index75 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _index76 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_344_v96).length));
          let _rhs73 = (((_348_v98).contains(_347_v97)) ? ((_348_v98).get(_347_v97)) : (_223_v7));
          let _rhs74 = (_341_v93).f22;
          let _rhs75 = _343_v95;
          let _rhs76 = new _dafny.CodePoint('i'.codePointAt(0));
          let _lhs70 = _341_v93;
          let _lhs71 = _240_v19;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_240_v19).length));
          let _lhs73 = _344_v96;
          let _lhs74 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_344_v96).length));
          _lhs70.f23 = _rhs73;
          _lhs71[_lhs72] = _rhs74;
          _343_v95 = _rhs75;
          _lhs73[_lhs74] = _rhs76;
        }
      }
      process.stdout.write((_215_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_216_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_217_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_218_v3)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_219_v4)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_222_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_223_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_225_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(879),true).updateUnsafe(new BigNumber(3),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_226_v10, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_227_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_228_v12, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(879), new BigNumber(879), new BigNumber(879), new BigNumber(879), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_229_v13, _dafny.Seq.of(new BigNumber(879), new BigNumber(879)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_230_globalState).f0)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_230_globalState).f1)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_230_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f4)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_230_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_globalState.f8).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(318),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_globalState).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_230_globalState).f10, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(879), new BigNumber(879), new BigNumber(879), new BigNumber(879), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_230_globalState).f12, _dafny.Seq.of(new BigNumber(879), new BigNumber(879)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_230_globalState).f13, _dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_230_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_233_v14).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_v15)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v18).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_v19)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_241_v20).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v21)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_243_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v30).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v31).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_257_v32).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v52).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v53).f20)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_299_v54).dtor_cf19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_300_v55).equals(_dafny.Set.fromElements(true, false))));
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
        return "D0.DC1" + "(" + this.cf1.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + this.cf0.toVerbatimString(true) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.Seq.UnicodeFromString(""));
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
    static create_DC4(cf7, cf8, cf9) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC5(cf10, cf11) {
      let $dt = new D1(1);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D1(2);
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(3);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6";
      } else if (this.$tag === 3) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, _dafny.ZERO, _dafny.Map.Empty);
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
    static create_DC8(cf13, cf14, cf15, cf16, cf17) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14 && this.cf15 === other.cf15 && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.Map.Empty, [], false, false, _dafny.ZERO);
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
    static create_DC10(cf19) {
      let $dt = new D3(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC9(cf18) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false);
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
    static create_DC13(cf24, cf25, cf26) {
      let $dt = new D4(1);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
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
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
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
        return other.$tag === 0 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(false, false, _dafny.ZERO);
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
    static create_DC14(cf27) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf27) + ")";
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf29, cf30) {
      let $dt = new D6(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC17(cf31) {
      let $dt = new D6(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC18(cf32) {
      let $dt = new D6(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC15(cf28) {
      let $dt = new D6(3);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf29 === other.cf29 && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf32 === other.cf32;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf28 === other.cf28;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(false, false);
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
    static create_DC19(cf33) {
      let $dt = new D7(0);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf35, cf36, cf37, cf38, cf39) {
      let $dt = new D8(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC22(cf40, cf41, cf42) {
      let $dt = new D8(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC23() {
      let $dt = new D8(2);
      return $dt;
    }
    static create_DC20(cf34) {
      let $dt = new D8(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC24(cf43) {
      let $dt = new D8(4);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get is_DC24() { return this.$tag === 4; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC23";
      } else if (this.$tag === 3) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36 && this.cf37 === other.cf37 && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf40 === other.cf40 && this.cf41 === other.cf41 && this.cf42 === other.cf42;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(_dafny.ZERO, false, false, false, _dafny.ZERO);
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
    static create_DC26(cf45, cf46) {
      let $dt = new D9(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC27(cf47, cf48) {
      let $dt = new D9(1);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC25(cf44) {
      let $dt = new D9(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf44) + ")";
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
        return other.$tag === 1 && this.cf47 === other.cf47 && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC26(_dafny.ZERO, null);
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
    static create_DC28(cf49) {
      let $dt = new D10(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf49 === other.cf49;
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
    static create_DC30(cf51, cf52) {
      let $dt = new D11(0);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC31(cf53, cf54, cf55, cf56) {
      let $dt = new D11(1);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC29(cf50) {
      let $dt = new D11(2);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC32(cf57) {
      let $dt = new D11(3);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get is_DC32() { return this.$tag === 3; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54) && _dafny.areEqual(this.cf55, other.cf55) && this.cf56 === other.cf56;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf57, other.cf57);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC30(false, _dafny.ZERO);
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
    static create_DC34(cf59, cf60, cf61, cf62) {
      let $dt = new D12(0);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC35(cf63, cf64, cf65, cf66, cf67) {
      let $dt = new D12(1);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC33(cf58) {
      let $dt = new D12(2);
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC36(cf68) {
      let $dt = new D12(3);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get is_DC36() { return this.$tag === 3; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60) && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf63 === other.cf63 && this.cf64 === other.cf64 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66 && this.cf67 === other.cf67;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf58 === other.cf58;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf68, other.cf68);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC34([], _dafny.ZERO, [], new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC37(cf69) {
      let $dt = new D13(0);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC39(cf71, cf72) {
      let $dt = new D14(0);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC38(cf70) {
      let $dt = new D14(1);
      $dt.cf70 = cf70;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf70() { return this.cf70; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC39" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf70) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf71 === other.cf71 && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf70, other.cf70);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC39(false, _dafny.ZERO);
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
    static create_DC41(cf74, cf75, cf76, cf77) {
      let $dt = new D15(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC42(cf78, cf79, cf80) {
      let $dt = new D15(1);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC40(cf73) {
      let $dt = new D15(2);
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC43(cf81) {
      let $dt = new D15(3);
      $dt.cf81 = cf81;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get is_DC43() { return this.$tag === 3; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf81() { return this.cf81; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC41" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + this.cf77.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf81) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf74, other.cf74) && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf78, other.cf78) && this.cf79 === other.cf79 && _dafny.areEqual(this.cf80, other.cf80);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf81, other.cf81);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC41(_dafny.ZERO, _dafny.Seq.of(), _dafny.Seq.of(), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC45(cf83, cf84, cf85) {
      let $dt = new D16(0);
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC46(cf86) {
      let $dt = new D16(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC44(cf82) {
      let $dt = new D16(2);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC47(cf87) {
      let $dt = new D16(3);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get is_DC47() { return this.$tag === 3; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC46" + "(" + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC44" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC47" + "(" + _dafny.toString(this.cf87) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf83 === other.cf83 && _dafny.areEqual(this.cf84, other.cf84) && _dafny.areEqual(this.cf85, other.cf85);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf86, other.cf86);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf87, other.cf87);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC45(false, _dafny.ZERO, _dafny.ZERO);
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
      this.f2 = false;
      this.f5 = false;
      this.f8 = _dafny.Map.Empty;
      this.f15 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f0 = [];
      this._f1 = [];
      this._f3 = _dafny.ZERO;
      this._f4 = [];
      this._f6 = false;
      this._f7 = _dafny.ZERO;
      this._f9 = _dafny.Map.Empty;
      this._f10 = _dafny.Seq.of();
      this._f11 = false;
      this._f12 = _dafny.Seq.of();
      this._f13 = _dafny.Seq.UnicodeFromString("");
      this._f14 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f27 = [];
      this._f26 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f26, f27) {
      let _this = this;
      (_this)._f26 = f26;
      (_this).f27 = f27;
      return;
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f16 = false;
      this._f17 = _dafny.Seq.of();
      this._f28 = _dafny.ZERO;
      this._f29 = _dafny.ZERO;
      this._f30 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f29, f30, f16, f17, f28) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f28 = f28;
      return;
    }
    fm13(p0, globalState) {
      let _this = this;
      return new BigNumber(519);
    };
    fm14(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(((_this).f29).minus((_this).f29), new BigNumber((function () {
        let _coll32 = new _dafny.Set();
        for (const _compr_32 of _dafny.IntegerRange(new BigNumber(847), new BigNumber(900))) {
          let _349_v0 = _compr_32;
          if (((new BigNumber(847)).isLessThanOrEqualTo(_349_v0)) && ((_349_v0).isLessThan(new BigNumber(900)))) {
            _coll32.add((_349_v0).plus((_this).f29));
          }
        }
        return _coll32;
      }()).length));
    };
    fm15(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.D4.create_DC12(_this.f16, _this.f16, (_this).f28)).dtor_cf23;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let _350_v0;
      _350_v0 = new BigNumber(138);
      let _rhs77 = (((_this.f16) ? (p1) : (_this.f16))) || (p2);
      let _rhs78 = _module.__default.fm3(globalState);
      let _lhs75 = _this;
      _lhs75.f16 = _rhs77;
      _350_v0 = _rhs78;
      let _351_v1;
      _351_v1 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f29).minus(new BigNumber(253)),p1);
      _351_v1 = (_351_v1).update((_this).f30, _this.f16);
      let _352_v2;
      let _nw67 = Array((new BigNumber(9)).toNumber()).fill(false);
      _352_v2 = _nw67;
      let _index77 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_352_v2).length));
      (_352_v2)[_index77] = _this.f16;
      let _353_v3;
      _353_v3 = new _dafny.CodePoint('r'.codePointAt(0));
      let _354_v4;
      _354_v4 = _dafny.Seq.of(((_dafny.ZERO).minus((_this).fm14((_this).f28, new BigNumber((p0).length), _353_v3, globalState))).minus(_350_v0));
      let _index78 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_352_v2).length));
      let _rhs79 = !(p1);
      let _rhs80 = _this.f16;
      let _rhs81 = !(_350_v0).isEqualTo((_this).f29);
      let _rhs82 = (_this).f17;
      let _lhs76 = globalState;
      let _lhs77 = globalState;
      let _lhs78 = _352_v2;
      let _lhs79 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_352_v2).length));
      _lhs76.f2 = _rhs79;
      _lhs77.f5 = _rhs80;
      _lhs78[_lhs79] = _rhs81;
      _354_v4 = _rhs82;
      let _355_v5;
      let _nw68 = new _module.C0();
      _nw68.__ctor(new BigNumber((_module.__default.fm16((_this).f30, globalState)).cardinality()), _352_v2);
      _355_v5 = _nw68;
      let _356_v6;
      _356_v6 = _module.D0.create_DC0(p0);
      let _pat_let_tv11 = globalState;
      let _source3 = function (_pat_let8_0) {
        return function (_357_dt__update__tmp_h0) {
          return function (_pat_let9_0) {
            return function (_358_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_358_dt__update_hcf0_h0);
            }(_pat_let9_0);
          }(_module.__default.fm1(_pat_let_tv11));
        }(_pat_let8_0);
      }(_356_v6);
      if (_source3.is_DC1) {
        let _359___mcc_h0 = (_source3).cf1;
        let _360_cf1 = _359___mcc_h0;
        let _361_v7;
        let _362_v8;
        let _363_v9;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_this).m12(globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _361_v7 = _out0;
        _362_v8 = _out1;
        _363_v9 = _out2;
        let _364_v10;
        _364_v10 = _dafny.Set.fromElements(new BigNumber((_360_cf1).length));
        let _365_v11;
        let _nw69 = new _module.C0();
        _nw69.__ctor(_module.__default.safeDivisionInt(_350_v0, new BigNumber((_364_v10).length)), _352_v2);
        _365_v11 = _nw69;
        (globalState).f2 = _363_v9;
        _350_v0 = _module.__default.safeModuloInt(((_355_v5).f26).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("mtmijpy")).length)), p3);
      } else if (_source3.is_DC2) {
        let _366___mcc_h1 = (_source3).cf2;
        let _367___mcc_h2 = (_source3).cf3;
        let _368___mcc_h3 = (_source3).cf4;
        let _369___mcc_h4 = (_source3).cf5;
        let _370_cf5 = _369___mcc_h4;
        let _371_cf4 = _368___mcc_h3;
        let _372_cf3 = _367___mcc_h2;
        let _373_cf2 = _366___mcc_h1;
        let _374_v12;
        let _nw70 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _374_v12 = _nw70;
        let _index79 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_374_v12).length));
        (_374_v12)[_index79] = _350_v0;
        let _index80 = _module.__default.safeIndex(new BigNumber(871), new BigNumber((_374_v12).length));
        (_374_v12)[_index80] = (_dafny.ZERO).minus((_this).f28);
        let _375_v13;
        _375_v13 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(691)), ((_376_v3) => function (_377_i0) {
          return _376_v3;
        })(_353_v3)), p0);
        (globalState).f5 = _module.__default.fm0((_375_v13).equals(_dafny.Set.fromElements(p0, p0)), p2, globalState);
        let _378_v14;
        let _nw71 = new _module.C0();
        _nw71.__ctor((_355_v5).f26, ((p2) ? (_370_cf5) : (_355_v5.f27)));
        _378_v14 = _nw71;
        _371_cf4 = (_378_v14).f26;
      } else {
        let _379___mcc_h5 = (_source3).cf0;
        let _380_cf0 = _379___mcc_h5;
        let _381_v15;
        _381_v15 = _dafny.Map.Empty.slice().updateUnsafe((_355_v5).f26,_352_v2);
        let _382_v16;
        _382_v16 = _dafny.Map.Empty.slice().updateUnsafe(_350_v0,(((_381_v15).contains((_this).f30)) ? ((_381_v15).get((_this).f30)) : (_355_v5.f27)));
        _382_v16 = (_382_v16).update(new BigNumber(297), _355_v5.f27);
        let _383_v17;
        _383_v17 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_352_v2)[_module.__default.safeIndex(new BigNumber(135), new BigNumber((_352_v2).length))]);
        let _384_v18;
        _384_v18 = _module.D1.create_DC4(p2, (_355_v5).f26, (_383_v17).Merge(_383_v17));
        let _385_v19;
        _385_v19 = _dafny.Map.Empty.slice().updateUnsafe(_353_v3,new BigNumber(((_this).f17).length));
        let _386_v20;
        _386_v20 = _dafny.Map.Empty.slice().updateUnsafe(p2,_385_v19);
        _384_v18 = _module.D1.create_DC4(false, new BigNumber((_386_v20).length), _383_v17);
        if ((false) && (((_355_v5).f26).isEqualTo((_355_v5).f26))) {
          _350_v0 = _module.__default.safeDivisionInt((_this).f28, new BigNumber((function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of ((_this).f17).Elements) {
              let _387_v21 = _compr_33;
              if (_dafny.Seq.contains((_this).f17, _387_v21)) {
                _coll33.push([_module.__default.safeModuloInt(_387_v21, _350_v0),(_this).f28]);
              }
            }
            return _coll33;
          }()).length));
          let _388_v22;
          let _nw72 = new _module.C0();
          _nw72.__ctor(_350_v0, _352_v2);
          _388_v22 = _nw72;
          let _389_v23;
          _389_v23 = _dafny.Map.Empty.slice().updateUnsafe((_352_v2)[_module.__default.safeIndex(new BigNumber(135), new BigNumber((_352_v2).length))],new BigNumber(443));
          (globalState).f5 = !(_389_v23).contains(p2);
          let _390_v24;
          let _nw73 = Array((new BigNumber(9)).toNumber()).fill(_module.D2.Default());
          _390_v24 = _nw73;
          let _391_v25;
          let _init9 = ((_392_cf0) => function (_393_i1) {
            return _392_cf0;
          })(_380_cf0);
          let _nw74 = Array((new BigNumber(18)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw74.length); _i0_9++) {
            _nw74[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _391_v25 = _nw74;
          let _index81 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_391_v25).length));
          (_391_v25)[_index81] = _dafny.Seq.Concat(_380_cf0, p0);
          let _394_v26;
          _394_v26 = _390_v24;
          let _index82 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_391_v25).length));
          let _rhs83 = ((_this.f16) ? (p1) : (_this.f16));
          let _rhs84 = (_394_v26);
          let _rhs85 = _353_v3;
          let _rhs86 = _dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_module.__default.fm17(globalState), _354_v4)).length), new BigNumber((p0).length)), _353_v3);
          let _rhs87 = p1;
          let _lhs80 = globalState;
          let _lhs81 = globalState;
          let _lhs82 = _391_v25;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_391_v25).length));
          let _lhs84 = _this;
          _lhs80.f5 = _rhs83;
          _390_v24 = _rhs84;
          _lhs81.f15 = _rhs85;
          _lhs82[_lhs83] = _rhs86;
          _lhs84.f16 = _rhs87;
          let _395_v27;
          _395_v27 = _dafny.Seq.of(_351_v1, _351_v1, _351_v1, _351_v1);
          let _396_v28;
          let _nw75 = Array((_dafny.ONE).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = (_395_v27)[_module.__default.safeIndex((_388_v22).f26, new BigNumber((_395_v27).length))];
          _396_v28 = _nw75;
          let _rhs88 = (_module.__default.safeModuloInt(p3, new BigNumber(151))).isLessThan((_this).f30);
          let _rhs89 = (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_383_v17).length),(_352_v2)[_module.__default.safeIndex(new BigNumber(135), new BigNumber((_352_v2).length))])).Merge(function () {
            let _coll34 = new _dafny.Map();
            for (const _compr_34 of ((_this).f17).Elements) {
              let _397_v29 = _compr_34;
              if (_dafny.Seq.contains((_this).f17, _397_v29)) {
                _coll34.push([(_397_v29).plus((_355_v5).f26),p1]);
              }
            }
            return _coll34;
          }())).length)).isLessThan((_this).fm14(p3, (_355_v5).f26, _353_v3, globalState));
          let _rhs90 = _396_v28;
          let _lhs85 = globalState;
          let _lhs86 = globalState;
          _lhs85.f2 = _rhs88;
          _lhs86.f5 = _rhs89;
          _396_v28 = _rhs90;
        } else {
          (globalState).f2 = (((_this).f17)[_module.__default.safeIndex(_350_v0, new BigNumber(((_this).f17).length))]).isLessThan(((_355_v5).f26).multipliedBy((_this).f29));
          (globalState).f15 = _353_v3;
          let _398_v30;
          let _nw76 = new _module.C0();
          _nw76.__ctor((_this).fm13(p1, globalState), _355_v5.f27);
          _398_v30 = _nw76;
          let _399_v31;
          _399_v31 = _dafny.MultiSet.fromElements((_398_v30).f26);
          _399_v31 = (_dafny.MultiSet.FromArray((_this).f17)).Intersect(_399_v31);
          (globalState).f5 = _dafny.Seq.IsPrefixOf((_this).f17, _354_v4);
        }
        let _400_v32;
        let _nw77 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _400_v32 = _nw77;
        let _401_v33;
        _401_v33 = _dafny.MultiSet.fromElements(_this.f16, _module.__default.fm0(!(p1), _module.__default.fm0(p2, _this.f16, globalState), globalState), true);
        let _402_v34;
        _402_v34 = _dafny.Map.Empty.slice().updateUnsafe(_400_v32,(_dafny.ZERO).minus(new BigNumber((_401_v33).cardinality())));
        _402_v34 = _402_v34;
      }
      let _403_v35;
      let _nw78 = Array((new BigNumber(10)).toNumber()).fill([]);
      _403_v35 = _nw78;
      let _index83 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_403_v35).length));
      (_403_v35)[_index83] = _352_v2;
      let _404_v36;
      _404_v36 = _dafny.Seq.of(_this.f16);
      let _pat_let_tv12 = _355_v5;
      let _pat_let_tv13 = _355_v5;
      let _pat_let_tv14 = _404_v36;
      let _pat_let_tv15 = p2;
      let _index84 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_403_v35).length));
      let _rhs91 = _dafny.Seq.Concat(_dafny.Seq.Concat(_404_v36, _404_v36), _404_v36);
      let _rhs92 = function (_source4) {
        if (_source4.is_DC12) {
          let _405___mcc_h6 = (_source4).cf21;
          let _406___mcc_h7 = (_source4).cf22;
          let _407___mcc_h8 = (_source4).cf23;
          let _408_cf23 = _407___mcc_h8;
          let _409_cf22 = _406___mcc_h7;
          let _410_cf21 = _405___mcc_h6;
          return (_pat_let_tv12).f26;
        } else if (_source4.is_DC13) {
          let _411___mcc_h9 = (_source4).cf24;
          let _412___mcc_h10 = (_source4).cf25;
          let _413___mcc_h11 = (_source4).cf26;
          let _414_cf26 = _413___mcc_h11;
          let _415_cf25 = _412___mcc_h10;
          let _416_cf24 = _411___mcc_h9;
          return _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv13).f26,_this.f16)).length), new BigNumber((_pat_let_tv14).length));
        } else {
          let _417___mcc_h12 = (_source4).cf20;
          let _418_cf20 = _417___mcc_h12;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_pat_let_tv15))).cardinality()));
        }
      }(_module.D4.create_DC13((_this).f28, (_this).f30, (_this).f30));
      let _rhs93 = (_this).f29;
      let _rhs94 = p3;
      let _rhs95 = _355_v5.f27;
      let _lhs87 = _403_v35;
      let _lhs88 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_403_v35).length));
      r0 = _rhs91;
      _350_v0 = _rhs92;
      _350_v0 = _rhs93;
      _350_v0 = _rhs94;
      _lhs87[_lhs88] = _rhs95;
      r0 = _dafny.Seq.of(p2);
      return r0;
    }
    m11(globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _hi2 = ((_this).f29).plus((_this).f30);
      for (let _419_i0 = (_this).f28; _419_i0.isLessThan(_hi2); _419_i0 = _419_i0.plus(_dafny.ONE)) {
        r2 = _module.__default.fm3(globalState);
        r0 = _module.__default.fm2(globalState);
        let _420_v0;
        let _nw79 = Array((new BigNumber(2)).toNumber()).fill(false);
        _420_v0 = _nw79;
        let _421_v1;
        let _nw80 = new _module.C0();
        _nw80.__ctor(_419_i0, _420_v0);
        _421_v1 = _nw80;
        let _422_v2;
        _422_v2 = _dafny.Seq.UnicodeFromString("anddf");
        let _423_v3;
        _423_v3 = _dafny.MultiSet.fromElements(_420_v0);
        let _424_v4;
        _424_v4 = _dafny.Map.Empty.slice().updateUnsafe(_422_v2,(_423_v3).Difference(_423_v3));
        _424_v4 = (_424_v4).update(((_this.f16) ? (_422_v2) : (_422_v2)), _423_v3);
      }
      let _hi3 = new BigNumber(((_this).f17).length);
      for (let _425_i1 = (_this).f28; _425_i1.isLessThan(_hi3); _425_i1 = _425_i1.plus(_dafny.ONE)) {
        let _426_v5;
        let _nw81 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _426_v5 = _nw81;
        let _427_v6;
        _427_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f17,(_this).f30);
        let _index85 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
        (_426_v5)[_index85] = (((_427_v6).contains((_this).f17)) ? ((_427_v6).get((_this).f17)) : (_425_i1));
        let _index86 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_426_v5).length));
        (_426_v5)[_index86] = new BigNumber(759);
        let _428_v7;
        _428_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber(13));
        let _429_v8;
        _429_v8 = _dafny.Seq.UnicodeFromString("lw");
        let _430_v9;
        _430_v9 = _dafny.MultiSet.fromElements(_428_v7, _428_v7, (_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_module.__default.fm3(globalState))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber((_429_v8).length))));
        let _431_v10;
        _431_v10 = _dafny.Set.fromElements(_this.f16);
        let _432_v11;
        _432_v11 = _dafny.MultiSet.fromElements((_this).f30, new BigNumber(393));
        let _index87 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
        let _index88 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_426_v5).length));
        let _rhs96 = new BigNumber((_430_v9).cardinality());
        let _rhs97 = new BigNumber((_429_v8).length);
        let _rhs98 = (!((_431_v10).IsSubsetOf(_431_v10))) && (!(_dafny.areEqual((_this).f17, (_this).f17)));
        let _rhs99 = _this.f16;
        let _rhs100 = (_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_432_v11)).contains(!(_this.f16) || (_module.__default.fm0(_this.f16, _this.f16, globalState)));
        let _lhs89 = _426_v5;
        let _lhs90 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
        let _lhs91 = _426_v5;
        let _lhs92 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_426_v5).length));
        let _lhs93 = globalState;
        let _lhs94 = globalState;
        let _lhs95 = globalState;
        _lhs89[_lhs90] = _rhs96;
        _lhs91[_lhs92] = _rhs97;
        _lhs93.f2 = _rhs98;
        _lhs94.f5 = _rhs99;
        _lhs95.f5 = _rhs100;
        let _433_v12;
        _433_v12 = _module.D1.create_DC4(_this.f16, (new BigNumber(43)).plus((_this).f29), _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16));
        let _source5 = _433_v12;
        if (_source5.is_DC4) {
          let _434___mcc_h0 = (_source5).cf7;
          let _435___mcc_h1 = (_source5).cf8;
          let _436___mcc_h2 = (_source5).cf9;
          let _437_cf9 = _436___mcc_h2;
          let _438_cf8 = _435___mcc_h1;
          let _439_cf7 = _434___mcc_h0;
          let _440_v13;
          _440_v13 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("hndbi"));
          let _index89 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
          (_426_v5)[_index89] = (_425_i1).plus(_module.__default.safeDivisionInt(new BigNumber((_440_v13).length), (_this).f30));
          let _441_v14;
          _441_v14 = _dafny.Seq.of(false);
          _438_cf8 = (new BigNumber((_441_v14).length)).plus((_this).f29);
          let _442_v15;
          _442_v15 = _dafny.Seq.of(_426_v5);
          let _index90 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
          (_426_v5)[_index90] = new BigNumber((_dafny.Seq.Concat(_442_v15, _442_v15)).length);
          let _443_v16;
          let _init10 = function (_444_i2) {
            return false;
          };
          let _nw82 = Array((new BigNumber(18)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw82.length); _i0_10++) {
            _nw82[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _443_v16 = _nw82;
          let _445_v17;
          let _nw83 = new _module.C0();
          _nw83.__ctor((_426_v5)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length))], _443_v16);
          _445_v17 = _nw83;
        } else if (_source5.is_DC5) {
          let _446___mcc_h3 = (_source5).cf10;
          let _447___mcc_h4 = (_source5).cf11;
          let _448_cf11 = _447___mcc_h4;
          let _449_cf10 = _446___mcc_h3;
          let _450_v18;
          _450_v18 = _dafny.MultiSet.fromElements(_448_cf11, !(_this.f16));
          let _451_v19;
          _451_v19 = _module.D4.create_DC12(_this.f16, ((_448_cf11) ? (_this.f16) : (_448_cf11)), new BigNumber((_450_v18).cardinality()));
          _451_v19 = _451_v19;
          let _452_v20;
          let _nw84 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _452_v20 = _nw84;
          _452_v20 = _452_v20;
          (_this).f16 = _this.f16;
          let _453_v21;
          let _init11 = ((_454_cf11) => function (_455_i3) {
            return _454_cf11;
          })(_448_cf11);
          let _nw85 = Array((new BigNumber(23)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw85.length); _i0_11++) {
            _nw85[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _453_v21 = _nw85;
          let _index91 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_453_v21).length));
          (_453_v21)[_index91] = true;
          let _index92 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_453_v21).length));
          (_453_v21)[_index92] = _this.f16;
        } else if (_source5.is_DC6) {
          let _456_v22;
          let _nw86 = Array((new BigNumber(19)).toNumber()).fill(false);
          _456_v22 = _nw86;
          let _457_v23;
          let _nw87 = new _module.C0();
          _nw87.__ctor((_this).f29, _456_v22);
          _457_v23 = _nw87;
          let _458_v24;
          _458_v24 = _dafny.MultiSet.fromElements(_this.f16, _dafny.areEqual(_dafny.Seq.of((_this).f30), _dafny.Seq.update((_this).f17, _module.__default.safeIndex((_this).f29, new BigNumber(((_this).f17).length)), _425_i1)), _this.f16, _this.f16, true);
          let _459_v25;
          _459_v25 = _dafny.Set.fromElements((_this).f29, new BigNumber(((_428_v7).Merge(_428_v7)).length));
          let _index93 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
          let _rhs101 = new BigNumber((_dafny.Set.fromElements(_429_v8, _dafny.Seq.Concat(_429_v8, _module.__default.fm1(globalState)), _429_v8, _dafny.Seq.Concat(_429_v8, _429_v8))).length);
          let _rhs102 = (_458_v24).Difference(_dafny.MultiSet.fromElements(_this.f16));
          let _rhs103 = (_459_v25).Intersect(_459_v25);
          let _lhs96 = _426_v5;
          let _lhs97 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
          _lhs96[_lhs97] = _rhs101;
          _458_v24 = _rhs102;
          _459_v25 = _rhs103;
          let _460_v26;
          _460_v26 = _dafny.Map.Empty.slice().updateUnsafe((_426_v5)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length))],_this.f16);
          let _461_v27;
          let _nw88 = new _module.C0();
          _nw88.__ctor(_module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_460_v26).length))), (_this).f29), _457_v23.f27);
          _461_v27 = _nw88;
          let _462_v28;
          _462_v28 = new _dafny.CodePoint('y'.codePointAt(0));
          (globalState).f15 = _462_v28;
        } else {
          let _463___mcc_h5 = (_source5).cf6;
          let _464_cf6 = _463___mcc_h5;
          let _index94 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length));
          (_426_v5)[_index94] = new BigNumber((_429_v8).length);
          (globalState).f5 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f28,_426_v5)).length)).isLessThan((_this).f30);
          let _465_v29;
          let _466_v30;
          let _467_v31;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = (_this).m12(globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _465_v29 = _out3;
          _466_v30 = _out4;
          _467_v31 = _out5;
          let _468_v32;
          _468_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC3(_464_cf6),_dafny.Seq.Concat(_465_v29, _465_v29));
          let _469_v33;
          _469_v33 = _module.D1.create_DC3(_464_cf6);
          let _470_v34;
          _470_v34 = new _dafny.CodePoint('q'.codePointAt(0));
          let _rhs104 = _dafny.Seq.update(_dafny.Seq.update((((_468_v32).contains(_469_v33)) ? ((_468_v32).get(_469_v33)) : (_429_v8)), _module.__default.safeIndex((_this).f30, new BigNumber(((((_468_v32).contains(_469_v33)) ? ((_468_v32).get(_469_v33)) : (_429_v8))).length)), _470_v34), _module.__default.safeIndex((_this).f28, new BigNumber((_dafny.Seq.update((((_468_v32).contains(_469_v33)) ? ((_468_v32).get(_469_v33)) : (_429_v8)), _module.__default.safeIndex((_this).f30, new BigNumber(((((_468_v32).contains(_469_v33)) ? ((_468_v32).get(_469_v33)) : (_429_v8))).length)), _470_v34)).length)), new _dafny.CodePoint('i'.codePointAt(0)));
          let _rhs105 = _this.f16;
          let _rhs106 = _module.__default.fm0(_467_v31, _466_v30, globalState);
          let _rhs107 = _466_v30;
          let _rhs108 = _467_v31;
          let _lhs98 = globalState;
          let _lhs99 = globalState;
          _465_v29 = _rhs104;
          _lhs98.f2 = _rhs105;
          r1 = _rhs106;
          r1 = _rhs107;
          _lhs99.f5 = _rhs108;
        }
        let _471_v36;
        _471_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_425_i1);
        let _472_v38;
        _472_v38 = _dafny.Seq.of(_this.f16);
        let _473_v39;
        _473_v39 = _dafny.Seq.of(_472_v38);
        let _474_v40;
        _474_v40 = _dafny.Seq.of((_471_v36).update(new BigNumber(((_473_v39)[_module.__default.safeIndex(new BigNumber((_472_v38).length), new BigNumber((_473_v39).length))]).length), (_this).f30));
        let _475_v42;
        let _nw89 = Array((new BigNumber(12)).toNumber());
        _nw89[(_dafny.ZERO).toNumber()] = function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_dafny.Seq.update((_this).f17, _module.__default.safeIndex((_this).f29, new BigNumber(((_this).f17).length)), (_this).f29)).Elements) {
            let _476_v35 = _compr_35;
            if (_dafny.Seq.contains(_dafny.Seq.update((_this).f17, _module.__default.safeIndex((_this).f29, new BigNumber(((_this).f17).length)), (_this).f29), _476_v35)) {
              _coll35.push([_module.__default.safeModuloInt(_476_v35, _425_i1),(_426_v5)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length))]]);
            }
          }
          return _coll35;
        }();
        _nw89[(_dafny.ONE).toNumber()] = (_471_v36).Merge(_471_v36);
        _nw89[(new BigNumber(2)).toNumber()] = (_471_v36).Merge(_471_v36);
        _nw89[(new BigNumber(3)).toNumber()] = function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of ((_this).f17).Elements) {
            let _477_v37 = _compr_36;
            if (_dafny.Seq.contains((_this).f17, _477_v37)) {
              _coll36.push([_module.__default.safeModuloInt(_477_v37, (_426_v5)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((_426_v5).length))]),(_dafny.ZERO).minus((_this).f30)]);
            }
          }
          return _coll36;
        }();
        _nw89[(new BigNumber(4)).toNumber()] = _471_v36;
        _nw89[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_425_i1,(_this).f30);
        _nw89[(new BigNumber(6)).toNumber()] = _471_v36;
        _nw89[(new BigNumber(7)).toNumber()] = (_474_v40)[_module.__default.safeIndex((((_428_v7).contains(_this.f16)) ? ((_428_v7).get(_this.f16)) : (_425_i1)), new BigNumber((_474_v40).length))];
        _nw89[(new BigNumber(8)).toNumber()] = function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of (_432_v11).Elements) {
            let _478_v41 = _compr_37;
            if ((_432_v11).contains(_478_v41)) {
              _coll37.push([(_478_v41).plus((_this).f29),new BigNumber((_472_v38).length)]);
            }
          }
          return _coll37;
        }();
        _nw89[(new BigNumber(9)).toNumber()] = _471_v36;
        _nw89[(new BigNumber(10)).toNumber()] = _471_v36;
        _nw89[(new BigNumber(11)).toNumber()] = (_471_v36).update((_this).f29, (_this).f29);
        _475_v42 = _nw89;
        _475_v42 = _475_v42;
        r2 = _module.__default.safeModuloInt(new BigNumber(60), (_425_i1).multipliedBy(_425_i1));
      }
      let _479_v43;
      let _nw90 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _479_v43 = _nw90;
      let _index95 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length));
      (_479_v43)[_index95] = (_this).f29;
      let _index96 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length));
      let _rhs109 = (_this).f30;
      let _rhs110 = (_this).f28;
      let _lhs100 = _479_v43;
      let _lhs101 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length));
      _lhs100[_lhs101] = _rhs109;
      r2 = _rhs110;
      let _480_v44;
      let _nw91 = Array((new BigNumber(27)).toNumber()).fill(false);
      _480_v44 = _nw91;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_480_v44).length))) {
        let _481_i4 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_481_i4)) && ((_481_i4).isLessThan(new BigNumber((_480_v44).length))))) {
          (_480_v44)[(_481_i4)] = _this.f16;
        }
      }
      let _482_v45;
      _482_v45 = _module.D4.create_DC13((_479_v43)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length))], _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_this.f16),_480_v44)).length), (_this).f28), new BigNumber(-224));
      let _source6 = _482_v45;
      if (_source6.is_DC12) {
        let _483___mcc_h6 = (_source6).cf21;
        let _484___mcc_h7 = (_source6).cf22;
        let _485___mcc_h8 = (_source6).cf23;
        let _486_cf23 = _485___mcc_h8;
        let _487_cf22 = _484___mcc_h7;
        let _488_cf21 = _483___mcc_h6;
        let _489_v46;
        _489_v46 = _dafny.Seq.UnicodeFromString("wivplsm");
        let _490_v47;
        _490_v47 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("wugmj"), _dafny.Seq.UnicodeFromString("cgl"), _dafny.Seq.UnicodeFromString("kamk"), _489_v46, _489_v46);
        let _491_v48;
        _491_v48 = _dafny.MultiSet.fromElements(_this.f16, false);
        let _492_v49;
        _492_v49 = _dafny.Map.Empty.slice().updateUnsafe(_490_v47,((_this.f16) ? (_dafny.MultiSet.fromElements(!(false), _488_cf21, _this.f16)) : (_491_v48)));
        _492_v49 = (_492_v49).update(_dafny.Seq.Concat(_module.__default.fm18((_dafny.ZERO).minus(_486_cf23), globalState), _490_v47), _491_v48);
        (_this).f16 = _this.f16;
        if (_488_cf21) {
          let _493_v50;
          _493_v50 = _dafny.Set.fromElements(((_this).f28).plus(new BigNumber(748)), (_this).f28, new BigNumber(811));
          let _494_v51;
          let _nw92 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
          _494_v51 = _nw92;
          let _index97 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_494_v51).length));
          (_494_v51)[_index97] = _dafny.Seq.of((_this).fm13(_488_cf21, globalState));
          let _index98 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_494_v51).length));
          let _rhs111 = _493_v50;
          let _rhs112 = _dafny.Seq.Concat((_this).f17, (_this).f17);
          let _lhs102 = _494_v51;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_494_v51).length));
          _493_v50 = _rhs111;
          _lhs102[_lhs103] = _rhs112;
          let _index99 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_480_v44).length));
          (_480_v44)[_index99] = _this.f16;
          let _495_v52;
          let _init12 = ((_496_cf22) => function (_497_i5) {
            return ((_496_cf22) ? (new _dafny.CodePoint('f'.codePointAt(0))) : (new _dafny.CodePoint('d'.codePointAt(0))));
          })(_487_cf22);
          let _nw93 = Array((new BigNumber(20)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw93.length); _i0_12++) {
            _nw93[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _495_v52 = _nw93;
          let _index100 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_495_v52).length));
          (_495_v52)[_index100] = new _dafny.CodePoint('h'.codePointAt(0));
          let _498_v53;
          _498_v53 = new _dafny.CodePoint('q'.codePointAt(0));
          let _index101 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_480_v44).length));
          let _index102 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_495_v52).length));
          let _rhs113 = !((_this).f29).isEqualTo((_this).f29);
          let _rhs114 = !(((_this).f29).isLessThan(_module.__default.safeDivisionInt((_this).f29, (_this).f29)));
          let _rhs115 = _498_v53;
          let _lhs104 = _480_v44;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(58), new BigNumber((_480_v44).length));
          let _lhs106 = _495_v52;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_495_v52).length));
          _lhs104[_lhs105] = _rhs113;
          r1 = _rhs114;
          _lhs106[_lhs107] = _rhs115;
          _498_v53 = new _dafny.CodePoint('i'.codePointAt(0));
          _493_v50 = _493_v50;
          r1 = !((_dafny.ZERO).minus(_486_cf23)).isEqualTo(((_494_v51)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((_494_v51).length))])[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_494_v51)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((_494_v51).length))]).length))]);
        } else {
          let _index103 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length));
          (_479_v43)[_index103] = (_this).f29;
          let _499_v54;
          _499_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f29);
          let _500_v55;
          let _init13 = function (_501_i6) {
            return (_this).f17;
          };
          let _nw94 = Array((new BigNumber(4)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw94.length); _i0_13++) {
            _nw94[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _500_v55 = _nw94;
          let _502_v56;
          _502_v56 = _dafny.Seq.of(true);
          let _503_v57;
          _503_v57 = _dafny.Map.Empty.slice().updateUnsafe(_488_cf21,_module.D2.create_DC8(_499_v54, _500_v55, _487_cf22, _487_cf22, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_502_v56).length),true)).length)));
          let _504_v58;
          _504_v58 = _module.D2.create_DC8(_499_v54, _500_v55, _488_cf21, _module.__default.fm0(_this.f16, _this.f16, globalState), _486_cf23);
          _503_v57 = (_503_v57).update(_module.__default.fm0(_this.f16, _487_cf22, globalState), _504_v58);
          _489_v46 = _module.__default.fm1(globalState);
          r2 = (_this).fm13(_487_cf22, globalState);
          (globalState).f5 = _487_cf22;
        }
        if (_488_cf21) {
          let _505_v59;
          let _nw95 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
          _505_v59 = _nw95;
          let _index104 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_505_v59).length));
          (_505_v59)[_index104] = _dafny.Set.fromElements(_489_v46);
          let _index105 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_505_v59).length));
          (_505_v59)[_index105] = _dafny.Set.fromElements(_489_v46);
          _487_cf22 = _488_cf21;
          let _506_v60;
          _506_v60 = new _dafny.CodePoint('p'.codePointAt(0));
          let _507_v61;
          _507_v61 = _dafny.MultiSet.fromElements(_506_v60, new _dafny.CodePoint('s'.codePointAt(0)));
          let _508_v62;
          _508_v62 = _dafny.Seq.of(_507_v61, (_507_v61).update(_506_v60, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("sey")).length))), _dafny.MultiSet.fromElements(_506_v60, _506_v60, _506_v60, _506_v60, _506_v60), _507_v61);
          let _509_v63;
          _509_v63 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_508_v62).length),_486_cf23);
          let _rhs116 = _509_v63;
          let _rhs117 = _487_cf22;
          let _lhs108 = _this;
          _509_v63 = _rhs116;
          _lhs108.f16 = _rhs117;
          let _510_v64;
          _510_v64 = _module.D0.create_DC0(_489_v46);
          _510_v64 = _510_v64;
          _509_v63 = (_509_v63).update((_this).f28, (_this).f29);
        } else {
          let _511_v65;
          _511_v65 = _module.D1.create_DC3((_491_v48).Difference(_491_v48));
          _511_v65 = _511_v65;
          let _512_v66;
          let _nw96 = new _module.C0();
          _nw96.__ctor((_this).f29, (_module.D0.create_DC2(new BigNumber((_489_v46).length), (_this).f28, (_479_v43)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length))], _480_v44)).dtor_cf5);
          _512_v66 = _nw96;
          let _513_v67;
          _513_v67 = _dafny.Set.fromElements(_488_cf21, _487_cf22);
          let _514_v68;
          _514_v68 = _module.D2.create_DC7(_513_v67);
          let _515_v69;
          let _nw97 = Array((new BigNumber(14)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = _514_v68;
          _nw97[(_dafny.ONE).toNumber()] = _514_v68;
          _nw97[(new BigNumber(2)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(3)).toNumber()] = _module.D2.create_DC7(_513_v67);
          _nw97[(new BigNumber(4)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(5)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(6)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(7)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(8)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(9)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(10)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(11)).toNumber()] = _module.D2.create_DC7(_513_v67);
          _nw97[(new BigNumber(12)).toNumber()] = _514_v68;
          _nw97[(new BigNumber(13)).toNumber()] = _514_v68;
          _515_v69 = _nw97;
          let _516_v70;
          _516_v70 = new _dafny.CodePoint('a'.codePointAt(0));
          let _517_v71;
          _517_v71 = _dafny.Map.Empty.slice().updateUnsafe(_515_v69,_module.__default.safeDivisionInt((_512_v66).f26, (_this).fm14((_512_v66).f26, (_479_v43)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length))], _516_v70, globalState)));
          _517_v71 = (_517_v71).update(_515_v69, new BigNumber(((_this).f17).length));
          _488_cf21 = _487_cf22;
          let _518_v72;
          let _nw98 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _518_v72 = _nw98;
          let _519_v73;
          _519_v73 = _dafny.Map.Empty.slice().updateUnsafe(false,_488_cf21);
          let _index106 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_518_v72).length));
          (_518_v72)[_index106] = _519_v73;
          let _index107 = _module.__default.safeIndex(new BigNumber(997), new BigNumber((_518_v72).length));
          (_518_v72)[_index107] = _519_v73;
        }
      } else if (_source6.is_DC13) {
        let _520___mcc_h9 = (_source6).cf24;
        let _521___mcc_h10 = (_source6).cf25;
        let _522___mcc_h11 = (_source6).cf26;
        let _523_cf26 = _522___mcc_h11;
        let _524_cf25 = _521___mcc_h10;
        let _525_cf24 = _520___mcc_h9;
        let _index108 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length));
        (_479_v43)[_index108] = (_this).f28;
        let _526_v74;
        _526_v74 = _dafny.Seq.of(_this.f16);
        _526_v74 = _dafny.Seq.Concat(_526_v74, _526_v74);
        let _527_v75;
        _527_v75 = _dafny.Set.fromElements(_479_v43);
        _527_v75 = ((_527_v75).Intersect(_527_v75)).Difference(_527_v75);
        let _528_v76;
        _528_v76 = _dafny.Seq.of((_479_v43)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length))], new BigNumber((_527_v75).length), (_479_v43)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length))]);
        let _529_v77;
        _529_v77 = _dafny.Seq.UnicodeFromString("js");
        let _530_v78;
        _530_v78 = _dafny.MultiSet.fromElements(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("ixeqyo"), _529_v77), _this.f16, (_this.f16) && (_this.f16), _this.f16);
        let _index109 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length));
        let _rhs118 = _526_v74;
        let _rhs119 = _528_v76;
        let _rhs120 = (_dafny.MultiSet.FromArray(_526_v74)).Union(_530_v78);
        let _rhs121 = (_this).f30;
        let _lhs109 = _479_v43;
        let _lhs110 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length));
        _526_v74 = _rhs118;
        _528_v76 = _rhs119;
        _530_v78 = _rhs120;
        _lhs109[_lhs110] = _rhs121;
      } else {
        let _531___mcc_h12 = (_source6).cf20;
        let _532_cf20 = _531___mcc_h12;
        let _533_v79;
        let _nw99 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
        _533_v79 = _nw99;
        let _index110 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_533_v79).length));
        (_533_v79)[_index110] = (_this).f17;
        let _index111 = _module.__default.safeIndex(new BigNumber(239), new BigNumber((_533_v79).length));
        (_533_v79)[_index111] = _dafny.Seq.Concat((_this).f17, (_this).f17);
        let _534_v80;
        _534_v80 = _dafny.Seq.UnicodeFromString("lju");
        _534_v80 = _534_v80;
        let _535_v81;
        _535_v81 = _module.D0.create_DC0(_dafny.Seq.Concat(_534_v80, _534_v80));
        _535_v81 = _535_v81;
        let _536_v82;
        _536_v82 = _dafny.Set.fromElements(false, _module.__default.fm0(!(false), _this.f16, globalState), true, _this.f16, !((_this).f28).isEqualTo((_this).f28));
        _536_v82 = _536_v82;
      }
      r1 = ((_this.f16) ? (_this.f16) : (_this.f16));
      let _537_v83;
      _537_v83 = new _dafny.CodePoint('c'.codePointAt(0));
      r0 = _537_v83;
      r1 = _this.f16;
      r2 = ((_479_v43)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_479_v43).length))]).minus((_dafny.ZERO).minus(new BigNumber(-271)));
      return [r0, r1, r2];
    }
    m12(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = false;
      let _538_v0;
      _538_v0 = _dafny.Seq.of(_this.f16, _this.f16);
      let _539_v1;
      let _nw100 = Array((new BigNumber(5)).toNumber());
      _nw100[(_dafny.ZERO).toNumber()] = false;
      _nw100[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements(_this.f16)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_538_v0));
      _nw100[(new BigNumber(2)).toNumber()] = ((_this.f16) ? (_this.f16) : (!(true)));
      _nw100[(new BigNumber(3)).toNumber()] = ((_this).f29).isLessThanOrEqualTo((_this).f29);
      _nw100[(new BigNumber(4)).toNumber()] = !((new BigNumber(((_this).f17).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of((_this).f30)).length)));
      _539_v1 = _nw100;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_539_v1).length))) {
        let _540_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_540_i0)) && ((_540_i0).isLessThan(new BigNumber((_539_v1).length))))) {
          (_539_v1)[(_540_i0)] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ucfop"), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_this).f29,_this.f16),_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_this.f16)).cardinality())))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ucfop")).length)), new _dafny.CodePoint('b'.codePointAt(0))), _dafny.Seq.UnicodeFromString("lpiib"));
        }
      }
      let _hi4 = (_this).f30;
      for (let _541_i1 = ((_this).f28).multipliedBy(new BigNumber(566)); _541_i1.isLessThan(_hi4); _541_i1 = _541_i1.plus(_dafny.ONE)) {
        if (_module.__default.fm0(_dafny.Seq.IsProperPrefixOf(_538_v0, _538_v0), _this.f16, globalState)) {
          let _542_v2;
          _542_v2 = new BigNumber(460);
          _542_v2 = (_this).f30;
          let _543_v3;
          _543_v3 = _dafny.Seq.UnicodeFromString("g");
          let _544_v4;
          _544_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_543_v3).length),false);
          let _545_v5;
          _545_v5 = new _dafny.CodePoint('p'.codePointAt(0));
          let _546_v6;
          _546_v6 = _dafny.Map.Empty.slice().updateUnsafe(_544_v4,_545_v5);
          _546_v6 = (_546_v6).update(_544_v4, _545_v5);
          (globalState).f5 = false;
          _542_v2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f28, (_this).f28));
          let _547_v7;
          _547_v7 = _module.D1.create_DC4(_this.f16, _542_v2, _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16));
          let _rhs122 = (_547_v7).dtor_cf8;
          let _rhs123 = _dafny.Seq.UnicodeFromString("rdpvkliok");
          let _rhs124 = _542_v2;
          _542_v2 = _rhs122;
          _543_v3 = _rhs123;
          _542_v2 = _rhs124;
        } else {
          let _548_v8;
          _548_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f28)));
          let _549_v9;
          let _init14 = function (_550_i2) {
            return (_this).f17;
          };
          let _nw101 = Array((new BigNumber(23)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw101.length); _i0_14++) {
            _nw101[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _549_v9 = _nw101;
          let _pat_let_tv16 = _549_v9;
          _548_v8 = ((_548_v8).Merge(_548_v8)).Merge((function (_pat_let10_0) {
            return function (_551_dt__update__tmp_h0) {
              return function (_pat_let11_0) {
                return function (_552_dt__update_hcf14_h0) {
                  return function (_pat_let12_0) {
                    return function (_553_dt__update_hcf16_h0) {
                      return _module.D2.create_DC8((_551_dt__update__tmp_h0).dtor_cf13, _552_dt__update_hcf14_h0, (_551_dt__update__tmp_h0).dtor_cf15, _553_dt__update_hcf16_h0, (_551_dt__update__tmp_h0).dtor_cf17);
                    }(_pat_let12_0);
                  }(true);
                }(_pat_let11_0);
              }(_pat_let_tv16);
            }(_pat_let10_0);
          }(_module.D2.create_DC8(_548_v8, _549_v9, _this.f16, _this.f16, (_this).f29))).dtor_cf13);
          let _554_v10;
          _554_v10 = new BigNumber(704);
          let _555_v11;
          _555_v11 = new _dafny.CodePoint('k'.codePointAt(0));
          let _rhs125 = _541_i1;
          let _rhs126 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ufg"), _module.__default.safeIndex(_module.__default.safeDivisionInt(_541_i1, (_this).f30), new BigNumber((_dafny.Seq.UnicodeFromString("ufg")).length)), _555_v11);
          let _rhs127 = ((_module.__default.fm0(_this.f16, _this.f16, globalState)) ? (_this.f16) : (true));
          let _lhs111 = globalState;
          _554_v10 = _rhs125;
          r0 = _rhs126;
          _lhs111.f2 = _rhs127;
          _554_v10 = ((_this).f28).minus((_this).f28);
          let _556_v12;
          let _init15 = ((_557_v11) => function (_558_i3) {
            return _557_v11;
          })(_555_v11);
          let _nw102 = Array((new BigNumber(19)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw102.length); _i0_15++) {
            _nw102[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _556_v12 = _nw102;
          let _559_v13;
          _559_v13 = _dafny.Map.Empty.slice().updateUnsafe(true,_556_v12);
          let _560_v14;
          _560_v14 = _dafny.Seq.UnicodeFromString("fj");
          let _561_v15;
          _561_v15 = _dafny.MultiSet.fromElements(_554_v10, (_this).f30, new BigNumber((_560_v14).length));
          let _562_v16;
          let _nw103 = Array((new BigNumber(17)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt((_this).fm14((_this).f29, (_this).f29, new _dafny.CodePoint('h'.codePointAt(0)), globalState), new BigNumber(-34));
          _nw103[(_dafny.ONE).toNumber()] = (_this).f30;
          _nw103[(new BigNumber(2)).toNumber()] = _module.__default.fm3(globalState);
          _nw103[(new BigNumber(3)).toNumber()] = _541_i1;
          _nw103[(new BigNumber(4)).toNumber()] = _554_v10;
          _nw103[(new BigNumber(5)).toNumber()] = ((_this).f29).multipliedBy(new BigNumber((_548_v8).length));
          _nw103[(new BigNumber(6)).toNumber()] = (_this).f28;
          _nw103[(new BigNumber(7)).toNumber()] = new BigNumber(84);
          _nw103[(new BigNumber(8)).toNumber()] = (_this).f30;
          _nw103[(new BigNumber(9)).toNumber()] = (_this).f30;
          _nw103[(new BigNumber(10)).toNumber()] = (_this).f29;
          _nw103[(new BigNumber(11)).toNumber()] = (_this).f28;
          _nw103[(new BigNumber(12)).toNumber()] = new BigNumber((_559_v13).length);
          _nw103[(new BigNumber(13)).toNumber()] = new BigNumber(((_561_v15).Difference(_561_v15)).cardinality());
          _nw103[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_554_v10), new BigNumber(210));
          _nw103[(new BigNumber(15)).toNumber()] = new BigNumber((_560_v14).length);
          _nw103[(new BigNumber(16)).toNumber()] = (_this).f30;
          _562_v16 = _nw103;
          let _rhs128 = _562_v16;
          let _rhs129 = _this.f16;
          let _rhs130 = _this.f16;
          let _lhs112 = globalState;
          let _lhs113 = globalState;
          _562_v16 = _rhs128;
          _lhs112.f2 = _rhs129;
          _lhs113.f5 = _rhs130;
          let _index112 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_539_v1).length));
          (_539_v1)[_index112] = _module.__default.fm0(_this.f16, _this.f16, globalState);
          let _index113 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_539_v1).length));
          let _rhs131 = !(_this.f16);
          let _rhs132 = _this.f16;
          let _rhs133 = _module.__default.fm0(!((_this).f28).isEqualTo((_this).f30), ((_this.f16) ? (_this.f16) : (_this.f16)), globalState);
          let _rhs134 = (_538_v0)[_module.__default.safeIndex(((_this.f16) ? (_module.__default.fm3(globalState)) : ((_this).f28)), new BigNumber((_538_v0).length))];
          let _lhs114 = _539_v1;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(695), new BigNumber((_539_v1).length));
          let _lhs116 = globalState;
          r2 = _rhs131;
          _lhs114[_lhs115] = _rhs132;
          _lhs116.f5 = _rhs133;
          r2 = _rhs134;
        }
        (globalState).f2 = _this.f16;
        let _563_v17;
        let _nw104 = new _module.C0();
        _nw104.__ctor((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f30,_this.f16)).length)).multipliedBy((_this).f28), _539_v1);
        _563_v17 = _nw104;
        let _564_v18;
        _564_v18 = _dafny.Map.Empty.slice().updateUnsafe(_539_v1,_539_v1);
        let _565_v19;
        let _nw105 = Array((new BigNumber(12)).toNumber());
        _nw105[(_dafny.ZERO).toNumber()] = _563_v17.f27;
        _nw105[(_dafny.ONE).toNumber()] = _539_v1;
        _nw105[(new BigNumber(2)).toNumber()] = _563_v17.f27;
        _nw105[(new BigNumber(3)).toNumber()] = _539_v1;
        _nw105[(new BigNumber(4)).toNumber()] = _539_v1;
        _nw105[(new BigNumber(5)).toNumber()] = _563_v17.f27;
        _nw105[(new BigNumber(6)).toNumber()] = ((_this.f16) ? (_539_v1) : (_563_v17.f27));
        _nw105[(new BigNumber(7)).toNumber()] = _539_v1;
        _nw105[(new BigNumber(8)).toNumber()] = _563_v17.f27;
        _nw105[(new BigNumber(9)).toNumber()] = (((_564_v18).contains(_563_v17.f27)) ? ((_564_v18).get(_563_v17.f27)) : (_563_v17.f27));
        _nw105[(new BigNumber(10)).toNumber()] = _539_v1;
        _nw105[(new BigNumber(11)).toNumber()] = _539_v1;
        _565_v19 = _nw105;
        let _nw106 = Array((new BigNumber(27)).toNumber()).fill([]);
        _565_v19 = _nw106;
      }
      let _566_v20;
      _566_v20 = new BigNumber(354);
      _566_v20 = (_this).f28;
      let _567_v21;
      _567_v21 = _dafny.Set.fromElements(_539_v1);
      let _568_v22;
      _568_v22 = _dafny.Set.fromElements(_this.f16, _this.f16);
      let _569_v23;
      _569_v23 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_568_v22);
      r2 = (((_567_v21).IsProperSubsetOf(_567_v21)) ? (_this.f16) : ((_569_v23).equals(_569_v23)));
      _566_v20 = new BigNumber(497);
      let _570_v24;
      _570_v24 = _dafny.Seq.UnicodeFromString("ldcs");
      let _571_v25;
      _571_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_570_v24).length),_this.f16);
      (globalState).f8 = (_571_v25).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f30,!(_this.f16)));
      r0 = _module.__default.fm1(globalState);
      let _572_v26;
      _572_v26 = _dafny.MultiSet.fromElements((_this).f28);
      r1 = ((_572_v26).Difference(_module.__default.fm19((_this).f28, globalState))).IsDisjointFrom(_572_v26);
      r2 = _this.f16;
      return [r0, r1, r2];
    }
    m13(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _573_v0;
      let _init16 = ((_574_p2) => function (_575_i0) {
        return _574_p2;
      })(p2);
      let _nw107 = Array((new BigNumber(25)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw107.length); _i0_16++) {
        _nw107[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _573_v0 = _nw107;
      let _576_v1;
      let _nw108 = new _module.C0();
      _nw108.__ctor(new BigNumber(876), _573_v0);
      _576_v1 = _nw108;
      let _577_v2;
      let _nw109 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _577_v2 = _nw109;
      _577_v2 = _577_v2;
      let _578_i1;
      _578_i1 = _dafny.ZERO;
      L2: {
        while (((_this.f16) ? (!(p2)) : (((_this).f28).isLessThanOrEqualTo((_this).f30)))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_578_i1)) {
              break L2;
            }
            _578_i1 = (_578_i1).plus(_dafny.ONE);
            r0 = (_this).f28;
            r0 = new BigNumber(965);
            (globalState).f2 = ((_this.f16) ? (p2) : (_this.f16));
            let _579_v3;
            _579_v3 = _dafny.Set.fromElements(p1, _this.f16, false, p2, _this.f16);
            let _580_v4;
            _580_v4 = _module.D2.create_DC7((_dafny.Set.fromElements(_this.f16)).Union(_579_v3));
            _580_v4 = _580_v4;
          }
        }
      }
      let _581_v5;
      _581_v5 = _dafny.Set.fromElements(p1, p2);
      let _582_v6;
      _582_v6 = _dafny.Seq.UnicodeFromString("curdxgmyi");
      let _583_v7;
      _583_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_581_v5).length),new BigNumber((_582_v6).length));
      let _584_v8;
      _584_v8 = _dafny.Seq.of(_583_v7, _583_v7);
      let _585_v10;
      _585_v10 = new _dafny.CodePoint('p'.codePointAt(0));
      let _586_v11;
      _586_v11 = _dafny.Map.Empty.slice().updateUnsafe(_585_v10,p1);
      let _587_v13;
      _587_v13 = _dafny.MultiSet.fromElements((_584_v8)[_module.__default.safeIndex((_this).f28, new BigNumber((_584_v8).length))], (function () {
        let _coll38 = new _dafny.Map();
        for (const _compr_38 of _dafny.IntegerRange(new BigNumber(357), new BigNumber(697))) {
          let _588_v9 = _compr_38;
          if (((new BigNumber(357)).isLessThanOrEqualTo(_588_v9)) && ((_588_v9).isLessThan(new BigNumber(697)))) {
            _coll38.push([(_588_v9).plus(new BigNumber(196)),(_576_v1).f26]);
          }
        }
        return _coll38;
      }()).update((_this).f30, new BigNumber((_586_v11).length)), _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_576_v1).f26), function () {
        let _coll39 = new _dafny.Map();
        for (const _compr_39 of _dafny.IntegerRange(new BigNumber(983), new BigNumber(580))) {
          let _589_v12 = _compr_39;
          if (((new BigNumber(983)).isLessThanOrEqualTo(_589_v12)) && ((_589_v12).isLessThan(new BigNumber(580)))) {
            _coll39.push([(_589_v12).multipliedBy(new BigNumber(323)),(_576_v1).f26]);
          }
        }
        return _coll39;
      }(), _583_v7);
      (globalState).f5 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_584_v8, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f30,new BigNumber(-324)), _583_v7, _583_v7)))).IsDisjointFrom(_587_v13);
      let _590_v15;
      let _nw110 = Array((new BigNumber(17)).toNumber());
      _nw110[(_dafny.ZERO).toNumber()] = (_this).f28;
      _nw110[(_dafny.ONE).toNumber()] = (_this).f30;
      _nw110[(new BigNumber(2)).toNumber()] = (_this).f30;
      _nw110[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_576_v1).f26, (_this).f28);
      _nw110[(new BigNumber(4)).toNumber()] = (_576_v1).f26;
      _nw110[(new BigNumber(5)).toNumber()] = (_576_v1).f26;
      _nw110[(new BigNumber(6)).toNumber()] = ((_this).f28).plus((_dafny.ZERO).minus((_576_v1).f26));
      _nw110[(new BigNumber(7)).toNumber()] = ((p1) ? ((_576_v1).f26) : (new BigNumber(700)));
      _nw110[(new BigNumber(8)).toNumber()] = (_this).f30;
      _nw110[(new BigNumber(9)).toNumber()] = new BigNumber(-614);
      _nw110[(new BigNumber(10)).toNumber()] = (_this).f28;
      _nw110[(new BigNumber(11)).toNumber()] = (_this).fm13(_this.f16, globalState);
      _nw110[(new BigNumber(12)).toNumber()] = (_this).f30;
      _nw110[(new BigNumber(13)).toNumber()] = (_this).f28;
      _nw110[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((function () {
        let _coll40 = new _dafny.Set();
        for (const _compr_40 of _dafny.IntegerRange(new BigNumber(965), new BigNumber(-1))) {
          let _591_v14 = _compr_40;
          if (((new BigNumber(965)).isLessThanOrEqualTo(_591_v14)) && ((_591_v14).isLessThan(new BigNumber(-1)))) {
            _coll40.add((_591_v14).minus((_576_v1).f26));
          }
        }
        return _coll40;
      }()).length), new BigNumber(40));
      _nw110[(new BigNumber(15)).toNumber()] = (_this).f28;
      _nw110[(new BigNumber(16)).toNumber()] = (_this).f28;
      _590_v15 = _nw110;
      let _index114 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_590_v15).length));
      (_590_v15)[_index114] = ((_dafny.ZERO).minus(new BigNumber((_582_v6).length))).plus((_this).f29);
      let _592_v16;
      let _init17 = ((_593_v1) => function (_594_i2) {
        return _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_593_v1).f26));
      })(_576_v1);
      let _nw111 = Array((new BigNumber(2)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw111.length); _i0_17++) {
        _nw111[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _592_v16 = _nw111;
      let _595_v17;
      _595_v17 = _dafny.MultiSet.fromElements((_this).f29);
      let _596_v18;
      _596_v18 = _dafny.MultiSet.fromElements(_595_v17, _595_v17);
      let _index115 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_592_v16).length));
      (_592_v16)[_index115] = (_596_v18).Union(_596_v18);
      let _597_v19;
      let _nw112 = Array((new BigNumber(18)).toNumber()).fill([]);
      _597_v19 = _nw112;
      let _598_v20;
      let _nw113 = Array((_dafny.ONE).toNumber());
      _nw113[(_dafny.ZERO).toNumber()] = _585_v10;
      _598_v20 = _nw113;
      let _index116 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_597_v19).length));
      (_597_v19)[_index116] = _598_v20;
      let _599_v21;
      _599_v21 = _dafny.Seq.of(_595_v17, _595_v17);
      let _600_v22;
      _600_v22 = _dafny.Seq.of(p2);
      let _index117 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_590_v15).length));
      let _index118 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_592_v16).length));
      let _index119 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_597_v19).length));
      let _rhs135 = new BigNumber(-533);
      let _rhs136 = (_this).f29;
      let _rhs137 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_599_v21, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(463), (_dafny.ZERO).minus((_576_v1).f26)))))).Difference(_596_v18);
      let _rhs138 = new BigNumber((_dafny.Seq.Concat(_600_v22, _module.__default.fm20(_module.__default.fm0(p1, p1, globalState), _this.f16, p1, _module.__default.fm3(globalState), globalState))).length);
      let _rhs139 = ((p2) ? (_598_v20) : (_598_v20));
      let _lhs117 = _590_v15;
      let _lhs118 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_590_v15).length));
      let _lhs119 = _592_v16;
      let _lhs120 = _module.__default.safeIndex(new BigNumber(969), new BigNumber((_592_v16).length));
      let _lhs121 = _597_v19;
      let _lhs122 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_597_v19).length));
      _lhs117[_lhs118] = _rhs135;
      r1 = _rhs136;
      _lhs119[_lhs120] = _rhs137;
      r0 = _rhs138;
      _lhs121[_lhs122] = _rhs139;
      let _601_v23;
      _601_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(globalState),p1);
      let _602_v24;
      let _nw114 = new _module.C0();
      _nw114.__ctor(((p1) ? (new BigNumber((_601_v23).length)) : ((_576_v1).f26)), _576_v1.f27);
      _602_v24 = _nw114;
      r0 = ((!(p1)) ? (((_dafny.ZERO).minus((_this).f28)).plus((_this).f28)) : (new BigNumber(((_this).f17).length)));
      let _603_v25;
      _603_v25 = _dafny.MultiSet.fromElements(_600_v22);
      let _604_v26;
      _604_v26 = _dafny.MultiSet.fromElements((_this).f17);
      let _605_v27;
      _605_v27 = _dafny.Seq.of((_this).f17);
      r1 = new BigNumber(((_603_v25).update(_600_v22, _module.__default.abs((((_604_v26).contains((_605_v27)[_module.__default.safeIndex(((_this).f17)[_module.__default.safeIndex((_576_v1).f26, new BigNumber(((_this).f17).length))], new BigNumber((_605_v27).length))])) ? ((_604_v26).get((_605_v27)[_module.__default.safeIndex(((_this).f17)[_module.__default.safeIndex((_576_v1).f26, new BigNumber(((_this).f17).length))], new BigNumber((_605_v27).length))])) : ((_this).fm15((_this).f17, new BigNumber((_600_v22).length), new BigNumber((_582_v6).length), new _dafny.CodePoint('i'.codePointAt(0)), globalState)))))).cardinality());
      return [r0, r1];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f25 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f25) {
      let _this = this;
      (_this)._f25 = f25;
      return;
    }
    fm12(globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(_dafny.Seq.of(true, false))).Difference((_dafny.Set.fromElements(_dafny.Seq.of(true, true), _dafny.Seq.of(!(false), false), _dafny.Seq.of(true), _dafny.Seq.of(!(true), true, false, true, false), _dafny.Seq.of(false, true))).Difference(_dafny.Set.fromElements(_dafny.Seq.of(true, false))));
    };
    m9(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _606_v0;
      _606_v0 = _dafny.Set.fromElements((_this).f25);
      let _607_i0;
      _607_i0 = _dafny.ZERO;
      L3: {
        while (!((((_this).f25).multipliedBy(new BigNumber((_606_v0).length))).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_this).f25, new BigNumber(873))))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_607_i0)) {
              break L3;
            }
            _607_i0 = (_607_i0).plus(_dafny.ONE);
            let _608_v1;
            _608_v1 = _dafny.MultiSet.fromElements((new BigNumber(389)).multipliedBy((_this).f25), ((_this).f25).plus((_this).f25));
            let _609_v2;
            _609_v2 = new _dafny.CodePoint('r'.codePointAt(0));
            let _610_v3;
            _610_v3 = _dafny.MultiSet.fromElements(_609_v2, _609_v2, _609_v2);
            let _611_v4;
            _611_v4 = false;
            let _612_v5;
            _612_v5 = _dafny.Set.fromElements(_611_v4);
            _608_v1 = (_dafny.MultiSet.fromElements((_this).f25, new BigNumber(((_610_v3).update(new _dafny.CodePoint('c'.codePointAt(0)), _module.__default.abs(new BigNumber((_612_v5).length)))).cardinality()), (_this).f25)).Intersect((_608_v1).Difference(p2));
            let _613_v6;
            _613_v6 = _dafny.Map.Empty.slice().updateUnsafe(_611_v4,(_this).f25);
            let _614_v7;
            _614_v7 = _dafny.MultiSet.fromElements((_613_v6).update(_611_v4, (_this).f25));
            let _615_v8;
            _615_v8 = _module.D4.create_DC11((_614_v7).Difference(_614_v7));
            let _source7 = _615_v8;
            if (_source7.is_DC12) {
              let _616___mcc_h0 = (_source7).cf21;
              let _617___mcc_h1 = (_source7).cf22;
              let _618___mcc_h2 = (_source7).cf23;
              let _619_cf23 = _618___mcc_h2;
              let _620_cf22 = _617___mcc_h1;
              let _621_cf21 = _616___mcc_h0;
              let _622_v9;
              let _nw115 = Array((new BigNumber(5)).toNumber()).fill(false);
              _622_v9 = _nw115;
              let _623_v10;
              let _nw116 = new _module.C0();
              _nw116.__ctor(_module.__default.safeModuloInt((_this).f25, (_this).f25), _622_v9);
              _623_v10 = _nw116;
              (globalState).f5 = _611_v4;
              let _624_v11;
              _624_v11 = _dafny.Seq.UnicodeFromString("rlvehw");
              let _625_v12;
              _625_v12 = _dafny.Seq.of(false, _611_v4, _621_cf21, _620_cf22);
              let _rhs140 = (_625_v12)[_module.__default.safeIndex(_619_cf23, new BigNumber((_625_v12).length))];
              let _rhs141 = _dafny.Seq.Concat(_624_v11, _module.__default.fm1(globalState));
              let _lhs123 = globalState;
              _lhs123.f5 = _rhs140;
              _624_v11 = _rhs141;
              let _626_v13;
              let _init18 = function (_627_i1) {
                return (_627_i1).plus((_dafny.ZERO).minus((_this).f25));
              };
              let _nw117 = Array((new BigNumber(6)).toNumber());
              for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw117.length); _i0_18++) {
                _nw117[_i0_18] = _init18(new BigNumber(_i0_18));
              }
              _626_v13 = _nw117;
              let _628_v14;
              _628_v14 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_620_cf22), _625_v12);
              let _629_v15;
              _629_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_623_v10).f26).plus(new BigNumber((_628_v14).cardinality())),_626_v13);
              let _630_v16;
              _630_v16 = _dafny.Map.Empty.slice().updateUnsafe(_611_v4,p0);
              _626_v13 = (((_629_v15).contains((_dafny.ZERO).minus(((_this).f25).minus(_619_cf23)))) ? ((_629_v15).get((_dafny.ZERO).minus(((_this).f25).minus(_619_cf23)))) : ((((_630_v16).contains(false)) ? ((_630_v16).get(false)) : (p0))));
            } else if (_source7.is_DC13) {
              let _631___mcc_h3 = (_source7).cf24;
              let _632___mcc_h4 = (_source7).cf25;
              let _633___mcc_h5 = (_source7).cf26;
              let _634_cf26 = _633___mcc_h5;
              let _635_cf25 = _632___mcc_h4;
              let _636_cf24 = _631___mcc_h3;
              let _637_v17;
              let _nw118 = Array((new BigNumber(2)).toNumber());
              _nw118[(_dafny.ZERO).toNumber()] = _611_v4;
              _nw118[(_dafny.ONE).toNumber()] = true;
              _637_v17 = _nw118;
              let _638_v18;
              _638_v18 = _module.D0.create_DC2((_dafny.ZERO).minus((_this).f25), (_this).f25, _634_cf26, _637_v17);
              let _pat_let_tv17 = _635_cf25;
              let _pat_let_tv18 = _636_cf24;
              _635_cf25 = (function (_pat_let13_0) {
                return function (_639_dt__update__tmp_h0) {
                  return function (_pat_let14_0) {
                    return function (_640_dt__update_hcf4_h0) {
                      return function (_pat_let15_0) {
                        return function (_641_dt__update_hcf2_h0) {
                          return _module.D0.create_DC2(_641_dt__update_hcf2_h0, (_639_dt__update__tmp_h0).dtor_cf3, _640_dt__update_hcf4_h0, (_639_dt__update__tmp_h0).dtor_cf5);
                        }(_pat_let15_0);
                      }(_pat_let_tv18);
                    }(_pat_let14_0);
                  }(_pat_let_tv17);
                }(_pat_let13_0);
              }(_638_v18)).dtor_cf4;
              (globalState).f2 = _611_v4;
              let _642_v19;
              _642_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_634_cf26),true);
              (globalState).f5 = (((_642_v19).contains(p2)) ? ((_642_v19).get(p2)) : (!(_635_cf25).isEqualTo(new BigNumber(543))));
              let _643_v20;
              _643_v20 = _dafny.Seq.UnicodeFromString("rsh");
              let _644_v21;
              _644_v21 = _module.D4.create_DC12(_611_v4, true, new BigNumber((_643_v20).length));
              let _rhs142 = _609_v2;
              let _rhs143 = _644_v21;
              let _lhs124 = globalState;
              _lhs124.f15 = _rhs142;
              _644_v21 = _rhs143;
            } else {
              let _645___mcc_h6 = (_source7).cf20;
              let _646_cf20 = _645___mcc_h6;
              r0 = (_this).f25;
              let _index120 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((p0).length));
              (p0)[_index120] = (_this).f25;
              let _647_v22;
              _647_v22 = _dafny.Seq.of(_611_v4);
              let _index121 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((p0).length));
              (p0)[_index121] = (_dafny.ZERO).minus((_dafny.ZERO).minus(((_this).f25).minus((((_608_v1).contains((_this).f25)) ? ((_608_v1).get((_this).f25)) : (new BigNumber((_647_v22).length))))));
              let _648_v23;
              let _nw119 = Array((new BigNumber(23)).toNumber()).fill(_module.D0.Default());
              _648_v23 = _nw119;
              let _649_v24;
              _649_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((p0)[_module.__default.safeIndex(new BigNumber(372), new BigNumber((p0).length))]),(_this).f25);
              let _650_v25;
              _650_v25 = _dafny.Map.Empty.slice().updateUnsafe(_649_v24,_611_v4);
              let _651_v26;
              _651_v26 = _dafny.Map.Empty.slice().updateUnsafe(_648_v23,(((_650_v25).contains(_649_v24)) ? ((_650_v25).get(_649_v24)) : (_611_v4)));
              _651_v26 = (_651_v26).update(_648_v23, _611_v4);
              let _652_v27;
              let _nw120 = Array((new BigNumber(24)).toNumber()).fill([]);
              _652_v27 = _nw120;
              let _653_v28;
              _653_v28 = _dafny.Map.Empty.slice().updateUnsafe(_611_v4,_611_v4);
              let _654_v29;
              let _nw121 = Array((new BigNumber(28)).toNumber());
              _nw121[(_dafny.ZERO).toNumber()] = (((_653_v28).contains(_611_v4)) ? ((_653_v28).get(_611_v4)) : (_611_v4));
              _nw121[(_dafny.ONE).toNumber()] = _611_v4;
              _nw121[(new BigNumber(2)).toNumber()] = false;
              _nw121[(new BigNumber(3)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(4)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(5)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(6)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(7)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(8)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(9)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(10)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(11)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(12)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(13)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(14)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(15)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(16)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(17)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(18)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(19)).toNumber()] = !(!(_611_v4));
              _nw121[(new BigNumber(20)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(21)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(22)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(23)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(24)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(25)).toNumber()] = _611_v4;
              _nw121[(new BigNumber(26)).toNumber()] = true;
              _nw121[(new BigNumber(27)).toNumber()] = _611_v4;
              _654_v29 = _nw121;
              let _index122 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_652_v27).length));
              (_652_v27)[_index122] = _654_v29;
              let _index123 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_652_v27).length));
              (_652_v27)[_index123] = _654_v29;
            }
            let _655_v30;
            let _nw122 = Array((new BigNumber(4)).toNumber()).fill(false);
            _655_v30 = _nw122;
            let _index124 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length));
            (_655_v30)[_index124] = _611_v4;
            let _index125 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length));
            (_655_v30)[_index125] = _611_v4;
            if (!(!(_611_v4)) || ((_module.__default.fm3(globalState)).isLessThan((_this).f25))) {
              let _656_v31;
              _656_v31 = _module.D3.create_DC10((_655_v30)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length))]);
              (globalState).f5 = (_656_v31).dtor_cf19;
              (globalState).f2 = false;
              let _657_v32;
              _657_v32 = _dafny.Map.Empty.slice().updateUnsafe(_611_v4,_612_v5);
              _612_v5 = (((((_657_v32).contains((_655_v30)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length))])) ? ((_657_v32).get((_655_v30)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length))])) : (_612_v5))).Union(_612_v5)).Intersect(_dafny.Set.fromElements(false));
              (globalState).f5 = (_655_v30)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length))];
              let _index126 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length));
              (_655_v30)[_index126] = ((_this).f25).isEqualTo(((_dafny.ZERO).minus((_this).f25)).plus((_dafny.ZERO).minus(new BigNumber((p2).cardinality()))));
            } else {
              let _658_v34;
              _658_v34 = _dafny.Map.Empty.slice().updateUnsafe(_613_v6,(_this).f25);
              let _659_v35;
              let _nw123 = new _module.C0();
              _nw123.__ctor(new BigNumber((function () {
                let _coll41 = new _dafny.Map();
                for (const _compr_41 of (_658_v34).Keys.Elements) {
                  let _660_v33 = _compr_41;
                  if ((_658_v34).contains(_660_v33)) {
                    _coll41.push([_660_v33,(_this).f25]);
                  }
                }
                return _coll41;
              }()).length), ((_611_v4) ? (_655_v30) : (_655_v30)));
              _659_v35 = _nw123;
              let _index127 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length));
              let _rhs144 = (_659_v35).f26;
              let _rhs145 = (_655_v30)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length))];
              let _lhs125 = _655_v30;
              let _lhs126 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_655_v30).length));
              r0 = _rhs144;
              _lhs125[_lhs126] = _rhs145;
              r0 = ((_659_v35).f26).multipliedBy((_659_v35).f26);
              let _661_v36;
              _661_v36 = _dafny.Map.Empty.slice().updateUnsafe(_609_v2,_module.D1.create_DC6());
              let _662_v37;
              _662_v37 = _module.D1.create_DC6();
              _661_v36 = (_661_v36).update(_609_v2, _662_v37);
              _655_v30 = _659_v35.f27;
            }
          }
        }
      }
      let _663_v38;
      _663_v38 = false;
      let _664_i2;
      _664_i2 = _dafny.ZERO;
      L4: {
        while (_663_v38) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_664_i2)) {
              break L4;
            }
            _664_i2 = (_664_i2).plus(_dafny.ONE);
            let _665_v39;
            let _nw124 = Array((new BigNumber(22)).toNumber()).fill(_module.D0.Default());
            _665_v39 = _nw124;
            let _666_v40;
            let _nw125 = Array((new BigNumber(16)).toNumber()).fill(false);
            _666_v40 = _nw125;
            let _index128 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_665_v39).length));
            (_665_v39)[_index128] = _module.D0.create_DC2((_this).f25, (_this).f25, (_this).f25, _666_v40);
            let _667_v41;
            _667_v41 = _module.D0.create_DC2((_this).f25, (((p2).contains((_this).f25)) ? ((p2).get((_this).f25)) : ((_this).f25)), ((_this).f25).plus(new BigNumber(730)), _666_v40);
            let _index129 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_665_v39).length));
            let _rhs146 = _667_v41;
            let _rhs147 = new BigNumber(-984);
            let _lhs127 = _665_v39;
            let _lhs128 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_665_v39).length));
            _lhs127[_lhs128] = _rhs146;
            r0 = _rhs147;
            let _index130 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length));
            (p0)[_index130] = _module.__default.fm3(globalState);
            let _index131 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length));
            (p0)[_index131] = new BigNumber(-683);
            let _668_v42;
            _668_v42 = _dafny.MultiSet.fromElements(_663_v38);
            let _669_v43;
            _669_v43 = _module.D1.create_DC3(_668_v42);
            _669_v43 = _669_v43;
            if (_663_v38) {
              let _670_v44;
              let _init19 = ((_671_p0) => function (_672_i3) {
                return (_672_i3).plus((_671_p0)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_671_p0).length))]);
              })(p0);
              let _nw126 = Array((new BigNumber(23)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw126.length); _i0_19++) {
                _nw126[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _670_v44 = _nw126;
              _670_v44 = _670_v44;
              let _index132 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length));
              (p0)[_index132] = ((new BigNumber((p2).cardinality())).multipliedBy(new BigNumber(257))).plus((p0)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length))]);
              let _673_v45;
              let _nw127 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _673_v45 = _nw127;
              let _index133 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_673_v45).length));
              (_673_v45)[_index133] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(372)), function (_674_i4) {
                return new _dafny.CodePoint('o'.codePointAt(0));
              });
              let _675_v46;
              _675_v46 = _dafny.Seq.UnicodeFromString("hmr");
              let _index134 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_673_v45).length));
              (_673_v45)[_index134] = _675_v46;
              let _676_v47;
              _676_v47 = _dafny.Seq.of((_this).f25);
              let _677_v48;
              let _nw128 = new _module.C1();
              _nw128.__ctor((_this).f25, (p0)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length))], _663_v38, _676_v47, (p0)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length))]);
              _677_v48 = _nw128;
              let _678_v49;
              _678_v49 = _dafny.Map.Empty.slice().updateUnsafe(!(_663_v38),_677_v48);
              let _rhs148 = _663_v38;
              let _rhs149 = _678_v49;
              let _lhs129 = globalState;
              _lhs129.f2 = _rhs148;
              _678_v49 = _rhs149;
              r0 = ((_677_v48.f16) ? ((((_668_v42).contains(_677_v48.f16)) ? ((_668_v42).get(_677_v48.f16)) : (_module.__default.fm3(globalState)))) : (((p0)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length))]).multipliedBy((p0)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length))])));
            } else {
              _663_v38 = !(_663_v38);
              let _679_v50;
              _679_v50 = _dafny.Seq.of(!(_663_v38));
              let _index135 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length));
              (p0)[_index135] = _module.__default.safeDivisionInt((_this).f25, new BigNumber((_679_v50).length));
              let _680_v51;
              _680_v51 = _dafny.Map.Empty.slice().updateUnsafe(_663_v38,_663_v38);
              let _681_v52;
              _681_v52 = _module.D1.create_DC4(!(_663_v38), (_this).f25, _680_v51);
              let _682_v53;
              _682_v53 = _dafny.Map.Empty.slice().updateUnsafe(_681_v52,_606_v0);
              _682_v53 = (_682_v53).update(_681_v52, _dafny.Set.fromElements((_this).f25));
              let _683_v54;
              let _nw129 = Array((new BigNumber(24)).toNumber()).fill([]);
              _683_v54 = _nw129;
              let _684_v55;
              _684_v55 = _module.D6.create_DC15(_683_v54);
              _683_v54 = (_684_v55).dtor_cf28;
              let _index136 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((p0).length));
              (p0)[_index136] = (_this).f25;
            }
          }
        }
      }
      r0 = ((_663_v38) ? ((_this).f25) : (_module.__default.safeDivisionInt((_this).f25, (_this).f25)));
      let _685_v56;
      let _nw130 = Array((new BigNumber(13)).toNumber()).fill(false);
      _685_v56 = _nw130;
      let _686_v57;
      let _nw131 = new _module.C0();
      _nw131.__ctor((_this).f25, _685_v56);
      _686_v57 = _nw131;
      let _index137 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_685_v56).length));
      (_685_v56)[_index137] = _663_v38;
      let _index138 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_685_v56).length));
      (_685_v56)[_index138] = (((_this).f25).multipliedBy(new BigNumber(-768))).isLessThanOrEqualTo((_this).f25);
      let _687_v58;
      _687_v58 = new _dafny.CodePoint('a'.codePointAt(0));
      (globalState).f15 = _687_v58;
      let _688_v59;
      let _nw132 = Array((new BigNumber(20)).toNumber()).fill(_module.D2.Default());
      _688_v59 = _nw132;
      let _689_v60;
      _689_v60 = _dafny.Map.Empty.slice().updateUnsafe(_688_v59,(_685_v56)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_685_v56).length))]);
      let _690_v61;
      _690_v61 = _688_v59;
      r0 = new BigNumber((_dafny.Set.fromElements((_this).f25, new BigNumber(((_689_v60).Merge((_689_v60).update(_690_v61, (_685_v56)[_module.__default.safeIndex(new BigNumber(641), new BigNumber((_685_v56).length))]))).length))).length);
      return r0;
    }
    m10(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let _691_v0;
      _691_v0 = new _dafny.CodePoint('k'.codePointAt(0));
      (globalState).f15 = _691_v0;
      r0 = _module.__default.fm3(globalState);
      let _692_v1;
      _692_v1 = _module.D6.create_DC16(!(!(p0)), p0);
      r0 = function (_source8) {
        if (_source8.is_DC16) {
          let _693___mcc_h0 = (_source8).cf29;
          let _694___mcc_h1 = (_source8).cf30;
          let _695_cf30 = _694___mcc_h1;
          let _696_cf29 = _693___mcc_h0;
          return (_this).f25;
        } else if (_source8.is_DC17) {
          let _697___mcc_h2 = (_source8).cf31;
          let _698_cf31 = _697___mcc_h2;
          return (_this).f25;
        } else if (_source8.is_DC18) {
          let _699___mcc_h3 = (_source8).cf32;
          let _700_cf32 = _699___mcc_h3;
          return ((_this).f25).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), function (_701_i0) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          })).length));
        } else {
          let _702___mcc_h4 = (_source8).cf28;
          let _703_cf28 = _702___mcc_h4;
          return _module.__default.safeModuloInt(new BigNumber(584), (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f25)));
        }
      }(((p0) ? (_692_v1) : (_692_v1)));
      let _hi5 = (_this).f25;
      for (let _704_i1 = (_this).f25; _704_i1.isLessThan(_hi5); _704_i1 = _704_i1.plus(_dafny.ONE)) {
        let _705_v2;
        _705_v2 = _dafny.Seq.of(new BigNumber(459));
        let _706_v3;
        let _nw133 = new _module.C1();
        _nw133.__ctor((_this).f25, (_this).f25, p0, _705_v2, _704_i1);
        _706_v3 = _nw133;
        (globalState).f2 = false;
        let _707_v4;
        _707_v4 = _dafny.Seq.UnicodeFromString("vbabos");
        let _708_v5;
        _708_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p0)).length),_707_v4)).length));
        let _709_v6;
        let _nw134 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
        _709_v6 = _nw134;
        let _710_v7;
        _710_v7 = _dafny.Set.fromElements((_706_v3).f30, (_this).f25);
        let _711_v8;
        _711_v8 = _dafny.Map.Empty.slice().updateUnsafe(_706_v3,p0);
        let _712_v9;
        _712_v9 = _module.D1.create_DC5(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_710_v7).length),new BigNumber(99))).length), (((_711_v8).contains(_706_v3)) ? ((_711_v8).get(_706_v3)) : (p0)));
        let _713_v10;
        _713_v10 = _module.D2.create_DC8(_708_v5, _709_v6, p0, (_712_v9).dtor_cf11, (_dafny.ZERO).minus((_this).f25));
        let _714_v11;
        _714_v11 = _dafny.MultiSet.fromElements(((_706_v3).f30).plus((_706_v3).f29), (_706_v3).fm14((_this).f25, (_706_v3).f30, _691_v0, globalState), (_713_v10).dtor_cf17);
        _714_v11 = _714_v11;
        let _715_v12;
        _715_v12 = _module.D1.create_DC6();
        _715_v12 = _715_v12;
      }
      let _716_v13;
      _716_v13 = _dafny.Seq.UnicodeFromString("lskhib");
      _716_v13 = _dafny.Seq.Concat(_716_v13, _dafny.Seq.UnicodeFromString("otcfiuf"));
      r0 = (new BigNumber((_716_v13).length)).multipliedBy((_this).f25);
      r0 = (_this).f25;
      let _717_v14;
      let _nw135 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _717_v14 = _nw135;
      r1 = (((p0) && (!(p0))) ? (_717_v14) : (_717_v14));
      return [r0, r1];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f16 = false;
      this._f17 = _dafny.Seq.of();
      this._f24 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f24, f16, f17) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm11(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f24;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let _718_v1;
      let _init20 = function (_719_i0) {
        return function () {
          let _coll42 = new _dafny.Set();
          for (const _compr_42 of _dafny.IntegerRange(new BigNumber(757), new BigNumber(81))) {
            let _720_v0 = _compr_42;
            if (((new BigNumber(757)).isLessThanOrEqualTo(_720_v0)) && ((_720_v0).isLessThan(new BigNumber(81)))) {
              _coll42.add((_720_v0).plus(new BigNumber((_dafny.MultiSet.fromElements(_this.f16)).cardinality())));
            }
          }
          return _coll42;
        }();
      };
      let _nw136 = Array((new BigNumber(15)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw136.length); _i0_20++) {
        _nw136[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _718_v1 = _nw136;
      let _721_v2;
      _721_v2 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,_this.f16)).length));
      let _index139 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_718_v1).length));
      (_718_v1)[_index139] = _721_v2;
      let _index140 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_718_v1).length));
      let _rhs150 = _721_v2;
      let _rhs151 = _this.f16;
      let _rhs152 = (((_721_v2).IsProperSubsetOf(_721_v2)) ? ((p3).isLessThan(p3)) : (_this.f16));
      let _lhs130 = _718_v1;
      let _lhs131 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_718_v1).length));
      let _lhs132 = globalState;
      let _lhs133 = globalState;
      _lhs130[_lhs131] = _rhs150;
      _lhs132.f5 = _rhs151;
      _lhs133.f2 = _rhs152;
      (_this).f16 = p1;
      let _722_v3;
      let _nw137 = Array((_dafny.ONE).toNumber()).fill(false);
      _722_v3 = _nw137;
      let _723_v4;
      _723_v4 = _dafny.Set.fromElements(_722_v3, _722_v3);
      let _724_i1;
      _724_i1 = _dafny.ZERO;
      L5: {
        while (((_723_v4).Difference(_723_v4)).contains(_722_v3)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_724_i1)) {
              break L5;
            }
            _724_i1 = (_724_i1).plus(_dafny.ONE);
            let _725_v5;
            _725_v5 = _module.D3.create_DC9((_this).f17);
            let _726_v6;
            _726_v6 = _dafny.Map.Empty.slice().updateUnsafe(_722_v3,p3);
            let _727_v7;
            _727_v7 = _dafny.Map.Empty.slice().updateUnsafe(_725_v5,new BigNumber(((_726_v6).Merge(_726_v6)).length));
            _727_v7 = (_727_v7).update(_725_v5, p3);
            let _728_v8;
            let _nw138 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
            _728_v8 = _nw138;
            let _729_v9;
            _729_v9 = _dafny.Seq.of(_this.f16, p1);
            let _index141 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_728_v8).length));
            (_728_v8)[_index141] = new BigNumber((_729_v9).length);
            let _index142 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_728_v8).length));
            (_728_v8)[_index142] = p3;
            let _730_v10;
            _730_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(false));
            _730_v10 = (_730_v10).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-629)), function (_731_i2) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            }), _this.f16);
            let _index143 = _module.__default.safeIndex(new BigNumber(940), new BigNumber((_728_v8).length));
            (_728_v8)[_index143] = (_728_v8)[_module.__default.safeIndex(new BigNumber(940), new BigNumber((_728_v8).length))];
          }
        }
      }
      let _732_v11;
      _732_v11 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f24);
      let _733_v12;
      _733_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16);
      let _734_v13;
      _734_v13 = _dafny.Seq.of(_733_v12);
      let _735_v14;
      _735_v14 = _dafny.Map.Empty.slice().updateUnsafe(_732_v11,((_this.f16) ? (_734_v13) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), ((_736_v12) => function (_737_i3) {
        return _736_v12;
      })(_733_v12)))));
      _735_v14 = (_735_v14).update(_732_v11, _734_v13);
      let _738_v15;
      _738_v15 = new BigNumber(-688);
      _738_v15 = (_dafny.ZERO).minus((_738_v15).multipliedBy(_738_v15));
      let _index144 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_722_v3).length));
      (_722_v3)[_index144] = _this.f16;
      let _index145 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_722_v3).length));
      let _rhs153 = _this.f16;
      let _rhs154 = !(_module.__default.fm0(p1, _this.f16, globalState));
      let _lhs134 = globalState;
      let _lhs135 = _722_v3;
      let _lhs136 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_722_v3).length));
      _lhs134.f5 = _rhs153;
      _lhs135[_lhs136] = _rhs154;
      let _739_v16;
      _739_v16 = _dafny.Seq.of(p1);
      r0 = _dafny.Seq.Concat(_739_v16, _739_v16);
      return r0;
    }
    m8(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Map.Empty;
      let _740_v0;
      _740_v0 = new BigNumber(69);
      let _741_v1;
      _741_v1 = _dafny.Seq.of(_this.f16, !(_this.f16));
      let _742_v2;
      _742_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-279),_module.__default.safeDivisionInt(_740_v0, new BigNumber((_741_v1).length)));
      let _743_v3;
      _743_v3 = _dafny.Map.Empty.slice().updateUnsafe(false,_740_v0);
      _742_v2 = (_742_v2).update(((_this).f17)[_module.__default.safeIndex(_740_v0, new BigNumber(((_this).f17).length))], (_740_v0).multipliedBy(new BigNumber((_743_v3).length)));
      let _744_i0;
      _744_i0 = _dafny.ZERO;
      L6: {
        while (_module.__default.fm0(_this.f16, _this.f16, globalState)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_744_i0)) {
              break L6;
            }
            _744_i0 = (_744_i0).plus(_dafny.ONE);
            if ((_module.__default.safeDivisionInt(new BigNumber(((_this).f24).length), (_dafny.ZERO).minus(_740_v0))).isEqualTo(new BigNumber(878))) {
              (globalState).f2 = _this.f16;
              let _745_v4;
              let _nw139 = Array((new BigNumber(18)).toNumber()).fill(false);
              _745_v4 = _nw139;
              let _index146 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length));
              (_745_v4)[_index146] = !(_this.f16);
              let _index147 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length));
              let _rhs155 = ((!(_this.f16)) ? (_742_v2) : (_742_v2));
              let _rhs156 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(((_this).f24).length)), _740_v0);
              let _rhs157 = _this.f16;
              let _lhs137 = _745_v4;
              let _lhs138 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length));
              _742_v2 = _rhs155;
              _740_v0 = _rhs156;
              _lhs137[_lhs138] = _rhs157;
              let _746_v5;
              let _nw140 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
              _746_v5 = _nw140;
              let _747_v6;
              _747_v6 = new _dafny.CodePoint('w'.codePointAt(0));
              let _index148 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_746_v5).length));
              (_746_v5)[_index148] = _dafny.Map.Empty.slice().updateUnsafe(_740_v0,_747_v6);
              let _748_v7;
              _748_v7 = _module.D3.create_DC10((_745_v4)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length))]);
              let _749_v8;
              _749_v8 = _dafny.Set.fromElements(_this.f16, _this.f16);
              let _750_v9;
              _750_v9 = _dafny.Map.Empty.slice().updateUnsafe(_740_v0,_747_v6);
              let _index149 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_746_v5).length));
              let _rhs158 = ((((_745_v4)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length))]) ? ((_745_v4)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length))]) : (!((_745_v4)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length))])))) === ((_749_v8).contains((_745_v4)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length))]));
              let _rhs159 = ((_module.__default.fm0(!(_this.f16), _this.f16, globalState)) ? (_750_v9) : ((_750_v9).Merge(_750_v9)));
              let _rhs160 = _748_v7;
              let _lhs139 = globalState;
              let _lhs140 = _746_v5;
              let _lhs141 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_746_v5).length));
              _lhs139.f5 = _rhs158;
              _lhs140[_lhs141] = _rhs159;
              _748_v7 = _rhs160;
              r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_740_v0).multipliedBy(_740_v0)));
              let _751_v10;
              _751_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,(_745_v4)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_745_v4).length))]);
              let _752_v11;
              _752_v11 = _module.D1.create_DC4(false, (_dafny.ZERO).minus(_740_v0), _751_v10);
              let _753_v12;
              _753_v12 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(832)), (_743_v3).Merge(_743_v3), _743_v3);
              let _754_v13;
              _754_v13 = _dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)));
              let _755_v14;
              _755_v14 = _module.D4.create_DC11(_753_v12);
              let _rhs161 = _752_v11;
              let _rhs162 = (_dafny.MultiSet.fromElements(_743_v3, _743_v3)).Intersect((_755_v14).dtor_cf20);
              let _rhs163 = (_740_v0).minus(_740_v0);
              let _rhs164 = _740_v0;
              let _rhs165 = _dafny.Seq.update(_dafny.Seq.of(((_this).f24)[_module.__default.safeIndex(_740_v0, new BigNumber(((_this).f24).length))], new _dafny.CodePoint('a'.codePointAt(0))), _module.__default.safeIndex((_740_v0).multipliedBy(_740_v0), new BigNumber((_dafny.Seq.of(((_this).f24)[_module.__default.safeIndex(_740_v0, new BigNumber(((_this).f24).length))], new _dafny.CodePoint('a'.codePointAt(0)))).length)), ((_this.f16) ? (_747_v6) : ((((_750_v9).contains(_740_v0)) ? ((_750_v9).get(_740_v0)) : (_747_v6)))));
              _752_v11 = _rhs161;
              _753_v12 = _rhs162;
              _740_v0 = _rhs163;
              _740_v0 = _rhs164;
              _754_v13 = _rhs165;
            } else {
              let _rhs166 = _740_v0;
              let _rhs167 = new BigNumber(-554);
              _740_v0 = _rhs166;
              r1 = _rhs167;
              (globalState).f2 = _this.f16;
              let _756_v15;
              _756_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16);
              _741_v1 = (((((_756_v15).contains(_this.f16)) ? ((_756_v15).get(_this.f16)) : (_this.f16))) ? (_741_v1) : (_741_v1));
              let _757_v16;
              _757_v16 = _dafny.Set.fromElements((_dafny.ZERO).minus((((_743_v3).contains(_this.f16)) ? ((_743_v3).get(_this.f16)) : (_740_v0))), _740_v0);
              let _758_v17;
              _758_v17 = _dafny.MultiSet.fromElements(_757_v16);
              _758_v17 = _758_v17;
              let _759_v18;
              let _nw141 = Array((new BigNumber(28)).toNumber()).fill(false);
              _759_v18 = _nw141;
              let _760_v19;
              let _nw142 = new _module.C0();
              _nw142.__ctor(_740_v0, _759_v18);
              _760_v19 = _nw142;
            }
            let _761_v20;
            _761_v20 = _module.D4.create_DC12(_this.f16, _module.__default.fm0(_this.f16, _this.f16, globalState), _740_v0);
            let _pat_let_tv19 = _740_v0;
            let _pat_let_tv20 = _740_v0;
            let _pat_let_tv21 = _761_v20;
            let _pat_let_tv22 = globalState;
            let _762_v21;
            let _nw143 = Array((new BigNumber(17)).toNumber());
            _nw143[(_dafny.ZERO).toNumber()] = _761_v20;
            _nw143[(_dafny.ONE).toNumber()] = _module.D4.create_DC12(_this.f16, _this.f16, _module.__default.fm3(globalState));
            _nw143[(new BigNumber(2)).toNumber()] = _module.__default.fm21(_740_v0, _740_v0, globalState);
            _nw143[(new BigNumber(3)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(4)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(5)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(6)).toNumber()] = function (_pat_let16_0) {
              return function (_763_dt__update__tmp_h0) {
                return function (_pat_let17_0) {
                  return function (_764_dt__update_hcf23_h0) {
                    return _module.D4.create_DC12((_763_dt__update__tmp_h0).dtor_cf21, (_763_dt__update__tmp_h0).dtor_cf22, _764_dt__update_hcf23_h0);
                  }(_pat_let17_0);
                }(new BigNumber((_dafny.Seq.of(_pat_let_tv19)).length));
              }(_pat_let16_0);
            }(_761_v20);
            _nw143[(new BigNumber(7)).toNumber()] = _module.D4.create_DC12(_this.f16, _this.f16, _740_v0);
            _nw143[(new BigNumber(8)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(9)).toNumber()] = function (_pat_let18_0) {
              return function (_765_dt__update__tmp_h1) {
                return function (_pat_let19_0) {
                  return function (_766_dt__update_hcf23_h1) {
                    return _module.D4.create_DC12((_765_dt__update__tmp_h1).dtor_cf21, (_765_dt__update__tmp_h1).dtor_cf22, _766_dt__update_hcf23_h1);
                  }(_pat_let19_0);
                }(_pat_let_tv20);
              }(_pat_let18_0);
            }(_module.__default.fm21(_740_v0, _740_v0, globalState));
            _nw143[(new BigNumber(10)).toNumber()] = _module.D4.create_DC12(_this.f16, _this.f16, (_module.__default.fm21(_740_v0, _740_v0, globalState)).dtor_cf23);
            _nw143[(new BigNumber(11)).toNumber()] = function (_pat_let20_0) {
              return function (_767_dt__update__tmp_h2) {
                return function (_pat_let21_0) {
                  return function (_768_dt__update_hcf22_h0) {
                    return function (_pat_let22_0) {
                      return function (_771_dt__update_hcf21_h0) {
                        return _module.D4.create_DC12(_771_dt__update_hcf21_h0, _768_dt__update_hcf22_h0, (_767_dt__update__tmp_h2).dtor_cf23);
                      }(_pat_let22_0);
                    }(_module.__default.fm0(_this.f16, (function (_pat_let23_0) {
                      return function (_769_dt__update__tmp_h3) {
                        return function (_pat_let24_0) {
                          return function (_770_dt__update_hcf21_h1) {
                            return _module.D4.create_DC12(_770_dt__update_hcf21_h1, (_769_dt__update__tmp_h3).dtor_cf22, (_769_dt__update__tmp_h3).dtor_cf23);
                          }(_pat_let24_0);
                        }(!(_this.f16));
                      }(_pat_let23_0);
                    }(_pat_let_tv21)).dtor_cf21, _pat_let_tv22));
                  }(_pat_let21_0);
                }(!(_this.f16));
              }(_pat_let20_0);
            }(_761_v20);
            _nw143[(new BigNumber(12)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(13)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(14)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(15)).toNumber()] = _761_v20;
            _nw143[(new BigNumber(16)).toNumber()] = _761_v20;
            _762_v21 = _nw143;
            let _index150 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_762_v21).length));
            (_762_v21)[_index150] = _761_v20;
            let _index151 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_762_v21).length));
            (_762_v21)[_index151] = _module.D4.create_DC12(false, _module.__default.fm0(_this.f16, _this.f16, globalState), _module.__default.safeModuloInt(_740_v0, _740_v0));
            let _772_v22;
            _772_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(_740_v0), _740_v0),!(_this.f16));
            _772_v22 = (_772_v22).update(_740_v0, _this.f16);
            let _773_v23;
            _773_v23 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f16) ? (_this.f16) : (_this.f16)),(_740_v0).isLessThanOrEqualTo((((_743_v3).contains(true)) ? ((_743_v3).get(true)) : (_740_v0))));
            let _774_v24;
            _774_v24 = _dafny.Seq.of((_773_v23).update(_this.f16, _this.f16));
            let _775_v25;
            _775_v25 = _module.D4.create_DC13(new BigNumber((_742_v2).length), _740_v0, _740_v0);
            _773_v23 = (_773_v23).Merge((_774_v24)[_module.__default.safeIndex((_775_v25).dtor_cf24, new BigNumber((_774_v24).length))]);
          }
        }
      }
      (_this).f16 = _this.f16;
      let _776_v26;
      let _nw144 = Array((new BigNumber(20)).toNumber()).fill(false);
      _776_v26 = _nw144;
      let _index152 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
      (_776_v26)[_index152] = _this.f16;
      let _777_v27;
      _777_v27 = _module.D0.create_DC2(_740_v0, _740_v0, _740_v0, _776_v26);
      let _778_v28;
      _778_v28 = _dafny.MultiSet.fromElements(_module.D0.create_DC2(_740_v0, _740_v0, _740_v0, _776_v26));
      let _index153 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
      (_776_v26)[_index153] = !(_778_v28).contains(_777_v27);
      (globalState).f5 = !((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]);
      if (!((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]) || (_this.f16)) {
        _740_v0 = _740_v0;
        let _779_v29;
        _779_v29 = _dafny.MultiSet.fromElements((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]);
        if (!(_779_v29).contains((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))])) {
          let _780_v30;
          _780_v30 = new _dafny.CodePoint('w'.codePointAt(0));
          (globalState).f15 = ((_this.f16) ? (new _dafny.CodePoint('m'.codePointAt(0))) : (_780_v30));
          let _781_v31;
          _781_v31 = _dafny.Seq.UnicodeFromString("dita");
          _781_v31 = _dafny.Seq.update((_this).f24, _module.__default.safeIndex((_dafny.ZERO).minus(_740_v0), new BigNumber(((_this).f24).length)), _780_v30);
          let _782_v32;
          let _nw145 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _782_v32 = _nw145;
          let _783_v33;
          _783_v33 = _module.D2.create_DC8(_743_v3, _782_v32, false, (_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))], _740_v0);
          let _784_v34;
          _784_v34 = _module.D2.create_DC8(_743_v3, (_783_v33).dtor_cf14, _this.f16, (_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))], new BigNumber((_743_v3).length));
          let _785_v35;
          _785_v35 = _module.D6.create_DC17(_784_v34);
          let _786_v36;
          _786_v36 = _dafny.MultiSet.fromElements(_785_v35);
          let _787_v37;
          _787_v37 = _dafny.Seq.of(_786_v36);
          let _788_v38;
          _788_v38 = _787_v37;
          let _rhs168 = _module.__default.fm3(globalState);
          let _rhs169 = (_787_v37);
          r0 = _rhs168;
          _787_v37 = _rhs169;
          (globalState).f2 = _this.f16;
          let _789_v39;
          let _nw146 = Array((new BigNumber(24)).toNumber());
          _nw146[(_dafny.ZERO).toNumber()] = _780_v30;
          _nw146[(_dafny.ONE).toNumber()] = _780_v30;
          _nw146[(new BigNumber(2)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(3)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(4)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(5)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(6)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(7)).toNumber()] = (((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]) ? ((_781_v31)[_module.__default.safeIndex(_740_v0, new BigNumber((_781_v31).length))]) : (_780_v30));
          _nw146[(new BigNumber(8)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(9)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(10)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(11)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(12)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('g'.codePointAt(0));
          _nw146[(new BigNumber(14)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(15)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(16)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(17)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(18)).toNumber()] = _module.__default.fm2(globalState);
          _nw146[(new BigNumber(19)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(20)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(21)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(22)).toNumber()] = _780_v30;
          _nw146[(new BigNumber(23)).toNumber()] = _780_v30;
          _789_v39 = _nw146;
          _789_v39 = _789_v39;
        } else {
          let _790_v40;
          let _nw147 = Array((new BigNumber(20)).toNumber()).fill(_module.D6.Default());
          _790_v40 = _nw147;
          let _791_v41;
          _791_v41 = _module.D6.create_DC16(true, _this.f16);
          let _index154 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_790_v40).length));
          (_790_v40)[_index154] = _791_v41;
          let _index155 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_790_v40).length));
          (_790_v40)[_index155] = _module.D6.create_DC16((_740_v0).isLessThan(new BigNumber(((_this).f24).length)), _this.f16);
          let _792_v42;
          let _init21 = function (_793_i1) {
            return _dafny.Seq.Concat((_this).f17, (_this).f17);
          };
          let _nw148 = Array((new BigNumber(25)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw148.length); _i0_21++) {
            _nw148[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _792_v42 = _nw148;
          let _794_v43;
          _794_v43 = _dafny.Map.Empty.slice().updateUnsafe(_741_v1,_740_v0);
          let _index156 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_792_v42).length));
          (_792_v42)[_index156] = _dafny.Seq.of(_740_v0, new BigNumber((_794_v43).length), _740_v0);
          let _index157 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_792_v42).length));
          (_792_v42)[_index157] = _module.__default.fm17(globalState);
          let _795_v44;
          _795_v44 = _dafny.Set.fromElements((_dafny.ZERO).minus(_740_v0), _740_v0, _740_v0, (_dafny.ZERO).minus(_740_v0), _740_v0);
          (_this).f16 = !(!((_795_v44).IsProperSubsetOf(_795_v44)));
          let _796_v45;
          let _nw149 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _796_v45 = _nw149;
          let _797_v46;
          _797_v46 = _dafny.Set.fromElements(_741_v1, _741_v1);
          let _798_v47;
          _798_v47 = _dafny.Map.Empty.slice().updateUnsafe(_796_v45,_797_v46);
          _798_v47 = (_798_v47).update(_796_v45, _dafny.Set.fromElements(_741_v1, (_module.__default.fm22(_this.f16, _dafny.MultiSet.FromArray((_792_v42)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_792_v42).length))]), (_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))], globalState)).dtor_cf34));
          let _799_v48;
          let _nw150 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _799_v48 = _nw150;
          let _index158 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_799_v48).length));
          (_799_v48)[_index158] = _740_v0;
          let _index159 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_799_v48).length));
          let _rhs170 = _740_v0;
          let _rhs171 = _779_v29;
          let _rhs172 = new BigNumber(-58);
          let _rhs173 = _777_v27;
          let _lhs142 = _799_v48;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_799_v48).length));
          r1 = _rhs170;
          _779_v29 = _rhs171;
          _lhs142[_lhs143] = _rhs172;
          _777_v27 = _rhs173;
        }
        if (!((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))])) {
          let _800_v49;
          let _nw151 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _800_v49 = _nw151;
          let _rhs174 = _module.__default.fm3(globalState);
          let _rhs175 = _800_v49;
          let _rhs176 = _740_v0;
          let _rhs177 = (((_this.f16) ? ((_dafny.ZERO).minus(_740_v0)) : (_740_v0))).plus(_740_v0);
          _740_v0 = _rhs174;
          _800_v49 = _rhs175;
          r0 = _rhs176;
          _740_v0 = _rhs177;
          r1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat((_this).f24, (_this).f24), _module.__default.safeIndex(_740_v0, new BigNumber((_dafny.Seq.Concat((_this).f24, (_this).f24)).length)), new _dafny.CodePoint('x'.codePointAt(0))), (_this).f24)).length);
          r0 = new BigNumber(960);
          let _801_v50;
          let _nw152 = Array((new BigNumber(17)).toNumber()).fill([]);
          _801_v50 = _nw152;
          let _index160 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_801_v50).length));
          (_801_v50)[_index160] = _800_v49;
          let _index161 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_801_v50).length));
          (_801_v50)[_index161] = _800_v49;
          let _802_v51;
          _802_v51 = _dafny.Map.Empty.slice().updateUnsafe((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))],(_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]);
          let _803_v52;
          _803_v52 = _dafny.Map.Empty.slice().updateUnsafe(_802_v51,((_this.f16) ? (_800_v49) : ((_801_v50)[_module.__default.safeIndex(new BigNumber(834), new BigNumber((_801_v50).length))])));
          let _804_v53;
          let _init22 = ((_805_v51) => function (_806_i2) {
            return (_805_v51).update(_this.f16, _this.f16);
          })(_802_v51);
          let _nw153 = Array((new BigNumber(24)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw153.length); _i0_22++) {
            _nw153[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _804_v53 = _nw153;
          let _index162 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_804_v53).length));
          (_804_v53)[_index162] = _802_v51;
          let _index163 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_804_v53).length));
          let _rhs178 = _803_v52;
          let _rhs179 = ((_802_v51).Merge(_802_v51)).Merge((_module.__default.fm23(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_740_v0), _740_v0, _740_v0), !(!((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))])), globalState)).update(_this.f16, (_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]));
          let _lhs144 = _804_v53;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_804_v53).length));
          _803_v52 = _rhs178;
          _lhs144[_lhs145] = _rhs179;
        } else {
          let _rhs180 = function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of _dafny.IntegerRange(new BigNumber(147), new BigNumber(327))) {
              let _807_v54 = _compr_43;
              if (((new BigNumber(147)).isLessThanOrEqualTo(_807_v54)) && ((_807_v54).isLessThan(new BigNumber(327)))) {
                _coll43.push([_module.__default.safeModuloInt(_807_v54, new BigNumber((_779_v29).cardinality())),new BigNumber(((_this).f24).length)]);
              }
            }
            return _coll43;
          }();
          let _rhs181 = _740_v0;
          _742_v2 = _rhs180;
          r1 = _rhs181;
          let _808_v55;
          _808_v55 = _dafny.Seq.UnicodeFromString("he");
          _808_v55 = _808_v55;
          let _809_v56;
          let _init23 = ((_810_v0) => function (_811_i3) {
            return (_811_i3).plus(_810_v0);
          })(_740_v0);
          let _nw154 = Array((new BigNumber(22)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw154.length); _i0_23++) {
            _nw154[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _809_v56 = _nw154;
          let _init24 = ((_812_v0) => function (_813_i4) {
            return _module.__default.safeDivisionInt(_813_i4, _812_v0);
          })(_740_v0);
          let _nw155 = Array((new BigNumber(25)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw155.length); _i0_24++) {
            _nw155[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _809_v56 = _nw155;
          r1 = _740_v0;
          _743_v3 = (_743_v3).Merge(_dafny.Map.Empty.slice().updateUnsafe(!((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]),_740_v0));
        }
        let _814_v57;
        _814_v57 = _module.D8.create_DC23();
        let _815_v58;
        _815_v58 = _dafny.Set.fromElements(_814_v57, _814_v57, _814_v57);
        let _816_v59;
        _816_v59 = new _dafny.CodePoint('s'.codePointAt(0));
        let _817_v60;
        _817_v60 = _dafny.Map.Empty.slice().updateUnsafe(_816_v59,new BigNumber(((_this).f24).length));
        let _818_v61;
        _818_v61 = _dafny.MultiSet.fromElements((_815_v58).Union(_815_v58), _dafny.Set.fromElements(_814_v57, _module.__default.fm24(((_this).f17)[_module.__default.safeIndex(new BigNumber((_817_v60).length), new BigNumber(((_this).f17).length))], _816_v59, globalState)), _815_v58, _815_v58, _dafny.Set.fromElements(_814_v57));
        let _819_v62;
        _819_v62 = _dafny.Seq.of((new BigNumber(((_this).f17).length)).minus(_740_v0));
        let _index164 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
        let _rhs182 = _module.__default.fm25(true, _740_v0, globalState);
        let _rhs183 = new BigNumber(675);
        let _rhs184 = (((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]) ? (!(_this.f16)) : (_this.f16));
        let _rhs185 = _819_v62;
        let _rhs186 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_this.f16, _this.f16), _741_v1), _741_v1);
        let _lhs146 = _776_v26;
        let _lhs147 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
        _818_v61 = _rhs182;
        _740_v0 = _rhs183;
        _lhs146[_lhs147] = _rhs184;
        _819_v62 = _rhs185;
        _741_v1 = _rhs186;
        let _820_v63;
        let _init25 = ((_821_v0) => function (_822_i5) {
          return (_822_i5).multipliedBy(_821_v0);
        })(_740_v0);
        let _nw156 = Array((new BigNumber(9)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw156.length); _i0_25++) {
          _nw156[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _820_v63 = _nw156;
        let _index165 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_820_v63).length));
        (_820_v63)[_index165] = _740_v0;
        let _index166 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_820_v63).length));
        (_820_v63)[_index166] = (_dafny.ZERO).minus(_740_v0);
      } else {
        let _823_v64;
        _823_v64 = _dafny.Set.fromElements(_dafny.Seq.of(_module.__default.fm3(globalState)), _dafny.Seq.of((_dafny.ZERO).minus(_740_v0)), _dafny.Seq.of(_740_v0));
        if ((_823_v64).IsProperSubsetOf(_823_v64)) {
          let _824_v65;
          let _nw157 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _824_v65 = _nw157;
          let _index167 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_824_v65).length));
          (_824_v65)[_index167] = (_dafny.ZERO).minus(_740_v0);
          let _index168 = _module.__default.safeIndex(new BigNumber(750), new BigNumber((_824_v65).length));
          (_824_v65)[_index168] = (((_743_v3).contains((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))])) ? ((_743_v3).get((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))])) : ((_740_v0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("isodoen")).length))));
          (globalState).f2 = ((_740_v0).isLessThan((_824_v65)[_module.__default.safeIndex(new BigNumber(750), new BigNumber((_824_v65).length))])) || (false);
          let _index169 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
          (_776_v26)[_index169] = !((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]);
          _743_v3 = (_743_v3).update(_this.f16, (_824_v65)[_module.__default.safeIndex(new BigNumber(750), new BigNumber((_824_v65).length))]);
          _module.__default.m0((_824_v65)[_module.__default.safeIndex(new BigNumber(750), new BigNumber((_824_v65).length))], (_this).f24, (function (_pat_let25_0) {
            return function (_825_dt__update__tmp_h5) {
              return function (_pat_let26_0) {
                return function (_826_dt__update_hcf39_h0) {
                  return _module.D8.create_DC21((_825_dt__update__tmp_h5).dtor_cf35, (_825_dt__update__tmp_h5).dtor_cf36, (_825_dt__update__tmp_h5).dtor_cf37, (_825_dt__update__tmp_h5).dtor_cf38, _826_dt__update_hcf39_h0);
                }(_pat_let26_0);
              }(new BigNumber(856));
            }(_pat_let25_0);
          }(_module.__default.fm26(globalState))).dtor_cf35, _740_v0, globalState);
        } else {
          let _827_v66;
          _827_v66 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16);
          let _828_v67;
          let _nw158 = new _module.C1();
          _nw158.__ctor((_740_v0).multipliedBy(_740_v0), _740_v0, _module.__default.fm0((((_827_v66).contains((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))])) ? ((_827_v66).get((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))])) : (_this.f16)), false, globalState), (_this).f17, _740_v0);
          _828_v67 = _nw158;
          _740_v0 = new BigNumber(853);
          r0 = (_828_v67).f30;
          let _829_v68;
          let _init26 = ((_830_v67) => function (_831_i6) {
            return _module.__default.safeDivisionInt(_831_i6, (_830_v67).f30);
          })(_828_v67);
          let _nw159 = Array((new BigNumber(7)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw159.length); _i0_26++) {
            _nw159[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _829_v68 = _nw159;
          let _832_v69;
          _832_v69 = _module.D6.create_DC18(_829_v68);
          let _index170 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
          let _index171 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
          let _rhs187 = (_832_v69).dtor_cf32;
          let _rhs188 = !(!((new BigNumber((_741_v1).length)).isLessThanOrEqualTo(_740_v0)));
          let _rhs189 = (_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))];
          let _rhs190 = (_741_v1)[_module.__default.safeIndex(_740_v0, new BigNumber((_741_v1).length))];
          let _rhs191 = _829_v68;
          let _lhs148 = _776_v26;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
          let _lhs150 = globalState;
          let _lhs151 = _776_v26;
          let _lhs152 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length));
          _829_v68 = _rhs187;
          _lhs148[_lhs149] = _rhs188;
          _lhs150.f5 = _rhs189;
          _lhs151[_lhs152] = _rhs190;
          _829_v68 = _rhs191;
          let _833_v70;
          let _nw160 = new _module.C2();
          _nw160.__ctor(_740_v0);
          _833_v70 = _nw160;
          _833_v70 = ((_this.f16) ? (_833_v70) : (_833_v70));
        }
        let _834_v71;
        _834_v71 = _dafny.Seq.UnicodeFromString("wmq");
        let _835_v72;
        _835_v72 = new _dafny.CodePoint('i'.codePointAt(0));
        _834_v71 = _dafny.Seq.update(_834_v71, _module.__default.safeIndex(_740_v0, new BigNumber((_834_v71).length)), _835_v72);
        _740_v0 = (_740_v0).minus(_740_v0);
        let _836_v73;
        _836_v73 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]);
        let _837_v74;
        _837_v74 = _dafny.Map.Empty.slice().updateUnsafe((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))],_776_v26);
        let _838_v75;
        _838_v75 = _dafny.MultiSet.fromElements(new BigNumber((_837_v74).length), _740_v0);
        (globalState).f5 = (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_740_v0, new BigNumber((_741_v1).length)), _module.__default.safeIndex((_dafny.ZERO).minus(_740_v0), new BigNumber((_dafny.Seq.of(_740_v0, new BigNumber((_741_v1).length))).length)), new BigNumber((_836_v73).length))).length), _740_v0, (_dafny.ZERO).minus(_740_v0), _740_v0)).IsProperSubsetOf((((_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))]) ? (_838_v75) : (_dafny.MultiSet.fromElements(new BigNumber((_834_v71).length)))));
        (globalState).f2 = (_776_v26)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_776_v26).length))];
      }
      r0 = _740_v0;
      r1 = _740_v0;
      let _839_v76;
      _839_v76 = _dafny.Seq.of(_module.__default.fm17(globalState), _dafny.Seq.of(_740_v0, _740_v0), (_this).f17, (_this).f17);
      let _840_v77;
      _840_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_839_v76)[_module.__default.safeIndex(_740_v0, new BigNumber((_839_v76).length))]).length),false);
      r2 = _840_v77;
      return [r0, r1, r2];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this.f23 = false;
      this._f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22, f23) {
      let _this = this;
      (_this)._f22 = f22;
      (_this).f23 = f23;
      return;
    }
    fm9(p0, globalState) {
      let _this = this;
      return _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(533), new BigNumber(795)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f22, (_this).f22)).cardinality()))), _dafny.Seq.of(_dafny.Seq.of((_this).f22, (_this).f22, (_this).f22, (_this).f22, (_this).f22), _dafny.Seq.Create(_module.__default.abs(new BigNumber(726)), function (_841_i0) {
        return (_this).f22;
      }))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-128)), function (_842_i1) {
        return _dafny.Seq.of((_this).f22);
      }));
    };
    fm10(p0, p1, globalState) {
      let _this = this;
      if ((_module.D1.create_DC5(new BigNumber((_dafny.Seq.of(_this.f23, false)).length), false)).dtor_cf11) {
        return _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_dafny.MultiSet.fromElements((_this).f22, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_module.D2.create_DC7(_dafny.Set.fromElements(_this.f23)))).length),(_this).f22)).length)));
      } else if (_this.f23) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, _this.f23, _this.f23, _this.f23, _this.f23)).length),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f22,_this.f23)).length), new BigNumber(702), (_dafny.ZERO).minus((_this).f22), (_this).f22, new BigNumber(62)));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(508),_dafny.MultiSet.fromElements((_this).f22));
      }
    };
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _843_v0;
      _843_v0 = new _dafny.CodePoint('j'.codePointAt(0));
      let _844_v1;
      _844_v1 = _dafny.Seq.UnicodeFromString("hvsewr");
      (globalState).f2 = !_dafny.Seq.contains(_844_v1, _843_v0);
      if (_this.f23) {
        let _845_v2;
        _845_v2 = _dafny.MultiSet.fromElements(p1);
        let _846_v3;
        _846_v3 = _module.D1.create_DC3(_845_v2);
        let _source9 = _846_v3;
        if (_source9.is_DC4) {
          let _847___mcc_h0 = (_source9).cf7;
          let _848___mcc_h1 = (_source9).cf8;
          let _849___mcc_h2 = (_source9).cf9;
          let _850_cf9 = _849___mcc_h2;
          let _851_cf8 = _848___mcc_h1;
          let _852_cf7 = _847___mcc_h0;
          let _853_v4;
          let _init27 = ((_854_p2) => function (_855_i0) {
            return (_855_i0).plus(_854_p2);
          })(p2);
          let _nw161 = Array((new BigNumber(7)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw161.length); _i0_27++) {
            _nw161[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _853_v4 = _nw161;
          let _index172 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_853_v4).length));
          (_853_v4)[_index172] = (((_this).fm9(p2, globalState)) ? (new BigNumber(368)) : (_module.__default.fm3(globalState)));
          let _index173 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_853_v4).length));
          let _rhs192 = _851_cf8;
          let _rhs193 = p2;
          let _rhs194 = p2;
          let _rhs195 = _851_cf8;
          let _rhs196 = (new BigNumber(-974)).multipliedBy((_this).f22);
          let _lhs153 = _853_v4;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_853_v4).length));
          _lhs153[_lhs154] = _rhs192;
          r2 = _rhs193;
          _851_cf8 = _rhs194;
          r2 = _rhs195;
          r2 = _rhs196;
          _850_cf9 = (_850_cf9).update(_852_cf7, _this.f23);
          let _856_v5;
          _856_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(772),(_this).f22);
          let _857_v6;
          _857_v6 = _dafny.Seq.of(_852_cf7);
          _856_v5 = (_856_v5).update((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_857_v6).length),p0)).length))), _851_cf8);
          let _858_v7;
          _858_v7 = _dafny.Map.Empty.slice().updateUnsafe(_843_v0,_this.f23);
          let _859_v8;
          _859_v8 = _module.D9.create_DC25(_858_v7);
          let _860_v9;
          _860_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_859_v8).dtor_cf44).length),false);
          let _861_v10;
          let _nw162 = Array((new BigNumber(16)).toNumber());
          _nw162[(_dafny.ZERO).toNumber()] = (((_860_v9).contains((_853_v4)[_module.__default.safeIndex(new BigNumber(375), new BigNumber((_853_v4).length))])) ? ((_860_v9).get((_853_v4)[_module.__default.safeIndex(new BigNumber(375), new BigNumber((_853_v4).length))])) : (p1));
          _nw162[(_dafny.ONE).toNumber()] = !(_this.f23);
          _nw162[(new BigNumber(2)).toNumber()] = _852_cf7;
          _nw162[(new BigNumber(3)).toNumber()] = p0;
          _nw162[(new BigNumber(4)).toNumber()] = p0;
          _nw162[(new BigNumber(5)).toNumber()] = _this.f23;
          _nw162[(new BigNumber(6)).toNumber()] = _852_cf7;
          _nw162[(new BigNumber(7)).toNumber()] = _this.f23;
          _nw162[(new BigNumber(8)).toNumber()] = p0;
          _nw162[(new BigNumber(9)).toNumber()] = p1;
          _nw162[(new BigNumber(10)).toNumber()] = p1;
          _nw162[(new BigNumber(11)).toNumber()] = false;
          _nw162[(new BigNumber(12)).toNumber()] = false;
          _nw162[(new BigNumber(13)).toNumber()] = p1;
          _nw162[(new BigNumber(14)).toNumber()] = _this.f23;
          _nw162[(new BigNumber(15)).toNumber()] = p1;
          _861_v10 = _nw162;
          let _862_v11;
          let _nw163 = new _module.C0();
          _nw163.__ctor((_853_v4)[_module.__default.safeIndex(new BigNumber(375), new BigNumber((_853_v4).length))], _861_v10);
          _862_v11 = _nw163;
        } else if (_source9.is_DC5) {
          let _863___mcc_h3 = (_source9).cf10;
          let _864___mcc_h4 = (_source9).cf11;
          let _865_cf11 = _864___mcc_h4;
          let _866_cf10 = _863___mcc_h3;
          let _867_v12;
          _867_v12 = _dafny.Seq.of(false, _this.f23);
          let _868_v13;
          _868_v13 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
          let _869_v14;
          _869_v14 = _dafny.Set.fromElements((_this).f22);
          let _870_v15;
          _870_v15 = _dafny.Map.Empty.slice().updateUnsafe(_869_v14,(((_868_v13).contains(p1)) ? ((_868_v13).get(p1)) : (!(false))));
          let _871_v16;
          _871_v16 = _dafny.Seq.of(p2, new BigNumber((_870_v15).length));
          let _872_v17;
          let _nw164 = new _module.C1();
          _nw164.__ctor(new BigNumber((_867_v12).length), new BigNumber((_868_v13).length), !(p1), _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm3(globalState)), _871_v16), (_this).f22);
          _872_v17 = _nw164;
          (globalState).f2 = !(_865_cf11);
          let _873_v18;
          _873_v18 = _dafny.Map.Empty.slice().updateUnsafe((_872_v17).f30,p0);
          let _874_v19;
          let _nw165 = Array((new BigNumber(8)).toNumber());
          _nw165[(_dafny.ZERO).toNumber()] = p1;
          _nw165[(_dafny.ONE).toNumber()] = _865_cf11;
          _nw165[(new BigNumber(2)).toNumber()] = _865_cf11;
          _nw165[(new BigNumber(3)).toNumber()] = p0;
          _nw165[(new BigNumber(4)).toNumber()] = false;
          _nw165[(new BigNumber(5)).toNumber()] = _865_cf11;
          _nw165[(new BigNumber(6)).toNumber()] = (_867_v12)[_module.__default.safeIndex((_872_v17).f30, new BigNumber((_867_v12).length))];
          _nw165[(new BigNumber(7)).toNumber()] = (((_873_v18).contains(new BigNumber((_867_v12).length))) ? ((_873_v18).get(new BigNumber((_867_v12).length))) : (_module.__default.fm0(p0, p1, globalState)));
          _874_v19 = _nw165;
          _874_v19 = _874_v19;
          let _875_v20;
          let _nw166 = new _module.C0();
          _nw166.__ctor((_872_v17).fm13(_this.f23, globalState), _874_v19);
          _875_v20 = _nw166;
        } else if (_source9.is_DC6) {
          r2 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(447)), function (_876_i1) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          })).length);
          let _877_v21;
          let _nw167 = Array((new BigNumber(21)).toNumber()).fill(false);
          _877_v21 = _nw167;
          let _index174 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_877_v21).length));
          (_877_v21)[_index174] = (p0) === (p0);
          let _index175 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_877_v21).length));
          (_877_v21)[_index175] = !(true);
          (globalState).f2 = p0;
          let _878_v22;
          let _init28 = ((_879_p2) => function (_880_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe(false,_879_p2);
          })(p2);
          let _nw168 = Array((new BigNumber(6)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw168.length); _i0_28++) {
            _nw168[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _878_v22 = _nw168;
          let _881_v23;
          _881_v23 = _dafny.Map.Empty.slice().updateUnsafe(_843_v0,_878_v22);
          _881_v23 = (_881_v23).update(_module.__default.fm2(globalState), _878_v22);
        } else {
          let _882___mcc_h5 = (_source9).cf6;
          let _883_cf6 = _882___mcc_h5;
          let _884_v24;
          _884_v24 = _dafny.MultiSet.fromElements(_844_v1, _844_v1);
          let _885_v25;
          _885_v25 = _dafny.Seq.of(new BigNumber((_884_v24).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(234)), ((_886_v0) => function (_887_i3) {
            return _886_v0;
          })(_843_v0))).length));
          let _888_v26;
          _888_v26 = _module.D3.create_DC10(p1);
          let _889_v27;
          _889_v27 = _dafny.Set.fromElements(false, !((_888_v26).dtor_cf19));
          let _890_v28;
          let _nw169 = new _module.C1();
          _nw169.__ctor((_dafny.ZERO).minus((_this).f22), p2, p0, _885_v25, new BigNumber((_889_v27).length));
          _890_v28 = _nw169;
          let _891_v29;
          _891_v29 = _dafny.MultiSet.fromElements(_890_v28);
          r2 = _module.__default.safeDivisionInt((((_891_v29).contains(_890_v28)) ? ((_891_v29).get(_890_v28)) : ((_890_v28).fm14(p2, (_890_v28).f30, _843_v0, globalState))), _module.__default.safeModuloInt((_890_v28).f29, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_890_v28).f29,p2)).update(p2, (_890_v28).f29)).length)));
          let _892_v30;
          _892_v30 = _module.D8.create_DC23();
          let _893_v31;
          _893_v31 = _module.D8.create_DC24(((p0) ? (_892_v30) : (_892_v30)));
          _893_v31 = _893_v31;
          let _894_v32;
          _894_v32 = _dafny.Map.Empty.slice().updateUnsafe(_844_v1,(_this).fm9(p2, globalState));
          _894_v32 = (_894_v32).update(_dafny.Seq.UnicodeFromString("weptditry"), _this.f23);
          let _895_v33;
          _895_v33 = _dafny.Map.Empty.slice().updateUnsafe(!(false),p0);
          let _896_v34;
          _896_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_895_v33).update(p1, p1)).length),p2);
          (globalState).f5 = ((((_896_v34).contains(p2)) ? ((_896_v34).get(p2)) : (p2))).isLessThan((_890_v28).fm15(_dafny.Seq.update(_885_v25, _module.__default.safeIndex((_this).f22, new BigNumber((_885_v25).length)), (_890_v28).f30), new BigNumber((_889_v27).length), (_890_v28).f30, _843_v0, globalState));
        }
        let _897_v35;
        let _nw170 = Array((new BigNumber(22)).toNumber()).fill(false);
        _897_v35 = _nw170;
        let _index176 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
        (_897_v35)[_index176] = p0;
        let _898_v36;
        let _nw171 = Array((new BigNumber(24)).toNumber());
        _nw171[(_dafny.ZERO).toNumber()] = _897_v35;
        _nw171[(_dafny.ONE).toNumber()] = _897_v35;
        _nw171[(new BigNumber(2)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(3)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(4)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(5)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(6)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(7)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(8)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(9)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(10)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(11)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(12)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(13)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(14)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(15)).toNumber()] = ((p0) ? (_897_v35) : (_897_v35));
        _nw171[(new BigNumber(16)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(17)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(18)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(19)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(20)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(21)).toNumber()] = _897_v35;
        _nw171[(new BigNumber(22)).toNumber()] = ((false) ? (_897_v35) : (_897_v35));
        _nw171[(new BigNumber(23)).toNumber()] = _897_v35;
        _898_v36 = _nw171;
        let _899_v37;
        _899_v37 = _dafny.Seq.of(p2);
        let _900_v38;
        _900_v38 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(569));
        let _901_v39;
        _901_v39 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_899_v37)[_module.__default.safeIndex((_this).f22, new BigNumber((_899_v37).length))]);
        let _index177 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
        let _rhs197 = ((p2).isLessThan((_dafny.ZERO).minus((_this).f22))) || (p1);
        let _rhs198 = p2;
        let _rhs199 = new BigNumber((_dafny.Seq.update(_899_v37, _module.__default.safeIndex(_module.__default.safeModuloInt((_this).f22, new BigNumber((_dafny.Set.fromElements(new BigNumber(996), new BigNumber((_844_v1).length))).length)), new BigNumber((_899_v37).length)), (((_900_v38).contains(p2)) ? ((_900_v38).get(p2)) : (_module.__default.fm3(globalState))))).length);
        let _rhs200 = (_module.__default.fm3(globalState)).isLessThan(new BigNumber((_901_v39).length));
        let _rhs201 = _898_v36;
        let _lhs155 = globalState;
        let _lhs156 = _897_v35;
        let _lhs157 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
        _lhs155.f2 = _rhs197;
        r2 = _rhs198;
        r2 = _rhs199;
        _lhs156[_lhs157] = _rhs200;
        _898_v36 = _rhs201;
        if (_dafny.areEqual(_dafny.Seq.Concat(_844_v1, _module.__default.fm1(globalState)), _844_v1)) {
          let _902_v40;
          let _nw172 = new _module.C3();
          _nw172.__ctor(_844_v1, _this.f23, _module.__default.fm17(globalState));
          _902_v40 = _nw172;
          let _903_v41;
          _903_v41 = _dafny.Map.Empty.slice().updateUnsafe(_843_v0,_this.f23);
          let _904_v42;
          _904_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC25(_903_v41),(p2).minus((_this).f22));
          let _905_v44;
          _905_v44 = _dafny.Map.Empty.slice().updateUnsafe(_843_v0,p2);
          let _906_v45;
          _906_v45 = _module.D9.create_DC25(function () {
  let _coll44 = new _dafny.Map();
  for (const _compr_44 of (_905_v44).Keys.Elements) {
    let _907_v43 = _compr_44;
    if ((_905_v44).contains(_907_v43)) {
      _coll44.push([_907_v43,p0]);
    }
  }
  return _coll44;
}());
          let _index178 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
          let _index179 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
          let _rhs202 = p0;
          let _rhs203 = p0;
          let _rhs204 = (_904_v42).update(_906_v45, new BigNumber(539));
          let _rhs205 = (_this).f22;
          let _lhs158 = _897_v35;
          let _lhs159 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
          let _lhs160 = _897_v35;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
          _lhs158[_lhs159] = _rhs202;
          _lhs160[_lhs161] = _rhs203;
          _904_v42 = _rhs204;
          r2 = _rhs205;
          let _908_v46;
          let _nw173 = Array((new BigNumber(16)).toNumber()).fill([]);
          _908_v46 = _nw173;
          let _909_v47;
          let _nw174 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _909_v47 = _nw174;
          let _index180 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_908_v46).length));
          (_908_v46)[_index180] = _909_v47;
          let _index181 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_908_v46).length));
          (_908_v46)[_index181] = _909_v47;
          let _910_v48;
          _910_v48 = _dafny.Seq.of((_897_v35)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length))], p0, false, (_dafny.Set.fromElements((_this).f22)).contains((_this).f22));
          let _911_v51;
          let _init29 = ((_912_v37) => function (_913_i4) {
            return (_dafny.Set.fromElements(function () {
              let _coll45 = new _dafny.Set();
              for (const _compr_45 of (_912_v37).Elements) {
                let _914_v49 = _compr_45;
                if (_dafny.Seq.contains(_912_v37, _914_v49)) {
                  _coll45.add((_914_v49).plus((_dafny.ZERO).minus(new BigNumber(-821))));
                }
              }
              return _coll45;
            }(), function () {
              let _coll46 = new _dafny.Set();
              for (const _compr_46 of _dafny.IntegerRange(new BigNumber(375), new BigNumber(298))) {
                let _915_v50 = _compr_46;
                if (((new BigNumber(375)).isLessThanOrEqualTo(_915_v50)) && ((_915_v50).isLessThan(new BigNumber(298)))) {
                  _coll46.add((_915_v50).minus((_this).f22));
                }
              }
              return _coll46;
            }())).Difference(_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f22)));
          })(_899_v37);
          let _nw175 = Array((new BigNumber(10)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw175.length); _i0_29++) {
            _nw175[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _911_v51 = _nw175;
          let _916_v52;
          _916_v52 = _dafny.Set.fromElements((_dafny.ZERO).minus(p2));
          let _917_v53;
          _917_v53 = _dafny.Set.fromElements(_916_v52, _916_v52);
          let _index182 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_911_v51).length));
          (_911_v51)[_index182] = _917_v53;
          let _918_v54;
          _918_v54 = _dafny.Map.Empty.slice().updateUnsafe(_844_v1,new _dafny.CodePoint('g'.codePointAt(0)));
          let _919_v56;
          _919_v56 = _dafny.Seq.of(_844_v1);
          let _index183 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_908_v46).length));
          let _index184 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_911_v51).length));
          let _rhs206 = _909_v47;
          let _rhs207 = _dafny.Seq.of(false);
          let _rhs208 = _917_v53;
          let _rhs209 = new BigNumber(-22);
          let _rhs210 = (function () {
            let _coll47 = new _dafny.Map();
            for (const _compr_47 of (_919_v56).Elements) {
              let _920_v55 = _compr_47;
              if (_dafny.Seq.contains(_919_v56, _920_v55)) {
                _coll47.push([_920_v55,new _dafny.CodePoint('e'.codePointAt(0))]);
              }
            }
            return _coll47;
          }()).Merge(_918_v54);
          let _lhs162 = _908_v46;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_908_v46).length));
          let _lhs164 = _911_v51;
          let _lhs165 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_911_v51).length));
          _lhs162[_lhs163] = _rhs206;
          _910_v48 = _rhs207;
          _lhs164[_lhs165] = _rhs208;
          r2 = _rhs209;
          _918_v54 = _rhs210;
          let _921_v57;
          _921_v57 = _dafny.Seq.of(_909_v47);
          let _922_v58;
          let _nw176 = Array((new BigNumber(16)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_921_v57, _921_v57);
          _nw176[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_921_v57, _dafny.Seq.of(_909_v47));
          _nw176[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_921_v57, _921_v57);
          _nw176[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_909_v47);
          _nw176[(new BigNumber(4)).toNumber()] = _921_v57;
          _nw176[(new BigNumber(5)).toNumber()] = ((p1) ? (_921_v57) : (_921_v57));
          _nw176[(new BigNumber(6)).toNumber()] = _921_v57;
          _nw176[(new BigNumber(7)).toNumber()] = _921_v57;
          _nw176[(new BigNumber(8)).toNumber()] = _921_v57;
          _nw176[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_921_v57, _921_v57);
          _nw176[(new BigNumber(10)).toNumber()] = _921_v57;
          _nw176[(new BigNumber(11)).toNumber()] = _921_v57;
          _nw176[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_909_v47), _921_v57);
          _nw176[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_909_v47);
          _nw176[(new BigNumber(14)).toNumber()] = _dafny.Seq.of((_908_v46)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_908_v46).length))], _909_v47);
          _nw176[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_921_v57, _921_v57);
          _922_v58 = _nw176;
          let _index185 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_922_v58).length));
          (_922_v58)[_index185] = _dafny.Seq.Concat(_921_v57, _921_v57);
          let _index186 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_922_v58).length));
          (_922_v58)[_index186] = _921_v57;
        } else {
          let _index187 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
          let _rhs211 = (p2).minus(p2);
          let _rhs212 = !(!_dafny.Seq.contains(_899_v37, ((_this).f22).minus((_this).f22)));
          let _lhs166 = _897_v35;
          let _lhs167 = _module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length));
          r2 = _rhs211;
          _lhs166[_lhs167] = _rhs212;
          (globalState).f2 = ((_897_v35)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length))]) && ((_dafny.MultiSet.fromElements(p0)).IsSubsetOf(_845_v2));
          let _923_v59;
          _923_v59 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_this.f23);
          let _924_v60;
          _924_v60 = _dafny.Map.Empty.slice().updateUnsafe(((_897_v35)[_module.__default.safeIndex(new BigNumber(738), new BigNumber((_897_v35).length))]) && (_this.f23),_module.D1.create_DC4(_this.f23, p2, _923_v59));
          _924_v60 = (_924_v60).Merge(_924_v60);
          _897_v35 = _897_v35;
          let _925_v61;
          _925_v61 = _dafny.Set.fromElements(_843_v0);
          let _926_v62;
          _926_v62 = _dafny.Map.Empty.slice().updateUnsafe(_843_v0,false);
          let _927_v64;
          _927_v64 = _dafny.Map.Empty.slice().updateUnsafe((_925_v61).Difference(function () {
            let _coll48 = new _dafny.Set();
            for (const _compr_48 of (_926_v62).Keys.Elements) {
              let _928_v63 = _compr_48;
              if ((_926_v62).contains(_928_v63)) {
                _coll48.add(_928_v63);
              }
            }
            return _coll48;
          }()),_897_v35);
          let _rhs213 = (p2).minus(p2);
          let _rhs214 = p2;
          let _rhs215 = (_this).f22;
          let _rhs216 = (((_927_v64).contains(_module.__default.fm27(globalState))) ? ((_927_v64).get(_module.__default.fm27(globalState))) : (_897_v35));
          let _rhs217 = (_this).f22;
          r2 = _rhs213;
          r2 = _rhs214;
          r2 = _rhs215;
          _897_v35 = _rhs216;
          r2 = _rhs217;
        }
        r2 = _module.__default.fm3(globalState);
        let _929_v65;
        _929_v65 = _dafny.Seq.of(false);
        _929_v65 = _dafny.Seq.Concat(_929_v65, _929_v65);
      } else {
        let _930_v66;
        _930_v66 = _dafny.Set.fromElements(p2);
        let _931_v67;
        _931_v67 = _dafny.Set.fromElements(new BigNumber((_930_v66).length), p2);
        let _932_v68;
        _932_v68 = _dafny.Set.fromElements(_931_v67, _931_v67, _930_v66);
        let _933_v69;
        _933_v69 = _dafny.Seq.of(_dafny.Set.fromElements(_930_v66));
        let _934_v70;
        _934_v70 = _dafny.Seq.of(new BigNumber(-236), p2);
        let _935_v71;
        let _nw177 = new _module.C1();
        _nw177.__ctor(p2, p2, _this.f23, _934_v70, new BigNumber((_844_v1).length));
        _935_v71 = _nw177;
        let _936_v72;
        _936_v72 = _dafny.Map.Empty.slice().updateUnsafe(_935_v71,(_this).f22);
        let _937_v73;
        _937_v73 = _dafny.Map.Empty.slice().updateUnsafe(_936_v72,_dafny.Set.fromElements(_dafny.Set.fromElements((_935_v71).f30, (_935_v71).f30, (_935_v71).f30)));
        let _938_v76;
        _938_v76 = _dafny.Seq.of(_930_v66);
        let _939_v79;
        let _nw178 = Array((new BigNumber(24)).toNumber());
        _nw178[(_dafny.ZERO).toNumber()] = _932_v68;
        _nw178[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_931_v67);
        _nw178[(new BigNumber(2)).toNumber()] = (_933_v69)[_module.__default.safeIndex(p2, new BigNumber((_933_v69).length))];
        _nw178[(new BigNumber(3)).toNumber()] = (_932_v68).Union(_932_v68);
        _nw178[(new BigNumber(4)).toNumber()] = (_932_v68).Union(_932_v68);
        _nw178[(new BigNumber(5)).toNumber()] = _module.__default.fm28(p1, (_this).f22, p0, _843_v0, globalState);
        _nw178[(new BigNumber(6)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(7)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(8)).toNumber()] = (((_937_v73).contains(_936_v72)) ? ((_937_v73).get(_936_v72)) : (function () {
          let _coll49 = new _dafny.Set();
          for (const _compr_49 of (_module.__default.fm29(p1, p2, p0, globalState)).Elements) {
            let _940_v74 = _compr_49;
            if (_dafny.Seq.contains(_module.__default.fm29(p1, p2, p0, globalState), _940_v74)) {
              _coll49.add(_940_v74);
            }
          }
          return _coll49;
        }()));
        _nw178[(new BigNumber(9)).toNumber()] = function () {
          let _coll50 = new _dafny.Set();
          for (const _compr_50 of (_dafny.Map.Empty.slice().updateUnsafe(_931_v67,(_935_v71).f30)).Keys.Elements) {
            let _941_v75 = _compr_50;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_931_v67,(_935_v71).f30)).contains(_941_v75)) {
              _coll50.add(_941_v75);
            }
          }
          return _coll50;
        }();
        _nw178[(new BigNumber(10)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(11)).toNumber()] = ((p0) ? (_932_v68) : (_932_v68));
        _nw178[(new BigNumber(12)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(13)).toNumber()] = (_932_v68).Difference(_dafny.Set.fromElements(_931_v67, _dafny.Set.fromElements(p2), _931_v67, _930_v66, _930_v66));
        _nw178[(new BigNumber(14)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(15)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(16)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(17)).toNumber()] = function () {
          let _coll51 = new _dafny.Set();
          for (const _compr_51 of (_938_v76).Elements) {
            let _942_v77 = _compr_51;
            if (_dafny.Seq.contains(_938_v76, _942_v77)) {
              _coll51.add(_942_v77);
            }
          }
          return _coll51;
        }();
        _nw178[(new BigNumber(18)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(19)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(20)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(21)).toNumber()] = _932_v68;
        _nw178[(new BigNumber(22)).toNumber()] = function () {
          let _coll52 = new _dafny.Set();
          for (const _compr_52 of (_dafny.Map.Empty.slice().updateUnsafe(_931_v67,p0)).Keys.Elements) {
            let _943_v78 = _compr_52;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_931_v67,p0)).contains(_943_v78)) {
              _coll52.add(_943_v78);
            }
          }
          return _coll52;
        }();
        _nw178[(new BigNumber(23)).toNumber()] = (_dafny.Set.fromElements(_930_v66)).Difference(_932_v68);
        _939_v79 = _nw178;
        let _index188 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_939_v79).length));
        (_939_v79)[_index188] = _dafny.Set.fromElements(_931_v67);
        let _index189 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_939_v79).length));
        (_939_v79)[_index189] = _module.__default.fm28(_this.f23, (_935_v71).f29, !(p0) || (p1), _843_v0, globalState);
        let _944_v80;
        _944_v80 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f22);
        let _945_v81;
        let _nw179 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
        _945_v81 = _nw179;
        let _946_v82;
        _946_v82 = _module.D2.create_DC8(_944_v80, _945_v81, _this.f23, p1, (_935_v71).f30);
        let _947_v83;
        _947_v83 = _module.D6.create_DC17(_946_v82);
        let _948_v84;
        _948_v84 = _dafny.MultiSet.fromElements(_947_v83, _module.D6.create_DC17(_946_v82));
        let _949_v85;
        _949_v85 = _dafny.Seq.of(_948_v84);
        let _950_v86;
        _950_v86 = _949_v85;
        let _951_v87;
        _951_v87 = _dafny.Map.Empty.slice().updateUnsafe(_950_v86,((p0) ? (_843_v0) : (_843_v0)));
        let _rhs218 = (_934_v70)[_module.__default.safeIndex((_this).f22, new BigNumber((_934_v70).length))];
        let _rhs219 = (((_951_v87).contains(_950_v86)) ? ((_951_v87).get(_950_v86)) : (_843_v0));
        let _lhs168 = globalState;
        r2 = _rhs218;
        _lhs168.f15 = _rhs219;
        _843_v0 = new _dafny.CodePoint('t'.codePointAt(0));
        _843_v0 = _843_v0;
        let _952_v88;
        let _nw180 = new _module.C1();
        _nw180.__ctor((_935_v71).f29, p2, p1, _934_v70, (_this).f22);
        _952_v88 = _nw180;
      }
      (globalState).f2 = !(_this.f23);
      let _953_v89;
      _953_v89 = _dafny.Map.Empty.slice().updateUnsafe(p1,_843_v0);
      let _954_v90;
      _954_v90 = _dafny.Set.fromElements(p0, !(_module.__default.fm0(false, false, globalState)));
      let _955_v91;
      let _nw181 = Array((new BigNumber(3)).toNumber());
      _nw181[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_953_v89).length), (_this).f22);
      _nw181[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(p2, p2);
      _nw181[(new BigNumber(2)).toNumber()] = new BigNumber((_954_v90).length);
      _955_v91 = _nw181;
      let _956_v92;
      let _nw182 = Array((new BigNumber(11)).toNumber()).fill(false);
      _956_v92 = _nw182;
      let _957_v93;
      _957_v93 = _dafny.Map.Empty.slice().updateUnsafe(_956_v92,_module.__default.fm3(globalState));
      let _958_v94;
      _958_v94 = _dafny.Seq.of(p1, p1);
      let _959_v95;
      _959_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(647),_dafny.Seq.of((_958_v94)[_module.__default.safeIndex(p2, new BigNumber((_958_v94).length))], p1));
      let _960_v96;
      _960_v96 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((((_957_v93).contains(_956_v92)) ? ((_957_v93).get(_956_v92)) : (new BigNumber((_959_v95).length)))),(_this).f22);
      let _index190 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_955_v91).length));
      (_955_v91)[_index190] = new BigNumber((_960_v96).length);
      let _index191 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_955_v91).length));
      (_955_v91)[_index191] = p2;
      let _961_i5;
      _961_i5 = _dafny.ZERO;
      L7: {
        while (p1) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_961_i5)) {
              break L7;
            }
            _961_i5 = (_961_i5).plus(_dafny.ONE);
            _956_v92 = _956_v92;
            let _962_v97;
            _962_v97 = _dafny.MultiSet.fromElements(_this.f23);
            _962_v97 = (_962_v97).Union(_962_v97);
            let _963_v98;
            _963_v98 = _dafny.MultiSet.fromElements((_955_v91)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_955_v91).length))], new BigNumber((_958_v94).length));
            let _964_v99;
            _964_v99 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,p1);
            let _index192 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_956_v92).length));
            (_956_v92)[_index192] = ((_963_v98).update(new BigNumber((_964_v99).length), _module.__default.abs((_dafny.ZERO).minus((_955_v91)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_955_v91).length))])))).IsDisjointFrom(_963_v98);
            let _index193 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_956_v92).length));
            (_956_v92)[_index193] = ((new BigNumber(391)).minus((_955_v91)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_955_v91).length))])).isLessThan(_module.__default.safeModuloInt((_this).f22, (_this).f22));
            let _965_v100;
            _965_v100 = _module.D4.create_DC13(p2, (_dafny.ZERO).minus((_955_v91)[_module.__default.safeIndex(new BigNumber(578), new BigNumber((_955_v91).length))]), (_this).f22);
            let _index194 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_955_v91).length));
            (_955_v91)[_index194] = (_965_v100).dtor_cf25;
          }
        }
      }
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_956_v92).length))) {
        let _966_i6 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_966_i6)) && ((_966_i6).isLessThan(new BigNumber((_956_v92).length))))) {
          (_956_v92)[(_966_i6)] = p0;
        }
      }
      let _967_v101;
      _967_v101 = _dafny.MultiSet.fromElements((_this).f22, new BigNumber(628), new BigNumber(437), (_dafny.ZERO).minus(p2), new BigNumber((_844_v1).length));
      let _968_v102;
      _968_v102 = _dafny.Set.fromElements(_967_v101);
      r0 = (function () {
        let _coll53 = new _dafny.Set();
        for (const _compr_53 of (_968_v102).Elements) {
          let _969_v103 = _compr_53;
          if ((_968_v102).contains(_969_v103)) {
            _coll53.add(_969_v103);
          }
        }
        return _coll53;
      }()).Union(((_this.f23) ? (_968_v102) : (_968_v102)));
      r1 = p1;
      r2 = (_this).f22;
      return [r0, r1, r2];
    }
    m7(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _970_v0;
      _970_v0 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_this.f23);
      let _971_v1;
      _971_v1 = _module.D1.create_DC4(_this.f23, (_this).f22, _970_v0);
      let _972_v2;
      _972_v2 = _dafny.Seq.of(_971_v1, _971_v1, _971_v1);
      _972_v2 = _dafny.Seq.Concat(_dafny.Seq.of(_971_v1), _dafny.Seq.Concat(_972_v2, _972_v2));
      let _973_v3;
      _973_v3 = _module.D8.create_DC21((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f22)), _this.f23, _this.f23, !(_this.f23), (_this).f22);
      let _974_i0;
      _974_i0 = _dafny.ZERO;
      L8: {
        while (function (_source10) {
          if (_source10.is_DC21) {
            let _979___mcc_h0 = (_source10).cf35;
            let _980___mcc_h1 = (_source10).cf36;
            let _981___mcc_h2 = (_source10).cf37;
            let _982___mcc_h3 = (_source10).cf38;
            let _983___mcc_h4 = (_source10).cf39;
            let _984_cf39 = _983___mcc_h4;
            let _985_cf38 = _982___mcc_h3;
            let _986_cf37 = _981___mcc_h2;
            let _987_cf36 = _980___mcc_h1;
            let _988_cf35 = _979___mcc_h0;
            return _986_cf37;
          } else if (_source10.is_DC22) {
            let _989___mcc_h5 = (_source10).cf40;
            let _990___mcc_h6 = (_source10).cf41;
            let _991___mcc_h7 = (_source10).cf42;
            let _992_cf42 = _991___mcc_h7;
            let _993_cf41 = _990___mcc_h6;
            let _994_cf40 = _989___mcc_h5;
            return !(_993_cf41.f16);
          } else if (_source10.is_DC23) {
            return _this.f23;
          } else if (_source10.is_DC20) {
            let _995___mcc_h8 = (_source10).cf34;
            let _996_cf34 = _995___mcc_h8;
            return _this.f23;
          } else {
            let _997___mcc_h9 = (_source10).cf43;
            let _998_cf43 = _997___mcc_h9;
            return _dafny.Seq.IsProperPrefixOf((_module.D3.create_DC9(_dafny.Seq.of((_this).f22))).dtor_cf18, _dafny.Seq.of((_this).f22, (_dafny.ZERO).minus((_this).f22)));
          }
        }(_973_v3)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_974_i0)) {
              break L8;
            }
            _974_i0 = (_974_i0).plus(_dafny.ONE);
            (_this).f23 = ((_this).f22).isLessThanOrEqualTo((_this).f22);
            let _975_v4;
            let _nw183 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
            _975_v4 = _nw183;
            let _976_v5;
            _976_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_this).f22);
            let _index195 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_975_v4).length));
            (_975_v4)[_index195] = (new BigNumber((_976_v5).length)).minus((_this).f22);
            let _977_v6;
            let _nw184 = Array((new BigNumber(20)).toNumber()).fill([]);
            _977_v6 = _nw184;
            let _978_v7;
            let _nw185 = Array((new BigNumber(2)).toNumber()).fill([]);
            _978_v7 = _nw185;
            let _index196 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_977_v6).length));
            (_977_v6)[_index196] = _978_v7;
            let _index197 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_975_v4).length));
            let _index198 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_977_v6).length));
            let _rhs220 = (_dafny.ZERO).minus((_this).f22);
            let _rhs221 = _this.f23;
            let _rhs222 = _this.f23;
            let _rhs223 = ((_this).f22).isEqualTo((_this).f22);
            let _rhs224 = _978_v7;
            let _lhs169 = _975_v4;
            let _lhs170 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_975_v4).length));
            let _lhs171 = globalState;
            let _lhs172 = globalState;
            let _lhs173 = globalState;
            let _lhs174 = _977_v6;
            let _lhs175 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_977_v6).length));
            _lhs169[_lhs170] = _rhs220;
            _lhs171.f5 = _rhs221;
            _lhs172.f2 = _rhs222;
            _lhs173.f2 = _rhs223;
            _lhs174[_lhs175] = _rhs224;
            (globalState).f2 = ((_975_v4)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_975_v4).length))]).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_this).f22, (_this).f22));
            (globalState).f15 = new _dafny.CodePoint('f'.codePointAt(0));
          }
        }
      }
      r0 = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f23,false)).length), ((_this).f22).multipliedBy(new BigNumber(858)));
      let _hi6 = (_this).f22;
      for (let _999_i1 = (_this).f22; _999_i1.isLessThan(_hi6); _999_i1 = _999_i1.plus(_dafny.ONE)) {
        r2 = !(_this.f23) || (_this.f23);
        r2 = _this.f23;
        let _1000_v8;
        _1000_v8 = _dafny.Seq.of(_this.f23);
        _1000_v8 = _dafny.Seq.of(_dafny.Seq.IsPrefixOf(_1000_v8, _dafny.Seq.update(_1000_v8, _module.__default.safeIndex(_module.__default.fm3(globalState), new BigNumber((_1000_v8).length)), _this.f23)));
        let _1001_v9;
        let _nw186 = Array((new BigNumber(7)).toNumber());
        _nw186[(_dafny.ZERO).toNumber()] = _973_v3;
        _nw186[(_dafny.ONE).toNumber()] = _973_v3;
        _nw186[(new BigNumber(2)).toNumber()] = _module.D8.create_DC21((_this).f22, _this.f23, _this.f23, true, _999_i1);
        _nw186[(new BigNumber(3)).toNumber()] = _973_v3;
        _nw186[(new BigNumber(4)).toNumber()] = _973_v3;
        _nw186[(new BigNumber(5)).toNumber()] = _973_v3;
        _nw186[(new BigNumber(6)).toNumber()] = _973_v3;
        _1001_v9 = _nw186;
        let _1002_v10;
        let _nw187 = new _module.C2();
        _nw187.__ctor((_this).f22);
        _1002_v10 = _nw187;
        let _1003_v11;
        _1003_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_1002_v10);
        let _pat_let_tv23 = _1003_v11;
        let _index199 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1001_v9).length));
        (_1001_v9)[_index199] = ((true) ? (function (_pat_let27_0) {
          return function (_1004_dt__update__tmp_h0) {
            return function (_pat_let28_0) {
              return function (_1005_dt__update_hcf35_h0) {
                return _module.D8.create_DC21(_1005_dt__update_hcf35_h0, (_1004_dt__update__tmp_h0).dtor_cf36, (_1004_dt__update__tmp_h0).dtor_cf37, (_1004_dt__update__tmp_h0).dtor_cf38, (_1004_dt__update__tmp_h0).dtor_cf39);
              }(_pat_let28_0);
            }(new BigNumber((_pat_let_tv23).length));
          }(_pat_let27_0);
        }(_973_v3)) : (_module.__default.fm26(globalState)));
        let _1006_v12;
        let _nw188 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _1006_v12 = _nw188;
        let _index200 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1006_v12).length));
        (_1006_v12)[_index200] = ((_this.f23) ? ((_1002_v10).f25) : ((_dafny.ZERO).minus((_1002_v10).f25)));
        let _1007_v13;
        _1007_v13 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(_this.f23), _1000_v8));
        let _1008_v14;
        _1008_v14 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,new BigNumber(491));
        let _1009_v15;
        _1009_v15 = _dafny.Map.Empty.slice().updateUnsafe(_999_i1,_this.f23);
        let _1010_v16;
        let _nw189 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _1010_v16 = _nw189;
        let _1011_v17;
        _1011_v17 = _module.D2.create_DC8(_1008_v14, _1010_v16, _this.f23, _this.f23, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("qjqdbdskh")).length)));
        let _1012_v18;
        let _nw190 = Array((new BigNumber(25)).toNumber());
        _nw190[(_dafny.ZERO).toNumber()] = _this.f23;
        _nw190[(_dafny.ONE).toNumber()] = _this.f23;
        _nw190[(new BigNumber(2)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(3)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(4)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(5)).toNumber()] = !(_this.f23);
        _nw190[(new BigNumber(6)).toNumber()] = !(_1008_v14).equals(_1008_v14);
        _nw190[(new BigNumber(7)).toNumber()] = false;
        _nw190[(new BigNumber(8)).toNumber()] = (((_1000_v8)[_module.__default.safeIndex((_1002_v10).f25, new BigNumber((_1000_v8).length))]) ? (_this.f23) : ((((_1009_v15).contains(_999_i1)) ? ((_1009_v15).get(_999_i1)) : ((_this).fm9(_999_i1, globalState)))));
        _nw190[(new BigNumber(9)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(10)).toNumber()] = false;
        _nw190[(new BigNumber(11)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(12)).toNumber()] = !(new BigNumber(499)).isEqualTo(_999_i1);
        _nw190[(new BigNumber(13)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(14)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(15)).toNumber()] = !(_this.f23) || (_this.f23);
        _nw190[(new BigNumber(16)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(17)).toNumber()] = true;
        _nw190[(new BigNumber(18)).toNumber()] = !((new BigNumber(59)).isLessThan(_999_i1));
        _nw190[(new BigNumber(19)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(20)).toNumber()] = ((_this.f23) ? ((_1011_v17).dtor_cf15) : (_this.f23));
        _nw190[(new BigNumber(21)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(22)).toNumber()] = !(!(_this.f23));
        _nw190[(new BigNumber(23)).toNumber()] = _this.f23;
        _nw190[(new BigNumber(24)).toNumber()] = _this.f23;
        _1012_v18 = _nw190;
        let _1013_v19;
        _1013_v19 = _dafny.MultiSet.fromElements(_999_i1);
        let _1014_v20;
        _1014_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1012_v18,_1012_v18);
        let _index201 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1001_v9).length));
        let _index202 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1006_v12).length));
        let _rhs225 = _module.D8.create_DC21((_dafny.ZERO).minus(((true) ? ((_dafny.ZERO).minus((_this).f22)) : ((_dafny.ZERO).minus(_999_i1)))), !(_1013_v19).equals(_1013_v19), _this.f23, _this.f23, _999_i1);
        let _rhs226 = (_this).f22;
        let _rhs227 = _999_i1;
        let _rhs228 = _1007_v13;
        let _rhs229 = (((_1014_v20).contains(_1012_v18)) ? ((_1014_v20).get(_1012_v18)) : (_1012_v18));
        let _lhs176 = _1001_v9;
        let _lhs177 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1001_v9).length));
        let _lhs178 = _1006_v12;
        let _lhs179 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1006_v12).length));
        _lhs176[_lhs177] = _rhs225;
        r0 = _rhs226;
        _lhs178[_lhs179] = _rhs227;
        _1007_v13 = _rhs228;
        _1012_v18 = _rhs229;
      }
      let _1015_v21;
      _1015_v21 = _module.D4.create_DC12(_this.f23, _this.f23, (_this).f22);
      let _source11 = _1015_v21;
      if (_source11.is_DC12) {
        let _1016___mcc_h10 = (_source11).cf21;
        let _1017___mcc_h11 = (_source11).cf22;
        let _1018___mcc_h12 = (_source11).cf23;
        let _1019_cf23 = _1018___mcc_h12;
        let _1020_cf22 = _1017___mcc_h11;
        let _1021_cf21 = _1016___mcc_h10;
        let _1022_v22;
        _1022_v22 = _dafny.Seq.UnicodeFromString("eh");
        r1 = _1022_v22;
        let _1023_v23;
        let _nw191 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
        _1023_v23 = _nw191;
        let _1024_v24;
        _1024_v24 = _dafny.Seq.of(!(_1021_cf21));
        let _index203 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1023_v23).length));
        (_1023_v23)[_index203] = _1024_v24;
        let _index204 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_1023_v23).length));
        (_1023_v23)[_index204] = _1024_v24;
        r2 = !_dafny.areEqual(_1022_v22, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("x"), _1022_v22));
        let _1025_v25;
        _1025_v25 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f22);
        _1025_v25 = (_1025_v25).update((_1021_cf21) && (true), (_this).f22);
      } else if (_source11.is_DC13) {
        let _1026___mcc_h13 = (_source11).cf24;
        let _1027___mcc_h14 = (_source11).cf25;
        let _1028___mcc_h15 = (_source11).cf26;
        let _1029_cf26 = _1028___mcc_h15;
        let _1030_cf25 = _1027___mcc_h14;
        let _1031_cf24 = _1026___mcc_h13;
        let _1032_v26;
        _1032_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1030_cf25,false);
        let _1033_v27;
        _1033_v27 = _dafny.MultiSet.fromElements(_1031_cf24, new BigNumber((_1032_v26).length));
        let _1034_v28;
        _1034_v28 = _dafny.Seq.of(new BigNumber((_module.__default.fm23(_1033_v27, _this.f23, globalState)).length), new BigNumber((_1032_v26).length));
        let _1035_v29;
        let _nw192 = Array((new BigNumber(28)).toNumber()).fill(false);
        _1035_v29 = _nw192;
        let _1036_v30;
        let _nw193 = new _module.C0();
        _nw193.__ctor(new BigNumber((_dafny.Seq.Concat(_1034_v28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), ((_1037_cf25) => function (_1038_i2) {
          return _1037_cf25;
        })(_1030_cf25)))).length), _1035_v29);
        _1036_v30 = _nw193;
        (globalState).f2 = _this.f23;
        (globalState).f2 = _this.f23;
        if (_this.f23) {
          let _1039_v31;
          _1039_v31 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_1031_cf24);
          let _1040_v32;
          _1040_v32 = _module.D3.create_DC9(_dafny.Seq.of(new BigNumber((_1039_v31).length), (_1036_v30).f26, (_1036_v30).f26));
          let _pat_let_tv24 = _1034_v28;
          let _pat_let_tv25 = _1034_v28;
          let _pat_let_tv26 = _1034_v28;
          let _1041_v33;
          let _nw194 = Array((new BigNumber(15)).toNumber());
          _nw194[(_dafny.ZERO).toNumber()] = _1040_v32;
          _nw194[(_dafny.ONE).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(2)).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(3)).toNumber()] = function (_pat_let29_0) {
            return function (_1042_dt__update__tmp_h1) {
              return function (_pat_let30_0) {
                return function (_1043_dt__update_hcf18_h0) {
                  return _module.D3.create_DC9(_1043_dt__update_hcf18_h0);
                }(_pat_let30_0);
              }(_pat_let_tv24);
            }(_pat_let29_0);
          }(_1040_v32);
          _nw194[(new BigNumber(4)).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(5)).toNumber()] = function (_pat_let31_0) {
            return function (_1044_dt__update__tmp_h2) {
              return function (_pat_let32_0) {
                return function (_1045_dt__update_hcf18_h1) {
                  return _module.D3.create_DC9(_1045_dt__update_hcf18_h1);
                }(_pat_let32_0);
              }(_pat_let_tv25);
            }(_pat_let31_0);
          }(_1040_v32);
          _nw194[(new BigNumber(6)).toNumber()] = function (_pat_let33_0) {
            return function (_1046_dt__update__tmp_h3) {
              return function (_pat_let34_0) {
                return function (_1047_dt__update_hcf18_h2) {
                  return _module.D3.create_DC9(_1047_dt__update_hcf18_h2);
                }(_pat_let34_0);
              }(_pat_let_tv26);
            }(_pat_let33_0);
          }(_1040_v32);
          _nw194[(new BigNumber(7)).toNumber()] = ((_this.f23) ? (_1040_v32) : (_1040_v32));
          _nw194[(new BigNumber(8)).toNumber()] = ((_this.f23) ? (_1040_v32) : (_1040_v32));
          _nw194[(new BigNumber(9)).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(10)).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(11)).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(12)).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(13)).toNumber()] = _1040_v32;
          _nw194[(new BigNumber(14)).toNumber()] = _1040_v32;
          _1041_v33 = _nw194;
          let _index205 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1041_v33).length));
          (_1041_v33)[_index205] = _1040_v32;
          let _index206 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1041_v33).length));
          (_1041_v33)[_index206] = _1040_v32;
          (_1036_v30).f27 = ((_dafny.areEqual(_module.__default.fm17(globalState), _1034_v28)) ? (_1035_v29) : (_1035_v29));
          let _1048_v34;
          _1048_v34 = _dafny.Seq.UnicodeFromString("bglc");
          let _1049_v35;
          let _nw195 = new _module.C3();
          _nw195.__ctor(_1048_v34, _this.f23, _dafny.Seq.Concat(_1034_v28, _1034_v28));
          _1049_v35 = _nw195;
          _1049_v35 = _1049_v35;
          (globalState).f2 = _this.f23;
          let _1050_v36;
          let _nw196 = new _module.C0();
          _nw196.__ctor(_1031_cf24, _1035_v29);
          _1050_v36 = _nw196;
        } else {
          r3 = (_1030_cf25).minus((_1036_v30).f26);
          r1 = _module.__default.fm1(globalState);
          let _1051_v37;
          let _nw197 = new _module.C2();
          _nw197.__ctor((_1036_v30).f26);
          _1051_v37 = _nw197;
          _1031_cf24 = (_1036_v30).f26;
          let _1052_v38;
          let _nw198 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _1052_v38 = _nw198;
          let _index207 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1052_v38).length));
          (_1052_v38)[_index207] = _1034_v28;
          let _index208 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1052_v38).length));
          (_1052_v38)[_index208] = _1034_v28;
        }
      } else {
        let _1053___mcc_h16 = (_source11).cf20;
        let _1054_cf20 = _1053___mcc_h16;
        (globalState).f2 = _this.f23;
        let _1055_v39;
        _1055_v39 = _dafny.Set.fromElements(_this.f23);
        if (((_dafny.Set.fromElements(_this.f23, _this.f23)).Union(_1055_v39)).IsSubsetOf(_1055_v39)) {
          let _1056_v40;
          let _nw199 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1056_v40 = _nw199;
          let _nw200 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1056_v40 = _nw200;
          let _1057_v41;
          _1057_v41 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f22);
          let _1058_v42;
          _1058_v42 = _dafny.Seq.UnicodeFromString("ap");
          _1057_v41 = (_1057_v41).update((_this).fm9((_this).f22, globalState), new BigNumber((_1058_v42).length));
          r0 = (_this).f22;
          let _rhs230 = true;
          let _rhs231 = _1056_v40;
          let _rhs232 = (_this).f22;
          let _rhs233 = _this.f23;
          let _lhs180 = globalState;
          let _lhs181 = globalState;
          _lhs180.f5 = _rhs230;
          _1056_v40 = _rhs231;
          r0 = _rhs232;
          _lhs181.f2 = _rhs233;
          r2 = _this.f23;
        } else {
          r2 = _this.f23;
          let _1059_v43;
          _1059_v43 = _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("o"));
          let _1060_v44;
          _1060_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1059_v43,_this.f23);
          let _1061_v46;
          _1061_v46 = _dafny.Set.fromElements(_1059_v43);
          let _1062_v47;
          _1062_v47 = _dafny.Set.fromElements(function () {
            let _coll54 = new _dafny.Map();
            for (const _compr_54 of (_1061_v46).Elements) {
              let _1063_v45 = _compr_54;
              if ((_1061_v46).contains(_1063_v45)) {
                _coll54.push([_1063_v45,_this.f23]);
              }
            }
            return _coll54;
          }());
          (_this).f23 = !(_1062_v47).contains(_1060_v44);
          let _1064_v48;
          _1064_v48 = _module.D1.create_DC5((new BigNumber(35)).minus(new BigNumber(174)), _this.f23);
          let _1065_v49;
          _1065_v49 = _dafny.MultiSet.fromElements(((true) ? (_this.f23) : (_this.f23)));
          let _1066_v50;
          _1066_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(222),_this.f23);
          let _1067_v51;
          _1067_v51 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_this).f22);
          let _rhs234 = _module.__default.fm0(false, !(_1066_v50).equals((_1066_v50).update((_this).f22, _this.f23)), globalState);
          let _rhs235 = _1064_v48;
          let _rhs236 = (_1065_v49).update(_this.f23, _module.__default.abs((((_1067_v51).contains(_this.f23)) ? ((_1067_v51).get(_this.f23)) : ((_this).f22))));
          let _lhs182 = globalState;
          _lhs182.f2 = _rhs234;
          _1064_v48 = _rhs235;
          _1065_v49 = _rhs236;
          let _1068_v52;
          let _nw201 = Array((new BigNumber(3)).toNumber());
          _nw201[(_dafny.ZERO).toNumber()] = _this.f23;
          _nw201[(_dafny.ONE).toNumber()] = _this.f23;
          _nw201[(new BigNumber(2)).toNumber()] = (_this.f23) && (_this.f23);
          _1068_v52 = _nw201;
          _1068_v52 = ((_this.f23) ? (_1068_v52) : (_1068_v52));
          let _1069_v53;
          let _nw202 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _1069_v53 = _nw202;
          let _index209 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_1069_v53).length));
          (_1069_v53)[_index209] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_this.f23)).length);
          let _1070_v54;
          let _nw203 = new _module.C2();
          _nw203.__ctor(new BigNumber((_970_v0).length));
          _1070_v54 = _nw203;
          let _1071_v55;
          let _nw204 = Array((new BigNumber(5)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1071_v55 = _nw204;
          let _1072_v56;
          _1072_v56 = _dafny.MultiSet.fromElements(_module.D1.create_DC5((_1070_v54).f25, _this.f23), _1064_v48);
          let _index210 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1071_v55).length));
          (_1071_v55)[_index210] = _1072_v56;
          let _1073_v57;
          _1073_v57 = _dafny.Seq.UnicodeFromString("fs");
          let _index211 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_1069_v53).length));
          let _index212 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1071_v55).length));
          let _rhs237 = (_this).f22;
          let _rhs238 = _1070_v54;
          let _rhs239 = (_module.__default.fm30(_this.f23, _dafny.Seq.UnicodeFromString("ut"), true, _this.f23, globalState)).Difference(_1072_v56);
          let _rhs240 = _this.f23;
          let _rhs241 = _module.__default.safeDivisionInt(new BigNumber((_1073_v57).length), ((_dafny.ZERO).minus((_1070_v54).f25)).minus(new BigNumber((_dafny.Seq.of(_this.f23, !(_this.f23))).length)));
          let _lhs183 = _1069_v53;
          let _lhs184 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_1069_v53).length));
          let _lhs185 = _1071_v55;
          let _lhs186 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1071_v55).length));
          let _lhs187 = _this;
          _lhs183[_lhs184] = _rhs237;
          _1070_v54 = _rhs238;
          _lhs185[_lhs186] = _rhs239;
          _lhs187.f23 = _rhs240;
          r0 = _rhs241;
        }
        let _1074_v58;
        _1074_v58 = _dafny.Seq.UnicodeFromString("gao");
        let _1075_v59;
        _1075_v59 = new _dafny.CodePoint('x'.codePointAt(0));
        _module.__default.m0(new BigNumber((_dafny.Seq.update(((_this.f23) ? (_1074_v58) : (_1074_v58)), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f22), new BigNumber((((_this.f23) ? (_1074_v58) : (_1074_v58))).length)), _1075_v59)).length), _1074_v58, (_this).f22, _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_this.f23)).length), new BigNumber(992)), globalState);
        let _1076_v60;
        _1076_v60 = _dafny.Seq.of((_this).f22);
        let _rhs242 = (_this).f22;
        let _rhs243 = _this.f23;
        let _rhs244 = new _dafny.CodePoint('w'.codePointAt(0));
        let _rhs245 = ((_this.f23) ? ((_this).f22) : ((_this).f22));
        let _rhs246 = _dafny.Seq.Concat(_1076_v60, _dafny.Seq.Concat(_1076_v60, _1076_v60));
        let _lhs188 = globalState;
        let _lhs189 = globalState;
        r0 = _rhs242;
        _lhs188.f2 = _rhs243;
        _lhs189.f15 = _rhs244;
        r0 = _rhs245;
        _1076_v60 = _rhs246;
      }
      let _1077_v61;
      let _nw205 = Array((new BigNumber(29)).toNumber()).fill(_module.D2.Default());
      _1077_v61 = _nw205;
      let _1078_v62;
      _1078_v62 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,(_this).f22);
      let _1079_v63;
      let _nw206 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
      _1079_v63 = _nw206;
      let _1080_v64;
      _1080_v64 = _module.D2.create_DC8(_1078_v62, _1079_v63, _this.f23, _this.f23, (_this).f22);
      let _index213 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1077_v61).length));
      (_1077_v61)[_index213] = _1080_v64;
      let _index214 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1077_v61).length));
      (_1077_v61)[_index214] = _1080_v64;
      r0 = new BigNumber(-398);
      let _1081_v65;
      _1081_v65 = _dafny.Seq.UnicodeFromString("m");
      r1 = _1081_v65;
      r2 = _this.f23;
      let _1082_v66;
      _1082_v66 = _dafny.Seq.of(_this.f23, _this.f23, _this.f23, (_this.f23) === (_this.f23), (((_970_v0).contains(false)) ? ((_970_v0).get(false)) : (_this.f23)));
      r3 = new BigNumber((_1082_v66).length);
      return [r0, r1, r2, r3];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f16 = false;
      this._f17 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f16, f17) {
      let _this = this;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm7(p0, globalState) {
      let _this = this;
      if (_this.f16) {
        return _dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)));
      } else {
        return (_module.D0.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(117)), function (_1083_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}))).dtor_cf1;
      }
    };
    fm8(p0, globalState) {
      let _this = this;
      return _dafny.Seq.IsPrefixOf((_this).f17, (_module.D3.create_DC9((_this).f17)).dtor_cf18);
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      (_this).f16 = !(p2) || (_module.__default.fm0(p2, p1, globalState));
      let _1084_v0;
      let _nw207 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
      _1084_v0 = _nw207;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1084_v0).length))) {
        let _1085_i0 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1085_i0)) && ((_1085_i0).isLessThan(new BigNumber((_1084_v0).length))))) {
          (_1084_v0)[(_1085_i0)] = (_1085_i0).plus(new BigNumber(8));
        }
      }
      let _1086_v1;
      let _nw208 = new _module.C4();
      _nw208.__ctor(p3, true);
      _1086_v1 = _nw208;
      let _1087_v2;
      let _nw209 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
      _1087_v2 = _nw209;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1087_v2).length))) {
        let _1088_i1 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1088_i1)) && ((_1088_i1).isLessThan(new BigNumber((_1087_v2).length))))) {
          (_1087_v2)[(_1088_i1)] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-728)), function (_1089_i2) {
            return new BigNumber(59);
          });
        }
      }
      let _1090_v3;
      let _nw210 = Array((new BigNumber(11)).toNumber()).fill(_module.D0.Default());
      _1090_v3 = _nw210;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1090_v3).length))) {
        let _1091_i3 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1091_i3)) && ((_1091_i3).isLessThan(new BigNumber((_1090_v3).length))))) {
          (_1090_v3)[(_1091_i3)] = _module.D0.create_DC0(p0);
        }
      }
      let _hi7 = (_1086_v1).f22;
      for (let _1092_i4 = p3; _1092_i4.isLessThan(_hi7); _1092_i4 = _1092_i4.plus(_dafny.ONE)) {
        let _1093_v4;
        _1093_v4 = new BigNumber(577);
        _1093_v4 = p3;
        let _1094_v5;
        let _1095_v6;
        let _1096_v7;
        let _1097_v8;
        let _out6;
        let _out7;
        let _out8;
        let _out9;
        let _outcollector2 = (_1086_v1).m7(globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _out9 = _outcollector2[3];
        _1094_v5 = _out6;
        _1095_v6 = _out7;
        _1096_v7 = _out8;
        _1097_v8 = _out9;
        let _1098_v9;
        let _nw211 = Array((new BigNumber(14)).toNumber()).fill(false);
        _1098_v9 = _nw211;
        let _index215 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1098_v9).length));
        (_1098_v9)[_index215] = p2;
        let _index216 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_1084_v0).length));
        (_1084_v0)[_index216] = (_1086_v1).f22;
        let _1099_v10;
        _1099_v10 = _dafny.Set.fromElements(_1092_i4);
        let _1100_v11;
        let _nw212 = Array((new BigNumber(6)).toNumber());
        _nw212[(_dafny.ZERO).toNumber()] = _1099_v10;
        _nw212[(_dafny.ONE).toNumber()] = _1099_v10;
        _nw212[(new BigNumber(2)).toNumber()] = _module.__default.fm31(_this.f16, globalState);
        _nw212[(new BigNumber(3)).toNumber()] = _module.__default.fm31(p2, globalState);
        _nw212[(new BigNumber(4)).toNumber()] = _module.__default.fm31(_1086_v1.f23, globalState);
        _nw212[(new BigNumber(5)).toNumber()] = (_1099_v10).Intersect(_1099_v10);
        _1100_v11 = _nw212;
        let _index217 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1100_v11).length));
        (_1100_v11)[_index217] = _1099_v10;
        let _1101_v12;
        _1101_v12 = new _dafny.CodePoint('h'.codePointAt(0));
        let _index218 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1098_v9).length));
        let _index219 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_1084_v0).length));
        let _index220 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1100_v11).length));
        let _rhs247 = _this.f16;
        let _rhs248 = _1101_v12;
        let _rhs249 = new BigNumber(716);
        let _rhs250 = _1099_v10;
        let _lhs190 = _1098_v9;
        let _lhs191 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1098_v9).length));
        let _lhs192 = globalState;
        let _lhs193 = _1084_v0;
        let _lhs194 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_1084_v0).length));
        let _lhs195 = _1100_v11;
        let _lhs196 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_1100_v11).length));
        _lhs190[_lhs191] = _rhs247;
        _lhs192.f15 = _rhs248;
        _lhs193[_lhs194] = _rhs249;
        _lhs195[_lhs196] = _rhs250;
        let _1102_v13;
        _1102_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1086_v1).f22,_1094_v5);
        let _1103_v14;
        let _nw213 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1103_v14 = _nw213;
        let _1104_v15;
        _1104_v15 = _dafny.Seq.of(_this.f16);
        let _1105_v16;
        _1105_v16 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f16),_1103_v14);
        let _1106_v17;
        _1106_v17 = _dafny.Map.Empty.slice().updateUnsafe(p3,(((_1105_v16).contains(p1)) ? ((_1105_v16).get(p1)) : (_1103_v14)));
        let _index221 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1098_v9).length));
        let _rhs251 = ((_1104_v15)[_module.__default.safeIndex(_1097_v8, new BigNumber((_1104_v15).length))]) && ((_this).fm8(_dafny.Seq.UnicodeFromString("vjxld"), globalState));
        let _rhs252 = _dafny.Map.Empty.slice().updateUnsafe(_1097_v8,(_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1094_v5, _1093_v4))));
        let _rhs253 = (((_1106_v17).contains(_1092_i4)) ? ((_1106_v17).get(_1092_i4)) : (((p2) ? (_1103_v14) : (_1103_v14))));
        let _lhs197 = _1098_v9;
        let _lhs198 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1098_v9).length));
        _lhs197[_lhs198] = _rhs251;
        _1102_v13 = _rhs252;
        _1103_v14 = _rhs253;
      }
      let _1107_v18;
      _1107_v18 = _dafny.Seq.of(_this.f16);
      r0 = _1107_v18;
      return r0;
    }
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      if (!(_this.f16) || (_this.f16)) {
        if (_this.f16) {
          let _1108_v0;
          _1108_v0 = _dafny.Seq.of(_module.__default.fm19(new BigNumber(125), globalState));
          let _1109_v1;
          _1109_v1 = new BigNumber(42);
          let _1110_v2;
          _1110_v2 = _dafny.Seq.of(_this.f16, !(_this.f16), _this.f16, _this.f16);
          let _1111_v3;
          _1111_v3 = _dafny.MultiSet.fromElements(_this.f16, _this.f16);
          let _1112_v4;
          _1112_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1111_v3).cardinality()),!(_this.f16));
          let _1113_v5;
          _1113_v5 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1114_v6;
          _1114_v6 = _dafny.MultiSet.fromElements(_1113_v5);
          let _1115_v7;
          _1115_v7 = _dafny.Map.Empty.slice().updateUnsafe(true,_1109_v1);
          let _1116_v8;
          let _nw214 = Array((new BigNumber(28)).toNumber());
          _nw214[(_dafny.ZERO).toNumber()] = _1109_v1;
          _nw214[(_dafny.ONE).toNumber()] = new BigNumber(203);
          _nw214[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt(_1109_v1, _1109_v1);
          _nw214[(new BigNumber(3)).toNumber()] = (_1109_v1).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_1109_v1)).length)));
          _nw214[(new BigNumber(4)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(5)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(6)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(7)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(8)).toNumber()] = (_1109_v1).minus(new BigNumber((_dafny.Seq.of(_this.f16, _this.f16)).length));
          _nw214[(new BigNumber(9)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(10)).toNumber()] = new BigNumber((_1110_v2).length);
          _nw214[(new BigNumber(11)).toNumber()] = (new BigNumber(187)).plus(_1109_v1);
          _nw214[(new BigNumber(12)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(13)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(14)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(15)).toNumber()] = (new BigNumber((_1112_v4).length)).multipliedBy(new BigNumber(-267));
          _nw214[(new BigNumber(16)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(17)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(18)).toNumber()] = new BigNumber((_1111_v3).cardinality());
          _nw214[(new BigNumber(19)).toNumber()] = (_1109_v1).minus(_1109_v1);
          _nw214[(new BigNumber(20)).toNumber()] = (_1109_v1).minus(((_this).f17)[_module.__default.safeIndex(_1109_v1, new BigNumber(((_this).f17).length))]);
          _nw214[(new BigNumber(21)).toNumber()] = new BigNumber(-718);
          _nw214[(new BigNumber(22)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(23)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(24)).toNumber()] = _1109_v1;
          _nw214[(new BigNumber(25)).toNumber()] = _module.__default.safeDivisionInt((((_1114_v6).contains(_1113_v5)) ? ((_1114_v6).get(_1113_v5)) : (_1109_v1)), new BigNumber((_1115_v7).length));
          _nw214[(new BigNumber(26)).toNumber()] = (_1109_v1).multipliedBy(new BigNumber(((_this).f17).length));
          _nw214[(new BigNumber(27)).toNumber()] = _1109_v1;
          _1116_v8 = _nw214;
          let _index222 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length));
          (_1116_v8)[_index222] = _1109_v1;
          let _1117_v9;
          _1117_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1110_v2);
          let _index223 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length));
          let _rhs254 = _dafny.Seq.Concat(((_this.f16) ? (_1108_v0) : (_1108_v0)), _module.__default.fm32(new BigNumber(102), false, _this.f16, _1113_v5, globalState));
          let _rhs255 = ((new BigNumber((_1117_v9).length)).multipliedBy(_1109_v1)).minus(_module.__default.fm3(globalState));
          let _lhs199 = _1116_v8;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length));
          _1108_v0 = _rhs254;
          _lhs199[_lhs200] = _rhs255;
          let _index224 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length));
          (_1116_v8)[_index224] = (_1116_v8)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length))];
          let _1118_v10;
          _1118_v10 = _dafny.Seq.UnicodeFromString("klmwlqbvh");
          _module.__default.m0(_1109_v1, _1118_v10, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nvw"), _1118_v10)).length), _module.__default.safeDivisionInt((_1116_v8)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length))], (_1116_v8)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length))]), globalState);
          let _index225 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length));
          (_1116_v8)[_index225] = ((_1116_v8)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length))]).multipliedBy((_1109_v1).multipliedBy((_1116_v8)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length))]));
          let _index226 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1116_v8).length));
          (_1116_v8)[_index226] = (_dafny.ZERO).minus(_1109_v1);
        } else {
          let _1119_v11;
          _1119_v11 = _dafny.ZERO;
          let _1120_v12;
          _1120_v12 = _dafny.Seq.UnicodeFromString("sf");
          _1119_v11 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1120_v12, _module.__default.fm1(globalState)), _1120_v12)).length);
          _1119_v11 = _1119_v11;
          _1119_v11 = _1119_v11;
          let _1121_v13;
          _1121_v13 = _module.D0.create_DC1(_dafny.Seq.UnicodeFromString("xvkvrqxb"));
          let _1122_v14;
          _1122_v14 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1123_v15;
          _1123_v15 = _dafny.MultiSet.fromElements(new BigNumber((((_this.f16) ? (_1120_v12) : (_dafny.Seq.update((_1121_v13).dtor_cf1, _module.__default.safeIndex(new BigNumber(819), new BigNumber(((_1121_v13).dtor_cf1).length)), _1122_v14)))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(185)), ((_1124_v11) => function (_1125_i0) {
            return _1124_v11;
          })(_1119_v11))).length));
          _1123_v15 = _dafny.MultiSet.fromElements((_1119_v11).multipliedBy(_1119_v11), _1119_v11, _1119_v11);
          let _1126_v16;
          let _nw215 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _1126_v16 = _nw215;
          _1126_v16 = ((_this.f16) ? (_1126_v16) : (_1126_v16));
        }
        let _source12 = _module.__default.fm33((_this).f17, globalState);
        if (_source12.is_DC4) {
          let _1127___mcc_h0 = (_source12).cf7;
          let _1128___mcc_h1 = (_source12).cf8;
          let _1129___mcc_h2 = (_source12).cf9;
          let _1130_cf9 = _1129___mcc_h2;
          let _1131_cf8 = _1128___mcc_h1;
          let _1132_cf7 = _1127___mcc_h0;
          let _1133_v17;
          _1133_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1132_cf7,new BigNumber(-752));
          _1133_v17 = _1133_v17;
          let _1134_v18;
          _1134_v18 = _dafny.Seq.UnicodeFromString("m");
          _1131_cf8 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1134_v18, _1134_v18), _dafny.Seq.UnicodeFromString("xlh"))).length);
          let _1135_v19;
          _1135_v19 = _dafny.MultiSet.fromElements(_1131_cf8);
          let _1136_v20;
          let _nw216 = new _module.C1();
          _nw216.__ctor(_1131_cf8, _1131_cf8, (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1131_cf8))).IsDisjointFrom(_1135_v19), _dafny.Seq.Concat((_this).f17, (_this).f17), _module.__default.safeDivisionInt(_1131_cf8, _1131_cf8));
          _1136_v20 = _nw216;
          let _1137_v21;
          let _init30 = ((_1138_v18) => function (_1139_i1) {
            return _1138_v18;
          })(_1134_v18);
          let _nw217 = Array((new BigNumber(8)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw217.length); _i0_30++) {
            _nw217[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1137_v21 = _nw217;
          let _index227 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1137_v21).length));
          (_1137_v21)[_index227] = _dafny.Seq.Concat(_1134_v18, _1134_v18);
          let _index228 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1137_v21).length));
          let _rhs256 = _1136_v20;
          let _rhs257 = _1134_v18;
          let _lhs201 = _1137_v21;
          let _lhs202 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_1137_v21).length));
          _1136_v20 = _rhs256;
          _lhs201[_lhs202] = _rhs257;
          let _1140_v22;
          _1140_v22 = _module.D1.create_DC5((((_1133_v17).contains(_1132_cf7)) ? ((_1133_v17).get(_1132_cf7)) : (new BigNumber(((_1137_v21)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_1137_v21).length))]).length))), _1132_cf7);
          _module.__default.m0(_1131_cf8, _1134_v18, _module.__default.safeDivisionInt((_1140_v22).dtor_cf10, (_1136_v20).f30), _1131_cf8, globalState);
        } else if (_source12.is_DC5) {
          let _1141___mcc_h3 = (_source12).cf10;
          let _1142___mcc_h4 = (_source12).cf11;
          let _1143_cf11 = _1142___mcc_h4;
          let _1144_cf10 = _1141___mcc_h3;
          let _1145_v23;
          _1145_v23 = _dafny.Seq.of(_1143_cf11);
          let _1146_v24;
          _1146_v24 = _dafny.Seq.UnicodeFromString("rnjvqpd");
          let _1147_v25;
          let _nw218 = new _module.C3();
          _nw218.__ctor(_1146_v24, _this.f16, (_this).f17);
          _1147_v25 = _nw218;
          let _1148_v26;
          _1148_v26 = _dafny.MultiSet.fromElements(_1144_cf10, new BigNumber((_dafny.Set.fromElements(_1147_v25)).length));
          let _1149_v27;
          _1149_v27 = _dafny.Set.fromElements(_module.__default.fm3(globalState));
          let _1150_v28;
          _1150_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1143_cf11,_1144_cf10);
          let _1151_v29;
          let _nw219 = Array((new BigNumber(21)).toNumber());
          _nw219[(_dafny.ZERO).toNumber()] = _1144_cf10;
          _nw219[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1144_cf10);
          _nw219[(new BigNumber(2)).toNumber()] = new BigNumber((_1145_v23).length);
          _nw219[(new BigNumber(3)).toNumber()] = new BigNumber((_1146_v24).length);
          _nw219[(new BigNumber(4)).toNumber()] = new BigNumber(575);
          _nw219[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_1144_cf10, new BigNumber(((_this).f17).length))).cardinality());
          _nw219[(new BigNumber(6)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(7)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(8)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(9)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(10)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(11)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(12)).toNumber()] = new BigNumber((_1145_v23).length);
          _nw219[(new BigNumber(13)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(14)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(15)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(16)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(17)).toNumber()] = (((_1148_v26).contains(new BigNumber(396))) ? ((_1148_v26).get(new BigNumber(396))) : (new BigNumber((_1149_v27).length)));
          _nw219[(new BigNumber(18)).toNumber()] = new BigNumber((_1146_v24).length);
          _nw219[(new BigNumber(19)).toNumber()] = _1144_cf10;
          _nw219[(new BigNumber(20)).toNumber()] = (((_1150_v28).contains(_this.f16)) ? ((_1150_v28).get(_this.f16)) : (_1144_cf10));
          _1151_v29 = _nw219;
          let _1152_v30;
          _1152_v30 = _dafny.MultiSet.fromElements(((_1143_cf11) ? (_1151_v29) : (_1151_v29)));
          let _1153_v31;
          _1153_v31 = _dafny.Seq.of(_1152_v30);
          let _1154_v32;
          _1154_v32 = _dafny.Seq.of(_1146_v24);
          _1152_v30 = (_1153_v31)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_1154_v32, _module.__default.fm18(_1144_cf10, globalState))).length), new BigNumber((_1153_v31).length))];
          let _1155_v33;
          let _nw220 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1155_v33 = _nw220;
          let _index229 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1155_v33).length));
          (_1155_v33)[_index229] = _1146_v24;
          let _index230 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1155_v33).length));
          let _rhs258 = _1146_v24;
          let _rhs259 = _dafny.Seq.Concat(_1146_v24, (_1147_v25).f24);
          let _rhs260 = _1143_cf11;
          let _lhs203 = _1155_v33;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1155_v33).length));
          let _lhs205 = globalState;
          _1146_v24 = _rhs258;
          _lhs203[_lhs204] = _rhs259;
          _lhs205.f5 = _rhs260;
          _1144_cf10 = _1144_cf10;
          (globalState).f2 = !(_this.f16);
        } else if (_source12.is_DC6) {
          let _1156_v34;
          _1156_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16)).length));
          let _1157_v35;
          _1157_v35 = _dafny.Seq.of(_this.f16);
          _1156_v34 = (_1156_v34).update(!(_this.f16), new BigNumber((_1157_v35).length));
          let _1158_v36;
          _1158_v36 = new BigNumber(39);
          let _1159_v37;
          _1159_v37 = _module.D8.create_DC21(_1158_v36, _this.f16, _this.f16, false, _1158_v36);
          let _1160_v38;
          _1160_v38 = _module.D4.create_DC12(_this.f16, _this.f16, (_1159_v37).dtor_cf39);
          let _1161_v39;
          let _init31 = function (_1162_i2) {
            return _this.f16;
          };
          let _nw221 = Array((new BigNumber(23)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw221.length); _i0_31++) {
            _nw221[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1161_v39 = _nw221;
          let _1163_v40;
          let _nw222 = new _module.C0();
          _nw222.__ctor((_1158_v36).multipliedBy((_1160_v38).dtor_cf23), _1161_v39);
          _1163_v40 = _nw222;
          let _1164_v41;
          _1164_v41 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1163_v40).f26),_this.f16);
          let _1165_v42;
          _1165_v42 = new _dafny.CodePoint('x'.codePointAt(0));
          (_this).f16 = !(((_this.f16) || ((((_1164_v41).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1165_v42)).length))) ? ((_1164_v41).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1165_v42)).length))) : (_this.f16)))) === (_this.f16));
          let _rhs261 = !((_1158_v36).isLessThanOrEqualTo((_1163_v40).f26));
          let _rhs262 = (_dafny.ZERO).minus(_1158_v36);
          let _lhs206 = globalState;
          _lhs206.f2 = _rhs261;
          _1158_v36 = _rhs262;
        } else {
          let _1166___mcc_h5 = (_source12).cf6;
          let _1167_cf6 = _1166___mcc_h5;
          let _1168_v43;
          let _nw223 = Array((new BigNumber(9)).toNumber()).fill(false);
          _1168_v43 = _nw223;
          let _init32 = function (_1169_i3) {
            return _this.f16;
          };
          let _nw224 = Array((new BigNumber(15)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw224.length); _i0_32++) {
            _nw224[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1168_v43 = _nw224;
          let _1170_v44;
          _1170_v44 = new BigNumber(-578);
          let _1171_v45;
          _1171_v45 = _dafny.Seq.UnicodeFromString("lbnnl");
          _module.__default.m0(_1170_v44, _1171_v45, _1170_v44, ((false) ? (_1170_v44) : (_1170_v44)), globalState);
          let _1172_v46;
          let _nw225 = new _module.C4();
          _nw225.__ctor(_1170_v44, _this.f16);
          _1172_v46 = _nw225;
          let _1173_v47;
          _1173_v47 = _dafny.MultiSet.fromElements(_1172_v46, _1172_v46, _1172_v46);
          let _1174_v49;
          _1174_v49 = _dafny.Seq.of(_module.__default.fm17(globalState));
          let _1175_v50;
          _1175_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1167_cf6).cardinality()),_1174_v49);
          let _1176_v51;
          _1176_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1173_v47,(_dafny.ZERO).minus(new BigNumber((function () {
            let _coll55 = new _dafny.Map();
            for (const _compr_55 of ((((_1175_v50).contains((_1172_v46).f22)) ? ((_1175_v50).get((_1172_v46).f22)) : (_1174_v49))).Elements) {
              let _1177_v48 = _compr_55;
              if (_dafny.Seq.contains((((_1175_v50).contains((_1172_v46).f22)) ? ((_1175_v50).get((_1172_v46).f22)) : (_1174_v49)), _1177_v48)) {
                _coll55.push([_1177_v48,_1172_v46.f23]);
              }
            }
            return _coll55;
          }()).length)));
          let _1178_v52;
          _1178_v52 = _dafny.Seq.of(_1172_v46);
          _1176_v51 = (_1176_v51).update(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1178_v52, _1178_v52)), _module.__default.fm3(globalState));
          let _1179_v53;
          _1179_v53 = _module.D1.create_DC5((_1172_v46).f22, _1172_v46.f23);
          (_this).f16 = ((_1179_v53).dtor_cf10).isLessThanOrEqualTo(((_1172_v46).f22).minus((_1172_v46).f22));
        }
        let _1180_v54;
        _1180_v54 = new BigNumber(649);
        _1180_v54 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.fm3(globalState)), (_dafny.ZERO).minus(_1180_v54));
        let _1181_v55;
        let _nw226 = new _module.C2();
        _nw226.__ctor(_1180_v54);
        _1181_v55 = _nw226;
        (globalState).f2 = !(!(((_this.f16) ? (_this.f16) : (_this.f16))));
      } else {
        let _1182_v56;
        _1182_v56 = new BigNumber(119);
        _1182_v56 = _module.__default.safeDivisionInt(_1182_v56, _1182_v56);
        let _1183_v57;
        _1183_v57 = _dafny.Set.fromElements(_module.__default.fm0(_this.f16, _this.f16, globalState));
        let _1184_v58;
        _1184_v58 = _dafny.Seq.of(true, (_1183_v57).IsDisjointFrom(_1183_v57), _this.f16, ((_dafny.ZERO).minus(_1182_v56)).isLessThan(_1182_v56));
        (_this).f16 = (_1184_v58)[_module.__default.safeIndex(_1182_v56, new BigNumber((_1184_v58).length))];
        let _1185_v59;
        _1185_v59 = _dafny.MultiSet.fromElements(_1182_v56);
        let _1186_v60;
        _1186_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1185_v59,_this.f16);
        _1186_v60 = (_1186_v60).update(((_dafny.MultiSet.fromElements(_1182_v56)).update(_1182_v56, _module.__default.abs(_1182_v56))).update(_1182_v56, _module.__default.abs(new BigNumber((_dafny.Seq.update((_this).f17, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_1182_v56)), new BigNumber(((_this).f17).length)), _1182_v56)).length))), !(!(_this.f16)));
        let _1187_v61;
        _1187_v61 = _dafny.Seq.UnicodeFromString("wbufmbxty");
        let _source13 = _module.__default.fm34(_1182_v56, _1182_v56, _1183_v57, new BigNumber((_1187_v61).length), globalState);
        if (_source13.is_DC4) {
          let _1188___mcc_h6 = (_source13).cf7;
          let _1189___mcc_h7 = (_source13).cf8;
          let _1190___mcc_h8 = (_source13).cf9;
          let _1191_cf9 = _1190___mcc_h8;
          let _1192_cf8 = _1189___mcc_h7;
          let _1193_cf7 = _1188___mcc_h6;
          let _1194_v62;
          _1194_v62 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1195_v63;
          _1195_v63 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(_1182_v56), _1182_v56),_1194_v62);
          _1195_v63 = (_1195_v63).update(_1182_v56, _1194_v62);
          let _1196_v64;
          let _nw227 = new _module.C2();
          _nw227.__ctor(_module.__default.fm3(globalState));
          _1196_v64 = _nw227;
          _1187_v61 = _dafny.Seq.Concat(_1187_v61, _1187_v61);
          let _1197_v65;
          let _nw228 = Array((new BigNumber(4)).toNumber()).fill([]);
          _1197_v65 = _nw228;
          let _1198_v66;
          let _init33 = ((_1199_v64) => function (_1200_i4) {
            return (_1200_i4).multipliedBy((_1199_v64).f25);
          })(_1196_v64);
          let _nw229 = Array((new BigNumber(16)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw229.length); _i0_33++) {
            _nw229[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1198_v66 = _nw229;
          let _index231 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_1197_v65).length));
          (_1197_v65)[_index231] = _1198_v66;
          let _1201_v67;
          _1201_v67 = _dafny.MultiSet.fromElements(_1193_cf7, _1193_cf7);
          let _index232 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_1197_v65).length));
          let _rhs263 = (_1182_v56).plus((_dafny.ZERO).minus(((_this.f16) ? ((_dafny.ZERO).minus(_1182_v56)) : ((_1196_v64).f25))));
          let _rhs264 = _1198_v66;
          let _rhs265 = (_this).fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-637)), ((_1202_v62) => function (_1203_i5) {
            return _1202_v62;
          })(_1194_v62)), globalState);
          let _rhs266 = _dafny.Seq.of((_this.f16) && (_this.f16), _1193_cf7, ((_module.D8.create_DC21(new BigNumber(62), !(_this.f16), _this.f16, _1193_cf7, new BigNumber((_1201_v67).cardinality()))).dtor_cf35).isEqualTo(_1192_cf8), true);
          let _rhs267 = (_1201_v67).IsSubsetOf((_1201_v67).Union(_dafny.MultiSet.FromArray(_1184_v58)));
          let _lhs207 = _1197_v65;
          let _lhs208 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_1197_v65).length));
          let _lhs209 = _this;
          let _lhs210 = globalState;
          _1192_cf8 = _rhs263;
          _lhs207[_lhs208] = _rhs264;
          _lhs209.f16 = _rhs265;
          _1184_v58 = _rhs266;
          _lhs210.f2 = _rhs267;
        } else if (_source13.is_DC5) {
          let _1204___mcc_h9 = (_source13).cf10;
          let _1205___mcc_h10 = (_source13).cf11;
          let _1206_cf11 = _1205___mcc_h10;
          let _1207_cf10 = _1204___mcc_h9;
          let _1208_v68;
          _1208_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1207_cf10,(_this).f17);
          let _1209_v69;
          _1209_v69 = new _dafny.CodePoint('d'.codePointAt(0));
          let _1210_v70;
          _1210_v70 = _dafny.Map.Empty.slice().updateUnsafe(_1209_v69,new BigNumber(483));
          _1208_v68 = (_1208_v68).update(new BigNumber((function () {
            let _coll56 = new _dafny.Set();
            for (const _compr_56 of (_1210_v70).Keys.Elements) {
              let _1211_v71 = _compr_56;
              if ((_1210_v70).contains(_1211_v71)) {
                _coll56.add(_1211_v71);
              }
            }
            return _coll56;
          }()).length), (_this).f17);
          _1206_cf11 = _this.f16;
          let _1212_v72;
          let _init34 = ((_1213_cf10) => function (_1214_i6) {
            return _module.__default.safeModuloInt(_1214_i6, _1213_cf10);
          })(_1207_cf10);
          let _nw230 = Array((new BigNumber(11)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw230.length); _i0_34++) {
            _nw230[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1212_v72 = _nw230;
          let _index233 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1212_v72).length));
          (_1212_v72)[_index233] = new BigNumber((_dafny.Seq.UnicodeFromString("hfgkuwyr")).length);
          let _index234 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1212_v72).length));
          let _rhs268 = ((_1182_v56).minus(_1182_v56)).minus(new BigNumber(-640));
          let _rhs269 = ((_this).f17)[_module.__default.safeIndex(_1182_v56, new BigNumber(((_this).f17).length))];
          let _rhs270 = _1187_v61;
          let _rhs271 = ((_this.f16) ? (_1182_v56) : (_1182_v56));
          let _lhs211 = _1212_v72;
          let _lhs212 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_1212_v72).length));
          _lhs211[_lhs212] = _rhs268;
          _1207_cf10 = _rhs269;
          _1187_v61 = _rhs270;
          _1207_cf10 = _rhs271;
          _1182_v56 = ((_this.f16) ? (new BigNumber((_dafny.Seq.Concat(_1187_v61, _1187_v61)).length)) : ((new BigNumber((_1187_v61).length)).multipliedBy(new BigNumber(-682))));
        } else if (_source13.is_DC6) {
          let _1215_v73;
          _1215_v73 = new _dafny.CodePoint('w'.codePointAt(0));
          (globalState).f15 = _1215_v73;
          _1182_v56 = _1182_v56;
          let _1216_v74;
          _1216_v74 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,true);
          let _1217_v75;
          _1217_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1182_v56,false);
          _1216_v74 = (_1216_v74).update((((_1217_v75).contains(_1182_v56)) ? ((_1217_v75).get(_1182_v56)) : (_this.f16)), _module.__default.fm0(_this.f16, _this.f16, globalState));
          (globalState).f2 = _module.__default.fm0(!(_this.f16), _this.f16, globalState);
        } else {
          let _1218___mcc_h11 = (_source13).cf6;
          let _1219_cf6 = _1218___mcc_h11;
          let _rhs272 = _1182_v56;
          let _rhs273 = _1187_v61;
          _1182_v56 = _rhs272;
          _1187_v61 = _rhs273;
          let _1220_v76;
          _1220_v76 = _dafny.Set.fromElements(_1182_v56);
          let _1221_v77;
          _1221_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1219_cf6).cardinality()),new BigNumber((_1220_v76).length));
          _1182_v56 = (_1182_v56).plus((new BigNumber((_1221_v77).length)).plus(_1182_v56));
          _1182_v56 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1182_v56, _1182_v56));
          _1182_v56 = new BigNumber((((_1219_cf6).Intersect(_1219_cf6)).Difference(_1219_cf6)).cardinality());
        }
        let _1222_v78;
        _1222_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this.f16) || (_this.f16),_this.f16);
        _1222_v78 = (_1222_v78).update(_this.f16, _this.f16);
      }
      let _1223_v79;
      let _init35 = function (_1224_i8) {
        return _this.f16;
      };
      let _nw231 = Array((new BigNumber(14)).toNumber());
      for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw231.length); _i0_35++) {
        _nw231[_i0_35] = _init35(new BigNumber(_i0_35));
      }
      _1223_v79 = _nw231;
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1223_v79).length))) {
        let _1225_i7 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1225_i7)) && ((_1225_i7).isLessThan(new BigNumber((_1223_v79).length))))) {
          (_1223_v79)[(_1225_i7)] = ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),new BigNumber(-381))).length)).minus(new BigNumber(836))).isEqualTo(((true) ? (new BigNumber(319)) : (new BigNumber(-786))));
        }
      }
      _1223_v79 = _1223_v79;
      let _1226_v80;
      _1226_v80 = new _dafny.CodePoint('n'.codePointAt(0));
      let _1227_v81;
      _1227_v81 = _dafny.Seq.of(((_this.f16) ? (_1226_v80) : (_1226_v80)), _1226_v80, _1226_v80);
      let _1228_v82;
      _1228_v82 = new BigNumber(516);
      (globalState).f15 = (_1227_v81)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_1228_v82, _1228_v82), new BigNumber((_1227_v81).length))];
      let _1229_v83;
      _1229_v83 = _dafny.Seq.of(_this.f16);
      let _1230_v84;
      _1230_v84 = _module.D0.create_DC1(_1227_v81);
      let _1231_v85;
      _1231_v85 = _dafny.Set.fromElements((_1229_v83)[_module.__default.safeIndex(new BigNumber(((_1230_v84).dtor_cf1).length), new BigNumber((_1229_v83).length))], !(!(true)), !(_this.f16));
      if ((_1231_v85).IsProperSubsetOf(_1231_v85)) {
        let _1232_v86;
        let _init36 = ((_1233_v82) => function (_1234_i9) {
          return (_1234_i9).plus(_1233_v82);
        })(_1228_v82);
        let _nw232 = Array((new BigNumber(18)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw232.length); _i0_36++) {
          _nw232[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1232_v86 = _nw232;
        let _index235 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length));
        (_1232_v86)[_index235] = _1228_v82;
        let _1235_v87;
        _1235_v87 = _dafny.MultiSet.fromElements(_1226_v80, (_1227_v81)[_module.__default.safeIndex(_1228_v82, new BigNumber((_1227_v81).length))], new _dafny.CodePoint('j'.codePointAt(0)));
        let _index236 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length));
        (_1232_v86)[_index236] = new BigNumber((_1235_v87).cardinality());
        let _1236_v88;
        _1236_v88 = _dafny.Set.fromElements(_1228_v82, (_1232_v86)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length))]);
        let _1237_v90;
        _1237_v90 = _module.D9.create_DC27(_this.f16, _1236_v88);
        let _1238_v91;
        _1238_v91 = _dafny.Seq.of((_1237_v90).dtor_cf48, _dafny.Set.fromElements(_1228_v82));
        let _1239_v92;
        _1239_v92 = _dafny.MultiSet.fromElements(_1236_v88, function () {
          let _coll57 = new _dafny.Set();
          for (const _compr_57 of ((_this).f17).Elements) {
            let _1240_v89 = _compr_57;
            if (_dafny.Seq.contains((_this).f17, _1240_v89)) {
              _coll57.add(_module.__default.safeModuloInt(_1240_v89, new BigNumber(564)));
            }
          }
          return _coll57;
        }(), (_1236_v88).Difference(_1236_v88), (_1238_v91)[_module.__default.safeIndex(_1228_v82, new BigNumber((_1238_v91).length))], _1236_v88);
        _1239_v92 = _1239_v92;
        let _1241_v93;
        let _nw233 = Array((new BigNumber(10)).toNumber());
        _nw233[(_dafny.ZERO).toNumber()] = _1226_v80;
        _nw233[(_dafny.ONE).toNumber()] = _1226_v80;
        _nw233[(new BigNumber(2)).toNumber()] = _1226_v80;
        _nw233[(new BigNumber(3)).toNumber()] = _1226_v80;
        _nw233[(new BigNumber(4)).toNumber()] = _1226_v80;
        _nw233[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw233[(new BigNumber(6)).toNumber()] = _1226_v80;
        _nw233[(new BigNumber(7)).toNumber()] = _1226_v80;
        _nw233[(new BigNumber(8)).toNumber()] = _1226_v80;
        _nw233[(new BigNumber(9)).toNumber()] = _1226_v80;
        _1241_v93 = _nw233;
        _1241_v93 = _1241_v93;
        let _index237 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length));
        let _rhs274 = _this.f16;
        let _rhs275 = (_1232_v86)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length))];
        let _rhs276 = (_module.__default.fm3(globalState)).isLessThanOrEqualTo((_1232_v86)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length))]);
        let _lhs213 = globalState;
        let _lhs214 = _1232_v86;
        let _lhs215 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length));
        let _lhs216 = globalState;
        _lhs213.f5 = _rhs274;
        _lhs214[_lhs215] = _rhs275;
        _lhs216.f2 = _rhs276;
        let _index238 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length));
        (_1232_v86)[_index238] = (_1232_v86)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1232_v86).length))];
      } else {
        let _1242_v94;
        _1242_v94 = _module.D0.create_DC0(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dlth"), _dafny.Seq.UnicodeFromString("ilunfi")));
        _1242_v94 = _1242_v94;
        let _1243_v95;
        let _nw234 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _1243_v95 = _nw234;
        let _index239 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length));
        (_1223_v79)[_index239] = (_this).fm8(_1227_v81, globalState);
        let _index240 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length));
        let _rhs277 = _1226_v80;
        let _rhs278 = _1243_v95;
        let _rhs279 = _this.f16;
        let _lhs217 = globalState;
        let _lhs218 = _1223_v79;
        let _lhs219 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length));
        _lhs217.f15 = _rhs277;
        _1243_v95 = _rhs278;
        _lhs218[_lhs219] = _rhs279;
        if ((_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))]) {
          let _1244_v96;
          let _nw235 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1244_v96 = _nw235;
          let _1245_v97;
          _1245_v97 = _module.D0.create_DC2(_1228_v82, _1228_v82, new BigNumber((_dafny.Set.fromElements(_1228_v82)).length), _1244_v96);
          let _1246_v98;
          let _nw236 = new _module.C0();
          _nw236.__ctor(_1228_v82, (_1245_v97).dtor_cf5);
          _1246_v98 = _nw236;
          let _1247_v99;
          _1247_v99 = _dafny.Seq.of(_1244_v96, _1244_v96);
          let _1248_v100;
          _1248_v100 = _dafny.Map.Empty.slice().updateUnsafe((_1229_v83)[_module.__default.safeIndex((_dafny.ZERO).minus((_1246_v98).f26), new BigNumber((_1229_v83).length))],(_1246_v98).f26);
          let _1249_v101;
          _1249_v101 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1244_v96);
          let _1250_v102;
          let _nw237 = Array((new BigNumber(18)).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = (((_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))]) ? ((_1247_v99)[_module.__default.safeIndex(new BigNumber((_1248_v100).length), new BigNumber((_1247_v99).length))]) : (_1244_v96));
          _nw237[(_dafny.ONE).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(2)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(3)).toNumber()] = (((_1249_v101).contains(false)) ? ((_1249_v101).get(false)) : (_1246_v98.f27));
          _nw237[(new BigNumber(4)).toNumber()] = _1244_v96;
          _nw237[(new BigNumber(5)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(6)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(7)).toNumber()] = _1244_v96;
          _nw237[(new BigNumber(8)).toNumber()] = _1244_v96;
          _nw237[(new BigNumber(9)).toNumber()] = _1244_v96;
          _nw237[(new BigNumber(10)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(11)).toNumber()] = _1244_v96;
          _nw237[(new BigNumber(12)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(13)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(14)).toNumber()] = _1244_v96;
          _nw237[(new BigNumber(15)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(16)).toNumber()] = _1246_v98.f27;
          _nw237[(new BigNumber(17)).toNumber()] = _1246_v98.f27;
          _1250_v102 = _nw237;
          let _index241 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1250_v102).length));
          (_1250_v102)[_index241] = _1244_v96;
          let _index242 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1250_v102).length));
          let _rhs280 = (new BigNumber((_dafny.Seq.of(_this.f16)).length)).plus((_dafny.ZERO).minus(_1228_v82));
          let _rhs281 = _1226_v80;
          let _rhs282 = _1246_v98.f27;
          let _lhs220 = globalState;
          let _lhs221 = _1250_v102;
          let _lhs222 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_1250_v102).length));
          _1228_v82 = _rhs280;
          _lhs220.f15 = _rhs281;
          _lhs221[_lhs222] = _rhs282;
          let _1251_v104;
          _1251_v104 = _dafny.Map.Empty.slice().updateUnsafe(_1228_v82,(_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))]);
          let _1252_v106;
          _1252_v106 = _dafny.Map.Empty.slice().updateUnsafe((function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of (_1251_v104).Keys.Elements) {
              let _1253_v103 = _compr_58;
              if ((_1251_v104).contains(_1253_v103)) {
                _coll58.push([(_1253_v103).multipliedBy(new BigNumber(586)),(_module.D6.create_DC16(_this.f16, (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))])).dtor_cf29]);
              }
            }
            return _coll58;
          }()).Merge(function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(709), new BigNumber(931))) {
              let _1254_v105 = _compr_59;
              if (((new BigNumber(709)).isLessThanOrEqualTo(_1254_v105)) && ((_1254_v105).isLessThan(new BigNumber(931)))) {
                _coll59.push([_module.__default.safeDivisionInt(_1254_v105, (_1246_v98).f26),(_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))]]);
              }
            }
            return _coll59;
          }()),_module.__default.safeDivisionInt(new BigNumber(748), (_1246_v98).f26));
          _1252_v106 = (_1252_v106).update((_1251_v104).update((_1246_v98).f26, false), (_1246_v98).f26);
          let _1255_v107;
          _1255_v107 = _dafny.Set.fromElements(_1231_v85, _1231_v85, _1231_v85, _1231_v85, _1231_v85);
          let _1256_v108;
          let _init37 = ((_1257_v80) => function (_1258_i10) {
            return _1257_v80;
          })(_1226_v80);
          let _nw238 = Array((new BigNumber(17)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw238.length); _i0_37++) {
            _nw238[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1256_v108 = _nw238;
          let _1259_v109;
          _1259_v109 = _1256_v108;
          let _index243 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length));
          let _rhs283 = _module.__default.fm0(_this.f16, _this.f16, globalState);
          let _rhs284 = !(_1255_v107).equals(_1255_v107);
          let _rhs285 = (((!((_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))])) ? (new BigNumber((_dafny.Seq.of(_1256_v108, _1256_v108, _1256_v108, (_1259_v109), _1256_v108)).length)) : (_1228_v82))).plus(_1228_v82);
          let _rhs286 = _this.f16;
          let _rhs287 = !((_1246_v98).f26).isEqualTo(_1228_v82);
          let _lhs223 = _1223_v79;
          let _lhs224 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length));
          let _lhs225 = globalState;
          let _lhs226 = _this;
          let _lhs227 = globalState;
          _lhs223[_lhs224] = _rhs283;
          _lhs225.f5 = _rhs284;
          _1228_v82 = _rhs285;
          _lhs226.f16 = _rhs286;
          _lhs227.f2 = _rhs287;
          let _index244 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_1243_v95).length));
          (_1243_v95)[_index244] = new BigNumber(724);
          let _1260_v110;
          let _nw239 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1260_v110 = _nw239;
          let _index245 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_1260_v110).length));
          (_1260_v110)[_index245] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), ((_1261_v80) => function (_1262_i11) {
            return _1261_v80;
          })(_1226_v80)), _1227_v81);
          let _1263_v111;
          _1263_v111 = _dafny.Set.fromElements(_module.__default.fm3(globalState), (_1246_v98).f26);
          let _1264_v112;
          _1264_v112 = _module.D8.create_DC21(_1228_v82, _this.f16, (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))], (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))], (_1246_v98).f26);
          let _index246 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1243_v95).length));
          (_1243_v95)[_index246] = (((_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))]) ? (new BigNumber((_1263_v111).length)) : ((_1264_v112).dtor_cf35));
          let _index247 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_1243_v95).length));
          let _index248 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_1260_v110).length));
          let _index249 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1243_v95).length));
          let _rhs288 = (_1246_v98).f26;
          let _rhs289 = _1227_v81;
          let _rhs290 = _1248_v100;
          let _rhs291 = (_1228_v82).multipliedBy((_1246_v98).f26);
          let _lhs228 = _1243_v95;
          let _lhs229 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_1243_v95).length));
          let _lhs230 = _1260_v110;
          let _lhs231 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_1260_v110).length));
          let _lhs232 = _1243_v95;
          let _lhs233 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1243_v95).length));
          _lhs228[_lhs229] = _rhs288;
          _lhs230[_lhs231] = _rhs289;
          _1248_v100 = _rhs290;
          _lhs232[_lhs233] = _rhs291;
        } else {
          let _1265_v113;
          let _nw240 = Array((new BigNumber(20)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = _1227_v81;
          _nw240[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_1266_v80) => function (_1267_i12) {
            return _1266_v80;
          })(_1226_v80));
          _nw240[(new BigNumber(2)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("tgjcs");
          _nw240[(new BigNumber(4)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(5)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1227_v81, _1227_v81);
          _nw240[(new BigNumber(7)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(8)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("fyruluxc");
          _nw240[(new BigNumber(10)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(317)), ((_1268_v80) => function (_1269_i13) {
            return _1268_v80;
          })(_1226_v80));
          _nw240[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1227_v81, _1227_v81);
          _nw240[(new BigNumber(13)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), ((_1270_v80) => function (_1271_i14) {
            return _1270_v80;
          })(_1226_v80));
          _nw240[(new BigNumber(15)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(16)).toNumber()] = _dafny.Seq.UnicodeFromString("hfvxlqw");
          _nw240[(new BigNumber(17)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(18)).toNumber()] = _1227_v81;
          _nw240[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("evndkdyhw");
          _1265_v113 = _nw240;
          let _index250 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1265_v113).length));
          (_1265_v113)[_index250] = _dafny.Seq.Concat(_1227_v81, _1227_v81);
          let _index251 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1265_v113).length));
          let _rhs292 = _module.__default.fm3(globalState);
          let _rhs293 = _1229_v83;
          let _rhs294 = _dafny.Seq.Concat(_1227_v81, _dafny.Seq.Create(_module.__default.abs(new BigNumber(205)), ((_1272_v80) => function (_1273_i15) {
            return _1272_v80;
          })(_1226_v80)));
          let _lhs234 = _1265_v113;
          let _lhs235 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_1265_v113).length));
          _1228_v82 = _rhs292;
          r0 = _rhs293;
          _lhs234[_lhs235] = _rhs294;
          let _1274_v114;
          _1274_v114 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(944)).plus(_1228_v82),_this.f16);
          let _1275_v115;
          let _nw241 = Array((_dafny.ONE).toNumber()).fill(false);
          _1275_v115 = _nw241;
          let _1276_v116;
          let _nw242 = new _module.C0();
          _nw242.__ctor(new BigNumber(((_1265_v113)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_1265_v113).length))]).length), _1275_v115);
          _1276_v116 = _nw242;
          let _1277_v117;
          _1277_v117 = _module.D9.create_DC26(_1228_v82, _1276_v116);
          _1274_v114 = (_1274_v114).update((_1277_v117).dtor_cf45, (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))]);
          let _1278_v118;
          let _nw243 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
          _1278_v118 = _nw243;
          let _1279_v119;
          _1279_v119 = _dafny.Map.Empty.slice().updateUnsafe(((_1276_v116).f26).plus((_1276_v116).f26),_1278_v118);
          _1279_v119 = (_1279_v119).update(_1228_v82, _1278_v118);
          _1228_v82 = (_1228_v82).plus((_1276_v116).f26);
          let _1280_v120;
          _1280_v120 = _module.D8.create_DC21((_1276_v116).f26, (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))], _this.f16, (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))], (_1276_v116).f26);
          let _1281_v121;
          _1281_v121 = _dafny.MultiSet.fromElements((_1280_v120).dtor_cf35, new BigNumber(141));
          let _1282_v122;
          _1282_v122 = _dafny.MultiSet.fromElements(_1226_v80, _1226_v80, _1226_v80);
          _1228_v82 = (new BigNumber((_1281_v121).cardinality())).multipliedBy(new BigNumber((_1282_v122).cardinality()));
        }
        if ((_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))]) {
          let _1283_v123;
          let _nw244 = new _module.C3();
          _nw244.__ctor(_1227_v81, _this.f16, _dafny.Seq.Concat(_dafny.Seq.of(_1228_v82), _dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), ((_1284_v83) => function (_1285_i16) {
            return new BigNumber((_1284_v83).length);
          })(_1229_v83))));
          _1283_v123 = _nw244;
          let _1286_v124;
          _1286_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1228_v82,_1228_v82);
          let _1287_v125;
          _1287_v125 = _dafny.Set.fromElements(_1286_v124);
          let _1288_v126;
          let _nw245 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1288_v126 = _nw245;
          let _1289_v127;
          _1289_v127 = _dafny.Map.Empty.slice().updateUnsafe((_1287_v125).IsProperSubsetOf(_module.__default.fm35(_1228_v82, new _dafny.CodePoint('o'.codePointAt(0)), globalState)),_1288_v126);
          _1289_v127 = (_1289_v127).update((_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))], _1288_v126);
          (_this).f16 = (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))];
          _1228_v82 = _1228_v82;
          _1228_v82 = _1228_v82;
        } else {
          let _1290_v128;
          _1290_v128 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ikbah"),_module.__default.fm0(_this.f16, (_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))], globalState));
          let _1291_v129;
          _1291_v129 = _dafny.Map.Empty.slice().updateUnsafe((_this.f16) || (_this.f16),(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("fhohi"),_this.f16)).Merge(_1290_v128));
          _1291_v129 = (_1291_v129).update((_1223_v79)[_module.__default.safeIndex(new BigNumber(954), new BigNumber((_1223_v79).length))], _1290_v128);
          let _1292_v130;
          _1292_v130 = _dafny.MultiSet.fromElements(_1227_v81, _1227_v81, _1227_v81);
          let _1293_v131;
          _1293_v131 = _dafny.Seq.of(_1227_v81, _1227_v81);
          let _rhs295 = _this.f16;
          let _rhs296 = _1228_v82;
          let _rhs297 = (_dafny.MultiSet.FromArray(_1293_v131)).update(_dafny.Seq.UnicodeFromString("ljveyjw"), _module.__default.abs(_module.__default.safeModuloInt(_1228_v82, _1228_v82)));
          let _rhs298 = _1228_v82;
          let _lhs236 = globalState;
          _lhs236.f5 = _rhs295;
          _1228_v82 = _rhs296;
          _1292_v130 = _rhs297;
          _1228_v82 = _rhs298;
          (_this).f16 = false;
          _1228_v82 = (_1228_v82).multipliedBy((_dafny.ZERO).minus(_1228_v82));
          let _1294_v132;
          _1294_v132 = _dafny.MultiSet.fromElements(_1228_v82);
          (globalState).f2 = (_1294_v132).contains(_1228_v82);
        }
        let _rhs299 = _1228_v82;
        let _rhs300 = !((_dafny.ZERO).minus(_1228_v82)).isEqualTo(_1228_v82);
        let _lhs237 = _this;
        _1228_v82 = _rhs299;
        _lhs237.f16 = _rhs300;
      }
      let _1295_v133;
      _1295_v133 = _module.D8.create_DC24(_module.D8.create_DC23());
      _1295_v133 = _1295_v133;
      r0 = _dafny.Seq.of(_this.f16);
      return r0;
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f21 = _module.D0.Default();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f21) {
      let _this = this;
      (_this)._f21 = f21;
      return;
    }
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.of(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("lc")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), function (_1296_i0) {
        return new BigNumber((_dafny.Set.fromElements(!(!(true)), false, !(false), false)).length);
      })).length)));
    };
    m3(p0, globalState) {
      let _this = this;
      let _1297_v0;
      _1297_v0 = _dafny.Seq.UnicodeFromString("qpaxlupea");
      let _1298_v1;
      _1298_v1 = new BigNumber(76);
      let _hi8 = (_1298_v1).multipliedBy(_1298_v1);
      for (let _1299_i0 = new BigNumber((_1297_v0).length); _1299_i0.isLessThan(_hi8); _1299_i0 = _1299_i0.plus(_dafny.ONE)) {
        let _1300_v2;
        let _nw246 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _1300_v2 = _nw246;
        let _index252 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_1300_v2).length));
        (_1300_v2)[_index252] = (_1299_i0).plus(_1299_i0);
        let _index253 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_1300_v2).length));
        (_1300_v2)[_index253] = _1298_v1;
        let _1301_v3;
        _1301_v3 = _dafny.Set.fromElements(p0);
        let _1302_v4;
        _1302_v4 = _module.D2.create_DC7(_1301_v3);
        let _1303_v5;
        _1303_v5 = _dafny.MultiSet.fromElements(new BigNumber((((_1302_v4).dtor_cf12).Union(_1301_v3)).length), new BigNumber(441));
        let _1304_v6;
        _1304_v6 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1305_v7;
        _1305_v7 = _dafny.Set.fromElements(_1304_v6, _1304_v6, _module.__default.fm2(globalState));
        _1298_v1 = (((_1303_v5).contains(new BigNumber((_1305_v7).length))) ? ((_1303_v5).get(new BigNumber((_1305_v7).length))) : ((_1298_v1).multipliedBy(_1299_i0)));
        let _index254 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_1300_v2).length));
        (_1300_v2)[_index254] = (_dafny.ZERO).minus(_1299_i0);
        let _1306_v8;
        _1306_v8 = _module.D1.create_DC6();
        _1306_v8 = _1306_v8;
      }
      let _hi9 = new BigNumber(169);
      for (let _1307_i1 = _1298_v1; _1307_i1.isLessThan(_hi9); _1307_i1 = _1307_i1.plus(_dafny.ONE)) {
        let _1308_v9;
        _1308_v9 = _dafny.MultiSet.fromElements(p0);
        let _1309_v10;
        _1309_v10 = _module.D1.create_DC3(_1308_v9);
        let _pat_let_tv27 = _1308_v9;
        let _source14 = function (_pat_let35_0) {
          return function (_1310_dt__update__tmp_h0) {
            return function (_pat_let36_0) {
              return function (_1311_dt__update_hcf6_h0) {
                return _module.D1.create_DC3(_1311_dt__update_hcf6_h0);
              }(_pat_let36_0);
            }(_pat_let_tv27);
          }(_pat_let35_0);
        }(_1309_v10);
        if (_source14.is_DC4) {
          let _1312___mcc_h0 = (_source14).cf7;
          let _1313___mcc_h1 = (_source14).cf8;
          let _1314___mcc_h2 = (_source14).cf9;
          let _1315_cf9 = _1314___mcc_h2;
          let _1316_cf8 = _1313___mcc_h1;
          let _1317_cf7 = _1312___mcc_h0;
          let _1318_v11;
          _1318_v11 = _dafny.Set.fromElements(_1298_v1);
          let _1319_v12;
          _1319_v12 = _dafny.Set.fromElements(_1318_v11, _1318_v11);
          let _1320_v13;
          _1320_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(175),_1307_i1);
          (globalState).f2 = ((_dafny.Set.fromElements(_1318_v11)).Intersect(_dafny.Set.fromElements(function () {
            let _coll60 = new _dafny.Set();
            for (const _compr_60 of (_1320_v13).Keys.Elements) {
              let _1321_v14 = _compr_60;
              if ((_1320_v13).contains(_1321_v14)) {
                _coll60.add(_module.__default.safeDivisionInt(_1321_v14, new BigNumber((_dafny.Seq.UnicodeFromString("mgmhqd")).length)));
              }
            }
            return _coll60;
          }()))).IsProperSubsetOf(_1319_v12);
          let _1322_v15;
          let _1323_v16;
          let _1324_v17;
          let _1325_v18;
          let _out10;
          let _out11;
          let _out12;
          let _out13;
          let _outcollector3 = (_this).m4(globalState);
          _out10 = _outcollector3[0];
          _out11 = _outcollector3[1];
          _out12 = _outcollector3[2];
          _out13 = _outcollector3[3];
          _1322_v15 = _out10;
          _1323_v16 = _out11;
          _1324_v17 = _out12;
          _1325_v18 = _out13;
          let _1326_v19;
          _1326_v19 = _dafny.Seq.of(_1297_v0, _dafny.Seq.UnicodeFromString("sluvhgu"));
          let _1327_v20;
          _1327_v20 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1328_v21;
          _1328_v21 = _dafny.MultiSet.fromElements(_dafny.Seq.update((_1326_v19)[_module.__default.safeIndex(_1316_cf8, new BigNumber((_1326_v19).length))], _module.__default.safeIndex(new BigNumber((_1297_v0).length), new BigNumber(((_1326_v19)[_module.__default.safeIndex(_1316_cf8, new BigNumber((_1326_v19).length))]).length)), _1327_v20), _1297_v0, _1297_v0);
          _1328_v21 = _1328_v21;
          let _1329_v22;
          _1329_v22 = _dafny.Seq.of(new BigNumber(-967));
          let _1330_v23;
          let _nw247 = new _module.C1();
          _nw247.__ctor(new BigNumber(919), new BigNumber((_1297_v0).length), _1317_cf7, _1329_v22, _1316_cf8);
          _1330_v23 = _nw247;
          let _1331_v24;
          _1331_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1323_v16,_1330_v23);
          let _1332_v25;
          _1332_v25 = _module.D8.create_DC22(_1322_v15, _1330_v23, _1324_v17);
          let _1333_v26;
          _1333_v26 = _dafny.Seq.of(_1330_v23, (_1332_v25).dtor_cf41);
          let _1334_v27;
          _1334_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1322_v15,new BigNumber((_1329_v22).length));
          let _1335_v28;
          _1335_v28 = _dafny.Map.Empty.slice().updateUnsafe((((_1331_v24).contains(_1317_cf7)) ? ((_1331_v24).get(_1317_cf7)) : ((_1333_v26)[_module.__default.safeIndex(new BigNumber((_1334_v27).length), new BigNumber((_1333_v26).length))])),_1330_v23.f16);
          _1335_v28 = (_1335_v28).update(_1330_v23, _1323_v16);
        } else if (_source14.is_DC5) {
          let _1336___mcc_h3 = (_source14).cf10;
          let _1337___mcc_h4 = (_source14).cf11;
          let _1338_cf11 = _1337___mcc_h4;
          let _1339_cf10 = _1336___mcc_h3;
          let _1340_v29;
          let _nw248 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1340_v29 = _nw248;
          let _index255 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1340_v29).length));
          (_1340_v29)[_index255] = _1307_i1;
          let _index256 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1340_v29).length));
          (_1340_v29)[_index256] = _module.__default.safeModuloInt(_1339_cf10, _1339_cf10);
          (globalState).f2 = ((p0) ? (p0) : (!(p0)));
          let _1341_v30;
          _1341_v30 = _dafny.Seq.of(_1338_cf11);
          let _1342_v31;
          _1342_v31 = _dafny.Set.fromElements(p0, false);
          _1341_v30 = _dafny.Seq.Concat(_1341_v30, ((_1338_cf11) ? (_1341_v30) : (_dafny.Seq.update(_1341_v30, _module.__default.safeIndex(new BigNumber((_1342_v31).length), new BigNumber((_1341_v30).length)), _1338_cf11))));
          let _1343_v32;
          _1343_v32 = _dafny.Set.fromElements(_1298_v1);
          (globalState).f2 = _module.__default.fm0((_1343_v32).IsProperSubsetOf(_1343_v32), _module.__default.fm0(_module.__default.fm0(p0, p0, globalState), p0, globalState), globalState);
        } else if (_source14.is_DC6) {
          let _1344_v33;
          _1344_v33 = new _dafny.CodePoint('a'.codePointAt(0));
          let _pat_let_tv28 = _1344_v33;
          (globalState).f5 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_1297_v0, _module.__default.safeIndex(_1298_v1, new BigNumber((_1297_v0).length)), (function (_pat_let37_0) {
            return function (_1345_dt__update__tmp_h1) {
              return function (_pat_let38_0) {
                return function (_1346_dt__update_hcf50_h0) {
                  return _module.D11.create_DC29(_1346_dt__update_hcf50_h0);
                }(_pat_let38_0);
              }(_pat_let_tv28);
            }(_pat_let37_0);
          }(_module.D11.create_DC29(_1344_v33))).dtor_cf50), _1297_v0);
          _1298_v1 = _1307_i1;
          let _1347_v34;
          _1347_v34 = _dafny.Set.fromElements(_1344_v33, new _dafny.CodePoint('u'.codePointAt(0)));
          _1347_v34 = _1347_v34;
          let _1348_v35;
          let _nw249 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1348_v35 = _nw249;
          let _1349_v36;
          _1349_v36 = _dafny.Seq.of(p0);
          let _1350_v37;
          let _nw250 = Array((new BigNumber(2)).toNumber());
          _nw250[(_dafny.ZERO).toNumber()] = new BigNumber((_1349_v36).length);
          _nw250[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1307_i1);
          _1350_v37 = _nw250;
          let _1351_v38;
          _1351_v38 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_1348_v35) : (_1348_v35)),_1350_v37);
          _1351_v38 = (_1351_v38).update(_1348_v35, _1350_v37);
        } else {
          let _1352___mcc_h5 = (_source14).cf6;
          let _1353_cf6 = _1352___mcc_h5;
          let _1354_v39;
          _1354_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1307_i1);
          let _1355_v40;
          _1355_v40 = new _dafny.CodePoint('t'.codePointAt(0));
          let _1356_v41;
          _1356_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1298_v1,p0);
          let _1357_v42;
          _1357_v42 = _dafny.Seq.of(new BigNumber((_1356_v41).length));
          let _1358_v43;
          _1358_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1355_v40,_1357_v42);
          let _1359_v44;
          _1359_v44 = _dafny.Seq.of(_1357_v42);
          let _1360_v45;
          let _nw251 = Array((new BigNumber(18)).toNumber());
          _nw251[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_1298_v1);
          _nw251[(_dafny.ONE).toNumber()] = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(888)), function (_1361_i2) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          })).length));
          _nw251[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_1298_v1);
          _nw251[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(925)), ((_1362_v1) => function (_1363_i3) {
            return _1362_v1;
          })(_1298_v1));
          _nw251[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_1307_i1);
          _nw251[(new BigNumber(5)).toNumber()] = (((_1358_v43).contains(_1355_v40)) ? ((_1358_v43).get(_1355_v40)) : (_1357_v42));
          _nw251[(new BigNumber(6)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(new BigNumber(668));
          _nw251[(new BigNumber(8)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(9)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(10)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(11)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(12)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(13)).toNumber()] = (_1359_v44)[_module.__default.safeIndex(_1298_v1, new BigNumber((_1359_v44).length))];
          _nw251[(new BigNumber(14)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(15)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(16)).toNumber()] = _1357_v42;
          _nw251[(new BigNumber(17)).toNumber()] = (_module.D3.create_DC9(_1357_v42)).dtor_cf18;
          _1360_v45 = _nw251;
          let _1364_v46;
          _1364_v46 = _module.D2.create_DC8(_1354_v39, _1360_v45, p0, false, _1307_i1);
          let _1365_v47;
          _1365_v47 = _dafny.Map.Empty.slice().updateUnsafe((_1364_v46).dtor_cf15,_1297_v0);
          _1298_v1 = (_1307_i1).minus(new BigNumber((_1365_v47).length));
          let _1366_v48;
          let _init38 = ((_1367_p0) => function (_1368_i4) {
            return _dafny.Seq.of(_1367_p0, _1367_p0);
          })(p0);
          let _nw252 = Array((new BigNumber(21)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw252.length); _i0_38++) {
            _nw252[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1366_v48 = _nw252;
          let _1369_v49;
          _1369_v49 = _dafny.Seq.of(false, p0);
          let _index257 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1366_v48).length));
          (_1366_v48)[_index257] = _1369_v49;
          let _index258 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1366_v48).length));
          (_1366_v48)[_index258] = _1369_v49;
          let _1370_v50;
          let _nw253 = Array((new BigNumber(2)).toNumber()).fill([]);
          _1370_v50 = _nw253;
          let _1371_v51;
          let _nw254 = Array((new BigNumber(29)).toNumber()).fill(false);
          _1371_v51 = _nw254;
          let _index259 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1370_v50).length));
          (_1370_v50)[_index259] = _1371_v51;
          let _index260 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_1370_v50).length));
          (_1370_v50)[_index260] = _1371_v51;
          let _index261 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_1366_v48).length));
          (_1366_v48)[_index261] = _1369_v49;
        }
        let _1372_v52;
        _1372_v52 = _dafny.MultiSet.fromElements(_1298_v1, ((p0) ? (_1307_i1) : (_1298_v1)), (new BigNumber(-937)).plus(_1307_i1), _1298_v1);
        let _1373_v53;
        _1373_v53 = _dafny.Seq.of(new BigNumber(-710), _1298_v1);
        let _1374_v54;
        _1374_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1373_v53)).cardinality()),p0);
        let _1375_v55;
        _1375_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v54,_1307_i1);
        let _1376_v56;
        _1376_v56 = _dafny.Seq.of(_1297_v0);
        _1372_v52 = (_1372_v52).update(new BigNumber((_1375_v55).length), _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat((_1376_v56)[_module.__default.safeIndex(_1298_v1, new BigNumber((_1376_v56).length))], _1297_v0)).length))));
        let _1377_v57;
        let _nw255 = Array((new BigNumber(26)).toNumber()).fill([]);
        _1377_v57 = _nw255;
        let _1378_v58;
        _1378_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1377_v57,_dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), function (_1379_i5) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }));
        let _1380_v59;
        _1380_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1298_v1,_1377_v57);
        let _1381_v60;
        _1381_v60 = _dafny.Seq.of(p0);
        let _1382_v61;
        _1382_v61 = _module.D12.create_DC33((((_1380_v59).contains(new BigNumber((_1381_v60).length))) ? ((_1380_v59).get(new BigNumber((_1381_v60).length))) : (_1377_v57)));
        _1297_v0 = (((_1378_v58).contains((_1382_v61).dtor_cf58)) ? ((_1378_v58).get((_1382_v61).dtor_cf58)) : (_dafny.Seq.UnicodeFromString("utfdflun")));
        let _1383_v62;
        _1383_v62 = _dafny.Set.fromElements(_1307_i1);
        let _1384_v63;
        let _nw256 = Array((new BigNumber(20)).toNumber());
        _nw256[(_dafny.ZERO).toNumber()] = p0;
        _nw256[(_dafny.ONE).toNumber()] = p0;
        _nw256[(new BigNumber(2)).toNumber()] = (p0) === (p0);
        _nw256[(new BigNumber(3)).toNumber()] = (_1308_v9).IsDisjointFrom(_1308_v9);
        _nw256[(new BigNumber(4)).toNumber()] = p0;
        _nw256[(new BigNumber(5)).toNumber()] = p0;
        _nw256[(new BigNumber(6)).toNumber()] = !(p0);
        _nw256[(new BigNumber(7)).toNumber()] = p0;
        _nw256[(new BigNumber(8)).toNumber()] = p0;
        _nw256[(new BigNumber(9)).toNumber()] = p0;
        _nw256[(new BigNumber(10)).toNumber()] = !(!(((p0) ? (p0) : (p0))));
        _nw256[(new BigNumber(11)).toNumber()] = false;
        _nw256[(new BigNumber(12)).toNumber()] = (_1298_v1).isLessThan(_1298_v1);
        _nw256[(new BigNumber(13)).toNumber()] = false;
        _nw256[(new BigNumber(14)).toNumber()] = p0;
        _nw256[(new BigNumber(15)).toNumber()] = (p0) === (p0);
        _nw256[(new BigNumber(16)).toNumber()] = p0;
        _nw256[(new BigNumber(17)).toNumber()] = (_1383_v62).IsProperSubsetOf(_module.__default.fm31(p0, globalState));
        _nw256[(new BigNumber(18)).toNumber()] = p0;
        _nw256[(new BigNumber(19)).toNumber()] = (new BigNumber((_1381_v60).length)).isLessThanOrEqualTo((_1373_v53)[_module.__default.safeIndex(_1307_i1, new BigNumber((_1373_v53).length))]);
        _1384_v63 = _nw256;
        let _index262 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1384_v63).length));
        (_1384_v63)[_index262] = p0;
        let _index263 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1384_v63).length));
        (_1384_v63)[_index263] = (_1372_v52).contains(new BigNumber(-628));
      }
      let _hi10 = _1298_v1;
      for (let _1385_i6 = new BigNumber((_dafny.Seq.UnicodeFromString("tfmtbfldy")).length); _1385_i6.isLessThan(_hi10); _1385_i6 = _1385_i6.plus(_dafny.ONE)) {
        (globalState).f5 = ((true) ? (p0) : (false));
        let _1386_v64;
        let _init39 = ((_1387_v1) => function (_1388_i7) {
          return (_1388_i7).plus(_1387_v1);
        })(_1298_v1);
        let _nw257 = Array((new BigNumber(18)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw257.length); _i0_39++) {
          _nw257[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1386_v64 = _nw257;
        let _1389_v65;
        _1389_v65 = _dafny.Seq.of(_1386_v64, _1386_v64);
        _1389_v65 = _1389_v65;
        let _1390_v66;
        _1390_v66 = _dafny.MultiSet.fromElements(p0, p0, p0);
        let _1391_v67;
        _1391_v67 = _module.D11.create_DC30(p0, _1385_i6);
        _1298_v1 = (((_1390_v66).contains((_1391_v67).dtor_cf51)) ? ((_1390_v66).get((_1391_v67).dtor_cf51)) : (_1298_v1));
        _1298_v1 = _1385_i6;
      }
      let _hi11 = _module.__default.safeDivisionInt(_1298_v1, new BigNumber((_dafny.Seq.UnicodeFromString("mkxb")).length));
      for (let _1392_i8 = _1298_v1; _1392_i8.isLessThan(_hi11); _1392_i8 = _1392_i8.plus(_dafny.ONE)) {
        let _1393_v68;
        _1393_v68 = new _dafny.CodePoint('o'.codePointAt(0));
        (globalState).f2 = !_dafny.Seq.contains(_1297_v0, _1393_v68);
        let _1394_v69;
        _1394_v69 = _dafny.Seq.of(_1298_v1, _1392_i8);
        let _1395_v70;
        _1395_v70 = _module.D3.create_DC9(_1394_v69);
        let _1396_v71;
        _1396_v71 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1395_v70).dtor_cf18);
        _1298_v1 = new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(((((_1396_v71).contains(p0)) ? ((_1396_v71).get(p0)) : ((_this).fm6(_module.__default.fm3(globalState), p0, _1393_v68, _dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), function (_1397_i9) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }), globalState)))).length))).cardinality());
        let _1398_v72;
        _1398_v72 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1392_i8);
        let _1399_v73;
        let _nw258 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _1399_v73 = _nw258;
        let _1400_v74;
        _1400_v74 = _module.D2.create_DC8(_1398_v72, _1399_v73, p0, p0, _1298_v1);
        let _pat_let_tv29 = p0;
        let _pat_let_tv30 = _1399_v73;
        let _1401_v75;
        _1401_v75 = _module.D6.create_DC17(function (_pat_let39_0) {
  return function (_1402_dt__update__tmp_h2) {
    return function (_pat_let40_0) {
      return function (_1403_dt__update_hcf15_h0) {
        return function (_pat_let41_0) {
          return function (_1404_dt__update_hcf14_h0) {
            return _module.D2.create_DC8((_1402_dt__update__tmp_h2).dtor_cf13, _1404_dt__update_hcf14_h0, _1403_dt__update_hcf15_h0, (_1402_dt__update__tmp_h2).dtor_cf16, (_1402_dt__update__tmp_h2).dtor_cf17);
          }(_pat_let41_0);
        }(_pat_let_tv30);
      }(_pat_let40_0);
    }(_pat_let_tv29);
  }(_pat_let39_0);
}(_1400_v74));
        let _1405_v76;
        _1405_v76 = _dafny.MultiSet.fromElements(_1401_v75, _1401_v75);
        let _1406_v77;
        _1406_v77 = _dafny.Seq.of((_1405_v76).update(_1401_v75, _module.__default.abs(_1298_v1)));
        let _1407_v78;
        _1407_v78 = _1406_v77;
        _1407_v78 = _1407_v78;
        let _1408_v79;
        let _nw259 = new _module.C2();
        _nw259.__ctor(_1298_v1);
        _1408_v79 = _nw259;
        let _1409_v80;
        _1409_v80 = _dafny.Set.fromElements(_1408_v79, _1408_v79, _1408_v79, _1408_v79);
        _1409_v80 = _1409_v80;
      }
      let _hi12 = _1298_v1;
      for (let _1410_i10 = _1298_v1; _1410_i10.isLessThan(_hi12); _1410_i10 = _1410_i10.plus(_dafny.ONE)) {
        (globalState).f5 = !(false);
        if ((_1298_v1).isLessThanOrEqualTo((_1298_v1).plus(_1410_i10))) {
          let _1411_v81;
          _1411_v81 = _dafny.Set.fromElements(_module.__default.fm0(p0, p0, globalState), false);
          let _1412_v82;
          _1412_v82 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("tfw")).length), new BigNumber((_dafny.Seq.of(new BigNumber((_1411_v81).length))).length), new BigNumber((_module.__default.fm20(p0, !(p0), p0, _1410_i10, globalState)).length), new BigNumber(-685), _1410_i10);
          let _1413_v83;
          _1413_v83 = _module.D1.create_DC5(_1298_v1, false);
          let _1414_v84;
          let _nw260 = new _module.C1();
          _nw260.__ctor(_1410_i10, (((_1412_v82).contains((_1413_v83).dtor_cf10)) ? ((_1412_v82).get((_1413_v83).dtor_cf10)) : (_1410_i10)), ((!(p0)) ? (p0) : (p0)), _dafny.Seq.Concat(_dafny.Seq.of(_1410_i10), _dafny.Seq.of(_1298_v1, _1298_v1, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(787))).cardinality()), _1410_i10)), _1410_i10);
          _1414_v84 = _nw260;
          let _1415_v85;
          let _init40 = function (_1416_i11) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(483)), function (_1417_i12) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            });
          };
          let _nw261 = Array((new BigNumber(14)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw261.length); _i0_40++) {
            _nw261[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1415_v85 = _nw261;
          let _index264 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1415_v85).length));
          (_1415_v85)[_index264] = _dafny.Seq.Concat(_1297_v0, _1297_v0);
          let _index265 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1415_v85).length));
          let _rhs301 = _1298_v1;
          let _rhs302 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vbdbgrpij"), _1297_v0);
          let _lhs238 = _1415_v85;
          let _lhs239 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1415_v85).length));
          _1298_v1 = _rhs301;
          _lhs238[_lhs239] = _rhs302;
          _1298_v1 = _1410_i10;
          let _1418_v86;
          _1418_v86 = _dafny.Seq.of((_1414_v84).f29);
          _1298_v1 = (((_1412_v82).contains((_1414_v84).f30)) ? ((_1412_v82).get((_1414_v84).f30)) : (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1418_v86,p0)).length), new BigNumber(-373))));
          let _1419_v87;
          let _nw262 = Array((new BigNumber(12)).toNumber()).fill(false);
          _1419_v87 = _nw262;
          let _1420_v88;
          let _nw263 = new _module.C0();
          _nw263.__ctor((_1414_v84).f30, _1419_v87);
          _1420_v88 = _nw263;
        } else {
          (globalState).f15 = _module.__default.fm2(globalState);
          let _1421_v89;
          let _nw264 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1421_v89 = _nw264;
          let _index266 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1421_v89).length));
          (_1421_v89)[_index266] = _1297_v0;
          let _1422_v90;
          _1422_v90 = new _dafny.CodePoint('a'.codePointAt(0));
          let _index267 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_1421_v89).length));
          (_1421_v89)[_index267] = _dafny.Seq.update(_1297_v0, _module.__default.safeIndex(new BigNumber(-864), new BigNumber((_1297_v0).length)), _1422_v90);
          let _1423_v91;
          _1423_v91 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_1298_v1, _module.__default.fm3(globalState))).length));
          let _1424_v92;
          _1424_v92 = _dafny.Seq.of(p0);
          let _1425_v93;
          _1425_v93 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("prrne"),_1410_i10);
          let _1426_v94;
          _1426_v94 = _dafny.Map.Empty.slice().updateUnsafe((_1424_v92)[_module.__default.safeIndex(_1298_v1, new BigNumber((_1424_v92).length))],new BigNumber(((_1425_v93).update(_dafny.Seq.UnicodeFromString("pn"), _1410_i10)).length));
          let _1427_v95;
          let _init41 = ((_1428_i10) => function (_1429_i13) {
            return _dafny.Seq.of(_1428_i10);
          })(_1410_i10);
          let _nw265 = Array((new BigNumber(25)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw265.length); _i0_41++) {
            _nw265[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1427_v95 = _nw265;
          let _1430_v96;
          _1430_v96 = _module.D2.create_DC8(_1426_v94, _1427_v95, p0, p0, _1410_i10);
          let _1431_v97;
          _1431_v97 = _dafny.Set.fromElements(_1430_v96, _1430_v96);
          let _1432_v98;
          _1432_v98 = _module.D9.create_DC27(p0, ((p0) ? (_1423_v91) : (_dafny.Set.fromElements(new BigNumber((_1431_v97).length)))));
          let _1433_v99;
          _1433_v99 = _dafny.MultiSet.fromElements(p0);
          let _rhs303 = (new BigNumber(214)).isLessThanOrEqualTo(new BigNumber((_1433_v99).cardinality()));
          let _rhs304 = _1432_v98;
          let _rhs305 = _1410_i10;
          let _lhs240 = globalState;
          _lhs240.f2 = _rhs303;
          _1432_v98 = _rhs304;
          _1298_v1 = _rhs305;
          (globalState).f5 = p0;
          _1422_v90 = _1422_v90;
        }
        let _1434_v100;
        _1434_v100 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),(_dafny.ZERO).minus(_1298_v1));
        _1298_v1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_1297_v0).length), new BigNumber((_1434_v100).length)), _module.__default.safeModuloInt(_1298_v1, _1410_i10)));
        let _1435_v101;
        let _nw266 = Array((new BigNumber(8)).toNumber());
        _nw266[(_dafny.ZERO).toNumber()] = p0;
        _nw266[(_dafny.ONE).toNumber()] = p0;
        _nw266[(new BigNumber(2)).toNumber()] = p0;
        _nw266[(new BigNumber(3)).toNumber()] = p0;
        _nw266[(new BigNumber(4)).toNumber()] = p0;
        _nw266[(new BigNumber(5)).toNumber()] = p0;
        _nw266[(new BigNumber(6)).toNumber()] = false;
        _nw266[(new BigNumber(7)).toNumber()] = p0;
        _1435_v101 = _nw266;
        let _1436_v102;
        let _nw267 = Array((new BigNumber(6)).toNumber());
        _nw267[(_dafny.ZERO).toNumber()] = _1435_v101;
        _nw267[(_dafny.ONE).toNumber()] = ((_this).f21).dtor_cf5;
        _nw267[(new BigNumber(2)).toNumber()] = _1435_v101;
        _nw267[(new BigNumber(3)).toNumber()] = _1435_v101;
        _nw267[(new BigNumber(4)).toNumber()] = _1435_v101;
        _nw267[(new BigNumber(5)).toNumber()] = _1435_v101;
        _1436_v102 = _nw267;
        let _1437_v103;
        _1437_v103 = _module.D6.create_DC15(_1436_v102);
        let _pat_let_tv31 = _1436_v102;
        let _1438_v104;
        _1438_v104 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(705),function (_pat_let42_0) {
          return function (_1439_dt__update__tmp_h3) {
            return function (_pat_let43_0) {
              return function (_1440_dt__update_hcf28_h0) {
                return _module.D6.create_DC15(_1440_dt__update_hcf28_h0);
              }(_pat_let43_0);
            }(_pat_let_tv31);
          }(_pat_let42_0);
        }(_1437_v103));
        _1438_v104 = (_1438_v104).update((_1410_i10).minus(_1410_i10), _1437_v103);
      }
      let _1441_v105;
      _1441_v105 = _dafny.MultiSet.fromElements(_1298_v1);
      let _1442_v106;
      _1442_v106 = _module.D4.create_DC12(p0, !(p0), _1298_v1);
      let _1443_v107;
      _1443_v107 = _dafny.MultiSet.fromElements(p0);
      let _1444_v108;
      _1444_v108 = _dafny.Seq.of(_1298_v1);
      let _1445_v109;
      let _nw268 = new _module.C3();
      _nw268.__ctor(_1297_v0, p0, _1444_v108);
      _1445_v109 = _nw268;
      let _1446_v110;
      _1446_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1445_v109,p0);
      let _1447_v111;
      _1447_v111 = _module.D8.create_DC21(_1298_v1, p0, true, p0, _1298_v1);
      let _1448_v112;
      _1448_v112 = _dafny.Set.fromElements(_1298_v1);
      let _1449_v113;
      _1449_v113 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1298_v1);
      let _1450_v114;
      _1450_v114 = _module.D11.create_DC31(p0, _1298_v1, _1298_v1, p0);
      let _1451_v115;
      _1451_v115 = _dafny.Seq.of((_1450_v114).dtor_cf53);
      let _1452_v116;
      _1452_v116 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(505),new BigNumber(617));
      let _1453_v117;
      let _nw269 = Array((new BigNumber(23)).toNumber());
      _nw269[(_dafny.ZERO).toNumber()] = (_1298_v1).plus(new BigNumber((_1297_v0).length));
      _nw269[(_dafny.ONE).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(2)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(3)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(4)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(5)).toNumber()] = (((_1441_v105).contains(_1298_v1)) ? ((_1441_v105).get(_1298_v1)) : ((_1442_v106).dtor_cf23));
      _nw269[(new BigNumber(6)).toNumber()] = (_1298_v1).plus(new BigNumber((_1443_v107).cardinality()));
      _nw269[(new BigNumber(7)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(8)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_1298_v1);
      _nw269[(new BigNumber(10)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_1298_v1);
      _nw269[(new BigNumber(12)).toNumber()] = new BigNumber(((_1446_v110).Merge((_1446_v110).update(_1445_v109, p0))).length);
      _nw269[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_1447_v111).dtor_cf37,new BigNumber((_1448_v112).length))).Merge(_1449_v113)).length));
      _nw269[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_1298_v1);
      _nw269[(new BigNumber(15)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(16)).toNumber()] = new BigNumber(((_module.__default.fm36(_1298_v1, p0, globalState)).Intersect(_1443_v107)).cardinality());
      _nw269[(new BigNumber(17)).toNumber()] = ((p0) ? (_1298_v1) : (_1298_v1));
      _nw269[(new BigNumber(18)).toNumber()] = (_1298_v1).plus(new BigNumber((_1451_v115).length));
      _nw269[(new BigNumber(19)).toNumber()] = _1298_v1;
      _nw269[(new BigNumber(20)).toNumber()] = new BigNumber((_1452_v116).length);
      _nw269[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(_1298_v1, new BigNumber(153));
      _nw269[(new BigNumber(22)).toNumber()] = _1298_v1;
      _1453_v117 = _nw269;
      let _index268 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1453_v117).length));
      (_1453_v117)[_index268] = _1298_v1;
      let _index269 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_1453_v117).length));
      (_1453_v117)[_index269] = (_dafny.ZERO).minus(_module.__default.fm3(globalState));
      return;
    }
    m4(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1454_v0;
      _1454_v0 = false;
      if (_1454_v0) {
        let _source15 = (_this).f21;
        if (_source15.is_DC1) {
          let _1455___mcc_h0 = (_source15).cf1;
          let _1456_cf1 = _1455___mcc_h0;
          r3 = _module.__default.fm3(globalState);
          let _1457_v1;
          _1457_v1 = new BigNumber(408);
          r3 = (_1457_v1).multipliedBy(_1457_v1);
          let _1458_v2;
          _1458_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(18),_1457_v1);
          let _1459_v3;
          _1459_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1454_v0,new BigNumber((_1458_v2).length));
          let _1460_v4;
          _1460_v4 = _dafny.MultiSet.fromElements(_1459_v3);
          let _1461_v5;
          _1461_v5 = _dafny.Seq.of(_1460_v4, _dafny.MultiSet.fromElements(_1459_v3));
          let _1462_v6;
          _1462_v6 = _module.D4.create_DC11((_1461_v5)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1461_v5).length))]);
          _1462_v6 = _1462_v6;
          let _1463_v7;
          _1463_v7 = _dafny.Seq.of(_1457_v1);
          let _1464_v8;
          _1464_v8 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1465_v9;
          _1465_v9 = _module.D0.create_DC1(_dafny.Seq.update(_1456_cf1, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("mmnwsjn"), _module.__default.safeIndex(_1457_v1, new BigNumber((_dafny.Seq.UnicodeFromString("mmnwsjn")).length)), _1464_v8)).length), new BigNumber((_1456_cf1).length)), _1464_v8));
          let _1466_v10;
          _1466_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1457_v1,_1465_v9);
          let _1467_v11;
          let _nw270 = new _module.C1();
          _nw270.__ctor(new BigNumber(931), _1457_v1, _1454_v0, _1463_v7, new BigNumber((_1466_v10).length));
          _1467_v11 = _nw270;
          let _1468_v12;
          _1468_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1457_v1,_1467_v11);
          _1468_v12 = (_1468_v12).update(new BigNumber((_1458_v2).length), _1467_v11);
        } else if (_source15.is_DC2) {
          let _1469___mcc_h1 = (_source15).cf2;
          let _1470___mcc_h2 = (_source15).cf3;
          let _1471___mcc_h3 = (_source15).cf4;
          let _1472___mcc_h4 = (_source15).cf5;
          let _1473_cf5 = _1472___mcc_h4;
          let _1474_cf4 = _1471___mcc_h3;
          let _1475_cf3 = _1470___mcc_h2;
          let _1476_cf2 = _1469___mcc_h1;
          let _1477_v13;
          _1477_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1475_cf3,_1454_v0);
          let _1478_v14;
          _1478_v14 = _dafny.Seq.UnicodeFromString("vxaubieu");
          r2 = (((_1477_v13).contains((_dafny.ZERO).minus((_dafny.ZERO).minus((_1476_cf2).plus(new BigNumber((_1478_v14).length)))))) ? ((_1477_v13).get((_dafny.ZERO).minus((_dafny.ZERO).minus((_1476_cf2).plus(new BigNumber((_1478_v14).length)))))) : (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-21)), ((_1479_cf4) => function (_1480_i0) {
            return _1479_cf4;
          })(_1474_cf4)), _dafny.Seq.of(_1476_cf2))));
          let _1481_v15;
          _1481_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1454_v0,_1478_v14);
          let _1482_v16;
          _1482_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1481_v15).length),_1475_cf3);
          let _1483_v17;
          _1483_v17 = _dafny.Set.fromElements(_1482_v16);
          _1483_v17 = (_1483_v17).Union(_1483_v17);
          _1477_v13 = _1477_v13;
          let _1484_v18;
          _1484_v18 = _dafny.Seq.of(_1454_v0, !(_1454_v0));
          let _index270 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_1473_cf5).length));
          (_1473_cf5)[_index270] = (_1484_v18)[_module.__default.safeIndex(new BigNumber((_1484_v18).length), new BigNumber((_1484_v18).length))];
          let _1485_v19;
          _1485_v19 = new _dafny.CodePoint('m'.codePointAt(0));
          let _index271 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_1473_cf5).length));
          let _rhs306 = _1485_v19;
          let _rhs307 = _1454_v0;
          let _rhs308 = ((_dafny.ZERO).minus(_1476_cf2)).multipliedBy(_module.__default.safeModuloInt(_1474_cf4, new BigNumber((_1484_v18).length)));
          let _rhs309 = _1485_v19;
          let _lhs241 = globalState;
          let _lhs242 = _1473_cf5;
          let _lhs243 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_1473_cf5).length));
          let _lhs244 = globalState;
          _lhs241.f15 = _rhs306;
          _lhs242[_lhs243] = _rhs307;
          _1475_cf3 = _rhs308;
          _lhs244.f15 = _rhs309;
        } else {
          let _1486___mcc_h5 = (_source15).cf0;
          let _1487_cf0 = _1486___mcc_h5;
          (globalState).f5 = _1454_v0;
          _1487_cf0 = _1487_cf0;
          r3 = new BigNumber((_1487_cf0).length);
          let _1488_v20;
          let _nw271 = Array((new BigNumber(14)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1488_v20 = _nw271;
          let _1489_v21;
          _1489_v21 = new BigNumber(-675);
          let _1490_v22;
          _1490_v22 = _dafny.MultiSet.fromElements(_1489_v21, _1489_v21);
          let _1491_v23;
          _1491_v23 = _dafny.Seq.of(_1489_v21, _1489_v21, _1489_v21);
          let _index272 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1488_v20).length));
          (_1488_v20)[_index272] = (_dafny.MultiSet.fromElements(new BigNumber((_1490_v22).cardinality()))).Intersect(_dafny.MultiSet.FromArray(_1491_v23));
          let _index273 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1488_v20).length));
          (_1488_v20)[_index273] = _1490_v22;
        }
        let _1492_v24;
        let _nw272 = Array((new BigNumber(24)).toNumber());
        _1492_v24 = _nw272;
        let _1493_v25;
        _1493_v25 = new BigNumber(645);
        let _1494_v26;
        let _nw273 = new _module.C4();
        _nw273.__ctor(new BigNumber(52), _1454_v0);
        _1494_v26 = _nw273;
        let _1495_v27;
        _1495_v27 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1496_v28;
        _1496_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1494_v26,new _dafny.CodePoint('m'.codePointAt(0)));
        let _1497_v29;
        _1497_v29 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1494_v26,_1495_v27), _1496_v28, _1496_v28);
        let _1498_v30;
        _1498_v30 = _dafny.Seq.of(new BigNumber(((_1497_v29)[_module.__default.safeIndex((_1494_v26).f22, new BigNumber((_1497_v29).length))]).length));
        let _1499_v31;
        let _nw274 = new _module.C1();
        _nw274.__ctor(_1493_v25, _1493_v25, true, _1498_v30, new BigNumber((_1498_v30).length));
        _1499_v31 = _nw274;
        let _index274 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_1492_v24).length));
        (_1492_v24)[_index274] = _1499_v31;
        let _index275 = _module.__default.safeIndex(new BigNumber(804), new BigNumber((_1492_v24).length));
        (_1492_v24)[_index275] = _1499_v31;
        let _1500_v32;
        let _nw275 = Array((new BigNumber(12)).toNumber()).fill(false);
        _1500_v32 = _nw275;
        let _1501_v33;
        _1501_v33 = _dafny.MultiSet.fromElements(_1500_v32, _1500_v32);
        let _1502_v34;
        let _nw276 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _1502_v34 = _nw276;
        let _1503_v35;
        _1503_v35 = _dafny.Set.fromElements(_1502_v34);
        let _1504_v36;
        _1504_v36 = _module.D11.create_DC30(_1494_v26.f23, (_1494_v26).f22);
        let _1505_v37;
        _1505_v37 = _dafny.MultiSet.fromElements(!((_1504_v36).dtor_cf51), _1494_v26.f23, _1494_v26.f23);
        let _1506_v38;
        _1506_v38 = _dafny.Set.fromElements(_1494_v26.f23, !(_1454_v0));
        let _1507_v39;
        _1507_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1494_v26.f23,(_1494_v26).f22);
        let _1508_v40;
        let _nw277 = Array((new BigNumber(16)).toNumber());
        _nw277[(_dafny.ZERO).toNumber()] = _1493_v25;
        _nw277[(_dafny.ONE).toNumber()] = (_1499_v31).f30;
        _nw277[(new BigNumber(2)).toNumber()] = new BigNumber((_1501_v33).cardinality());
        _nw277[(new BigNumber(3)).toNumber()] = (_1499_v31).fm15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), ((_1509_v31) => function (_1510_i1) {
          return (_1509_v31).f30;
        })(_1499_v31)), new BigNumber((_1503_v35).length), _1493_v25, _1495_v27, globalState);
        _nw277[(new BigNumber(4)).toNumber()] = _1493_v25;
        _nw277[(new BigNumber(5)).toNumber()] = (_1499_v31).f30;
        _nw277[(new BigNumber(6)).toNumber()] = _1493_v25;
        _nw277[(new BigNumber(7)).toNumber()] = new BigNumber((_1505_v37).cardinality());
        _nw277[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(((_1494_v26.f23) ? (_1493_v25) : (new BigNumber((_1506_v38).length))));
        _nw277[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1500_v32,_1454_v0)).length);
        _nw277[(new BigNumber(10)).toNumber()] = new BigNumber(32);
        _nw277[(new BigNumber(11)).toNumber()] = (_1499_v31).f30;
        _nw277[(new BigNumber(12)).toNumber()] = new BigNumber(-898);
        _nw277[(new BigNumber(13)).toNumber()] = (_1499_v31).f29;
        _nw277[(new BigNumber(14)).toNumber()] = (_1494_v26).f22;
        _nw277[(new BigNumber(15)).toNumber()] = (((_1507_v39).contains(_1454_v0)) ? ((_1507_v39).get(_1454_v0)) : (new BigNumber(116)));
        _1508_v40 = _nw277;
        _1502_v34 = _1502_v34;
        let _1511_v41;
        _1511_v41 = _dafny.MultiSet.fromElements((_1499_v31).f30, (_1499_v31).f29);
        let _index276 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1502_v34).length));
        (_1502_v34)[_index276] = new BigNumber(((_1511_v41).Intersect(_1511_v41)).cardinality());
        let _index277 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1502_v34).length));
        (_1502_v34)[_index277] = (_1499_v31).f29;
        let _1512_v42;
        _1512_v42 = _dafny.Seq.UnicodeFromString("rhpqpdgw");
        _1493_v25 = (_dafny.ZERO).minus(new BigNumber((_1512_v42).length));
      } else {
        let _1513_v43;
        _1513_v43 = _dafny.Seq.UnicodeFromString("aebefuri");
        let _1514_v44;
        _1514_v44 = _module.D0.create_DC0(_1513_v43);
        _1514_v44 = _1514_v44;
        let _1515_v45;
        _1515_v45 = new BigNumber(700);
        let _1516_v46;
        _1516_v46 = _dafny.Seq.of(_1515_v45);
        let _1517_v47;
        let _nw278 = new _module.C4();
        _nw278.__ctor(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1516_v46).length))).cardinality()), ((_1454_v0) ? (_1454_v0) : (_module.__default.fm0(_1454_v0, _1454_v0, globalState))));
        _1517_v47 = _nw278;
        let _1518_v48;
        let _nw279 = new _module.C2();
        _nw279.__ctor(new BigNumber((_1516_v46).length));
        _1518_v48 = _nw279;
        let _1519_v49;
        _1519_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1454_v0,_1517_v47.f23);
        let _1520_v50;
        _1520_v50 = _dafny.Seq.of((((_1519_v49).contains(_1454_v0)) ? ((_1519_v49).get(_1454_v0)) : (_1454_v0)), _1454_v0);
        let _1521_v51;
        _1521_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1520_v50).length),_1518_v48);
        let _1522_v52;
        _1522_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1454_v0,new BigNumber((_1521_v51).length));
        let _1523_v53;
        _1523_v53 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1524_v54;
        _1524_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1523_v53,_1454_v0);
        let _1525_v55;
        let _nw280 = new _module.C1();
        _nw280.__ctor(new BigNumber((_1522_v52).length), (_1518_v48).f25, (((_1524_v54).contains(_1523_v53)) ? ((_1524_v54).get(_1523_v53)) : (_1517_v47.f23)), _1516_v46, (_1518_v48).f25);
        _1525_v55 = _nw280;
        let _1526_v56;
        let _nw281 = new _module.C3();
        _nw281.__ctor(_1513_v43, _1517_v47.f23, ((_1454_v0) ? (_1516_v46) : ((_this).fm6(new BigNumber(822), !(!(_1517_v47.f23)), _1523_v53, _1513_v43, globalState))));
        _1526_v56 = _nw281;
      }
      let _1527_v57;
      _1527_v57 = new _dafny.CodePoint('h'.codePointAt(0));
      let _1528_v58;
      _1528_v58 = new BigNumber(461);
      let _1529_v59;
      _1529_v59 = _dafny.Seq.of(_1454_v0);
      let _1530_v60;
      let _nw282 = new _module.C1();
      _nw282.__ctor(_1528_v58, _1528_v58, _1454_v0, _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_1529_v59)).cardinality()), _1528_v58, new BigNumber(-126)), _1528_v58);
      _1530_v60 = _nw282;
      let _1531_v61;
      _1531_v61 = _module.D8.create_DC22(_1454_v0, _1530_v60, _1454_v0);
      let _1532_v62;
      _1532_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1527_v57,(_1531_v61).dtor_cf40);
      let _1533_v63;
      _1533_v63 = _module.D9.create_DC25(_1532_v62);
      let _1534_v64;
      _1534_v64 = _dafny.Set.fromElements(_1533_v63, _1533_v63, _1533_v63);
      let _1535_v65;
      _1535_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1533_v63,_module.__default.fm3(globalState));
      let _1536_v67;
      let _nw283 = new _module.C5();
      _nw283.__ctor(!(_1534_v64).equals(function () {
        let _coll61 = new _dafny.Set();
        for (const _compr_61 of (_1535_v65).Keys.Elements) {
          let _1537_v66 = _compr_61;
          if ((_1535_v65).contains(_1537_v66)) {
            _coll61.add(_1537_v66);
          }
        }
        return _coll61;
      }()), _dafny.Seq.of(_1528_v58, _1528_v58, _1528_v58, _1528_v58));
      _1536_v67 = _nw283;
      _1528_v58 = _module.__default.safeDivisionInt(_1528_v58, _1528_v58);
      let _1538_v68;
      _1538_v68 = _dafny.Seq.UnicodeFromString("kp");
      let _1539_v69;
      _1539_v69 = _module.D0.create_DC1(_1538_v68);
      let _1540_v70;
      _1540_v70 = _dafny.MultiSet.fromElements(new BigNumber(987));
      _1539_v69 = _module.__default.fm37(_1454_v0, new BigNumber((_1540_v70).cardinality()), _1454_v0, globalState);
      let _hi13 = _1528_v58;
      for (let _1541_i2 = _module.__default.fm3(globalState); _1541_i2.isLessThan(_hi13); _1541_i2 = _1541_i2.plus(_dafny.ONE)) {
        let _1542_v71;
        _1542_v71 = _module.D3.create_DC10((_1529_v59)[_module.__default.safeIndex(_1541_i2, new BigNumber((_1529_v59).length))]);
        let _source16 = _1542_v71;
        if (_source16.is_DC10) {
          let _1543___mcc_h6 = (_source16).cf19;
          let _1544_cf19 = _1543___mcc_h6;
          let _1545_v72;
          _1545_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1528_v58,_module.__default.safeDivisionInt(_1528_v58, (_dafny.ZERO).minus(_1541_i2)));
          _1545_v72 = (_1545_v72).update(_1541_i2, ((_1530_v60).f17)[_module.__default.safeIndex(new BigNumber((_1538_v68).length), new BigNumber(((_1530_v60).f17).length))]);
          r3 = _1528_v58;
          r3 = _1528_v58;
          let _1546_v73;
          _1546_v73 = _dafny.Set.fromElements(_1528_v58);
          let _1547_v74;
          _1547_v74 = _dafny.Set.fromElements(_1546_v73);
          let _1548_v75;
          _1548_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1527_v57,_1527_v57);
          let _1549_v76;
          let _init42 = ((_1550_v60) => function (_1551_i3) {
            return !(_1550_v60.f16);
          })(_1530_v60);
          let _nw284 = Array((new BigNumber(4)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw284.length); _i0_42++) {
            _nw284[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1549_v76 = _nw284;
          let _1552_v77;
          _1552_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1541_i2,_1549_v76);
          let _1553_v78;
          let _init43 = ((_1554_i2) => function (_1555_i4) {
            return (_1555_i4).multipliedBy(_1554_i2);
          })(_1541_i2);
          let _nw285 = Array((new BigNumber(11)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw285.length); _i0_43++) {
            _nw285[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1553_v78 = _nw285;
          let _1556_v79;
          _1556_v79 = _module.D12.create_DC34((((_1552_v77).contains(new BigNumber(63))) ? ((_1552_v77).get(new BigNumber(63))) : (_1549_v76)), _1541_i2, _1553_v78, _1527_v57);
          let _1557_v80;
          let _nw286 = Array((new BigNumber(26)).toNumber());
          _nw286[(_dafny.ZERO).toNumber()] = _1530_v60.f16;
          _nw286[(_dafny.ONE).toNumber()] = _1544_cf19;
          _nw286[(new BigNumber(2)).toNumber()] = _1530_v60.f16;
          _nw286[(new BigNumber(3)).toNumber()] = !(_1530_v60.f16);
          _nw286[(new BigNumber(4)).toNumber()] = (_module.__default.fm3(globalState)).isLessThan(_1541_i2);
          _nw286[(new BigNumber(5)).toNumber()] = (_1547_v74).equals(_1547_v74);
          _nw286[(new BigNumber(6)).toNumber()] = _module.__default.fm0(false, true, globalState);
          _nw286[(new BigNumber(7)).toNumber()] = _module.__default.fm0(_1454_v0, false, globalState);
          _nw286[(new BigNumber(8)).toNumber()] = (_1541_i2).isLessThanOrEqualTo(_1528_v58);
          _nw286[(new BigNumber(9)).toNumber()] = false;
          _nw286[(new BigNumber(10)).toNumber()] = _1544_cf19;
          _nw286[(new BigNumber(11)).toNumber()] = _1544_cf19;
          _nw286[(new BigNumber(12)).toNumber()] = !_dafny.areEqual(_dafny.Seq.of(_1527_v57, (((_1548_v75).contains((_1556_v79).dtor_cf62)) ? ((_1548_v75).get((_1556_v79).dtor_cf62)) : (_module.__default.fm2(globalState))), _1527_v57, _1527_v57, _1527_v57), _1538_v68);
          _nw286[(new BigNumber(13)).toNumber()] = !(_1454_v0) || (!(_1454_v0));
          _nw286[(new BigNumber(14)).toNumber()] = ((_1530_v60.f16) ? (_1530_v60.f16) : (_1530_v60.f16));
          _nw286[(new BigNumber(15)).toNumber()] = _1530_v60.f16;
          _nw286[(new BigNumber(16)).toNumber()] = _1530_v60.f16;
          _nw286[(new BigNumber(17)).toNumber()] = _1530_v60.f16;
          _nw286[(new BigNumber(18)).toNumber()] = _1544_cf19;
          _nw286[(new BigNumber(19)).toNumber()] = !(_1530_v60.f16) || (_1530_v60.f16);
          _nw286[(new BigNumber(20)).toNumber()] = _1454_v0;
          _nw286[(new BigNumber(21)).toNumber()] = (_1536_v67).fm8(_1538_v68, globalState);
          _nw286[(new BigNumber(22)).toNumber()] = false;
          _nw286[(new BigNumber(23)).toNumber()] = false;
          _nw286[(new BigNumber(24)).toNumber()] = false;
          _nw286[(new BigNumber(25)).toNumber()] = ((!(!(_1454_v0))) ? (_1530_v60.f16) : (_1454_v0));
          _1557_v80 = _nw286;
          let _index278 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1549_v76).length));
          (_1549_v76)[_index278] = _1530_v60.f16;
          let _index279 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_1549_v76).length));
          (_1549_v76)[_index279] = _1544_cf19;
        } else {
          let _1558___mcc_h7 = (_source16).cf18;
          let _1559_cf18 = _1558___mcc_h7;
          (globalState).f15 = _1527_v57;
          _1540_v70 = _1540_v70;
          let _1560_v81;
          _1560_v81 = _dafny.MultiSet.fromElements(_1454_v0, _1454_v0, false);
          let _1561_v82;
          _1561_v82 = _module.D1.create_DC3((_dafny.MultiSet.FromArray(_1529_v59)).Intersect(_1560_v81));
          _1561_v82 = ((_1530_v60.f16) ? (_1561_v82) : (_1561_v82));
          let _1562_v83;
          _1562_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1528_v58,_1454_v0);
          let _1563_v84;
          _1563_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1527_v57,(_1562_v83).Merge(_1562_v83));
          _1563_v84 = (_1563_v84).update(_1527_v57, (_module.__default.fm38(_1533_v63, _1541_i2, globalState)).Merge(_1562_v83));
        }
        let _1564_v85;
        let _nw287 = new _module.C1();
        _nw287.__ctor(_1541_i2, _1528_v58, _1530_v60.f16, (_1530_v60).f17, _1528_v58);
        _1564_v85 = _nw287;
        let _1565_v86;
        _1565_v86 = _module.D1.create_DC5(_1528_v58, _1454_v0);
        _1454_v0 = (_module.__default.safeDivisionInt((_1564_v85).f30, new BigNumber((_dafny.Set.fromElements((_1536_v67).fm8(_1538_v68, globalState), true)).length))).isLessThan((_1565_v86).dtor_cf10);
        let _1566_v87;
        _1566_v87 = _dafny.Set.fromElements(_1454_v0, _1530_v60.f16, _1530_v60.f16);
        let _1567_v88;
        _1567_v88 = _module.D11.create_DC31(_1454_v0, (_1564_v85).f29, (_dafny.ZERO).minus((_1564_v85).fm14(_1541_i2, _1528_v58, _1527_v57, globalState)), _1454_v0);
        let _1568_v89;
        _1568_v89 = _dafny.Set.fromElements(_1567_v88);
        let _1569_v90;
        _1569_v90 = _dafny.Seq.of(_1568_v89);
        let _1570_v91;
        _1570_v91 = _dafny.Map.Empty.slice().updateUnsafe((_1564_v85).f29,_1530_v60.f16);
        let _1571_v92;
        let _init44 = ((_1572_v60) => function (_1573_i5) {
          return _1572_v60.f16;
        })(_1530_v60);
        let _nw288 = Array((new BigNumber(22)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw288.length); _i0_44++) {
          _nw288[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1571_v92 = _nw288;
        let _1574_v93;
        let _nw289 = new _module.C0();
        _nw289.__ctor(_1528_v58, _1571_v92);
        _1574_v93 = _nw289;
        let _1575_v94;
        let _nw290 = Array((new BigNumber(9)).toNumber());
        _nw290[(_dafny.ZERO).toNumber()] = new BigNumber((_1566_v87).length);
        _nw290[(_dafny.ONE).toNumber()] = _1541_i2;
        _nw290[(new BigNumber(2)).toNumber()] = _1541_i2;
        _nw290[(new BigNumber(3)).toNumber()] = new BigNumber(((_1569_v90)[_module.__default.safeIndex((_1564_v85).f29, new BigNumber((_1569_v90).length))]).length);
        _nw290[(new BigNumber(4)).toNumber()] = _1528_v58;
        _nw290[(new BigNumber(5)).toNumber()] = _1528_v58;
        _nw290[(new BigNumber(6)).toNumber()] = _1528_v58;
        _nw290[(new BigNumber(7)).toNumber()] = (_1564_v85).f29;
        _nw290[(new BigNumber(8)).toNumber()] = (_module.D9.create_DC26(new BigNumber((_1570_v91).length), _1574_v93)).dtor_cf45;
        _1575_v94 = _nw290;
        let _1576_v95;
        _1576_v95 = _module.D6.create_DC18(_1575_v94);
        let _source17 = _1576_v95;
        if (_source17.is_DC16) {
          let _1577___mcc_h8 = (_source17).cf29;
          let _1578___mcc_h9 = (_source17).cf30;
          let _1579_cf30 = _1578___mcc_h9;
          let _1580_cf29 = _1577___mcc_h8;
          let _1581_v96;
          let _nw291 = new _module.C4();
          _nw291.__ctor((_1564_v85).f29, _1579_cf30);
          _1581_v96 = _nw291;
          r2 = _1579_cf30;
          let _1582_v97;
          _1582_v97 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),(_dafny.ZERO).minus(_1541_i2));
          let _1583_v98;
          _1583_v98 = _module.D11.create_DC29(_1527_v57);
          _1582_v97 = (_1582_v97).update((_1583_v98).dtor_cf50, (_dafny.ZERO).minus((_1564_v85).f29));
          _1528_v58 = (_module.D1.create_DC5((_1564_v85).f29, false)).dtor_cf10;
        } else if (_source17.is_DC17) {
          let _1584___mcc_h10 = (_source17).cf31;
          let _1585_cf31 = _1584___mcc_h10;
          _1528_v58 = (_1564_v85).f30;
          let _1586_v99;
          let _nw292 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
          _1586_v99 = _nw292;
          let _index280 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_1586_v99).length));
          (_1586_v99)[_index280] = _1566_v87;
          let _index281 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_1586_v99).length));
          (_1586_v99)[_index281] = _1566_v87;
          _1528_v58 = _module.__default.safeDivisionInt((_1564_v85).f29, (_1564_v85).f30);
          r3 = (new BigNumber(((function () {
            let _coll62 = new _dafny.Map();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(119), new BigNumber(403))) {
              let _1587_v100 = _compr_62;
              if (((new BigNumber(119)).isLessThanOrEqualTo(_1587_v100)) && ((_1587_v100).isLessThan(new BigNumber(403)))) {
                _coll62.push([(_1587_v100).multipliedBy((_1564_v85).f29),_1530_v60.f16]);
              }
            }
            return _coll62;
          }()).Merge((_1570_v91).update((_1574_v93).f26, _1454_v0))).length)).plus((_1564_v85).f29);
        } else if (_source17.is_DC18) {
          let _1588___mcc_h11 = (_source17).cf32;
          let _1589_cf32 = _1588___mcc_h11;
          let _1590_v101;
          _1590_v101 = _dafny.Map.Empty.slice().updateUnsafe((_1564_v85).f30,new _dafny.CodePoint('m'.codePointAt(0)));
          let _1591_v102;
          let _nw293 = Array((new BigNumber(16)).toNumber());
          _nw293[(_dafny.ZERO).toNumber()] = (((_1590_v101).contains((_1564_v85).f29)) ? ((_1590_v101).get((_1564_v85).f29)) : (_1527_v57));
          _nw293[(_dafny.ONE).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(2)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(3)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(4)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(5)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(6)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(7)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(8)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(9)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(10)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(11)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
          _nw293[(new BigNumber(13)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(14)).toNumber()] = _1527_v57;
          _nw293[(new BigNumber(15)).toNumber()] = _1527_v57;
          _1591_v102 = _nw293;
          let _index282 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_1591_v102).length));
          (_1591_v102)[_index282] = _1527_v57;
          let _index283 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_1591_v102).length));
          (_1591_v102)[_index283] = _1527_v57;
          let _index284 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1575_v94).length));
          (_1575_v94)[_index284] = ((_1574_v93).f26).multipliedBy(_1541_i2);
          let _1592_v103;
          let _nw294 = Array((new BigNumber(7)).toNumber()).fill([]);
          _1592_v103 = _nw294;
          let _1593_v104;
          _1593_v104 = _dafny.Map.Empty.slice().updateUnsafe((_1564_v85).f30,(_1530_v60).f17);
          let _index285 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1575_v94).length));
          let _rhs310 = (_dafny.ZERO).minus((_1564_v85).fm15((((_1593_v104).contains(new BigNumber((_1540_v70).cardinality()))) ? ((_1593_v104).get(new BigNumber((_1540_v70).cardinality()))) : ((_1530_v60).f17)), (_1528_v58).minus((_1574_v93).f26), (_1574_v93).f26, _1527_v57, globalState));
          let _rhs311 = _1592_v103;
          let _rhs312 = (_1564_v85).f30;
          let _lhs245 = _1575_v94;
          let _lhs246 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_1575_v94).length));
          _lhs245[_lhs246] = _rhs310;
          _1592_v103 = _rhs311;
          _1528_v58 = _rhs312;
          let _1594_v105;
          let _nw295 = new _module.C4();
          _nw295.__ctor(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(241)), ((_1595_v57) => function (_1596_i6) {
            return _1595_v57;
          })(_1527_v57)), _dafny.Seq.update(_1538_v68, _module.__default.safeIndex((_1564_v85).f29, new BigNumber((_1538_v68).length)), _1527_v57))).length), ((_1530_v60.f16) ? (false) : (_1454_v0)));
          _1594_v105 = _nw295;
          _1528_v58 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_1575_v94)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1575_v94).length))], new BigNumber(100)), (_1564_v85).f29);
        } else {
          let _1597___mcc_h12 = (_source17).cf28;
          let _1598_cf28 = _1597___mcc_h12;
          let _1599_v106;
          _1599_v106 = _dafny.MultiSet.fromElements(_1530_v60.f16);
          let _1600_v107;
          _1600_v107 = _module.D0.create_DC2((_1564_v85).f29, (new BigNumber((_module.__default.fm20(!(_1454_v0), _1530_v60.f16, !(_1530_v60.f16), (_module.D4.create_DC12(_1454_v0, true, new BigNumber((_1599_v106).cardinality()))).dtor_cf23, globalState)).length)).multipliedBy((_1564_v85).f30), new BigNumber((_1529_v59).length), _1574_v93.f27);
          _1600_v107 = _1600_v107;
          let _1601_v108;
          let _init45 = ((_1602_v57) => function (_1603_i7) {
            return _1602_v57;
          })(_1527_v57);
          let _nw296 = Array((new BigNumber(19)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw296.length); _i0_45++) {
            _nw296[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1601_v108 = _nw296;
          let _1604_v109;
          let _nw297 = new _module.C1();
          _nw297.__ctor((_1574_v93).f26, (_1574_v93).f26, _1530_v60.f16, _dafny.Seq.update((_1530_v60).f17, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_1601_v108, _1601_v108)).cardinality()), new BigNumber(((_1530_v60).f17).length)), new BigNumber(875)), (_1564_v85).f30);
          _1604_v109 = _nw297;
          let _1605_v110;
          _1605_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1604_v109,!(((_1530_v60.f16) ? (_1604_v109.f16) : (_1604_v109.f16))));
          let _1606_v111;
          _1606_v111 = _dafny.Seq.of(_1604_v109);
          _1605_v110 = (_1605_v110).update((_1606_v111)[_module.__default.safeIndex((_1564_v85).f29, new BigNumber((_1606_v111).length))], _1604_v109.f16);
          (_1530_v60).f16 = _1530_v60.f16;
          let _1607_v112;
          _1607_v112 = _dafny.Seq.of(_1538_v68);
          _1529_v59 = _dafny.Seq.update(_1529_v59, _module.__default.safeIndex((_1564_v85).f30, new BigNumber((_1529_v59).length)), !((_1564_v85).f29).isEqualTo(new BigNumber((_1607_v112).length)));
        }
      }
      let _1608_v113;
      _1608_v113 = _dafny.Map.Empty.slice().updateUnsafe(_1528_v58,_1528_v58);
      let _hi14 = _1528_v58;
      for (let _1609_i8 = (new BigNumber((_1608_v113).length)).multipliedBy(new BigNumber(-870)); _1609_i8.isLessThan(_hi14); _1609_i8 = _1609_i8.plus(_dafny.ONE)) {
        let _1610_v114;
        _1610_v114 = _module.D8.create_DC21(_1528_v58, (_1609_i8).isLessThanOrEqualTo(_1609_i8), !(_1528_v58).isEqualTo(_1528_v58), _dafny.Seq.contains(_1529_v59, _1454_v0), (_dafny.ZERO).minus(new BigNumber((_1538_v68).length)));
        let _source18 = _1610_v114;
        if (_source18.is_DC21) {
          let _1611___mcc_h13 = (_source18).cf35;
          let _1612___mcc_h14 = (_source18).cf36;
          let _1613___mcc_h15 = (_source18).cf37;
          let _1614___mcc_h16 = (_source18).cf38;
          let _1615___mcc_h17 = (_source18).cf39;
          let _1616_cf39 = _1615___mcc_h17;
          let _1617_cf38 = _1614___mcc_h16;
          let _1618_cf37 = _1613___mcc_h15;
          let _1619_cf36 = _1612___mcc_h14;
          let _1620_cf35 = _1611___mcc_h13;
          r3 = _1616_cf39;
          _1616_cf39 = _1616_cf39;
          let _1621_v115;
          _1621_v115 = _dafny.Map.Empty.slice().updateUnsafe(_1530_v60.f16,_1620_cf35);
          let _1622_v116;
          _1622_v116 = _dafny.Seq.of(_1621_v115);
          _1622_v116 = _1622_v116;
          (_1530_v60).f16 = _1530_v60.f16;
        } else if (_source18.is_DC22) {
          let _1623___mcc_h18 = (_source18).cf40;
          let _1624___mcc_h19 = (_source18).cf41;
          let _1625___mcc_h20 = (_source18).cf42;
          let _1626_cf42 = _1625___mcc_h20;
          let _1627_cf41 = _1624___mcc_h19;
          let _1628_cf40 = _1623___mcc_h18;
          (globalState).f5 = _1454_v0;
          let _1629_v117;
          _1629_v117 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_1628_cf40, _1628_cf40, globalState),_1538_v68);
          let _1630_v118;
          _1630_v118 = _dafny.Seq.of(new BigNumber((_1629_v117).length), (_1528_v58).multipliedBy(_1528_v58), _1528_v58);
          _1630_v118 = ((false) ? (_dafny.Seq.update((_1627_cf41).f17, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_1627_cf41).f17).length)), new BigNumber(((_1627_cf41).f17).length)), _1609_i8)) : (_module.__default.fm17(globalState)));
          let _1631_v119;
          let _nw298 = new _module.C3();
          _nw298.__ctor(_dafny.Seq.UnicodeFromString("gj"), (_1536_v67).fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), ((_1632_v57) => function (_1633_i9) {
            return _1632_v57;
          })(_1527_v57)), globalState), (_1530_v60).f17);
          _1631_v119 = _nw298;
          let _1634_v120;
          _1634_v120 = _dafny.Map.Empty.slice().updateUnsafe(_1528_v58,_1530_v60.f16);
          let _1635_v121;
          let _nw299 = new _module.C1();
          _nw299.__ctor(_1609_i8, new BigNumber(((_1631_v119).f24).length), _1530_v60.f16, _1630_v118, _module.__default.safeModuloInt(new BigNumber((_1634_v120).length), new BigNumber((_1540_v70).cardinality())));
          _1635_v121 = _nw299;
        } else if (_source18.is_DC23) {
          let _1636_v122;
          _1636_v122 = _dafny.Map.Empty.slice().updateUnsafe(_1609_i8,_dafny.Seq.UnicodeFromString("jvrpwreug"));
          let _1637_v123;
          _1637_v123 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-794),_1530_v60.f16);
          let _1638_v125;
          _1638_v125 = _dafny.Seq.of(_1608_v113, function () {
            let _coll63 = new _dafny.Map();
            for (const _compr_63 of _dafny.IntegerRange(new BigNumber(918), new BigNumber(620))) {
              let _1639_v124 = _compr_63;
              if (((new BigNumber(918)).isLessThanOrEqualTo(_1639_v124)) && ((_1639_v124).isLessThan(new BigNumber(620)))) {
                _coll63.push([(_1639_v124).multipliedBy(_1609_i8),_1609_i8]);
              }
            }
            return _coll63;
          }());
          let _1640_v126;
          let _nw300 = new _module.C3();
          _nw300.__ctor(_1538_v68, _1530_v60.f16, (_1530_v60).f17);
          _1640_v126 = _nw300;
          let _1641_v127;
          _1641_v127 = _module.D12.create_DC35(false, _1530_v60.f16, (_1638_v125)[_module.__default.safeIndex(_1528_v58, new BigNumber((_1638_v125).length))], _1530_v60.f16, _1640_v126);
          _1636_v122 = (_1636_v122).update((_dafny.ZERO).minus(new BigNumber(((_1637_v123).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1528_v58,(_1641_v127).dtor_cf63))).length)), (_1640_v126).f24);
          let _1642_v128;
          let _out14;
          _out14 = (_1536_v67).m5(globalState);
          _1642_v128 = _out14;
          _1538_v68 = _module.__default.fm1(globalState);
          let _1643_v129;
          let _nw301 = Array((new BigNumber(11)).toNumber()).fill([]);
          _1643_v129 = _nw301;
          let _1644_v130;
          _1644_v130 = _module.D6.create_DC15(_1643_v129);
          let _1645_v131;
          let _init46 = ((_1646_v60) => function (_1647_i10) {
            return (_1646_v60.f16) || (_1646_v60.f16);
          })(_1530_v60);
          let _nw302 = Array((new BigNumber(24)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw302.length); _i0_46++) {
            _nw302[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1645_v131 = _nw302;
          let _index286 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1645_v131).length));
          (_1645_v131)[_index286] = _dafny.Seq.IsPrefixOf((_1640_v126).f24, _dafny.Seq.UnicodeFromString("uncqkk"));
          let _index287 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_1645_v131).length));
          (_1645_v131)[_index287] = _module.__default.fm0(_1530_v60.f16, true, globalState);
          let _index288 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1645_v131).length));
          let _index289 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_1645_v131).length));
          let _rhs313 = _1644_v130;
          let _rhs314 = _1530_v60.f16;
          let _rhs315 = !((((_dafny.ZERO).minus(_1528_v58)).isEqualTo(new BigNumber((_1642_v128).length))) && (_1454_v0));
          let _lhs247 = _1645_v131;
          let _lhs248 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_1645_v131).length));
          let _lhs249 = _1645_v131;
          let _lhs250 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_1645_v131).length));
          _1644_v130 = _rhs313;
          _lhs247[_lhs248] = _rhs314;
          _lhs249[_lhs250] = _rhs315;
        } else if (_source18.is_DC20) {
          let _1648___mcc_h21 = (_source18).cf34;
          let _1649_cf34 = _1648___mcc_h21;
          let _1650_v132;
          _1650_v132 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1609_i8).plus(_1609_i8)),!(!(((_1530_v60.f16) ? (_1530_v60.f16) : ((_1536_v67).fm8(_dafny.Seq.UnicodeFromString("pwy"), globalState))))));
          _1650_v132 = (_1650_v132).update(_1609_i8, true);
          r3 = (new BigNumber(794)).plus(_1609_i8);
          let _1651_v133;
          let _init47 = ((_1652_v60, _1653_v68) => function (_1654_i11) {
            return ((_1652_v60.f16) ? (_1653_v68) : ((_module.D0.create_DC0(_1653_v68)).dtor_cf0));
          })(_1530_v60, _1538_v68);
          let _nw303 = Array((new BigNumber(19)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw303.length); _i0_47++) {
            _nw303[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1651_v133 = _nw303;
          let _index290 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length));
          (_1651_v133)[_index290] = _dafny.Seq.Concat(_1538_v68, _1538_v68);
          let _1655_v134;
          _1655_v134 = _dafny.Set.fromElements((((_1650_v132).contains(new BigNumber(836))) ? ((_1650_v132).get(new BigNumber(836))) : (_1454_v0)), _1530_v60.f16, _1530_v60.f16);
          let _index291 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length));
          let _rhs316 = _1530_v60.f16;
          let _rhs317 = _1538_v68;
          let _rhs318 = _1655_v134;
          let _rhs319 = _1528_v58;
          let _lhs251 = globalState;
          let _lhs252 = _1651_v133;
          let _lhs253 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length));
          _lhs251.f5 = _rhs316;
          _lhs252[_lhs253] = _rhs317;
          _1655_v134 = _rhs318;
          r3 = _rhs319;
          let _1656_v135;
          let _init48 = ((_1657_i8) => function (_1658_i12) {
            return (_1658_i12).multipliedBy(_1657_i8);
          })(_1609_i8);
          let _nw304 = Array((new BigNumber(13)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw304.length); _i0_48++) {
            _nw304[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1656_v135 = _nw304;
          let _1659_v136;
          _1659_v136 = _dafny.Map.Empty.slice().updateUnsafe(true,_1609_i8);
          let _1660_v137;
          _1660_v137 = _module.D3.create_DC9(_module.__default.fm17(globalState));
          let _1661_v138;
          _1661_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1609_i8,_1527_v57);
          let _1662_v139;
          _1662_v139 = _dafny.Map.Empty.slice().updateUnsafe(_1660_v137,new BigNumber((_1661_v138).length));
          let _1663_v141;
          let _nw305 = Array((new BigNumber(25)).toNumber());
          _nw305[(_dafny.ZERO).toNumber()] = (_1530_v60).f17;
          _nw305[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(727)), ((_1664_i8) => function (_1665_i13) {
            return _1664_i8;
          })(_1609_i8));
          _nw305[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_1609_i8);
          _nw305[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(new BigNumber(((_1651_v133)[_module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length))]).length));
          _nw305[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), ((_1666_i8) => function (_1667_i14) {
            return _1666_i8;
          })(_1609_i8));
          _nw305[(new BigNumber(5)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(6)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(7)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(8)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(9)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(10)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(11)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(_1528_v58, _1528_v58, _1609_i8);
          _nw305[(new BigNumber(13)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(14)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_1609_i8, _1528_v58, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(320)), ((_1668_v57) => function (_1669_i15) {
            return _1668_v57;
          })(_1527_v57))).length));
          _nw305[(new BigNumber(16)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(17)).toNumber()] = _dafny.Seq.update((_1530_v60).f17, _module.__default.safeIndex(_1609_i8, new BigNumber(((_1530_v60).f17).length)), new BigNumber((_1662_v139).length));
          _nw305[(new BigNumber(18)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(new BigNumber((function () {
            let _coll64 = new _dafny.Set();
            for (const _compr_64 of _dafny.IntegerRange(new BigNumber(263), new BigNumber(597))) {
              let _1670_v140 = _compr_64;
              if (((new BigNumber(263)).isLessThanOrEqualTo(_1670_v140)) && ((_1670_v140).isLessThan(new BigNumber(597)))) {
                _coll64.add((_1670_v140).multipliedBy(_1609_i8));
              }
            }
            return _coll64;
          }()).length));
          _nw305[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_1609_i8);
          _nw305[(new BigNumber(21)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(22)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(23)).toNumber()] = (_1530_v60).f17;
          _nw305[(new BigNumber(24)).toNumber()] = (_1530_v60).f17;
          _1663_v141 = _nw305;
          let _1671_v142;
          _1671_v142 = _dafny.Seq.of(_1663_v141, _1663_v141, _1663_v141);
          let _1672_v143;
          _1672_v143 = _module.D2.create_DC8(_1659_v136, (_1671_v142)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1538_v68, _module.__default.safeIndex(_1528_v58, new BigNumber((_1538_v68).length)), _1527_v57)).length), new BigNumber((_1671_v142).length))], _1530_v60.f16, !(_1454_v0), _1609_i8);
          let _1673_v144;
          _1673_v144 = _dafny.Map.Empty.slice().updateUnsafe(_1530_v60.f16,_1530_v60.f16);
          let _1674_v145;
          _1674_v145 = _dafny.Map.Empty.slice().updateUnsafe((_1672_v143).dtor_cf16,(_1673_v144).Merge(_1673_v144));
          let _1675_v146;
          _1675_v146 = _dafny.Map.Empty.slice().updateUnsafe(_1538_v68,_1656_v135);
          let _index292 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length));
          let _rhs320 = _dafny.Seq.update(_dafny.Seq.Concat((_1651_v133)[_module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length))], _1538_v68), _module.__default.safeIndex(_1609_i8, new BigNumber((_dafny.Seq.Concat((_1651_v133)[_module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length))], _1538_v68)).length)), _1527_v57);
          let _rhs321 = (((_1675_v146).contains(_1538_v68)) ? ((_1675_v146).get(_1538_v68)) : (_1656_v135));
          let _rhs322 = (((_1674_v145).update(!(_1530_v60.f16), _dafny.Map.Empty.slice().updateUnsafe(_1530_v60.f16,_1530_v60.f16))).Merge(_1674_v145)).Merge(_1674_v145);
          let _lhs254 = _1651_v133;
          let _lhs255 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_1651_v133).length));
          _lhs254[_lhs255] = _rhs320;
          _1656_v135 = _rhs321;
          _1674_v145 = _rhs322;
        } else {
          let _1676___mcc_h22 = (_source18).cf43;
          let _1677_cf43 = _1676___mcc_h22;
          let _1678_v147;
          let _nw306 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
          _1678_v147 = _nw306;
          let _1679_v148;
          _1679_v148 = _dafny.Map.Empty.slice().updateUnsafe(_1454_v0,_1530_v60.f16);
          let _index293 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_1678_v147).length));
          (_1678_v147)[_index293] = _dafny.Map.Empty.slice().updateUnsafe(!((((_1679_v148).contains(_1530_v60.f16)) ? ((_1679_v148).get(_1530_v60.f16)) : (_1530_v60.f16))),_1528_v58);
          let _1680_v149;
          _1680_v149 = _dafny.Map.Empty.slice().updateUnsafe(_1530_v60.f16,new BigNumber((_dafny.Seq.UnicodeFromString("phnuuwd")).length));
          let _index294 = _module.__default.safeIndex(new BigNumber(15), new BigNumber((_1678_v147).length));
          (_1678_v147)[_index294] = (_1680_v149).Merge(_1680_v149);
          let _1681_v150;
          _1681_v150 = _dafny.Map.Empty.slice().updateUnsafe(_1536_v67,_1530_v60.f16);
          _1681_v150 = (_1681_v150).update(_1536_v67, _1530_v60.f16);
          let _1682_v151;
          let _init49 = ((_1683_v68) => function (_1684_i16) {
            return (_1684_i16).multipliedBy(new BigNumber((_1683_v68).length));
          })(_1538_v68);
          let _nw307 = Array((new BigNumber(6)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw307.length); _i0_49++) {
            _nw307[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1682_v151 = _nw307;
          let _index295 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1682_v151).length));
          (_1682_v151)[_index295] = _1528_v58;
          let _index296 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_1682_v151).length));
          (_1682_v151)[_index296] = _1609_i8;
          (_1530_v60).f16 = true;
        }
        let _1685_v152;
        _1685_v152 = _dafny.Set.fromElements(_1530_v60.f16, ((_1454_v0) ? (_1530_v60.f16) : (_1530_v60.f16)));
        _1685_v152 = (_1685_v152).Union(_1685_v152);
        _1528_v58 = new BigNumber(((((_1454_v0) || (!(_1454_v0))) ? ((_1530_v60).f17) : ((_1530_v60).f17))).length);
        _1528_v58 = (_1609_i8).plus(_1528_v58);
      }
      let _1686_v153;
      _1686_v153 = _dafny.Map.Empty.slice().updateUnsafe(_1530_v60.f16,_1528_v58);
      let _1687_v154;
      _1687_v154 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1686_v153).length),_1454_v0);
      let _1688_v155;
      _1688_v155 = _dafny.Seq.of(_1687_v154, _1687_v154);
      r0 = _dafny.Seq.IsProperPrefixOf(_1688_v155, _dafny.Seq.Concat(_1688_v155, _1688_v155));
      let _1689_v156;
      _1689_v156 = _dafny.MultiSet.fromElements(_1686_v153, _1686_v153, _dafny.Map.Empty.slice().updateUnsafe(_1530_v60.f16,_1528_v58));
      let _pat_let_tv32 = _1540_v70;
      let _pat_let_tv33 = _1540_v70;
      let _pat_let_tv34 = _1530_v60;
      let _pat_let_tv35 = _1530_v60;
      r1 = function (_source19) {
        if (_source19.is_DC12) {
          let _1690___mcc_h23 = (_source19).cf21;
          let _1691___mcc_h24 = (_source19).cf22;
          let _1692___mcc_h25 = (_source19).cf23;
          let _1693_cf23 = _1692___mcc_h25;
          let _1694_cf22 = _1691___mcc_h24;
          let _1695_cf21 = _1690___mcc_h23;
          return (_pat_let_tv32).IsSubsetOf(_pat_let_tv33);
        } else if (_source19.is_DC13) {
          let _1696___mcc_h26 = (_source19).cf24;
          let _1697___mcc_h27 = (_source19).cf25;
          let _1698___mcc_h28 = (_source19).cf26;
          let _1699_cf26 = _1698___mcc_h28;
          let _1700_cf25 = _1697___mcc_h27;
          let _1701_cf24 = _1696___mcc_h26;
          return _pat_let_tv34.f16;
        } else {
          let _1702___mcc_h29 = (_source19).cf20;
          let _1703_cf20 = _1702___mcc_h29;
          return _pat_let_tv35.f16;
        }
      }(_module.D4.create_DC11(_1689_v156));
      r2 = !(_1530_v60.f16);
      r3 = _module.__default.fm3(globalState);
      return [r0, r1, r2, r3];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f16 = false;
      this._f17 = _dafny.Seq.of();
      this._f20 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f20, f16, f17) {
      let _this = this;
      (_this)._f20 = f20;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let _1704_v0;
      _1704_v0 = new BigNumber(33);
      _1704_v0 = _module.__default.fm3(globalState);
      (globalState).f5 = !(_this.f16) || ((p3).isEqualTo(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_1704_v0))).cardinality())));
      let _1705_v1;
      let _nw308 = Array((new BigNumber(11)).toNumber()).fill(false);
      _1705_v1 = _nw308;
      let _1706_v2;
      _1706_v2 = _module.D0.create_DC2(_1704_v0, _1704_v0, _1704_v0, _1705_v1);
      let _hi15 = _module.__default.safeModuloInt(new BigNumber(929), (_1706_v2).dtor_cf3);
      for (let _1707_i0 = (p3).plus(_1704_v0); _1707_i0.isLessThan(_hi15); _1707_i0 = _1707_i0.plus(_dafny.ONE)) {
        let _1708_v3;
        _1708_v3 = _dafny.Seq.of(p0);
        let _1709_v4;
        let _nw309 = Array((new BigNumber(24)).toNumber());
        _nw309[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("w");
        _nw309[(_dafny.ONE).toNumber()] = p0;
        _nw309[(new BigNumber(2)).toNumber()] = p0;
        _nw309[(new BigNumber(3)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), function (_1710_i1) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        });
        _nw309[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-237)), function (_1711_i2) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        });
        _nw309[(new BigNumber(5)).toNumber()] = p0;
        _nw309[(new BigNumber(6)).toNumber()] = p0;
        _nw309[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("magnsasim");
        _nw309[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("edeabhm");
        _nw309[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("xiq");
        _nw309[(new BigNumber(10)).toNumber()] = p0;
        _nw309[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("mul");
        _nw309[(new BigNumber(12)).toNumber()] = p0;
        _nw309[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(p0, p0);
        _nw309[(new BigNumber(14)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(619)), function (_1712_i3) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        });
        _nw309[(new BigNumber(15)).toNumber()] = ((false) ? (_dafny.Seq.UnicodeFromString("wbhpqphi")) : (p0));
        _nw309[(new BigNumber(16)).toNumber()] = p0;
        _nw309[(new BigNumber(17)).toNumber()] = p0;
        _nw309[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(p0, _module.__default.fm1(globalState));
        _nw309[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(p0, p0);
        _nw309[(new BigNumber(20)).toNumber()] = p0;
        _nw309[(new BigNumber(21)).toNumber()] = (_1708_v3)[_module.__default.safeIndex(_1704_v0, new BigNumber((_1708_v3).length))];
        _nw309[(new BigNumber(22)).toNumber()] = p0;
        _nw309[(new BigNumber(23)).toNumber()] = p0;
        _1709_v4 = _nw309;
        _1709_v4 = _1709_v4;
        let _1713_v5;
        _1713_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f16);
        _1713_v5 = (_1713_v5).update(_dafny.Seq.Concat(p0, p0), _this.f16);
        let _1714_v6;
        let _nw310 = new _module.C3();
        _nw310.__ctor(p0, p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), ((_1715_i0) => function (_1716_i4) {
          return _1715_i0;
        })(_1707_i0)));
        _1714_v6 = _nw310;
        let _1717_v7;
        _1717_v7 = _module.D11.create_DC30(p2, (_dafny.ZERO).minus(_1704_v0));
        _1704_v0 = ((_1717_v7).dtor_cf52).minus(p3);
      }
      if (p1) {
        let _1718_v8;
        _1718_v8 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1719_v9;
        _1719_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1718_v8,p3);
        _1704_v0 = new BigNumber((_1719_v9).length);
        _1705_v1 = _1705_v1;
        let _1720_v10;
        _1720_v10 = _dafny.Seq.of(_1704_v0);
        let _index297 = _module.__default.safeIndex(new BigNumber(109), new BigNumber(((_this).f20).length));
        ((_this).f20)[_index297] = new BigNumber(252);
        let _index298 = _module.__default.safeIndex(new BigNumber(109), new BigNumber(((_this).f20).length));
        let _rhs323 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f17, (_this).f17), _dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_1721_p0) => function (_1722_i5) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_this.f16)).length)),_dafny.Seq.of(new BigNumber((_1721_p0).length)))).length);
        })(p0)));
        let _rhs324 = _module.__default.fm3(globalState);
        let _rhs325 = p3;
        let _lhs256 = (_this).f20;
        let _lhs257 = _module.__default.safeIndex(new BigNumber(109), new BigNumber(((_this).f20).length));
        _1720_v10 = _rhs323;
        _1704_v0 = _rhs324;
        _lhs256[_lhs257] = _rhs325;
        let _1723_v11;
        _1723_v11 = _dafny.Set.fromElements(_module.__default.fm0(p2, p1, globalState));
        let _1724_v12;
        _1724_v12 = _dafny.Set.fromElements(((_this).f20)[_module.__default.safeIndex(new BigNumber(109), new BigNumber(((_this).f20).length))], p3, p3, _module.__default.fm3(globalState));
        let _1725_v13;
        _1725_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1723_v11,_1724_v12);
        _1725_v13 = (_1725_v13).update(_1723_v11, _1724_v12);
        let _1726_v14;
        let _init50 = function (_1727_i6) {
          return (_1727_i6).multipliedBy(((_this).f20)[_module.__default.safeIndex(new BigNumber(109), new BigNumber(((_this).f20).length))]);
        };
        let _nw311 = Array((new BigNumber(26)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw311.length); _i0_50++) {
          _nw311[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1726_v14 = _nw311;
        _1726_v14 = _1726_v14;
      } else {
        let _1728_v15;
        let _nw312 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _1728_v15 = _nw312;
        let _1729_v16;
        _1729_v16 = new _dafny.CodePoint('q'.codePointAt(0));
        let _1730_v17;
        _1730_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1729_v16,p3);
        let _1731_v18;
        _1731_v18 = _dafny.Set.fromElements(p3);
        let _nw313 = Array((new BigNumber(12)).toNumber());
        _nw313[(_dafny.ZERO).toNumber()] = _1704_v0;
        _nw313[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update(((p2) ? (_dafny.Seq.UnicodeFromString("eio")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), function (_1732_i7) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }))), _module.__default.safeIndex(_1704_v0, new BigNumber((((p2) ? (_dafny.Seq.UnicodeFromString("eio")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), function (_1733_i7) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        })))).length)), _1729_v16)).length);
        _nw313[(new BigNumber(2)).toNumber()] = (((_1730_v17).contains(_1729_v16)) ? ((_1730_v17).get(_1729_v16)) : (_1704_v0));
        _nw313[(new BigNumber(3)).toNumber()] = _1704_v0;
        _nw313[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(p3, new BigNumber((p0).length));
        _nw313[(new BigNumber(5)).toNumber()] = ((p2) ? (_1704_v0) : (p3));
        _nw313[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p3));
        _nw313[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(_1704_v0, _1704_v0);
        _nw313[(new BigNumber(8)).toNumber()] = p3;
        _nw313[(new BigNumber(9)).toNumber()] = new BigNumber((_1731_v18).length);
        _nw313[(new BigNumber(10)).toNumber()] = p3;
        _nw313[(new BigNumber(11)).toNumber()] = ((p2) ? (new BigNumber(956)) : (new BigNumber((_1731_v18).length)));
        _1728_v15 = _nw313;
        let _1734_v19;
        _1734_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("hwadxtdi"),p1);
        let _1735_v20;
        _1735_v20 = _dafny.Seq.of(_this.f16);
        let _1736_v21;
        _1736_v21 = _module.D4.create_DC12(p2, _module.__default.fm0(p1, p2, globalState), new BigNumber((_1735_v20).length));
        _1734_v19 = (_1734_v19).update(_dafny.Seq.Concat(p0, p0), (_1736_v21).dtor_cf21);
        let _1737_v22;
        _1737_v22 = _dafny.Seq.UnicodeFromString("gmrqcudhj");
        _1737_v22 = _1737_v22;
        let _1738_v23;
        let _nw314 = new _module.C1();
        _nw314.__ctor(p3, _module.__default.safeDivisionInt(p3, new BigNumber((_1731_v18).length)), p2, (_this).f17, _1704_v0);
        _1738_v23 = _nw314;
        let _1739_v24;
        _1739_v24 = _dafny.Seq.of(_1704_v0);
        let _1740_v25;
        _1740_v25 = _dafny.Map.Empty.slice().updateUnsafe((_1738_v23).f29,!(_this.f16));
        let _1741_v26;
        let _nw315 = new _module.C0();
        _nw315.__ctor((_1738_v23).f29, _1705_v1);
        _1741_v26 = _nw315;
        let _1742_v27;
        _1742_v27 = _dafny.MultiSet.fromElements(p1);
        let _rhs326 = _dafny.Seq.update(_module.__default.fm17(globalState), _module.__default.safeIndex((_module.D9.create_DC26(new BigNumber((_1740_v25).length), _1741_v26)).dtor_cf45, new BigNumber((_module.__default.fm17(globalState)).length)), (_1704_v0).minus((_1738_v23).f30));
        let _rhs327 = (_dafny.ZERO).minus(((_1741_v26).f26).minus(_1704_v0));
        let _rhs328 = _module.__default.fm0(((p2) ? (!(_this.f16)) : (p1)), (_1731_v18).IsProperSubsetOf(_1731_v18), globalState);
        let _rhs329 = (_1742_v27).IsSubsetOf(_1742_v27);
        let _rhs330 = _1734_v19;
        let _lhs258 = globalState;
        let _lhs259 = globalState;
        _1739_v24 = _rhs326;
        _1704_v0 = _rhs327;
        _lhs258.f5 = _rhs328;
        _lhs259.f2 = _rhs329;
        _1734_v19 = _rhs330;
      }
      let _1743_v28;
      _1743_v28 = _module.D0.create_DC0(_dafny.Seq.UnicodeFromString("egdqwsb"));
      if (_dafny.Seq.IsProperPrefixOf((_1743_v28).dtor_cf0, p0)) {
        let _1744_v29;
        _1744_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1705_v1,_this.f16);
        let _1745_v30;
        _1745_v30 = _1744_v29;
        let _1746_v31;
        let _nw316 = Array((new BigNumber(21)).toNumber());
        _nw316[(_dafny.ZERO).toNumber()] = _1744_v29;
        _nw316[(_dafny.ONE).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1705_v1,p1);
        _nw316[(new BigNumber(3)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(4)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(5)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(6)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(7)).toNumber()] = (_1744_v29).Merge(_1744_v29);
        _nw316[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1705_v1,p1);
        _nw316[(new BigNumber(9)).toNumber()] = (_1744_v29).Merge(_1744_v29);
        _nw316[(new BigNumber(10)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(11)).toNumber()] = (_1745_v30);
        _nw316[(new BigNumber(12)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(13)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(14)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(15)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(16)).toNumber()] = (_1744_v29).update(_1705_v1, p2);
        _nw316[(new BigNumber(17)).toNumber()] = _1744_v29;
        _nw316[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1705_v1,!(_this.f16));
        _nw316[(new BigNumber(19)).toNumber()] = (_1744_v29).update(_1705_v1, p2);
        _nw316[(new BigNumber(20)).toNumber()] = _1744_v29;
        _1746_v31 = _nw316;
        let _index299 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_1746_v31).length));
        (_1746_v31)[_index299] = (_1744_v29).Merge(_1744_v29);
        let _index300 = _module.__default.safeIndex(new BigNumber(632), new BigNumber((_1746_v31).length));
        (_1746_v31)[_index300] = _1744_v29;
        let _1747_v32;
        _1747_v32 = _dafny.Set.fromElements((_this).f20, (_this).f20);
        let _1748_v33;
        let _nw317 = new _module.C5();
        _nw317.__ctor(_this.f16, _dafny.Seq.update((_this).f17, _module.__default.safeIndex(new BigNumber((_1747_v32).length), new BigNumber(((_this).f17).length)), p3));
        _1748_v33 = _nw317;
        let _index301 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
        (_1705_v1)[_index301] = p2;
        let _index302 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
        (_1705_v1)[_index302] = _this.f16;
        let _1749_v34;
        _1749_v34 = _dafny.Map.Empty.slice().updateUnsafe(true,_1704_v0);
        let _1750_v35;
        let _init51 = function (_1751_i8) {
          return (_this).f17;
        };
        let _nw318 = Array((new BigNumber(9)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw318.length); _i0_51++) {
          _nw318[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _1750_v35 = _nw318;
        let _1752_v36;
        _1752_v36 = _module.D2.create_DC8(_1749_v34, _1750_v35, !(_this.f16), p1, p3);
        let _source20 = _module.D6.create_DC17(_1752_v36);
        if (_source20.is_DC16) {
          let _1753___mcc_h0 = (_source20).cf29;
          let _1754___mcc_h1 = (_source20).cf30;
          let _1755_cf30 = _1754___mcc_h1;
          let _1756_cf29 = _1753___mcc_h0;
          let _1757_v37;
          _1757_v37 = _dafny.Seq.UnicodeFromString("lutqh");
          let _index303 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
          let _rhs331 = _1755_cf30;
          let _rhs332 = (_1704_v0).plus(_1704_v0);
          let _rhs333 = (new BigNumber((_dafny.Seq.Concat(p0, _1757_v37)).length)).isLessThanOrEqualTo(_1704_v0);
          let _rhs334 = _dafny.Seq.Concat((_1748_v33).fm7(_dafny.Seq.of(_1756_cf29, _1756_cf29), globalState), _dafny.Seq.UnicodeFromString("soa"));
          let _lhs260 = globalState;
          let _lhs261 = _1705_v1;
          let _lhs262 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
          _lhs260.f2 = _rhs331;
          _1704_v0 = _rhs332;
          _lhs261[_lhs262] = _rhs333;
          _1757_v37 = _rhs334;
          let _index304 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
          (_1705_v1)[_index304] = p1;
          let _1758_v38;
          _1758_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(globalState),_dafny.Seq.of((_this).f20));
          let _1759_v39;
          _1759_v39 = _dafny.Seq.of((_this).f20);
          _1758_v38 = (_1758_v38).update(p3, _dafny.Seq.Concat(_1759_v39, _1759_v39));
          _1755_cf30 = p1;
        } else if (_source20.is_DC17) {
          let _1760___mcc_h2 = (_source20).cf31;
          let _1761_cf31 = _1760___mcc_h2;
          _1704_v0 = ((new BigNumber(677)).minus((_dafny.ZERO).minus(p3))).minus((_1704_v0).plus(p3));
          let _1762_v40;
          _1762_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(_this).f20);
          let _1763_v41;
          _1763_v41 = _dafny.Seq.of((_this).f20, (_this).f20, (_this).f20, (_this).f20, (_this).f20);
          let _1764_v42;
          _1764_v42 = _dafny.MultiSet.fromElements((_this).f20, (_this).f20, (_this).f20, (((_1762_v40).contains((_1705_v1)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length))])) ? ((_1762_v40).get((_1705_v1)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length))])) : ((_this).f20)), (_1763_v41)[_module.__default.safeIndex(p3, new BigNumber((_1763_v41).length))]);
          _1764_v42 = ((_1764_v42).Union(_1764_v42)).Union(_1764_v42);
          let _1765_v43;
          _1765_v43 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          _1762_v40 = (_1762_v40).update(!((((_1765_v43).contains(p0)) ? ((_1765_v43).get(p0)) : (p2))), (_this).f20);
          let _1766_v44;
          let _nw319 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1766_v44 = _nw319;
          let _1767_v45;
          let _nw320 = new _module.C0();
          _nw320.__ctor(p3, _1766_v44);
          _1767_v45 = _nw320;
          let _index305 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
          let _rhs335 = (true) && (p2);
          let _rhs336 = _1767_v45;
          let _lhs263 = _1705_v1;
          let _lhs264 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
          _lhs263[_lhs264] = _rhs335;
          _1767_v45 = _rhs336;
        } else if (_source20.is_DC18) {
          let _1768___mcc_h3 = (_source20).cf32;
          let _1769_cf32 = _1768___mcc_h3;
          let _index306 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1769_cf32).length));
          (_1769_cf32)[_index306] = p3;
          let _index307 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1769_cf32).length));
          (_1769_cf32)[_index307] = p3;
          (globalState).f2 = p1;
          let _1770_v46;
          let _nw321 = new _module.C4();
          _nw321.__ctor(new BigNumber(235), _this.f16);
          _1770_v46 = _nw321;
          (globalState).f5 = (_1705_v1)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length))];
        } else {
          let _1771___mcc_h4 = (_source20).cf28;
          let _1772_cf28 = _1771___mcc_h4;
          let _index308 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length));
          (_1705_v1)[_index308] = (_1705_v1)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length))];
          let _1773_v47;
          _1773_v47 = _module.D0.create_DC1(p0);
          (globalState).f5 = _dafny.Seq.IsPrefixOf(p0, ((p2) ? ((_1773_v47).dtor_cf1) : (p0)));
          let _1774_v48;
          _1774_v48 = _dafny.MultiSet.fromElements(_1704_v0, p3);
          let _1775_v49;
          _1775_v49 = _dafny.MultiSet.fromElements((_1774_v48).Intersect(_1774_v48), _module.__default.fm19(_1704_v0, globalState), _dafny.MultiSet.FromArray(_module.__default.fm17(globalState)));
          _1775_v49 = _module.__default.fm39((_1705_v1)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(498)), ((_1776_p3) => function (_1777_i9) {
            return _1776_p3;
          })(p3)), (p0)[_module.__default.safeIndex(p3, new BigNumber((p0).length))], globalState);
          (_this).m2(globalState);
        }
        let _1778_v50;
        _1778_v50 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
        let _rhs337 = !(!(!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iuu"), p0), p0))));
        let _rhs338 = (((_1778_v50).contains(_1704_v0)) ? ((_1778_v50).get(_1704_v0)) : ((((_1705_v1)[_module.__default.safeIndex(new BigNumber(489), new BigNumber((_1705_v1).length))]) ? (_1704_v0) : (_1704_v0))));
        let _rhs339 = p3;
        let _lhs265 = globalState;
        _lhs265.f5 = _rhs337;
        _1704_v0 = _rhs338;
        _1704_v0 = _rhs339;
      } else {
        let _1779_v51;
        _1779_v51 = new _dafny.CodePoint('q'.codePointAt(0));
        if ((p3).isLessThanOrEqualTo((_1704_v0).multipliedBy((_module.D12.create_DC34(_1705_v1, p3, (_this).f20, _1779_v51)).dtor_cf60))) {
          let _1780_v52;
          _1780_v52 = _module.D11.create_DC31(p1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(721)), ((_1781_v51) => function (_1782_i10) {
  return _1781_v51;
})(_1779_v51)), p0)).length), _1704_v0, p2);
          _1780_v52 = _1780_v52;
          let _1783_v53;
          _1783_v53 = _dafny.MultiSet.fromElements(_1704_v0);
          _module.__default.m0((((_1783_v53).contains(p3)) ? ((_1783_v53).get(p3)) : ((_dafny.ZERO).minus(_module.__default.fm3(globalState)))), _dafny.Seq.UnicodeFromString("re"), ((_this).f17)[_module.__default.safeIndex(_module.__default.fm3(globalState), new BigNumber(((_this).f17).length))], _1704_v0, globalState);
          (globalState).f2 = _this.f16;
          let _1784_v54;
          _1784_v54 = _dafny.Seq.UnicodeFromString("ykvixfv");
          _1784_v54 = _dafny.Seq.Concat(_1784_v54, _dafny.Seq.UnicodeFromString("ir"));
          let _1785_v55;
          _1785_v55 = _dafny.Set.fromElements(p2);
          let _index309 = _module.__default.safeIndex(new BigNumber(455), new BigNumber(((_this).f20).length));
          ((_this).f20)[_index309] = new BigNumber(((_1785_v55).Intersect(_1785_v55)).length);
          let _1786_v56;
          _1786_v56 = _dafny.Seq.of((_this).f17, _module.__default.fm17(globalState), (_this).f17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-271)), ((_1787_v0) => function (_1788_i11) {
            return (_dafny.ZERO).minus(_1787_v0);
          })(_1704_v0)), (_this).f17);
          let _index310 = _module.__default.safeIndex(new BigNumber(455), new BigNumber(((_this).f20).length));
          let _rhs340 = new BigNumber((p0).length);
          let _rhs341 = _1786_v56;
          let _rhs342 = _1784_v54;
          let _lhs266 = (_this).f20;
          let _lhs267 = _module.__default.safeIndex(new BigNumber(455), new BigNumber(((_this).f20).length));
          _lhs266[_lhs267] = _rhs340;
          _1786_v56 = _rhs341;
          _1784_v54 = _rhs342;
        } else {
          (globalState).f5 = !(p1);
          (globalState).f5 = _dafny.areEqual(_dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(914)), ((_1789_v51) => function (_1790_i12) {
            return _1789_v51;
          })(_1779_v51))), p0);
          _1704_v0 = p3;
          let _1791_v57;
          let _nw322 = Array((new BigNumber(3)).toNumber());
          _nw322[(_dafny.ZERO).toNumber()] = _1705_v1;
          _nw322[(_dafny.ONE).toNumber()] = ((p1) ? (_1705_v1) : (_1705_v1));
          _nw322[(new BigNumber(2)).toNumber()] = _1705_v1;
          _1791_v57 = _nw322;
          let _index311 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1791_v57).length));
          (_1791_v57)[_index311] = _1705_v1;
          let _index312 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_1791_v57).length));
          (_1791_v57)[_index312] = _1705_v1;
          _1704_v0 = _1704_v0;
        }
        if ((_1704_v0).isLessThan(_1704_v0)) {
          let _1792_v58;
          _1792_v58 = _module.D2.create_DC7(_dafny.Set.fromElements(_this.f16, _this.f16, p1));
          let _1793_v59;
          _1793_v59 = _dafny.Seq.of(_1792_v58);
          let _1794_v60;
          _1794_v60 = _dafny.Seq.of(_1793_v59, _1793_v59);
          let _1795_v61;
          _1795_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray((_1794_v60)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_1794_v60).length))]),p3);
          let _1796_v62;
          _1796_v62 = _dafny.Set.fromElements(p2, _this.f16);
          let _1797_v63;
          _1797_v63 = _dafny.MultiSet.fromElements(_module.D2.create_DC7(_1796_v62));
          _1795_v61 = (_1795_v61).update(_1797_v63, _1704_v0);
          let _1798_v64;
          _1798_v64 = _dafny.Seq.of((_this).f17);
          let _1799_v65;
          let _nw323 = new _module.C1();
          _nw323.__ctor((_dafny.ZERO).minus(p3), _1704_v0, p1, (_1798_v64)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_1798_v64).length))], new BigNumber(154));
          _1799_v65 = _nw323;
          _1799_v65 = _1799_v65;
          (_this).m2(globalState);
          (globalState).f2 = _1799_v65.f16;
          let _1800_v66;
          _1800_v66 = _dafny.Seq.of((_1704_v0).plus(p3));
          let _1801_v68;
          _1801_v68 = _dafny.MultiSet.fromElements(_1779_v51);
          let _rhs343 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1800_v66, _1800_v66), (_1799_v65).f17);
          let _rhs344 = (new BigNumber((function () {
            let _coll65 = new _dafny.Map();
            for (const _compr_65 of (_1801_v68).Elements) {
              let _1802_v67 = _compr_65;
              if ((_1801_v68).contains(_1802_v67)) {
                _coll65.push([_1802_v67,new BigNumber((_1800_v66).length)]);
              }
            }
            return _coll65;
          }()).length)).plus(_module.__default.safeModuloInt(_1704_v0, new BigNumber((p0).length)));
          let _rhs345 = ((false) || (true)) === (_this.f16);
          let _rhs346 = p3;
          let _rhs347 = true;
          let _lhs268 = globalState;
          let _lhs269 = globalState;
          _1800_v66 = _rhs343;
          _1704_v0 = _rhs344;
          _lhs268.f5 = _rhs345;
          _1704_v0 = _rhs346;
          _lhs269.f5 = _rhs347;
        } else {
          let _1803_v69;
          _1803_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1804_v70;
          _1804_v70 = _module.D1.create_DC4(!(_this.f16), _1704_v0, _1803_v69);
          let _1805_v71;
          _1805_v71 = _dafny.Seq.of((_1804_v70).dtor_cf9, _1803_v69, _1803_v69);
          let _index313 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length));
          (_1705_v1)[_index313] = !_dafny.areEqual(_1805_v71, _1805_v71);
          let _index314 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length));
          (_1705_v1)[_index314] = p2;
          let _1806_v72;
          _1806_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1704_v0,p3);
          let _1807_v73;
          let _nw324 = new _module.C3();
          _nw324.__ctor(_dafny.Seq.UnicodeFromString("t"), p1, _dafny.Seq.of(new BigNumber(((_this).f17).length)));
          _1807_v73 = _nw324;
          let _1808_v74;
          _1808_v74 = _module.D12.create_DC35(p1, (p3).isLessThan((_dafny.ZERO).minus(_1704_v0)), _1806_v72, (_1705_v1)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length))], _1807_v73);
          let _index315 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length));
          let _index316 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length));
          let _rhs348 = p1;
          let _rhs349 = (_1705_v1)[_module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length))];
          let _rhs350 = _this.f16;
          let _rhs351 = _1808_v74;
          let _lhs270 = globalState;
          let _lhs271 = _1705_v1;
          let _lhs272 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length));
          let _lhs273 = _1705_v1;
          let _lhs274 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_1705_v1).length));
          _lhs270.f5 = _rhs348;
          _lhs271[_lhs272] = _rhs349;
          _lhs273[_lhs274] = _rhs350;
          _1808_v74 = _rhs351;
          let _1809_v75;
          _1809_v75 = _dafny.Set.fromElements((_dafny.ZERO).minus(_1704_v0), new BigNumber(406));
          let _rhs352 = p2;
          let _rhs353 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((p3).multipliedBy(new BigNumber((_1809_v75).length)), (_1704_v0).multipliedBy(p3)));
          let _lhs275 = _this;
          _lhs275.f16 = _rhs352;
          _1704_v0 = _rhs353;
          let _1810_v76;
          let _1811_v77;
          let _1812_v78;
          let _out15;
          let _out16;
          let _out17;
          let _outcollector4 = (_1807_v73).m8(globalState);
          _out15 = _outcollector4[0];
          _out16 = _outcollector4[1];
          _out17 = _outcollector4[2];
          _1810_v76 = _out15;
          _1811_v77 = _out16;
          _1812_v78 = _out17;
          let _1813_v79;
          let _nw325 = new _module.C2();
          _nw325.__ctor(_1704_v0);
          _1813_v79 = _nw325;
        }
        let _1814_v80;
        _1814_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1705_v1,_module.__default.fm0(p2, _this.f16, globalState));
        let _1815_v81;
        _1815_v81 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1814_v80,false));
        let _1816_v82;
        let _nw326 = new _module.C1();
        _nw326.__ctor(new BigNumber((_1815_v81).length), _1704_v0, p2, (_this).f17, _1704_v0);
        _1816_v82 = _nw326;
        let _1817_v83;
        _1817_v83 = _dafny.Seq.of(_1816_v82, _1816_v82);
        let _1818_v84;
        _1818_v84 = _dafny.Seq.of((_1817_v83)[_module.__default.safeIndex((_1816_v82).f30, new BigNumber((_1817_v83).length))], _1816_v82, _1816_v82);
        let _1819_v85;
        _1819_v85 = _dafny.MultiSet.fromElements((_1816_v82).f30, new BigNumber(128));
        let _1820_v86;
        let _nw327 = Array((new BigNumber(14)).toNumber());
        _nw327[(_dafny.ZERO).toNumber()] = _1816_v82;
        _nw327[(_dafny.ONE).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(2)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(3)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(4)).toNumber()] = (_1818_v84)[_module.__default.safeIndex(new BigNumber((_1819_v85).cardinality()), new BigNumber((_1818_v84).length))];
        _nw327[(new BigNumber(5)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(6)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(7)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(8)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(9)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(10)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(11)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(12)).toNumber()] = _1816_v82;
        _nw327[(new BigNumber(13)).toNumber()] = (_1818_v84)[_module.__default.safeIndex(new BigNumber(943), new BigNumber((_1818_v84).length))];
        _1820_v86 = _nw327;
        let _index317 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1820_v86).length));
        (_1820_v86)[_index317] = _1816_v82;
        let _index318 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_1820_v86).length));
        (_1820_v86)[_index318] = _1816_v82;
        let _1821_v87;
        let _nw328 = new _module.C1();
        _nw328.__ctor(_module.__default.safeModuloInt(new BigNumber(((_this).f17).length), new BigNumber(-478)), (_dafny.ZERO).minus(((_this).f17)[_module.__default.safeIndex(_1704_v0, new BigNumber(((_this).f17).length))]), false, _dafny.Seq.update((_this).f17, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("jstytmew")).length), new BigNumber(((_this).f17).length)), (_1816_v82).f30), (_1816_v82).f29);
        _1821_v87 = _nw328;
        _1821_v87 = _1821_v87;
        let _1822_v88;
        let _nw329 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1822_v88 = _nw329;
        let _index319 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_1822_v88).length));
        (_1822_v88)[_index319] = _1779_v51;
        let _index320 = _module.__default.safeIndex(new BigNumber(508), new BigNumber((_1822_v88).length));
        (_1822_v88)[_index320] = _1779_v51;
      }
      let _hi16 = new BigNumber(((_this).f17).length);
      for (let _1823_i13 = _1704_v0; _1823_i13.isLessThan(_hi16); _1823_i13 = _1823_i13.plus(_dafny.ONE)) {
        let _1824_v89;
        let _init52 = ((_1825_v0) => function (_1826_i14) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1825_v0,(_dafny.ZERO).minus(new BigNumber(-813)));
        })(_1704_v0);
        let _nw330 = Array((new BigNumber(6)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw330.length); _i0_52++) {
          _nw330[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1824_v89 = _nw330;
        let _1827_v90;
        _1827_v90 = _dafny.Set.fromElements(new BigNumber(-759));
        let _1828_v91;
        _1828_v91 = _dafny.Set.fromElements(p3, new BigNumber((_1827_v90).length), p3);
        let _1829_v92;
        _1829_v92 = _dafny.MultiSet.fromElements(_1704_v0);
        let _1830_v93;
        _1830_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1704_v0,_1829_v92);
        let _rhs354 = !(_1827_v90).contains(new BigNumber((p0).length));
        let _rhs355 = _module.__default.safeDivisionInt(new BigNumber((_1830_v93).length), _1823_i13);
        let _rhs356 = _1824_v89;
        let _rhs357 = p1;
        let _lhs276 = _this;
        let _lhs277 = globalState;
        _lhs276.f16 = _rhs354;
        _1704_v0 = _rhs355;
        _1824_v89 = _rhs356;
        _lhs277.f2 = _rhs357;
        (globalState).f5 = _this.f16;
        let _1831_v94;
        _1831_v94 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,p1);
        let _1832_v95;
        _1832_v95 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),new BigNumber((_1831_v94).length));
        _1704_v0 = (((_1832_v95).contains(_module.__default.fm3(globalState))) ? ((_1832_v95).get(_module.__default.fm3(globalState))) : (p3));
        if (false) {
          let _1833_v96;
          _1833_v96 = _module.D1.create_DC3(_dafny.MultiSet.fromElements(p1));
          let _1834_v97;
          _1834_v97 = _dafny.Map.Empty.slice().updateUnsafe((_1704_v0).minus(new BigNumber((p0).length)),_1833_v96);
          _1834_v97 = (_1834_v97).update(_1823_i13, _1833_v96);
          let _1835_v98;
          _1835_v98 = _dafny.MultiSet.fromElements(_1705_v1, _1705_v1, _1705_v1, _1705_v1, _1705_v1);
          let _rhs358 = !((((_1831_v94).contains(((p2) ? (_this.f16) : (p2)))) ? ((_1831_v94).get(((p2) ? (_this.f16) : (p2)))) : (_this.f16)));
          let _rhs359 = _1835_v98;
          let _lhs278 = globalState;
          _lhs278.f5 = _rhs358;
          _1835_v98 = _rhs359;
          let _index321 = _module.__default.safeIndex(new BigNumber(903), new BigNumber(((_this).f20).length));
          ((_this).f20)[_index321] = _module.__default.safeModuloInt(_1704_v0, _1704_v0);
          let _1836_v99;
          _1836_v99 = new _dafny.CodePoint('g'.codePointAt(0));
          let _index322 = _module.__default.safeIndex(new BigNumber(903), new BigNumber(((_this).f20).length));
          ((_this).f20)[_index322] = (((_this).f17)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_1836_v99, _1836_v99, _1836_v99)).cardinality()), new BigNumber(((_this).f17).length))]).minus(p3);
          let _index323 = _module.__default.safeIndex(new BigNumber(903), new BigNumber(((_this).f20).length));
          ((_this).f20)[_index323] = _1823_i13;
          let _index324 = _module.__default.safeIndex(new BigNumber(903), new BigNumber(((_this).f20).length));
          ((_this).f20)[_index324] = ((_this).f20)[_module.__default.safeIndex(new BigNumber(903), new BigNumber(((_this).f20).length))];
        } else {
          (globalState).f5 = p2;
          let _1837_v100;
          _1837_v100 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1704_v0);
          _1837_v100 = _1837_v100;
          let _pat_let_tv36 = _1704_v0;
          let _pat_let_tv37 = _1705_v1;
          let _pat_let_tv38 = _1705_v1;
          let _1838_v101;
          let _nw331 = new _module.C6();
          _nw331.__ctor(function (_pat_let44_0) {
            return function (_1843_dt__update__tmp_h1) {
              return function (_pat_let49_0) {
                return function (_1844_dt__update_hcf5_h1) {
                  return function (_pat_let50_0) {
                    return function (_1845_dt__update_hcf3_h1) {
                      return _module.D0.create_DC2((_1843_dt__update__tmp_h1).dtor_cf2, _1845_dt__update_hcf3_h1, (_1843_dt__update__tmp_h1).dtor_cf4, _1844_dt__update_hcf5_h1);
                    }(_pat_let50_0);
                  }(_1823_i13);
                }(_pat_let49_0);
              }(_pat_let_tv38);
            }(_pat_let44_0);
          }(function (_pat_let45_0) {
            return function (_1839_dt__update__tmp_h0) {
              return function (_pat_let46_0) {
                return function (_1840_dt__update_hcf3_h0) {
                  return function (_pat_let47_0) {
                    return function (_1841_dt__update_hcf2_h0) {
                      return function (_pat_let48_0) {
                        return function (_1842_dt__update_hcf5_h0) {
                          return _module.D0.create_DC2(_1841_dt__update_hcf2_h0, _1840_dt__update_hcf3_h0, (_1839_dt__update__tmp_h0).dtor_cf4, _1842_dt__update_hcf5_h0);
                        }(_pat_let48_0);
                      }(_pat_let_tv37);
                    }(_pat_let47_0);
                  }((_dafny.ZERO).minus(new BigNumber(-269)));
                }(_pat_let46_0);
              }(_pat_let_tv36);
            }(_pat_let45_0);
          }(_1706_v2)));
          _1838_v101 = _nw331;
          let _1846_v102;
          _1846_v102 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (new BigNumber(-131)) : (p3)),((false) ? (_1838_v101) : (_1838_v101)));
          let _1847_v103;
          _1847_v103 = _dafny.Seq.of(_1838_v101);
          _1838_v101 = (((_1846_v102).contains(new BigNumber(153))) ? ((_1846_v102).get(new BigNumber(153))) : ((_1847_v103)[_module.__default.safeIndex((((_1832_v95).contains(p3)) ? ((_1832_v95).get(p3)) : ((_dafny.ZERO).minus(new BigNumber((_1837_v100).length)))), new BigNumber((_1847_v103).length))]));
          let _1848_v104;
          _1848_v104 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1849_v105;
          _1849_v105 = _dafny.MultiSet.fromElements(!(!_dafny.Seq.contains(p0, _1848_v104)));
          _1849_v105 = (_1849_v105).Intersect((_1849_v105).update(_this.f16, _module.__default.abs(_module.__default.fm3(globalState))));
          let _1850_v106;
          _1850_v106 = _dafny.Seq.of(p2);
          let _1851_v107;
          _1851_v107 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_dafny.Seq.Concat(_1850_v106, _1850_v106));
          _1851_v107 = (_1851_v107).update(_dafny.areEqual(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(77)), ((_1852_v104) => function (_1853_i15) {
            return _1852_v104;
          })(_1848_v104))), _1850_v106);
        }
      }
      let _1854_v108;
      _1854_v108 = _dafny.Seq.of(false);
      r0 = _1854_v108;
      return r0;
    }
    m2(globalState) {
      let _this = this;
      let _1855_v0;
      _1855_v0 = new BigNumber(411);
      let _1856_v1;
      _1856_v1 = _dafny.Seq.of(_this.f16);
      let _1857_v2;
      _1857_v2 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_1856_v1).length)), _1855_v0, _1855_v0, _1855_v0, _1855_v0);
      let _1858_v3;
      let _nw332 = Array((new BigNumber(25)).toNumber());
      _nw332[(_dafny.ZERO).toNumber()] = _this.f16;
      _nw332[(_dafny.ONE).toNumber()] = _this.f16;
      _nw332[(new BigNumber(2)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(3)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(4)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(5)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(6)).toNumber()] = false;
      _nw332[(new BigNumber(7)).toNumber()] = (_1857_v2).IsSubsetOf(_1857_v2);
      _nw332[(new BigNumber(8)).toNumber()] = (_1855_v0).isLessThan(_1855_v0);
      _nw332[(new BigNumber(9)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(10)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(11)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(12)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(13)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(14)).toNumber()] = _module.__default.fm0(!(_this.f16), _this.f16, globalState);
      _nw332[(new BigNumber(15)).toNumber()] = (_1857_v2).IsSubsetOf(_1857_v2);
      _nw332[(new BigNumber(16)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(17)).toNumber()] = (_this.f16) === (_this.f16);
      _nw332[(new BigNumber(18)).toNumber()] = (_1856_v1)[_module.__default.safeIndex(_1855_v0, new BigNumber((_1856_v1).length))];
      _nw332[(new BigNumber(19)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(20)).toNumber()] = false;
      _nw332[(new BigNumber(21)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(22)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(23)).toNumber()] = _this.f16;
      _nw332[(new BigNumber(24)).toNumber()] = _this.f16;
      _1858_v3 = _nw332;
      let _1859_v4;
      _1859_v4 = _module.D1.create_DC5(new BigNumber(520), false);
      let _index325 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length));
      (_1858_v3)[_index325] = ((_this.f16) ? (_this.f16) : ((_1859_v4).dtor_cf11));
      let _1860_v5;
      _1860_v5 = _module.D0.create_DC2(new BigNumber(184), new BigNumber((_1857_v2).cardinality()), _1855_v0, _1858_v3);
      let _1861_v6;
      _1861_v6 = _dafny.Seq.UnicodeFromString("bv");
      let _index326 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length));
      let _rhs360 = (_1860_v5).dtor_cf2;
      let _rhs361 = ((_1855_v0).minus(_1855_v0)).isEqualTo(new BigNumber((_1861_v6).length));
      let _lhs279 = _1858_v3;
      let _lhs280 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length));
      _1855_v0 = _rhs360;
      _lhs279[_lhs280] = _rhs361;
      let _1862_i0;
      _1862_i0 = _dafny.ZERO;
      L9: {
        while (_this.f16) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1862_i0)) {
              break L9;
            }
            _1862_i0 = (_1862_i0).plus(_dafny.ONE);
            let _1863_v7;
            _1863_v7 = new _dafny.CodePoint('r'.codePointAt(0));
            let _1864_v8;
            _1864_v8 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pdxlfemr"),_1863_v7);
            _1864_v8 = (_1864_v8).update(_1861_v6, _1863_v7);
            _1858_v3 = _1858_v3;
            if ((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))]) {
              let _index327 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length));
              ((_this).f20)[_index327] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(533)), ((_1865_v0) => function (_1866_i1) {
                return _1865_v0;
              })(_1855_v0))).length);
              let _1867_v9;
              let _init53 = ((_1868_v6) => function (_1869_i2) {
                return _1868_v6;
              })(_1861_v6);
              let _nw333 = Array((new BigNumber(12)).toNumber());
              for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw333.length); _i0_53++) {
                _nw333[_i0_53] = _init53(new BigNumber(_i0_53));
              }
              _1867_v9 = _nw333;
              let _index328 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_1867_v9).length));
              (_1867_v9)[_index328] = _dafny.Seq.UnicodeFromString("xmofc");
              let _index329 = _module.__default.safeIndex(new BigNumber(832), new BigNumber(((_this).f20).length));
              ((_this).f20)[_index329] = _1855_v0;
              let _1870_v10;
              _1870_v10 = _dafny.Seq.of(_1861_v6, _1861_v6);
              let _1871_v11;
              _1871_v11 = _dafny.Set.fromElements(_this.f16, (_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))], _this.f16);
              let _index330 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length));
              let _index331 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_1867_v9).length));
              let _index332 = _module.__default.safeIndex(new BigNumber(832), new BigNumber(((_this).f20).length));
              let _rhs362 = _1855_v0;
              let _rhs363 = _dafny.Seq.Concat((_1870_v10)[_module.__default.safeIndex(_1855_v0, new BigNumber((_1870_v10).length))], _1861_v6);
              let _rhs364 = new BigNumber((_1871_v11).length);
              let _rhs365 = (!((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))]) || ((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))])) && ((_1855_v0).isLessThanOrEqualTo(_1855_v0));
              let _rhs366 = (_1855_v0).minus(((_dafny.ZERO).minus(_1855_v0)).multipliedBy(_1855_v0));
              let _lhs281 = (_this).f20;
              let _lhs282 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length));
              let _lhs283 = _1867_v9;
              let _lhs284 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_1867_v9).length));
              let _lhs285 = globalState;
              let _lhs286 = (_this).f20;
              let _lhs287 = _module.__default.safeIndex(new BigNumber(832), new BigNumber(((_this).f20).length));
              _lhs281[_lhs282] = _rhs362;
              _lhs283[_lhs284] = _rhs363;
              _1855_v0 = _rhs364;
              _lhs285.f2 = _rhs365;
              _lhs286[_lhs287] = _rhs366;
              _1855_v0 = ((((_this.f16) ? (true) : ((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))]))) ? ((((_1857_v2).contains(((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))])) ? ((_1857_v2).get(((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))])) : (((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))]))) : (_1855_v0));
              let _1872_v12;
              let _nw334 = new _module.C0();
              _nw334.__ctor(new BigNumber(742), _1858_v3);
              _1872_v12 = _nw334;
              let _1873_v13;
              _1873_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1872_v12,((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))]);
              let _1874_v14;
              let _nw335 = Array((new BigNumber(21)).toNumber());
              _nw335[(_dafny.ZERO).toNumber()] = ((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))];
              _nw335[(_dafny.ONE).toNumber()] = ((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))];
              _nw335[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-983)), function (_1875_i3) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              })).length));
              _nw335[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))]);
              _nw335[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.of((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))])).length);
              _nw335[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_1855_v0)).cardinality());
              _nw335[(new BigNumber(6)).toNumber()] = ((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))];
              _nw335[(new BigNumber(7)).toNumber()] = _1855_v0;
              _nw335[(new BigNumber(8)).toNumber()] = ((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))];
              _nw335[(new BigNumber(9)).toNumber()] = new BigNumber(420);
              _nw335[(new BigNumber(10)).toNumber()] = new BigNumber((_1873_v13).length);
              _nw335[(new BigNumber(11)).toNumber()] = _1855_v0;
              _nw335[(new BigNumber(12)).toNumber()] = _1855_v0;
              _nw335[(new BigNumber(13)).toNumber()] = ((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))];
              _nw335[(new BigNumber(14)).toNumber()] = (_1872_v12).f26;
              _nw335[(new BigNumber(15)).toNumber()] = _1855_v0;
              _nw335[(new BigNumber(16)).toNumber()] = new BigNumber(-696);
              _nw335[(new BigNumber(17)).toNumber()] = _module.__default.fm3(globalState);
              _nw335[(new BigNumber(18)).toNumber()] = _1855_v0;
              _nw335[(new BigNumber(19)).toNumber()] = (_1872_v12).f26;
              _nw335[(new BigNumber(20)).toNumber()] = ((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))];
              _1874_v14 = _nw335;
              let _1876_v15;
              _1876_v15 = _dafny.Map.Empty.slice().updateUnsafe(true,_1874_v14);
              _1876_v15 = (_1876_v15).update(_this.f16, _1874_v14);
              (globalState).f5 = (_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))];
              let _1877_v17;
              _1877_v17 = _dafny.Set.fromElements((_1872_v12).f26, (_1872_v12).f26, new BigNumber((function () {
                let _coll66 = new _dafny.Set();
                for (const _compr_66 of _dafny.IntegerRange(new BigNumber(-397), new BigNumber(-384))) {
                  let _1878_v16 = _compr_66;
                  if (((new BigNumber(-397)).isLessThanOrEqualTo(_1878_v16)) && ((_1878_v16).isLessThan(new BigNumber(-384)))) {
                    _coll66.add(_module.__default.safeModuloInt(_1878_v16, (_1872_v12).f26));
                  }
                }
                return _coll66;
              }()).length));
              let _1879_v18;
              _1879_v18 = _module.D9.create_DC26(new BigNumber(((_1870_v10)[_module.__default.safeIndex((_1872_v12).f26, new BigNumber((_1870_v10).length))]).length), _1872_v12);
              let _index333 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length));
              let _index334 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length));
              let _rhs367 = (((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))]) ? (_1872_v12) : ((_1879_v18).dtor_cf46));
              let _rhs368 = (_1872_v12).f26;
              let _rhs369 = (_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))];
              let _rhs370 = (_1877_v17).Difference(_dafny.Set.fromElements(((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))], ((_this).f20)[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length))], _1855_v0, (_1872_v12).f26, _1855_v0));
              let _lhs288 = (_this).f20;
              let _lhs289 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_this).f20).length));
              let _lhs290 = _1858_v3;
              let _lhs291 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length));
              _1872_v12 = _rhs367;
              _lhs288[_lhs289] = _rhs368;
              _lhs290[_lhs291] = _rhs369;
              _1877_v17 = _rhs370;
            } else {
              _1861_v6 = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm2(globalState), _1863_v7), _dafny.Seq.Concat(_1861_v6, _1861_v6));
              let _1880_v19;
              _1880_v19 = _dafny.Seq.of(_1861_v6);
              let _1881_v20;
              _1881_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1880_v19,_module.__default.safeModuloInt(_1855_v0, _1855_v0));
              _1881_v20 = (_1881_v20).update(_1880_v19, _1855_v0);
              _1861_v6 = ((true) ? (_dafny.Seq.UnicodeFromString("yerr")) : (_1861_v6));
              let _1882_v21;
              _1882_v21 = _dafny.Map.Empty.slice().updateUnsafe((((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))]) ? ((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))]) : (_this.f16)),(new BigNumber((_dafny.Set.fromElements((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))], (_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))])).length)).minus(new BigNumber(37)));
              _1882_v21 = (_1882_v21).update(!(_this.f16), _module.__default.fm3(globalState));
              (globalState).f5 = (_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))];
            }
            let _1883_v22;
            _1883_v22 = _dafny.Set.fromElements(_1855_v0, _1855_v0, (_dafny.ZERO).minus(_1855_v0));
            let _1884_v23;
            _1884_v23 = _module.D9.create_DC27(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("yxw"), _1861_v6), _1883_v22);
            _1884_v23 = _1884_v23;
          }
        }
      }
      let _1885_v24;
      let _nw336 = new _module.C2();
      _nw336.__ctor((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1855_v0, _1855_v0)));
      _1885_v24 = _nw336;
      _1858_v3 = _1858_v3;
      let _hi17 = _1855_v0;
      for (let _1886_i4 = new BigNumber(676); _1886_i4.isLessThan(_hi17); _1886_i4 = _1886_i4.plus(_dafny.ONE)) {
        _1855_v0 = _module.__default.safeModuloInt(_1886_i4, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1885_v24).f25,_1886_i4)).length));
        let _1887_v25;
        _1887_v25 = _dafny.MultiSet.fromElements(_this.f16);
        let _1888_v26;
        _1888_v26 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_1858_v3)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1858_v3).length))]), _1887_v25);
        _1887_v25 = (_1888_v26)[_module.__default.safeIndex((_1885_v24).f25, new BigNumber((_1888_v26).length))];
        let _1889_v27;
        _1889_v27 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1890_v28;
        _1890_v28 = _dafny.Set.fromElements(_1889_v27);
        _1890_v28 = _module.__default.fm27(globalState);
        let _1891_v29;
        let _nw337 = new _module.C6();
        _nw337.__ctor(_1860_v5);
        _1891_v29 = _nw337;
      }
      _1855_v0 = new BigNumber((_dafny.Seq.Concat((_this).f17, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-893)), function (_1892_i5) {
        return new BigNumber(109);
      }), (_this).f17))).length);
      return;
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f16 = false;
      this._f17 = _dafny.Seq.of();
      this.f19 = _dafny.Set.Empty;
      this._f18 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
    set f16(value) {
      let _this = this;
      _this._f16 = value;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    __ctor(f18, f19, f16, f17) {
      let _this = this;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return _this.f16;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let _hi18 = p3;
      for (let _1893_i0 = p3; _1893_i0.isLessThan(_hi18); _1893_i0 = _1893_i0.plus(_dafny.ONE)) {
        let _1894_v0;
        _1894_v0 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        _1894_v0 = (_1894_v0).update(_1893_i0, !(p1));
        (globalState).f2 = p2;
        let _1895_v1;
        _1895_v1 = new BigNumber(960);
        _1895_v1 = _1893_i0;
        let _1896_v2;
        let _nw338 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _1896_v2 = _nw338;
        let _1897_v3;
        _1897_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_1896_v2);
        _1897_v3 = (_1897_v3).update(_this.f16, _1896_v2);
      }
      let _hi19 = p3;
      for (let _1898_i1 = p3; _1898_i1.isLessThan(_hi19); _1898_i1 = _1898_i1.plus(_dafny.ONE)) {
        let _1899_v4;
        _1899_v4 = new BigNumber(284);
        let _1900_v5;
        let _nw339 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1900_v5 = _nw339;
        let _1901_v6;
        let _init54 = ((_1902_p2) => function (_1903_i2) {
          return _1902_p2;
        })(p2);
        let _nw340 = Array((new BigNumber(20)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw340.length); _i0_54++) {
          _nw340[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1901_v6 = _nw340;
        let _index335 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1901_v6).length));
        (_1901_v6)[_index335] = (_dafny.Set.fromElements(p1, p2)).IsDisjointFrom(_this.f19);
        let _1904_v7;
        _1904_v7 = _module.D0.create_DC0(_module.__default.fm1(globalState));
        let _1905_v8;
        _1905_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1904_v7);
        let _index336 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1901_v6).length));
        let _rhs371 = new BigNumber((_1905_v8).length);
        let _rhs372 = _1900_v5;
        let _rhs373 = !(p2);
        let _lhs292 = _1901_v6;
        let _lhs293 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_1901_v6).length));
        _1899_v4 = _rhs371;
        _1900_v5 = _rhs372;
        _lhs292[_lhs293] = _rhs373;
        _1899_v4 = new BigNumber(-278);
        let _1906_v9;
        let _nw341 = new _module.C6();
        _nw341.__ctor(_module.D0.create_DC2(((_this).f17)[_module.__default.safeIndex(p3, new BigNumber(((_this).f17).length))], p3, _1899_v4, _1901_v6));
        _1906_v9 = _nw341;
        let _1907_v10;
        _1907_v10 = _dafny.Seq.of(_1901_v6);
        (globalState).f5 = _dafny.Seq.contains(_1907_v10, _1901_v6);
      }
      let _1908_v11;
      let _nw342 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _1908_v11 = _nw342;
      let _1909_v12;
      _1909_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1908_v11);
      let _hi20 = (_dafny.ZERO).minus(new BigNumber(((_1909_v12).update(_this.f16, _1908_v11)).length));
      for (let _1910_i3 = p3; _1910_i3.isLessThan(_hi20); _1910_i3 = _1910_i3.plus(_dafny.ONE)) {
        let _1911_v13;
        _1911_v13 = _dafny.Seq.of(p3);
        _1911_v13 = (_this).f17;
        (globalState).f5 = true;
        let _1912_v14;
        let _nw343 = new _module.C1();
        _nw343.__ctor(new BigNumber((_this.f19).length), p3, p1, (_this).f17, p3);
        _1912_v14 = _nw343;
        let _1913_v15;
        _1913_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f19).length),_1912_v14);
        _1913_v15 = (_1913_v15).update(p3, _1912_v14);
        if (_this.f16) {
          let _1914_v16;
          _1914_v16 = _module.D3.create_DC10(true);
          let _1915_v17;
          _1915_v17 = _dafny.Seq.of(_1912_v14.f16, _1912_v14.f16, (_1914_v16).dtor_cf19);
          r0 = _dafny.Seq.update(_1915_v17, _module.__default.safeIndex((p3).minus(new BigNumber(835)), new BigNumber((_1915_v17).length)), p2);
          let _1916_v18;
          let _nw344 = Array((new BigNumber(20)).toNumber()).fill(false);
          _1916_v18 = _nw344;
          let _1917_v19;
          _1917_v19 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1918_v20;
          _1918_v20 = _dafny.MultiSet.fromElements(_1917_v19);
          let _index337 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1916_v18).length));
          (_1916_v18)[_index337] = (_dafny.MultiSet.fromElements(_1917_v19)).equals(_1918_v20);
          let _1919_v21;
          _1919_v21 = _dafny.Seq.of(_dafny.Seq.update(p0, _module.__default.safeIndex((_1912_v14).f28, new BigNumber((p0).length)), _1917_v19), (_this).f18, p0);
          let _1920_v22;
          _1920_v22 = _module.D0.create_DC1(p0);
          let _1921_v23;
          _1921_v23 = _dafny.Set.fromElements((_this).f18, p0, (_1919_v21)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber((_dafny.Set.fromElements((_1912_v14).f28)).length))).length)), new BigNumber((_1919_v21).length))], _dafny.Seq.UnicodeFromString("apnbo"), _dafny.Seq.update(_dafny.Seq.Concat((_this).f18, (_1920_v22).dtor_cf1), _module.__default.safeIndex((_1912_v14).f28, new BigNumber((_dafny.Seq.Concat((_this).f18, (_1920_v22).dtor_cf1)).length)), _1917_v19));
          let _1922_v24;
          _1922_v24 = new BigNumber(320);
          let _index338 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1916_v18).length));
          let _rhs374 = _1912_v14.f16;
          let _rhs375 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("ktslxunay"), (_this).f18);
          let _rhs376 = _1921_v23;
          let _rhs377 = _dafny.Seq.Concat(_1915_v17, _1915_v17);
          let _rhs378 = _1910_i3;
          let _lhs294 = _1916_v18;
          let _lhs295 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1916_v18).length));
          let _lhs296 = globalState;
          _lhs294[_lhs295] = _rhs374;
          _lhs296.f2 = _rhs375;
          _1921_v23 = _rhs376;
          _1915_v17 = _rhs377;
          _1922_v24 = _rhs378;
          _1911_v13 = (_1912_v14).f17;
          let _1923_v25;
          _1923_v25 = _dafny.Seq.UnicodeFromString("jco");
          let _1924_v26;
          _1924_v26 = _dafny.Set.fromElements(_1908_v11);
          let _1925_v27;
          _1925_v27 = _dafny.MultiSet.fromElements(p2);
          let _1926_v28;
          _1926_v28 = _dafny.MultiSet.fromElements(new BigNumber(((_1912_v14).f17).length), new BigNumber((_1925_v27).cardinality()), _1910_i3);
          let _rhs379 = _dafny.Seq.UnicodeFromString("xfcilcoa");
          let _rhs380 = _1924_v26;
          let _rhs381 = ((_dafny.MultiSet.fromElements(_1922_v24)).update(new BigNumber(-137), _module.__default.abs((_dafny.ZERO).minus(p3)))).equals(_1926_v28);
          let _lhs297 = globalState;
          _1923_v25 = _rhs379;
          _1924_v26 = _rhs380;
          _lhs297.f2 = _rhs381;
          let _1927_v29;
          _1927_v29 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).fm5(globalState));
          _1927_v29 = (_1927_v29).update(p2, true);
        } else {
          let _1928_v30;
          _1928_v30 = new _dafny.CodePoint('b'.codePointAt(0));
          (globalState).f15 = _1928_v30;
          let _1929_v31;
          _1929_v31 = new BigNumber(418);
          let _rhs382 = (_1912_v14).f28;
          let _rhs383 = _module.__default.fm3(globalState);
          _1929_v31 = _rhs382;
          _1929_v31 = _rhs383;
          let _1930_v32;
          let _nw345 = Array((new BigNumber(13)).toNumber()).fill(false);
          _1930_v32 = _nw345;
          let _index339 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_1930_v32).length));
          (_1930_v32)[_index339] = ((_1912_v14).f28).isLessThan((_1912_v14).f28);
          let _index340 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_1930_v32).length));
          (_1930_v32)[_index340] = p2;
          let _rhs384 = (new BigNumber(((_this).f18).length)).minus((p3).multipliedBy(_1910_i3));
          let _rhs385 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("lvfjtwdb")).length), ((_1912_v14).f28).plus(_1910_i3));
          let _rhs386 = (!(_1912_v14.f16)) === ((p1) && (!((_1930_v32)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_1930_v32).length))])));
          let _rhs387 = (_dafny.ZERO).minus(_1929_v31);
          let _lhs298 = globalState;
          _1929_v31 = _rhs384;
          _1929_v31 = _rhs385;
          _lhs298.f2 = _rhs386;
          _1929_v31 = _rhs387;
          let _1931_v33;
          let _nw346 = new _module.C3();
          _nw346.__ctor((_this).f18, _1912_v14.f16, (_1912_v14).f17);
          _1931_v33 = _nw346;
        }
      }
      if (_this.f16) {
        let _1932_v34;
        _1932_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1908_v11,p0);
        let _1933_v35;
        _1933_v35 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _1934_v36;
        _1934_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,(((_1933_v35).contains(p1)) ? ((_1933_v35).get(p1)) : (p0)));
        _1932_v34 = (_1932_v34).update(_1908_v11, (((_1933_v35).contains(p2)) ? ((_1933_v35).get(p2)) : (p0)));
        let _1935_v37;
        _1935_v37 = _dafny.Seq.of(_this.f16);
        let _1936_v38;
        _1936_v38 = _dafny.Set.fromElements(new BigNumber((_1935_v37).length), (_dafny.ZERO).minus(p3), new BigNumber(-403));
        _1936_v38 = (_1936_v38).Union(_module.__default.fm31(p2, globalState));
        let _1937_v39;
        let _init55 = function (_1938_i4) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        };
        let _nw347 = Array((new BigNumber(20)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw347.length); _i0_55++) {
          _nw347[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _1937_v39 = _nw347;
        let _1939_v40;
        _1939_v40 = _dafny.Seq.of(_1937_v39, _1937_v39, _1937_v39, _1937_v39);
        _1939_v40 = _dafny.Seq.Concat(_1939_v40, _1939_v40);
        let _1940_v41;
        let _nw348 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1940_v41 = _nw348;
        let _1941_v42;
        _1941_v42 = _module.D11.create_DC31(p1, new BigNumber(919), p3, p1);
        let _1942_v43;
        _1942_v43 = _dafny.Map.Empty.slice().updateUnsafe(_module.D11.create_DC32(_1941_v42),_1940_v41);
        let _1943_v44;
        _1943_v44 = _module.D11.create_DC32(_1941_v42);
        let _1944_v45;
        _1944_v45 = _dafny.MultiSet.fromElements(new BigNumber(762));
        let _1945_v46;
        _1945_v46 = _module.D0.create_DC2(new BigNumber((_1944_v45).cardinality()), p3, p3, _1940_v41);
        let _rhs388 = (((_1942_v43).contains(_1943_v44)) ? ((_1942_v43).get(_1943_v44)) : ((_1945_v46).dtor_cf5));
        let _rhs389 = _1908_v11;
        _1940_v41 = _rhs388;
        _1908_v11 = _rhs389;
        let _1946_v47;
        _1946_v47 = new BigNumber(-83);
        let _1947_v48;
        _1947_v48 = _module.D11.create_DC29(_module.__default.fm2(globalState));
        let _pat_let_tv39 = _1941_v42;
        let _rhs390 = _1946_v47;
        let _rhs391 = _1947_v48;
        let _rhs392 = _1946_v47;
        let _rhs393 = function (_pat_let51_0) {
          return function (_1948_dt__update__tmp_h0) {
            return function (_pat_let52_0) {
              return function (_1949_dt__update_hcf57_h0) {
                return _module.D11.create_DC32(_1949_dt__update_hcf57_h0);
              }(_pat_let52_0);
            }(_pat_let_tv39);
          }(_pat_let51_0);
        }(_1943_v44);
        let _rhs394 = true;
        let _lhs299 = _this;
        _1946_v47 = _rhs390;
        _1947_v48 = _rhs391;
        _1946_v47 = _rhs392;
        _1943_v44 = _rhs393;
        _lhs299.f16 = _rhs394;
      } else {
        if (p2) {
          let _1950_v49;
          _1950_v49 = _dafny.Seq.UnicodeFromString("ebxvgp");
          let _1951_v50;
          _1951_v50 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1950_v49);
          let _1952_v51;
          _1952_v51 = _dafny.Seq.of(_1950_v49);
          let _1953_v52;
          _1953_v52 = _dafny.MultiSet.fromElements(p1);
          _1950_v49 = (((_1951_v50).contains(_module.__default.safeModuloInt(p3, _module.__default.fm3(globalState)))) ? ((_1951_v50).get(_module.__default.safeModuloInt(p3, _module.__default.fm3(globalState)))) : (((_this.f16) ? ((_1952_v51)[_module.__default.safeIndex(new BigNumber((_1953_v52).cardinality()), new BigNumber((_1952_v51).length))]) : ((_1952_v51)[_module.__default.safeIndex(p3, new BigNumber((_1952_v51).length))]))));
          let _1954_v53;
          let _nw349 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1954_v53 = _nw349;
          let _index341 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1954_v53).length));
          (_1954_v53)[_index341] = (false) && (p2);
          let _1955_v54;
          _1955_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,p1);
          let _1956_v55;
          _1956_v55 = _module.D1.create_DC5(new BigNumber(((_1955_v54).update(p2, p1)).length), p2);
          let _1957_v56;
          _1957_v56 = _dafny.Seq.of(_this.f16, _module.__default.fm0(_this.f16, _this.f16, globalState), _this.f16, p2);
          let _1958_v57;
          let _nw350 = new _module.C5();
          _nw350.__ctor(_this.f16, (_this).f17);
          _1958_v57 = _nw350;
          let _1959_v58;
          _1959_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1958_v57,p3);
          let _index342 = _module.__default.safeIndex(new BigNumber(819), new BigNumber((_1954_v53).length));
          (_1954_v53)[_index342] = (((_1956_v55).dtor_cf11) ? ((_1957_v56)[_module.__default.safeIndex((((_1959_v58).contains(_1958_v57)) ? ((_1959_v58).get(_1958_v57)) : (p3)), new BigNumber((_1957_v56).length))]) : (p1));
          let _1960_v59;
          _1960_v59 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(((_this).f18).length));
          let _1961_v60;
          let _nw351 = Array((new BigNumber(20)).toNumber());
          _nw351[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(p3, (_dafny.ZERO).minus(p3), p3, p3, p3);
          _nw351[(_dafny.ONE).toNumber()] = _dafny.Seq.of(new BigNumber(-740));
          _nw351[(new BigNumber(2)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(3)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(4)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(5)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(6)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(7)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(8)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(p3, p3);
          _nw351[(new BigNumber(10)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(11)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(12)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(p3);
          _nw351[(new BigNumber(14)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(15)).toNumber()] = _dafny.Seq.update((_this).f17, _module.__default.safeIndex(new BigNumber(474), new BigNumber(((_this).f17).length)), p3);
          _nw351[(new BigNumber(16)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(17)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(18)).toNumber()] = (_this).f17;
          _nw351[(new BigNumber(19)).toNumber()] = (_this).f17;
          _1961_v60 = _nw351;
          let _1962_v61;
          _1962_v61 = _module.D2.create_DC8(_1960_v59, _1961_v60, p1, _this.f16, p3);
          let _1963_v62;
          _1963_v62 = _module.D6.create_DC17(_1962_v61);
          _1963_v62 = _1963_v62;
          let _1964_v63;
          _1964_v63 = _dafny.Set.fromElements(p3);
          let _1965_v64;
          _1965_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1964_v63,_1908_v11);
          _1965_v64 = (_1965_v64).update(function () {
            let _coll67 = new _dafny.Set();
            for (const _compr_67 of (_1964_v63).Elements) {
              let _1966_v65 = _compr_67;
              if ((_1964_v63).contains(_1966_v65)) {
                _coll67.add(_module.__default.safeDivisionInt(_1966_v65, new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(439),!(true)))).cardinality()),false))).length)));
              }
            }
            return _coll67;
          }(), _1908_v11);
          let _1967_v66;
          let _nw352 = Array((_dafny.ONE).toNumber());
          _nw352[(_dafny.ZERO).toNumber()] = _this.f19;
          _1967_v66 = _nw352;
          _1967_v66 = _1967_v66;
        } else {
          (globalState).f2 = _dafny.Seq.IsProperPrefixOf(p0, (_this).f18);
          let _1968_v67;
          _1968_v67 = _module.D11.create_DC31(_this.f16, p3, p3, true);
          _1968_v67 = _module.D11.create_DC31(_this.f16, p3, p3, _this.f16);
          let _1969_v68;
          _1969_v68 = new BigNumber(248);
          _1969_v68 = _1969_v68;
          _1969_v68 = ((!(_this.f16)) ? (_module.__default.fm3(globalState)) : (new BigNumber(6)));
          (globalState).f15 = ((_this).f18)[_module.__default.safeIndex(_1969_v68, new BigNumber(((_this).f18).length))];
        }
        let _1970_v69;
        let _init56 = function (_1971_i5) {
          return false;
        };
        let _nw353 = Array((new BigNumber(15)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw353.length); _i0_56++) {
          _nw353[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _1970_v69 = _nw353;
        let _index343 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1970_v69).length));
        (_1970_v69)[_index343] = _module.__default.fm0(p2, p2, globalState);
        let _index344 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1970_v69).length));
        (_1970_v69)[_index344] = p2;
        let _1972_v70;
        _1972_v70 = new BigNumber(-337);
        let _rhs395 = ((_dafny.ZERO).minus(p3)).multipliedBy(_1972_v70);
        let _rhs396 = _module.__default.safeDivisionInt(new BigNumber(15), p3);
        _1972_v70 = _rhs395;
        _1972_v70 = _rhs396;
        let _1973_v71;
        _1973_v71 = _dafny.Seq.UnicodeFromString("dlvn");
        _1973_v71 = _dafny.Seq.Concat((_this).f18, (_this).f18);
        _1972_v70 = (new BigNumber(((_this).f17).length)).plus(p3);
      }
      for (const _guard_loop_9 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1908_v11).length))) {
        let _1974_i6 = _guard_loop_9;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1974_i6)) && ((_1974_i6).isLessThan(new BigNumber((_1908_v11).length))))) {
          (_1908_v11)[(_1974_i6)] = (_1974_i6).multipliedBy(p3);
        }
      }
      let _1975_v72;
      _1975_v72 = new BigNumber(-430);
      let _1976_v73;
      _1976_v73 = _dafny.Seq.UnicodeFromString("dmsoc");
      let _rhs397 = ((_this.f16) ? (_1908_v11) : (_1908_v11));
      let _rhs398 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm20(_this.f16, !(p2), true, _1975_v72, globalState), _dafny.Seq.of(p1, _this.f16))).length);
      let _rhs399 = p2;
      let _rhs400 = !(_this.f16);
      let _rhs401 = p0;
      let _lhs300 = globalState;
      let _lhs301 = globalState;
      _1908_v11 = _rhs397;
      _1975_v72 = _rhs398;
      _lhs300.f5 = _rhs399;
      _lhs301.f5 = _rhs400;
      _1976_v73 = _rhs401;
      let _1977_v74;
      _1977_v74 = _dafny.Seq.of(p2, _dafny.Seq.IsPrefixOf(p0, _dafny.Seq.UnicodeFromString("c")), p1);
      r0 = _1977_v74;
      return r0;
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
